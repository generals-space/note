# SVN服务器搭建-多版本库配置

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
## 禁止匿名访问
anon-access = none
auth-access = write
## 统一使用密码文件
password-db = /svn/passwd
authz-db = /svn/authz
## 权限域名，很重要，写你的工程名
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