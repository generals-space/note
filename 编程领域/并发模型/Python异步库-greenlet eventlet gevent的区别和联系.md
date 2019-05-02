# Python异步库-greenlet eventlet gevent的区别和联系

`greenlet`应该是一种概念或思想, 中文名叫**微线程**, `eventlet`和`gevent`是两种不同的实现.

关于这两者, 这篇文章[[gevent源码分析] 深度分析gevent运行流程](http://blog.csdn.net/yueguanghaidao/article/details/24281751)在开头有如下解释, 很精辟.

gevent是一个高性能网络库, 底层是libevent, 1.0版本之后是libev, 核心是greenlet. gevent和eventlet是近亲, 唯一不同的是eventlet是自己实现的事件驱动, 而gevent是使用libev. 两者都有广泛的应用, 如openstack底层网络通信使用eventlet, goagent是使用gevent. 
