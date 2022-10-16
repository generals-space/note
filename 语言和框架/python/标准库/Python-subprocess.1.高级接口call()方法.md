# Python-subprocess(一)高级接口call()方法

参考文章

1. [python子进程模块subprocess详解与应用实例 之一](http://blog.csdn.net/fireroll/article/details/39153831)
2. [python子进程模块subprocess详解与应用实例 之二](http://blog.csdn.net/fireroll/article/details/39153947)
3. [python子进程模块subprocess详解与应用实例 之三](http://blog.csdn.net/fireroll/article/details/39153991)
4. [subprocess官方文档](https://docs.python.org/2.7/library/subprocess.html)
5. [Python subprocess模块学习总结](http://www.jb51.net/article/48086.htm)
6. [python subprocess.Popen 监控控制台输出](http://blog.csdn.net/mldxs/article/details/8555819)

subprocess最早在2.4版本中引入. 用来生成子进程, 并可以通过管道连接它们的输入/输出/错误, 以及获得它们的返回值. 

简单来说, 它就是fork + exec创建子进程方式的封装. 传统方式中, 我们创建子进程任务, 首先要fork一个子进程, 放弃父进程, 然后在子进程中选择通过管道连接双方的标准IO, 设置执行用户及运行目录等信息. 这一切, 使用subprocess就可以完成. 如果熟悉这个流程, 就不必再手动写大片代码了.

参考文章1, 2, 3其实就是参考文章4也就是官方文档的中文翻译而已, 不过也是够良心了.

调用subprocess启动子进程的推荐方式是使用下面的便利功能. 当这些还不能满足需求时, 才需要使用底层的Popen接口. 

## 1. call()方法

```py
subprocess.call(args, *, stdin=None, stdout=None, stderr=None, shell=False)
```

运行由args指定的命令(`*`表示可以是目标命令的参数, 与命令在同一列表中), 直到命令结束后, 返回调用命令的退出码. 

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

~~第一种`call()`函数的调用方式指定了目标命令字符串与`shell = True`参数, 而第一种则只指定了目标命令的list形式变量(不能指定`shell = True`哦, 倒是不会出错, 只是可能不会得到你期望的结果...貌似只接受list中的第一个成员作为命令去执行)~~

~~猜测没有显式指定`shell = True`的执行方式, 都是由subprocess直接执行的, 不过也没有太确切的验证方式. 反正list形式无法执行`history`与`exit`这样的shell内置命令, 因为根本找不到这两个命令在哪.~~

shell默认为False, 在Linux下, shell = False时, Popen调用os.execvp()执行args指定的程序；shell = True时, 如果args是字符串, Popen直接调用系统的Shell来执行args指定的程序, 如果args是一个序列, 则args的第一项是定义程序命令字符串, 其它项是调用系统Shell时的附加参数. 

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

## 2. `check_call()`与`check_output()`方法

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
