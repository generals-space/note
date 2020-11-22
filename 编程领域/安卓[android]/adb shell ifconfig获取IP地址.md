# adb shell ifconfig获取IP地址

参考文章

1. [adb 命令获取手机ip地址](https://www.jianshu.com/p/17fcca5d8ea4)
2. [在linux下如何用正则表达式执行ifconfig命令，只提取IP地址！](https://blog.51cto.com/gotoo/1979005)

```console
$ adb shell ifconfig wlan0
wlan0     Link encap:UNSPEC    Driver icnss
          inet addr:172.18.239.24  Bcast:172.18.239.63  Mask:255.255.255.192
          inet6 addr: fe80::2247:daff:fec1:9d66/64 Scope: Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:58784355 errors:0 dropped:6 overruns:0 frame:0
          TX packets:82482288 errors:0 dropped:480 overruns:0 carrier:0
          collisions:0 txqueuelen:3000
          RX bytes:41817901306 TX bytes:89632700208
```

```console
$ adb shell ifconfig wlan0 | grep 'inet addr' | awk -F'[ :]+' '{print $4}'
172.18.239.24
```
