# SVN服务器搭建

参考文章

[Linux搭建SVN 服务器](https://my.oschina.net/lionel45/blog/298305)

[Linux服务器配置——搭建SVN服务器](http://blog.csdn.net/a649518776/article/details/39433883)

系统版本: CentOS7

SVN版本: 1.9.4

## 1. 安装

源码包中存在`INSTALL`文件, 有详细的安装步骤.

### 1.1 依赖安装

svn内部数据存储在sqlite中, 所以需要安装sqlite.

```
$ yum install -y \
autoconf libtool gcc gcc-c++ \
apr-devel apr-util-devel \
sqlite sqlite-devel \
libz-devel openssl-devel
```

### 1.2 编译, 安装

编译安装subversion.

```
$ ./configure --prefix=/usr/local/svn
make && make install
```

### 1.3 配置, 启动

首先使用svn命令创建版本库, 用于存储来自客户端的提交代码.

```
$ mkdir -p /svn
$ /usr/local/svn/bin/svnadmin create /svn
$ ls
conf  db  format  hooks  locks  README.txt
```

`/svn`下的子目录用途说明:

- hooks目录：放置hook脚本文件的目录

- locks目录：用来放置subversion的db锁文件和db_logs锁文件的目录，用来追踪存取文件库的客户端

- format文件：是一个文本文件，里面只放了一个整数，表示当前文件库配置的版本号

- conf目录：是这个版本库的配置文件（版本库的用户访问账号、权限等）

修改svn版本库的配置, 该目录下有4个文件.

```
$ ls
authz  hooks-env.tmpl  passwd  svnserve.conf
```

编辑`svnserve.conf`文件, 设置权限信息及验证文件

```
[general]
## 控制非鉴权用户访问版本库的权限. 
anon-access = none
## 控制鉴权用户访问版本库的权限 
auth-access = write
## 指定用户名口令文件名。 
password-db = /svn/conf/passwd
## 指定权限配置文件名，通过该文件可以实现以路径为基础的访问控制。 
authz-db = /svn/conf/authz
## 指定版本库的认证域，即在登录时提示的认证域名称。若两个版本库的认证域相同，建议使用相同的用户名口令数据文件 
realm = My Test Repository         
```


编辑`passwd`, 添加用户及其密码

```
[users]
general = 123456
test = 123456
```

编辑`authz`, 修改用户所属组, 并配置组权限

```
## 这里组名可以自定义, 不只是admin或user两种.
[groups]
admin = general
user = test

[/]
@admin = rw
@user = r
* =
```

> 注意：对用户配置文件(`passwd`, `authz`)的修改立即生效，不必重启svn服务。 

启动`svn`服务, 默认端口3960.

```
## `-d`: 以daemon形式运行, `-r`: 以/svn为根路径
$ svnserve -d -r /svn
```

## 使用方法

首次连接需要输入验证信息

```
$ svn list svn://172.32.100.136/
Authentication realm: <svn://172.32.100.136:3690> 80d199a0-5573-4433-b1b6-2c56e037b6dc
Password for 'root': 
Authentication realm: <svn://172.32.100.136:3690> 80d199a0-5573-4433-b1b6-2c56e037b6dc
Username: general 
Password for 'general': 

-----------------------------------------------------------------------
ATTENTION!  Your password for authentication realm:

   <svn://172.32.100.136:3690> 80d199a0-5573-4433-b1b6-2c56e037b6dc

can only be stored to disk unencrypted!  You are advised to configure
your system so that Subversion can store passwords encrypted, if
possible.  See the documentation for details.

You can avoid future appearances of this warning by setting the value
of the 'store-plaintext-passwords' option to either 'yes' or 'no' in
'/root/.subversion/servers'.
-----------------------------------------------------------------------
Store password unencrypted (yes/no)? yes
```

目前版本库中还没有文件, 所以list会得到空目录

```
$ svn list svn://172.32.100.136/
```

创建本地版本库并提交.

```
$ svn checkout svn://172.32.100.136/ ./project
Checked out revision 0.
$ cd project/
$ ls
$ vim test.txt
$ svn add ./test.txt 
A         test.txt
$ svn commit -m '测试文件'
Adding         test.txt
Transmitting file data .
Committed revision 1.
```

再次查询, 可以看到我们提交过的文件.

```
$ svn list svn://172.32.100.136/
test.txt
```
