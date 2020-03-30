# go get设置ss代理cow

参考文章

1. [教你如何让go get畅通无阻](https://studygolang.com/articles/9490)

2. [Cow](https://github.com/cyfdecyf/cow/)

`go get`可以设置代理, 不过只能设置`http/https`代理, 但是`http/https`代理是没法翻墙的, 只能靠ss等工具实现, 但是ss开放的端口不能被`go get`直接使用, 于是就有了`Cow`工具的用处.

`Cow`可以将本地的ss端口做一个转发, 将socks5端口映射为http端口.

github仓库中有提供下载方法, 然后就是配置文件了, `Cow`的配置文件行数很多, 注释也很多, 不过其实只需要修改2个字段: `listen`和`proxy`.

- `listen`: 表示cow本身提供的http代理端口, `go get`要设置的环境变量地址就是这个值

- `proxy`: 表示本地ss监听的端口, 一般是1080. cow会将listen端口发来的请求通过proxy地址转发出去.

```ini
listen = http://0.0.0.0:7777
proxy = socks5://127.0.0.1:1080
```

然后设置环境变量, 在windows下, 只能通过`计算机` -> `属性` -> `高级系统设置` -> `环境变量`.

```ini
http_proxy=http://127.0.0.1:7777
https_proxy=http://127.0.0.1:7777
```

设置完成后可以在命令行中通过`echo %proxy_http%`可以打印出变量值.

> `git`也同样使用这两个环境变量, 所以也可以极大加快github的反应速度.

