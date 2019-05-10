# 多线程threading库使用(三)Condition

参考文章

1. [Python模块学习：threading 多线程控制和处理](http://python.jobbole.com/81546/)

2. [Java并发编程——锁与可重入锁](http://www.jianshu.com/p/007bd7029faf)

3. [Python中使用threading.Event协调线程的运行](http://blog.csdn.net/cnweike/article/details/40821283)

可以把`Condiftion`理解为一把高级的锁, 它提供了比`Lock`, `RLock`更高级的功能, 允许我们能够控制复杂的线程同步问题. 

`threadiong.Condition`在内部维护一个锁对象(默认是`RLock`), 可以在创建`Condigtion`对象的时候把锁对象作为参数传入. 

`Condition`也提供了`acquire`, `release`方法, 其含义与常规锁的`acquire`, `release`方法一致, 其实它只是简单的调用内部锁对象的对应的方法而已. `Condition`还提供了如下方法:

- `Condition.wait([timeout])`: 暂时释放内部所占用的锁, 同时线程被挂起, 直至接收到通知被唤醒或超时(如果提供了`timeout`参数的话). 当线程被唤醒并**重新占有锁**的时候, 程序才会继续执行下去. 
- `Condition.notify()`: 唤醒一个挂起的线程(如果存在挂起的线程). 注意: `notify()`方法不会释放所占用的锁. 
- `Condition.notify_all()`: 唤醒所有挂起的线程(如果存在挂起的线程). 注意: 这些方法不会释放所占用的锁. 

> **注意**: 这些方法只有在占用锁(`acquire`)之后才能调用, 否则将会报`RuntimeError`异常.

`Condition`类的应用场景为, 线程之间存在 **时序关系**. 简单来说, 就是线程需要能够进行简单的相互通信, 同时能够在某些条件下挂起/唤醒.

如下示例, 模拟了一个捉迷藏的游戏.

假设这个游戏由两个人来玩, 一个藏(Hider), 一个找(Seeker). 游戏的规则如下: 

1. 游戏开始之后, Seeker先把自己眼睛蒙上, 蒙上眼睛后, 就通知Hider; 
2. Hider接收通知后开始找地方将自己藏起来, 藏好之后, 再通知Seeker可以找了; 
3. Seeker接收到通知之后, 就开始找Hider. 

Hider和Seeker都是独立的个体(线程), 在游戏过程中, 两者之间的行为有一定的时序关系, 可以通过Condition来控制. 

```py
#!encoding: utf-8
import time
import threading
class Seeker(threading.Thread):
    def __init__(self, cond, name):
        super(Seeker, self).__init__()
        self.cond = cond
        self.name = name
    
    def run(self):
        ## Condition对象内部的锁为RLock, 所以可以在Seeker与Hider中同时acquire
        self.cond.acquire()
        print(self.name + ': 我已经把眼睛蒙上了')
        ## 唤醒挂起的线程hider, 然后挂起自身, 等待Hider藏好
        self.cond.notify() 
        
        self.cond.wait()

        print(self.name + ': 我找到你了 ~_~')
        self.cond.notify()
        self.cond.release()
        print(self.name + ': 我赢了')

class Hider(threading.Thread):
    def __init__(self, cond, name):
        super(Hider, self).__init__()
        self.cond = cond
        self.name = name

    def run(self):
        ## Condition对象内部的锁为RLock, 所以可以在Seeker与Hider中同时acquire
        self.cond.acquire()
        ## 挂起的同时释放锁, 它需要一个notify通知, 等待Seeker把眼蒙上
        self.cond.wait() 

        print(self.name + ': 我已经藏好了, 你快来找我吧')
        self.cond.notify()
        
        self.cond.wait()
        
        self.cond.release() 
        print(self.name + ': 被你找到了, 哎~~~')

cond = threading.Condition()
seeker = Seeker(cond, 'seeker')
hider = Hider(cond, 'hider')
hider.start()
seeker.start()

```

注意: 在对`Condition`对象进行`wait`, `notify`等操作时, 需要首先`acquire`锁. 当不再需要这个对象进行挂起/唤醒操作时, 要像普通锁一样`release`.

执行结果

```
seeker: 我已经把眼睛蒙上了
hider: 我已经藏好了, 你快来找我吧
seeker: 我找到你了 ~_~
seeker: 我赢了
hider: 被你找到了, 哎~~~
```
