# Python高级语法-for...in...if语句.md

参考文章

[Python中在for循环中嵌套使用if和else语句的技巧](http://www.jb51.net/article/86987.htm)

Python中，`for...if...`语句一种简洁的**构建List**的方法，从`for`给定的List中选择出满足`if`条件的元素组成新的List

简单示例

```python 
#!/usr/bin/env python
#!coding:utf-8

oldlist = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

## 从oldlist中选出奇数元素
newlist = [x for x in oldlist if x % 2]
print(newlist)
```

执行结果为

```
[1, 3, 5, 7, 9]
```

在`newlist = ...`这一句中, 第一个`x`表示即将成为newlist成员的元素, 它需要满足之后的`if`条件. 第2, 3个`x`表示oldlist中的所有成员元素, 遍历过程中为每个x都做`if`语句的检测, 满足条件的x将变成返回到第1个`x`的位置.

------

还有一种, oldlist是一个列表/字典, 从其中选择部分成员/键作为新列表的成员.

```python
#!/usr/bin/env python
#!coding:utf-8

oldlist = { 
    '1': 'a',
    '2': 'b',
    '3': 'c',
    '4': 'd',
    '5': 'e',
    '6': 'f',
    '7': 'g',
    '8': 'h',
    '9': 'i' 
}

## 从oldlist中选出奇数键对应的值
newlist = [oldlist.get(x) for x in oldlist if int(x) % 2]
print(newlist)
newlist = [{x: oldlist.get(x)} for x in oldlist if int(x) % 2]
print(newlist)
```

## 局限

很多时候我想在列表推导式中使用自定义的方法去处理目标列表对象, 但是如果方法里有抛出异常, 从而过滤掉某个列表成员, 列表推导式就无法实现了. 这种情况下, 列表推导式默认返回`None`, 这可不是我想要的.

我曾想过在自定义方法外设置一个全局变量, 在方法内部根据是否有异常情况对这个变量进行修改, 然后在列表推导式中通过if子句判断这个变量的True/False决定是否放弃这次对列表成员的求值, 但是很不幸, 并不能完成. `flag`标识一但为False, 列表推导式就无法进行. 可以说, 推导式会先进行一次判断, 再决定本次遍历是否继续.

```py
flag = False
def func(n):
    global flag
    print(n)
    if n % 2:
        flag = True
        return n
    else:
        flag = False
        return

a = [1, 2, 3, 4]

print([func(n) for n in a if flag == True])
```