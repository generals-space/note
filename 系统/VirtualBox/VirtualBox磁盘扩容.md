# VirtualBox磁盘扩容

参考文章

1. [VirtualBox上Centos7磁盘扩容](https://blog.csdn.net/haeydy/article/details/89447689)

virtualbox 固定大小磁盘可以先复制一份为动态分配, 然后更换磁盘, 再扩容.

在 virtualbox GUI 界面上扩容后, centos 进入系统会发现`df -h`并未更新, 但是`fdisk -l`可以看到, 需要手动执行命令进行 lvm 逻辑卷操作, 见参考文章1.
