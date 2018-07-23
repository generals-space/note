# Linux-curl&wget忽略https证书

参考文章

1. [curl wget 不验证证书进行https请求](https://blog.csdn.net/bytxl/article/details/46989667)

一些自签名证书的网站不被浏览器信任, 同时也不会被`curl`, `wget`命令行工具信息, 但是我们可以忽略对目标网站证书的验证.

```
wget --no-check-certificate https://x.x.x.x
curl -k https://x.x.x.x
```