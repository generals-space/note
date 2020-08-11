# Python执行shell命令

参考文章

1. [python中如何调用shell 中OS.SYSTEM等方法](http://blog.csdn.net/gray13/article/details/7044453)
    - `os.system()`, `os.popen()`, `subprocess.Popen()`基本介绍
2. [subprocess 模块](https://www.cnblogs.com/bigberg/p/7136952.html)
    - subprocess 详解, 用于替代 `os.system()`, `os.popen()` 等一系列方法.

## 1. os.system()

`os.system()`实际上是使用C标准库函数`system()`实现的, 

```py
import os
result = os.system('ping -c 30 www.baidu.com')
print(result)
```

参考文章1说, 无法保存命令的执行结果, 因为在实际执行时, `ping`的输出会直接打印在终端上, 无法通过 result 捕获, 也没有其他像是`stdout`的属性来指定输出目标, 看起来像是和当前进程的标准输入输出和错误绑定在了一起.

```console
$ python3 main.py
PING www.a.shifen.com (180.101.49.11) 56(84) bytes of data.
64 bytes from 180.101.49.11 (180.101.49.11): icmp_seq=1 ttl=128 time=11.9 ms
64 bytes from 180.101.49.11 (180.101.49.11): icmp_seq=2 ttl=128 time=22.1 ms
```

另外, 参考文章1说ta运行时会调用`bash`创建一个新的 shell 执行目标命令, 没看懂是为什么. 如下

```console
$ ps -ef | grep ping
root      88104  88093  0 16:57 pts/1    00:00:00 ping -c 30 www.baidu.com
root      88152 102735  0 16:57 pts/0    00:00:00 grep --color=auto ping
$ ps -ef | grep 88093
root      88093  77390  0 16:57 pts/1    00:00:00 python3 main.py
root      88104  88093  0 16:57 pts/1    00:00:00 ping -c 30 www.baidu.com
root      88227 102735  0 16:57 pts/0    00:00:00 grep --color=auto 88093
```

`python`进程明显是`ping`进程的父进程, 并不存在一个 shell.

`result`是类似返回码的`int`数值, 但是ta的值与真正的退出码又不相等. 如下

```py
import os
result = os.system('lsx')
print(result)
```

```console
$ python3 main.py
sh: lsx: 未找到命令
32512
```

但是如果直接在`bash`中执行`lsx`, 得到的值如下

```console
$ lsx
-bash: lsx: 未找到命令
$ echo $?
127
```

## 2. 

在一个子shell中运行command命令, 并返回command命令执行完毕后的退出状态. 这实际上是使用C标准库函数system()实现的. 这个函数在执行command命令时需要重新打开一个终端, 并且无法保存command命令的执行结果. 

os.popen()得到的是file read对象, 需要对其进行读取`read()`的操作才可以看到执行的输出, 适合`ps`, `ls`这种命令.
