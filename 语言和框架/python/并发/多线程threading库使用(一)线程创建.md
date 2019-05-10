# 多线程threading库使用

参考文章

1. [Python模块学习：threading 多线程控制和处理](http://python.jobbole.com/81546/)

`threading`是比`thread`更高级的库, 封装的更好, 接口也更简单.

## 1. 创建线程 - threading.Thread类

`Thread`是`threading`模块中最重要的类之一, 可以使用它来创建线程. 常用有两种方式:

1. 通过继承`Thread`类, 创建一个子类, 重写它的`run`方法; 
2. 创建一个`threading.Thread`对象, 在它的初始化函数`__init__`中将可调用对象作为参数传入. 

下面分别举例说明. 

### 1.1 创建`threading.Thread`子类

```py
#!encoding: utf-8
import threading, time, random
count = 0
class Counter(threading.Thread):
    def __init__(self, lock, threadName):
        '''
        @function: 初始化对象. 
        @param lock: 锁对象. 
        @param threadName: 线程名称. 
        '''
        # 注意: 一定要显式的调用父类的构造函数
        super(Counter, self).__init__(name = threadName)  
        self.lock = lock

    def run(self):
        '''
        @function: 重写父类run方法, 在线程启动后执行该方法内的代码. 
        '''
        global count
        self.lock.acquire()
        for i in xrange(10000):
            count = count + 1
        self.lock.release()

lock = threading.Lock()
for i in range(5): 
    Counter(lock, "thread-" + str(i)).start()

# 确保线程都执行完毕
time.sleep(2)	
print count
```

在代码中, 我们创建了一个`Counter`类, 它继承了`threading.Thread`. 初始化函数接收两个参数, 一个是琐对象, 另一个是线程的名称. 在`Counter`中, 重写了从父类继承的`run`方法, `run`方法将一个全局变量逐一的增加10000. 在接下来的代码中, 创建了五个`Counter`对象, 分别调用其`start`方法. 最后打印结果. 

这里要说明一下`run`方法和`start`方法, 它们都是从`Thread`父类继承而来的.

- `run()`方法将在线程开启后执行, 可以把相关的逻辑写到`run`方法中（通常把`run`方法称为活动[Activity]. ）; 
- `start()`方法用于启动线程. 

### 1.2 传入

再看看另外一种创建线程的方法: 

```py
#!encoding: utf-8
import threading, time, random
count = 0
lock = threading.Lock()
def counter():
    '''
    @function: 将全局变量count 逐一的增加10000. 
    '''
    global count, lock
    lock.acquire()
    for i in xrange(10000):
        count = count + 1
    lock.release()
for i in range(5):
    threading.Thread(target = counter, args = (), name = 'thread-' + str(i)).start()
time.sleep(2)
#确保线程都执行完毕
print count
```

在这段代码中, 我们定义了方法`counter`, 它将全局变量`count`逐一的增加10000. 然后创建了5个`Thread`对象, 把函数对象`counter`作为参数传给它的初始化函数, 再调用`Thread`对象的`start`方法, 线程启动后将执行counter函数. 

这里有必要介绍一下`threading.Thread`类的初始化函数原型: 

```py
def __init__(self, group=None, target=None, name=None, args=(), kwargs={})
```

1. `group`是预留的, 用于将来扩展; 
2. `target`是一个可调用对象（也称为活动[activity]）, 在线程启动后执行, 上面代码中指的就是`counter`函数; 
3. `name`是线程的名字. 默认值为`Thread-N`, `N`是一个数字. 可自定义;
4. `args`和`kwargs`分别表示调用`target`时的参数列表和字典;

## 2. Threading类成员属性

`Thread`类还定义了以下常用方法与属性: 

### 1.

`Thread.getName()`

`Thread.setName()`

`Thread.name`

用于获取和设置线程的名称。

### 2.

`Thread.ident`

获取线程的标识符。线程标识符是一个非零整数，只有在调用了`start()`方法之后该属性才有效，否则它只返回None。

### 3.

`Thread.is_alive()`

`Thread.isAlive()`

判断线程是否是激活的`（alive）`。从调用`start()`方法启动线程，到`run()`方法执行完毕或遇到未处理异常而中断这段时间内，线程是激活的。

### 4. 

`Thread.activeCount()`

得到通过`threading`模块创建的所有正在运行的线程个数(包括主线程).

### 5.

`Thread.setDaemon(1)`

设置子线程随着主线程程的退出而结束.

在终端直接运行python脚本时, 如果不手动设置这个值, `Ctrl + C`无法返回终端, 因为子线程无法释放.

### 6. 

`Thread.join([timeout])`

调用`Thread.join`将会使主调线程堵塞，直到被调用线程运行结束或超时。参数`timeout`是一个数值类型，表示超时时间，如果未提供该参数，那么主调线程将一直堵塞到被调线程结束。

下面举个例子说明`join()`的使用: 

```py
#!encoding: utf-8
import threading, time
def doWaiting():
    print 'start waiting:', time.strftime('%H:%M:%S')
    time.sleep(3)
    print 'stop waiting', time.strftime('%H:%M:%S')
thread1 = threading.Thread(target = doWaiting)
thread1.start()
time.sleep(1)  #确保线程thread1已经启动
print 'start join'
thread1.join()	#将一直堵塞，直到thread1运行结束。
print 'end join'
```

写一下我自己的理解.

在多线程的应用场景下, 主线程与各个子线程各自独立运行(当然, 操作共享变量需要加锁), 在主线程中创建子线程并启动后, 主线程继续向下执行, 不会因为子线程在执行而阻塞.

以上面创建线程的两种方式的示例代码为例. 在print打印`count`变量之前, 主线程等待了2秒钟, 为什么?

可以实验一下, 把`time.sleep(2)`移除, 你会发现, `print count`会立即执行, 而`count`值输出为0.

因为子线程的独立运行, 并不会阻塞主线程, 主线程一路到底. 但是整个程序还不会退出, 因为整个进程的标准输入/输出/错误被所有子线程共享, 所以只有子线程完成, 命令行才会退出.

但你不知道子线程什么时候结束. `time.sleep(2)`只是说所有线程的计算会在2秒内完成, 所以等待2秒, 如果这个时间难以估算呢? 难道我要给它充足的时间去跑吗? 

这就是`join`的作用了.