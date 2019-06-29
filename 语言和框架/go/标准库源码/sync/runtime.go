package sync

import "unsafe"

// 本包内的函数只是声明(declarition), 实际实现的代码在runtime包中定义.

// Semacquire操作在*s<=0时会阻塞, >0时会将其减一然后返回.(与类unix系统提供的系统调用行为相似)
// semaphore被设计为用于同步的简单的P操作, golang并不支持开发者直接使用这个机制.
func runtime_Semacquire(s *uint32)

// SemacquireMutex与Semacquire相似.
// 后者是将排序的协程按照FIFO的顺序入队, 是将新成员加入等待队列尾.
// 前者则多了一个配置荐, 如果lifo为true, 则将新成员加入等待队列头.
func runtime_SemacquireMutex(s *uint32, lifo bool)

// Semrelease自动递增*s, 并通知一个因为调用Semacquire而阻塞的处于等待状态的协程.
// 与C语言中操作系统提供的信号量机制一致, 但是C语言中的互斥锁是用futex实现的...
// 它是一个简单的唤醒原语, 供同步库使用, 不应直接使用.
// 如果handoff为true, 将计数值直接传递给第一个等待协程.
// ...如果为false, 则随机选一个唤醒???
func runtime_Semrelease(s *uint32, handoff bool)

// Approximation of notifyList in runtime/sema.go. Size and alignment must
// agree.
type notifyList struct {
	wait   uint32
	notify uint32
	lock   uintptr
	head   unsafe.Pointer
	tail   unsafe.Pointer
}

// 函数定义在runtime/sema.go

func runtime_notifyListAdd(l *notifyList) uint32

func runtime_notifyListWait(l *notifyList, t uint32)

func runtime_notifyListNotifyAll(l *notifyList)

func runtime_notifyListNotifyOne(l *notifyList)

// 保证当前(sync包)与runtime的notifyList占用空间相同.
func runtime_notifyListCheck(size uintptr)

func init() {
	var n notifyList
	runtime_notifyListCheck(unsafe.Sizeof(n))
}

// Active spinning runtime support.
// runtime_canSpin 返回当前是否可以自旋.
func runtime_canSpin(i int) bool

// runtime_doSpin 开始自旋.
func runtime_doSpin()

func runtime_nanotime() int64
