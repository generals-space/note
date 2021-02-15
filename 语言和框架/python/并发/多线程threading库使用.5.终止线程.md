# 多线程threading库使用.5.终止线程

参考文章

1. [Python多线程之线程创建和终止](https://blog.csdn.net/suipingsp/article/details/40342939)
    - `threading.Event()`传入子线程对象, 作为控制变量
2. [Is there any way to kill a Thread?](https://stackoverflow.com/questions/323972/is-there-any-way-to-kill-a-thread)
    - 高票答案也是使用`threading.Event()`, 扩展`Thread`类
3. [How to stop a looping thread in Python?](https://stackoverflow.com/questions/18018033/how-to-stop-a-looping-thread-in-python)
    - 可以直接对`Thread`实例对象设置属性`flag`, 然后在线程函数内执行`getattr(t, "flag")`, 得到该属性的值.
    - `threading.Event`对象
    - 如果要控制多个线程的终止, 可能还是`threading.Evnet`比较方便.
    - 值得收藏
4. [python下使用ctypes获取threading线程id](http://xiaorui.cc/archives/3017)
    - `ctypes.cdll.LoadLibrary('libc.so.6')`
5. [How to obtain a Thread id in Python?](https://stackoverflow.com/questions/919897/how-to-obtain-a-thread-id-in-python)
    - `/usr/include/x86_64-linux-gnu/asm/unistd_64.h`系统调用映射表
6. [How can I kill a particular thread of a process?](https://unix.stackexchange.com/questions/1066/how-can-i-kill-a-particular-thread-of-a-process)
7. [How to kill a single thread from Terminal in Ubuntu?](https://askubuntu.com/questions/608343/how-to-kill-a-single-thread-from-terminal-in-ubuntu)
8. [linux进程与线程之间区别 --------------进程是系统资源分配的最小单位，线程是进程执行的最小单位](https://blog.csdn.net/alpha_love/article/details/62247914)
    - 如果说进程是一个资源管家, 负责从主人那里要资源的话, 那么线程就是干活的苦力. 
    - 一个管家必须完成一项工作, 就需要最少一个苦力, 也就是说, 一个进程最少包含一个线程, 也可以包含多个线程. 
9. [python原生结束线程的方法](https://www.cnblogs.com/jefferybest/archive/2011/10/09/2204050.html)
    - 主线程先创建一个`caller`线程, 然后用`caller`线程创建目标子线程(此时`setDaemon(True)`), 到时只用`event`通知`caller`线程结束就可以...
    - 思路清奇
10. [Python强制关闭线程的一种办法](https://www.oschina.net/question/172446_2159505)
    - `ctypes.pythonapi.PyThreadState_SetAsyncExc()`设置异常
    - 评论中提到在线程内部调用`sys.exit()`, 同样需要扩展一下`threading.Thread`类, 但我没试过, 貌似无效.
11. [Python | Different ways to kill a Thread](https://www.geeksforgeeks.org/python-different-ways-to-kill-a-thread/)
    - 算是把上面所有的情况都考虑到了吧...
12. [Python | Different ways to kill a Thread](https://gist.github.com/dogterbox/4bc54bf499c5a43b7fd6de4e4dde336e)
    - 参考文章11的备份文章
13. [不要粗暴的销毁python线程](http://xiaorui.cc/archives/4302)
    - 对 python 不提供终止线程的原因做了说明

> 进程是最小的系统资源分配单位, 线程是最小的 CPU 调度单位

## 场景描述

我有一个通过常规 web 框架创建的 http api 服务, 提供了一个接口, 用户每向这个接口提交一些数据, 就会在后端通过`threading.Thread`创建一个线程无限循环去处理ta.

由于需要在主线程控制这些子线程的增删, 所以我需要一个在不结束主线程的前提下结束指定子线程的方法(不要想`setDaemon()`了).

## thread.Event 终止标记

参考文章1, 2, 3都提到使用一个主线程子线程共享的"全局变量", 或是直接用`threading.Event`对象当作子线程关闭的开关. 如下

```py
def thread_worker(event:threading.Event, arg):
    while not event.wait(1):
        print ("working on %s" % arg)
    print("Stopping as you wish.")
```

但是在我的场景中, 线程中嵌套了很多个 while 循环, 且包含很多耗时操作, 类似如下代码.

```py
def thread_worker(event:threading.Event, arg):
    while True:
        ## do something
        for i in range(30):
            ## do something
            time.sleep(5)
        ## do something
        for j in range(30):
            ## do something
            time.sleep(5)
```

这就导致对 event 的判断可能要添加到很多地方, 而且在`time.sleep()`期间(真实代码中的耗时操作), 这个标记是无法立刻生效的, 目标线程可能要等待耗时操作完成后才有机会判断 event 对象, 这对于我来说是不可接受的.

## kill signal 单独杀死线程

由于`threading.Event`没办法直接终止目标线程, 所以我考虑使用类似于`kill`的方法, 对目标 thread 发送一个`ctrl-c`的信号, 或者`SIGTERM`啥的.

首先要得到目标线程的线程id...

但是`threading.Thread`实例对象的`name`和`ident`属性只是一个标识而已, 跟线程id根本没啥关系.

参考文章4, 5中提到使用`ctype`库从底层获取目标线程id, 实验可行.

```py
import ctypes
import threading

libc = ctypes.cdll.LoadLibrary('libc.so.6')

## 系统调用参考号 /usr/include/x86_64-linux-gnu/asm/unistd_64.h
SYS_gettid = 186

class ExtendThread(threading.Thread):
    '''
    主要是 tid 的获取, 因为 threading.Thread 的 name/ident 属性与线程 id 不是一回事.
    '''
    _tid:int = 0

    def __init__(self, *args, **kwargs):
        super(ExtendThread, self).__init__(*args, **kwargs)
        self._tid = libc.syscall(SYS_gettid)

    def get_tid(self):
        return self._tid

## thread = threading.Thread(target = task_module.run, args = (device_id), name = device_id)
thread = ExtendThread(target = task_module.run, args = (device_id), name = device_id)
thread.start()
```

...但是! linux 下根本没有在进程之外单独终止线程的方法 (ó﹏ò｡)

见参考文章6, 7.

另外看了下`_thread`标准库, 没有任何可取之处, 不愧是要被废弃的库...

