# Linux命令技巧

## 1. 使用`nc`命令传输文件

可以绕过scp, 不只可以传输文本文件, 不过不可以传目录, 打包之后倒是可以的.

远程主机(监听端口任意, 不过客户端需要指定, 与其保持相同):

```shell
nc -l 4567 > /tmp/cache.tar.gz
```

本地主机:

```shell
nc 远程主机IP 4567 < redis.tar.gz
```

传输完成之后双方自动退出.

## 2. 强制性操作

这些方法一般用于脚本中, 以防止脚本执行被打断. 在实际系统中并不建议直接使用强制手段.

### 2.1 强制复制

```shell
cp -f 源文件 目标文件
```

`-r`选项可以用于复制目录. **`scp`命令也是如此**.

`-f --force`选项如果表示目标文件已存在, 先将其删除再重新复制一次, 如果没有这个选项, 将会提示你是否确定要覆盖目标文件.

> PS: 有的时候即使使用了`-f`选项, 也依然会得到提示, 无法直接覆盖, 为什么?

> 这有可能是因为系统默认对`cp`命令使用了别名`alias cp='cp -i'`, `-i`选项将禁止覆盖目标文件. 直接使用`alias`查看当前所有的别名.

> 解决方法是使用`\cp -f 源文件 目标文件`, 使用`\cp`复制时不走`alias`;

> 另外, 也可以使用`unalias cp; cp -f 源文件 目标文件`, 这是临时取消`cp`的`alias`(不过这个方法可以影响整个终端, 感觉有点风险).

### 2.2 强制删除

```shell
rm -rf 目标文件/目录
```

`-r`选项可用于删除目录;

`-f`选项可以防止删除原来不存在的目标文件/目录时命令中止的情况.

## 2.3 强制创建目录

```shell
mkdir -p /tmp/first/second
```

`-p`命令可以防止当创建一个目录, 而你指定的它的父目录也不存在时的错误情况. 它会在创建指定的目录之前, 将所需要的父目录创建出来.

## 3. wget下载jdk

参考文章

[Linux 使用wget 命令下载JDK的方法](http://www.oschina.net/code/snippet_875267_44726)

```
wget --no-check-certificate --no-cookies --header "Cookie: oraclelicense=accept-securebackup-cookie" http://download.oracle.com/otn-pub/java/jdk/7u71-b14/jdk-7u71-linux-x64.rpm
```

## 4. 使用rz/sz命令上传/下载文件

通过在XShell的远程服务器的shell命令行中执行`rz`, `sz`可以方便地执行下载/上传操作, 而不使用filezilla或xftp这样的文件客户端.

这两个命令默认没有在linux系统中, 首先使用yum安装

```
yum install -y lrzsz
```

上传/下载大文件使用`rz`或`sz`的`-e`选项, 否则可能会乱码...md, 有时候还要加`-b`选项, 文本文件可能什么都不用加...每种还不一样.

另外, rz上传文件时, 如果没有当前目录的写权限, 或是当前目录存在同名文件, 会显示传输失败;

而且, sz下载文件时, 下载目录也必须是xshell启动用户有写权限的才行. 比如我并不是以`Administor`身份登录, 而是以一个普通用户A身份, 如果A启动xshell, 是没有办法将文件下载到C盘的(C盘的Users/A的个人目录还是可以的). 当sz, xshell弹出对话框让你选择下载路径时, 如果你选择一个没有写权限的C盘路径, xshell将不会有任何反应, 也没有报错, 当时困扰了很久. 这一点SecureCRT做的好一点, 它是提前设置的下载目录, 当选择目录如果没有写权限, 是没有办法将其设置为下载路径的.

## 5. 删除乱码文件

许多文件的名称显示为乱码, 无法选中也无法删除, 比如:

```
$ ll
total 20
-rw------- 1 root root 19151    Jul  1      20:22 anaconda-post.log
-rw-r--r-- 1 root root     0        Sep 22    05:48 ÿ3Ҩg[1ϕ??????eҬ
-rw------- 1 root root     0        Jul  1      20:20 yum.log
```

这里提供一个方法.

首先使用`ls`的`-i`选项, 取到目标文件的`inode`编号.

```
$ ls -li
total 20
8388946 -rw------- 1 root root 19151 Jul  1   20:22 anaconda-post.log
8427830 -rw-r--r-- 1 root root         0 Sep 22 05:48 ÿ3Ҩg[1ϕ??????eҬ
8388947 -rw------- 1 root root         0 Jul  1   20:20 yum.log
```

然后通过`find`命令删除它

```
$ ls -li
total 20
8388946 -rw------- 1 root root 19151 Jul  1 20:22 anaconda-post.log
8427830 -rw-r--r-- 1 root root     0 Sep 22   05:48 ÿ3Ҩg[1ϕ??????eҬ
8388947 -rw------- 1 root root     0 Jul  1     20:20 yum.log
$ find -inum 8427830 -delete
$ ls
anaconda-post.log  yum.log
```

## 6. Linux生成随机密码

###  6.1 使用mkpasswd命令

这个命令不是单独存在, 而是在`expect`包里, 要安装它就直接装`expect`.

```
yum -y install expect
```

可以设置密码长度, 特殊字符个数等信息.


## 8. 查看进程启动时间及运行时间

参考文章

[linux下查看一个进程的启动时间和运行时间](http://www.cnblogs.com/fengbohello/p/4111206.html)

```
## -A表示所有进程, -o表示输出格式(stime: start time, 启动时间; etime: elapsed time, 消逝的时间, 即运行时间, args: 启动命令及参数)
## stime如果超过一年就只能显示年的数字而不能再显示日期, 运行时间可以看到启动的天数和精确到秒级的计算结果
$ ps -A -o pid,stime,etime,args
  PID STIME     ELAPSED COMMAND
    1  2014 846-23:09:53 /sbin/init
11883 Jan29 254-23:59:04 java -Xbootclasspath/a:. -Denv=product -Ddubbo.properties.file=conf/product/dubbo.properties -Djava.a
12767 Mar07 216-16:42:12 java -Xbootclasspath/a:. -Denv=product -Ddubbo.properties.file=conf/product/dubbo.properties -Djava.a
14552 Jan29 254-23:55:53 java -Xbootclasspath/a:. -Denv=product -Ddubbo.properties.file=conf/product/dubbo.properties -Djava.a
15185 Jan29 254-23:55:07 java -Xbootclasspath/a:. -Denv=product -Ddubbo.properties.file=conf/product/dubbo.properties -Djava.a
15813  2015 400-00:22:20 java -Xbootclasspath/a:. -Denv=product -Ddubbo.properties.file=conf/product/dubbo.properties -Djava.a
```

## 9. curl使用

### 9.1 输出信息及格式设置

curl命令内置了许多输出，如状态码，抓取速度，总时间等，可通过`-w`选项选择性输出.

```shell
## 输出抓取百度首页的平均速度
$ curl -s -o baidu -w "%{speed_download}\n" www.baidu.com
61669.000
## 平均速度与总时间
$ curl -s -o baidu -w "--%{speed_download}--%{time_total}--\n" www.baidu.com
--96451.000--0.025--
```

参数`-s`是为了防止curl的默认输出, 包括响应时间, 下载速度和下载进度等, 不然显示会很杂乱.

### 9.2 使用代理

使用代理(使用10.10.10.10:10的代理访问google)

```
$ curl -x 10.10.10.10:10 www.google.com
```

使用wget达到同样的效果

```
## 选项`-Y`: 是否使用代理; `-e`执行命令
$ wget -Y on -e 'http_proxy=http://10.10.10.10:10' 'www.google.com'
```

## 10 获取命令执行时间

bash内置了一个time命令，功能比较少，`/usr/bin/time`是具有更强大功能的另一个命令，可以有格式化输出。例如`/usr/bin/time -f %e 待测命令`

time的默认输出是在stderr中的, 有时用`var=$(time [option] command [arguments])`进行变量赋值时会得到空值.

使用下面的命令可以解决这个问题.

```
$ var=$(/usr/bin/time -f %e curl -s -o baidu www.baidu.com 2>&1)
$ echo $var
```

curl的`-s`选项必不可少，不然curl的输出会扰乱变量var的赋直. 另外，注意`$()`的包裹范围，把`2>&1`也圈进去了.


## 12. useradd创建带密码的用户

参考文章

[linux下创建带密码的用户](http://blog.csdn.net/dliyuedong/article/details/24228599)

`useradd`命令有一个`-p`选项, 可在创建用户的同时为其指定密码. 但这个密码是经过`crypt`加密过的, 即必须是密文(验证时使用对应的明文). `man`手册中对其的介绍如下

```
    -p, --password PASSWORD
        The encrypted password, as returned by crypt(3). The default is to disable the password.

        Note: This option is not recommended because the password (or encrypted password) will be visible by users listing the processes.

        You should make sure the password respects the system's password policy.
```

需要注意的是, `crypt`并不是一个命令, 而是一个头文件(使用`whereis crypt`可以查看到相关信息), 它在`openssl-devel`包中. 我们并没有办法使用它.

替代方法是使用`openssl`命令. `openssl`有一个子命令`passwd`用来生成密码. 它的相关帮助信息如下

```
$ openssl passwd --help
Usage: passwd [options] [passwords]
where options are
-crypt             standard Unix password algorithm (default)
-1                 MD5-based password algorithm
-apr1              MD5-based password algorithm, Apache variant
-salt string       use provided salt
-in file           read passwords from file
-stdin             read passwords from stdin
-noverify          never verify when reading password from terminal
-quiet             no warnings
-table             format output as table
-reverse           switch table columns
```

其中`-crypt`, `-1`与`-apr1`三者是并列关系, 如果不指定, 则默认为`-crypt`的加密方式, 这也是系统密码的默认算法. 

我们可以使用如下方法获得加密后的密码

```
$ openssl passwd 123456
VGmG1B363td7o
```

然后创建目标用户, 如`general`

```
$ useradd general -p VGmG1B363td7o
```

之后就可以使用`general`及其密码`123456`进行登录或切换用户等操作了.

更为简便一点的方式

```
$ useradd general -p $(openssl passwd 123456)
```

除了这种方式, 也可以交互式从终端输入密码, 或是从文件读取然后批量生成.

```
$ openssl passwd
Password: <=输入123456
Verifying - Password: <=输入123456
hE/MhvW1qosEs
```

**注意: 虽然明文都是123456, 但加密后的密文并不相同. 即使如此, 创建用户时指定这些密文为密码, 其明文都会是123456, 不必担心**

从文件读取

```
$ cat passwd
123456
123456
## 这里得到的密码依然不一样哦
$ openssl passwd -in ./passwd
mEX6A8pyoI2do
UtFubbGdts3fg
```

## 13. Linux下查看高CPU占用率线程

参考文章

[Linux下如何查看高CPU占用率线程](http://itindex.net/detail/45450-linux-cpu-%E7%BA%BF%E7%A8%8B)

在Linux下`top`工具可以显示cpu的平均利用率(user,nice,system,idle,iowait,irq,softirq,etc.)，可以显示每个cpu的利用率。但是无法显示每个线程的cpu利用率情况， 
这时就可能出现这种情况，总的 cpu 利用率中`user`或`system`很高，但是用进程的 cpu 占用率进行排序时，没有进程的`user`或`system`与之对应。

如下图, 服务被入侵, 植入了挖矿服务, 杀掉`minerd`服务后CPU占用依然很高, 猜测是存在后台进程一直在检测, 但是`top`没法看到哪一个进程CPU占用率如此高.

![](https://gitimg.generals.space/bc3643ce87a37194cd61427bb0939ffa.png)

可以用下面的命令将 cpu 占用率高的线程找出来: 

```
$ ps H -eo user,pid,ppid,tid,time,%cpu,cmd --sort=%cpu
```

这个命令首先指定参数'H'，显示线程相关的信息，格式输出中包含:user,pid,ppid,tid,time,%cpu,cmd，然后再用`%cpu`字段进行排序。这样就可以找到占用处理器的线程了。

查到的结果如下图.

![](https://gitimg.generals.space/ceb8d634b41e796e3b6c98a8750ee88d.png)
