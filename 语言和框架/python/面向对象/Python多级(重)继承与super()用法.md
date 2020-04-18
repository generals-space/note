# Python中多级(重)继承与super()用法

<!tags!>: <!多级继承!> <!多重继承!> <!隔代调用!>

参考文章

1. [Python中多继承与super()用法](http://www.jackyshen.com/2015/08/19/multi-inheritance-with-super-in-Python/)

名词解释

> 多重继承: 同时继承多个类

> 多级继承: B继承A, C又继承B...

## 1. 经典类与新式类继承

Python类分为两种, 一种叫`经典类`, 一种叫`新式类`, 两种都支持多继承, 只不写法不太一样.

如下是简单的经典类与新式类示例代码.

### 1.1 经典类

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

### 1.2 新式类

经典类的基类不需要继承任何类, 但是新式类要求基类**一定要继承`object`**, 这样就可以使用`super()`函数来调用父类中的函数, 

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

## 4. 多重继承中同名方法继承问题

假设有两个父类A与B, 两者都包含方法`display`, 如果子类C同时继承A与B, 在未覆写`display`方法前, C的实例使用的是哪个类的成员方法?

实验代码如下

```py
class A():
    def display(self):
        print('A...')

class B():
    def display(self):
        print('B...')

## class C(A, B):
class C(B, A):
    pass

c = C()
c.display()
```

答案是, 取决于继承顺序的先后, 在子类C没有覆写`display`方法的情况下, C将继承第一个拥有`display`方法的父类方法. 尝试一下将继承顺序调换一个你就会明白.

另外, 如果需要同时执行两个父类中的`display`方法, 可以在C类中覆写`display`, 并使用如下代码分别调用两者.

```py
class C(B, A):
    def display(self):
    	A.display(self)
        B.display(self)
```

> 注意`A.display()`中的`self`参数.

## FAQ

### 1.

```
    super(C, self).__init__()
TypeError: must be type, not classobj
```

解决方法: 使用`super()`方法时, 基类必须继承`object`.