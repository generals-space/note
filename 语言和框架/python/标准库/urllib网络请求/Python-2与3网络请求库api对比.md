# Python-2与3网络请求库api对比

参考文章

1. [Python3中urllib使用介绍](https://blog.csdn.net/duxu24/article/details/77414298)

| python2               | python3                  |
| :-------------------- | :----------------------- |
| `urllib2`             | `urllib.request`         |
| `urllib.quote`        | `urllib.request.quote`   |
| `urlparse`            | `urllib.parse`           |
| `urlopen`             | `urllib.request.urlopen` |
| `urlencode`           | `urllib.parse.urlencode` |
| `urllib2.Request`     | `urllib.request.Request` |
| `cookielib.CookieJar` | `http.CookieJar`         |

> 注意: python3 `urllib`的`request`包导入方式应为`import urllib.request`, 而`from urllib import request`这种会出错, 显示找不到`request`属性.
