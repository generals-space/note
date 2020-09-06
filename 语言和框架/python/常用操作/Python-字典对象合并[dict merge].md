# Python-字典对象合并

参考文章

1. [Python中如何实现两个字典合并](http://www.pythoner.com/13.html)

虽然觉得字典合并的操作貌似非常简单, 但还是想找一个可用的方法, 就像js里的`Object.assign()`.

```py
dict1 = {
    'name':'xxx', 
    'age': 21 
}
dict2 = { 
    'name': 'general',
    'birthday': '2017-11-22'
}
```

## 1.

```py
## items()的结果为dict_items对象, 实际是列表对象的变种, 但不支持直接相加
print(dict1.items())    ## dict_items([('name', 'xxx'), ('age', 21)])
new_dic = dict(list(dict1.items()) + list(dict2.items()))
print(new_dic)          ## {'name': 'general', 'age': 21, 'birthday': '2017-11-22'}
```

## 2. 

```py
new_dic = dict(dict1, **dict2)
print(new_dic) ## {'name': 'general', 'age': 21, 'birthday': '2017-11-22'}
```

> 这种方法对`dict1`, `dict2`原始值没有影响.

## 3. 

```py
new_dic = dict1.copy()
new_dic.update(dict2)
print(new_dic) ## {'name': 'general', 'age': 21, 'birthday': '2017-11-22'}
```

或

```py
new_dic = dict(dict1)
new_dic.update(dict2)
print(new_dic) ## {'name': 'general', 'age': 21, 'birthday': '2017-11-22'}
```

------

不过都是浅拷贝.

建议使用第2种方法.
