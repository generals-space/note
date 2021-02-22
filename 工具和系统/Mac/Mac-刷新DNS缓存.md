# Mac-刷新DNS缓存

参考文章

1. [Mac下如何清除DNS缓存？](https://www.zhihu.com/question/19679715)
2. [Mac OS X 清除DNS缓存](https://www.cnblogs.com/qq952693358/p/9126860.html)

阿里云上修改了某个域名的解析地址, 但是 mac 上一直无法更新(其他 linux 主机上已经更新了), 需要手动执行一下清除缓存的命令.

MacOS: 10.15.7

按照参考文章1所说, 不同版本需要使用不同的命令, 只执行`sudo dscacheutil -flushcache`时, 不生效.

按照参考文章2所说, 执行如下命令就可以了.

```
sudo killall -HUP mDNSResponder
sudo killall mDNSResponderHelper
sudo dscacheutil -flushcache
```
