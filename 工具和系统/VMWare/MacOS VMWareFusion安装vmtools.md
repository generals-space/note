# MacOS VMWareFusion安装vmtools

参考文章

1. [macos 下 vmware fusion 安装 vmware tools](https://blog.csdn.net/fangkailove/article/details/106164372)
2. [Vmware Fusion中CentOS 7中安装Vmware Tools](https://www.jibing57.com/2019/04/26/manually-install-vmeare-tools-on-centos-7/)
    - `mount /dev/cdrom /mnt/cdrom`
3. [vmware + centos 7安装vmtools时提示The path "" is not a valid path to the xxx kernel header](https://blog.csdn.net/liu1340308350/article/details/80824696)
    - `yum install kernel-headers kernel-devel gcc`
    - `/lib/modules/3.10.0-1160.24.1.el7.x86_64/build/include`
4. [#./vmware-install.pl踩点：](https://www.erlo.vip/share/2/15676.html)
    - `yum install kernel*`
5. [解决vmware fusion + centos 7安装vmtools时提示The path "" is not a valid path to the xxx kernel headers.](https://blog.csdn.net/sirchenhua/article/details/49719659)
    - 安装完内核头, 需要重启生效.

从 CDROM 加载 vmtool 工具

![](https://gitee.com/generals-space/gitimg/raw/master/793695b1c5daa6e33b039d7ccf9a8509.png)

![](https://gitee.com/generals-space/gitimg/raw/master/4ce0c02e6c9659a64bfe271feefb50ed.png)

然后就入命令行, 挂载

```console
$ ls /mnt/cdrom

$ mount | grep iso9660
$ mount /dev/cdrom /mnt/cdrom
mount: /dev/sr0 写保护，将以只读方式挂载
$ ls /mnt/cdrom
manifest.txt  run_upgrader.sh  VMwareTools-10.3.10-13959562.tar.gz  vmware-tools-upgrader-32  vmware-tools-upgrader-64
```

拷贝出来, 解压执行

```conole
$ ls
VMwareTools-10.3.10-13959562.tar.gz
$ tar -zxf ./VMwareTools-10.3.10-13959562.tar.gz
$ ls
VMwareTools-10.3.10-13959562.tar.gz  vmware-tools-distrib
$ cd vmware-tools-distrib/
$ ls
bin  caf  doc  etc  FILES  INSTALL  installer  lib  vgauth  vmware-install.pl
$ ./vmware-install.pl
```

如果是CentOS最小化安装, 则系统没有perl, gcc和kernel-devel, 需要先安装

```
yum install -y perl gcc make kernel-headers kernel-devel
```

否则会陷入死循环.

```
The path "" is not a valid path to the 3.10.0-1062.el7.x86_64 kernel headers.
Would you like to change it? [yes]

INPUT: [yes]  default

Enter the path to the kernel header files for the 3.10.0-1062.el7.x86_64
kernel?

```

...但结果还是不行.

...重启了, 还是不行.

按照参考文章3, 找到了`/lib/modules/3.10.0-1162.24.1.el7.x86_64`目录, 发现ta下面的`build`子文件是个软链接, 但是链接目标目录是`/usr/src/kernels/3.10.0-1162.24.1.el7.x86_64`不存在了. 但是在`/usr/src/kernels/`目录下发现了其他版本的子目录, 于是我尝试重建了一下这个软链接.

```
ln -s /usr/src/kernels/3.10.0-1160.24.1.el7.x86_64 /lib/modules/3.10.0-1062.el7.x86_64/build
```

没错, 还是不行.

后来我找到了参考文章4, 直接`yum install kernel*`, 再重启, 终于可以了...

