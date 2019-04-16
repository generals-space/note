# 多线程threading库使用(二)锁

参考文章

1. [Python模块学习：threading 多线程控制和处理](http://python.jobbole.com/81546/)

2. [Java并发编程——锁与可重入锁](http://www.jianshu.com/p/007bd7029faf)

3. [Python中使用threading.Event协调线程的运行](http://blog.csdn.net/cnweike/article/details/40821283)

上一篇文档中简单的使用了一个`threading.Lock`类, 在操作`count`共享变量时, 首先获得这个锁, 操作完成后再释放. 这也是线程锁的基本使用方法.

## 1. Lock与RLock

其实在`threading`模块中，定义两种类型的琐：`threading.Lock`和`threading.RLock`。它们之间有一点细微的区别，通过比较下面两段代码来说明：

```py
import threading
lock = threading.Lock()	#Lock对象
lock.acquire()
lock.acquire()  #产生了死琐。
lock.release()
lock.release()
```

```py
import threading
rLock = threading.RLock()  #RLock对象
rLock.acquire()
rLock.acquire()	#在同一线程内，程序不会堵塞。
rLock.release()
rLock.release()
```

这两种琐的主要区别是：`RLock`允许在同一线程中被多次`acquire`。而`Lock`却不允许这种情况。注意：如果使用`RLock`，那么`acquire`和`release`必须成对出现, 不可嵌套.

> `RLock`学名为`可重入锁`, 而`Lock`这种不可重入锁, 有些地方称之为`自旋锁`.

可重入意味着：线程可以进入任何一个它**已经拥有的锁**所同步着的代码块。如下示例

```py
lock = RLock()

def outer():
    lock.lock
    inner()
    lock.unlock()

def inner():
    lock.lock()
    ## do something
    lock.unlock()
```

还是容易理解的.

可重入锁的原理是, `RLock`内部维护着一个`Lock`和一个`counter`变量，`counter`记录了`acquire`的次数，从而使得资源可以被多次`require`。直到一个线程所有的`acquire`都被`release`，其他的线程才能获得资源。

也就是说, 它其实还是只获得了一次锁, 但却可以修正普通锁的放置位置问题(到底应该把锁加在调用函数中, 还是放在主调函数体中...)

## 2. Condition

可以把`Condiftion`理解为一把高级的琐，它提供了比`Lock`, `RLock`更高级的功能，允许我们能够控制复杂的线程同步问题。

`threadiong.Condition`在内部维护一个琐对象（默认是`RLock`），可以在创建`Condigtion`对象的时候把琐对象作为参数传入。

`Condition`也提供了`acquire`, `release`方法，其含义与常规琐的`acquire`, `release`方法一致，其实它只是简单的调用内部琐对象的对应的方法而已。`Condition`还提供了如下方法(特别要注意：这些方法只有在占用琐(`acquire`)之后才能调用，否则将会报`RuntimeError`异常。).

`Condition`类的常用方法:

`Condition.wait([timeout])`

`wait`方法, 暂时释放内部所占用的琐，同时线程被挂起，直至接收到通知被唤醒或超时（如果提供了`timeout`参数的话）。当线程被唤醒并**重新占有琐**的时候，程序才会继续执行下去。

`Condition.notify()`

唤醒一个挂起的线程（如果存在挂起的线程）。注意：`notify()`方法不会释放所占用的琐。

`Condition.notify_all()`

`Condition.notifyAll()`

唤醒所有挂起的线程（如果存在挂起的线程）。注意：这些方法不会释放所占用的琐。

------

`Condition`类的应用场景为, 线程之间存在**时序关系**. 简单来说, 就是线程需要能够进行简单的相互通信, 同时能够在某些条件下挂起/唤醒.

如下示例, 模拟了一个捉迷藏的游戏.

假设这个游戏由两个人来玩，一个藏(Hider)，一个找(Seeker)。游戏的规则如下：

1. 游戏开始之后，Seeker先把自己眼睛蒙上，蒙上眼睛后，就通知Hider；

2. Hider接收通知后开始找地方将自己藏起来，藏好之后，再通知Seeker可以找了； 

3. Seeker接收到通知之后，就开始找Hider。

Hider和Seeker都是独立的个体(线程)，在游戏过程中，两者之间的行为有一定的时序关系，可以通过Condition来控制。

```py
#!encoding: utf-8
import threading, time
class Seeker(threading.Thread):
    def __init__(self, cond, name):
        super(Seeker, self).__init__()
        self.cond = cond
        self.name = name
    
    def run(self):
        self.cond.acquire()
        print self.name + ': 我已经把眼睛蒙上了'
        self.cond.notify()  ## 唤醒挂起的线程hider, 然后挂起自身
        self.cond.wait()
        print self.name + ': 我找到你了 ~_~'
        self.cond.notify()
        self.cond.release()
        print self.name + ': 我赢了'

class Hider(threading.Thread):
    def __init__(self, cond, name):
        super(Hider, self).__init__()
        self.cond = cond
        self.name = name
    def run(self):
        self.cond.acquire()
        self.cond.wait()    ## 挂起的同时释放锁, 所以就算Hider对象先运行, 也没关系, 它需要一个notify通知
        print self.name + ': 我已经藏好了，你快来找我吧'
        self.cond.notify()
        self.cond.wait()
        self.cond.release() 
        print self.name + ': 被你找到了，哎~~~'

cond = threading.Condition()
seeker = Seeker(cond, 'seeker')
hider = Hider(cond, 'hider')
seeker.start()
hider.start()
```

注意, 在对`Condition`对象进行`wait`, `notify`等操作时, 需要首先获得(`acquire`)锁. 当不再需要这个对象进行挂起/唤醒操作时, 要像普通锁一样释放.

## 3. Event

`Event`实现与`Condition`类似的功能，不过比`Condition`简单一点。它通过维护内部的标识符来实现线程间的同步问题。（`threading.Event`和`.NET`中的`System.Threading.ManualResetEvent`类实现同样的功能。）

`Event.wait([timeout])`

堵塞线程，直到`Event`对象内部标识位被设为`True`或超时（如果提供了参数`timeout`）。

`Event.set()`

将标识位设为`Ture`

`Event.clear()`

将标识伴设为`False`

`Event.isSet()`

判断标识位是否为`Ture`

`Event`与`Condition`类似, 但没有获得与释放的操作. 

`threading.Event`机制类似于一个线程向其它多个线程发号施令的模式，其它线程都会持有一个`threading.Event`的对象，这些线程都会等待这个事件的“发生”，如果此事件一直不发生，那么这些线程将会阻塞，直至事件的“发生”。

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