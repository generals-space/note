# python模块导入

参考文章

[python语法31[module/package+import]](http://www.cnblogs.com/itech/archive/2010/06/20/1760345.html)

## 1. 引言

首先介绍两个概念

### 1.1 module(模块)

通常module为一个文件，直接使用`import`来导入就好了。可以作为module的文件类型有`.py`、`.pyo`、`.pyc`、`.pyd`、`.so`、`.dll`。

### 1.2 package(包)

通常package总是一个目录，可以使用`import`导入包，或者`from...import...`来导入包中的部分模块。包目录下为首的一个文件便是 `__init__.py`。然后是一些模块文件和子目录，假如子目录中也有 `__init__.py` 那么它就是这个包的子包。

## 2. module(模块)

一个源代码文件作为模块, 在其他主调程序的源码文件中通过`import spam`语句就可以将这个文件作为模块导入。系统在导入模块时，会做以下三件事：

1. 为源代码文件中定义的对象创建一个`名字空间`，通过这个名字空间可以访问到模块中定义的函数及变量。

2. 在新创建的名字空间里执行源代码文件.

3. 创建一个以源代码文件为名称的`对象`，该对象引用模块的名字空间，这样就可以通过这个对象访问模块中的函数及变量

以如下代码为例, person.py文件作为调用模块, caller.py作为主调程序, 两者放在同一目录下.

```python
## person.py模块
#!/usr/bin/python
a = 30
def greet():
    print('hello world')
class Person:
    def greet(self):
        print('hello Niko')
b = Person()
```

```python
## caller.py文件, 导入person模块
#!/usr/bin/python
#!coding:utf-8
import person

## person模块中的属性
print(person.a)
person.greet()
c = person.Person()
c.greet()
```

```
$ python caller.py
30
hello world
hello Niko
```

### 2.1 模块的其他调用方法

#### 2.1.1 用逗号分割模块名称可以同时导入多个模块

```python
## caller.py
## /usr/bin/python
import os, sys, person
```

#### 2.1.2 使用 as 关键字可以改变模块的引用对象的名称

```python
## /usr/bin/python
import person as leader
## 调用person模块中的greet函数(不是Person类中的greet方法)
leader.greet()
```

#### 2.1.3 使用from语句可以将模块中的对象直接导入到当前的名字空间.

from语句不创建一个到模块名字空间的引用对象，而是把被导入模块的一个或多个对象直接放入当前的名字空间, 可以直接使用。

```python
## /usr/bin/python
from person import a, greet, Person
print(a)
greet()
Person().greet()
```

from语句也支持逗号分割的对象，也可以使用星号(\*)代表模块中除下划线开头的所有对象。

```python
from person import a, greet
```

```python
from person import *
```

`from`语句可以和`as`也可以结合使用

```python
from person import Person as Leader

c = Leader()
c.greet()
```

#### 2.1.4 限制外部文件访问

如果一个模块如果定义有列表变量`__all__`，则`from module import *` 语句只能导入`__all__`列表中存在的对象。

```python
## person.py模块
#!/usr/bin/python

## 限制访问列表, 注意成员类型都为字符串
__all__ = ['a', 'Person']
a = 30
def greet():
    print('hello world')
class Person:
    def greet(self):
        print('hello Niko')
b = Person()
```

```python
#!/usr/bin/python
#!coding:utf-8
## from person import a, greet, Person
from person import *
print(a)
## 调用greet()函数会出错
## greet()
Person().greet()
```

但是这种形式只能限制以`from module import *`语句的导入方式, 如果`from person import a, greet, Person`直接指定模块需要导入的对象, 那么`__all__`是不起作用的. 另外, 在`__all__`存在的情况下, 主调程序中使用`import person`导入模块本身是不可行的, 无法使用任何模块中定义的对象.

```python
#!/usr/bin/python
#!coding:utf-8
## from person import a, greet, Person
from person import a, greet, Person
print(a)
## 依然可以调用greet()函数
greet()
Person().greet()
```

```python
#!/usr/bin/python
#!coding:utf-8
## from person import a, greet, Person
import person
## 注意, 以下任何一行都会出错!!!
print(a)
greet()
Person().greet()
```
