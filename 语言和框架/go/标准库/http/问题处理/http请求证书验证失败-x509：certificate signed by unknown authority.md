# x509：certificate signed by unknown authority

参考文章

1. [go调用Https时出错:certificate signed by unknown authority](http://www.pangxieke.com/linux/go-certificate-signed-by-unknown-authority.html)
2. [关于CentOS6访问https证书错误返回x509: certificate signed by unknown authority](http://lilybug.cn/post/63.html)
    - `update-ca-trust`不生效

## 问题描述

发起 http 请求失败(根本得不到响应)

```go
    resp, err := client.Get("https://localhost:8081")
```

```log
Get "": x509: certificate signed by unknown authority (possibly because of "x509: cannot verify signature: algorithm unimplemented" while trying to verify candidate authority certificate "ZeroSSL RSA Domain Secure Site CA")
```

## 解决方法

参考文章2中的更新本地证书的方法不管用

```
yum install -y ca-certificates
update-ca-trust
```

关闭客户端对服务器证书的校验.

```go
    tr := &http.Transport{
        TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    resp, err := client.Get("https://localhost:8081")
```

成功
