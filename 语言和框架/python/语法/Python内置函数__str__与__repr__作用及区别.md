# Python内置函数__str__与__repr__作用及区别

参考文章

1. []()

`__str__`与`__repr__`都是类的成员方法, 网卡有许多人说前者是给用户看的, 后者是给程序员看的...白痴, 都是给程序员看的好吗?

以如下代码为例

```py
#!/usr/bin/env python
#!encoding: utf-8

class A(object):
    def __init__(self, name):
	self.name = name

    def __str__(self):
	return self.name + ' in __str__'

    def __repr__(self):
	return self.name + ' in __repr__'

a = A('general')
print(a)                            ## general in __str__
```

在python命令行中输入同样的代码, 通过`print`和直接打印两种方法查看到的是不一样的.

```py
>>> class A(object):
...     def __init__(self, name):
...         self.name = name
...     def __str__(self):
...         return self.name + ' in __str__'
...     def __repr__(self):
...         return self.name + ' in __repr__'
... 
>>> a = A('general')
>>> print(a)
general in __str__
>>> a
general in __repr__
```

直接打印`a`将得到`__repr__`方法中的结果, 但是这在代码里是完全不可能出现的. 所以`__repr__`的结果只会在命令行里出现, 而这个函数的目的, 就是提供一个类实例对象的简单展示, 而不是直接把类实例转换成一个字符串.

如果想看`__repr__`的结果, 可以通过python内置函数`repr()`完成, 与之相对的是`str()`函数, 它们分别直接输出`__repr__`和`__str__`函数的返回值.

似乎还有一个`__unicode__`方法, 作用和`__str__`一样, python2中推荐使用`__unicode__`, python3中推荐`__str__`, 貌似也有一个与之对应的`unicode()`内置函数, 以后再说.
