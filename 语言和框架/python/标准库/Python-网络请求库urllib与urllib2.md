# Python-网络请求库urllib与urllib2

参考文章

1. [urllib官方文档](https://docs.python.org/2.7/library/urllib.html#module-urllib)

2. [urlib2官方文档](https://docs.python.org/2.7/library/urllib2.html?highlight=urllib#module-urllib2)

3. [Python urllib与urllib2](http://www.cnblogs.com/wang-can/p/3591116.html)

4. 

> Python中包含了两个网络模块，分别是urllib与urllib2，urllib2是urllib的升级版，拥有更强大的功能。

关于`urllib2`比`urllib`更强大的传言, 我目前只发现前者在使用`urlopen()`函数访问url时可以自定义请求头, 比如添加Cookie字段模拟登录, 其他的我还真没看出来.

## 1. urllib

> 官方文档提示: `urllib`模块在Python 3中被拆分成`urllib.request`, `urllib.parse`和`urllib.error`. Python 3中的`urllib.request.urlopen()`等同于`urllib2.urlopen()`, `urllib.urlopen()`已经被移除了.

### 1.1 urllib.urlopen

普通http访问请求, GET/POST

```py
urllib.urlopen(url[, data[, proxies[, context]]])
```

1. `url`: 目标url字符串

2. `data`: POST请求中携带的数据, 其格式必须为`application/x-www-form-urlencoded`

3. `proxies`: 代理服务器配置

4. `context`: `2.7.9`版本新加的特性, 是一个`ssl.SSLContext`实例, 在创建`https`连接时使用.

注意: 当不传入`data`参数时, 请求类型默认为GET, 当传入`data`时, 此次请求自动转换为POST类型.

另外, 关于`data`的类型为`application/x-www-form-urlencoded`这种, 就是在html页面中单纯通过`form`表单和`submit`按钮原生提交的格式. 同时浏览器会静默在请求头中添加`Content-Type`字段.

与之对应的是我们通常在使用ajax时直接以json格式发送请求. 这种情况下请求头中`Content-Type`的类型为`application/json`.

`urllib.urlopen()`的`data`参数只能为原生form类型, 因为它不能携带额外的请求头. 而`urllib2.urlopen()`可以, 到时在请求头中加上`Content-Type`字段, 其值为`application/json`即可.

但话说回来, 怎么构造`application/x-www-form-urlencoded`格式的参数变量?

`urllib`提供了一个方法`urlencode()`来做这件事.


### 1.2 urllib.urlencode

```py
urllib.urlencode(query[, doseq])
```

将传入的字典类型变量, 或是二元组列表变量转化为`urlopen()`函数可以接受的字符串类型, 结果大概就是`a=1&b=2&c=3`这种吧.

**示例1**

```py
>>> import urllib
>>> dic = {'name': 'general', 'age': 24, 'sex': 'male'}
>>> urllib.urlencode(dic)
'age=24&name=general&sex=male'
>>> 
```

**示例2**

```py
>>> import urllib
>>> lis = [('id', 2), ('name', 'general'), ('id', 3), ('name', 'jiangming')]
>>> urllib.urlencode(lis)
'id=2&name=general&id=3&name=jiangming'
>>> 
```

关于示例2的应用场景, 想想批量删除接口中, 希望获取到前端传入的多个目标对象id, 只能通过这种形式.

### 1.3 urllib.urlretrieve

下载文件函数

```py
urllib.urlretrieve(url[, filename[, reporthook[, data]]])
```

1. `url`: 目标资源url

2. `filename`: 本地路径, 用于存放下载的数据

3. `reporthook`: 回调函数, 调用时机还不太清楚, 反正用得也不多

4. `data`: 同`urlopen()`中的`data`一样, 也是POST请求中携带的数据对象

```py
>>> import urllib
>>> urllib.urlretrieve('https://gitimg.generals.space/54e181029ee23ae664a10fa3ef1ad5b9.png', '/tmp/1.png')
('/tmp/1.png', <httplib.HTTPMessage instance at 0x7f0bceee1878>)
```

### 1.4 urllib中的一些字符编码的辅助函数

`urllib.quote(string[, safe])`: 对字符串进行编码。参数safe指定了不需要编码的字符;

`urllib.unquote(string)`: 对字符串进行解码；

`urllib.quote_plus(string [, safe ])`: 与`urllib.quote`类似，但这个方法用'+'来替换' '，而quote用'%20'来代替' '

`urllib.unquote_plus(string )`: 对字符串进行解码；

`urllib.urlencode(query[, doseq])`: 将dict或者包含两个元素的元组列表转换成url参数。例如 字典{'name': 'dark-bull', 'age': 200}将被转换为"name=dark-bull&age=200"

`urllib.pathname2url(path)`: 将本地路径转换成url路径；

`urllib.url2pathname(path)`: 将url路径转换成本地路径

## 2. urllib2

### 2.1 urllib2.urlopen

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

### 2.2 urllib2.Request

```py
urllib2.Request(url[, data][, headers][, origin_req_host][, unverifiable])
```

> 如果响应头中有`Content-Type`字段为`text/json`类型, 貌似就不用在`urlopen()`后调用`read()`而是直接得到字符串.<???>