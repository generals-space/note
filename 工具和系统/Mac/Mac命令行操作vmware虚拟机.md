# Mac命令行操作vmware虚拟机

参考文章

1. [VMware命令行打开虚拟机（加密密码）及相关文档](https://blog.csdn.net/weixin_40277264/article/details/107712827)
2. [vmrun命令行的使用（VMWare虚拟机）](https://blog.csdn.net/weixin_40277264/article/details/107712827)

启动虚拟机

```
vmrun start ./Virtual\ Machines.localized/CentOS\ 7\ 64\ 位.vmwarevm
```

查看正在运行的虚拟机

```console
$ vmrun list
Total running VMs: 2
/Users/general/Virtual Machines.localized/CentOS 7 64 位.vmwarevm/CentOS 7 64 位.vmx
/Users/general/Virtual Machines.localized/kw01/kw01.vmx
```

# 冷重启虚拟机 | 热重启虚拟机

```
vmrun reset "/Users/general/Virtual Machines.localized/kw01/kw01.vmx" hard | soft
```

> 建议使用`soft`.

