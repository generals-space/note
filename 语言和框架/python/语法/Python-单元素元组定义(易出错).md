# Python-单元素元组定义(易出错)

参考文章

1. [程序员必知的Python陷阱与缺陷列表](http://www.cnblogs.com/xybaby/p/7183854.html)

```py
>>> a = (1, 2)
>>> type(a)
<class 'tuple'>
>>> type(())
<class 'tuple'>
```

如果只有一个元素

```py
>>> a = (1)
>>> type(a)
<class 'int'>
```

...怪不得在使用数据库驱动时传入变量总是出现问题.

表示单元素元组的正确方法是:

```py
>>> a = (1,)
>>> type(a)
<class 'tuple'>
>>>
```

逗号`,`必不可少.