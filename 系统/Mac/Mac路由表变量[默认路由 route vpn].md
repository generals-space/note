安装了公司的VPN工具, 一拨vpn就修改了默认路由, 劫持了全部流量, 连本地的虚拟机都没法访问了...

```log
general@MacBook-Pro:~$ netstat -rn
Routing tables

Internet:
Destination        Gateway            Flags               Netif Expire
default            192.168.43.1       UGScg                 en0
default            link#22            UCSIg               utun4
10.37.129.2        ca.89.f3.da.ee.66  UHLWIi                lo0
117.159.206.192/32 link#22            UCS                 utun4
127                127.0.0.1          UCS                   lo0
127.0.0.1          192.168.43.1       UGHS                  en0
127.0.0.1/32       link#22            UCSI                utun4
169.254            link#15            UCS                   en0      !
172.18.80.137/32   link#22            UCS                 utun4
172.18.80.138/31   link#22            UCS                 utun4
172.20.119.182/32  link#22            UCS                 utun4
172.20.126.193/32  link#22            UCS                 utun4
172.20.126.202/32  link#22            UCS                 utun4
172.22.16/21       link#22            UCS                 utun4
172.22.24/22       link#22            UCS                 utun4
172.22.28/23       link#22            UCS                 utun4
172.22.30/24       link#22            UCS                 utun4
172.22.33/24       link#22            UCS                 utun4
172.22.34/23       link#22            UCS                 utun4
172.22.36/22       link#22            UCS                 utun4
172.22.40/24       link#22            UCS                 utun4
172.22.128/20      link#22            UCS                 utun4
172.22.160.126/32  link#22            UCS                 utun4
172.22.161/24      link#22            UCS                 utun4
172.22.161.200     link#22            UHW3I               utun4      7
192.168.7.60       192.168.7.60       UH                  utun4
```

经过对比, 发现连接上vpn后多出了如下网络设备.

```log
$ ifconfig | tail
utun4: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 1400
	inet 192.168.7.63 --> 192.168.7.63 netmask 0xffffffff
```

> utun设备的地址, 在每次连接vpn后都会变动.

ok, 现在需要删除默认路由.

```bash
sudo route delete default -ifscope utun4
```

一定要加`-ifscope`, 因为在默认路由项中, `utun4`的网关地址为`link#22`, 但这玩意根本不是一个合法的地址.

具体的我也不懂, 貌似跟链路层有关, 以后再说吧.
