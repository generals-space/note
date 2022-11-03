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
		// 创建无缓冲channel, 无缓冲channel是没有 buf 队列的,
		// 读协程的读取操作将直接从写协程处获取数据, 
		// 这部分操作在 recv() 函数中, 可以去查看一下.
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
		print(
			"makechan: chan=", c, 
			"; elemsize=", elem.size, 
			"; elemalg=", elem.alg, 
			"; dataqsiz=", size, 
			"\n",
		)
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

////////////////////////////////////////////////////////////
// send 部分
////////////////////////////////////////////////////////////

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

//////////////////////////////////////////////////////
// recv 部分
//////////////////////////////////////////////////////

// 以下是实现 select + channel 机制的部分, 由编译器负责翻译.

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
