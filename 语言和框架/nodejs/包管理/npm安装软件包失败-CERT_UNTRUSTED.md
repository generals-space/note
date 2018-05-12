# npm安装软件包失败-CERT_UNTRUSTED

参考文章

1. [npm 设置淘宝的npm源后更新包提示证书错误](https://segmentfault.com/q/1010000005079915)

npm安装第三方包失败, 查看它提供的日志时找到`Error: CERT_UNTRUSTED`错误, 应该是一种证书错误.

按照参考文章中[dreamstu](https://segmentfault.com/u/dreamstu)的回答, 执行`npm config set strict-ssl false`可以安装成功(如果已经有`node_modules`目录最好删掉重新安装).