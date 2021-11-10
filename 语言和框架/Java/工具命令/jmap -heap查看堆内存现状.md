# jmap -heap查看堆内存现状

参考文章

1. [【JVM】jmap命令详解----查看JVM内存使用详情](https://www.cnblogs.com/sxdcgaq8080/p/11089664.html)

```
jmap 139961  # pid = 139961
# 查看进程的内存映像信息,类似 Solaris pmap 命令。
jmap -heap pid
# 查看进程的详细内存占用，包括每个区域大小和使用大小
jmap -histo:live 33320
# 查看存活对象大小， 单位是字节
```
