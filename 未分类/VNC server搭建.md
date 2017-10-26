# VNC server搭建

参考文章

1. [怎样在 CentOS 7.0 上安装和配置 VNC 服务器](http://www.linuxidc.com/Linux/2015-04/116725.htm)

2. [CentOS 中 YUM 安装桌面环境](http://cnzhx.net/blog/centos-yum-install-desktop/)

3. [（总结）CentOS Linux下VNC Server远程桌面配置详解](http://www.ha97.com/4634.html)

## 图形桌面

首先安装桌面环境, 这个话题参考文章2中讲解的十分清楚, 值得一看.

```
## yum groupinstall "X Window System"
yum groupinstall "GNOME Desktop"
yum install -y  tigervnc-server
```

一般服务器都是通过字符界面启动的, 这种情况下是没有办法启动vnc的, 也没有办法启动诸如`xmanager`这种图形桌面. 我们需要将系统设置为从图形界面启动.

参考文章2中讲述了如何修改启动类型.

```
centos 6: sed -i 's/id:3:initdefault:/id:5:initdefault:/' /etc/inittab
centos 7: systemctl set-default graphical.target
```

### 基本认识

然后启动vnc server.

```
$ vncserver :1
New '服务器IP:1 (启动用户)' desktop is 服务器IP:1
```

首次执行这个操作, 会在用户家目录下创建`.vnc`目录, 其中vnc server实例的日志和pid都在这里, 还有配置文件什么的.

注意, 这个`:1`没有什么特殊的地方, 你可以把它当作一个vnc server实例的别名, 一般来说, `:n`就表示其端口为`590n`. 在执行`vncserver`时不显式指定`:n`的话它会自动创建一个合适的.

vncserver可以由任意用户启动, 一个用户也可以启动多个vnc server, 配置文件默认就在其用户home目录下的`.vnc`下.

连接上哪一个vnc实例, 就会看到哪个用户的登录界面, 登录后获得该用户的会话及相应权限.

#### 查看/结束进程

`-list`选项可以查看当前用户启动的vnc实例

```
$ vncserver -list

TigerVNC server sessions:

X DISPLAY #	PROCESS ID
:1		23440
:2		27121
```

`-kill`可以停止目标实例, 如

```
$ vncserver -kill :1
Killing Xvnc process ID 23440

```

### 关于密码

通过vnc连接时, 一般只会让用户指定目标IP:端口, 可能会弹出密码输入框, **此时的密码是建立vnc连接的密码, 并不是用户本身的系统密码.**

这个密码可以不设置(...应用), 要设置的话, 方法为, 在vnc server的启动用户下执行如下命令.

```
[root@localhost .vnc]# vncpasswd
Password:
Verify:
```

这个操作将在`~/.vnc/`目录下创建一个`passwd`文件, 所有这个用户启动的vnc server, 默认都会经过这里配置的密码验证.

OK, 在建立连接后就能看到系统登录界面了, 这需要该系统用户设置过密码, 通过`passwd`设置.

其实通过这种方法启动服务已经很不错了, 哦, 再加上开机启动.

------

## systemctl启动

这个方法我没搞成功, 通过systemctl启动的vnc server, 连接上后背景是黑的, 还没有标题栏...如下

![](https://gitimg.generals.space/fc7dba6a462af812957e9523f1b6f36b.png)

首先拷贝服务脚本, 注意这只是模板, 内容需要自己修改

```
$ cp /lib/systemd/system/vncserver@.service /usr/lib/systemd/system/vncserver@:1.service
```

把其中的`<USER>`字段修改成可以通过VNC登录的普通用户名, 该用户需要事先存在于系统中(当然, root也是可以的, 要注意pid文件的路径)

注意, 服务脚本文件中的`@:1`是必须的. 因为脚本内容中有如下代码

```
ExecStart=/usr/bin/vncserver %i
PIDFile=/root/.vnc/%H%i.pid
ExecStop=-/usr/bin/vncserver -kill %i
```

其中`%H`与`%i`是systemd的内置变量, 前者表示本机`hostname`值, 后者表示服务实例的`@`与`.service`后缀之前的值, 正好是`:1`.

然后启动服务.

```
$ systemctl daemon-reload
systemctl enable vncserver@:1.service
systemctl start vncserver@:1.service
```

注意防火墙开启VNC的端口

```
$ firewall-cmd --permanent --add-service vnc-server
success
$ systemctl restart firewalld.service
$ 
```

------

关于`xstartup`文件, `vncserver`命令可以通过`-xstartup 文件路径`指定这个文件的路径, 我这边实验的时候没看出加不加这些东西有什么区别. 

这里保留一下, 可以在VNC显示鼠标为黑色叉号时修改一下作为参考.

```
[root@localhost .vnc]# cat xstartup 
#!/bin/sh

# Uncomment the following two lines for normal desktop:
unset SESSION_MANAGER
exec /etc/X11/xinit/xinitrc

[ -x /etc/vnc/xstartup ] && exec /etc/vnc/xstartup
[ -r $HOME/.Xresources ] && xrdb $HOME/.Xresources
xsetroot -solid grey
vncconfig -iconic &
xterm -geometry 80x24+10+10 -ls -title "$VNCDESKTOP Desktop" &
gnome-session &
#startkde &
```

这里的`gnome-session &`我觉得的是在装多个桌面环境时要加的

### 分辨率/DPI设置

参考文章3中有详细解释.

最简单的一种

```
$ vncserver :1 -geometry 1920x1080 -depth 24
```

这种需要在实例启动时指定, 也可以写在`/etc/sysconfig/vncserver`中, 后者没有实验过.