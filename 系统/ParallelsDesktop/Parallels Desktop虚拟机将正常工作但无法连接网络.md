参考文章

1. [Parallels Desktop虚拟机将正常工作但无法连接网络](https://www.gzycdzsw.com/blog/36)
2. [Parallels Desktop虚拟机将正常工作但无法连接网络](https://zhuanlan.zhihu.com/p/459475822)
3. [Mac｜Parallels Desktop 17 无法连接网络及执行该操作失败的解决方案](https://zhuanlan.zhihu.com/p/509319903)

MacOS: Big Sur
Parallels Desktop: 16.0.0 (48916)

参考文章1, 2, 3说的方法其实都一样, 需要改2个文件`dispatcher.desktop.xml`和`network.desktop.xml`.

在我实际操作时, `network.desktop.xml`中并没有`<UseKextless>1</UseKextless>`这一行, 需要手动添加一下, 并改为0.

