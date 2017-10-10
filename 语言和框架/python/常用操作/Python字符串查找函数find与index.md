# Python字符串查找函数find与index

参考文章

1. [python字符串字串查找 find和index方法](http://outofmemory.cn/code-snippet/6682/python-string-find-or-index)

python 字符串查找有4个方法

1. find

2. index

3. rfind方法

4. rindex方法.

## 1. find()/rfind()

查找子字符串, 若找到返回从0开始的下标值, 若找不到返回-1

```py
info = 'abca'
print info.find('a')        ##从下标0开始, 查找在字符串里第一个出现的子串, 返回0

info = 'abca'
print info.find('a', 1)     ##从下标1开始, 查找在字符串里第一个出现的子串: 返回结果3

info = 'abca'
print info.find('333')      ##返回-1, 查找不到返回-1
```

## 2. index()/rindex

python 的index方法是在字符串里查找子串第一次出现的位置, 类似字符串的find方法, 不过比find方法更好的是, 如果查找不到子串, 会抛出异常, 而不是返回-1

```py
info = 'abca'
print info.index('a')
print info.index('33')
```

rfind和rindex方法用法和上面一样, 只是从字符串的末尾开始查找。