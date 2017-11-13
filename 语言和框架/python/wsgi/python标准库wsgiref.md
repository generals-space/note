

参考文章

1. [WSGI初探及wsgiref简单实现](http://blog.csdn.net/laughing2333/article/details/51288660)

2. [Pthon WSGI心得](http://blog.csdn.net/liukeforever/article/details/6892428)

3. [WSGI接口 - 廖雪峰的官方网站](https://www.liaoxuefeng.com/wiki/001374738125095c955c1e6d8bb493182103fac9270762a000/001386832689740b04430a98f614b6da89da2157ea3efe2000)

## 协议分析

`wsgiref.simple_server`中有一个`make_server`方法用于创建http服务器, 同时要将遵循了wsgi协议的应用接口传入.

```py
from wsgiref.simple_server import make_server, demo_app
application = demo_app

server = make_server("0.0.0.0", 8000, application)
server.handle_request()
```

其中`make_server`创建的server实例对象默认为`WSGIServer`类的实例, 并且处理接口默认为`WSGIRequestHandler`.

`WSGI Server`底层依赖于http server到socket server, 暂时不深究.

在`handle_request`运行后, 服务器开始接受客户端请求, 需要说明的是, `handle_request`的执行结果就是实例化了一个`WSGIRequestHandler`类, 接受的请求都将由这个类处理.

这个类将获取环境变量等数据, 并且实例化一个`ServerHandler`类, 然后把结果传给这个实例. 容易混淆的是, 这个实例也有一个run方法, 这个方法将`finish_response`. 响应完成.

## wsgiref简单使用

wsgiref模块提供了一个简单的的WSGI Server和WSGI Application, 

```py
#!encoding: utf-8
from wsgiref.simple_server import make_server, demo_app

application = demo_app

# 参数分别为服务器的IP地址和端口, 以及应用程序.
server = make_server("0.0.0.0", 8000, application)

# 处理一个request之后立即退出程序
server.handle_request()
```

访问这个地址, 你会得到一个'Hello World'的页面, 并且还输出了许多系统信息, 如下.

![](https://gitimg.generals.space/d6925b7fb8930931c9e921db787511bd.png)

server进程只接受一次请求就退出(因为`handle_request`已经完成了).

OK, 在这种程度之上, 我们至少还有两个要求: 1. 循环接受请求, 不退出; 2. 应用程序(即`application`)自定义, 至少要能接受request参数, 生成response响应.

第1个条件很容易满足, 直接使用`serve_forever`替换掉`handle_request`就可以;

至于第2个, 我们只要达到wsgi协议的要求就可以. demo_app就是这个一个简单的应用, wsgi协议要求

1. 应用程序必须是一个可调用的对象, 因此, 应用程序可以是一个函数, 一个类, 或者一个重载了`__call__`的类的实例;

2. 应用程序必须接受两个参数并且要按照位置顺序, 分别是`environ`（环境变量）, 以及`start_response`函数（负责将响应的status code, headers写进缓冲区但不返回给客户端）. 

3. 应用程序返回的结果必须是一个可迭代的对象.

参考文章1中按照函数, 类, 重载了`__call__`成员的类实例分别给出了简单的application实现...相当给力.

我们以类实例为例

```py
from wsgiref.simple_server import make_server

class instance_app:
    """
    当使用类的实例作为应用程序, application = instance_app(), not instance_app
    """
    def __call__(self, environ, start_response):
        status = "200 OK"
        response_headers = [('Content-type', 'text/plain')]
        start_response(status, response_headers)
        return ["Instantiate : My Own Hello World!"]

server = make_server("0.0.0.0", 8000, instance_app())
server.serve_forever()
```

重新访问, 就会得到我们定义的"Instantiate : My Own Hello World!"的输出.

能够看出, result的列表就是常规的响应体, 即`response`对象, 那么`request`对象呢? 我们如何获得请求参数? 

答案就在`environ`对象中, 我们可以打印这变量的内容. 它包括如下成员

- `QUERY_STRING`: GET请求参数

- `HTTP_HOST`: 请求头的的HOST值

- `SERVER_PORT`: 请求端口

...许多http请求头都会被格式化成`environ`对象的成员属性, 省了我们很多事情, 另外通过`start_response`可以设置响应码与响应头, 当然, 在纯粹的业务编程中这部分应该是框架提供的功能, 不需要我们手写了.

值得注意的是, `start_respon`是`ServerHandler`类提供的方法, 至于`ServerHandler`是什么时候传入的, 是我们接下来要研究的事情.

## make_server分析

我们可以理解, `make_server`实际上创建了一个http服务器实例, 当它执行`serve_forever`时正式开始工作. 实际上它不只是一个http服务器(像nginx一样那种), 而是一个WSGI服务器, 当然后者实际上是基于来实现的.

我们可以查看一个`make_server`的源码, 纯粹的函数, 不是类方法.

```py
def make_server(host, port, app, server_class=WSGIServer, handler_class=WSGIRequestHandler):
    """Create a new WSGI server listening on `host` and `port` for `app`"""
    server = server_class((host, port), handler_class)
    server.set_app(app)
    return server
```

`WSGIServer`就是我们要找的WSGI服务器类, 另外`WSGIRequestHandler`应该就是为我们对request与response进行初步处理的类. 这两个类也都定义在`wsgiref.simple_server`模块中.

> 它们是理解django, flask等内置服务器的基础, 它们或多或少的都对`WSGIServer`与`WSGIRequestHandler`进行过处理.

`makr_server`的返回值是`WSGIServer`的实例, 那么`serve_forever`就是我们下一步要分析的(其实应该分析`handle_request`的, `serve_forever`无非就是一个死循环).

好吧, `WSGIServer` -> `HTTPServer` -> `SocketServer.TCPServer` -> `SocketServer.BaseServer`这个继承链有点长, 其实`handle_request`和`serve_forever`都是在`BaseServer`类中定义的.