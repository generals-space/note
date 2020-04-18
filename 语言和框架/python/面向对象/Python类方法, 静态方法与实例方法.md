# Python类方法, 静态方法与实例方法

参考文章

1. [Python类方法、静态方法与实例方法](https://www.cnblogs.com/blackmatrix/p/5606364.html)

类中定义的方法有3种: 实例方法, 静态方法与类方法; 而类中定义的属性有两种: 类属性与实例属性.

## 实例方法与实例属性

```py
class A():
    ## 实例属性
    a = 100
    def __init__(self):
        pass

    def setProps(self, a, b):
        self.a = a
        self.b = b

a = A()

print('##############################################################')
## 在没有显式声明self.a时, 实例方法会把类属性当成实例属性来用的.
print(a.a) ## 100 
## 不存在实例属性self.b, 也不存在类属性A.b, 这里会报错
## print(a.b) ## AttributeError: 'A' object has no attribute 'b'

## 打印类属性
print(A.a) ## 100
## 不存在的类属性
## print(A.b) ## AttributeError: type object 'A' has no attribute 'b'

print('##############################################################')
a.setProps(1, 2)
## 声明了self.a, 就不再使用类属性A.a的值了
print(a.a) ## 1 
print(a.b) ## 2

## 打印类属性
print(A.a) ## 100
## 类属性A.b仍然不存在, 可见实例属性不能被当作类属性来用(反过来却可以)
## print(A.b) ## AttributeError: type object 'A' has no attribute 'b'

print('##############################################################')
## 可以在类外创建和修改类属性和实例属性
a.a = 10
print(a.a) ## 10
a.c = 20
print(a.c) ## 20

A.a = 90
print(A.a) ## 90
A.b = 80
print(A.b) ## 80
```

## 类方法与类属性

听说在python中比较少使用...类方法需要使用`@classmethod`装饰器声明, 第一个参数必须要为`cls`, 表示类对象, 与实例方法的`self`有异曲同工之妙, 咳.

```py
class A():
    def __init__(self):
        pass
    @classmethod
    def cls_func(cls):
        print(type(cls), cls)

## 通过类对象调用类方法, 不需要传入第一个参数
A.cls_func() ## <class 'type'> <class '__main__.A'>

## 也可以通过实例调用类方法
a = A()
a.cls_func() ## <class 'type'> <class '__main__.A'>
```

关于类方法, 如果存在(多级)继承, 通过子类调用父类的类方法时, 由于传入的`cls`对象一定是主调类对象, 所以类方法中的操作需要考虑`cls`的取值, 以免出错.

## 静态方法

使用`@staticmethod`装饰器声明, 无需传入固定参数(比如`self`, `cls`), 可以通过类和实例调用. 

但是由于没有`cls`和`self`参数, 所以静态方法无法对类或实例对象产生任何影响, 可以说 **静态方法就是定义在类中的普通方法**.

> 这一点与C++和Java这种强对象语言不同, 我记得在C++和Java中, 静态方法是与静态成员配合使用的, 同一个类的不同实例对象通过静态方法操作静态成员, 而静态成员的值在各实例对象中是共享的.

```py
class A():
    def __init__(self):
        pass
    @staticmethod
    def static_func():
        print('hello world')

A.static_func() ## hello world

a = A()
a.static_func() ## hello world
```

**注意**

Python 2中, 如果一个类的方法不需要`self`参数, 必须声明为静态方法, 即加上`@staticmethod`装饰器, 才能不带实例调用它. 

Python 3中, 如果一个类的方法不需要`self`参数, 不再需要显式声明为静态方法, 但是这样的话只能通过类去调用这个方法, 如果使用实例调用这个方法会引发异常. 

如下是在python 3下的测试代码.

```py
class A():
    def __init__(self):
        pass
    ## @staticmethod
    def static_func():
        print('hello world')

A.static_func() ## hello world

a = A()
## a.static_func() ## TypeError: static_func() takes 0 positional arguments but 1 was given

```

因为没有声明为静态方法, 所以在通过实例对象调用时, 会隐式地传入`self`参数, 但是`static_func`不接受任何参数, 于是出现了异常.
