# MacOS-realvnc连接失败too many security failures

参考文章

1. [VNC连接报错“too many security failures”](https://www.cnblogs.com/jiading/articles/12695171.html)

## 问题描述

![](https://gitee.com/generals-space/gitimg/raw/master/8e26fda125cde5e37b1380ed1d16ccd7.png)

## 解决方案

按照参考文章1所说, 停止重启就可以了.

不过我这边是通过MacOS自带的服务管理启动的, 无法使用`vncserver -kill :1`杀进程, 所以只能用常规`kill`命令完成.

```console
$ ps -ef | grep vnc
    0   355     1   0 22Mar22 ??         0:28.11 /Library/vnc/vncserver -service
  501 60512     1   0 10:10AM ??         0:00.08 /Library/vnc/vncagent service
  501 60518 60512   0 10:10AM ??         0:00.50 /Library/vnc/VNC Server.app/Contents/MacOS/vncserverui service 11
```

我们先kill "vncserver -service"这个服务, 然后依次kill下面两个, 等待ta自动重新启动就可以了.

```
sudo kill 355
```
