## golang的演进历史

参考文章

1. [深入理解golang 的栈](https://www.jianshu.com/p/7ec9acca6480)
    - golang动态调整栈空间的解决方案: 分段栈与连续栈
2. [Go 系列文章 8: select](http://xargin.com/go-select/)
    - select在应用层的使用示例及技巧
    - select源码分析
3. [关于Golang GC的一些误解--真的比Java GC更领先吗？](https://zhuanlan.zhihu.com/p/77943973)
4. [Goroutine并不是可剥夺的](https://www.eaglexiang.org/deprived_goroutine)
5. [Go 语言设计与实现 - 7.1 内存分配器](https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-memory-allocator/)
    - 在 Go 语言 1.10 以前的版本，堆区的内存空间都是连续的；但是在 1.11 版本，Go 团队使用稀疏的堆内存空间替代了连续的内存，解决了连续内存带来的限制以及在特殊场景下可能出现的问题。

1. 处理栈空间的弹性方案, 1.4之前使用的是分段栈, 从1.4开始使用连续栈. (来自参考文章1)
2. select的实现源码, 在1.11之前一直保留有hselect结构和newselect()构造函数, 从1.11开始, 移除了这些. (来自参考文章2)
3. gc机制.
    - 1.3以前是简单的标记清除法(需要STW), 从1.3开始gc为并行操作(分离了标记和清理的操作, 标记过程STW, 清理过程并发执行.(来自参考文章3))). 
    - 1.5开始应用三色标记法(回收过程主要有四个阶段, 其中, 标记和清理都并发执行的, 但标记阶段的前后需要STW一定时间来做GC的准备工作和栈的re-scan.). 
    - 1.8版本在引入混合屏障rescan来降低mark termination的时间.
4. golang从1.5版本实现了自举.
    - 自举, 即使用golang编译golang, 而替代了C编译器.
    - 自举的好处在于, golang编译器的官方开发者将直接使用golang来写代码, 他们将对golang本身特性更熟悉, 而不是用C写编译器的贼6, 对自家语言反而没那么熟悉(就像游戏开发者没法把自家游戏打通关一样, 这一点甚至不如高端玩家).
5. golang从1.4实现了抢占调度.
    - 在1.4之前, 假设一个协程在执行CPU密集任务, 没有机会进行网络请求, 读写硬盘, time.Sleep()等IO操作时, M与G将长时间关联无法解绑, 这将导致该G对象无法被归还. 如果所有M对象都绑定了执行CPU密集任务的G, Goroutine调度器等于不存在. 见参考文章4.
6. 在 Go 语言 1.10 以前的版本，堆区的内存空间都是连续的；但是在 1.11 版本，Go 团队使用稀疏的堆内存空间替代了连续的内存，解决了连续内存带来的限制以及在特殊场景下可能出现的问题。
