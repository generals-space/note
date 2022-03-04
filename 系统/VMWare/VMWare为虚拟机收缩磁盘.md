# VMWare为虚拟机收缩磁盘

参考文章

1. [关于VMware虚拟机磁盘收缩的几种方法](https://www.cnblogs.com/5201351/p/4290401.html)

宿主机: Win10
VMWare: 15.1
虚拟机: CentOS 7

20210417: MacOS宿主机, VMwareFusion, CentOS7虚拟机, 实践可行.

方法一实验有效.

原来虚拟机实际占用11G左右, 但是在宿主机上查看已经使用了近50G.

```
# df -h
Filesystem      Size  Used Avail Use% Mounted on
devtmpfs        894M     0  894M   0% /dev
tmpfs           910M     0  910M   0% /dev/shm
tmpfs           910M  9.9M  900M   2% /run
tmpfs           910M     0  910M   0% /sys/fs/cgroup
/dev/sda3        78G   11G   68G  14% /
/dev/sda1       297M  254M   44M  86% /boot
tmpfs           182M     0  182M   0% /run/user/0
```

![](https://gitee.com/generals-space/gitimg/raw/master/D0C47E7E9C4392D6586130F96949A4A8.png)

按照参考文章1中所说, 使用方法一.

```
vmware-toolbox-cmd disk shrink /
```

![](https://gitee.com/generals-space/gitimg/raw/master/7D42926F3F8599E44EE890F55EF564C3.png)

执行完成后, ssh连接会断开, 然后vmware会显示如下对话框.

![](https://gitee.com/generals-space/gitimg/raw/master/0D3EB817E3DA97F1F62D3CC78EBDFA70.png)

当流程走完后, 再连接(我尝试时还重启了一下, 不然连不上, 虚拟机终端也无响应)

```
# df -h
Filesystem      Size  Used Avail Use% Mounted on
devtmpfs        894M     0  894M   0% /dev
tmpfs           910M     0  910M   0% /dev/shm
tmpfs           910M  9.9M  900M   2% /run
tmpfs           910M     0  910M   0% /sys/fs/cgroup
/dev/sda3        78G   11G   68G  14% /
/dev/sda1       297M  254M   44M  86% /boot
tmpfs           182M     0  182M   0% /run/user/0
```


![](https://gitee.com/generals-space/gitimg/raw/master/C123498CB9A13E31D798E6ADE9250873.png)

此时该虚拟机在宿主机上就只占用13.4G了.

