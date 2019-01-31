# Python高阶函数-map

参考文章

1. [python的map()函数](http://www.cnblogs.com/XXCXY/p/5180237.html)

## 1. 基本应用

`map()`是python内置的高阶函数, 它接收一个函数`f`和一个`list`, 并通过把函数`f`依次作用在`list`的每个元素上, 得到一个新的`list`并返回. 

例如, 对于list [1, 2, 3, 4, 5, 6, 7, 8, 9]

如果希望把list的每个元素都作平方, 就可以用map()函数：

![](https://gitee.com/generals-space/gitimg/raw/master/3273034c16996f965ba8ca5cd273fffb.jpg)

因此, 我们只需要传入函数f(x)=x*x, 就可以利用map()函数完成这个计算：

```py
def f(x):
    return x*x

print map(f, [1, 2, 3, 4, 5,  6,  7,  8,  9])
```

输出结果：

```
[1, 4, 9, 10, 25, 36, 49, 64, 81]
```

注意：`map()`函数不改变原有的`list`, 而是返回一个新的`list`. 

利用`map()`函数, 可以把一个`list`加工成另一个`list`, 只需要传入转换函数. 

由于list包含的元素可以是任何类型. 因此, `map()`不仅仅可以处理只包含数值的list, 事实上它可以处理包含任意类型的 list, 只要传入的函数f可以处理这种数据类型. 

## 2. 进阶应用

map()的第2+个函数其实不一定非得是list类型, 只要是可以迭代的都可以, 如list, tuple甚至是dict. 不过map的返回只能是list类型就对了.

示例如下


```py
>>> def f(x):
...     return x + 1
... 
>>> map(f, (1, 3, 6))
[2, 4, 7]
>>> 
```

至于字典类型调用map()函数的我还没找到好一点的示例. 看起来应该是map把字典变量的所有key当成list传进去了.

```py
>>> def g(x):
...     return x
... 
>>> map(g, {'a': 123, 'b': 'abc'})
['a', 'b']
>>> 
```

------

**如果map函数调用的处理函数接受不只一个参数呢???**

```py
def f(x, y):
    return x * y

map(f, [1, 3, 6], [2, 4, 9])
[2, 12, 54]
```

good. 可以得出, **处理函数f的参数个数必须是list列表的个数**.

那么, 如果list组数与f的参数个数不同怎么办? 多出来的list会不会有问题?? 各个list的成员个数不一样怎么破???

```py
>>> def f(x, y):
...     return x * y
... 
>>> map(f, [1, 3, 6], [2, 4, 9])
[2, 12, 54]
>>> map(f, [1, 3, 6])
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: f() takes exactly 2 arguments (1 given)
>>> map(f, [1, 3, 6], [2, 4, 9], [4, 5 , 8])
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: f() takes exactly 2 arguments (3 given)
>>> map(f, [1, 3, 6], [2, 4])
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
  File "<stdin>", line 2, in f
TypeError: unsupported operand type(s) for *: 'int' and 'NoneType'
```

我觉得, list个数一定要与f的参数个数匹配, 否则就会出错而且没得商量.

不过, 如果改动一下处理函数f的定义, 比如加上参数默认值的设置, 会不会可以让传入的list成员个数不同? 

```py
>>> def f(x, y = 2):
...     return x * y
... 
>>> map(f, [1, 3, 7], [6, 4])
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
  File "<stdin>", line 2, in f
TypeError: unsupported operand type(s) for *: 'int' and 'NoneType'
```

ok, 我错了...