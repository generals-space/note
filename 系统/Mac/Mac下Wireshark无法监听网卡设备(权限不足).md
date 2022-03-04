# Mac下Wireshark无法监听网卡设备(权限不足)

参考文章

1. [Wireshark 抓包遇到 you don’t have permission to capture on that device mac 错误的解决方案](https://www.cnblogs.com/Mr-zyh/p/7684727.html)
2. [Wireshark for mac 权限](https://www.jianshu.com/p/73652a2375e5)

![](https://gitee.com/generals-space/gitimg/raw/master/624afb2e1f7a8995b854e9357e6c3895.jpg)

> `en0`是我电脑上的wifi网卡.

```console
$ ll /dev/bpf*
crw-------  1 root  wheel   23,   0  1 28 22:09 /dev/bpf0
crw-------  1 root  wheel   23,   1  1 28 22:09 /dev/bpf1
crw-------  1 root  wheel   23,   2  1 30 00:23 /dev/bpf2
crw-------  1 root  wheel   23,   3  1 30 00:23 /dev/bpf3
crw-------  1 root  wheel   23,   4  1 28 22:40 /dev/bpf4
```

解决方法有两种, 要么把`bpf*`的权限修改为`755`, 要么把ta们的属主改为当前的登录用户, 如`general`.

然后重启wireshark.
