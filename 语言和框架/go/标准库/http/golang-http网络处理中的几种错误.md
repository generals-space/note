# golang网络处理中的几种错误

参考文章

1. [Go 网络处理中的几种错误](https://romatic.net/post/go_net_errors/)
2. [net/http: request canceled (Client.Timeout exceeded while reading body)](https://github.com/golang/go/issues/37916)
3. [[译]Go net/http 超时机制完全手册](https://colobu.com/2016/07/01/the-complete-guide-to-golang-net-http-timeouts/)
    - 超厉害的文章, 值得收藏
    - 已转载 <!-- <!link!>: {2d6e6f21-4a54-44e1-9e37-4ad78d43c17c} -->

我最近在用 elasticsearch 的 exporter 的时候, 发现 ta 在请求 es 的某个 http 接口时发生了`Client.Timeout exceeded while reading body`的错误, 尤其是, 这个错误发生几次后内存会发生泄露, 导致了 oom.

按照参考文章1中的说法, 发生这个错误时, 连接是正常的, 且得到了正常的响应体, 但是在处理响应体时到了设置的超时时间, 所以连接断开, 处理操作就失败了...

```go
c := &http.Client{  
    Timeout: 15 * time.Second,
}
resp, err := c.Get("https://blog.filippo.io/")
```

## client 端

### 找不到服务器（no such host）的几种情况

**域名不存在，瞄了下代码，应该是 dns 包返回的**

```
Get http://a.b/abc: dial tcp: lookup a.b: no such host
```

**ip 不合法不会直接检查，也会返回同样错误**

```
Get http://127.0.0.1888:8080/abc: dial tcp: lookup 127.0.0.1888: no such host
```

**端口瞎填会直接报错，都不会发请求**

```
Get http://127.0.0.1:65536/abc: dial tcp: address 65536: invalid port
```

------

**拒绝连接，对方端口未监听、进程挂掉等等**

```
Get http://127.0.0.1:8080/abc: dial tcp 127.0.0.1:8080: connect: connection refused
```

**建立连接超时**

```
Get http://127.0.0:8080/abc: net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)
```

**等待返回 header， 一般是接口还在处理逻辑，没有返回任何数据；或者对方只是个普通的 tcp 服务，但不是 http**

```
Get http://127.0.0.1:8080/abc: net/http: request canceled (Client.Timeout exceeded while awaiting headers)
```

**客户端读取超时：已建立好连接，已经开始返回数据，但是body 太大太慢**

```
wait_test.go:48: net/http: request canceled (Client.Timeout exceeded while reading body)
```

## server 端

**客户端主动断开连接，服务器端在调用 Write(p []byte) (n int, err error) 时会返回**

```
wait_test.go:21: write tcp 127.0.0.1:8080->127.0.0.1:49290: write: broken pipe
```

**客户端主动断开连接，通常会直接使用 ctx.Done() 检测到，这个时候 ctx.Err() 里会拿到这个信息：**

```
context canceled
```
