# uwsgi模块mules

参考文章

1. [uWSGI中文文档 - mules子系统](http://uwsgi-docs-zh.readthedocs.io/zh_CN/latest/Mules.html)

2. [uWSGI官方文档 - mules子系统](https://uwsgi-docs.readthedocs.io/en/latest/Mules.html)

## 1. 关于mules

uwsgi提供了多进程方式运行我们的python工程, 主进程自然起调度作用, 那实际上的处理进程(也就是worker进程)才量运行我们python代码逻辑地方, uwsgi把用户请求转发到worker进程上, 由我们编写的代码处理并生成响应.

这是一般的worker的作用.

而mule, 是和worker同等级的进程, 它也是worker, 但用户的请求不会被转发给它, 那么它存在的意义就是做一些非交互性的事情, 比如定时任务, 目录/文件变动的监控等.

## 2. 配置

Mule有两种模式.

1. 纯信号模式（默认模式）。在这种模式下，mule像正常的worker那样加载你的应用。它们只能响应 uWSGI signals。

2. 编程模式。在这种模式下，mule与你的应用分开加载一个程序.

### 2.1 纯信号模式

就像uwsgi的`--workers`(或`--processes`)选项一样, mule进程也是要多少有多少的, 随便配就行.

比如

```
$ uwsgi --socket :3031 --mule --mule
```

或


```ini
[uwsgi]
socket = :3001
mule = true
mule = true
```

### 2.2 编程模式

如果你想让mule进程运行独立于你python工程之外的脚本, 就用这种方式.

```
$ uwsgi --socket :3031 --mule=scripts.py --mule
```

```ini
[uwsgi]
socket = :3001
mule = scripts.py
mule = true
```

...有点眼熟啊.

## 3. 使用

### 3.1 纯信号模式

纯信号模式会为每一个`--mule`选项添加一个序号, `mule0`, `mule1`等(`mule`默认指向`mule0`).

你可以在自己的工程中指定某一个函数在`muleN`中运行, 纯信号模式下运行的函数一般有三种作用:

1. 定时器, 不过只能设置每隔n秒执行一次. cron形式的任务需要`uwsgi.add_cron()`函数.

2. 信号映射, 就是定义事件机制中的事件处理函数, 你可以在程序的其他地方触发这个信号来激活它的运行

2. 目录/文件变动监控

官方文档中有如下示例代码

```py
import uwsgi
from uwsgidecorators import timer, signal, filemon

## 在第一个mule进程中运行定时器
@timer(30, target='mule')
def hello(signum):
    print "Hi! I am responding to signal %d, running on mule %d" % (signum, uwsgi.mule_id())

## 映射信号17的处理函数到mule2
@signal(17, target='mule2')
def i_am_mule2(signum):
    print "Greetings! I am running in mule number two."

## 监控/tmp目录的变动然后推送给所有的mules进程
@filemon('/tmp', target='mules')
def tmp_modified(signum):
    print "/tmp has been modified. I am mule %d!" % uwsgi.mule_id()
```

> `uwsgidecorators`模块需要pip下载, 它只是给`uwsgi`的一些内置方法包装了一下.

### 3.2 编程模式

很多时候我想自己开线程控制后台任务, 写一个定时任务模式, uwsgi加载就好了, 毕竟用uwsgi的定时器强耦合感觉很不爽.

官方文档的示例代码

```py
import uwsgi
from threading import Thread
import time

def loop1():
    while True:
        print "loop1: Waiting for messages... yawn."
        message = uwsgi.mule_get_msg()
        print message

def loop2():
    print "Hi! I am loop2."
    while True:
        time.sleep(2)
        print "This is a thread!"

t = Thread(target=loop2)
t.daemon = True
t.start()

if __name__ == '__main__':
    loop1()
```

mmp, `loop1`貌似是必须存在的. mule进程貌似会循环加载我们的脚本, 而`uwsgi.mule_get_msg()`可以让这个循环阻塞(能实现这个效果的还有`uwsgi.signal_wait()`). 只有这样我们才能实现自己的逻辑...

## 4. 总结

```py
#!encoding: utf-8

## 下面这句加在定时任务模块的末尾...判断是否运行在uwsgi模式下, 然后阻塞mule主进程.
try:
    import uwsgi
    while True:
        sig = uwsgi.signal_wait()
        print(sig)
except Exception as err:
    pass
```