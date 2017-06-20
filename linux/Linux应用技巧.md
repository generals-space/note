# Linux应用技巧

## 1. dnf或是yum安装指定版本的包

一般版本指定应该是安装出错时的提示, 自行安装时不会需要某指定版本的. 比如fedora 22安装某软件出错提示如下:

```
Error: package mariadb-devel-1:10.0.17-1.fc22.x86_64 requires mariadb-libs(x86-64) = 1:10.0.17-1.fc22, but none of the providers can be installed.
```

fedora默认的软件版本都比较新, rpm查询时得到

```shell
$ rpm -qa 'mariadb-libs'
mariadb-libs-10.0.21-1.fc22.x86_64
```

根据给出的包与其版本的格式(`包名-版本号-系统平台-位数`)尝试安装一下指定版本

```shell
$ dnf install mariadb-libs-1:10.0.17-1.fc22.x86_64
Last metadata expiration check performed 16:24:36 ago on Mon Feb 22 17:01:22 2016.
Error: cannot install both mariadb-libs-1:10.0.17-1.fc22.x86_64 and mariadb-libs-1:10.0.21-1.fc22.x86_64
(try to add '--allowerasing' to command line to replace conflicting packages)
```

这个情况的出现应该就是当其已经安装的与将要安装的包冲突了, 根据提示, 使用`--allowerasing`选项可强制"降级安装".

```
$ dnf install mariadb-1:10.0.17-1.fc22.x86_64 --allowerasing
Last metadata expiration check performed 16:34:02 ago on Mon Feb 22 17:01:22 2016.
Dependencies resolved.
```

当然, 如果提示有包依赖依然需要一层一层解决的.

## 2. 创建用户时指定已经存在的目录作为其home目录

创建用户时, 使用`-d`选项为其指定一个已经存在的目录作为home目录而不是自动创建. 但是要知道, 用户home目录的属主是其本身而且权限一般是700. 所以使用`root`执行这样的操作时, 该用户不存在并且就算创建了, 也不见得该用户拥有此目录的读写权限, 因而可能会报如下错误

```
[root@localhost ~]# mkdir /home/general
[root@localhost ~]# useradd -d /home/general general
adduser: warning: the home directory already exists.
Not copying any file from skel directory into it.
```

但是此时`general`已经被创建了, 并且在`/etc/passwd`文件中其home目录的确也已经指定到`/home/general`中, 但是使用`su`且换为`general`用户时, 命令提示符会呈现一个很原始的状态

```
[root@localhost ~]# cat /etc/passwd | grep general
general:x:1001:1001::/home/general:/bin/bash
[root@localhost ~]# su - general
-bash-4.3$
```

这是因为, 创建新用户的操作, 除了在`/etc/passwd`中添加一行外, 还要从`/etc/skel`目录下拷贝`.bashrc`等模板文件到用户home目录. 指定已经存在的目录作为新建用户的home目录时, 没有拷贝这些模板文件.

所以我们要做的是, 将此home目录的属主该为目标用户, 并修改其权限为`700`, 然后拷贝`/etc/skel`下的模板文件到home目录.

```
[root@localhost ~]# cp /etc/skel/* /home/general
[root@localhost ~]# chown -R general:general /home/general
[root@localhost ~]# chmod 700 /home/general
[root@localhost ~]# su - general
[general@localhost ~]$
```

## 3 强制踢出其他正在SSH登陆的用户

### 3.1 踢人前提

首先, linux系统下root用户可强制踢出其它登录用户, 如果同时有两个root在终端登录, 其中任何一个都可以踢掉另一个;

然后, 任何普通用户都可以踢掉自己, 如果同一个普通帐户在多个终端登录时, 可以互相踢. 但是, 被踢的用户不能用`su`转成其他用户, 否则需要root权限才能踢除. 比如两个用户使用同一个普通帐户`A`登录不同终端, 其中第1个人使用`su`转成了帐户B, 第2个人将无法将其踢除, 除非他有root权限.

最后, 如果要踢掉使用`tty`方式登录的用户(包括普通用户)必须要root权限(ubuntu下是如此, 还没有在其他地方测试过), tty方式登录说明该用户有能力接近服务器本机而不是通过远程ssh登录, 可能这样权限的确大些.

### 3.2 实施方法

1 . 首先可用`w`命令查看登录用户信息, `who`命令也可以, 不过信息不如`w`的输出丰富.

```
general@ubuntu:~$ w
 05:54:18 up 13 min,  5 users,  load average: 0.00, 0.17, 0.18
USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
general  tty1                      05:50    4:10   0.10s  0.09s -bash
general  :0       :0               05:46   ?xdm?  29.35s  0.17s init --user
general  pts/7    :0               05:46    6:50   0.09s  0.09s bash
general  pts/1    192.168.138.1    05:48    6:10   0.07s  0.07s -bash
general  pts/24   192.168.138.1    05:54    2.00s  0.10s  0.00s w
```

其中, `tty1`是使用`Ctrl+Alt+F1`登录的, `pts/7, 1, 24`是通过ssh或ubuntu提供的伪终端登录的.

2 . 可以用`who am i`命令查看此时自己属于哪个终端(小心别踢错了).

```
general@ubuntu:~$ who am i
general  pts/24       2016-02-28 05:54 (192.168.138.1)
```

3 . 踢人的命令格式为`pkill -kill -t 终端名`或是`pkill -9 -t 终端名`

```
general@ubuntu:~$ pkill -kill -t pts/1
general@ubuntu:~$ w
 06:26:29 up 46 min,  4 users,  load average: 0.00, 0.04, 0.13
USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
general  tty1                      05:50   36:21   0.10s  0.09s -bash
general  :0       :0               05:46   ?xdm?  41.26s  0.19s init --user
general  pts/7    :0               05:46   39:01   0.09s  0.09s bash
general  pts/24   192.168.138.1    05:54    5.00s  0.12s  0.01s w
```

## 4. 安装内核头

```
$ yum install linux-headers-$(uname -r)  
```

## 6. hostname

执行以下命令, 当前hostname不会变化, 但是新的会话终端会有变化

```
$ hostname 新的hostname
```