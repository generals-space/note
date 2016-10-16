# FTP服务器配置

## 1. 访客模式(guest)

>　　这种方式在网上普遍被称作ftp的`虚拟用户`方式. 基本原理就是, 在ftp服务器上创建一个用户A, 然后`vsftpd`服务以这个A用户的身份创建任意个虚拟用户(这些虚拟用户只对`vsftpd`服务有效, 实际上并不存在于系统中), 虚拟用户所拥有的权限至多为A的权限, 并且可以为这些虚拟用户单独设置其各自的用户主目录, 用户读写权限, 会话失效时间, 文件传输速率等等. 所以这种方式搭建的ftp服务器使用方便, 且非常安全.

首先, **关闭防火墙与SELinux**.

然后安装`vsftpd`

```
$ yum -y install vsftpd*  db4-*
```

建立vsftp虚拟宿主用户，虚拟用户并不是系统用户，在系统中是不存在的，但它们需要一个实际的系统用户作为宿主用户。可以设置为`nologin`的登陆模式.

```
$ useradd ftpuser -s /sbin/nologin
```

### 1.1 配置文件理解

接下来开始配置vsftp. 通过yum安装的vsftpd服务, 配置文件默认在`/etc/vsftpd`目录下.

```
$ pwd
/etc/vsftpd
$ ls
ftpusers  user_list  vsftpd.conf  vsftpd_conf_migrate.sh
```

其中, `ftpusers`与`user_list`中都为系统用户列表, 不同的是, `ftpusers`文件中列出的是绝对禁止通过ftp方式登陆的系统用户; 而`user_list`中的用户, 需要根据`vsftpd.conf`文件中的`userlist_deny={YES|NO}`字段判断. 其值取`YES`时, 只有该文件列出的用户允许通过ftp登陆, `NO`时, 该文件中的用户连同`ftpusers`中列出的用户都将被禁止.

而`vsftpd.conf`为`vsftpd`服务的主配置文件.

### 1.2 创建虚拟用户

我们先新建一个虚拟用户列表文件`viruser`, 在这个列表中的用户是允许ftp登陆的合法用户. 其格式为`一行用户名, 一行密码`. 如下

```
$ pwd
/etc/vsftpd
$ cat viruser
ftpuser1
123456
general
123456
```

然后执行`db_load -T -t hash -f /etc/vsftpd/virusers /etc/vsftpd/virusers.db`, 这将会以`virusers`文件为准生成一个`virusers.db`文件, 用于ftp的登陆验证.

然后编辑`/etc/pam.d/vsftpd`, **将其中原来的内容注释掉或删除**, 添加如下几行

```
auth sufficient /lib64/security/pam_userdb.so db=/etc/vsftpd/virusers
account sufficient /lib64/security/pam_userdb.so db=/etc/vsftpd/virusers
```

一般64位系统为`/lib64/security/pam_userdb.so`, 而32位系统为`/lib/security/pam_userdb.so`. 另外, `db=/etc/vsftpd/virusers`中`virusers`并不是指`db=/etc/vsftpd/virusers`文件本身, 而是 **名为`virusers`并且后缀为`.db`的文件**.

注意: ftp虚拟用户登陆验证的实际文件不是`virusers`, 而之后也可以将此文件删除或限制其读权限, 以免其中明文存储的用户名密码被窃取.

------

然后为虚拟用户创建各自的配置文件及ftp主目录, 以`general`这个虚拟用户为例.

```
## 用以存储虚拟用户的配置文件
$ mkdir /etc/vsftpd/viruser.d
## 虚拟用户配置文件名需要与用户名相同, 而且不可以加任何后缀!!!
$ touch /etc/vsftpd/viruser.d/general
## 虚拟用户general的ftp主目录, 并将权限赋予ftp宿主用户`ftpuser`
$ mkdir -p /opt/vsftpd/general
$ chown -R /opt/vsftpd
```

### 1.3 配置虚拟用户

首先, 在`/etc/vsftpd/vsftpd.conf`文件中添加如下语句

```
## vsftpd服务的验证将通过PAM模块.
pam_service_name=vsftpd
userlist_enable=YES
tcp_wrappers=YES

## 设定启用虚拟用户功能
guest_enable=YES
## 虚拟用户的配置目录
user_config_dir=/etc/vsftpd/viruser.d
## 指定虚拟用户的宿主用户, 就是我们开始时创建的`ftpuser`
guest_username=ftpuser
## 设定虚拟用户的权限符合他们的宿主用户, 这一句也很重要!!!
virtual_use_local_privs=YES
## 限制虚拟用户只能在他们的主目录而不能进入其他目录
chroot_local_user=YES
```

然后编辑`/etc/vsftpd/viruser.d/general`, 添加

```
## 如果不添加这一行, 则用户的主目录为其ftp的宿主用户家目录
local_root=/opt/vsftpd/general
```

启动`vsftpd`服务.

```
service vsfptd start
```

### 1.4 测试ftp登陆

在客户端上通过ftp尝试登陆.

```
$ ftp
ftp> open 172.17.0.4
Connected to 172.17.0.4 (172.17.0.4).
220 (vsFTPd 2.2.2)
Name (172.17.0.4:general): general
331 Please specify the password.
Password:
230 Login successful.
Remote system type is UNIX.
Using binary mode to transfer files.
ftp> ls
227 Entering Passive Mode (172,17,0,4,116,226).
150 Here comes the directory listing.
226 Directory send OK.
ftp> mkdir first
257 "/opt/vsftpd/general/first" created
ftp> quit
221 Goodbye.
```

到ftp服务端`/opt/vsfptd/general`下查看, 会发现测试登陆时创建的`first`目录的所属用户为`ftpuser`, 正是我们创建的ftp服务的宿主用户. 这也是`虚拟`用户的涵义所在.

------

**扩展**

以下是一些灵活的权限控制, 分别在`/etc/vsftpd/vsftpd.conf`与`/etc/viruser.d/*`下的配置文件中有效, 不过前者中是全局性的, 后者中是单独对某一虚拟用户起作用的, 有些配置后者可以覆盖前者.

```
## 限制虚拟用户只能在其主目录中操作. NO的话用户可以去其他比如`/etc`等目录
chroot_local_user=YES
```

在`/etc/viruser.d/general`文件中有效(注意: `viruser.d`这个目录不是一定的, 而是通过`/etc/vsftpd.conf`中`user_config_dir`字段指定的). 另外, 对虚拟用户各自的配置文件的修改貌似不需要重启`vsftpd`服务, 即时生效, 但是客户端需要重新连接.

```
## 指定虚拟用户的主目录, 用户属主需要是其宿主用户
local_root /opt/vsftpd/general
## 限制用户的下载速度, 单位是B, 0为不限制
local_max_rate=0
```

------

命令控制

`cmds_allowed`字段可以限制客户端可以执行的命令, 在`/etc/vsftpd/vsftpd.conf`与`/etc/vsftpd/viruser.d/*`文件中都有效, 但以前者为准. 其格式为

```
cmds_allowed=FEAT,REST,CWD,LIST,MDTM,NLST,PASS,PASV,PORT,PWD,QUIT,RETR,SIZE,STOR,TYPE,USER,ACCT,APPE,CDUP,HELP,MODE,NOOP,REIN,STAT,STOU,STRU,SYST,MKD
```

注意: 不能又空格和换行, 另外命令名称必须为大写, 小写的命令将被拒绝执行. 只有显式的指定的命令允许被执行, 未列出的将被拒绝.

```
# ABOR - abort a file transfer 取消文件传输
# CWD - change working directory 更改目录
# DELE - delete a remote file 删除文件
# LIST - list remote files 列出目录内容
# MDTM - return the modification time of a file 返回文件的更新时间
# MKD - make a remote directory 新建文件夹
# NLST - name list of remote directory
# PASS - send password
# PASV - enter passive mode
# PORT - open a data port 打开一个传输端口
# PWD - print working directory 显示当前工作目录
# QUIT - terminate the connection 退出
# RETR - retrieve a remote file 下载文件
# RMD - remove a remote directory
# RNFR - rename from
# RNTO - rename to
# SITE - site-specific commands
# SIZE - return the size of a file 返回文件大小
# STOR - store a file on the remote host 上传文件
# TYPE - set transfer type
# USER - send username

# less common commands:
# ACCT* - send account information
# APPE - append to a remote file
# CDUP - CWD to the parent of the current directory
# HELP - return help on using the server
# MODE - set transfer mode
# NOOP - do nothing
# REIN* - reinitialize the connection
# STAT - return server status
# STOU - store a file uniquely
# STRU - set file transfer structure
# SYST - return system type
```
