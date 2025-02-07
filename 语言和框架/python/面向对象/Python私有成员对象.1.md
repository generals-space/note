# Python私有成员对象(一)

参考文章

1. [Python 中特殊变量/方法命名规则说明(特别是私有变量)及使用实例](http://www.cnblogs.com/king-sun/p/4361998.html)

2. [关于Python的class的私有变量扎压 Python的命名机制](http://blog.csdn.net/zsuguangh/article/details/6207469)

## 1. 命名规范

特殊变量命名

1. `_xxx`: 以单下划线开头的表示的是`protected`类型的变量. 即保护类型只能允许其本身与子类进行访问. 若内部变量标示, 如： 当使用`from M import`时, 不会将以一个下划线开头的对象引入 . 

2. `__xxx`: 双下划线的表示的是私有类型的变量. 只能允许这个类本身进行访问了, 连子类也不可以用于命名一个类属性(类变量), 调用时名字被改变(在类`FooBar`内部, `__boo`变成`_FooBar__boo`, 如`self._FooBar__boo`)

3. `__xxx__`定义的是特例方法. 用户控制的命名空间内的变量或是属性, 如`__init__`, `__import__`或是`__file__`. 只有当文档有说明时使用, 不要自己定义这类变量.  (就是说这些是python内部定义的变量名)

## 2. private类型

> 私有的成员函数可以被类内部的(公有或私有)成员函数调用, 间接的对外部提供服务接口, 但是私有成员函数不对外提供服务

```py
#!/usr/bin/env python
#!encoding: utf-8

class A(object):
    def __init__(self):
        self.__private()            ## 调用自身的私有对象

    def __private(self):
        print 'A的__private()'

class B(A):
    def __private(self):
        print 'B的__private()'

a = A()
b = B()

a.__private()                       ## 这里其实有点违规了, 私有方法不能在实例外调用的.
b.__private()
```

执行它, 得到如下输出

```py
A的__private()                                   ## A类的构造方法, 不解释
A的__private()                                   ## B类没有覆盖A类的构造方法, 所以也调用了`self.__private()`
Traceback (most recent call last):              ## 类外调用私有方法的错误姿势...
  File "pri.py", line 18, in <module>
    a.__private()
AttributeError: 'A' object has no attribute '__private'
```

## 2.1 第一个问题

首先, B类在实例化时, 默认继承A类的构造方法...但是构造方法调用的...还是A类的`__private()`???

...WTF!!!

把`__private()`换成普通的public成员方法结果就不一样, B类在实例化时会老老实实调用它自己的方法.

```py
#!/usr/bin/env python
#!encoding: utf-8

class A(object):
    def __init__(self):
        self.public()

    def public(self):
        print 'A的public()'

class B(A):
    def public(self):
        print 'B的public()'

a = A()                 ## 输出: A的public()
b = B()                 ## 输出: B的public()
```

一个我认为还可算合理的解释: 私有变量的情况, b在实例化时, 调用的是A类的构造方法(因为B类没有覆盖它), 构造方法中调用了`__private()`私有方法, 而私有方法是不可被继承, 也不可以在非类本身实例中访问的, 所以调用的还是A类的`__private()`

```py
#!/usr/bin/env python
#!encoding: utf-8

class A(object):
    def __init__(self):
        self.__private()            ## 调用自身的私有对象

    def __private(self):
        print 'A的__private()'
    def sayHi(self):
        self.__private()

class B(A):
    def __private(self):
        print 'B的__private()'

a = A()                             ## A的__private()
b = B()                             ## A的__private()

b.sayHi()                           ## A的__private()
```

这真tm是一个匪夷所思的问题<???>

### 2.2 第二个问题

为什么`a.__private()`会提示`'A' object has no attribute '__private'`??

这是python的私有变量轧压(这个翻译好拗口)机制, 英文是(`private name mangling`)

你可以理解为, python其实并没有实现继承时的`private`, `protected`这些机制, 私有变量没办法被非类本身实例调用是因为, python在运行时把私有变量/方法的名字改了, 所以子类实例找不到了.

...感觉好自欺欺人啊怎么破? (⊙﹏⊙)

不信? 试试下面的代码

```py
#!/usr/bin/env python
#!encoding: utf-8

class A(object):
    def __private(self):
        print 'A的__private()'
class B(A):
    def sayHi(self):
        print('B的sayHi()')
        print(A._A__private(self))
b = B()
b.sayHi()
```

```
B的sayHi()
A的__private()
None                                ## ..哦, 这个None不知道从哪冒出来的, 先忽略
```

看到了? A类的`__private()`被重命名为`_A__private()`...

当然, 在类方法内部还是可以通过`self.__private()`调用的, 但是在类外, 就只能通过`a._A__private()`这种方法调用了.

关于这一点, 你可以通过查看`dir(a)`查看实例a的可用属性有哪些, 你会发现`__private()`不存在.

这种**自欺欺人**的机制被称为`私有变量扎压`, 这翻译真是醉了, 它的英文原文是`private name mangling`, 但是直译起来...完全没毛病(⊙ˍ⊙)

## 3. 私有变量轧压机制详解

关于这个...命名转换机制, 默认原则为`__私有变量名(方法名)` -> `_类名__私有变量名(方法名)`

参考文章2中解释的十分详尽.

1. Python把以**两个或以上下划线字符开头**且**没有以两个或以上下划线结尾**的变量当作私有变量;

2. 私有变量会在代码生成之前被转换为长格式(变为**公有**);

注意: 因为轧压会使标识符(私有方法名)变长, 当超过255的时候, Python会切断, 要注意因此引起的命名冲突

有一种讨打的情况是, 如果类名全部由下划线组成时...没错, 像下面这种情况. 这种轧压机制就不再生效.

```
class _(object):
class __(object):
class ____(object):
```

具体的可以自己实现一下, 这种情况下在类方法外部也可以通过`类名.__私有方法名`调用了.

实际上第2节的第一个问题已经解决了, 在代码运行前, 变成了下面这个样子

```py
#!/usr/bin/env python
#!encoding: utf-8

class A(object):
    def __init__(self):
        self._A__private()            ## 调用自身的私有对象

    def _A__private(self):
        print 'A的__private()'

class B(A):
    def _B__private(self):
        print 'B的__private()'

a = A()
b = B()

a.__private()
b.__private()
```

A类的构造方法中调用的其实是`_A__private()`这个方法, B类中覆盖的`__private()`...没覆盖对地方. 所以B在实例时调用的还是`_A__private()`..呵呵