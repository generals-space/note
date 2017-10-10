# powershell网络操作

## 1. 访问网页/下载文件

```ps1
(New-Object System.Net.WebClient).downloadstring('https://www.baidu.com')
(New-Object System.Net.WebClient).downloadfile('https://www.baidu.com', 'index.html')
```

其中`downloadstring()`方法会在控制台直接输出页面的内容, 而`downloadfile()`则必须指定下载文件的路径.

> 注意: `downloadfile()`方法下载的文件默认存放在`user`目录, 就算指定相对路径, 也是基于`user`目录的, 这点与linux下不同.

------

还有一种, 使用`invoke-webrequest`方法.

```ps
$ Invoke-WebRequest -usebasicparsing 'https://www.baidu.com'


StatusCode        : 200
StatusDescription : OK
Content           : <html>
                    <head>
                        <script>
                                location.replace(location.href.replace("https://","http://"));
                        </script>
                    </head>
                    <body>
                        <noscript><meta http-equiv="refresh" content="0;url=http://www.baidu.com/"></...
RawContent        : HTTP/1.1 200 OK
                    Connection: keep-alive
                    Content-Length: 227
                    Content-Type: text/html
                    Date: Mon, 15 May 2017 10:20:27 GMT
                    Last-Modified: Mon, 08 May 2017 03:48:00 GMT
                    Set-Cookie: BD_NOT_HTTPS=1; pa...
Forms             :
Headers           : {[Connection, keep-alive], [Content-Length, 227], [Content-Type, text/html], [Date, Mon, 15 May 201
                    7 10:20:27 GMT]...}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        :
RawContentLength  : 227
```

上面使用`invoke-webrequest`直接输出了请求响应, 当我们希望将其输出到文件时还需要指定`-outfile`参数.

```ps
$ Invoke-WebRequest -usebasicparsing -uri 'https://www.baidu.com' -outfile 'index.html'
$ ls


    目录: C:\Users\general


Mode                LastWriteTime         Length Name
----                -------------         ------ ----
...
-a----        2017/5/15     20:08            227 index.html
```