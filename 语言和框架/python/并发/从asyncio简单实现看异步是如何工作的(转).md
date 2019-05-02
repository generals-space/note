# 从 asyncio 简单实现看异步是如何工作的(转)

原文链接

[从 asyncio 简单实现看异步是如何工作的](https://ipfans.github.io/2016/02/simple-implement-asyncio-to-understand-how-async-works/)

by ipfans

注：请使用3.5+版本运行以下代码。

## 1. 先从例子看起

首先我们来看一个 socket 通讯的例子，这个例子我们可以在官方 `socket` 模块的文档中找到部分原型代码：

```py
# echo.py
from socket import *  # 是的，这是一个不好的写法
def echo_server(address):
    sock = socket(AF_INET, SOCK_STREAM)
    sock.setsockopt(SOL_SOCKET, SO_REUSEADDR, 1)
    sock.bind(address)
    sock.listen(5)
    while True:
        client, addr = sock.accept()
        print("connect from", addr)
        echo_handler(client)
def echo_handler(client):
    while True:
        data = client.recv(10000)
        if not data:
            break
        client.send(str.encode("Got: ") + data)
    print("connection closed.")
if __name__ == '__main__':
    echo_server(('', 25000))
```

但是同步模式会有一个问题，当进行通讯是阻塞的，当一个连接占用时就会阻碍其他连接的继续，这个时候应该怎么更快的运行呢？

## 2. 回顾历史

在`asyncio`出现之前，我们都是怎么提高效率的呢？首先想到的方法就是多线程处理：

```py
# echo_thread.py
from socket import *
import _thread
def echo_server(address):
    sock = socket(AF_INET, SOCK_STREAM)
    sock.setsockopt(SOL_SOCKET, SO_REUSEADDR, 1)
    sock.bind(address)
    sock.listen(5)
    while True:
        client, addr = sock.accept()
        print("connect from", addr)
        _thread.start_new_thread(echo_handler, (client,))
def echo_handler(client):
    while True:
        data = client.recv(10000)
        if not data:
            break
        client.send(str.encode("Got: ") + data)
    print("connection closed.")
if __name__ == '__main__':
    echo_server(('', 25000))
```

当然了，我们都知道多线程之下总是会有一些问题的。那么还有更好的方案吗？如果你了解过**C10k问题**，你一定听过`epoll`, `kqueue`之类的大名。那么，能在 Python 中使用这些功能吗？答案是肯定的。那就是[select](https://docs.python.org/3.5/library/select.html)。


```py
# echo_select.py
from socket import *
import select
def echo_server(address):
    sock = socket(AF_INET, SOCK_STREAM)
    sock.setsockopt(SOL_SOCKET, SO_REUSEADDR, 1)
    sock.bind(address)
    sock.listen(5)
    input = [sock,]
    while True:
        r, _, _ = select.select(input, [], [])
        for s in r:
            if s == sock:
                client, addr = sock.accept()
                print("connect from", addr)
                echo_handler(client)
def echo_handler(client):
    while True:
        data = client.recv(10000)
        if not data:
            break
        client.send(str.encode("Got: ") + data)
    print("connection closed.")
if __name__ == '__main__':
    echo_server(('', 25000))
```

相比`_thread`来说，`select` 更加底层，提供了最基础的等待 IO 完成功能。但是缺点是这个功能太单一了，这也就是为什么后面语言提供了 `asyncio`。最早应该是 [python-dev](https://mail.python.org/pipermail/python-ideas/2012-May/015223.html)中提出了要在标准库中添加基于 `select` 的异步 IO 功能。之后 Python 在 3.4 版本之中就加入了 [selectors](https://docs.python.org/3.5/library/selectors.html) 与 [`asyncio](https://docs.python.org/3.5/library/asyncio.html) 库用于异步 IO。

其他的方法还有 `gevent`、`Twisted`、`Tornado` 等等的方案，这里就不多赘述了。(在3.4的时候我一直觉得 `yield form` 太丑陋了，相对我宁愿继续用 `Tornado` 的 `yield` 方式。当然这个更加主观的原因吧，不过现在 `async/await` 方式明显让我又让我爱上了。）

## 3. 从同步到 asyncio

那么如何在`asyncio`框架下如何实现异步`socket`通讯的例子呢？事实上官方文档中提供了两个比较高层封装过的 asyncio 库例子

1. [TCP echo server protocol](https://docs.python.org/3.5/library/asyncio-protocol.html#tcp-echo-server-protocol)

2. [TCP echo server using streams](https://docs.python.org/3.5/library/asyncio-stream.html#tcp-echo-server-using-streams)

这两个例子采用的是 `asyncio`的`socket` 通讯高级别封装，似乎与我们同步代码相差有点远。这里我们实际例子中使用了更加底层的

[Low-level socket operations](https://docs.python.org/3.5/library/asyncio-eventloop.html#low-level-socket-operations)。

这个更接近于我们在同步状态下使用 `socket` 的代码。

```py
# aecho.py
from socket import *
import asyncio

loop = asyncio.get_event_loop()

async def echo_server(address):
    sock = socket(AF_INET, SOCK_STREAM)
    sock.setsockopt(SOL_SOCKET, SO_REUSEADDR, 1)
    sock.bind(address)
    sock.listen(5)
    sock.setblocking(False) # 设置非阻塞
    while True:
        client, addr = await loop.sock_accept(sock)
        print("connect from", addr)
        loop.create_task(echo_handler(client))

async def echo_handler(client):
    with client:
        while True:
            data = await loop.sock_recv(client, 10000)
            if not data:
                break
            await loop.sock_sendall(client, str.encode("Got: ") + data)
    print("connection closed")
loop.create_task(echo_server(('', 25000)))
loop.run_forever()
```

其中遇到的`create_task`会相对同步状态下无法对应，这个方法用于安排一个异步任务的执行，将一个异步方法封装为`future`对象。其他的`Event Loop`中的功能基本与传统的程序相同。

## 4. 从asyncio到自己的实现

那么在`asyncio.event_loop`中到底发生了什么呢？我们可以尝试用自己的程序实现一下。

如果你阅读过[PEP-0492](https://www.python.org/dev/peps/pep-0492/)，你就知道，实际上 Python 的协程是通过生成器实现的。

```py
# async_yield.py
from types import coroutine
@coroutine
def read_wait(sock):
    yield "read_wait", sock  # 为什么有个 read_wait？等下介绍
```

下面来模拟实际调用：

```
python -i async_yield.py
>>> f = read_wait("somesocket")
>>> f
<generator object read_wait at 0x10200d5c8>
>>> f.send(None)
('read_wait', 'somesocket')
>>> f.send(None)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
StopIteration
```

如果不了解`send()`与`StopIteration` 作用的话，请参考 `PEP-0492` 中相关的描述。接下来继续完善 `write` 方法，并且实现我们自己的 `Loop`。

```py
# async_yield.py
from types import coroutine
from collections import deque
from selectors import DefaultSelector, EVENT_READ, EVENT_WRITE
@coroutine
def read_wait(sock):
    yield "read_wait", sock
@coroutine
def write_wait(sock):
    yield "write_wait", sock
class Loop(object):
    def __init__(self):
        self.ready = deque()
        self.selector = DefaultSelector()
    async def sock_recv(self, sock, maxbytes):
        await read_wait(sock)
        return sock.recv(maxbytes)
    async def sock_accept(self, sock):
        await read_wait(sock)
        return sock.accept()
    async def sock_sendall(self, sock, data):
        while data:
            await write_wait(sock)
            n = sock.send(data)
            data = data[n:]
    def create_task(self, coro):
        self.ready.append(coro)
    def run_forever(self):
        while True:
            while not self.ready:
                events = self.selector.select()
                for key, _ in events:
                    self.ready.append(key.data)
                    self.selector.unregister(key.fileobj)
            while self.ready:
                self.current_task = self.ready.popleft()
                try:
                    op, *args = self.current_task.send(None)
                    getattr(self, op)(*args)
                except StopIteration:
                    pass
    def read_wait(self, sock):
        self.selector.register(sock, EVENT_READ, self.current_task)
    def write_wait(self, sock):
        self.selector.register(sock, EVENT_WRITE, self.current_task)
```

对于之前一节中的`aecho.py`文件，我们只需要修改一下导入模块与 loop 的获取方法即可：


```py
# pecho.py
from socket import *
import async_yield
loop = async_yield.Loop()
async def echo_server(address):
    sock = socket(AF_INET, SOCK_STREAM)
    sock.setsockopt(SOL_SOCKET, SO_REUSEADDR, 1)
    sock.bind(address)
    sock.listen(5)
    sock.setblocking(False)  # 设置非阻塞模式
    while True:
        client, addr = await loop.sock_accept(sock)
        print("connect from", addr)
        loop.create_task(echo_handler(client))
async def echo_handler(client):
    with client:
        while True:
            data = await loop.sock_recv(client, 10000)
            if not data:
                break
            await loop.sock_sendall(client, str.encode("Got: ") + data)
    print("connection closed")
loop.create_task(echo_server(('', 25000)))
loop.run_forever()
```

## 5. async_yield 发生了什么？

首先，我们定义了两个协程函数`read_wait`和`write_wait`，分别用于相应处理读取操作与写入操作。其中返回了一个`tuple`类型数据，用于在`op, *args = self.current_task.send(None)`中填充方法名和参数，之后在`getattr(self, op)(*args)`中进行分别调用。

下面Loop类实现了在`pecho`中用到的所有异步函数。初始化时的`self.ready`用于存储协程的调用序列。该序列通过`create_task`添加协程到队列中。

在`run_forever`中，如果目前队列为空，则通过`self.selector.select()`提取一个事件放入队列处理，若队列存在通过`self.current_task.send(None)`通知事件发送，从而调用对应的事件功能。你也可以在`op, *args = self.current_task.send(None)`后添加`print(op)`获取实时的调用情况。

## 结语

事实上这篇文章的思路是基于`@dabeaz`在`Python Brasil`上的 keynote 整理而来。dabeaz 还有另外一个非常不错的基于 select 的异步库，名字叫做curio，是一个了解实现异步库的很好教程。

最后讲个段子，之前有人开玩笑，蟒爹开发一个功能，之后大家都不会正确使用，直到 dabeaz 站出来告诉大家如何正确使用新功能。在写这篇文章的时候虽然很想找出来出处，但是似乎找不到了…