# 真正的并行计算-concurrency.futures模块

参考文章

1. [python异步并发模块concurrent.futures简析](http://lovesoo.org/analysis-of-asynchronous-concurrent-python-module-concurrent-futures.html)
2. [python concurrent.futures](https://www.cnblogs.com/kangoroo/p/7628092.html)
3. [How to kill/cancel/stop running executor future in python ThreadPoolExecutor? future.cancel() is returning False](https://stackoverflow.com/questions/71464095/how-to-kill-cancel-stop-running-executor-future-in-python-threadpoolexecutor-fu)

`concurrency.futures`的两个并发模型: `ProcessPoolExecutor`与`ThreadPoolExecutor`都继承自`Executor`抽象类, 实现了ta的`submit`, `map`和`shutdown`方法. 

`Executor`实现了`__enter__`和`__exit__`使得其对象可以使用`with`操作符.

```py
with ProcessPoolExecutor(max_workers=4) as pool:
    results = list(pool.map(gcd, numbers))
```

查看源码你会发现, 在`__exit__`方法中调用了`self.shutdown()`方法, 所以不难理解, `shutdown`类似于文件操作中的`close`方法.

接下来看看`submit()`方法的使用, 以`ProcessPoolExecutor`模型为例.

## 1. `submit()`方法与`futures`对象

```py
import time
from concurrent.futures import ProcessPoolExecutor

start = time.time()
future_list = []
pool = ProcessPoolExecutor(max_workers=4)
for pair in numbers:
    ## submit函数返回future对象
    future = pool.submit(gcd, pair)
    future_list.append(future)

## 此时任务并未开始执行
print('%.3f seconds.' % (time.time() - start))
## result()方法返回future对象的返回值, 可授受一个timeout参数.
results = [future.result() for future in future_list]
end = time.time()

print(results)
print('Took %.3f seconds.' % (end - start))
```

`submit()`方法返回一个`futures`对象, 类似于一个任务对象, ta提供了跟踪任务执行状态的方法. 如判断任务是否执行中future.running(), 判断任务是否执行完成future.done()等等. 

`future`对象还可以手动取消`cancel()`, 以及添加回调函数. `Executor`对象会在指定`future`对象完成后调用指定回调函数, 同时传入该future对象. 下面我们看看回调函数的用法.

```py
import time
from concurrent.futures import ProcessPoolExecutor

def callback(future):
    '''
    这里无法得到更多信息了, 比如任务参数pair等.
    future对象貌似只绑定了结果队列, 与任务队列已经不相关了.
    '''
    print(future.result())

start = time.time()
future_list = []
pool = ProcessPoolExecutor(max_workers=4)
for pair in numbers:
    ## submit函数返回future对象
    future = pool.submit(gcd, pair)
    ## 注册回调函数
    future.add_done_callback(callback)
    future_list.append(future)

## 此时任务并未开始执行
print('%.3f seconds.' % (time.time() - start))
## result()方法返回future对象的返回值, 可授受一个timeout参数.
results = [future.result() for future in future_list]
end = time.time()

print(results)
print('Took %.3f seconds.' % (end - start))
```

输出

```
0.021 seconds.
5
5
1
1
2
3
[1, 5, 1, 5, 2, 3]
Took 0.577 seconds.
```

注意: 与`result()`方法不同, **回调函数的调用不是有序的**

## 2. `as_completed()`与`wait()`方法

### 2.1 `as_completed()`

`as_completed`方法传入`futures`迭代器和`timeout`两个参数, 返回在`timeout`时间内完成的`futures`迭代器.

默认`timeout=None`，阻塞等待任务执行完成，并返回执行完成的`future`对象迭代器，迭代器是通过`yield`实现的。 

`timeout>0`，等待`timeout`时间，如果`timeout`时间到仍有任务未能完成，不再执行并抛出异常`TimeoutError`

```py
import time
from concurrent.futures import ProcessPoolExecutor, as_completed

start = time.time()
future_list = []
## 为了查看as_completed的效果, 这里指定max_workers为1.
pool = ProcessPoolExecutor(max_workers=1)
for pair in numbers:
    ## submit函数返回future对象
    future = pool.submit(gcd, pair)
    future_list.append(future)

print('=========================')
## 此时任务并未开始执行
for future in future_list:
    print('执行中:%s, 已完成:%s' % (future.running(), future.done()))
print('=========================')
## as_completed返回的与future_list类型相同, 不过是已经完成了的future.
for future in as_completed(future_list, timeout=0.5):
    print('执行中:%s, 已完成:%s' % (future.running(), future.done()))

results = [future.result() for future in future_list]
end = time.time()

print(results)
print('Took %.3f seconds.' % (end - start))
```

执行ta, 输出

```log
=========================
执行中:True, 已完成:False
执行中:False, 已完成:False
执行中:False, 已完成:False
执行中:False, 已完成:False
执行中:False, 已完成:False
执行中:False, 已完成:False
=========================
执行中:False, 已完成:True
执行中:False, 已完成:True
执行中:False, 已完成:True
Traceback (most recent call last):
  File "curr.py", line 36, in <module>
    for future in as_completed(future_list, timeout=0.5):
  File "/usr/local/Cellar/python/3.6.5/Frameworks/Python.framework/Versions/3.6/lib/python3.6/concurrent/futures/_base.py", line 238, in as_completed
    len(pending), total_futures))
concurrent.futures._base.TimeoutError: 3 (of 6) futures unfinished
```

### 2.2 `wait()`

wait方法的参数与`as_completed`相似, 但多了一个`return_when`参数, 返回一个tuple(元组)，tuple中包含两个set(集合)，一个是`completed`(已完成的)另外一个是`uncompleted`(未完成的)。

使用`wait`方法的一个优势就是获得更大的自由度，它接收的第三个参数`return_when`, 可选值有:

- `FIRST_COMPLETED`
- `FIRST_EXCEPTION`
- `ALL_COMPLETED`: 默认

```py
import time
from concurrent.futures import ProcessPoolExecutor, wait, ALL_COMPLETED

start = time.time()
future_list = []
pool = ProcessPoolExecutor(max_workers=1)
for pair in numbers:
    ## submit函数返回future对象
    future = pool.submit(gcd, pair)
    future_list.append(future)

print('=========================')
## 此时任务并未开始执行
for future in future_list:
    print('执行中:%s, 已完成:%s' % (future.running(), future.done()))
print('=========================')
completed, uncompleted = wait(future_list, timeout=0.5, return_when=ALL_COMPLETED)
for d in completed:
    print('执行中:%s, 已完成:%s' % (d.running(), d.done()))
    print(d.result())
print('=========================')
results = [future.result() for future in future_list]
end = time.time()

print(results)
print('Took %.3f seconds.' % (end - start))
```

输出

```log
=========================
执行中:True, 已完成:False
执行中:True, 已完成:False
执行中:False, 已完成:False
执行中:False, 已完成:False
执行中:False, 已完成:False
执行中:False, 已完成:False
=========================
执行中:False, 已完成:True
5
执行中:False, 已完成:True
1
执行中:False, 已完成:True
1
=========================
[1, 5, 1, 5, 2, 3]
Took 0.840 seconds.
```

可以看到, 与`as_completed`不同, `wait`的`timeout`参数不会报异常, 未完成的futures对象会出现在`uncompleted`集合中.

这是`return_when`取`ALL_COMPLETED`的情况, 如果换成`FIRST_COMPLETED`, 将会得到如下结果

```log
=========================
执行中:True, 已完成:False
执行中:True, 已完成:False
执行中:True, 已完成:False
执行中:False, 已完成:False
执行中:False, 已完成:False
执行中:False, 已完成:False
=========================
执行中:False, 已完成:True
1
=========================
[1, 5, 1, 5, 2, 3]
Took 0.893 seconds.
```

这是`wait`在当第一个任务完成时就返回的情况.

## 终止异常线程

20241128

在使用多线程, 尤其是线程池时, 经常会担心某些线程陷入死循环, 或是无法正常结束的状态, 导致`wait()`时无法结束.

但是无法结束反而是好事, 这样表象会很明显, 容易追查. 如果`wait()`时配合了`timeout`参数, 时间到了只处理了`completed`状态的线程结果, 留下那些异常的线程一直在运行, 占用cpu资源, 且会造成线程泄露.

Future对象提供了`cancel()`方法, 比如如下

```py
for item in uncompleted:
    item.cancel()
```

但是对于处于 running 状态的线程, 是无效的. Future对象又没有 thread_id, 根本无法追踪.

在网上翻遍了, 没有解决方案, 开发者需要自行在线程函数内部, 可能会发生死循环的地方埋点, 接收来自外界的信号, 自行退出(类似于 golang 的 context).
