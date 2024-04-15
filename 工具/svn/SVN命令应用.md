# SVN命令应用

[官方文档](http://svnbook.red-bean.com/)

安装svn客户端(好像还是服务器)

```
yum install subversion
svn help
```

## 1. 获取

svn要先`checkout`获得远程文件, 然后本地操作. 第一次连接服务器会进行身份认证.

```
svn checkout 服务器地址/文件路径 本地路径
```

checkout后会在此本地目录记录下对应的远程路径, 再次执行`svn list`就不用再加远程路径了.

svn可以单独同步某一个子目录, 不必同步根目录下的所有目录/文件

## 2. 删除

svn删除操作分为`本地删除`和`远程删除`两种, 而且这两个是明显隔开的.

单纯本地删除无法删掉远程文件, commit也不行

```
svn del 本地文件路径
svn commit
```

远程删除还需要写删除理由(随便写), 删除远程文件不会删除本地文件, 只有下一次拉取更新时本地文件才会消失.

```
svn del svn服务器远程文件路径 -m '删除理由'
```

## 3. 新增

```
## 将目标文件添加到更新列表, 如果不执行add命令, 则新增文件只会存在本地, 不会被提交到远程仓库
svn add 文件/目录
## 然后将更新提交到服务器, 必须使用-m选项, 否则提交出错
svn commit -m '更新信息'
```

## 4. 创建远程仓库目录

当想要在远程仓库指定目录下创建一个子目录, 然后再单独checkout这个子目录, 可以执行如下命令

```
$ svn mkdir -m 'shell脚本代码片段' https://SVN地址/snippet
```

由于是在远程创建, 所以需要日志信息, 使用`-m`指定. 然后再checkout即可.

如果在一个已经checkout的svn目录下创建新目录. 可以先`mkdir 目录名`再使用`svn add ./目录名`将其加入svn版本库中. 也可以直接使用`svn mkdir 目录名`.

------

## 5. checkout单独的文件

参考文章

[svn checkout单个文件](http://www.cnblogs.com/leaven/p/4386915.html)

我想单独下载一个目录下的指定文件(因为它所在的目录太大了, 没有必要为了一个文件下载全部的内容). 但`svn checkout path`中的path必须是一个目录, 如果是一个文件的话就会报如下错误.

```
svn: URL 你要下载的文件路径 refers to a file, not a directory
```

解决方法是, 下载目标文件所在的空目录到指定目录, 然后update指定的文件.

```
## 执行如下命令, 你会得到一个空目录(不过还是包含了.svn目录的)
$ svn checkout --depth=empty 目标文件所在目录(一般是绝对路径) 本地指定目录
## 然后update指定文件(可以是相对路径)
$ svn update 指定文件
```

## 6. svn info查询

包括checkout的svn具体路径, SVN根目录路径, 当前版本号等信息

```
[root@localhost serverInfo]# svn info
Path: .
Working Copy Root Path: /mnt/hgfs/Work/SVN/serverInfo
URL: https://svn.mopo.com:9991/svn/noc/%E5%80%BC%E7%8F%AD%E7%9B%91%E6%8E%A7/%E9%80%9A%E7%94%A8/2016%E4%B8%9A%E5%8A%A1%E8%BF%90%E7%BB%B4%E8%B5%84%E6%BA%90%E5%90%8C%E6%AD%A5
Repository Root: https://svn.mopo.com:9991/svn/noc
Repository UUID: eb8496a6-1e08-11df-818b-ad2196546191
Revision: 17190
Node Kind: directory
Schedule: normal
Last Changed Author: general
Last Changed Rev: 17190
Last Changed Date: 2016-12-20 08:17:01 +0800 (Tue, 20 Dec 2016)
```

## 7. 客户端清除用户登录信息

参考文章

[如何清除SVN的用户名和密码](http://jingyan.baidu.com/article/d45ad148ed12c469552b801b.html)

svn客户端保存了我们连接svn服务器时的登录信息, 如果需要切换身份连接时, 需要先将登录信息删除.

**windows下(TortoiseSVN)**

右键->tortoisesvn->setting

弹出窗口左侧->Saved Data->右侧Authentication-> clear

**Linux**

`rm -rf ~/.subversion/auth/svn.simple`