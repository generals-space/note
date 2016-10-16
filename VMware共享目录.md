# VMware共享目录

## 1. Linux

VMware软件中开启共享目录后, 无需重启直接就可以在linux虚拟机的`/mnt/hgfs`目录下看到共享目录.

如果没有, 则需要重新安装`vmware-tools`, 完成后也将立即看到结果.

要注意的是, VMware挂载的`vmware-tools`是在`/dev/cdrom`下, 但没有办法直接进入, 需要首先将其挂载到一个存在的目录下.

```
mount /dev/cdrom /mnt
```

这个目录是只读的, 将安装包拷贝到本地目录后解除挂载

```
umount /mnt
```