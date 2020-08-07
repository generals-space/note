# 多线程threading库使用(四)Event

参考文章

1. [Python模块学习：threading 多线程控制和处理](http://python.jobbole.com/81546/)

2. [Java并发编程——锁与可重入锁](http://www.jianshu.com/p/007bd7029faf)

3. [Python中使用threading.Event协调线程的运行](http://blog.csdn.net/cnweike/article/details/40821283)

`Event`实现与`Condition`类似的功能，不过比`Condition`简单一点. 它通过维护内部的标识符来实现线程间的同步问题. (`threading.Event`和`.NET`中的`System.Threading.ManualResetEvent`类实现同样的功能. )

- `Event.wait([timeout])`: 堵塞线程，直到`Event`对象内部标识位被设为`True`或超时(如果提供了参数`timeout`). 
- `Event.set()`: 将标识位设为`Ture`
- `Event.clear()`: 将标识伴设为`False`
- `Event.isSet()`: 判断标识位是否为`Ture`

`Event`与`Condition`类似, 但没有获得与释放的操作. 

`threading.Event`机制类似于一个线程向其它多个线程发号施令的模式，其它线程都会持有一个`threading.Event`的对象，这些线程都会等待这个事件的“发生”，如果此事件一直不发生，那么这些线程将会阻塞，直至事件的“发生”. 

下面是用`Event`实现的捉迷藏游戏, 有点问题, 暂时还没想清楚.<???>

```py
#!encoding: utf-8
import threading, time
class Seeker(threading.Thread):
    def __init__(self, cond, name):
        super(Seeker, self).__init__()
        self.cond = cond
        self.name = name
    
    def run(self):
        print self.name + ': 我已经把眼睛蒙上了'
        self.cond.set()         ## 发出通知, hider开始运行
        self.cond.wait()        ## 这里再次wait貌似是无意义的, 为什么???
        print self.name + ': 我找到你了 ~_~'
        print self.name + ': 我赢了'
        
class Hider(threading.Thread):
    def __init__(self, cond, name):
        super(Hider, self).__init__()
        self.cond = cond
        self.name = name
    def run(self):
        self.cond.wait()        ## 先挂起等待
        print self.name + ': 我已经藏好了，你快来找我吧'
        
cond = threading.Event()
seeker = Seeker(cond, 'seeker')
hider = Hider(cond, 'hider')
seeker.start()
hider.start()
```

上面的代码执行结果如下

```
$ python test.py 
seeker: 我已经把眼睛蒙上了
seeker: 我找到你了 ~_~
seeker: 我赢了
hider: 我已经藏好了，你快来找我吧
```

可以看到hider最初的`wait`是成功的, seeker在"蒙眼"操作后的通知`set`也是成功的, 但是seeker本身再次`wait`就无效了. 

暂时还无法理解这一点, 也许是`Event`对象就不是这么用的? 只能一次性通知?