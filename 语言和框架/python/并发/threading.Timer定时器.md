# threading.Timer定时器

参考文章

1. [Python 线程（六）：Timer(定时器)](https://www.cnblogs.com/wang-can/p/3582051.html)

`Timer`是`Threading`的子类, 一定时间后调用指定函数. 但是这只类似于js中的`setTimeout`, 如果想实现每隔一段时间就调用一次函数的话, 就要在Timer调用的函数中, 再次设置`Timer`.

## 单次调用

```py
import threading
from datetime import datetime

def greet(name = ''):
    print(datetime.now(), name)

print(datetime.now(), 'start')
timer = threading.Timer(2, greet, kwargs={'name': 'general'})
timer.run()

```

执行结果

```
2019-06-05 23:55:35.406373 start
2019-06-05 23:55:37.406543 general
```

> 貌似`Timer`对象的`run()`与`start()`作用一样?

## 循环调用

```py
import threading
from datetime import datetime

def greet(name = ''):
    print(datetime.now(), name)
    global timer
    timer = threading.Timer(2, greet, kwargs={'name': name})
    timer.start()

print(datetime.now(), 'start')
timer = threading.Timer(2, greet, kwargs={'name': 'general'})
timer.run()

```

执行结果(需要手动ctrl-c停止)

```
2019-06-05 23:58:38.580981 start
2019-06-05 23:58:40.581260 general
2019-06-05 23:58:42.582343 general
2019-06-05 23:58:44.583403 general
```

## 关于取消

作为`Thread`的子类, `Timer`拥有线程对象的全部方法. ta还多了一个`cancel()`方法, 作为终止的方法.

~~如果定时器要执行的函数需要运行较长的时间, 就可以调用这个方法手动取消.~~

好吧我错了, `cancel()`不能强制停止一个正在运行的`Timer`函数. 如下

```py
import threading
import time
from datetime import datetime

def stop_loop():
    ## 5s后结束循环
    time.sleep(5)
    global timer
    timer.cancel()
    print(datetime.now(), 'stop')

def loop():
    while True:
        print(datetime.now(), 'working')
        time.sleep(2)

print(datetime.now(), 'start')
thread = threading.Thread(target=stop_loop)
thread.start()

timer = threading.Timer(2, loop)
timer.run()

```

执行结果(需要手动ctrl-c停止, `cancel()`完全无效)

```
2019-06-06 00:13:48.466526 start
2019-06-06 00:13:50.468301 working
2019-06-06 00:13:52.468327 working
2019-06-06 00:13:53.468300 stop
2019-06-06 00:13:54.469059 working
2019-06-06 00:13:56.469094 working
```

按照文档所说, `cancel()`只能解除等待执行的定时器, 就是说, 在定时器没有运行目标函数前取消. 上面的代码可以修改为

```py
import threading
import time
from datetime import datetime

def stop_greet():
    ## 2s后结束监听
    time.sleep(2)
    global timer
    timer.cancel()
    print(datetime.now(), 'stop')

def greet():
    print(datetime.now(), 'working')

print(datetime.now(), 'start')
thread = threading.Thread(target=stop_greet)
thread.start()

timer = threading.Timer(5, greet)
timer.run()

```

执行结果为

```
2019-06-06 00:17:35.579713 start
2019-06-06 00:17:37.582705 stop
```

可以看到, 目标函数`greet`没机会执行就终止了.

## 总结

总感觉`Timer`在运行的时候像执行了`join`一样, 可以阻塞等待子线程结束???
