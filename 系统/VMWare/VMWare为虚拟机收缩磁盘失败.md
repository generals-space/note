# VMWare为虚拟机收缩磁盘失败

参考文章

1. [关于VMware虚拟机磁盘收缩的几种方法](https://www.cnblogs.com/5201351/p/4290401.html)
2. [vmware fusion清理linux虚拟机磁盘](https://zhuanlan.zhihu.com/p/275907188)
     - `vmware-toolbox-cmd disk shrink /`宿主机可能会崩溃
     - `vmware-toolbox-cmd disk shrinkonly`换成这个
     - 收缩磁盘包括两部分, 整理和收缩, 前者包含了整理的过程, 而后者只是收缩

宿主机: Win10
VMWare: 15.1
虚拟机: CentOS 7

按照方法一操作时竟然失败了, 出现如下报错.

```
$ vmware-toolbox-cmd disk shrink /
Shrink disk is disabled for this virtual machine.

Shrinking is disabled for linked clones, parents of linked clones, 
pre-allocated disks, snapshots, or due to other factors. 
See the User's manual for more information.
Unable to find partition /
```

重启虚拟机, 宿主机, 都不能解决. 我也尝试过方法三...简直被害惨, 虚拟机磁盘被扩充到100%, 所在分区被占满, 也尝试过vmware软件的磁盘工具, 碎片整理和收缩都试过了, 还是没办法减下来.

后来尝试了方法四.

```
F:\Work-CentOS7>d:\vmware\vmware-vdiskmanager -k Ori-CentOS7-cl2.vmdk
  Shrink: 100% done.
Shrink completed successfully.

F:\Work-CentOS7>d:\vmware\vmware-vdiskmanager -k Ori-CentOS7-cl2-000001.vmdk
  Shrink: 100% done.
Shrink completed successfully.

F:\Work-CentOS7>d:\vmware\vmware-vdiskmanager -k Ori-CentOS7-cl2-000002.vmdk
  Shrink: 100% done.
Shrink completed successfully.
```

成了.

------

不过其实在上面的成功之前, 还有一次报错.

```
F:\Work-CentOS7>d:\vmware\vmware-vdiskmanager -k Ori-CentOS7-cl2-000001.vmdk
Failed to shrink the disk 'Ori-CentOS7-cl2-000001.vmdk' : An error occurred while writing a file; the disk is full. Data has not been saved. Free some disk space and try again (0x8).
```

这是因为我前面把虚拟机所在分区占满了, 把其他的东西移走, 给shrink操作腾出一点空间即可.
