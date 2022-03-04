# git编译问题

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
yum install tcl tk gettext
```

编译完成后, 执行`git clone`或`git push`等与远程仓库交互的操作时, 报错如下

```
$ git push origin
fatal: Unable to find remote helper for 'https'
```

问题分析

编译时未安装`curl-devel`库

解决方法

安装`curl-devel`, 重新`configure`并重新编译.
