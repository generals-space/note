# 多线程threading库使用(二)锁Lock与RLock

参考文章

1. [Python模块学习：threading 多线程控制和处理](http://python.jobbole.com/81546/)

2. [Java并发编程——锁与可重入锁](http://www.jianshu.com/p/007bd7029faf)

上一篇文档中简单的使用了一个`threading.Lock`类, 在操作`count`共享变量时, 首先获得这个锁, 操作完成后再释放. 这也是线程锁的基本使用方法.

其实在`threading`模块中, 定义两种类型的锁: `threading.Lock`和`threading.RLock`. 它们之间有一点细微的区别, 通过比较下面两段代码来说明: 

```py
import threading
lock = threading.Lock()	#Lock对象
lock.acquire()
lock.acquire()  #产生了死锁. 
lock.release()
lock.release()
```

```py
import threading
rLock = threading.RLock()  #RLock对象
rLock.acquire()
rLock.acquire()	#在同一线程内, 程序不会堵塞. 
rLock.release()
rLock.release()
```

这两种锁的主要区别是: `RLock`允许在同一线程中被多次`acquire`. 而`Lock`却不允许这种情况. 注意: 如果使用`RLock`, 那么`acquire`和`release`必须成对出现, 不可嵌套.

> `RLock`学名为`可重入锁`, 而`Lock`这种不可重入锁, 有些地方称之为`自旋锁(linux内核)`, 更常见的是叫做`互斥锁`.

可重入意味着: 线程可以进入任何一个它 **已经拥有的锁**所同步着的代码块. 如下示例

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

可重入锁的原理是, `RLock`内部维护着一个`Lock`和一个`counter`变量, `counter`记录了`acquire`的次数, 从而使得资源可以被多次`require`. 直到一个线程所有的`acquire`都被`release`, 其他的线程才能获得资源. 

也就是说, 它其实还是只获得了一次锁, 但却可以修正普通锁的放置位置问题(到底应该把锁加在调用函数中, 还是放在主调函数体中...)
