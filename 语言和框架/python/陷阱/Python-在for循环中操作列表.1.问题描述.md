# Python-在for循环中操作列表.1.问题描述

参考文章

1. [Python的list循环遍历中，删除数据的正确方法](http://www.cnblogs.com/bananaplan/p/remove-listitem-while-iterating.html)
2. [为什么python中不建议在for循环中修改列表？](https://www.zhihu.com/question/49098374)

python: 3.7.2

## 1. 引言

这里讲到的操作是在循环过程中删除或增加元素这些, 而不是修改成员变量的值的操作, (交换列表元素也没什么问题). 

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

倒是新增成员不会报错, 如下

```py
a = [1, 2, 3, 4]
for i in range(len(a)):
    if a[i] == 3:
        a.append(5)
print(a)
```

> 注意: `append()`是在列表末尾而不是元素3后面追加.

执行它, 你将得到如下结果

```
[1, 2, 3, 4, 5]
```

看起来还蛮正常的, 但是, 其实它也是问题的.

下面我们分析一下.

## 2. 问题研究

### 1. pop操作

我们尝试将元素3从列表`[1, 2, 3, 4]`中移除.

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

可以看到, 将a的成员`3`移出后, 列表长度变成了`3`, 但是下一次循环的索引也是`3`(索引是从0开始的), 实际上超出了列表的长度, 就报错了.

也就是说, 在for循环开始之前, 循环次数就确定下来, 然后按照这个索引去遍历.

### 2. append操作

别以为上面在for循环内`append()`没报错就觉得的没问题, 我们把代码改成如下看看

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
