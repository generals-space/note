# Python执行shell命令.3.subprocess

参考文章

1. [python中如何调用shell 中OS.SYSTEM等方法](http://blog.csdn.net/gray13/article/details/7044453)
    - `os.system()`, `os.popen()`, `subprocess.Popen()`基本介绍
2. [subprocess 模块](https://www.cnblogs.com/bigberg/p/7136952.html)
    - subprocess 详解, 用于替代 `os.system()`, `os.popen()` 等一系列方法.

可以说, `subprocess.Popen()`是最强大的 shell 命令执行库了, 以后就直接使用这个好了...不纠结

以下几个方法都是对`Popen()`的封装, 使用起来会更简单一些.

## run()

```py
subprocess.run(['ls', '/tmp'])
subprocess.run(args = ['ls', '/tmp'])
subprocess.run('ping -c 30 www.baidu.com', shell=True)
```

上面的几个很像`os.system()`, 直接与当前python进程的标准输入输出绑定, 无法获取.

`run()`返回结果为`CompletedProcess(args='ls /tmp', returncode=0)`, 有很多属性信息可以用.

我比较了下`shell`设置为`True`与`False`两种情况, 发现在进程关系上并没有什么不同, 都与下面的结构相似, 唯一的区别应该就是命令传入的方式有所不同了...

```console
$ ps -ef | grep ping
root      42921  42910  0 18:57 pts/1    00:00:00 ping -c 30 www.baidu.com
root      43002 102735  0 18:57 pts/0    00:00:00 grep --color=auto ping
$ ps -ef | grep 42910
root      42910  77390  0 18:57 pts/1    00:00:00 python3 main.py
root      42921  42910  0 18:57 pts/1    00:00:00 ping -c 30 www.baidu.com
root      43063 102735  0 18:57 pts/0    00:00:00 grep --color=auto 42910
```

------

还有一点区别.

```
subprocess.run(args = ['lsx', '/tmp'])
subprocess.run('ls /tmp', shell=True)
```

前者会抛出异常, `FileNotFoundError: [Errno 2] No such file or directory: 'lsx': 'lsx'`

而后者则是会把标准错误中的信息打印出来(同样没法捕获), `/bin/sh: lsx: 未找到命令`, 不影响整个程序.

## call()

```py
subprocess.call(['ls', '/tmp'])
subprocess.call(args = ['ls', '/tmp'])
subprocess.call('ping -c 30 www.baidu.com', shell=True)
```

与`run()`相似的是, 执行结果同样直接输出, 无法获取.

不同的是, `call()`的返回值为命令的退出码, 与直接执行这些命令的退出码是相同的.

可以看出, `call()`主要是用于判断目标命令执行成功与否的.

## 

还有几个, 不想写了, 感觉没什么用处, 详见参考文章2.

