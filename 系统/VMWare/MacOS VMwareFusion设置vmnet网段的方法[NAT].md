# MacOS VMwareFusion设置vmnet网段的方法[NAT]

参考文章

1. [vmware fusion设置vmnet网段的方法](https://www.cnblogs.com/iwantcomputer/archive/2013/04/19/8489821.html)

```
sudo vim /Library/Preferences/VMware\ Fusion/networking
```

更新网段

```
VERSION=1,0
answer VNET_1_DHCP yes
answer VNET_1_DHCP_CFG_HASH BA5B761F774CC031718E05D8A7F5446CF7033192
answer VNET_1_HOSTONLY_NETMASK 255.255.255.0
answer VNET_1_HOSTONLY_SUBNET 192.168.224.0
answer VNET_1_VIRTUAL_ADAPTER yes
answer VNET_8_DHCP yes
answer VNET_8_DHCP_CFG_HASH F7E917777F3D55D65F732985872CF8D5F1B08AD4
answer VNET_8_HOSTONLY_NETMASK 255.255.255.0
## 只要更改这一行就可以了
answer VNET_8_HOSTONLY_SUBNET 172.16.91.0
answer VNET_8_NAT yes
answer VNET_8_VIRTUAL_ADAPTER yes
add_bridge_mapping en0 2
```
