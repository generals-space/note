## golang的演进历史

参考文章

1. [深入理解golang 的栈](https://www.jianshu.com/p/7ec9acca6480)
    - golang动态调整栈空间的解决方案: 分段栈与连续栈
2. [Go 系列文章 8: select](http://xargin.com/go-select/)
    - select在应用层的使用示例及技巧
    - select源码分析

3. 处理栈空间的弹性方案, 1.4之前使用的是分段栈, 从1.4开始使用连续栈. (来自参考文章1)

4. select的实现源码, 在1.11之前一直保留有hselect结构和newselect()构造函数, 从1.11开始, 移除了这些. (来自参考文章2)

