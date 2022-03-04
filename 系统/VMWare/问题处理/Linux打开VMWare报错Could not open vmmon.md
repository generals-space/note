# Linux打开VMWare报错Could not open vmmon

1. [参考文章1](http://blog.csdn.net/gsying1474/article/details/40684071)

2. [参考文章2](https://communities.vmware.com/message/2442783)

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
