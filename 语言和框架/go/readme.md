[Go 语言高性能编程](https://geektutu.com/post/high-performance-go.html)
    - 在线电子书

[Go 语言设计与实现](https://draveness.me/golang/)
    - 源码分析
    - 基于 1.15
    - 插图完美, 思路清晰

[golang-nuts](https://groups.google.com/g/golang-nuts/)
    - golang 的发展历程
    - 可以搜索很多源码层面的问题

[disk io引起golang线程数暴涨的问题](http://xiaorui.cc/2018/04/23/disk-io%E5%BC%95%E8%B5%B7golang%E7%BA%BF%E7%A8%8B%E6%95%B0%E6%9A%B4%E6%B6%A8%E7%9A%84%E9%97%AE%E9%A2%98/)

1. golang的scheduler会帮你调度关联 PMG, 用户无法看到也无法创建原生系统线程, 且golang社区并没有要加入这个特性的意图. 如果有必要那么可以用cgo, 或者手动触发`runtime.LockOSThread`绑定.

3. golang写文件是阻塞, golang针对网络io使用了epoll事件机制来实现了异步非阻塞, 但磁盘io用不了这套, 因为磁盘始终是就绪状态, epoll是根据就绪通知, 就绪 & 非就绪状态. 现在常见的一些高性能框架都没有直接的解决磁盘io异步非阻塞问题. 像nginx之前读写磁盘会阻塞worker进程, 只有开启aio及线程池thread才会绕开阻塞问题. 可以说, 整个linux下, aio不是很完善, 最少比bsd和windows平台下稍微差点. 像windows下iocp就解决了磁盘io异步问题, 但他的原理关键也是线程池.

4. golang默认的最大线程数10000个线程, 这个是硬编码. 如果想要控制golang的pthread线程数可以使用 `runtime.SetMaxThreads()`

2. 每个pthread的stack栈空间是8MB, 当然是virt.(可以通过ulimit -a查看, `stack size`字段即是)
3. 一个socket连接一般占用8kb内存

## 问题???

单核主机上, 某个goroutine的缓冲区已满, 是否会影响其他goroutine? golang有办法将ta们调度起来吗?

DDos泛洪攻击, 耗尽的是主机的哪种资源? 端口, 内存? 连接数, 还是打开文件数?

## 20191230 reflect.DeepEqual()

1. [Go语言 bytes.Equal() 和 reflect.DeepEqual() 的不同](https://www.cnblogs.com/hanyu100/p/8717456.html)
2. [10x faster than reflect.DeepEqual](https://zhuanlan.zhihu.com/p/55654454)

## 20221025

golang 什么是协程泄露(Goroutine Leak), 以及 golang 是否需要协程池?
