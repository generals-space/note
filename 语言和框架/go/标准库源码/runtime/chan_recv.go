// 所有通过`<-`从channel中读取数据的代码, 最终都会调用`chanrecv1()`函数
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

// chanrecv 从 c 中读取数据, 并将读到的数据写入 ep .
// ep 可以为 nil, 表示不 care 接收到的结果(只把`<-`当成阻塞代码流程的一种手段)
// 这里的 ep 表示 <- 左侧接收数据的变量指针.
// 如果 block == false 且 channel 中没有数据时, 则直接返回(false, false), 
// 否则, 如果c被关闭, 
// ep 对象如果不为空, 则一定指向堆, 或是调用者的栈(如a <- chan, a为栈中的变量)
func chanrecv(c *hchan, ep unsafe.Pointer, block bool) (selected, received bool) {
	// raceenabled: don't need to check ep, 
	// as it is always on the stack
	// or is new memory allocated by reflect.

	if debugChan {
		print("chanrecv: chan=", c, "\n")
	}

	// ...什么情况下 c 可能为 nil?
	// var c chan int 这样, 只声明类型并未初始化?
	if c == nil {
		if !block {
			return
		}
		gopark(nil, nil, waitReasonChanReceiveNilChan, traceEvGoStop, 2)
		throw("unreachable")
	}

	// Fast path: check for failed non-blocking operation 
	// without acquiring the lock.
	//
	// After observing that the channel is not ready for receiving,
	// we observe that the channel is not closed. 
	// Each of these observations is a single word-sized read
	// (first c.sendq.first or c.qcount, and second c.closed).
	// Because a channel cannot be reopened, 
	// the later observation of the channel being not closed implies that
	// it was also not closed at the moment of the first observation. 
	// We behave as if we observed the channel at that moment
	// and report that the receive cannot proceed.
	//
	// The order of operations is important here: 
	// reversing the operations can lead to
	// incorrect behavior when racing with a close.
	

	// 下面的 if 块就是上面英文注释中的`Fast Path`捷径.
	// 如果是非阻塞读取操作(select + default 可以实现),
	// 且目标 channel 为无缓冲通道, 且此时没有写协程(无缓冲 channel 的读操作是直接到 写协程中取数据的)
	// 并且确认 channel 未关闭, 这种情况下根本没可能读到数据.
	// 则直接返回, 不需要加锁.
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
	// 如果 channel 已经被关闭且 channel 缓冲中没有数据了
	// 则直接返回, 如果 ep 位于 <- 左侧, 希望得到一个结果,
	// 则 typedmemclr 会为其赋上该类型的默认值.
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

	// 下面的操作中, channel 可能为 closed, 但不影响读操作.

	// sendq 队列非空, 表示有goroutine对象等待写数据
	// (此时 channel 必然已满). 
	// 那么跳过channel缓冲, 直接从写队列接收数据.
	if sg := c.sendq.dequeue(); sg != nil {
		// If buffer is size 0, receive value directly from sender. 
		// Otherwise, receive from head of queue and 
		// add sender's value to the tail of the queue 
		// (both map to the same buffer slot because the queue is full).
		// 
		recv(c, sg, ep, func() { unlock(&c.lock) }, 3)
		return true, true
	}

	// 运行到此处, 表示 sendq 队列为空, channel 至少没满.

	// qcount > 0, 表示 channel 有数据, 读协程可以不阻塞.
	// 于是直接从 buf 中取数据.
	if c.qcount > 0 {
		// Receive directly from queue
		qp := chanbuf(c, c.recvx)
		if raceenabled {
			raceacquire(qp)
			racerelease(qp)
		}
		if ep != nil {
			// 将 qp 处的数据拷贝至 ep, ...其实应该说是剪切吧.
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

	// 运行到此处, 表示 sendq 队列为空且 channel 也为空. 

	// 如果读协程是非阻塞的, 那么可以直接返回空了.
	// channel 的非阻塞如何实现? 只要在 select 中加入一个 default 就可以了.
	// select {
    // case msg := <-messages:
    //     fmt.Println("received message", msg)
    // default:
    //     fmt.Println("no message received")
    // }
	if !block {
		unlock(&c.lock)
		return false, false
	}

	// 此时读协程是读不数据的, 只能等待.
	// 因此需要将当前 goroutine 加入 c 的 recvq 队列, 并将其阻塞.
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
	// 这里 gp 和 mysg 还相互引用...
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

// caller: chanrecv() sendq 非空, channel 已满时调用.
// recv 会判断 channel 是否为无缓冲.
// 如果是, 则调用 recvDirect() 直接从写协程中拷贝数据
// 如果不是, 则读协程取 buf[recvx] 处的数据, 
// 然后将 sendq 要写的数据追加到 buf[sendx], 
// 不管怎样, 都需要将 sendq 中的写协程唤醒.
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
		// Queue is full. 
		// Take the item at the head of the queue. 
		// Make the sender enqueue
		// its item at the tail of the queue. 
		// Since the queue is full, those are both the same slot.
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
	// 不过 channel 是不是有缓冲, sendq 非空必然导致写协程阻塞,
	// 这里读完数据还需要将阻塞的写协程唤醒.
	sg.elem = nil
	gp := sg.g
	unlockf()
	gp.param = unsafe.Pointer(sg)
	if sg.releasetime != 0 {
		sg.releasetime = cputicks()
	}
	goready(gp, skip+1)
}

// caller: recv() 当 channel 为无缓冲队列, 且有写协程阻塞时, 
// 发生读操作, 则会调用此函数直接从写协程处取数据.
// 因为无缓冲 channel 中的 buf 是没有分配空间的.
func recvDirect(t *_type, sg *sudog, dst unsafe.Pointer) {
	// dst is on our stack or the heap, src is on another stack.
	// The channel is locked, so src will not move during this
	// operation.
	src := sg.elem
	typeBitsBulkBarrier(t, uintptr(dst), uintptr(src), t.size)
	memmove(dst, src, t.size)
}
