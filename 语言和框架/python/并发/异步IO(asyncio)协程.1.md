# 异步IO(asyncio)协程

参考文章

1. [Python黑魔法 --- 异步IO（ asyncio） 协程](https://www.jianshu.com/p/b5e347b3a17c)

## 1. python asyncio

网络模型有很多种，为了实现高并发也有很多方案，多线程，多进程。无论多线程和多进程，IO的调度更多取决于系统，而协程的方式，调度来自用户，用户可以在函数中`yield`一个状态。使用协程可以实现高效的并发任务。Python的在3.4中引入了协程的概念，可是这个还是以生成器对象为基础，3.5则确定了协程的语法。下面将简单介绍`asyncio`的使用。实现协程的不仅仅是`asyncio`，`tornado`和`gevent`都实现了类似的功能。

- `event_loop` 事件循环：程序开启一个无限的循环，程序员会把一些函数注册到事件循环上。当满足事件发生的时候，调用相应的协程函数。
- `coroutine` 协程：协程对象，指一个使用`async`关键字定义的函数，它的调用不会立即执行函数，而是会返回一个协程对象。**协程对象需要注册到事件循环，由事件循环调用**。
- `task` 任务：一个协程对象就是一个原生可以挂起的函数，任务则是对协程进一步封装，其中包含任务的各种状态。
- `future`： 代表将来执行或没有执行的任务的结果。它和task上没有本质的区别

`async/await 关键字`：`python3.5` 用于定义协程的关键字，`async`定义一个协程，`await`用于挂起阻塞的异步调用接口。

上述的概念单独拎出来都不好懂，比较他们之间是相互联系，一起工作。下面看例子，再回溯上述概念，更利于理解。

## 2. 定义一个协程

定义一个协程很简单，使用async关键字，就像定义普通函数一样：

```py
import time
import asyncio

now = lambda : time.time()

async def do_some_work(x):
    print('Waiting: ', x)

start = now()

coroutine = do_some_work(2)

loop = asyncio.get_event_loop()
loop.run_until_complete(coroutine)

print('TIME: ', now() - start)
```

通过`async`关键字定义一个协程（coroutine），协程也是一种对象。**协程不能直接运行，需要把协程加入到事件循环（loop），由后者在适当的时候调用协程**。

`asyncio.get_event_loop`方法可以创建一个事件循环，然后使用`run_until_complete`将协程注册到事件循环，并启动事件循环。因为本例只有一个协程，于是可以看见如下输出：

```
Waiting:  2
TIME:  0.0004658699035644531
```

## 3. 创建一个task

协程对象不能直接运行，在注册事件循环的时候，其实是`run_until_complete`方法将协程包装成为了一个任务（task）对象。所谓`task`对象是`Future`类的子类。保存了协程运行后的状态，用于未来获取协程的结果。

```py
import asyncio
import time

now = lambda : time.time()

async def do_some_work(x):
    print('Waiting: ', x)

start = now()

coroutine = do_some_work(2)
loop = asyncio.get_event_loop()
# task = asyncio.ensure_future(coroutine)
task = loop.create_task(coroutine)
print(task)
loop.run_until_complete(task)
print(task)
print('TIME: ', now() - start)
```

可以看到输出结果为：

```
<Task pending coro=<do_some_work() running at /Users/ghost/Rsj217/python3.6/async/async-main.py:17>>
Waiting:  2
<Task finished coro=<do_some_work() done, defined at /Users/ghost/Rsj217/python3.6/async/async-main.py:17> result=None>
TIME:  0.0003490447998046875
```

创建`task`后，`task`在加入事件循环之前是`pending`状态，因为`do_some_work`中没有耗时的阻塞操作，`task`很快就执行完毕了。后面打印的`finished`状态。

`asyncio.ensure_future(coroutine)`和`loop.create_task(coroutine)`都可以创建一个`task`，`run_until_complete`的参数是一个`futrue`对象。当传入一个协程，其内部会自动封装成`task`，`task`是`Future`的子类。`isinstance(task, asyncio.Future)`将会输出`True`。

## 4. 绑定回调

绑定回调，在`task`执行完毕的时候可以获取执行的结果，回调的最后一个参数是`future`对象，通过该对象可以获取协程返回值。如果回调需要多个参数，可以通过偏函数导入。

```py
import time
import asyncio

now = lambda : time.time()

async def do_some_work(x):
    print('Waiting: ', x)
    return 'Done after {}s'.format(x)

def callback(future):
    print('Callback: ', future.result())

start = now()

coroutine = do_some_work(2)
loop = asyncio.get_event_loop()
task = asyncio.ensure_future(coroutine)
task.add_done_callback(callback)
loop.run_until_complete(task)

print('TIME: ', now() - start)
```

关键点

```py
def callback(t, future):
    print('Callback:', t, future.result())

task.add_done_callback(functools.partial(callback, 2))
```

可以看到，`coroutine`执行结束时候会调用回调函数。并通过参数`future`获取协程执行的结果。我们创建的`task`和回调里的`future`对象，实际上是同一个对象。

## 5. future 与 result

回调一直是很多异步编程的恶梦，程序员更喜欢使用同步的编写方式写异步代码，以避免回调的恶梦。回调中我们使用了`future`对象的`result`方法。前面不绑定回调的例子中，我们可以看到`task`有`fiinished`状态。在那个时候，可以直接读取`task`的`result()`方法。

```py
import time
import asyncio

now = lambda : time.time()

async def do_some_work(x):
    print('Waiting {}'.format(x))
    return 'Done after {}s'.format(x)

start = now()

coroutine = do_some_work(2)
loop = asyncio.get_event_loop()
task = asyncio.ensure_future(coroutine)
loop.run_until_complete(task)

print('Task ret: {}'.format(task.result()))
print('TIME: {}'.format(now() - start))
```

可以看到输出的结果：

```py
Waiting:  2
Task ret:  Done after 2s
TIME:  0.0003650188446044922
```
