# VMWare设置共享目录

参考文章

1. [解决 VMware 虚拟机共享文件 但是找不到共享文件夹](https://blog.csdn.net/Neneolia/article/details/119567372)
    - windows宿主机, linux虚拟机

## 1. Linux访问windows宿主机

VMware软件中开启共享目录后, 无需重启直接就可以在linux虚拟机的`/mnt/hgfs`目录下看到共享目录.

如果没有, 则需要重新安装`vmware-tools`, 完成后也将立即看到结果.

要注意的是, VMware挂载的`vmware-tools`是在`/dev/cdrom`下, 但没有办法直接进入, 需要首先将其挂载到一个存在的目录下.

**注意: 不能挂载到`/mnt`目录下, 否则安装完成后光驱弹出操作会阻止hgfs目录的创建**

```
mount /dev/cdrom /opt/cdrom
```

这个目录是只读的, 将安装包拷贝到本地目录后解除挂载

```
umount /opt/cdrom
```
