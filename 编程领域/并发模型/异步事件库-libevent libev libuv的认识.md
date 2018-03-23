# 异步事件库-libevent libev libuv的认识

这三个是都C语言编写的, 存在已久的系统级异步事件库, 许多高级语言都用它们封装了自己的异步操作库. 比如nodejs原生支持异步语法, 是因为其底层使用了libuv; python的gevent异步库封装了`libev`(最初用的是`libevent`).

关于这3者, 初步入门的文章可以参考这个: [网络库libevent、libev、libuv对比](http://blog.csdn.net/lijinqi1987/article/details/71214974)

- libevent: 名气最大，应用最广泛，历史悠久的跨平台事件库；

- libev: 较`libevent`而言，设计更简练，性能更好，但对Windows支持不够好；

- libuv: 开发node的过程中需要一个跨平台的事件库，他们首选了libev，但又要支持Windows，故重新封装了一套，linux下用libev实现，Windows下用IOCP实现；

`libuv`和`libev`的区别十分容易理解. 关于`libevent`和`libev`, 这是两个经常拿来比较的库, 可以查看这篇文章: [[译]libev和libevent的设计差异](https://www.cnblogs.com/Lifehacker/p/whats_the_difference_between_libevent_and_libev_chinese.html), 这是libev作者写的一篇关于设计哲学的阐述. 简洁, 高效, 不追求大而全. 基本就是这些了. 

还有, 在上面那篇文章的最后提到另一篇文章[libev and libevent](https://blog.gevent.org/2011/04/28/libev-and-libevent/), 这是gevent的开发者博客上的文章, 简要说明了gevent从libevent切换到libev的决策过程.

> 回顾gevent，它实际需要的只是一个负责事件循环的C库，在上面的HTTP库和DNS库，都可以交由标准库强大得不得了的python完成。

最后, 这篇文章真值得一看, [使用 libevent 和 libev 提高网络应用性能——I/O模型演进变化史](http://blog.csdn.net/hguisu/article/details/38638183), 文章很长, 讲解得很详细.