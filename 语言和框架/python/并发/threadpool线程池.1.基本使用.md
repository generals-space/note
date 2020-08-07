# threadpool线程池.1.基本使用

参考文章

1. [python线程池（threadpool）模块使用笔记](http://www.cnblogs.com/xiaozi/p/6182990.html)

2. [threadpool官方文档](https://chrisarndt.de/projects/threadpool/api/)

网络上多数文章的第一个示例都是`threadpool`官方文档中的示例, 如下

```py
from threadpool import ThreadPool, makeRequests, putRequest

pool = ThreadPool(poolsize)
requests = makeRequests(func, list_of_args, callback)
[pool.putRequest(req) for req in requests]
pool.wait()
```

第1行: 定义了一个线程池, 表示最多可以创建`poolsize`个线程. 但是与相像中不同的是, 初始化时的线程数就是`poolsize`, 不会弹性变化;

第2行: 调用`makeRequests`创建了调用线程池处理的请求, 列表类型. 其中包括要调用的函数`func`, 以及函数相关参数`list_of_args`和回调函数`callback`, 其中回调函数可以不写, default是无, 也就是说`makeRequests`只需要2个参数就可以运行；

第3行: 用一个`for`循环将所有要运行多线程的请求扔进线程池, 等同于

```py
for req in requests:
    pool.putRequest(req) 
```

第4行是等待所有的线程完成后退出, 如果没有这一句, 子线程来不及执行主线程就退出了, 类似于`threading`的`join`.

------

真正有点别扭的是第2行, 如果你需要启动`n`个线程执行`func`函数, 那要`list_of_args`中有`n`个成员才行. 如果`list_of_args`是一个空列表, 那么`requests`列表会只有一个成员. 就是说, 同一个函数启动的个数是由`list_of_args`决定的.

...有点像`map`函数, 是吧?

```py
import threadpool
import time
pool = threadpool.ThreadPool(4)

def run(name):
    print('start...')
    time.sleep(10)
    print(name)
name_list = ['AA','BB','CC','DD']
reqs = threadpool.makeRequests(run, name_list)
[pool.putRequest(req) for req in reqs]
pool.wait()
```

执行它, 会得到如下结果

```
start...
start...
start...
start...
## 这里会等待10s
AA
BB
CC
DD
```
