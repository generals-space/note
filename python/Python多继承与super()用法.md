# Python中多继承与super()用法

<!tags!>: <!多级继承!> <!多重继承!> <!隔代调用!>

参考文章

1. [Python中多继承与super()用法](http://www.jackyshen.com/2015/08/19/multi-inheritance-with-super-in-Python/)

> 多重继承: 同时继承多个类

> 多级继承: B继承A, C又继承B...


## 1. 经典类与新式类继承

Python类分为两种, 一种叫经典类, 一种叫新式类, 两种都支持多继承.

如下是简单的经典类与新式类示例代码.

```py
#!/usr/bin/env python
#!coding:utf-8
## classic_class.py
## 经典类

class A():
    def __init__(self):
        print('A')
class B(A):
    def __init__(self):
        ## 这里为A.__init__()传入的self应该是一个对象实例
        ## 在子类中调用时, 这个参数应该是必须传入的吧???
        A.__init__(self)
        print('B')
class C(B, A):
    def __init__(self):
        A.__init__(self)
        B.__init__(self)
        print('C')

c = C()
```

执行它得到如下输出 

```
$ python classic_class.py 
A
A
B
C
```

但是在子类中显式调用父类的方法, 还是会显得耦合有点严重, 不符合python编程中的DRY原则. 所以推荐使用新式类.

经典类的基类不需要继承任何类, 但是新式类要求基类一定要继承`object`, 这样就可以使用`super()`函数来调用父类中的函数, 

```py
#!/usr/bin/env python
#!coding:utf-8
## new_class.py
## 新式类

class A(object):
    def __init__(self):
        print('A')
class B(A):
    def __init__(self):
        super(B, self).__init__()
        print('B')
class C(B, A):
    def __init__(self):
        super(C, self).__init__()
        print('C')
c = C()
```

得到如下输出

```
$ python new_class.py 
A
B
C
```

使用`super()`方法, 只会执行找到的**第一个**父类的成员方法, 当然, 得是拥有此成员方法的父类才行. 

------

如果还想强制调用其他父类的方法, 就只能使用经典类中的方法了.

```py
#!/usr/bin/env python
#!coding:utf-8
## new_class.py
## 新式类
class A(object):
    def __init__(self):
        print('A')
class B(A):
    def __init__(self):
        super(B, self).__init__()
        print('B')
class C(B, A):
    def __init__(self):
        ## 由于继承顺序的问题, 这里只会调用父类B的init方法
        super(C, self).__init__()
        ## 如果要调用父类A的init方法, 只能这样做
        A.__init__(self)
        print('C')

c = C()
```

输出如下

```
$ python new_class.py 
A
B
A           ## 这是父类A的init方法输出
C
```

## 2. 多级继承与隔代调用

...好吧其实这两个名词都是我自己瞎编的 ╮(╯▽╰)╭

想像一个场景, B继承A, C继承B, C需要在什么情况下才调用父类A的成员方法?

经典类的做法如下.

```py
#!/usr/bin/env python
#!coding:utf-8
class A():
    def __init__(self):
        print('A')
class B(A):
    def __init__(self):
        print('B')
class C(B):
    def __init__(self):
        ## 注意这里传入的self
        A.__init__(self)
        print('C')
c = C()
```

这样竟然是可行的(⊙ˍ⊙)...而且我竟无言以对...

输出如下

```
$ python test.py 
A
C
```

> 注意: 隔代调用时`self`对象参数的传入是必须的

新式类的做法如下

```py
#!/usr/bin/env python
#!coding:utf-8
class A(object):
    def __init__(self):
        print('A')
    def greet(self):
        print('hello')
class B(A):
    def __init__(self):
        print('B')
class C(B):
    def __init__(self):
        ## 注意这里传入的self
        super(C, self).greet()
        print('C')
c = C()
```

其输出为

```
$ python new_class.py 
hello
C
```

## 3. 多重继承与继承顺序

当多重继承与多级继承同时使用时, 比如如下场景.

```py
class A()
class B(A)
class C(B, A)
```

注意C类的继承顺序, 不管在什么情况下, 都需要将B写在A之前, 不然会报如下错误

```
Traceback (most recent call last):
  File "文件名", line 13, in <module>
    class C(A, B):
TypeError: Error when calling the metaclass bases
    Cannot create a consistent method resolution
order (MRO) for bases B, A
```

## FAQ

### 1.

```
    super(C, self).__init__()
TypeError: must be type, not classobj
```

解决方法: 使用`super()`方法时, 基类必须继承`object`.