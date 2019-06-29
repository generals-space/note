# mutex锁实现

参考文章

1. [golang之sync.Mutex互斥锁源码分析](https://www.jianshu.com/p/ffe646ada7b4)
    - 详细介绍了Mutex结构成员和Lock与Unlock方法的源码
2. [go sync.Mutex 设计思想与演化过程 （一）](https://www.cnblogs.com/niniwzw/archive/2013/06/24/3153955.html)
3. [golang 1.10 mutex互斥锁源码](https://blog.csdn.net/weixin_40318210/article/details/80301288)
    - 对Mutex源码的解读深入到了runtime包的sema.go层面, 探究了golang中信号量机制的实现, 十分深入.

Mutex: 互斥量的缩写. mutual(互相) exclusion(排斥, 独占)

> semaphore: 一个数值, 尝试获取semaphore的操作都会被加入等待队列.

参考文章1只介绍了`sync/mutex.go`, 但实际上mutex的实现还依赖了信号量semaphore. 所以还需要阅读同目录下的`sync/runtime.go`, 不过这里仅是声明, 实际定义的代码在runtime包里.

参考文章2给出了mutex的演进历史, 及使用semaphore实现的最简的mutex互斥锁示例.

但是单纯的信号量无法实现重入锁读写锁逻辑, 需要额外封装...不过性能呢?

目前mutex的实现中, 包含了两种模式: normal模式和starvation模式. 

一开始默认处于normal模式. 在normal模式中, 每个新加入竞争锁行列的协程都会直接参与到锁的竞争当中来. 

而处于starvation模式时, 所有新进入的协程都会直接被放入等待队列中挂起, 直到其所在队列之前的协程全部执行完毕. 

在normal模式中协程的挂起等待时间如果大于某个值, 就会进入starvation模式.

normal模式下各协程会在获取mutex中的sema成员时阻塞, 直到某个协程将其释放.

