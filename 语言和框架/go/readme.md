[Go 语言高性能编程](https://geektutu.com/post/high-performance-go.html)
    - 在线电子书

[Go 语言设计与实现](https://draveness.me/golang/)
    - 源码分析
    - 基于 1.15
    - 插图完美, 思路清晰

[golang-nuts](https://groups.google.com/g/golang-nuts/)
    - golang 的发展历程
    - 可以搜索很多源码层面的问题

[Go 语言原本](https://golang.design/under-the-hood/)
    - 基于 1.15

[深入解析go](https://tiancaiamao.gitbooks.io/go-internals/content/zh/06.3.html)
    - 源码级分析, 比雨痕讲的还细节

[Go语言圣经（中文版）](https://books.studygolang.com/gopl-zh/index.html)

[Go语言101](https://gfw.go101.org/article/101.html)
    - 很深很详细

[Go语言入门指南](https://www.kancloud.cn/kancloud/the-way-to-go/72432)
    - 实际应用中常见的陷阱和错误

[Go示例学](https://www.kancloud.cn/itfanr/go-by-example/81617)
    - 标准库的简洁示例代码

[go学习笔记](https://github.com/qyuhen/book)

[GO语言的进阶之路](http://www.cnblogs.com/yinzhengjie/tag/GO%E8%AF%AD%E8%A8%80%E7%9A%84%E8%BF%9B%E9%98%B6%E4%B9%8B%E8%B7%AF/)
[Golang编程经验总结](https://blog.csdn.net/yxw2014/article/details/43451625)

[GO 命令教程](https://www.kancloud.cn/cattong/go_command_tutorial/261347)
    - go build, run, test等子命令的各选项和参数用法

[Go-Questions](https://github.com/qcrao/Go-Questions)
    - 从问题切入, 串连 Go 语言相关的所有知识, 融会贯通

--------------------------------------------------------------------------------

[disk io引起golang线程数暴涨的问题](http://xiaorui.cc/2018/04/23/disk-io%E5%BC%95%E8%B5%B7golang%E7%BA%BF%E7%A8%8B%E6%95%B0%E6%9A%B4%E6%B6%A8%E7%9A%84%E9%97%AE%E9%A2%98/)

1. golang的scheduler会帮你调度关联 PMG, 用户无法看到也无法创建原生系统线程, 且golang社区并没有要加入这个特性的意图. 如果有必要那么可以用cgo, 或者手动触发`runtime.LockOSThread`绑定.

2. golang写文件是阻塞, golang针对网络io使用了epoll事件机制来实现了异步非阻塞, 但磁盘io用不了这套, 因为磁盘始终是就绪状态, epoll是根据就绪通知, 就绪 & 非就绪状态. 现在常见的一些高性能框架都没有直接的解决磁盘io异步非阻塞问题. 像nginx之前读写磁盘会阻塞worker进程, 只有开启aio及线程池thread才会绕开阻塞问题. 可以说, 整个linux下, aio不是很完善, 最少比bsd和windows平台下稍微差点. 像windows下iocp就解决了磁盘io异步问题, 但他的原理关键也是线程池.

3. golang默认的最大线程数10000个线程, 这个是硬编码. 如果想要控制golang的pthread线程数可以使用 `runtime.SetMaxThreads()`

4. 每个pthread的stack栈空间是8MB, 当然是virt.(可以通过ulimit -a查看, `stack size`字段即是)
5. 一个socket连接一般占用8kb内存

## 问题???

单核主机上, 某个goroutine的缓冲区已满, 是否会影响其他goroutine? golang有办法将ta们调度起来吗?

DDos泛洪攻击, 耗尽的是主机的哪种资源? 端口, 内存? 连接数, 还是打开文件数?

## 20191230 reflect.DeepEqual()

1. [Go语言 bytes.Equal() 和 reflect.DeepEqual() 的不同](https://www.cnblogs.com/hanyu100/p/8717456.html)
2. [10x faster than reflect.DeepEqual](https://zhuanlan.zhihu.com/p/55654454)
