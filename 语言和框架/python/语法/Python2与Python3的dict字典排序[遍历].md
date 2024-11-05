# Python2与Python3的dict字典排序

参考文章

1. [Python 字典为什么是无序的?](https://www.zhihu.com/question/24306558/answer/374075597)
    - Python 3.6改写了dict的内部算法, 因此3.6的dict是有序的, 在此版本之前皆是无序
    - 低版本希望有序dict可以用`collections.OrderedDict`, 这是早在2.7就实装了的容器, 它同样是遵循插入顺序.

直到 python3.6 开始, dict 的字典才变得有序, 且是按插入顺序. 在此之前的版本中, dict 都是无序的.

当然, 其实也不应该说是无序, 而应该叫"哈希序", 因为多次遍历同一个 dict 对象, 顺序都是相同的. 相比于 golang 的乱序, 每次遍历的顺序都不一样, 好了不知道多少倍.

```py
## 预置的字典与之后插入内容的字典遍历行为相同.
## dic = {
##     'c': 12,
##     'd': 34,
##     'b': 56,
##     'a': 78,
## }

dic = {}
dic['c'] = '12'
dic['d'] = '34'
dic['b'] = '56'
dic['a'] = '78'

## 下面的两种遍历方式相同
## for i in dic.keys():
##     print('key: %s, value %s' % (i, dic[i]))
for k, v in dic.items():
    print('key: %s, value %s' % (k, v))

```

python2的执行结果

```log
key: a, value 78
key: c, value 12
key: b, value 56
key: d, value 34
```

python3的执行结果

```log
key: c, value 12
key: d, value 34
key: b, value 56
key: a, value 78
```
