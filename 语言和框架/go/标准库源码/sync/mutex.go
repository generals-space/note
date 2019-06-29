// Package sync provides basic synchronization primitives such as mutual
// exclusion locks. Other than the Once and WaitGroup types, most are intended
// for use by low-level library routines. Higher-level synchronization is
// better done via channels and communication.
//
// Values containing the types defined in this package should not be copied.
package sync

import (
	"internal/race"
	"sync/atomic"
	"unsafe"
)

func throw(string) // 在runtime包中实现. 类似于C语言中.h中的函数声明, 实际定义在.c文件中.

// Mutex是mutual(相互) exclusion(独占, 排斥) lock. 默认为未占用状态.
//
// A Mutex must not be copied after first use.
type Mutex struct {
	state int32  // 运行状态. 将一个32位整数拆分为 当前阻塞的goroutine数(29位)|饥饿状态(1位)|唤醒状态(1位)|锁状态(1位) 的形式, 来简化字段设计
	sema  uint32 // semaphore 信号量, 实际用来阻塞协程的依据.
}

type Locker interface {
	Lock()
	Unlock()
}

const (
	mutexLocked      = 1 << iota // 1 0001 用最后一位表示当前对象锁的状态, 0-未锁住 1-已锁住
	mutexWoken                   // 2 0010 用倒数第二位表示当前对象是否被唤醒 0-唤醒 1-未唤醒
	mutexStarving                // 4 0100 用倒数第三位表示当前对象是否为饥饿模式, 0为正常模式, 1为饥饿模式.
	mutexWaiterShift = iota      // 3, 从倒数第四位往前的bit位表示在排队等待的goroutine数
	// 前三个数值是常规的位掩码定义将1分别左移0, 1, 2位, 此时iota值为2.
	// 接着mutexWaiterShift出现, 但ta的值直接为iota, 而不再是1 << iota, 所以ta直接取了iota的下一个值3.

	// Mutex的公平性.
	//
	// Mutex可以分为2种操作模式: 正常模式和饥饿模式
	// 正常模式下, 等待获取锁的协程按照FIFO顺序排队, 但是协程被唤醒后并不能立即得到锁, 而是需要与新到达的协程进行争夺.
	// 因为新到达的goroutine已经在CPU上运行了, 所以被唤醒的goroutine很大概率是争夺mutex锁是失败的.
	// 出现这样的情况时候, 被唤醒的goroutine需要排队在队列的前面(就是插队咯?).
	// 如果被唤醒的goroutine有超过1ms没有获取到mutex锁, 那么它就会将锁转换为饥饿模式.
	//
	// 饥饿模式下, 锁的所有权将由刚解锁的协程中直接移交给排在队列前面的协程.
	// 新到达的协程也不会去争夺, 即使ta已经被某个协程解锁且没有自旋.
	// 取而代之的是将自己放置到等待队列的末尾...(擦, 这太坑了吧)
	//
	// 一个处于等待队列中的协程获得锁后, 当ta发现自己已经是等待队列中最后一个成员(即之后没有新的等待协程),
	// 或是ta的等待时间不超过1ms(...说明锁的需求不是特别紧迫?), ta会将锁转换为正常模式.
	//
	// 当等待队列中的协程多次请求锁并且能成功获取到时, 正常模式的性能是相当不错的, 即使会阻塞一部分其他的协程继续等待.
	// 但总有特别倒霉的协程存在, 有很小的机率一直无法获取锁而导致极严重的延迟, 饥饿模式就是为避免这种最坏情况的出现.
	starvationThresholdNs = 1e6 // 1e6纳秒, 即1ms
)

// Lock 加锁
// 如果当前mutex对象正在被使用(即已经加过锁), 调用此函数的协程将被阻塞直接mutex可用.
func (m *Mutex) Lock() {
	// 这里是第一个调用Lock()函数时的情况, 直接进入这个if块, 将其标记为已占用然后就return了.
	// 如果m.state=0, 说明当前的对象还没有被锁住, 进行原子性赋值操作设置为mutexLocked状态, CompareAnSwapInt32返回true
	// 否则说明对象已被其他goroutine锁住, 不会进行原子赋值操作设置, CopareAndSwapInt32返回false, 继续往下执行.
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		if race.Enabled {
			race.Acquire(unsafe.Pointer(m))
		}
		return
	}

	var waitStartTime int64 // 开始等待的时间戳
	starving := false       // 饥饿模式标识
	awoke := false          // 唤醒标识
	iter := 0               // 自旋次数
	old := m.state          // 保存mutex锁当前的状态

	// 这个for循环使用了cas的同步模式, 即compare and swap.
	// 调用Lock()的协程不断尝试获取锁, 当mutex的状态为未占用, 就使用CAS将最后一位改为已占用.
	// 这个for循环其实就是各协程抢占的过程.
	for {
		// 1. 与操作, 这一表达式表示old处于locked状态, 但不是饥饿模式.
		// 2. runtime_canSpin(iter)判断是可以进入自旋锁(超过一定次数就无法再次进入).
		// 因为在饥饿模式下, 锁的所有权要直接移除给等待的协程, 无法通过抢占获取到.
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			// 主动自旋是有意义的. 尝试设置mutexwake标志, 通知解锁, 不要唤醒其他阻塞的协程.

			// 1. 确定目前仍未被唤醒
			// 2. 确定有其他协程在排队
			// 3. 如果m.state状态未变, 则将锁的状态标记为唤醒.
			if !awoke && old&mutexWoken == 0 &&
				old>>mutexWaiterShift != 0 &&
				atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
				awoke = true
			}
			// runtime_doSpin()进入自旋
			// 注意: 进入自旋锁后当前goroutine并不挂起, 仍然在占用cpu资源,
			// 所以重试一定次数后, 不会再进入自旋锁逻辑.
			runtime_doSpin()
			iter++        // 自旋次数加1
			old = m.state // 保存mutex对象状态
			continue
		}
		
		// 代码运行到这里, 说明要么锁处于饥饿模式, 要么无法进入自旋.

		// 猜测new用于确定当前协程有机会获得锁后要将锁对象设置成什么状态
		// 注意: 锁处于饥饿状态不一定被占用, 应该也存在移交过程中的短暂释放, 只不过新的协程只能排队不能抢占罢了
		new := old
		// 如果没有处于饥饿状态, 就只设置其为普通的lock状态
		if old&mutexStarving == 0 {
			new |= mutexLocked
		}
		// 如果锁被占用, 且处于饥饿模式, 则不尝试抢占.
		// 直接更新阻塞goroutine的数量, 等待锁的协程数目加1
		if old&(mutexLocked|mutexStarving) != 0 {
			new += 1 << mutexWaiterShift
		}
		// 当前协程将mutex转换成饥饿模式, 但需要mutex对象为被占用的状态, 否则不要操作.
		// Unlock expects that starving mutex has waiters, which will not
		// be true in this case.
		// 
		// 要进入这个条件, starving必须要为true.
		// 但在for循环的前面部分的代码中, 没有对starving做更改的操作.
		// 所以此条语句要满足, 必然不是第一个循环.
		if starving && old&mutexLocked != 0 {
			new |= mutexStarving
		}
		if awoke {
			// 当前协程已被唤醒(是因为得到了sema信号量?)
			if new&mutexWoken == 0 {
				throw("sync: inconsistent mutex state")
			}
			new &^= mutexWoken
		}

		// 只有在这个if块里, 才有可能跳出这个for循环.
		// 分析一下条件, 就是当m.state等于old的时候, 尝试将m.state更改为new.
		// 上面经过的步骤中, 对old的修改比较少, 所以还是要看这个CAS操作能否成功.
		// ...不过貌似一般都成功了.
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			// 所有协程都在通过CAS不断尝试重置m.state, 以运行到这里.
			// 到了这里, 其实可以说是对mutex的抢占成功了...???

			if old&(mutexLocked|mutexStarving) == 0 {
				// 到这里, 上面的CAS已经成功将m.state替换成了new
				// 如果之前old处于未加锁状态, 也未处于饥饿模式, 就可以跳出循环了.
				break // locked the mutex with CAS
			}
			// waitStartTime != 0, 表示这已经不是第一次循环了.
			// 下面在SemacquireMutex()的时候, 就可以把当前协程加入到等待队列头部.
			queueLifo := waitStartTime != 0
			if waitStartTime == 0 {
				waitStartTime = runtime_nanotime()
			}
			// 一般来说, 协程就是在这里阻塞的, 主要就是其中的信号量成员的作用.
			runtime_SemacquireMutex(&m.sema, queueLifo)
			// 此时当前协程已经获得了锁, 根据当前协程等待的时间以决定是否要转换为饥饿模式, 但不是在这一轮.
			starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
			old = m.state
			// 如果当前锁处于饥饿模式, 就可以跳出for循环了.
			if old&mutexStarving != 0 {
				// 如果当前协程已经被唤醒, 且mutex处于饥饿模式, 可以认定mutex的所有权是被移交给我们的.
				// 但是出现如下情况, 
				// 1. mutex仍处于Lock状态, 或是
				// 2. 等待队列中协程的个数为0(但当前协程本身就应该在等待队列中)
				// 说明mutex处于一种不一致的状态, 抛出异常.
				if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
					throw("sync: inconsistent mutex state")
				}
				delta := int32(mutexLocked - 1<<mutexWaiterShift)
				// 如果当前协程等待时间很短, starving为false, 
				// 或是等待队列中只有当前协程一个成员了, 
				// 那就退出饥饿模式
				if !starving || old>>mutexWaiterShift == 1 {
					// Critical to do it here and consider wait time.
					// Starvation mode is so inefficient, that two goroutines
					// can go lock-step infinitely once they switch mutex
					// to starvation mode.
					delta -= mutexStarving
				}
				atomic.AddInt32(&m.state, delta)
				break
			}
			// 如果mutex不是饥饿模式, 那么即使当前协程已经抢占到了sema信号量, 
			// 也只能将awoke设置为true, 在下一次循环处理.
			awoke = true
			iter = 0
		} else {
			old = m.state
		}
	}

	if race.Enabled {
		race.Acquire(unsafe.Pointer(m))
	}
}

// Unlock 解锁
// 对一个未加锁的mutex对象进行Unlock(), 会引发运行时错误.
// 一个已经被锁定的mutex对象并不会和各协程对象有关联.
// 因为有很多时候是在a协程中加锁, 再由b协程解锁.
// 不需要像channel一样把等待读写的协程对象保存在队列里.
func (m *Mutex) Unlock() {
	if race.Enabled {
		_ = m.state
		race.Release(unsafe.Pointer(m))
	}

	// new是解锁mutex时为其设置的状态.

	// 移除mutexLocked位, 表示移除锁的占用状态.
	new := atomic.AddInt32(&m.state, -mutexLocked)
	// 验证锁状态(如果m原本没有加锁, 上面的操作会造成if判断的结果不为0)
	// 对一个未上锁的mutex对象执行Unlock会抛出异常.
	if (new+mutexLocked)&mutexLocked == 0 {
		throw("sync: unlock of unlocked mutex")
	}
	// 此时已经解除了占用状态.

	// 判断是否处于正常模式
	if new&mutexStarving == 0 {
		old := new
		for {
			// 如果没有等待的协程, 或是有协程已经被唤醒, 又或者有协程占用了锁, 
			// 可以直接退出, 不必唤醒任何协程(即, 不必调用Semrelease). 
			if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
				return
			}
			// 现在有能力去唤醒某个协程了. 做3件事
			// 1. 协程数减1
			// 2. 设置状态为唤醒
			// 3. 释放sema信号量, 唤醒一个等待的协程
			new = (old - 1<<mutexWaiterShift) | mutexWoken
			if atomic.CompareAndSwapInt32(&m.state, old, new) {
				// 这里第2个参数为false, 随机选择一个等待协程唤醒.
				runtime_Semrelease(&m.sema, false)
				return
			}
			old = m.state
		}
	} else {
		// 如果锁处于饥饿模式, 将sema所有权交给第一个等待的协程.
		// 但是注意, 这里没有设置唤醒标志位...
		// 
		// 注意: 这里mutex还没有被设置占用, 需要由下一个将会获取ta的协程自己设置.
		// 但是互斥锁仍然被认为是锁定的, 如果互斥对象被设置, 所以新来的goroutines不会得到它.
		// ...为什么, 占用标志位已经被移除了啊?
		runtime_Semrelease(&m.sema, true)
	}
}
