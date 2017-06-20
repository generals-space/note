# VNC server搭建

参考文章

[怎样在 CentOS 7.0 上安装和配置 VNC 服务器](http://www.linuxidc.com/Linux/2015-04/116725.htm)

## RealVNC

```
yum groupinstall "GNOME Desktop"
yum install -y  tigervnc tigervnc-server

```

```
[root@localhost .vnc]# pwd
/root/.vnc
[root@localhost .vnc]# ll
total 24
-rw-r--r--  1 root root 1397 Aug  5  2016 localhost.localdomain:1.log
-rw-r--r--  1 root root    5 Aug  5  2016 localhost.localdomain:1.pid
-rw-r--r--  1 root root 1397 Jul 27  2015 localhost.localdomain:2.log
-rw-r--r--  1 root root    5 Jul 27  2015 localhost.localdomain:2.pid
-rw-------. 1 root root    8 Mar 23  2015 passwd
-rwxr-xr-x. 1 root root   92 Oct 21  2014 xstartup
```



```
$ cp /lib/systemd/system/vncserver@.service /etc/systemd/system/
```

把其中的`<USER>`字段修改成可以通过VNC登录的普通用户名(当然, root也是可以的)

```
$ systemctl daemon-reload
```

```
[root@localhost .vnc]# vncpasswd
Password:
Verify:
```


systemctl enable vncserver@:1.service
systemctl start vncserver@:1.service

```
$ firewall-cmd --permanent --add-service vnc-server
success
$ systemctl restart firewalld.service
$ 
```

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
# gnome-session &
#startkde &
twm &
# xfce4-session &

```