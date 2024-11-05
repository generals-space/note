# Python迭代器和生成器(一) - 迭代器

参考文章

1. [迭代器与生成器](http://python.jobbole.com/84527/)

2. [完全理解 Python 迭代对象、迭代器、生成器](http://python.jobbole.com/87805/)

迭代器（iterator）与生成器（generator）是 Python 中比较常用又很容易混淆的两个概念. 关于它们的关系, 我的初步认知是, 生成器可以更方便的将一个普通函数生成迭代器. 也可以说生成器是一种特殊的迭代器, 不过实现方式更加优雅.


## 1. for语句与可迭代对象（iterable object）

```py
for i in [1, 2, 3]:
    print(i)
```

```
1
2
3
```

```py
obj = {"a": 123, "b": 456}
for k in obj:
    print(k)
```

```
a
b
```

这些可以用在for语句进行循环的对象就是`可迭代对象`, 但它们可不是迭代器。

除了内置的数据类型（列表、元组、字符串、字典等）可以通过 for 语句进行迭代，我们也可以自己创建一个容器，包含一系列元素，可以通过 for 语句依次循环取出每一个元素，这种容器就是迭代器。(容器这个术语比较难以理解, 其实上面提到的诸如列表, 元组, 字符串, 字典, 都可以看作是容器, 容器就是一个数据结构.)

但是并不是我们随便创建一个容器, 就能当作迭代器的, 需要满足一些条件才行.

------

现在我们来理一理可迭代对象, 容器, 和迭代器的关系.

实现迭代器有3种方法, 第3种是生成器, 这个在下一节讲述. 前两种就是

1. 可迭代对象通过`iter()`内置函数得到迭代器.

2. 我们创建的容器实现迭代器协议后也可以得到迭代器对象.

## 2. 可迭代对象通过`iter()`内置函数得到迭代器

```py
ita = iter([1, 2, 3])
print(type(ita))                    ## <type 'listiterator'>
 
print(next(ita))                    ## 1
print(next(ita))                    ## 2
print(next(ita))                    ## 3
```

### 2.1 迭代器对象与for循环兼容

```py
ita = iter([1, 2, 3])
for i in ita:
    print(i)
```

注意: 一次性的

实际上, for循环的实现貌似都是基于迭代器的, 即最终都会转换成`for i in iter(列表|字典)`这种形式.

### 2.2 in和not in

迭代器和可迭代对象一样, 都可以通过`in`, `not in`判断元素是否存在

```py
>>> ita = iter([1, 2, 3])
>>> 2 in ita
True
>>> 2 in ita
False
```

注意: 也是一次性的(只能判断一次, 就会将`ita`全部遍历一遍, 之后就没用了)

### 2.3 常规索引无效

```py
>>> ita = iter([1, 2, 3])
>>> ita[0]
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: 'listiterator' object has no attribute '__getitem__'
```

## 3. 自定义容器

为了使我们创建的容器成为迭代器, 我们需要遵守如下协议

为容器对象添加 `__iter__()` 和 `next()` 方法. 

`__iter__()` 返回迭代器对象本身`self`，`next()`则返回每次调用 `next()` 或迭代时的元素；

```py
class Container:
    def __init__(self, start = 0, end = 0):
        self.start = start
        self.end = end
    def __iter__(self):
        print("[LOG] I made this iterator!")
        return self
    def next(self):
        print("[LOG] Calling __next__ method!")
        if self.start < self.end:
            i = self.start
            self.start += 1
            return i
        else:
            raise StopIteration()
c = Container(0, 5)
print(type(c))                      ## <type 'instance'>...竟然不是iterator
print(c.next())                     ## 0
print(c.next())                     ## 1
## for i in c:
##     print(i)
```

创建迭代器对象的好处是当序列长度很大时(比如一个几百万条数据的列表, for循环时...呃, 不敢想象)，可以减少内存消耗，因为每次只需要记录一个值即可.

这要求列表数据不是随机的, 而是可以计算的, 否则也没什么意义...突然感觉作用也不是很大啊...


目前没发现什么实际的应用场景, 可能在数学计算领域会常见一些. 因为迭代器和生成器的核心在于推导, 通过一个值可以得到下一个值, 而不是通过一个单纯的列表得到所有的成员.

在这种计算领域...斐波那契数列算是标准教程, 应该会用于矩阵运算, 百万级列表的生成, 存在列表里是不可想像的.