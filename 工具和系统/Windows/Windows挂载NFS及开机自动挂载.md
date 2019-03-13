# Windows挂载NFS

参考文章

1. [在windows下挂载nfs文件系统](https://blog.csdn.net/wykkunkun/article/details/79638288)

2. [Windows挂载NFS文件系统](https://www.cgtblog.com/wljs/1537.html)

首先是在windows主机上启用nfs客户端功能(普通win10桌面版不需要重启即可使用, 而在windows server 2012上需要重启才行).

挂载方式为, 在cmd中输入

```
mount \\192.168.1.100\mnt\medianfs Z:
```

注意IP前的两个反斜线, 之后共享路径的分隔符也是反斜线, 最后一个参数为在本机上的盘符, 必须要有冒号`:`. 这样就可以把`192.168.1.100`上共享的`/mnt/medianfs`目录挂载到本地的`Z`盘中了.

挂载的方法参考文章1中讲解的很清楚, 可以看一下.

------

关于开机自动挂载, 点击"计算机" -> 点击"映射网络驱动器" -> "输入网络共享文件路径" -> "完成".

其中驱动器目录可以自主选择, 另外要挂载的文件夹路径和mount命令中相同(`\\192.168.1.100\mnt\medianfs`), **登录时重新连接**是默认勾选的.