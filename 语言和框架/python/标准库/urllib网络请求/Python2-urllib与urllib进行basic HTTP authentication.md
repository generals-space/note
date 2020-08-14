# Python2-urllib与urllib进行basic HTTP authentication

参考文章

1. [curl 的用法指南](http://www.ruanyifeng.com/blog/2019/09/curl-reference.html)
    - `curl -u`参数的两种格式与`Authorization`请求头的格式
    - `curl -u 'bob:12345' https://google.com/login`
    - `curl https://bob:12345@google.com/login`
2. [Python urllib2, basic HTTP authentication, and tr.im](https://stackoverflow.com/questions/635113/python-urllib2-basic-http-authentication-and-tr-im)

```
curl -u elastic:changeme 127.0.0.1:9200/_cat/health
```

这里的`-u`为`--user`, url格式与`http://elastic:changeme@127.0.0.1:9200/_cat/health`相同.

在使用`urllib.urlopen()`时, 可以正常请求.

```py
>>> resp = urllib.urlopen('http://name:123456@www.baiudu.com')
>>> resp.getcode()
200
```

但是`urllib2.urlopen()`是无法识别`name:password@domain`这种url格式的, 而且p事特多.

```py
>>> resp = urllib2.urlopen('http://name:123456@www.baiudu.com')
Traceback (most recent call last):
## 省略
## ...
httplib.InvalidURL: nonnumeric port: '123456@www.baiudu.com'
## 或是
URLError: <urlopen error [Errno -2] Name or service not known>
```

> 有的平台是`httplib.InvalidURL`, 有的则是`URLError`异常.

参考文章2给出了通过设置`Authorization`请求头实现这类 url 请求的示例.

```py
import urllib2, base64

req = urllib2.Request('http://api.foursquare.com/v1/user')
base64string = base64.encodestring('%s:%s' % (username, password)).replace('\n', '')
req.add_header('Authorization', 'Basic %s' % base64string)   
resp = urllib2.urlopen(req)
```

实测有效.
