# Windows下查看网卡接口索引

<!keys!>: <!接口索引!>

参考文章

1. [Windows下查看网卡“接口索引”的方法](http://news.newhua.com/news/2010/0915/102207.shtml)

有时在路由命令添加路由时需指定网卡的接口索引，例如

```
route ADD 157.0.0.0 MASK 255.0.0.0  157.55.80.1 METRIC 3 IF 2
```

其中`IF`后面的数字用于指定网卡接口, 感觉类似于linux下的eth0。如果一台机器有多个网卡，如果知道每块网卡的索引编号呢？

用`ipconfig /all`是看不到的. 

## 1. arp

一种方法是使用`arp`命令, 如下 

```
[c:\~]$ arp -a
接口: 172.32.100.1 --- 0xc
  Internet 地址         物理地址              类型
  172.32.100.232        00-0c-29-9b-ba-0a     动态        
  172.32.100.255        ff-ff-ff-ff-ff-ff     静态        
  224.0.0.2             01-00-5e-00-00-02     静态        
   
接口: 10.96.0.114 --- 0x12
  Internet 地址         物理地址              类型
  192.168.0.2           bb-bb-bb-bb-bb-00     动态        
  192.168.0.3           bb-bb-bb-bb-bb-00     动态        

```

其中每块网卡（接口）IP地址后面跟随的16进制数字就是网卡的接口索引, 上述如`0xc`, `0x12`。

## 2. route

```
[c:\~]$ route print
===========================================================================
接口列表
  4...9e b6 d0 0f 85 2b ......Microsoft Wi-Fi Direct Virtual Adapter
 26...00 50 56 c0 00 01 ......VMware Virtual Ethernet Adapter for VMnet1
 12...00 50 56 c0 00 08 ......VMware Virtual Ethernet Adapter for VMnet8
  7...00 50 56 c0 00 02 ......VMware Virtual Ethernet Adapter for VMnet2
  3...9c b6 d0 0f 85 2b ......Killer Wireless-n/a/ac 1535 Wireless Network Adapter
 18...aa aa aa 2c a3 00 ......Shrew Soft Virtual Adapter
  1...........................Software Loopback Interface 1
 22...00 00 00 00 00 00 00 e0 Microsoft ISATAP Adapter
 17...00 00 00 00 00 00 00 e0 Microsoft Teredo Tunneling Adapter
 23...00 00 00 00 00 00 00 e0 Microsoft ISATAP Adapter #3
 25...00 00 00 00 00 00 00 e0 Microsoft ISATAP Adapter #4
 14...00 00 00 00 00 00 00 e0 Microsoft ISATAP Adapter #2
 20...00 00 00 00 00 00 00 e0 Microsoft ISATAP Adapter #5

```

接口列表的第一列字段就是网卡接口索引

## 3. netstat

```
[c:\~]$ netstat -nr
===========================================================================
接口列表
  4...9e b6 d0 0f 85 2b ......Microsoft Wi-Fi Direct Virtual Adapter
 26...00 50 56 c0 00 01 ......VMware Virtual Ethernet Adapter for VMnet1
 12...00 50 56 c0 00 08 ......VMware Virtual Ethernet Adapter for VMnet8
  7...00 50 56 c0 00 02 ......VMware Virtual Ethernet Adapter for VMnet2
  3...9c b6 d0 0f 85 2b ......Killer Wireless-n/a/ac 1535 Wireless Network Adapter
 18...aa aa aa 2c a3 00 ......Shrew Soft Virtual Adapter
  1...........................Software Loopback Interface 1
 22...00 00 00 00 00 00 00 e0 Microsoft ISATAP Adapter
 17...00 00 00 00 00 00 00 e0 Microsoft Teredo Tunneling Adapter
 23...00 00 00 00 00 00 00 e0 Microsoft ISATAP Adapter #3
 25...00 00 00 00 00 00 00 e0 Microsoft ISATAP Adapter #4
 14...00 00 00 00 00 00 00 e0 Microsoft ISATAP Adapter #2
 20...00 00 00 00 00 00 00 e0 Microsoft ISATAP Adapter #5

```