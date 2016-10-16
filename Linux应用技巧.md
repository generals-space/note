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

## 5. 为普通用户添加sudo权限

参考文章

[linux中sudo的用法和sudoers配置详解](http://www.bianceng.cn/OS/Linux/201410/45603.htm)

[linux下sudoers设置方法详解](http://www.ahlinux.com/start/cmd/457.html)

`/etc/sudoers`文件中存在如下行, 定义了root用户有权限以任何用户的身份执行主机上的所有命令, 这个文件由`sudo`工具提供, 大多数linux发行版中都默认安装.

```
root    ALL=(ALL)       ALL
```

### 5.1 赋予普通用记sudo权限

如果一个用户知道`root`用户的密码, 可以使用`sudo 命令`以root身份执行, 但如果我们不想这个用户拥有root密码, 并且想要其可以拥有`root`用户的全部权限, 可以在这个文件中添加这样一行

```
general    ALL=(ALL)       ALL
```

这样, 普通用户`general`就拥有了`sudo su -`切换成`root`用户的能力, 并且可以使用`sudo 命令`执行系统上所有命令, 并且只需要输入`general`用户本身的密码即可.

```
## 普通用户使用netstat命令无法查询所有用户的网络使用情况, 只能查看以自己的身份启动的进程.
general@localhost$ netstat -nlp
(Not all processes could be identified, non-owned process info
 will not be shown, you would have to be root to see it all.)
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address               Foreign Address             State       PID/Program name   
tcp        0      0 0.0.0.0:10130               0.0.0.0:*                   LISTEN      32209/python        
Active UNIX domain sockets (only servers)
Proto RefCnt Flags       Type       State         I-Node PID/Program name    Path
general@localhost$ sudo !!
sudo netstat -nlp
[sudo] password for jumpserver3: 
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address               Foreign Address             State       PID/Program name   
tcp        0      0 0.0.0.0:10130               0.0.0.0:*                   LISTEN      32209/python        
Active UNIX domain sockets (only servers)
Proto RefCnt Flags       Type       State         I-Node PID/Program name    Path
```

```
## 直接执行su需要知道root用户的密码
general@localhost$ su
Password: 
su: incorrect password
## 因为用户general获得了sudo权限, 所以使用sudo su可以切换到root, 并且只需要输入自己的密码.
general@localhost$ sudo su -
[sudo] password for jumpserver3: 
[root@b14e517d408b ~]# 
```

如果希望该用户可以连本身的密码都不用询问, 可以这样写

```
general    ALL=(ALL)       NOPASSWD:ALL
```

### 5.2 赋予用户指定权限

如果我们希望普通用户general只能执行部分root命令, 不想他能直接切换成root管理系统, 我们可以在`/etc/sudoers`文件为其显式指定可以执行的命令列表.

我们分析一下上面那一行配置的格式

```
general    ALL=(ALL)       ALL
```

- general: 表示被授权的用户, 如果是为`/etc/group`文件中存在的组进行授权, 使用`%组名`

- 第一个ALL: 表示所有来源(从任何主机连接进来)

- 第二个ALL: 表示所有用户

- 第三个ALL: 表示所有命令

我们为general添加`useradd`, `userdel`权限, 这样其将拥有权限执行这两条命令, 但无法再使用`sudo su`切换成root了.

```conf
general ALL=(root) /usr/sbin/useradd,/usr/sbin/userdel
```

同样也可以为general使用`NOPASSWD`标记实现免密码执行.

注意: 命令列表要使用绝对路径, 否则会报如下错误.

```
sudo useradd jiangming
sudo: >>> /etc/sudoers: syntax error near line 92 <<<
sudo: parse error in /etc/sudoers near line 92
sudo: no valid sudoers sources found, quitting
```

### 5.3 批量授权

如果待授权用户很多且不在同一用户组, 或者指定的命令太多, 这样sudo规则书写起来会很麻烦. 我们可以使用`sudoers`文件中提供的`XXX_Alias`系列命令指定一组变量, 可以是一组用户, 一组权限, 或一组命令等.

Alias使用方法如下

```conf
User_Alias 变量名=变量值
Runas_Alias 变量名=变量值
Host_Alias 变量名=变量值
Cmnd_Alias 变量名=变量值
```

变量名必须要以大写字母开头，而且只能包含有大写字母，数字，下划线.

而变量值是以逗号','分隔的数组，不过这四个别名表示的数组内容都会不同.

比如想要赋予普通用户general以网络配置相关的root级别命令, 可以添加如下行

```conf
Cmnd_Alias NETWORKING = /sbin/route, /sbin/ifconfig, /bin/ping, /sbin/dhclient, /usr/bin/net, /sbin/iptables, /usr/bin/rfcomm, /usr/bin/wvdial, /sbin/iwconfig, /sbin/mii-tool
```

然后使用上面讲到的, 将`NETWORKING`字段添加给general用户

```
general ALL=(root) NETWORKING
```

这样, general就可以拥有root级别的, 执行`NETWORKING`定义的包括`route`, `ifconfig`...等一系列网络相关的命令, 是不是很方便?

------

然后我们看一下这四种类型的字段, 变量值可以取哪些值

```
User：[!][username | #uid | %groupname | +netgroup | %:nonunix_group | User_Alias]
Runas：[!][username| #uid | %groupname | +netgroup | Runas_Alias]
Host：[!][hostname | ip_addr | network(/netmask)? |  netgroup | Host_Alias]
Cmnd：[!][commandname| directory | "sudoedit" | Cmnd_Alias]
```

感叹号`!`表示取反, 比如不包括指定主机, 不包括指定用户, 禁止执行的命令等.

### 5.4 通配符(未验证)

通配符只可以用在`主机名`、`文件路径`、`命令行的参数列表`中。下面是可用的通配符：

- *：匹配任意数量的字符

- ?：匹配一个任意字符

- [...]：匹配在范围内的一个字符

- [!...]：匹配不在范围内的一个字符

- \x：用于转义特殊字符

在使用通配符时有以下的注意点：

1. 使用[:alpha:]等通配符时，要转义冒号':'，如：[\:alpha\:]

2. 当通配符用于文件路径时，不能跨'/'匹配，如：/usr/bin/*能匹配/usr/bin/who但不能匹配/usr/bin/X11/xterm

3. 如果指令的参数列表是""时，匹配不包含任何参数的指令。

4. ALL这个关键字表示匹配所有情况。

### 5.5 更深层的用户规则(未验证)

用户规则定义的语法如下：

```conf
User_List Host_List=(Runas_List1:Runas_List2) SELinux_Spec Tag_Spec Cmnd_List,...
```

下面对上面的语法进行说明一下：

- `User_List`（必填项）：指的是该规则是针对哪些用户的。

- `Host_List`（必填项）：指的是该规则针对来自哪些主机的用户。

- `Runas_List1`（可选项）：表示可以用sudo -u来切换的用户

- `Runas_List2`（可选项）：表示可以用sudo -g来切换的用户组

- `SELinux_Spec`（可选项）：表示SELinux相关的选项，可选值为ROLE=role 或 TYPE=type。本人对SELinux不太熟，以后再补充这里吧。

- `Tag_Spec`（可选项）：用于控制后面Cmnd_List的一些选项啦，可选值有下面这些，具体可以查阅man手册

```
'NOPASSWD:' | 'PASSWD:' | 'NOEXEC:' | 'EXEC:' | 'SETENV:' | 'NOSETENV:' | 'LOG_INPUT:' | 'NOLOG_INPUT:' | 'LOG_OUTPUT:' | 'NOLOG_OUTPUT:'
```

- `...`（可选项）：表示可以有多个(`Runas_List1`:`Runas_List2`) `SELinux_Spec` `Tag_Spec` `Cmnd_List`段的意思。

注意：如果`Runas_List1`和`Runas_List2`都没填的话，默认是以root用户执行