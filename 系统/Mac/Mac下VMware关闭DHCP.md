# Mac下VMware关闭DHCP

参考文章

1. [Mac VMware fusion 专用网络关闭DHCP](https://blog.csdn.net/miiser/article/details/44419759)

`/Library/Preferences/VMware\ Fusion/networking`文件中存储了vmware提供的虚拟网络配置, 默认包括`vnet1`, `vnet8`两个(前者应该是`host-only`, 后者则是`NAT`).

```log
VERSION=1,0
answer VNET_1_DHCP yes
answer VNET_1_DHCP_CFG_HASH B4B7EF1463ADB73FAF55D8E47623A74DC52336B3
answer VNET_1_HOSTONLY_NETMASK 255.255.255.0
answer VNET_1_HOSTONLY_SUBNET 192.168.224.0
answer VNET_1_VIRTUAL_ADAPTER yes
answer VNET_8_DHCP yes
answer VNET_8_DHCP_CFG_HASH EB8A7BF799ABB2FD9AF330952C619B1EFBD5B383
answer VNET_8_HOSTONLY_NETMASK 255.255.255.0
answer VNET_8_HOSTONLY_SUBNET 172.16.91.0
answer VNET_8_NAT yes
answer VNET_8_VIRTUAL_ADAPTER yes
```

修改`VNET_8_DHCP`为`no`, 然后重启vmware即可.
