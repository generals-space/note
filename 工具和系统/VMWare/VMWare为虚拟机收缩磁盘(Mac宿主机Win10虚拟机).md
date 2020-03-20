# VMWare为虚拟机收缩磁盘(Mac宿主机Win10虚拟机)

参考文章

1. [关于VMware虚拟机磁盘收缩的几种方法](https://www.cnblogs.com/5201351/p/4290401.html)

宿主机: MacOS 10.14.5(Mojave)
VMWare Fusion: 专业版 11.5.0 (14634996)
虚拟机: Win10 1909

方法一实验同样有效.

原来虚拟机实际占用20G左右, 但是在宿主机上查看已经使用了近45G.

![](https://gitee.com/generals-space/gitimg/raw/master/D6B05AD6794CFDF88E83714E2D8A9A58.jpg)

按照参考文章1中所说, 使用方法一.

![](https://gitee.com/generals-space/gitimg/raw/master/13E818C5FB075C196F16B56A2004F273.jpg)

Windows的vmware tools压缩工具被安装在`C:\Program Files\VMware Tools`目录下, 可执行文件名为"VMwareToolboxCmd.exe".

当流程走完后, 再查看.

![](https://gitee.com/generals-space/gitimg/raw/master/D6B05AD6794CFDF88E83714E2D8A9A58.jpg)

此时该虚拟机在宿主机上就只占用16G了.

