# VirtualBox5.x没有vdfuse工具的解决方法

参考文章

1. [centos/ubuntu挂载vmdk、 vdi为块设备的方法(非vdfuse)](http://blog.51cto.com/zhangyu/1862716)

`vdfuse`命令是VirtualBox软件挂载VDI分区文件的一个工具，VirtualBox是一款能创建虚拟机的开源软件，vdi是它的默认磁盘格式。

但是这个工具只在Virtualbox 4.x中出现, 5.0+之后就被取消了. 参考文章1中有提到过替代方法, 但不是很想用. 

后来想到先安装一个4.x的virtualbox, 把其中的`vdfuse`可执行文件拷贝到安装了5.x的机器上. 当然, 单纯只拷贝这个文件是不行的, 它还依赖一些其他的共享库, 所以需要执行看看缺少哪些库文件, 都拷贝过来就可以了.