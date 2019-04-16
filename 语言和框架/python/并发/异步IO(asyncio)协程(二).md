# 异步IO(asyncio)协程(二)

参考文章

1. [Python黑魔法 --- 异步IO（ asyncio） 协程](https://www.jianshu.com/p/b5e347b3a17c)

## 1. 阻塞和await

使用`async`可以定义协程对象，使用`await`可以针对耗时的操作进行挂起，就像生成器里的`yield`一样，函数让出控制权。协程遇到`await`，事件循环将会挂起该协程，执行别的协程，直到其他的协程也挂起或者执行完毕，再进行下一个协程的执行。

耗时的操作一般是一些IO操作，例如网络请求，文件读取等。我们使用`asyncio.sleep`函数来模拟IO操作。协程的目的也是让这些IO操作异步化。

```
import asyncio
import time

now = lambda: time.time()

async def do_some_work(x):
    print('Waiting: ', x)
    await asyncio.sleep(x)
    return 'Done after {}s'.format(x)

start = now()

coroutine = do_some_work(2)
loop = asyncio.get_event_loop()
task = asyncio.ensure_future(coroutine)
loop.run_until_complete(task)

print('Task ret: ', task.result())
print('TIME: ', now() - start)
```

在`sleep`的时候，使用`await`让出控制权。即当遇到阻塞调用的函数的时候，使用`await`方法将协程的控制权让出，以便`loop`调用其他的协程。现在我们的例子就用耗时的阻塞操作了。

## 2. 并发和并行

并发和并行一直是容易混淆的概念。并发通常指有多个任务需要同时进行，并行则是同一时刻有多个任务执行。用上课来举例就是，并发情况下是一个老师在同一时间段辅助不同的人功课。并行则是好几个老师分别同时辅助多个学生功课。简而言之就是一个人同时吃三个馒头还是三个人同时分别吃一个的情况，吃一个馒头算一个任务。

`asyncio`实现并发，就需要多个协程来完成任务，每当有任务阻塞的时候就`await`，然后其他协程继续工作。创建多个协程的列表，然后将这些协程注册到事件循环中。

```py
import asyncio
import time

now = lambda: time.time()

async def do_some_work(x):
    print('Waiting: ', x)

    await asyncio.sleep(x)
    return 'Done after {}s'.format(x)

start = now()

coroutine1 = do_some_work(1)
coroutine2 = do_some_work(2)
coroutine3 = do_some_work(4)

tasks = [
    asyncio.ensure_future(coroutine1),
    asyncio.ensure_future(coroutine2),
    asyncio.ensure_future(coroutine3)
]

loop = asyncio.get_event_loop()
loop.run_until_complete(asyncio.wait(tasks))

for task in tasks:
    print('Task ret: ', task.result())

print('TIME: ', now() - start)
```

结果如下

```
Waiting:  1
Waiting:  2
Waiting:  4
Task ret:  Done after 1s
Task ret:  Done after 2s
Task ret:  Done after 4s
TIME:  4.003541946411133
```

总时间为4s左右。4s的阻塞时间，足够前面两个协程执行完毕。如果是同步顺序的任务，那么至少需要7s。此时我们使用了`aysncio`实现了并发。`asyncio.wait(tasks)`, 也可以使用`asyncio.gather(*tasks)` ,前者接受一个`task`列表，后者接收一堆`task`。

## 3. 协程嵌套

使用`async`可以定义协程，协程用于耗时的io操作，我们也可以封装更多的io操作过程，这样就实现了嵌套的协程，即一个协程中`await`了另外一个协程，如此连接起来。

```py
import asyncio
import time

now = lambda: time.time()

async def do_some_work(x):
    print('Waiting: ', x)

    await asyncio.sleep(x)
    return 'Done after {}s'.format(x)

async def main():
    coroutine1 = do_some_work(1)
    coroutine2 = do_some_work(2)
    coroutine3 = do_some_work(4)

    tasks = [
        asyncio.ensure_future(coroutine1),
        asyncio.ensure_future(coroutine2),
        asyncio.ensure_future(coroutine3)
    ]

    dones, pendings = await asyncio.wait(tasks)

    for task in dones:
        print('Task ret: ', task.result())

start = now()

loop = asyncio.get_event_loop()
loop.run_until_complete(main())

print('TIME: ', now() - start)
```

## 4. 协程停止

上面见识了协程的几种常用的用法，都是协程围绕着事件循环进行的操作。future对象有几个状态：

- Pending

- Running

- Done

- Cancelled

创建future的时候，task为`pending`，事件循环调用执行的时候当然就是`running`，调用完毕自然就是`done`，如果需要停止事件循环，就需要先把task取消。可以使用`asyncio.Task`获取事件循环的task.

```py
import asyncio
import time

now = lambda: time.time()

async def do_some_work(x):
    print('Waiting: ', x)

    await asyncio.sleep(x)
    return 'Done after {}s'.format(x)

coroutine1 = do_some_work(1)
coroutine2 = do_some_work(2)
coroutine3 = do_some_work(2)

tasks = [
    asyncio.ensure_future(coroutine1),
    asyncio.ensure_future(coroutine2),
    asyncio.ensure_future(coroutine3)
]

start = now()

loop = asyncio.get_event_loop()
try:
    loop.run_until_complete(asyncio.wait(tasks))
except KeyboardInterrupt as e:
    print(asyncio.Task.all_tasks())
    for task in asyncio.Task.all_tasks():
        print(task.cancel())
    loop.stop()
    loop.run_forever()
finally:
    loop.close()

print('TIME: ', now() - start)
```

启动事件循环之后，马上ctrl+c，会触发`run_until_complete`的执行异常 `KeyBorardInterrupt`。然后通过循环`asyncio.Task`取消`future`。可以看到输出如下：

```
Waiting:  1
Waiting:  2
Waiting:  2
{<Task pending coro=<do_some_work() running at ...
True
True
True
True
TIME:  0.8858370780944824
```

True表示cannel成功，`loop stop`之后还需要再次开启事件循环，最后在close，不然还会抛出异常：

```
Task was destroyed but it is pending!
task: <Task pending coro=<do_some_work() done,
```

循环task，逐个`cancel`是一种方案，可是正如上面我们把task的列表封装在main函数中，main函数外进行事件循环的调用。这个时候，main相当于最外出的一个task，那么处理包装的main函数即可。

```py
import asyncio
import time

now = lambda: time.time()

async def do_some_work(x):
    print('Waiting: ', x)

    await asyncio.sleep(x)
    return 'Done after {}s'.format(x)

async def main():
    coroutine1 = do_some_work(1)
    coroutine2 = do_some_work(2)
    coroutine3 = do_some_work(2)

    tasks = [
        asyncio.ensure_future(coroutine1),
        asyncio.ensure_future(coroutine2),
        asyncio.ensure_future(coroutine3)
    ]
    done, pending = await asyncio.wait(tasks)
    for task in done:
        print('Task ret: ', task.result())

start = now()

loop = asyncio.get_event_loop()
task = asyncio.ensure_future(main())
try:
    loop.run_until_complete(task)
except KeyboardInterrupt as e:
    print(asyncio.Task.all_tasks())
    print(asyncio.gather(*asyncio.Task.all_tasks()).cancel())
    loop.stop()
    loop.run_forever()
finally:
    loop.close()
```