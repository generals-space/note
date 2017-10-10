# SVN服务器应用场景

## 1. 使用apache代理svn服务

参考文章

[Apache+SVN搭建SVN服务器](http://blog.csdn.net/linux7985/article/details/49127427)

使用`svnserve -d`启动的svn服务, 客户端连接时的svn地址前缀为`svn://`, 而通过使用apache(应该也可以是nginx), svn可以通过`http://`前缀访问. 这样可以方便的使用域名, 甚至使用`https://`来连接svn服务器.

### 1.1 安装apache

首先需要安装`apr`, `apr-util`, `pcre`包, 这里不再写出. 但apache的`configure`选项需要额外包含`--enable-dav`以支持svn目录读写等操作.

```
$ ./configure --prefix=/usr/local/apache2 \
--with-apr=/usr/local/apr --with-apr-util=/usr/local/apr-util \
--with-z --with-pcre \
--enable-so --enable-ssl --enable-cgi --enable-rewrite --enable-deflate \
--enable-modules=most --enable-mpms-shared=all --with-mpm=event \
--enable-dav
```

在编译`subversion`时, 配置选项需要额外指出`--with-apxs`与`--with-apache-libexecdir`, 这会在apache的modules目录生成`mod_dav_svn.so`模块. 并且此时需要指定上面使用源码安装apache的`apr`与`apr-util`目录, 而不能再是使用yum装的包了, 否则apache启动会出错.

```
$ ./configure --prefix=/usr/local/svn \
--with-apr=/usr/local/apr --with-apr-util=/usr/local/apr-util \
--with-apxs=/usr/local/apache2/bin/apxs --with-apache-libexecdir
make && make install
```

之后配置apache. 这里假设已经存在subversion版本库, 路径为`/svn`. 编辑`conf/httpd.conf`, 先加载如下模块

```conf
LoadModule dav_module modules/mod_dav.so
LoadModule dav_svn_module modules/mod_dav_svn.so
LoadModule authz_svn_module modules/mod_authz_svn.so
```

然后添加如下配置到`httpd.conf`底部.

```conf
<Location />
        DAV svn
        ## svn版本库目录路径
        SVNPath /svn
        Require valid-user
        AuthType Basic
        ## checkout的时候在对话框的提示
        AuthName "svn for matu"
        ## 密码文件的位置,文件名随意, 与svn的密码文件不同
        AuthUserFile conf/svnPasswd
        ## 权限文件,文件名随意, 可以与svn的权限文件相同
        AuthzSVNAccessFile conf/svnAuthz
</Location>
```

使用apache自带的`htpasswd`命令创建用户密码, 如下命令会生成`httpd.conf`文件中`AuthUserFile`指令指定的`svnPasswd`文件, 文件中用户的密码都是经过加密后的字符串.

```
## 第一次创建时, 需要使用-c选项指明, 先创建svnPasswd文件
$ /usr/local/apache2/bin/htpasswd -c /usr/local/apache2/conf/svnPasswd general
New password: 
Re-type new password: 
Adding password for user general
## 之后添加其他用户时, 不必再使用-c选项.
$ /usr/local/apache2/bin/htpasswd /usr/local/apache2/conf/svnPasswd test
New password: 
Re-type new password: 
Adding password for user test
```

将svn版本库下的`conf/authz`拷贝到`apache`的`conf`目录下, 这个权限文件, svn与apache可以通用.

```
$ cp /svn/conf/authz /usr/local/apache2/conf/svnAuthz
```

启动apache.

> 注意: 使用apache提供svn服务时不必再使用`svnserve -d`启动(当然, 启动了也没有什么影响).

然后客户端通过http协议连接svn.

```
$ svn checkout  http://172.32.100.136/ ./project
Authentication realm: <http://172.32.100.136:80> 这里是提示
Password for 'root': 
Authentication realm: <http://172.32.100.136:80> 这里是提示
Username: general
Password for 'general': 

-----------------------------------------------------------------------
ATTENTION!  Your password for authentication realm:

   <http://172.32.100.136:80> 这里是提示

can only be stored to disk unencrypted!  You are advised to configure
your system so that Subversion can store passwords encrypted, if
possible.  See the documentation for details.

You can avoid future appearances of this warning by setting the value
of the 'store-plaintext-passwords' option to either 'yes' or 'no' in
'/root/.subversion/servers'.
-----------------------------------------------------------------------
Store password unencrypted (yes/no)? yes
A    project/test.txt
Checked out revision 1.
```

### FAQ

#### 1

```
httpd: Syntax error on line 141 of /usr/local/apache2/conf/httpd.conf: Cannot load modules/mod_dav_svn.so into server: /usr/local/svn/lib/libsvn_subr-1.so.0: undefined symbol: apr_crypto_block_cleanup
```

问题描述: 安装完apache, subversion, 建立了版本库, 并启动了svn服务. 在apache中配置了svn后启动时报上述错误.

原因分析: apache与subversion安装时指定的`apr`与`apr-util`不一致.

解决方法: 指定一致的`apr`与`apr-util`包路径, 重新编译apache或subversion.

#### 2

```
Invalid command 'AuthzSVNAccessFile', perhaps misspelled or defined by a module not included in the server configuration
```

问题描述: 在`httpd.conf`文件中添加svn的验证方式, 重启时报错.

原因分析: 未加载`mod_authz_svn`模块

解决方法: 在`httpd.conf`文件中添加`LoadModule authz_svn_module modules/mod_authz_svn.so`即可.

#### 3

```
$ svn checkout  http://172.32.100.136/ ./project
Authentication realm: <http://172.32.100.136:80> 这里是提示
Password for 'root': 
Authentication realm: <http://172.32.100.136:80> 这里是提示
Username: general
Password for 'general': 
Authentication realm: <http://172.32.100.136:80> 这里是提示
Username: general
Password for 'general': 
svn: E170001: Unable to connect to a repository at URL 'http://172.32.100.136'
svn: E170001: OPTIONS of 'http://172.32.100.136': authorization failed: Could not authenticate to server: rejected Basic challenge (http://172.32.100.136)
```

提示: 通过apache代理的svn服务, 认证文件是通过`htpasswd`命令生成的加密文件, 并不是像`svnserve`直接使用的明文文件. 需要注意.

## 2. 多版本库配置

参考文章

[SVN服务器多个项目的权限分组管理](http://blog.sina.com.cn/s/blog_62cd41130102v4ro.html)

创建多个版本库

```
$ mkdir -p /svn/project1 /svn/project2
$ /usr/local/svn/bin/svnadmin create /svn/project1 
$ /usr/local/svn/bin/svnadmin create /svn/project2
```

在`/svn`目录下创建共用的`passwd`与`authz`文件, 两文件的内容如下

```
$ cat /svn/passwd
[users]
general = 123456
test1 = 123456
test2 = 123456

$ cat /svn/authz
[groups]
admin = general
user = test1, test2
pro1_user = test1
pro2_user = test2

[project1:/]
@admin = rw
@pro1_user = rw
@user = r

[project2:/]
@admin = rw
@pro2_user = rw
@user = r

[/]
@admin = rw
@user = r
* =
```

可以看到, project1项目的用户为`test1`, project2项目的用户为`test2`. `[project1:/]`块表示, `admin`组的成员与`pro1_user`组的成员对project1项目有读写权限, 而普通的`user`组只有读权限. 最后的`[/]`更是指明了, 不在任意组中的用户无任何权限.

然后配置每个工程的验证文件, 将每个工程的用户认证及权限配置文件, 都修改成使用上面定义的共享文件. 以`project1`为例. `project1/conf/svnserve.conf`.

```
//禁止匿名访问
anon-access = none
auth-access = write
//统一使用密码文件
password-db = /svn/passwd
authz-db = /svn/authz
//权限域名，很重要，写你的工程名
realm = project1
```

启动svn服务, 指定根目录为`/svn`.

```
$ svnserve -d -r /svn
```

### 客户端访问

由于是多版本库, 所以客户端连接到svn服务器时, 就不能直接指定根目录了, 还需要指定项目名.

```
$ svn list svn://172.32.100.136:3690
svn: E210005: Unable to connect to a repository at URL 'svn://172.32.100.136'
svn: E210005: No repository found in 'svn://172.32.100.136'
$ svn list svn://172.32.100.136:3690/project1
Authentication realm: <svn://172.32.100.136:3690> 1954a4e8-b099-11e6-9346-692b51c33849
Password for 'root': 
Authentication realm: <svn://172.32.100.136:3690> 1954a4e8-b099-11e6-9346-692b51c33849
Username: test1
Password for 'test1': 
...
```

以`test1`用户分别`checkout`并提交项目`project1`与`project2`.

```
$ svn checkout svn://172.32.100.136:3690/project1 pro1
Checked out revision 0.
$ cd pro1
$ touch test1
$ svn add ./test1 
A         test1
$ svn commit -m '测试test1'
Adding         test1
Transmitting file data .
Committed revision 1.

$ svn checkout svn://172.32.100.136:3690/project2 pro2
Authentication realm: <svn://172.32.100.136:3690> 1a760466-b099-11e6-b213-99bb3d63cdaa
Password for 'root': 
Authentication realm: <svn://172.32.100.136:3690> 1a760466-b099-11e6-b213-99bb3d63cdaa
Username: test1
Password for 'test1': 

-----------------------------------------------------------------------
ATTENTION!  Your password for authentication realm:

   <svn://172.32.100.136:3690> 1a760466-b099-11e6-b213-99bb3d63cdaa

can only be stored to disk unencrypted!  You are advised to configure
your system so that Subversion can store passwords encrypted, if
possible.  See the documentation for details.

You can avoid future appearances of this warning by setting the value
of the 'store-plaintext-passwords' option to either 'yes' or 'no' in
'/root/.subversion/servers'.
-----------------------------------------------------------------------
Store password unencrypted (yes/no)? yes
Checked out revision 2.
$ cd pro2
$ touch test2
$ svn add ./test2 
A         test2
$ svn commit -m 'test2 by test1'
svn: E170001: Commit failed (details follow):
svn: E170001: Authorization failed
```

可以看到, 用户`test1`可以对`project1`进行`commit`, 但对`project2`只能执行`checkout`, `commit`的话会报错.

> 注意: 新增版本库, 新建用户等操作即时生效, 无需重启svn服务.

------

如果使用apache代理svn服务, 则不能再使用`SVNPath`字段, 而是需要使用`SVNParentPath`, 因为前者的值是一个指定的版本库, 而后者指定svn版本库的父目录.

```
SVNPath /svn/project1
```

改成

```
SVNParentPath /svn
```

## 3. 清除用户登录信息

参考文章

[如何清除SVN的用户名和密码](http://jingyan.baidu.com/article/d45ad148ed12c469552b801b.html)

svn客户端保存了我们连接svn服务器时的登录信息, 如果需要切换身份连接时, 需要先将登录信息删除.

**windows下(TortoiseSVN)**

右键->tortoisesvn->setting

弹出窗口左侧->Saved Data->右侧Authentication-> clear

**Linux**

`rm ~/.subversion/auth/svn.simple`