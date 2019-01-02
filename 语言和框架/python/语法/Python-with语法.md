# Python-with语法

参考文章

1. [python with关键字学习](https://www.cnblogs.com/Xjng/p/3927794.html)

2. [浅谈 Python 的 with 语句](https://www.ibm.com/developerworks/cn/opensource/os-cn-pythonwith/)

## 1. 打开文件

with语句时用于对`try except finally`的优化, 让代码更加美观.

例如常用的开发文件的操作, 用`try except finally`实现: 

```py
f=open('file_name','r')
try:
    r=f.read()
except:
    pass
finally:
    f.close()
```

打开文件的时候, 为了能正常释放文件的句柄, 都要加个`try`, 然后再`finally`里把f `close`掉, 但是这样的代码不美观, `finally`就像个尾巴, 一直托在后面, 尤其是当`try`里面的语句时几十行

用`with`的实现: 

```py
with open('file_name','r') as f:
    r=f.read()
```

这条语句就好简洁很多, 当`with`里面的语句产生异常的话, 也会正常关闭文件.

## 2. 上下文管理器

### 2.1 

除了打开文件, with语句还可以用于哪些地方呢？

with只适用于上下文管理器的调用, 除了文件外, with还支持`threading`、`decimal`等模块, 当然我们也可以自己定义可以给with调用的上下文管理器.

```py
class A():
    def __enter__(self):
        self.a=1
        return self
    def f(self):
        print 'f'
    def __exit__(self,a,b,c):
        print 'exit'
def func():
    return A()

with A() as a:
    1/0
    a.f()
    print a.a
```

使用类定义上下文管理器需要在类上定义`__enter__`和`__exit__`方法, 执行`with A() as a`: 语句时会先执行`__enter__`方法, 这个方法的返回值会赋值给后面的`a`变量, 当`with`里面的语句产生异常或正常执行完时, 都会调用类中的`__exit__`方法。

### 2.2 使用生成器定义上下文管理器

```py
from contextlib import contextmanager

@contextmanager
def demo():
    print '这里的代码相当于__enter__里面的代码'
    yield 'i ma value'
    print '这里的代码相当于__exit__里面的代码'

with demo() as value:
    print  value
```

### 2.3 自定义支持 closing 的对象

```py
class closing(object):
    def __init__(self, thing):
        self.thing = thing
    def __enter__(self):
        return self.thing
    def __exit__(self, *exc_info):
        self.thing.close()

class A():
    def __init__(self):
        self.thing=open('file_name','w')
    def f(self):
        print '运行函数'
    def close(self):
        self.thing.close()

with closing(A()) as a:
    a.f()
```

在开发的过程中, 会有很多对象在使用之后, 是需要执行一条或多条语句来进行关闭, 释放等操作的, 例如上面说的的文件, 还有数据库连接, 锁的获取等, 这些收尾的操作会让代码显得累赘, 也会造成由于程序异常跳出后, 没有执行到这些收尾操作, 而导致一些系统的异常, 还有就是很多程序员会忘记写上这些操作-_-!-_-!, 为了避免这些错误的产生, `with`语句就被生产出来了。`with`语句的作用就是让程序员不用写这些收尾的代码, 并且即使程序异常也会执行到这些代码（`finally`的作用）