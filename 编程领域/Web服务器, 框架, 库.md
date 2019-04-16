作者：李广胜
链接：https://www.zhihu.com/question/52574763/answer/131642822
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

现在一般python做web主要有三大部分，web server, web framework, async io.

web server: 承担端口监听和接受请求的任务

web framework 主要承担路由，业务逻辑等任务

有了web server和web framework基本就能运行了。

一般web framework库（比如flask），主要部分是web framework, 同时也自带一个性能不咋滴的web server，这样你在开发和调试时可以直接运行起来看看效果，但是在生产环境中，它自带的web server性能就不够用了。

除了这种性能不咋地的自带的web server，还有像gunicorn和uwsgi这种单独实现的性能强劲的web server，这种单独实现的web server和web framework配合起来用就可以提高整个应用的性能。

由于python的多线程基本就是个摆设，python web中一般用协程+异步IO的方式来实现并发。gevent就是一个协程+异步IO的库，其作用是将阻塞的应用变为非阻塞，来提高并发量。

总结，gunicorn和uwsgi是web server, flask或者bottle是web framework, gevent是async io。目前大部分测试下，uwsgi比gunicorn性能更好。uwsgi是C语言实现，gunicorn是纯Python实现，如果gunicorn有pypy加成应该性能会有所提高。