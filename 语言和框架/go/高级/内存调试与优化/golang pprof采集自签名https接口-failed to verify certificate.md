# golang pprof采集自签名https接口-failed to verify certificate

参考文章

1. [cmd/pprof: add HTTPS support with client certificates](https://github.com/golang/go/issues/20939)
    - `go tool pprof -seconds 5 https+insecure://192.168.99.100:32473/debug/pprof/profile`

go: 1.22.8

## 问题描述

程序运行是精简版的 docker 容器, 里面没有 go tool pprof 工具包, 而其web端口是由 nginx 做反向代理暴露出来的 https 链接, 本地访问时必须通过 https 协议完成.

又由于该证书是自签名, 本地开发环境不信任该证书, 因此会报如下错误

```log
[root@dc650b0ce3ca ~]# go tool pprof --seconds 30 https://192.168.72.2:3001/debug/pprof/heap
Fetching profile over HTTP from https://192.168.72.2:3001/debug/pprof/heap?seconds=30
Please wait... (30s)
https://192.168.72.2:3001/debug/pprof/heap: Get "https://192.168.72.2:3001/debug/pprof/heap?seconds=30": tls: failed to verify certificate: x509: cannot validate certificate for 192.168.72.2 because it doesn't contain any IP SANs
failed to fetch any source profiles
```

就类似于通过 curl 访问此类链接时需要指定`-k`参数忽略该证书一样, go tool 工具包是否提供了这样跳过验证的机制?

```log
$ curl https://192.168.72.2:3001/debug/pprof
curl: (60) Peer's Certificate issuer is not recognized.
More details here: http://curl.haxx.se/docs/sslcerts.html

curl performs SSL certificate verification by default, using a "bundle"
 of Certificate Authority (CA) public keys (CA certs). If the default
 bundle file isn't adequate, you can specify an alternate file
 using the --cacert option.
If this HTTPS server uses a certificate signed by a CA represented in
 the bundle, the certificate verification probably failed due to a
 problem with the certificate (it might be expired, or the name might
 not match the domain name in the URL).
If you'd like to turn off curl's verification of the certificate, use
 the -k (or --insecure) option.
```

## 解决方案

见参考文章1, 完美解决.
