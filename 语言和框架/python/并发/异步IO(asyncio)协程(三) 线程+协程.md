# 异步IO(asyncio)协程(三) 线程+协程

参考文章

1. [Python黑魔法 --- 异步IO（ asyncio） 协程](https://www.jianshu.com/p/b5e347b3a17c)

## 1. 不同线程的事件循环

很多时候，我们的事件循环用于注册协程，而有的协程需要动态的添加到事件循环中。一个简单的方式就是使用多线程。当前线程创建一个事件循环，然后在新建一个线程，在新线程中启动事件循环。当前线程不会被block。

```py
from threading import Thread

def start_loop(loop):
    asyncio.set_event_loop(loop)
    loop.run_forever()

def more_work(x):
    print('More work {}'.format(x))
    time.sleep(x)
    print('Finished more work {}'.format(x))

start = now()
new_loop = asyncio.new_event_loop()
t = Thread(target=start_loop, args=(new_loop,))
t.start()
print('TIME: {}'.format(time.time() - start))

new_loop.call_soon_threadsafe(more_work, 6)
new_loop.call_soon_threadsafe(more_work, 3)
```

启动上述代码之后，当前线程不会被block，新线程中会按照顺序执行`call_soon_threadsafe`方法注册的`more_work`方法，后者因为`time.sleep`操作是同步阻塞的，因此运行完毕`more_work`需要大致6 + 3.

## 2. 新线程协程

```py
import asyncio
import time
from threading import Thread

def start_loop(loop):
    asyncio.set_event_loop(loop)
    loop.run_forever()

async def do_some_work(x):
    print('Waiting {}'.format(x))
    await asyncio.sleep(x)
    print('Done after {}s'.format(x))

def more_work(x):
    print('More work {}'.format(x))
    time.sleep(x)
    print('Finished more work {}'.format(x))

start = time.time()
new_loop = asyncio.new_event_loop()
t = Thread(target=start_loop, args=(new_loop,))
t.start()
print('TIME: {}'.format(time.time() - start))

asyncio.run_coroutine_threadsafe(do_some_work(6), new_loop)
asyncio.run_coroutine_threadsafe(do_some_work(4), new_loop)
```

上述的例子，主线程中创建一个`new_loop`，然后在另外的子线程中开启一个无限事件循环。主线程通过`run_coroutine_threadsafe`新注册协程对象。这样就能在子线程中进行事件循环的并发操作，同时主线程又不会被block。一共执行的时间大概在6s左右。

输出为

```
TIME: 0.0010180473327636719
Waiting 6
Waiting 4
Done after 4s
Done after 6s
## 这里会卡住
```

协程执行完毕后会卡住, ctrl-c也不管用, 无法结束, 只能关闭当前终端. 因为协程的`run_forever()`使得子线程保持运行, 又没有设置`setDaemon(True)`选项, 主线程的输入无法传入给子线程.

## 3. master-worker主从模式

对于并发任务，通常是用生成消费模型，对队列的处理可以使用类似`master-worker`的方式，master主要用户获取队列的msg，worker用户处理消息。

为了简单起见，并且协程更适合单线程的方式，我们的主线程用来监听队列，子线程用于处理队列。这里使用redis的队列。主线程中有一个是无限循环，用户消费队列。

```py
while True:
    task = rcon.rpop("queue")
    if not task:
        time.sleep(1)
        continue
    asyncio.run_coroutine_threadsafe(do_some_work(int(task)), new_loop)
```

给队列添加一些数据：

```
127.0.0.1:6379[3]> lpush queue 2
(integer) 1
127.0.0.1:6379[3]> lpush queue 5
(integer) 1
127.0.0.1:6379[3]> lpush queue 1
(integer) 1
127.0.0.1:6379[3]> lpush queue 1
```

可以看见输出：

```
Waiting  2
Done 2
Waiting  5
Waiting  1
Done 1
Waiting  1
Done 1
Done 5
```

我们发起了一个耗时5s的操作，然后又发起了连个1s的操作，可以看见子线程并发的执行了这几个任务，其中5s awati的时候，相继执行了1s的两个任务。

## 4. 停止子线程

如果一切正常，那么上面的例子很完美。可是，需要停止程序，直接ctrl+c，会抛出`KeyboardInterrupt`错误，我们修改一下主循环：

```py
try:
    while True:
        task = rcon.rpop("queue")
        if not task:
            time.sleep(1)
            continue
        asyncio.run_coroutine_threadsafe(do_some_work(int(task)), new_loop)
except KeyboardInterrupt as e:
    print(e)
    new_loop.stop()
```

可是实际上并不好使，虽然主线程`try`了`KeyboardInterrupt`异常，但是子线程并没有退出，为了解决这个问题，可以设置子线程为守护线程，这样当主线程结束的时候，子线程也随机退出。

```py
new_loop = asyncio.new_event_loop()
t = Thread(target=start_loop, args=(new_loop,))
t.setDaemon(True)    # 设置子线程为守护线程
t.start()

try:
    while True:
        # print('start rpop')
        task = rcon.rpop("queue")
        if not task:
            time.sleep(1)
            continue
        asyncio.run_coroutine_threadsafe(do_some_work(int(task)), new_loop)
except KeyboardInterrupt as e:
    print(e)
    new_loop.stop()
```

线程停止程序的时候，主线程退出后，子线程也随机退出才了，并且停止了子线程的协程任务。

