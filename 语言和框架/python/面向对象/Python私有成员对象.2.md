# Python私有成员对象(二)

参考文章

1. [Python 中特殊变量/方法命名规则说明(特别是私有变量)及使用实例](http://www.cnblogs.com/king-sun/p/4361998.html)

2. [关于Python的class的私有变量扎压 Python的命名机制](http://blog.csdn.net/zsuguangh/article/details/6207469)

这种下划线的命名机制一般出现在类方法中, 但在Python中类是对象, 函数也是对象, 函数嵌套时此类下划线的命名约定是否生效? 模块内的方法定义呢?

函数嵌套时, 内部函数的角色等同于局部变量, 拥有同样的生命周期和作用域, 这个其实没必要考虑.

```py
def outer():
    def inner():
        pass
```

在模块定义中, 无非会有这样的假设: 单下划线`_`开始的函数变量只有同文件模块内的方法和相同package的其他模块函数使用(...呃, 说使用不如说导入合适, 毕竟只要能导入就应该能调用了). 双下划线`__`开始的函数/变量只有本文件内的函数能使用.

实验一下.

`a.py`

```py
#!/usr/bin/env python
#!encoding: utf-8

_a = 123

def _sayA():
    print('i am a.py\'s _sayA')
```

同目录下的`b.py`

```py
#!/usr/bin/env python
#!encoding: utf-8

from a import _a, _sayA

print(_a)       ## 123
_sayA()         ## i am a.py's _sayA
```

...完全没影响.

OK, 其实换成双下划线也是如此, 就不再演示了.

看来这种下划线命名法只适用于类方法/变量, 只有这种场景才有私有变量轧压机制...