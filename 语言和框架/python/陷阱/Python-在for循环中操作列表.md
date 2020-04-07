# Python-在for循环中操作列表

参考文章

1. [Python的list循环遍历中，删除数据的正确方法](http://www.cnblogs.com/bananaplan/p/remove-listitem-while-iterating.html)

2. [为什么python中不建议在for循环中修改列表？](https://www.zhihu.com/question/49098374)

## 1. 引言

这里讲到的操作是在循环过程中删除或增加元素这些, 而不是修改成员变量的值的操作, (交换列表元素也是可以的). 

举例如下

```py
a = [1, 2, 3, 4]
for i in range(len(a)):
    if a[i] == 3:
        a.pop(i)
print(a)
```

执行它会报错

```
Traceback (most recent call last):
  File "test.py", line 3, in <module>
    if a[i] == 3:
IndexError: list index out of range
```

但是新增成员是可以的, 如下

```py
a = [1, 2, 3, 4]
for i in range(len(a)):
    if a[i] == 3:
        a.append(4)
print(a)
```

执行它, 你将得到如下结果

```
[1, 2, 3, 4, 4]
```

看起来还蛮正常的, 但是, 其实它也是问题的.

下面我们分析一下.

## 2. 问题研究

### 1. pop操作

```py
a = [1, 2, 3, 4]
for i in range(len(a)):
    print(i)
    if a[i] == 3:
        a.pop(i)
    print(a)
```

执行它, 得到输出

```
0
[1, 2, 3, 4]
1
[1, 2, 3, 4]
2
[1, 2, 4]
3
Traceback (most recent call last):
  File "test.py", line 4, in <module>
    if a[i] == 3:
IndexError: list index out of range
```

可以看到, 将a的成员`3`移出后, 列表长度变成了`3`, 但是下一次循环的索引也是`3`, 实际上超出了列表的长度, 就报错了.

也就是说, 在for循环开始之前, 循环次数就确定下来, 然后按照这个索引去遍历.

### 2. append操作

别以为上面在tor循环内append没报错就觉得的没问题, 我们把代码改成如下看看

```py
a = [1, 2, 3, 4]
for i in range(len(a)):
    print(i)
    print(a[i])
    if a[i] == 3:
        a.append(5)
    print(a)
```

输出为

```
0
1
[1, 2, 3, 4]
1
2
[1, 2, 3, 4]
2
3
[1, 2, 3, 4, 5]
3
4
[1, 2, 3, 4, 5]
```

只循环了4次, 新增元素导致列表长度变长, 最后一个成员没有被遍历到.

------

在增删操作中, 列表本身元素会发生移动, 但索引仍然是按循环开始时确定好的来执行, 这样的执行结果肯定不会是我们想要的.

## 3. 解决办法

参考文章1中提到了两种解决方案

1. 改用`while`循环, 增删的同时手动维护一个索引变量

2. 倒序循环遍历

还有一种是遍历目标列表的一个副本, 但这种情况只适合增加, 不适合删除, 并不想采用它. 只有一点, 它的循环写成如下形式

```
a = [1, 2, 3, 4]
for i in a[:]
    ...
```

其中`a[:]`是原列表`a`的一份拷贝, 与`a`本身并无关系(有点像匿名函数呃), 它只能保证不会产生删除列表成员后的越界问题, 而且在列表过大时...呵呵.

下面我们讨论一下另外两种方案.

### 3.1 while循环

```py
a = [1, 2, 3, 4]

i = 0
while i < len(a):
    print(i)
    if a[i] == 3:
        a.pop(i)
        i -= 1
    i += 1
    print(a)

```

输出如下

```
0
[1, 2, 3, 4]
1
[1, 2, 3, 4]
2
[1, 2, 4]
2
[1, 2, 4]
```

列表是蛮正常的, 唯一的缺点就是索引变量得控制好, 我觉得这个地方会很容易出错的.

同样, append操作也很正常, 不会发生遗漏, 不过代码得这么写.

```
a = [1, 2, 3, 4]

i = 0
while i < len(a):
    print(i)
    print(a[i])
    if a[i] == 3:
        a.append(5)
    i += 1
    print(a)
```

这种方式有点反人类呃.

### 3.2 倒序循环

我觉得这种方式更讨打...我不想深入研究.

-----

倒是参数文章1的第2条评论提到的方法, 我很喜欢

> 也可以赋值为None, 然后删除完过滤一遍, 这种方法就是不用脑子, 不用反着计算索引.

妙啊...

删除的时候不要真的删除, 把值赋为`None`. 循环完毕后再一次把成员为`None`的删除了就行, 当然, 这次删除依然不能用for, 用while就简单了.

```py
while None in a:
    a.remove(None)
```

删完了就好了.