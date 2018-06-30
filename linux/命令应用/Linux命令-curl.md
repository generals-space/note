# Linux命令-curl

参考文章

1. [Linux curl命令参数详解](http://www.aiezu.com/system/linux/linux_curl_syntax.html)

2. [在linux下使用curl访问 多参数url GET参数问题](http://blog.csdn.net/sunbiao0526/article/details/6831327)

3. [shell curl 数据中含有空格 如何提交](https://blog.csdn.net/qq_25279717/article/details/71577313)

## 1. 请求参数中`&`的处理

假设url为`http://mywebsite.com/index.PHP?a=1&b=2&c=3`, web形式下访问url地址，使用`$_GET`是可以在后台获取到所有的参数

而在Linux下使用`curl http://mywebsite.com/index.php?a=1&b=2&c=3`, `$_GET`只能获取到参数`a`. 由于url中有`&`，其他参数获取不到，必须对&进行下转义才能`$_GET`获取到所有参数

`curl http://mywebsite.com/index.php?a=1\&b=2\&c=3`

## 2. 使用`-H`添加请求头信息

尤其是`proxy`, `ua`和`cookie`的配置.

可以分别使用`--cookie`, `--user-agent`配置, 也可以直接使用`-H 'User-Agent: UA字符串'`的形式.

```
curl -I -H 'Cookie: _ga=GA1.2.1337029376.1526882292; session_id=PvAvY-4CYs463Y;' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36' www.baidu.com
```

但是有一点需要注意, 'Cookie: _ga=GA1.2.1337029376.1526882292; session_id=PvAvY-4CYs463Y;'字符串不能作为一个变量传入, `User-Agent`同理, 因为这样curl执行会报错. 如下

```
cookie_str='Cookie: _ga=GA1.2.1337029376.1526882292; session_id=PvAvY-4CYs463Y;'
ua_str='User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36'
curl -I -H $cookie_str -H $ua_str www.baidu.com
curl: (6) Could not resolve host: _ga=GA1.2.1337029376.1526882292;
curl: (6) Could not resolve host: session_id=PvAvY-4CYs463Y;
curl: (6) Could not resolve host: Mozilla
curl: (6) Could not resolve host: (Macintosh;
curl: (6) Could not resolve host: Intel
curl: (6) Could not resolve host: Mac
curl: (6) Could not resolve host: OS
curl: (6) Could not resolve host: X
curl: (6) Could not resolve host: 10_13_2)
curl: (6) Could not resolve host: AppleWebKit
curl: (6) Could not resolve host: (KHTML,
curl: (6) Could not resolve host: like
curl: (6) Could not resolve host: Gecko)
curl: (6) Could not resolve host: Chrome
curl: (6) Could not resolve host: Safari
HTTP/1.1 200 OK
Accept-Ranges: bytes
Cache-Control: private, no-cache, no-store, proxy-revalidate, no-transform
Connection: Keep-Alive
Content-Length: 277
Content-Type: text/html
Date: Wed, 27 Jun 2018 11:22:19 GMT
Etag: "575e1f7c-115"
Last-Modified: Mon, 13 Jun 2016 02:50:36 GMT
Pragma: no-cache
Server: bfe/1.0.8.18
```

通过变量传入的字符串会被以空格分隔开, 所以会出错...不过写在行内也真够low的.