参考文章

1. [【我的架构师之路】- golang源码分析之select的底层实现](https://blog.csdn.net/qq_25870633/article/details/83339538)

2. [Go 系列文章 8: select](http://xargin.com/go-select/)
    - select在应用层的使用示例及技巧
3. [由浅入深聊聊Golang中select的实现机制](https://blog.csdn.net/u011957758/article/details/82230316)
    - 一对select的陷阱示例
    - select的使用及运行机制
    - select的源码解读, 源码解释的不太详细, 但文末给出了一张select底层实现的流程图, 很清晰

本来以为golang的select语句底层使用了linux的select/epoll机制, 是IO复用的最佳实践...我又错了.

网络上大部分对select的源码分析都停留在<=1.10版本, 实际上自从1.11后, 源码中就不再包含`hselect`结构了, 参考文章2中也提到了这一点.

参考文章3提到select的执行流程分为4步:

1. 创建select. 调用newselect(), 创建hselect对象.
2. 注册case.
3. 执行select选择case. 调用selectgo()
4. 释放select

可以在`reflect_rselect()`函数中查看, 可以说, 此函数是select执行的入口.
