# Linux内核编译

参考文章：

[ubuntu12.04 升级内核实战](http://bbs.pcbeta.com/forum.php?mod=viewthread&tid=1040977)

[ubuntu 内核编译](http://www.cnblogs.com/devil-91/archive/2012/07/23/2605568.html)

## 1. 编译环境

- 系统版本：Ubuntu14.04

- 内核版本：3.18.3

（另：在Ubuntu10.04下编译2.6.32.65版本步骤相同）

## 2. 准备工作

确认以下工具已安装：

- build-essential (基本的编程库（`gcc`, `make`等）

- kernel-package(Debian系统里生成`kernel-image`的一些配置文件和工具)

- libncurses5-dev(`make menuconfig`中要调用的)

- libqt3-headers(`make xconfig`要调用的)

下载并解压好的内核源码可以放在任意位置，建议在home目录下。

## 3. 编译开始

> **注：以下操作都是在内核源码根目录下进行**

### 3.1 清除残留文件(如果是第一次编译，可跳过此步)

命令行下输入：

```
sudo make mrproper
```

该命令的功能在于清除当前目录下残留的`.config`和`.o`文件，这些文件一般是以前编译时未清理而残留的。而对于第一次编译的代码来说，不存在这些残留文件，所以可以略过此步，但是如果该源代码以前被编译过，那么强烈建议执行此命令，否则后面可能会出现未知的问题.


### 3.2 配置编译选项

```
sudo make menuconfig
```

根据菜单提示，选择编译配置选项，并保存配置文件为`.config`(也可以直接复制现有的`.config`文件，一般在`/boot`下有当前内核所用的配置文件，以`config-*`开头)

### 3.3 清除编译中间文件(可不执行)

```
sudo make clean
```

### 3.4 生成新内核

```
sudo make bzImage
```

### 3.5 生成模块

```
sudo make modules
```

### 3.6 安装模块

```
sudo make modules_install
``` 

### 3.7 建立ramdisk映像文件

```
sudo mkinitramfs -o /boot/initrd-3.18.3.img 3.18.3
```

其中`initrd-3.18.3.img`的版本号可自定，而后面`3.18.3`则需要与当前正在编译的内核版本一致

### 3.8 安装内核

```
sudo make install
```

此时系统会把linux内核的镜像文件还有`System.map`拷贝到`/boot`下，然后会自动生成引导菜单。

### 3.9 配置grub引导程序

在`Ubuntu14.04`下编译`3.18.3`时grub是自动更新的，然而在`Ubuntu10.04`下编译`2.6.32.65时并没有更新。所以在下一步重启之前最好先检查一下`/boot/grub/grub.cfg`

步骤如下：

编辑`/boot/grub/grub.cfg`，参照已经存在的启动项和`/boot`下的文件名称，添加新的启动项，比如：

```
menuentry 'Ubuntu，Linux 3.18.3' --class ubuntu --class gnu-linux --class gnu --class os {
recordfail
gfxmode $linux_gfx_mode
insmod gzio
insmod part_msdos
insmod ext2
set root='(hd0,msdos1)'
search --no-floppy --fs-uuid --set=root ee7c3a4d-5305-46b1-807e-fa9f39a5d13e
linux /boot/vmlinuz-3.18.3 root=UUID=ee7c3a4d-5305-46b1-807e-fa9f39a5d13e ro quiet splash $vt_handoff
initrd /boot/initrd-3.18.3.img
}
```

## 4. 重新启动，编译完成

## 5. 拓展-关于二次编译

当已经编译过一次内核，暂时不要删除产生的中间文件。这样接下来再次编译时不必重新执行首次编译的全部步骤，时间也会缩短很多。

修改完源代码之后运行：

```
sudo make
sudo make install
```

然后重新启动，重新编译的过程应该会比初次编译快很多。
