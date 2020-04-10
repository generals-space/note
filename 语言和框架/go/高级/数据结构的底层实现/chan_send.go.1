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
	// 读写一个 nil 的 channel 会阻塞, 使用close()关闭ta时会发生panic.
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
	// 如果是非阻塞写入(select+default), 且 channel 未关闭, 
	// 且为无缓冲队列且没有读协程等待, 或有缓冲但数据已满,
	// (这种情况下根本没可能成功写入), 则直接返回写入失败.
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

	// recvq 队列非空, 表示有读协程. 
	if sg := c.recvq.dequeue(); sg != nil {
		// Found a waiting receiver. 
		// We pass the value we want to send directly to the receiver, 
		// bypassing the channel buffer (if any).
		// 由 send() 判断是否跳过 buf, 直接将数据发给读协程.
		send(c, sg, ep, func() { unlock(&c.lock) }, 3)
		return true
	}

	// 运行到此处, 表示recvq队列为空, 且 channel 未满, 可以写入
	if c.qcount < c.dataqsiz {
		// Space is available in the channel buffer. 
		// Enqueue the element to send.
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

	// 运行到此处, 表示缓存队列已满且接收消息队列 recvq 为空.

	// 如果是非阻塞写入, 那么直接返回失败, 因为写协程不接受挂起
	if !block {
		unlock(&c.lock)
		return false
	}

	// 将当前的goroutine加入到send队列
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
	// The sudog has a pointer to the stack object, 
	// but sudogs aren't considered as roots of the stack tracer.
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
