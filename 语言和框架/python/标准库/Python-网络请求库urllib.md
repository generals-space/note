# Python-网络请求库urllib

参考文章

1. [Python3中urllib使用介绍](https://blog.csdn.net/duxu24/article/details/77414298)

Py2.x：

- urllib库

- urllin2库

Py3.x：

- urllib库

变化：

在Pytho2.x中使用`import urllib2`——-对应的，在Python3.x中会使用`import urllib.request`，`urllib.error`。
在Pytho2.x中使用`import urllib`——-对应的，在Python3.x中会使用import urllib.request，urllib.error，urllib.parse。
在Pytho2.x中使用`import urlparse`——-对应的，在Python3.x中会使用import urllib.parse。
在Pytho2.x中使用`import urlopen`——-对应的，在Python3.x中会使用import urllib.request.urlopen。
在Pytho2.x中使用`import urlencode`——-对应的，在Python3.x中会使用import urllib.parse.urlencode。
在Pytho2.x中使用`import urllib.quote`——-对应的，在Python3.x中会使用import urllib.request.quote。
在Pytho2.x中使用`cookielib.CookieJar`——-对应的，在Python3.x中会使用http.CookieJar。
在Pytho2.x中使用`urllib2.Request`——-对应的，在Python3.x中会使用urllib.request.Request。

> 注意: python3 `urllib`的`request`包导入方式应为`import urllib.request`, 而`from urllib import request`这种会出错, 显示找不到`request`属性.