# SVN服务器搭建-使用apache代理svn服务

参考文章

[Apache+SVN搭建SVN服务器](http://blog.csdn.net/linux7985/article/details/49127427)

使用`svnserve -d`启动的svn服务, 客户端连接时的svn地址前缀为`svn://`, 而通过使用apache(应该也可以是nginx), svn可以通过`http://`前缀访问. 这样可以方便的使用域名, 甚至使用`https://`来连接svn服务器.

> 要用apache, 得先装apache再装svn...否则要做很多事情.

## 1. 安装apache

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

安装好svn之后配置apache. 这里假设已经存在subversion版本库, 路径为`/svn`. 编辑`conf/httpd.conf`, 先加载如下模块

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

## 2. FAQ

### 2.1

```
httpd: Syntax error on line 141 of /usr/local/apache2/conf/httpd.conf: Cannot load modules/mod_dav_svn.so into server: /usr/local/svn/lib/libsvn_subr-1.so.0: undefined symbol: apr_crypto_block_cleanup
```

问题描述: 安装完apache, subversion, 建立了版本库, 并启动了svn服务. 在apache中配置了svn后启动时报上述错误.

原因分析: apache与subversion安装时指定的`apr`与`apr-util`不一致.

解决方法: 指定一致的`apr`与`apr-util`包路径, 重新编译apache或subversion.

### 2.2

```
Invalid command 'AuthzSVNAccessFile', perhaps misspelled or defined by a module not included in the server configuration
```

问题描述: 在`httpd.conf`文件中添加svn的验证方式, 重启时报错.

原因分析: 未加载`mod_authz_svn`模块

解决方法: 在`httpd.conf`文件中添加`LoadModule authz_svn_module modules/mod_authz_svn.so`即可.

### 2.3

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
