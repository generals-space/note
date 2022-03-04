# VMware桥接模式bridge配置静态IP

参考文章

1. [VMware桥接网络模式配置静态IP详细步骤](https://blog.csdn.net/ganpuzhong42/article/details/77775145)

宿主机: win10
宿主机网络: wifi
VMware版本: 15
虚拟机系统: CentOS 8

本来我只主要设置了如下三个字段

```
BOOTPROTO=static
IPADDR=192.168.0.101
NETMASK=255.255.255.0
```

但是重启网络服务后发现只能ping通内网的地址.

按照参考文章1, 我在宿主机找到无线网卡对应的网关, 再设置`GATEWAY`字段再重启网络服务, 就可以了.
