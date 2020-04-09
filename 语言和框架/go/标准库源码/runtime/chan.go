// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// This file contains the implementation of Go channels.

// Invariants:
//  At least one of c.sendq and c.recvq is empty,
//  except for the case of an unbuffered channel with a single goroutine
//  blocked on it for both sending and receiving using a select statement,
//  in which case the length of c.sendq and c.recvq is limited only by the
//  size of the select statement.
//
// 在 hchan 对象中, c.sendq 和 c. recvq 至少有一个为空.
// 不过有一种情况除外.
// 
// 
// 
// For buffered channels, also:
//  c.qcount > 0 implies that c.recvq is empty.
//  c.qcount < c.dataqsiz implies that c.sendq is empty.
// 对于缓冲channel
// 当 c.qcount > 0 表示 c.recvq 为空
// recvq 中是等待读取的协程, channel 中还有数据留存说明没有多余的协程了.
// ...不过应该也存在 recvq 中同时出现 n 个协程来读取, 挨个读取也是需要时间的吧.
// 当 c.qcount < c.dataqsiz 表示 c.sendq 为空().
// sendq 中是等待写入的协程, 在 channel 未满的情况下, 不过有额外的写协程存在的.
// 当然, 与上面一样, 同时来 n 个协程写入, 挨个写入也是需要时间的啊.
// ...总感觉这样的说法不够严谨.

import (
	"runtime/internal/atomic"
	"runtime/internal/math"
	"unsafe"
)

const (
	maxAlign  = 8
	// 这是空的 hchan 结构体占用的空间大小吧
	hchanSize = unsafe.Sizeof(hchan{}) + uintptr(-int(unsafe.Sizeof(hchan{}))&(maxAlign-1))
	debugChan = false
)

type hchan struct {
	// 队列里面的数据总量
	// total data in the queue
	qcount uint 
	// 循环队列的容量, 与qcount相比, 应该是cap与len的区别
	// 注意, channel对象无法像slice一样扩容, 所以初始化后这个值是无法改变的
	// size of the circular queue
	dataqsiz uint
	// points to an array of dataqsiz elements
	// 实际存储数据的地方, 这其实是一个数组对象(首地址)
	buf      unsafe.Pointer 
	// 队列中的元素类型占用的字节数(即单个元素的大小)
	elemsize uint16
	closed   uint32
	// 队列中的元素的类型
	// element type
	elemtype *_type 
	// sendx/recvx为buf内的元素索引, 最大值为dataqsiz.
	// 当sendx增加至dataqsiz时, 就会被重置为0, 以此表示循环队列.
	sendx    uint   // send index
	recvx    uint   // receive index
	// 等待读取channel的G队列(其实是一个链表)
	// 这个链表如果不为空, 则说明 channel 中没有数据, 大家都在等.
	// list of recv waiters
	recvq waitq 
	// 等待向channel写数据的G队列
	// list of send waiters
	sendq waitq 

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	// 全局锁, 保护 hchan 结构中的所有字段.
	lock mutex
}

// sudog结构是对 G 对象的封装.
// 空 channel 下的读协程会挂在 recvq 队列中
// 满 channel 下的写协程会挂在 sendq 队列中
// waitq 拥有 enqueue 和 dequque 两个方法
type waitq struct {
	first *sudog
	last  *sudog
}

//go:linkname reflect_makechan reflect.makechan
func reflect_makechan(t *chantype, size int) *hchan {
	return makechan(t, size)
}

/*
type chantype struct {
	typ  _type
	elem *_type
	dir  uintptr
}
*/

// 应用层代码中通过make(chan XXX)创建channel对象,
// 在编译时编译器会重写然后执行这里的makechan(64)函数.
// 实际的流程在 makechan().
// @param t: channel里面保存的元素的数据类型
// @param size: 缓冲的容量(如果为0表示是非缓冲buffer)
func makechan64(t *chantype, size int64) *hchan {
	if int64(int(size)) != size {
		panic(plainError("makechan: size out of range"))
	}

	return makechan(t, int(size))
}

func makechan(t *chantype, size int) *hchan {
	elem := t.elem

	// compiler checks this but be safe.
	if elem.size >= 1 << 16 {
		throw("makechan: invalid channel element type")
	}
	if hchanSize % maxAlign != 0 || elem.align > maxAlign {
		throw("makechan: bad alignment")
	}

	// mem = channel中元素对象占用空间大小 * 队列容量size,
	// 即为channel对象所需的整体内存大小.
	// maxAlloc: 单次 malloc 内存分配的上限
	mem, overflow := math.MulUintptr(elem.size, uintptr(size))
	if overflow || mem > maxAlloc-hchanSize || size < 0 {
		panic(plainError("makechan: size out of range"))
	}

	// 着手创建hchan对象
	// Hchan does not contain pointers interesting for GC 
	// when elements stored in buf do not contain pointers.
	// buf points into the same allocation, elemtype is persistent.
	// SudoG's are referenced from their owning thread so they can't be collected.
	// TODO(dvyukov,rlh): Rethink when collector can move allocated objects.
	// mallocgc() 参数分别为: 空间大小, 类型, 是否需要清零.
	var c *hchan
	switch {
	case mem == 0:
		// 创建无缓冲channel, 无缓冲channel是没有 buf 队列的
		// Queue or element size is zero.
		c = (*hchan)(mallocgc(hchanSize, nil, true))
		// Race detector uses this location for synchronization.
		// 使用mallocgc创建hchan对象后, buf为普通的Pointer对象, 没有额外空间.
		c.buf = c.raceaddr()
	case elem.kind & kindNoPointers != 0:
		// channel中的对象不为指针, 此时同时为hchan对象及其buf成员申请内存.
		// 此时, 整个hchan对象中各成员的内存是连续的.
		// Elements do not contain pointers.
		// Allocate hchan and buf in one call.
		c = (*hchan)(mallocgc(hchanSize+mem, nil, true))
		// 这里的 add() 还是搜一下吧, 在 src/cmd 目录下, 不知道不同版本是不是相同的.
		// 以go 1.12 为例, 原型如下
		// func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
		// 	return unsafe.Pointer(uintptr(p) + x)
		// }
		// 就是做了一个加法, 计算一下从 c 的地址开始, hchanSize 之后的位置.
		// 就是为实际存储数据的数组起始地址.
		c.buf = add(unsafe.Pointer(c), hchanSize)
	default:
		// channel中的对象为指针类型, 则先创建hchan对象, 再为buf成员申请内存
		// 此时hchan与其buf成员的指针不连续.
		// Elements contain pointers.
		c = new(hchan)
		c.buf = mallocgc(mem, elem, true)
	}
	// 初始化channel对象的关键属性, 元素对象类型和占用空间, 以及队列缓冲区大小
	c.elemsize = uint16(elem.size)
	c.elemtype = elem
	c.dataqsiz = uint(size)

	if debugChan {
		print("makechan: chan=", c, "; elemsize=", elem.size, "; elemalg=", elem.alg, "; dataqsiz=", size, "\n")
	}
	return c
}

// chanbuf(c, i) 返回channel对象 c.buf 成员中第 i 个元素的地址
// chanbuf(c, i) is pointer to the i'th slot in the buffer.
// caller: chansend()
func chanbuf(c *hchan, i uint) unsafe.Pointer {
	// add 函数用于完成指针运算
	return add(c.buf, uintptr(i)*uintptr(c.elemsize))
}

// 所有通过`c <- x`向channel写数据的代码, 编译器都会指向这个函数, 并最终调用`chansend()`
// entry point for c <- x from compiled code
//go:nosplit
func chansend1(c *hchan, elem unsafe.Pointer) {
	chansend(c, elem, true, getcallerpc())
}

/*
 * generic single channel send/recv
 * If block is not nil,
 * then the protocol will not
 * sleep but return if it could not complete.
 *
 * sleep can wake up with g.param == nil
 * when a channel involved in the sleep has been closed. 
 * it is easiest to loop and re-run the operation; 
 * we'll see that it's now closed.
 */
func chansend(c *hchan, ep unsafe.Pointer, block bool, callerpc uintptr) bool {
	// channel为空, 什么情况下channel会为空?
	// 使用var c chan int声明变量时, c为nil, 因为没有分配空间.
	// 读写一个nil的channel会阻塞, 使用close()关闭ta时会发生panic.
	if c == nil {
		if !block {
			return false
		}
		// gopark()会挂起当前goroutine, 可以通过ta的第一个参数unlockf唤醒.
		// 但这里unlockf为nil, 所以会一直休眠.
		// 貌似gopark本身不是阻塞的? 因为ta不阻塞throw抛出异常?
		gopark(nil, nil, waitReasonChanSendNilChan, traceEvGoStop, 2)
		throw("unreachable")
	}

	if debugChan {
		print("chansend: chan=", c, "\n")
	}

	if raceenabled {
		racereadpc(c.raceaddr(), callerpc, funcPC(chansend))
	}

	// Fast path: check for failed non-blocking operation without acquiring the lock.
	//
	// After observing that the channel is not closed, we observe that the channel is
	// not ready for sending. Each of these observations is a single word-sized read
	// (first c.closed and second c.recvq.first or c.qcount depending on kind of channel).
	// Because a closed channel cannot transition from 'ready for sending' to
	// 'not ready for sending', even if the channel is closed between the two observations,
	// they imply a moment between the two when the channel was both not yet closed
	// and not ready for sending. We behave as if we observed the channel at that moment,
	// and report that the send cannot proceed.
	//
	// It is okay if the reads are reordered here: if we observe that the channel is not
	// ready for sending and then observe that it is not closed, that implies that the
	// channel wasn't closed during the first observation.
	if !block && c.closed == 0 && ((c.dataqsiz == 0 && c.recvq.first == nil) ||
		(c.dataqsiz > 0 && c.qcount == c.dataqsiz)) {
		return false
	}

	var t0 int64
	if blockprofilerate > 0 {
		t0 = cputicks()
	}

	lock(&c.lock)
	// 向一个已经关闭的channel写数据, 这里会抛出panic.
	if c.closed != 0 {
		unlock(&c.lock)
		panic(plainError("send on closed channel"))
	}
	// recvq队列非空, 表示有goroutine等待接收数据. 此时跳过channel缓存, 直接将数据发给recevier goroutine.
	if sg := c.recvq.dequeue(); sg != nil {
		// Found a waiting receiver. We pass the value we want to send
		// directly to the receiver, bypassing the channel buffer (if any).
		send(c, sg, ep, func() { unlock(&c.lock) }, 3)
		return true
	}
	// 运行到此处, 表示recvq队列为空, 且channel缓冲未满, 可以写入
	if c.qcount < c.dataqsiz {
		// Space is available in the channel buffer. Enqueue the element to send.
		// qp表示c.buf队列中可写入的位置(偏移)
		qp := chanbuf(c, c.sendx)
		// raceenabled是一个const值, 总为true
		if raceenabled {
			// 这里是对qp这个地址加锁???
			raceacquire(qp)
			racerelease(qp)
		}
		// 从ep拷贝一个c.elemtype类型的值到qp
		typedmemmove(c.elemtype, qp, ep)
		// sendx索引加1
		c.sendx++
		if c.sendx == c.dataqsiz {
			// 这是为了表示循环队列?
			c.sendx = 0
		}
		// channel的成员数量加1
		c.qcount++
		unlock(&c.lock)
		return true
	}

	if !block {
		unlock(&c.lock)
		return false
	}
	// 运行到此处, 表示缓存队列已满且接收消息队列recv为空,
	// 则将当前的goroutine加入到send队列
	// Block on the channel. Some receiver will complete our operation for us.
	gp := getg()
	mysg := acquireSudog()
	mysg.releasetime = 0
	if t0 != 0 {
		mysg.releasetime = -1
	}
	// No stack splits between assigning elem and enqueuing mysg
	// on gp.waiting where copystack can find it.
	mysg.elem = ep
	mysg.waitlink = nil
	mysg.g = gp
	mysg.isSelect = false
	mysg.c = c
	gp.waiting = mysg
	gp.param = nil
	c.sendq.enqueue(mysg)
	goparkunlock(&c.lock, waitReasonChanSend, traceEvGoBlockSend, 3)
	// Ensure the value being sent is kept alive until the receiver copies it out. 
	// The sudog has a pointer to the stack object, but sudogs aren't considered as roots of the stack tracer.
	KeepAlive(ep)

	// someone woke us up.
	if mysg != gp.waiting {
		throw("G waiting list is corrupted")
	}
	gp.waiting = nil
	if gp.param == nil {
		if c.closed == 0 {
			throw("chansend: spurious wakeup")
		}
		panic(plainError("send on closed channel"))
	}
	gp.param = nil
	if mysg.releasetime > 0 {
		blockevent(mysg.releasetime-t0, 2)
	}
	mysg.c = nil
	releaseSudog(mysg)
	return true
}

// send 发起一个对空缓冲的 channel 的 send 操作.
// 关于c必须为空这一情况, 由于 send() 函数是在当前channel拥有等待读取的进程被调用, 表示c中为空.
// 由发送方发送的ep会被拷贝到接收方sg, 之后sg会被唤醒(通过调用goready()函数).
// 调用此函数时, c必须为空且被加锁. 等到send()执行完成, 会调用unlockf()将其解锁.
// sg必须已经从c.recvq中通过dequeue()弹出.
// ep必须为非nil的值, 且指向堆内存或是调用者的栈空间.
// caller: chansend()
func send(c *hchan, sg *sudog, ep unsafe.Pointer, unlockf func(), skip int) {
	if raceenabled {
		if c.dataqsiz == 0 {
			// 如果是无缓冲channel
			racesync(c, sg)
		} else {
			// 对于有缓冲channel, 假装使用了buf, 但实际上是直接拷贝的.
			// Pretend we go through the buffer, even though we copy directly. 
			// Note that we need to increment the head/tail locations only when raceenabled.
			qp := chanbuf(c, c.recvx)
			raceacquire(qp)
			racerelease(qp)
			raceacquireg(sg.g, qp)
			racereleaseg(sg.g, qp)
			c.recvx++
			if c.recvx == c.dataqsiz {
				c.recvx = 0
			}
			c.sendx = c.recvx // c.sendx = (c.sendx+1) % c.dataqsiz
		}
	}
	// 执行数据拷贝, 不管有无缓冲, 这里都是真正进行数据交换的地方.
	if sg.elem != nil {
		sendDirect(c.elemtype, sg, ep)
		sg.elem = nil
	}
	// 唤醒gp, 即等待读取channel的协程对象
	gp := sg.g
	unlockf() // 解除c.lock全局锁
	gp.param = unsafe.Pointer(sg)
	if sg.releasetime != 0 {
		sg.releasetime = cputicks()
	}
	goready(gp, skip+1) 
}

// Sends and receives on unbuffered or empty-buffered channels are the
// only operations where one running goroutine writes to the stack of
// another running goroutine. The GC assumes that stack writes only
// happen when the goroutine is running and are only done by that
// goroutine. Using a write barrier is sufficient to make up for
// violating that assumption, but the write barrier has to work.
// typedmemmove will call bulkBarrierPreWrite, but the target bytes
// are not in the heap, so that will not help. We arrange to call
// memmove and typeBitsBulkBarrier instead.

func sendDirect(t *_type, sg *sudog, src unsafe.Pointer) {
	// src is on our stack, dst is a slot on another stack.

	// Once we read sg.elem out of sg, it will no longer
	// be updated if the destination's stack gets copied (shrunk).
	// So make sure that no preemption points can happen between read & use.
	dst := sg.elem
	typeBitsBulkBarrier(t, uintptr(dst), uintptr(src), t.size)
	// No need for cgo write barrier checks because dst is always
	// Go memory.
	memmove(dst, src, t.size)
}

func recvDirect(t *_type, sg *sudog, dst unsafe.Pointer) {
	// dst is on our stack or the heap, src is on another stack.
	// The channel is locked, so src will not move during this
	// operation.
	src := sg.elem
	typeBitsBulkBarrier(t, uintptr(dst), uintptr(src), t.size)
	memmove(dst, src, t.size)
}

// closechan 调用 close() 关闭 channel, 通过内部锁 c.lock 完成.
// 需要对 channel 的多种状态作出判断, 如未初始化的, 已被关闭过的等.
// 同时对正在等待读取或写入的协程解除联系, 并在最后唤醒由于等待而阻塞的协程,
// 此时如果有读协程, 则读协程读出为 channel 类型的默认值; 
// 如果有写协程, 则写协程会panic...
func closechan(c *hchan) {
	// 对一个空 channel 执行 close 会引发panic
	if c == nil {
		panic(plainError("close of nil channel"))
	}

	lock(&c.lock)
	// 如果目标 channel 已经被关闭, 多次调用close也会引发panic
	if c.closed != 0 {
		unlock(&c.lock)
		panic(plainError("close of closed channel"))
	}

	if raceenabled {
		callerpc := getcallerpc()
		racewritepc(c.raceaddr(), callerpc, funcPC(closechan))
		racerelease(c.raceaddr())
	}
	// 标记其为已关闭
	c.closed = 1

	var glist gList

	// 遍历 c.recvq 队列, 解除所有读协程的联系.
	for {
		sg := c.recvq.dequeue()
		if sg == nil {
			break
		}
		if sg.elem != nil {
			typedmemclr(c.elemtype, sg.elem)
			sg.elem = nil
		}
		if sg.releasetime != 0 {
			sg.releasetime = cputicks()
		}
		gp := sg.g
		gp.param = nil
		if raceenabled {
			raceacquireg(gp, c.raceaddr())
		}
		glist.push(gp)
	}

	// 遍历 c.recvq 队列, 解除所有写协程的联系.
	// 注意: 这会导致写协程的 panic
	for {
		sg := c.sendq.dequeue()
		if sg == nil {
			break
		}
		sg.elem = nil
		if sg.releasetime != 0 {
			sg.releasetime = cputicks()
		}
		gp := sg.g
		gp.param = nil
		if raceenabled {
			raceacquireg(gp, c.raceaddr())
		}
		glist.push(gp)
	}
	unlock(&c.lock)

	// Ready all Gs now that we've dropped the channel lock.
	// 解除 channel 的全局锁后, 唤醒所有读写协程
	for !glist.empty() {
		gp := glist.pop()
		gp.schedlink = 0
		goready(gp, 3)
	}
}

// 所有通过`<-`从channel中读取数据的代码, 最终都会调用`channel()`函数
// 其中直接通过 `<- c`不需要得到结果的操作, 编译器会将其指向`chanrecv1()`函数
// entry points for <- c from compiled code
//go:nosplit
func chanrecv1(c *hchan, elem unsafe.Pointer) {
	chanrecv(c, elem, true)
}

//go:nosplit
func chanrecv2(c *hchan, elem unsafe.Pointer) (received bool) {
	_, received = chanrecv(c, elem, true)
	return
}

// chanrecv receives on channel c and writes the received data to ep.
// ep may be nil, in which case received data is ignored.
// If block == false and no elements are available, returns (false, false).
// Otherwise, if c is closed, zeros *ep and returns (true, false).
// Otherwise, fills in *ep with an element and returns (true, true).

// chanrecv 从c中读取数据, 并将读到的数据写入ep.
// ep可以为nil, 表示不care接收到的结果(只把`<-`当成阻塞代码流程的一种手段)
// 如果block == false且channel中没有数据时, 则直接返回(false, false), 
// 否则, 如果c被关闭, 
// ep对象如果不为空, 则一定指向堆, 或是调用者的栈(如a <- chan, a为栈中的变量)
func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool) {
	// raceenabled: don't need to check ep, as it is always on the stack
	// or is new memory allocated by reflect.

	if debugChan {
		print("chanrecv: chan=", c, "\n")
	}

	// ...什么情况下 c 可能为 nil?
	if c == nil {
		if !block {
			return
		}
		gopark(nil, nil, waitReasonChanReceiveNilChan, traceEvGoStop, 2)
		throw("unreachable")
	}

	// Fast path: check for failed non-blocking operation without acquiring the lock.
	//
	// After observing that the channel is not ready for receiving, we observe that the
	// channel is not closed. Each of these observations is a single word-sized read
	// (first c.sendq.first or c.qcount, and second c.closed).
	// Because a channel cannot be reopened, the later observation of the channel
	// being not closed implies that it was also not closed at the moment of the
	// first observation. We behave as if we observed the channel at that moment
	// and report that the receive cannot proceed.
	//
	// The order of operations is important here: reversing the operations can lead to
	// incorrect behavior when racing with a close.
	if !block && (c.dataqsiz == 0 && c.sendq.first == nil ||
		c.dataqsiz > 0 && atomic.Loaduint(&c.qcount) == 0) &&
		atomic.Load(&c.closed) == 0 {
		return
	}

	var t0 int64
	if blockprofilerate > 0 {
		t0 = cputicks()
	}

	lock(&c.lock)
	// channel已经被关闭且channel缓冲中没有数据了
	if c.closed != 0 && c.qcount == 0 {
		if raceenabled {
			raceacquire(c.raceaddr())
		}
		unlock(&c.lock)
		if ep != nil {
			typedmemclr(c.elemtype, ep)
		}
		return true, false
	}
	// sendq队列非空, 表示有goroutine对象等待写数据. 那么跳过channel缓冲, 直接从写队列接收数据.
	if sg := c.sendq.dequeue(); sg != nil {
		// If buffer is size 0, receive value directly from sender. 
		// Otherwise, receive from head of queue and add sender's value to the tail of the queue 
		// (both map to the same buffer slot because the queue is full).
		recv(c, sg, ep, func() { unlock(&c.lock) }, 3)
		return true, true
	}
	// 运行到此处, 表示sendq队列为空, 但是channel缓冲队列存在成员.
	// 此时可以直接从缓冲队列接收数据.
	if c.qcount > 0 {
		// Receive directly from queue
		qp := chanbuf(c, c.recvx)
		if raceenabled {
			raceacquire(qp)
			racerelease(qp)
		}
		if ep != nil {
			typedmemmove(c.elemtype, ep, qp)
		}
		typedmemclr(c.elemtype, qp)
		c.recvx++
		if c.recvx == c.dataqsiz {
			c.recvx = 0
		}
		c.qcount--
		unlock(&c.lock)
		return true, true
	}

	if !block {
		unlock(&c.lock)
		return false, false
	}

	// 运行到此处, 表示sendq队列为空且缓冲队列也为空. 
	// 因此需要将当前goroutine加入c的recvq队列, 并将其阻塞.
	// no sender available: block on this channel.
	// getg()返回当前的G对象
	gp := getg()
	// acquireSudog()会返回一个sudog对象(recv/send队列(链表)里都是sudog对象)
	mysg := acquireSudog()
	mysg.releasetime = 0
	if t0 != 0 {
		mysg.releasetime = -1
	}
	// No stack splits between assigning elem and enqueuing mysg
	// on gp.waiting where copystack can find it.
	mysg.elem = ep
	mysg.waitlink = nil
	// 这里gp和mysg还相互引用...
	gp.waiting = mysg
	mysg.g = gp
	mysg.isSelect = false
	mysg.c = c
	gp.param = nil
	// 加入队列
	c.recvq.enqueue(mysg)
	// 让其休眠, 到此阻塞
	goparkunlock(&c.lock, waitReasonChanReceive, traceEvGoBlockRecv, 3)
	// 这里被唤醒, 继续执行
	if mysg != gp.waiting {
		throw("G waiting list is corrupted")
	}
	gp.waiting = nil
	if mysg.releasetime > 0 {
		blockevent(mysg.releasetime-t0, 2)
	}
	closed := gp.param == nil
	gp.param = nil
	mysg.c = nil
	releaseSudog(mysg)
	return true, !closed
}

// recv processes a receive operation on a full channel c.
// There are 2 parts:
// 1) The value sent by the sender sg is put into the channel
//    and the sender is woken up to go on its merry way.
// 2) The value received by the receiver (the current G) is
//    written to ep.
// For synchronous channels, both values are the same.
// For asynchronous channels, the receiver gets its data from
// the channel buffer and the sender's data is put in the
// channel buffer.
// Channel c must be full and locked. recv unlocks c with unlockf.
// sg must already be dequeued from c.
// A non-nil ep must point to the heap or the caller's stack.
func recv(c *hchan, sg *sudog, ep unsafe.Pointer, unlockf func(), skip int) {
	if c.dataqsiz == 0 {
		if raceenabled {
			racesync(c, sg)
		}
		if ep != nil {
			// copy data from sender
			recvDirect(c.elemtype, sg, ep)
		}
	} else {
		// Queue is full. Take the item at the
		// head of the queue. Make the sender enqueue
		// its item at the tail of the queue. Since the
		// queue is full, those are both the same slot.
		qp := chanbuf(c, c.recvx)
		if raceenabled {
			raceacquire(qp)
			racerelease(qp)
			raceacquireg(sg.g, qp)
			racereleaseg(sg.g, qp)
		}
		// copy data from queue to receiver
		if ep != nil {
			typedmemmove(c.elemtype, ep, qp)
		}
		// copy data from sender to queue
		typedmemmove(c.elemtype, qp, sg.elem)
		c.recvx++
		if c.recvx == c.dataqsiz {
			c.recvx = 0
		}
		c.sendx = c.recvx // c.sendx = (c.sendx+1) % c.dataqsiz
	}
	sg.elem = nil
	gp := sg.g
	unlockf()
	gp.param = unsafe.Pointer(sg)
	if sg.releasetime != 0 {
		sg.releasetime = cputicks()
	}
	goready(gp, skip+1)
}

// compiler implements
//
//	select {
//	case c <- v:
//		... foo
//	default:
//		... bar
//	}
//
// as
//
//	if selectnbsend(c, v) {
//		... foo
//	} else {
//		... bar
//	}
//
func selectnbsend(c *hchan, elem unsafe.Pointer) (selected bool) {
	return chansend(c, elem, false, getcallerpc())
}

// compiler implements
//
//	select {
//	case v = <-c:
//		... foo
//	default:
//		... bar
//	}
//
// as
//
//	if selectnbrecv(&v, c) {
//		... foo
//	} else {
//		... bar
//	}
//
func selectnbrecv(elem unsafe.Pointer, c *hchan) (selected bool) {
	selected, _ = chanrecv(c, elem, false)
	return
}

// compiler implements
//
//	select {
//	case v, ok = <-c:
//		... foo
//	default:
//		... bar
//	}
//
// as
//
//	if c != nil && selectnbrecv2(&v, &ok, c) {
//		... foo
//	} else {
//		... bar
//	}
//
func selectnbrecv2(elem unsafe.Pointer, received *bool, c *hchan) (selected bool) {
	// TODO(khr): just return 2 values from this function, now that it is in Go.
	selected, *received = chanrecv(c, elem, false)
	return
}

//go:linkname reflect_chansend reflect.chansend
func reflect_chansend(c *hchan, elem unsafe.Pointer, nb bool) (selected bool) {
	return chansend(c, elem, !nb, getcallerpc())
}

//go:linkname reflect_chanrecv reflect.chanrecv
func reflect_chanrecv(c *hchan, nb bool, elem unsafe.Pointer) (selected bool, received bool) {
	return false, chanrecv(c, elem, !nb)
}

//go:linkname reflect_chanlen reflect.chanlen
func reflect_chanlen(c *hchan) int {
	if c == nil {
		return 0
	}
	return int(c.qcount)
}

//go:linkname reflect_chancap reflect.chancap
func reflect_chancap(c *hchan) int {
	if c == nil {
		return 0
	}
	return int(c.dataqsiz)
}

//go:linkname reflect_chanclose reflect.chanclose
func reflect_chanclose(c *hchan) {
	closechan(c)
}

func (q *waitq) enqueue(sgp *sudog) {
	sgp.next = nil
	x := q.last
	if x == nil {
		sgp.prev = nil
		q.first = sgp
		q.last = sgp
		return
	}
	sgp.prev = x
	x.next = sgp
	q.last = sgp
}

func (q *waitq) dequeue() *sudog {
	for {
		sgp := q.first
		if sgp == nil {
			return nil
		}
		y := sgp.next
		if y == nil {
			q.first = nil
			q.last = nil
		} else {
			y.prev = nil
			q.first = y
			sgp.next = nil // mark as removed (see dequeueSudog)
		}

		// if a goroutine was put on this queue because of a
		// select, there is a small window between the goroutine
		// being woken up by a different case and it grabbing the
		// channel locks. Once it has the lock
		// it removes itself from the queue, so we won't see it after that.
		// We use a flag in the G struct to tell us when someone
		// else has won the race to signal this goroutine but the goroutine
		// hasn't removed itself from the queue yet.
		if sgp.isSelect {
			if !atomic.Cas(&sgp.g.selectDone, 0, 1) {
				continue
			}
		}

		return sgp
	}
}

func (c *hchan) raceaddr() unsafe.Pointer {
	// Treat read-like and write-like operations on the channel to happen at this address. 
	// Avoid using the address of qcount or dataqsiz, because the len() and cap() builtins read those addresses, 
	// and we don't want them racing with operations like close().
	return unsafe.Pointer(&c.buf)
}

func racesync(c *hchan, sg *sudog) {
	racerelease(chanbuf(c, 0))
	raceacquireg(sg.g, chanbuf(c, 0))
	racereleaseg(sg.g, chanbuf(c, 0))
	raceacquire(chanbuf(c, 0))
}
