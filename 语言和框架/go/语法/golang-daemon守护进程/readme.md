参考文章

1. [golang 创建守护进程](https://blog.csdn.net/ghost911_slb/article/details/7831540)
    - 3种方法创建独立进程: `exec.Command()`, `os.StartProcess()`, `syscall.RawSyscall()`, 不过第3种只说了思路.
2. [Golang 中创建守护进程的正确姿势](https://blog.csdn.net/Qiangks/article/details/86158129)

OS: CentOS 7
kernel: 3.10.0-1062.4.1.el7.x86_64
golang: 1.12.14

本实验中希望启动的后台进程是kubernetes的cni插件: `dhcp`(作用就是dhcp客户端), 使用`yum install kubernetes-cni`安装.

目标可执行程序在`/opt/cni/bin/dhcp`, 启动命令为`./dhcp daemon`.

如果执行时已经存在`/run/cni/dhcp.sock`文件, 会启动失败.

```
./dhcp daemon
2020/02/11 19:59:02 Error getting listener: listen unix /run/cni/dhcp.sock: bind: address already in use
```

命令行执行的话, 标准输入输出如下

```
dhcp    6287 root    0u      CHR              136,3      0t0       6 /dev/pts/3
dhcp    6287 root    1u      CHR              136,3      0t0       6 /dev/pts/3
dhcp    6287 root    2u      CHR              136,3      0t0       6 /dev/pts/3
```

虽然参考文章1中有评论说前两种方法需要添加额外的flag, 但我实验中发现并不需要, 参考文章2中也提出了这个疑问(`syscall.SysProcAttr.Setsid`字段).

