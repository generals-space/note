# Python-字符集编码转换

参考文章

1. [Python: 在Unicode和普通字符串之间转换](http://blog.csdn.net/u012448083/article/details/51918681)
2. [python将unicode和str互相转化](http://blog.csdn.net/huludan/article/details/59518325)
3. [python 中 unicode原样转成str, unicode-escape与string_escape](http://blog.csdn.net/lrz4119/article/details/45247611)
4. [根本解决Python2中unicode编码问题](https://blog.csdn.net/weixin_42989523/article/details/81873874)
5. [python 之路，致那些年，我们依然没搞明白的编码](https://www.cnblogs.com/alex3714/articles/7550940.html)
    - 最全面的讲解, 值得收藏.

环境: python 2.7

md, 目前我只见过python一种语言在定义字符串的时候可以选择定义成utf-8, unicode, ascii等格式...不明觉厉.

怎么说呢, 你只要知道, unicode范围最大, 兼容并包就好了...

如果非要排个序的话, 我觉得`unicode > utf-8 > ascii + 中文 + 各种符号`

要怎么理解才好呢? unicode范围最大, 但它每个字符所占的空间也大, 一个ascii字母和一个中文字符占的空间一样大, 在一些对数据体积十分敏感的项目中, unicode占用的空间是不可忍受的.

可是ascii包含的就那么几个字母和一些符号, 显示不了各国的语言字符.

utf-8则是一个取中妥协的方案, 里面的ascii字符还是占用比较小的空间, 其他的各国的文字字符则根据实际情况选择占用空间, 这样最方便.

> 本文讨论的是字符串变量的定义类型, 至于源文件开头添加的`#!encoding:utf-8`, 那是因为python解释器对**源码**读取时遵循的字符集默认为`ascii`.

------

```py
## 定义普通字符串, 默认是utf-8格式的
>>> a = 'abc'
## 定义unicode字符串
>>> b = u'abc'
## 输出
>>> a
'abc'
>>> b
u'abc'
## 比较, 相等, 好开心...
>>> a == b
True
```

然而, 还有中文问题, 再按照上面的步骤来一次呢?

```py
## 定义中文字符串
>>> a = '中国'
>>> b = u'中国'
## 好像有什么东西不一样了?
>>> a
'\xe4\xb8\xad\xe5\x9b\xbd'      ## 这是utf-8的编码
>>> b
u'\u4e2d\u56fd'                 ## 这是unicode的编码...
## 见鬼, 我怎么才能输出中文呢?
>>> print(a)
中国
>>> print(b)
中国
## 这下不相等了...
>>> a == b
__main__:1: UnicodeWarning: Unicode equal comparison failed to convert both arguments to Unicode - interpreting them as being unequal
False
```

在python中, 默认将字符串定义为utf-8格式的(很多语言都是), 对于字母字符串, 和其相应的ascii字符串完全相同; 遇到中文或其他语言字符, 各自占用其最小的空间.

但是, 如果将中文字符串(或包含中文的字符串)声明为unicode, 每个字符占用的空间...呵呵, 反正比utf-8格式的大.

## 2. 格式转换

so, 如何对不同格式的字符串进行转换呢? 下面是一个列表

### 2.1 将Unicode转换成普通的Python字符串 - 编码(encode)

- `unicodestring = u"Hello 中国"`
- `utf8string = unicodestring.encode("utf-8")`
- `asciistring = unicodestring.encode("ascii")`
- `isostring = unicodestring.encode("ISO-8859-1")`
- `utf16string = unicodestring.encode("utf-16")`

我们试试

```py
>>> a = '中国'
>>> b = u'中国'
>>> a == b
False
>>> c = b.encode('utf-8')
>>> a == c
True
>>> a
'\xe4\xb8\xad\xe5\x9b\xbd'
>>> b
u'\u4e2d\u56fd'
>>> c
'\xe4\xb8\xad\xe5\x9b\xbd'
```

### 2.2 将普通的Python字符串转换成Unicode - 解码(decode)

需要声明当前字符串格式...这要怎么知道??

- `plainstring1 = unicode(utf8string, "utf-8")`
- `plainstring2 = unicode(asciistring, "ascii")`
- `plainstring3 = unicode(isostring, "ISO-8859-1")`
- `plainstring4 = unicode(utf16string, "utf-16")`

再来

```py
>>> a = '中国'
>>> b = u'中国'
>>> c = unicode(a, 'utf-8')
>>> b == c
True
```

完美~
