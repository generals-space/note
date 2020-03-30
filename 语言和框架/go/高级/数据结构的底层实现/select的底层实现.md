参考文章

1. [【我的架构师之路】- golang源码分析之select的底层实现](https://blog.csdn.net/qq_25870633/article/details/83339538)
    - 优秀
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

select 语句执行时会对所有case中对应的chanel加锁
select 语句会创建select对象, 如果放在for循环中长期执行可能会频繁的分配内存

目前仍有一点不太明白, 在`selectgo()`函数中第3阶段, 因为没有可执行的case而陷入休眠, 之后有事件发生被唤醒后, 为什么要重新走一遍loop? 之后的操作是什么意思, 对应哪种情况? 参考文章3只是给出了这样的流程图, 但是并没有解释这样的目的.
