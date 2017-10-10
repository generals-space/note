title: Python标准库-os模块exec系函数
---

tags: 语言 python os模块 exec函数族

`os.exec*()`都只是posix系统调用的直接映射

## 1. 

```python
os.execv(program, cmdargs)
```

基本的`v`执行形式, 需要传入可执行文件路径, 以及用来运行程序的命令行参数字符的**列表**或**元组**.

示例:

```python
os.execv('/bin/ping', ['/bin/ping', '-c', '10', 'www.baidu.com'])
os.execv('/bin/ping', ['ping', '-c', '10', 'www.baidu.com'])
os.execv('/bin/ping', ('/bin/ping', '-c', '10', 'www.baidu.com'))
```

注意, `execv`函数的第一个参数为目标程序路径, 第二个参数(即列表参数)的第一个参数也是可执行文件的路径, 它将出现在目标程序的`argv[0]`的位置. 如果目标程序在环境变量中, 则第二个参数的第一个成员可以不写完整路径, 直接写程序名即可. 但是第一个参数必须是完整路径, 且是绝对路径.

如果目标程序不需要任何参数, 我们可以传入空列表, 此时目标程序的`argv[0]`等同于参数1. 当然也可以传入与参数1相同的列表成员

```python
os.execv('/bin/pwd', [''])
os.execv('/bin/pwd', ['/bin/pwd'])
```

## 2. 

```python
os.execl(program, cmdarg1, cmdarg2, ..., cmdargN)
```

基本的`l`执行形式, 需要传入可执行文件路径, 以及用来运行程序的命令行多个字符参数.

示例

```python
os.execl('/bin/ping', '/bin/ping', '-c', '10', 'www.baidu.com')
os.execl('/bin/ping', 'ping', '-c', '10', 'www.baidu.com')
```

同理, 目标程序参数为空时, 需要使用如下方式

```python
os.execl('/bin/pwd', '')
os.execl('/bin/pwd', '/bin/pwd')
```

## 3. 

```python
os.execvp(program, args)
```

`p`模式下, 基本的`v`执行形式, 需要传入可执行文件路径, 以及用来运行程序的命令行参数字符的列表或元组. 运行新程序的搜索路径为当前文件的**相对路径**.

```python
os.execvp('/bin/pwd', ['pwd'])
os.execvp('./test.py', ['test.py', 'hello', 'world'])
```

其中, `test.py`是与当前程序同目录的程序, 它的内容可如下

```python
#!/usr/bin/env python
#!coding:utf-8

import sys

print(sys.argv)
```

试着执行一下.

## 4. 

```python
os.execlp(program, cmdarg1, cmdarg2, ..., cmdargN)
```

`p`模式下, 基本的`l`执行形式, 需要传入可执行文件路径, 以及用来运行程序的命令行多个字符参数. 运行新程序的搜索路径为当前文件的搜索路径.

```python
os.execvp('/bin/pwd', ['pwd'])
```

## 5. 

```python
os.execve(program, args, env)
```

`e`模式下, 基本的`v`执行形式, 需要传入可执行文件路径, 以及用来运行程序的命令行参数字符的列表或元组. 最后还要传入运行新程序的需要的环境变量env**字典参数**.

使用这个函数, 我们可以为将要执行的目标程序设置环境变量而不污染全局的环境变量.

```python
os.execve('/root/test.py', ['test.py', 'hello', 'world'], {'EXEC_TEST': '1234'})
```

对应的`test.py`文件内容可为如下

```python
#!/usr/bin/env python
#!coding:utf-8

import sys
import os

print(sys.argv)
print(os.getenv('EXEC_TEST'))
```

## 6. 

```python
os.execle(program, cmdarg1, cmdarg2, ..., cmdargN, env)
```

`e`模式下,基本的`l`执行形式,需要传入可执行文件路径,以及用来运行程序的命令行多个字符参数.最后还要传入运行新程序的需要的环境变量env字典参数.

## 7.

```python
os.execvpe(program, args, env)
```

在`p`和`e`的组合模式下,基本的`v`执行形式,需要传入可执行文件路径,以及用来运行程序的命令行参数字符的列表或元组.最后还要传入运行新程序的需要的环境变量env字典参数.运行新程序的搜索路径为当前文件的搜索路径.

## 8.

```python
os.execlpe(program, cmdarg1, cmdarg2, ..., cmdargN, env)
```

在`p`和`e`的组合模式下, 基本的`l`执行形式, 需要传入可执行文件路径, 以及用来运行程序的命令行多个字符参数. 最后还要传入运行新程序的需要的环境变量env字典参数. 运行新程序的搜索路径为当前文件的搜索路径.