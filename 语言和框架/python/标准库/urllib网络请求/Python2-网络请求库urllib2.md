# Python-网络请求库urllib与urllib2

参考文章

1. [urllib官方文档](https://docs.python.org/2.7/library/urllib.html#module-urllib)

2. [urlib2官方文档](https://docs.python.org/2.7/library/urllib2.html?highlight=urllib#module-urllib2)

3. [Python urllib与urllib2](http://www.cnblogs.com/wang-can/p/3591116.html)
4. [Handling urllib2's timeout? - Python](https://stackoverflow.com/questions/2712524/handling-urllib2s-timeout-python)
    - urllib2 超时设置, 补测有效.

> Python中包含了两个网络模块，分别是urllib与urllib2，urllib2是urllib的升级版，拥有更强大的功能。

关于`urllib2`比`urllib`更强大的传言, 我目前只发现前者在使用`urlopen()`函数访问url时可以自定义请求头, 比如添加Cookie字段模拟登录, 其他的我还真没看出来.

## 1. `urllib2.urlopen()`

```py
urllib2.urlopen(url[, data[, timeout[, cafile[, capath[, cadefault[, context]]]]])
```

1. `url`: 可以直接是url字符串, 也可以是一个Request对象. 通过创建Request对象, 可以自定义请求头, 也可以携带POST请求数据.
2. `data`: ...感觉和Request对象的data参数有重复的地方, 不过如果你只想创建POST请求又不需要改请求头, 直接用这个data就好了.

```py
import urllib2
import json
obj = {
    'a': 123,
    'b': 234
}

reqHeader = {
    'Content-Type': 'application/json'
}
req = urllib2.Request(operateAddr, data = obj, headers = reqHeader) 
result = urllib2.urlopen(req).read()
```

## 2. urllib2.Request

```py
urllib2.Request(url[, data][, headers][, origin_req_host][, unverifiable])
```

> 如果响应头中有`Content-Type`字段为`text/json`类型, 貌似就不用在`urlopen()`后调用`read()`而是直接得到字符串.<???>

```py
import urllib2

url = 'http://www.baidu.com'
headers = {
    'User-Agent':"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0",
    'Referer':"http://www.baidu.com"
}
req = urllib2.Request(url, headers=headers)
res = urllib2.urlopen(req)
```
