# Mac信任网站证书[https ssl]

参考文章

1. [mac 浏览器(chrome, safari)信任自签名证书](https://www.cnblogs.com/ZhYQ-Note/p/8493848.html)

Mac下访问自签名网站时显示网站不安全.

![](https://gitee.com/generals-space/gitimg/raw/master/1b64b411b21e7dfae0c34ed821d740e9.png)

chrome 貌似没有导出证书的能力, 不过 safari 可以, 我们可以通过 safari 下载网站证书, 然后导入...linux???

以下是 safari 中的操作.

![](https://gitee.com/generals-space/gitimg/raw/master/04c8f2a9fa21cb7c366027ceb9b415cf.png)

![](https://gitee.com/generals-space/gitimg/raw/master/9b442a8600a372dd20c6af8348d288f4.png)

![](https://gitee.com/generals-space/gitimg/raw/master/ec829cb4a2e38d6d1395b57a08ff82c0.png)

信任这些网站需要管理权限, 所以需要输入密码.

![](https://gitee.com/generals-space/gitimg/raw/master/8d6ebf495cce177db2e81e4560c00555.png)

现在再查看证书, 发现已被当前用户信任.

![](https://gitee.com/generals-space/gitimg/raw/master/54ac761a839dd67c1e3decd42730baca.png)

使用"钥匙串(keychain)"工具可以查看并导出目标证书.

![](https://gitee.com/generals-space/gitimg/raw/master/695b023a257f4ba4ccf0c1790936fed0.png)

![](https://gitee.com/generals-space/gitimg/raw/master/0c609f6e37ba100b43945cc01f9da98f.png)

chrome 再次访问该网站, 还是显示"不安全", 不过还是显示已经被当前用户信任.

![](https://gitee.com/generals-space/gitimg/raw/master/3da76127f241e28639f23773f2f5c470.png)
