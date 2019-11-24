# Window安装OpenSSH

参考文章

1. [Windows安装OpenSSH服务](https://www.jianshu.com/p/6e5bc39d386e)
2. [官方文档 Install Win32 OpenSSH](https://github.com/PowerShell/Win32-OpenSSH/wiki/Install-Win32-OpenSSH)

win10教育版的可选功能无法打开(不是"没有可安装的功能", 而是打开就闪退, 从控制面板进的话刷不出可安装的功能), 只能手动安装. 正好powershell官方提供了openssh的安装方法, 按照教程走就行了, 还挺简单的.

```ps1
## 安装sshd服务
powershell.exe -ExecutionPolicy Bypass -File install-sshd.ps1
## 启动sshd服务
net start sshd
## 设置开机启动
Set-Service sshd -StartupType Automatic
```

```ps1
## 卸载
powershell.exe -ExecutionPolicy Bypass -File uninstall-sshd.ps1
```
