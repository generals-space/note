# Git 所遇问题

## 1. `.gitignore`文件无效

参考文章

[.gitignore无效, 不能过滤某些文件](http://blog.csdn.net/zhangjs0322/article/details/37658669)

利用.gitignore过滤文件, 如编译过程中的中间文件, 等等, 这些文件不需要被追踪管理.

问题描述:

在`.gitignore`添加`file1`文件, 以过滤该文件, 但是通过`git status`查看仍显示file1文件的修改状态.

原因分析:

`.gitignore`文件只对还没有加入版本管理的文件起作用, 如果之前已经用git把这些文件纳入了版本库并`commit`过, 就不起作用了.

解决方法:

需要在git库中删除该文件, 并更新(`git commit`). 然后再次`git status`查看状态, file1文件不再显示状态.

## 2.

`git clone`或`git push`操作时报错如下

```
git push origin
error: The requested URL returned error: 401 Unauthorized while accessing https://git.oschina.net/generals-space/ansible.git/info/refs

fatal: HTTP request failed
```

问题分析

git版本问题, 当前版本为`1.7.1`

解决办法

重新安装高版本的git即可, 最好是1.9+

## 3. 编译安装git时遇到的问题

### 3.1

编译时遇到问题

```
$ make install
GITGUI_VERSION = 0.20.GITGUI
    * new locations or Tcl/Tk interpreter
    GEN git-gui
    INDEX lib/
    * tclsh failed; using unoptimized loading
    MSGFMT    po/bg.msg make[1]: *** [po/bg.msg] Error 127
make: *** [all] Error 2
```

解决办法

```
yum install tcl  tk gettext
```

### 3.2

编译完成后, 执行`git clone`或`git push`等与远程仓库交互的操作时, 报错如下

```
$ git push origin
fatal: Unable to find remote helper for 'https'
```

问题分析

编译时未安装`curl-devel`库

解决方法

安装`curl-devel`, 重新`configure`并重新编译.

## 4. git无法提交空目录

git默认无法提交空目录, 解决办法是在空目录中创建类似`.gitignore`的隐藏文件占位, 一般命名为`.gitkeep`

## 5.

参考文章

[“Unable to find remote helper for 'https'” during git clone](http://stackoverflow.com/questions/8329485/unable-to-find-remote-helper-for-https-during-git-clone)

情境描述:

源码安装的git, 2.9.3, 第一次执行`git clone`时出现如下报错

```
git clone https://git.oschina.net/generals-space/work-grapher.git ./grapher/
fatal: Unable to find remote helper for 'https'
```

原因分析与解决方法

这是因为缺少`curl-devel`库的缘故, 需要首先`yum install curl-devel`, 然后重新`configure`并编译安装.