# MacOS下VMware Fusion为CentOS7扩容磁盘

<!tags!>: <!lvm!> <!LVM!>

参考文章

1. [Mac VMware Fusion 中修改 centos7 虚拟机的磁盘空间、扩容](https://blog.csdn.net/ifjgm003/article/details/101461585)

参考文章1介绍了2种方式进行扩容, 一种是新增磁盘, 然后新增挂载点, 一种是扩容原磁盘, 但需要借助LVM.

本来想用第1种种的, 但是在`fusion`中调整了一下磁盘容量(从20G调到了50G), 结果调不回去了, 还是用第2种吧...

------

几乎是完全照着参考文章1中的步骤来做的, 需要注意的有几点.

一是有一句"可以看到/dev/sda4的Id号为83，我们要将其改成8e(LVM卷文件系统的Id)"

以我在实验中的输出为例(`sda3`是扩容后用`fdisk /dev/sda`新增的分区).

```console
$ fdisk -l

磁盘 /dev/sda：53.7 GB, 53687091200 字节，104857600 个扇区
Units = 扇区 of 1 * 512 = 512 bytes
扇区大小(逻辑/物理)：512 字节 / 512 字节
I/O 大小(最小/最佳)：512 字节 / 512 字节
磁盘标签类型：dos
磁盘标识符：0x000a92da

   设备 Boot      Start         End      Blocks   Id  System
/dev/sda1   *        2048     2099199     1048576   83  Linux
/dev/sda2         2099200    41943039    19921920   8e  Linux LVM
/dev/sda3        41943040   104857599    31457280   83  Linux

...省略
```

观察上面输出的表格中的`Id`列, 这里的`Id`其实是文件系统类型的标识符, `sda2`的类型是`LVM`, 其`Id`为`8e`, 普通的就是`83`.

因为要借助`LVM`, 所以还需要把`sda3`的类型改成`LVM`, 继续按照操作来就行了, 再次输出就变成了`LVM`类型.

```
   设备 Boot      Start         End      Blocks   Id  System
/dev/sda1   *        2048     2099199     1048576   83  Linux
/dev/sda2         2099200    41943039    19921920   8e  Linux LVM
/dev/sda3        41943040   104857599    31457280   8e  Linux LVM
```

实验成功, 由于有很多地方需要配置结果验证, 这里就不给出脚本了, 下次有需要重新走一遍流程即可.

这种方式不需要写`fstab`, 重启后仍然有效.
