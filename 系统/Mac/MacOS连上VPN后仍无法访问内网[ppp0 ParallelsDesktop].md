# MacOS连上VPN后仍无法访问内网[ppp0 ParallelsDesktop]

参考文章

1. [macOS Monterey 连接vpn后无法联网问题，解决方案](https://www.jianshu.com/p/cca31b474c61)

MacOS: Catalina 10.15.7

## 问题描述

开启 ParallelsDesktop 虚拟机后, 再连宿主机上的 VPN, 显示连接上, 但是无法访问到公司内网, 其他网络倒一切正常.

要先把虚拟机关掉, 且终止 ParallelsDesktop 进程, 再连才可以.

## 排查思路

连接上vpn后, 会多出一个`ppp0`接口.

```log
$ ifconfig
ppp0: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 1280
	inet 10.10.200.16 --> 10.10.200.1 netmask 0xffffff00
```

但是查看路由表, 发现默认路由没变, 还是原来网关.

```log
$ route -n get default
   route to: default
destination: default
       mask: default
    gateway: 172.18.239.1
  interface: en0
      flags: <UP,GATEWAY,DONE,STATIC,PRCLONING>
 recvpipe  sendpipe  ssthresh  rtt,msec    rttvar  hopcount      mtu     expire
       0         0         0         0         0         0      1500         0
```

所以问题应该出现在路由表.

## 解决方法

参考文章1中所说的方法, 实践确认有效, 需要调整网络接口列表的顺序...

连接成功后, 调整VPN接口到Wifi接口之前, 应用之后就可以了.
