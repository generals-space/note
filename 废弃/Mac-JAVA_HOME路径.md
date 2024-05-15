# Mac-JAVA_HOME路径

参考文章

1. [MAC系统-JAVA_HOME配置](https://blog.csdn.net/q62506212/article/details/96448410)

MacOS: 10.15.4 (19E287)
Java: 1.8.0

Mac 下的 Java 1.8 是通过`dmg`安装包安装的, 不像 win 和 linux, 安装位置可以控制, 而且安装完成后也没有设置`JAVA_HOME`环境变量.

使用`whereis`查看`java`可执行文件的路径, 但这并不是`jdk`的路径, 软链接太多, 很烦.

参考文章1给出了详细的方法.

```
The $JAVA_HOME on Mac OS X should be found using the /usr/libexec/java_home command line tool on Mac OS X 10.5 or later.  
```

参照苹果的官方文档说明，可以使用`/usr/libexec/java_home`命令查看`JAVA_HOME`.

> `java_home`并不在`PATH`路径下, 所以需要使用绝对路径执行.

```console
$ /usr/libexec/java_home
/Library/Java/JavaVirtualMachines/jdk1.8.0_251.jdk/Contents/Home
$ /usr/libexec/java_home -V
Matching Java Virtual Machines (1):
    1.8.0_251, x86_64:	"Java SE 8"	/Library/Java/JavaVirtualMachines/jdk1.8.0_251.jdk/Contents/Home

/Library/Java/JavaVirtualMachines/jdk1.8.0_251.jdk/Contents/Home
```

可以设置`JAVA_HOME`如下

```bash
export JAVA_HOME=/Library/Java/JavaVirtualMachines/jdk1.8.0_251.jdk/Contents/Home
```
