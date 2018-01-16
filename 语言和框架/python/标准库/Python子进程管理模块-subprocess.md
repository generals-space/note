# Python子进程管理模块-subprocess

参考文章

1. [python子进程模块subprocess详解与应用实例 之一](http://blog.csdn.net/fireroll/article/details/39153831)

2. [python子进程模块subprocess详解与应用实例 之二](http://blog.csdn.net/fireroll/article/details/39153947)

3. [python子进程模块subprocess详解与应用实例 之三](http://blog.csdn.net/fireroll/article/details/39153991)

4. [subprocess官方文档](https://docs.python.org/2.7/library/subprocess.html)

5. [Python subprocess模块学习总结](http://www.jb51.net/article/48086.htm)

6. [python subprocess.Popen 监控控制台输出](http://blog.csdn.net/mldxs/article/details/8555819)

subprocess--子进程管理器​

## 1. subprocess 模块简介

subprocess最早在2.4版本中引入. 用来生成子进程，并可以通过管道连接它们的输入/输出/错误，以及获得它们的返回值。

简单来说, 它就是fork + exec创建子进程方式的封装. 传统方式中, 我们创建子进程任务, 首先要fork一个子进程, 放弃父进程, 然后在子进程中选择通过管道连接双方的标准IO, 设置执行用户及运行目录等信息. 这一切, 使用subprocess就可以完成. 如果熟悉这个流程, 就不必再手动写大片代码了.

参考文章1, 2, 3其实就是参考文章4也就是官方文档的中文翻译而已, 不过也是够良心了.

调用subprocess启动子进程的推荐方式是使用下面的便利功能。当这些还不能满足需求时，才需要使用底层的Popen接口。

## 1. 高级接口

### 1.1 call()方法

语法:

```py
subprocess.call(args, *, stdin=None, stdout=None, stderr=None, shell=False)
```

语义:
     运行由args指定的命令(`*`表示可以是目标命令的参数, 与命令在同一列表中)，直到命令结束后，返回调用命令的退出码。

示例:

```py
>>> import subprocess
>>> subprocess.call('ls -al /tmp', shell = True)
total 32
drwxrwxrwt.  7 root root 4096 Jun 29 03:06 .
dr-xr-xr-x. 19 root root 4096 May 23  2016 ..
-rw-rw-rw-   1 root root    6 Jun 27 17:03 sky.pid
0             ## 这个是命令执行的退出码
>>>
>>>
>>>
>>> subprocess.call(['ls', '-al', '/tmp'])
total 32
drwxrwxrwt.  7 root root 4096 Jun 29 03:06 .
dr-xr-xr-x. 19 root root 4096 May 23  2016 ..
-rw-rw-rw-   1 root root    6 Jun 27 17:03 sky.pid
0
```

~~第一种call()函数的调用方式指定了目标命令字符串与`shell = True`参数, 而第一种则只指定了目标命令的list形式变量(不能指定`shell = True`哦, 倒是不会出错, 只是可能不会得到你期望的结果...貌似只接受list中的第一个成员作为命令去执行)~~

~~猜测没有显式指定`shell = True`的执行方式, 都是由subprocess直接执行的, 不过也没有太确切的验证方式. 反正list形式无法执行`history`与`exit`这样的shell内置命令, 因为根本找不到这两个命令在哪.~~

shell默认为False，在Linux下，shell = False时, Popen调用os.execvp()执行args指定的程序；shell = True时，如果args是字符串，Popen直接调用系统的Shell来执行args指定的程序，如果args是一个序列，则args的第一项是定义程序命令字符串，其它项是调用系统Shell时的附加参数。

**简单来说, 就是指定`shell`为`False`时, 目标命令需要是列表形式, 如果`shell`为True, 目标命令可以直接是字符串, 就像在真正的shell环境中执行一样.**

官方文档中提到, 使用`shell = True`将会是一个安全隐患. 因为它可能引起shell注入攻击, 尤其是在直接运行读取到的用户输入时极其危险.

```py
>>> from subprocess import call
>>> filename = input("What file would you like to display?\n")
What file would you like to display?
non_existent; rm -rf / 
>>> call("cat " + filename, shell = True) ## 这里可是会跪的...
```

**call()函数的返回值应该就只有目标命令的退出码**, 上述示例中`ls`的输出是直接打印在屏幕上的, 没有办法取到.

官方文档中还说call()函数不要使用`stdout=PIPE`与`stderr=PIPE`, 因为可能由于子进程的输出行为导致死锁...我该骂一句sb吗? 如果有问题为什么还要提供这两个参数? 又不说明什么情况下才会死锁, 明明我自己试了两次都没什么问题, 未知情况什么的, 太闹心了.

### 1.2 `check_call()`与`check_output()`方法

这两者的作用基本上与`call()`一样, 没看出有什么大用...简单来说就是当目标命令的退出码不为0时这两个函数就会报异常...

```
>>> subprocess.call('ls /tmp; exit 1', shell = True)
sky.pid  sub.py
1
>>> subprocess.check_call('ls /tmp; exit 1', shell = True)
sky.pid  sub.py
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
  File "/usr/lib64/python2.7/subprocess.py", line 542, in check_call
    raise CalledProcessError(retcode, cmd)
subprocess.CalledProcessError: Command 'ls /tmp; exit 1' returned non-zero exit status 1
```

ok, 根据退出码决定报不报异常感觉还是太草率了, 太不负责任了, 所以我感觉这两个函数对我没什么用.

那`check_call()`与`check_output()`有什么区别呢? 

与`call()`一样, `check_call()`的返回值也是目标命令的退出码, 没法捕获输出. 但是`check_output()`可以, 不过没有退出码了, 二选一嘛.

```py
>>> a = subprocess.check_output('ls /tmp', shell = True)
>>> print(a)
sky.pid
sub.py
```

good...

## 2. 底层接口Popen

### 2.1 Popen()的使用方法

上面的几个函数都是基于Popen()的封装。这些封装的目的在于让我们容易使用子进程。当我们想要更个性化我们的需求的时候，就要转向Popen类，该类生成的对象用来代表子进程。

与上面的封装不同，Popen对象创建后，主程序不会自动等待子进程完成。我们必须调用对象的wait()方法，父进程才会等待 (也就是阻塞block)，举例：

```py
#!/usr/bin/env python
import subprocess

child = subprocess.call('ping -c 4 172.16.3.206', shell = True)
print('child complete')
```

你会发现, 在`call()`生成的ping子进程执行完成之前, print是无法打印信息的, 即发生了阻塞.

但是同样的调用方法, Popen()函数就不会阻塞, print几乎是立刻就打印出了信息.

```py
#!/usr/bin/env python
import subprocess

child = subprocess.Popen('ping -c 4 172.16.3.206', shell = True)
print('child complete')
```

为了让Popen()生成的子进程能够阻塞, 以便取到目标程序的输出, 我们需要使用`wait()`函数.

```py
#!/usr/bin/env python

import subprocess

child = subprocess.Popen('ping -c 4 172.16.3.206', shell = True)
child.wait()
print('child complete')
```

这下, 程序的运行结果就和call()函数相同了. 我们也发现了, Popen()的返回值是一个对象, 它拥有成员方法`wait()`. 实际上它正是suprocess生成的子进程对象.

我们可以直接操作这个子进程对象.

```
child.poll() # 检查子进程状态, 未结束返回None
child.kill() # 终止子进程
child.send_signal(signal) # 向子进程发送信号signal
child.terminate() # 终止子进程
```

这些操作我想应该是把子进程当成守护进程启动的了吧?

好吧我们来试试

```py
#!/usr/bin/env python

import subprocess
import time
child = subprocess.Popen('ping 172.16.3.206', shell = True, stdout = subprocess.PIPE, stderr = subprocess.PIPE)
time.sleep(10)
child.terminate()
print('child complete')
```

就这样, 执行的时候ping子进程会在后台执行, 并且未绑定标准IO, 不过我们可以使用child子进程对象对它进行操作.

再说一句, ping子进程貌似还脱离了当前session group, 因为当我关闭执行终端时, ping进程依然在运行. 竟然真的是把它当成守护进程运行的.

### 2.2 Popen获取目标命令输出

上面我们对Popen()的执行方式, 使得目标命令的输出都是打印在标准输出中, 却没有办法赋值到某个变量中. 

其实可以使用`subprocess.PIPE`选项将子进程与父进程的标准输入输出连接起来. 然后通过子进程对象的stdout读取. 示例如下

```py
#!/usr/bin/env python

import subprocess
import time
child = subprocess.Popen('ping -c 4 172.16.3.206', shell = True, stdout = subprocess.PIPE, stderr = subprocess.STDOUT)
child.wait()
print('%s' % child.stdout.read())
print('child complete')
```

由于要读取ping子进程的全部输出, 所以只能使用wait()函数等待它执行完成后再读取. `stderr = subprocess.STDOUT`这个参数非常重要, 它表示标准错误使用与标准输出相同的方法处理, 不然当目标子进程输出到标准错误时就无法捕获其信息了.

但是, 发现没有, ping子进程实际有4条输出, 至少每秒一条, 但是由于`wait()`存在, 我们必须得等到ping子进程完全结束才能读取它的输出. 但是如果目标命令的输出太多, 超过了缓冲区的大小的话, 就不知道会发生什么事情了.

正好我们可以使用poll方法检测子进程的运行状态, 然后实时读取它的输出, 直到结束.

```py
#!/usr/bin/env python

import subprocess
import time
child = subprocess.Popen('ping -c 4 172.16.3.206', shell = True, stdout = subprocess.PIPE, stderr = subprocess.PIPE)
while True:
    line = child.stdout.readline()
    print(line)
    if child.poll() != None:
        break
print('child complete')
```

不过好像好了点东西, 下面的那些没了

```
--- 172.16.3.206 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3000ms
rtt min/avg/max/mdev = 0.318/0.338/0.372/0.028 ms
```

不明觉厉...