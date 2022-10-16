# Python-xmlrpclib应用

参考文章

1. [xmlrcp学习 - python中使用xmlrpc](http://www.cnblogs.com/coderzh/archive/2008/12/03/1346994.html)

XML-RPC的全称是XML Remote Procedure Call，即**XML远程方法调用**。

它是一套允许运行在不同操作系统、不同环境的程序实现基于Internet过程调用的规范和一系列的实现。

这种远程过程调用使用http作为传输协议，XML作为传送信息的编码格式。Xml-Rpc的定义尽可能的保持了简单，但同时能够传送、处理、返回复杂的数据结构。

XML-RPC是工作在Internet上的远程过程调用协议。一个XML-RPC消息就是一个**请求体为xml**的`POST`请求，被调用的方法在服务器端执行并将执行结果以`xml`格式编码后返回。

## 1. 简单示例-远程函数调用

首先创建rpc server.

```py
#!/usr/bin/python
#!coding:utf-8

import SimpleXMLRPCServer

class MyObject:
    def sayHello(self):
        return "hello xmlrpc"

obj = MyObject()
server = SimpleXMLRPCServer.SimpleXMLRPCServer(("localhost", 80))
server.register_instance(obj)

print "Listening on port 80"
server.serve_forever()
```

运行它

```
[root@519f248ca8f2 ~]# python rpc_server.py 
Listening on port 80
```

使用`nmap`工具探测发现, rpc服务器监听的端口协议为`http`.

```
[root@519f248ca8f2 ~]# nmap -sS -p 80 localhost

Starting Nmap 5.51 ( http://nmap.org ) at 2017-02-04 01:32 UTC
Nmap scan report for localhost (127.0.0.1)
Host is up (-1700s latency).
Other addresses for localhost (not scanned): 127.0.0.1
PORT   STATE SERVICE
80/tcp open  http

Nmap done: 1 IP address (1 host up) scanned in 0.08 seconds
```

然后通过`xmlrpclib`库与远程rpc服务器进行通信.

```py
#!/usr/bin/python
#!coding:utf-8

import xmlrpclib

server = xmlrpclib.ServerProxy("http://localhost:80")
words = server.sayHello()
print "result: " + words
```

运行它, 结果如下

```
[root@519f248ca8f2 ~]# python ./rpc_client.py 
result: hello xmlrpc
```

同时, rpc server端的日志输出为

```log
127.0.0.1 - - [04/Feb/2017 09:34:25] "POST /RPC2 HTTP/1.1" 200 -
```

------

可以看出, rpc_client调用了rpc服务器的`sayHello`方法. 最初我对这种通信方式保持怀疑的态度, 因为同样的功能可以使用`restful`形式实现, 并且更加优雅美观, 都是由远程服务器提供了调用接口(API), 并且返回数据.

比如, 我们可以使用`/userid/1`从远程得到id为1的用户数据, 或者使用`?a=1&b=2`的方式传递参数. 可以想像, 服务端代码必须从request url中取出参数`userid`, `a`, `b`等, 然后进行操作, 这是显示易见的. 并且无论怎样, 服务端的响应数据都肯定是字符串类型, 即使是经过了格式化, 也只是符合某种语法(json, xml什么的)的字符串, 还是需要客户端加工一下的.

那么, 如果我想传递一个对象, 或是想直接获取一个对象呢? 想获取一个列表, 字典, 甚至是一个函数呢?

类似于`ajax`的方式传输给服务端一个json类型的字符串, 我想这就是`xmlrpclib`库的真正用法了, 只是省略了客户端/服务端序列化字符串的过程.

## 2. 进阶应用-传递/返回对象

```py
#!/usr/bin/python
#!coding:utf-8

import SimpleXMLRPCServer

class MyObject:
    def __init__(self):
        self.lists = ['a', 'b', 'c', 'd', 'e', 'f']
    def sayHello(self):
        return "hello xmlrpc"
    def getList(self, start):
        print self.lists[start:]
        return self.lists[start:]
obj = MyObject()
server = SimpleXMLRPCServer.SimpleXMLRPCServer(("localhost", 80))
server.register_instance(obj)

print "Listening on port 80"
server.serve_forever()
```

对应的客户端代码

```py
#!/usr/bin/python
#!coding:utf-8

import xmlrpclib

server = xmlrpclib.ServerProxy("http://localhost:80")
## words = server.sayHello()
lists = server.getList(2)
print "result: "
print lists
```

运行它们

服务端

```
[root@519f248ca8f2 ~]# python rpc_server.py 
Listening on port 80
```

客户端

```
[root@519f248ca8f2 ~]# python ./rpc_client.py 
result: 
['c', 'd', 'e', 'f']
```

对应服务器端的日志输出为

```
['c', 'd', 'e', 'f']
127.0.0.1 - - [04/Feb/2017 10:25:45] "POST /RPC2 HTTP/1.1" 200 -
```

这样, 我们直接可以从服务端获取一个列表对象, 而不是字符串.