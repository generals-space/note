# VMWare问题

>VMWare虚拟机总是, 突然之间就无法启动了...

>各种奇葩问题总结...

## 1. 关键词 `Could not open /dev/vmmon`

[参考文章1](http://blog.csdn.net/gsying1474/article/details/40684071)

[参考文章2](https://communities.vmware.com/message/2442783)

根据解决方法分析...这个问题应该是因为找不到系统模块`vmmon`(`modprobe`那个层级). 解决步骤如下.

```shell
cp /usr/lib/vmware/modules/source/vmmon.tar /tmp
cd /tmp
tar -xf ./vmmon.tar
cd vmmon-only
make
cp ./vmmon.ko /lib/modules/$(uname -r)/misc/vmmon.ko
modprobe vmmon
```

我自己的情况是, `/lib/modules/$(uname -r)/misc/vmmon.ko`文件已经存在, 而且`modprobe -l`能查看到有`misc/vmmon.ko`, 只是`lsmod`里没有`vmmon.ko`. 

这说明该模块可以被系统检测到, 只是不知道为什么被移除了. 只要再加进去就可以了吧. 但是尝试了以下3种方法, 都失败, 报相似的错误, 不知道是不是人品问题.

```shell
modprobe vmmon.ko
modprobe misc/vmmon.ko
insmod /lib/modules/$(uname -r)/misc/vmmon.ko

FATAL: Module vmmon not found error
```

于是按照上面参考文章中的方法, 覆盖掉原来的`vmmon.ko`文件(还是备份一下比较好), 但是`modprobe`还是不起作用, 换用`insmod`命令加文件绝对路径, 可以了.

## 2. 网络配置

使用`菜单->编辑->虚拟网络编辑器`对VMnet(x)进行配置后最好重启一个VMware软件而不是单独重启虚拟机, 否则会出现网络配置不生效, 生效了也无法正常进行网络连接的情况(尤其是IP, 路由, 防火墙都正常的时候, 就是无法访问外部网络).

## 3.

问题描述:

启动虚拟机时报如下错误:

```
无法打开内核设备"\\.\Global\vmx86": 系统找不到指定的文件. 是否在安装 VMware Workstation 后重新引导?

未能初始化监视器设备. 
```

原因分析:

VMware相关服务未启动, 找找看全部打开就行了.

## 4. XAMPP和VMware Workstation占用443端口冲突

[XAMPP和VMware Workstation占用443端口冲突的解决办法](http://www.weste.net/2014/10-28/99655.html)

今天安装了一个`VMware Workstation`，发现`XAMPP`的`Apache`就启动不了. 看了一下错误日志，似乎是VMware Workstation占用了443端口导致冲突引起的. 查看了一下，原来`VMware Workstation`有个共享虚拟机的服务，占用了443端口. 

对于单机安装虚拟机来说，这个功能没有用处，禁用掉就可以了. 操作步骤如下：

1. 打开VMware Workstation，点击菜单中的"编辑->首选项";

2. 找到左侧功能列表中的"共享虚拟机"，选择后，在右侧界面中点击"更改设置";

3. 这个时候，本来是disabled的"禁用共享"按钮就被激活了，点击"禁用共享"按钮，就可以将这个功能禁用了;

4. 如果还想使用此功能，可以将443端口修改成446或者其他端口都可以. 而且不需要关闭正在运行的虚拟机. 