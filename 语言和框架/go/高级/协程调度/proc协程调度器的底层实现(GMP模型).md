# proc协程调度器的底层实现(GMP模型)

参考文章

1. [【我的架构师之路】- golang源码分析之协程调度器底层实现( G、M、P)](https://blog.csdn.net/qq_25870633/article/details/83445946)
    - GMP模型概念及流程介绍
2. [深入Golang调度器之GMP模型](https://www.cnblogs.com/sunsky303/p/9705727.html)
    - 源码分析比较用心
3. [协程调度时机一：系统调用](https://zhuanlan.zhihu.com/p/29970624)
    - 用后妈比喻M对象, 很生动
    - 结婚, 离婚, 前夫, 前妻
4. [Go scheduler 概览](https://qcrao.com/ishare/go-scheduler/#truego-scheduler-%E6%A6%82%E8%A7%88)

M是machine的首字母, 在当前版本的golang中等同于系统线程.

M可以运行两种代码:

1. go代码, 即goroutine, M运行go代码需要一个P
2. 原生代码, 例如阻塞的syscall, M运行原生代码不需要P

M会从运行队列中取出G, 然后运行G, 如果G运行完毕或者进入休眠状态, 则从运行队列中取出下一个G运行, 周而复始。
有时候G需要调用一些无法避免阻塞的原生代码, 这时M会释放持有的P并进入阻塞状态, 其他M会取得这个P并继续运行队列中的G.
go需要保证有足够的M可以运行G, 不让CPU闲着, 也需要保证M的数量不能过多。通常创建一个M的原因是由于没有足够的M来关联P并运行其中可运行的G。而且运行时系统执行系统监控的时候，或者GC的时候也会创建M。

比如开启两个G, 第一个G运行到IO阻塞部分休眠并切换让出M, 当第二个G正在运行时, 第一个的IO已经处理好等待唤醒, 此时M被第二个G占用, 就需要再创建一个M去运行第一个G.

猜测G的数量 = M的数量(正在运行的G的数量) + 正在休眠的G的数量.
