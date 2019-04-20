# Python内置方法, 魔术方法

参考文章

1. [Python学习笔记二十---- 对象高级 - 163博客](http://blog.163.com/qimeizhen8808@126/blog/static/16511951820127220173667)

2. [python中的globals()、locals()、dir()、vars()、__dict__](https://blog.csdn.net/biorelive/article/details/47319923)

__init__(self,...)-----初始化对象，创建对象时调用（实例化对象时）

__del__(self,...)------释放对象，在对象被删除时调用

__new__(cls,*args.**kwd)----实例的生成

__str__(self)----------在使用print时调用，将内容转换为字符串

__getitem__(self,key)----获取序列的索引key所对应的值

__len__(self)------------调用内敛函数len时候调用的是它  ##什么意思?

__cmp__(src,dst)---------比较

__getattr__(s,name)------ 获取属性的值

__setattr__(s,name,val)---设置属性的值

__delattr__(s,name)-------删除name的属性

__getattribute__()--------与 __getattr__ 类似

__gt__(self,other)--------判断self对象是否大于other 对象

__lt__(self,other)--------判断self对象是否小于other 对象

__ge__(self,other)--------判断self对象是否大于或等于other 对象

__le__(self,other)--------判断self对象是否小于或等于other 对象

__eq__(self,other)--------判断 self 对象是否等于 other 对象

__call__(self,*args)------把实例对象作为函数调用