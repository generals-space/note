# Python字符串格式化

参考文章

1. [python 字符串格式化](http://www.cnblogs.com/xxby/p/5571620.html)
2. [Python format 格式化函数](http://www.runoob.com/python/att-string-format.html)

Python的字符串格式化有两种方式:

1. `%`占位符
2. `format()`函数配合`{}`

> 这两种方式在python2和python3中都可以使用.

## 1. `%`占位符

```
%[(name)][flags][width].[precision]typecode
```

各选项的涵义见参考文章1

常用的使用方法如下

```py
name = 'general'
age = 21
string = 'my name is %s, %d years old'
print(string % (name, age)) ## my name is general, 21 years old
```

打印字典对象

```py
user = {
    'name': 'general',
    'age': 21,
}
string = 'my name is %(name)s, %(age)d years old'
print(string % user) ## my name is general, 21 years old
```

直接输出`%`, 使用`%%`就可以了.

```py
string = '我是内容: %s, 我是百分号: %%'
print(string % 'general') ## 我是内容: general, 我是百分号: %
```

## 2. format()函数

Python2.6 开始，新增了一种格式化字符串的函数 `str.format()`，它增强了字符串格式化的功能。

基本语法是通过`{}`和`:`来代替以前的`%`.

```
[[fill]align][sign][#][0][width][,][.precision][type]
```

各选项的涵义见参考文章1

常用的使用方法如下

```py
name = 'general'
age = 21
string = 'my name is {:s}, {:d} years old'
print(string.format(name, age)) ## my name is general, 21 years old
```

打印字典对象

```py
user = {
    'name': 'general',
    'age': 21,
}
string = 'my name is {name:s}, {age:d} years old'
print(string.format(**user)) ## my name is general, 21 years old
```

> 注意: `format()`不支持直接把字典对象用`{:s}`打印出来, 这一点不如使用`'%s' % {'name': 'general'}`

直接打印`{}`而不转义, 可以使用双层的大括号`{{}}`.

```py
string = '我是内容: {}, 我是大括号: {{}}'
print(string .format('general')) ## 我是内容: general, 我是大括号: {}
```

------

经实验, `%`占位符方式貌似不支持将十进制整数转换成二进制, python2和3都不行.

```py
>>> print('%b' % 10)
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ValueError: unsupported format character 'b' (0x62) at index 1
```

只能使用`print('{:b}'.format(10))`这种形式.
