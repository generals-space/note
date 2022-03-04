# MacOS通过命令行启动 RealVNC server

参考文章

1. [Mac OS系统上用命令行方式启动VNC Server](https://blog.csdn.net/jollypigclub/article/details/48582203)

Mac OS X 上打开VNC Server服务(不带vnc密码):

```
sudo /System/Library/CoreServices/RemoteManagement/ARDAgent.app/Contents/Resources/kickstart -activate -configure -access -off -restart -agent -privs -all -allowAccessFor -allUsers
```

Mac OS X 上打开VNC Server服务(带vnc密码, 替换myVncPassword为自己的密码): 

```
sudo /System/Library/CoreServices/RemoteManagement/ARDAgent.app/Contents/Resources/kickstart -activate -configure -access -off -restart -agent -privs -all -allowAccessFor -allUsers -clientopts -setvncpw -vncpw myVncPassword
```

> 这里说的带不带 vnc 密码, 应该是独立于系统密码之外的 vnc 的密码.

