# golang-missing Location in call to Time.In异常

参考文章

1. [panic: time: missing Location in call to Time.In](https://www.cnblogs.com/siaslfslovewp/p/11219470.html)

用的是 busybox 容器, 里面没有 timezone 那些文件, 换成 centos 或是将宿主机上的`/usr/share/`
