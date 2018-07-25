# python模块导入

参考文章

[python语法31[module/package+import]](http://www.cnblogs.com/itech/archive/2010/06/20/1760345.html)

## 1. 引言

首先介绍两个概念

### 1.1 module(模块)

通常module是一个单独的python **文件**, 直接使用`import`来导入就好了. 可以作为module的文件类型有`.py`、`.pyo`、`.pyc`、`.pyd`、`.so`、`.dll`等. 

### 1.2 package(包)

通常package总是一个目录, 可以使用`import`导入包, 或者`from...import...`来导入包中的部分模块. 包目录下为首的一个文件便是 `__init__.py`. 然后是一些模块文件和子目录, 假如子目录中也有 `__init__.py` 那么它就是这个包的子包. 

python中的package必须包含一个`__init__.py`的文件, 如果 `__init__.py` 不存在, 这个目录就仅仅是一个目录, 而不是一个包, 它就不能被导入或者包含其它的模块和嵌套包. 

## 2. module(模块)

一个python源文件作为模块, 在其他主调程序中通过`import`语句就可以将这个文件作为模块导入. 系统在导入模块时, 会做以下三件事: 

1. 为模块源文件中定义的对象创建一个`命名空间`, 通过这个命名空间可以访问到模块中定义的函数及变量. 

2. 在新创建的命名空间里执行源代码文件.

3. 创建一个以模块源文件为名称的`对象`, 该对象引用模块的命名空间, 这样就可以通过这个对象访问模块中的函数及变量.

以如下代码为例, `person.py`文件作为调用模块, `caller.py`作为主调程序, 两者放在同一目录下.

```python
#!/usr/bin/python
#!coding:utf-8
## person.py模块
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

## caller.py文件, 导入person模块
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

#### 2.1.3 使用from语句可以将模块中的对象直接导入到当前的命名空间.

`from`语句不创建一个到模块命名空间的引用对象, 而是把被导入模块的一个或多个对象直接放入当前的命名空间, 可以直接使用. 

```python
## /usr/bin/python
from person import a, greet, Person
print(a)
greet()
Person().greet()
```

`from`语句也支持逗号分割的对象, 也可以使用星号(*)代表模块中除下划线开头的所有对象. 

```py
from person import a, greet
```

或

```py
from person import *
```

`from`语句可以和`as`也可以结合使用

```py
from person import Person as Leader

c = Leader()
c.greet()
```

------

使用`from`语句可以将模块中的对象直接导入到当前的命名空间, 使用`import`就不行了.

```py
import person.a
print(a)
```

将会得到如下报错.

```
$ python caller.py 
Traceback (most recent call last):
  File "caller.py", line 1, in <module>
    import person.a
ImportError: No module named a
```

也就是说, `import`单独使用时, 目标只能是模块, 或者包中的模块, 无法直接导入模块中的变量或对象.

#### 2.1.4 限制外部文件访问

如果一个模块如果定义有列表变量`__all__`, 则`from module import *` 语句只能导入`__all__`列表中存在的对象. 

```python
#!/usr/bin/python
#!coding:utf-8
## person.py模块

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

## 3. package(包)

package目录下的`__init__.py`, 在`import`操作时会自动运行, 并且如果packege目录下存在子目录, 那子目录的`__init__.py`也会执行.

### 3.1 简单的包导入方法

如下, `Test`是我们自定义的包, `testfile.py`文件是主调程序.

```
.
├── Test
│   ├── __init__.py
│   ├── test.py
└── testfile.py
```

```py
#!/usr/bin/python
#!coding:utf-8
## Test/__init__.py

print('__init__ in Test...')
```

```py
#!/usr/bin/python
#!coding:utf-8
## Test/test.py

test1 = 'testing...'
```

```py
#!/usr/bin/python
#!coding:utf-8
## testfile.py

## from Test import test
## print(test.test1)

## from Test.test import test1
## print(test1)

## import Test.test
## print(Test.test.test1)
```

上面三种import方式都是正确的, 执行`testfile.py`时都可以得到如下的输出.

```
$ python testfile.py 
__init__ in Test...
testing...
```

### 3.2 import package

上面示例中, `import`的直接宾语都是模块名. 要留心下面这种导入方式, 直接宾语为包名.

```py
import Test
print(Test.test.test1)
```

它会导致如下错误

```
$ python testfile.py 
__init__ in Test...
Traceback (most recent call last):
  File "testfile.py", line 12, in <module>
    print(Test.test.test1)
AttributeError: 'module' object has no attribute 'test'
```

就算在`Test/__init__.py`文件中明确写了`__all__ = ['test']`, 也无法解决.

这个示例说明, `import`可以直接导入一个package(简单点说就是目录), 但实际调用时不能通过`包名.模块名.对象`这种方式取得目标对象, 因为`包名.模块名`本身就是非法操作.

有效的做法是, 在`Test/__init__.py`文件中写入子模块的导入过程.

```py
#!/usr/bin/python
#!coding:utf-8
## Test/__init__.py

import test
print('__init__ in Test...')
```

这样, 再次执行`testfile.py`, 就可以得到正确的输出了.

同理, 当自定义的包中还包括子包和其他子模块时, 这样的操作就十分有效了.

> `from`的直接宾语可以是包名, 也可以是模块名, 而`import`的直接宾语可以是模块名, 也可以是模块内的变量名.