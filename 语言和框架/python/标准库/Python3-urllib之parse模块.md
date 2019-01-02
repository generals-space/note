# Python3-urllib之parse模块

参考文章

1. [python3 urllib.parse 常用函数](http://www.cnblogs.com/mengyu/p/7722883.html)

## 1. `urlparse`: 获取查询参数

```py
from urllib import parse

url = 'https://docs.python.org/3.5/library/urllib.parse.html?highlight=parse#module-urllib.parse'
result = parse.urlparse(url)

# 获取返回结果参数内容
print(result.query) ## highlight=parse
# 结果转换成字典
print(parse.parse_qs(result.query)) ## {'highlight': ['parse']}
# 结果转换成列表
print(parse.parse_qsl(result.query)) ## [('highlight', 'parse')]
```

`result`是一个`ParseResult`对象. ta的打印结果如下

```py
ParseResult(scheme='https', netloc='docs.python.org', path='/3.5/library/urllib.parse.html', params='', query='highlight=parse', fragment='module-urllib.parse')
```

## 2. `quote`与`unquote`编解码

```py
from urllib import parse
print(parse.quote('@')) ## %40
print(parse.unquote('%40')) ## @
```