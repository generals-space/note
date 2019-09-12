## golang的演进历史

参考文章

1. [深入理解golang 的栈](https://www.jianshu.com/p/7ec9acca6480)
    - golang动态调整栈空间的解决方案: 分段栈与连续栈
2. [Go 系列文章 8: select](http://xargin.com/go-select/)
    - select在应用层的使用示例及技巧
    - select源码分析
3. [关于Golang GC的一些误解--真的比Java GC更领先吗？](https://zhuanlan.zhihu.com/p/77943973)

1. 处理栈空间的弹性方案, 1.4之前使用的是分段栈, 从1.4开始使用连续栈. (来自参考文章1)
2. select的实现源码, 在1.11之前一直保留有hselect结构和newselect()构造函数, 从1.11开始, 移除了这些. (来自参考文章2)
3. gc机制, 1.3以前是简单的标记清除法(需要STW), 从1.3开始gc为并行操作(分离了标记和清理的操作, 标记过程STW, 清理过程并发执行.(来自参考文章3))). 1.5开始应用三色标记法(回收过程主要有四个阶段, 其中, 标记和清理都并发执行的, 但标记阶段的前后需要STW一定时间来做GC的准备工作和栈的re-scan.). 1.8版本在引入混合屏障rescan来降低mark termination的时间.
4. golang从1.5版本实现了自举, 即使用golang编译golang, 而替代了C编译器.
