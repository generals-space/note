# Python高阶函数-filter

参考文章

1. [python的filter()函数](http://www.cnblogs.com/XXCXY/p/5180256.html)

`filter()`是`Python`内置的另一个有用的高阶函数. 

`filter()`接受一个函数`f`和一个list(也可以是元组), 这个函数`f`的作用是对list中每个元素进行判断, **返回`True`或`False`**, `filter()`根据判断结果**自动过滤**掉不符合条件的元素, 返回由符合条件元素组成的新list(两个list的成员数量应该就不一样了哦). 

例如, 要从一个list `[1, 4, 6, 7, 9, 12, 17]`中删除偶数, 保留奇数, 首先, 要编写一个判断奇数的函数: 

```py
def is_odd(x):
    return x % 2 == 1
## 然后, 利用filter()过滤掉偶数: 
filter(is_odd, [1, 4, 6, 7, 9, 12, 17])
```

结果: [1, 7, 9, 17]

利用filter(), 可以完成很多有用的功能, 例如, 删除 None 或者空字符串: 

```py
def is_not_empty(s):
    return s and len(s.strip()) > 0
filter(is_not_empty, ['test', None, '', 'str', '  ', 'END'])
```

结果: ['test', 'str', 'END']

注意: `s.strip(char)`可以删除`s`字符串中开头、结尾处的`char`序列的字符. 

当char为空时, 默认删除空白符(包括'\n', '\r', '\t', ' '), 如下: 

```py
a = '     123'
a.strip()
```

结果:  '123'

```py
a='\t\t123\r\n'
a.strip()
```

结果: '123'