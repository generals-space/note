# VMware安装android安卓模拟器

参考文章

1. [vmware安装安卓](https://zhuanlan.zhihu.com/p/87633763)
2. [VMware虚拟机安装教程](http://bbs.phoenixstudio.org/cn/read.php?tid=7813)
    - 凤凰OS官方安装文档
3. [VMware 12安装凤凰os](https://jingyan.baidu.com/article/20095761dce379cb0721b41b.html)
    = 我最初是按照这篇文档安装凤凰OS的.

只这一篇就可以了.

MacOS: Mojave 10.14.5 (18F132)
VMware Fusion: 15.1

唯一的问题是启动时进入虚拟机时进入引导到总是直接进入文本模式, 需要输入reboot命令, 在下一次引导界面输入两次`e`, 然后再输入`nomodeset`, 回国, 然后按`b`键, 即可进入图形界面.

另外一点就是模拟器的网络问题, wlan中显示连接了`virtnet`, 但总是说已连接入网络但无法上网, 而且顶栏总wlan图标显示红叉❌. 但实际上是可以上的, 想办法跳过连接网络的步骤, 回到桌面, 打开chrome是可以访问百度的.

------

但基本就是个废物, 根本没办法安装国内的软件.

凤凰OS不错, 可以安装闲鱼, 拼多多等, 安装步骤与Android x86的步骤差不多.
