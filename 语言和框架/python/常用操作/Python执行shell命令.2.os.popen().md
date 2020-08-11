# Python执行shell命令

参考文章

1. [python中如何调用shell 中OS.SYSTEM等方法](http://blog.csdn.net/gray13/article/details/7044453)
    - `os.system()`, `os.popen()`, `subprocess.Popen()`基本介绍
2. [subprocess 模块](https://www.cnblogs.com/bigberg/p/7136952.html)
    - subprocess 详解, 用于替代 `os.system()`, `os.popen()` 等一系列方法.

## 基本用法

`os.popen()`的用法基本如下

```py
import os
result = os.popen('ls /tmp')
print(result.read())
```

`popen`会返回一个管道文件, 可以对这个文件进入读取操作. 

不只`read()`, 还有`readline()`, `readlines()`, 命令结果中所有以空格分隔的结果, 在`result`这个管道文件中, 都会变成按行存储下来.

## 标准错误

但是`result`只包括标准输入的信息, 如果命令执行出错, 标准错误中的信息将会直接打印出来.

```py
import os
result = os.popen('lsx')
print('===============')
```

```console
$ python3 main.py
===============
/bin/sh: lsx: 未找到命令
```

> 注意, 上面`print()`的横线先打印出来, 还是比较神奇的...

## `read()`的必要性

```py
import os
result = os.popen('ping -c 30 www.baidu.com')
## print(result.read())
```

如果执行上面的代码, 程序将会立刻结束, 不会有`ping`进程出现, 只有解除`print()`行的注释后, 才能正常进行`ping`操作, 这么看来, `read()`方法还有类似多线程中的`join()`的功能呢.

## `mode`参数-打开模式

参考文章1和2都说过, 我们能对`result`执行`read()`操作是因为, `popen`默认使用`r`模式打开管道文件, 因此ta是可读的, 所以我就尝试着试验了下用`w`模式去执行一个命令.

用`r`打开`result`文件是因为多数命令是数据打印在标准输出的, 那么用`w`打开文件则要考虑一些需要从标准输入读取数据的命令, 于是我想到了`passwd`

`passwd`默认是交互式输入, 如果想要在脚本中, 进行非交互式的操作, 需要使用如下命令.

```bash
echo "new_password" | passwd --stdin general
```

换成`popen()`

```py
result = os.popen('passwd --stdin general', mode='w')
result.write('123456')
result.close() ## 注意, 一定要有close(), 否则操作可能失败.
```

执行时结果如下

```console
$ python3 main.py
更改用户 general 的密码 。
passwd：所有的身份验证令牌已经成功更新。
```
