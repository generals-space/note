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
[root@20ce69da6dac tmp]# ll
total 20
-rw------- 1 root root 19151    Jul  1      20:22 anaconda-post.log
-rw-r--r-- 1 root root     0        Sep 22    05:48 ÿ3Ҩg[1ϕ??????eҬ
-rw------- 1 root root     0        Jul  1      20:20 yum.log
```

这里提供一个方法.

首先使用`ls`的`-i`选项, 取到目标文件的`inode`编号.

```
[root@20ce69da6dac tmp]# ls -li
total 20
8388946 -rw------- 1 root root 19151 Jul  1   20:22 anaconda-post.log
8427830 -rw-r--r-- 1 root root         0 Sep 22 05:48 ÿ3Ҩg[1ϕ??????eҬ
8388947 -rw------- 1 root root         0 Jul  1   20:20 yum.log
```

然后通过`find`命令删除它

```
[root@20ce69da6dac tmp]# ls -li
total 20
8388946 -rw------- 1 root root 19151 Jul  1 20:22 anaconda-post.log
8427830 -rw-r--r-- 1 root root     0 Sep 22   05:48 ÿ3Ҩg[1ϕ??????eҬ
8388947 -rw------- 1 root root     0 Jul  1     20:20 yum.log
[root@20ce69da6dac tmp]# find -inum 8427830 -delete
[root@20ce69da6dac tmp]# ls
anaconda-post.log  yum.log
```

## 6. Linux生成随机密码

###  6.1 使用mkpasswd命令

这个命令不是单独存在, 而是在`expect`包里, 要安装它就直接装`expect`.

```
yum -y install expect
```

可以设置密码长度, 特殊字符个数等信息.

## 7. find命令使用

### 7.1 按名称查找

```
## find 目标路径 -name '目标名称'
$ find /etc/ -name 'yum.conf'
```

可以使用`*`与`?`通配符查找

### 7.2 按照时间查找

```
## find 目标路径 {-atime|-ctime|-mtime|-amin|-cmin|-mmin} [-|+]num
```

使用atime/ctime/mtime时, 后面的num值将被视为天数; 使用amin/cmin/mmin时, 后面的num值将被视为分钟数, 也就是说这被当作num的单位. 关于a, c, m这三种时间的含义需要自行了解.

后面的`[-|+]num`, 表示计时的时间段. `-num`表示**从这个时间开始**, `+num`表示**到这个时间为止**, 不带符的话将被看作是限制完全符合这个时间点. 

例如, `-mtime -3`, 可以表示查找3天内被修改过的文件(4天前被修改过的文件就不会出现啦).

```
## 查找最近一次被修改是在1000天前或者更早的文件
$ find /var/log/httpd -mtime +1000
## 如果不加加号或减号就表示是1000前那天被修改过的文件
```

### 7.3 执行命令

参考文章

[每天一个linux命令（20）：find命令之exec](http://www.cnblogs.com/peida/archive/2012/11/14/2769248.html)

我们使用`find`, 很多时候并不是单单只是想看看有哪些文件而已, 比如删掉查出来的很早的文件, 或是查看这些文件的详细信息. 这个时候可以使用`find`的`exec`参数.

`-exec`参数后面跟的是普通的bash命令，它的终止是以`;`为结束标志的，所以命令后面的分号是不可缺少的，考虑到各个系统中分号会有不同的意义，所以前面加反斜杠, 而`{}`代表查找出的文件。

```
## 显示详细信息
find ./ -mtime +1000 -exec ls -l {} \;
## 或者删掉它们
find ./ -mtime +1000 -exec rm {} \;
```

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

## 11 crontab定时任务

### 11.1 权限控制

crontab 用来任务定时调度，在Linux下可以通过创建文件`/etc/cron.allow`或者`/etc/cron.deny` 
来控制权限，如果`/etc/cron.allow`文件存在，那么只有这个文件中列出的用户可以使用`cron`，同时 
`/etc/cron.deny`文件被忽略； 如果`/etc/cron.allow`文件不存在，那么文件`/cron.deny`中列出的用户将不能用使用`cron`。

添加要限制的用户，只需要写入用户名即可。

### 11.2 书写格式

执行`crontab -l`可以查看当前用户的定时任务列表, `crontab -e`可以编辑当前用户的过时任务

任务列表最终保存在`/var/spool/cron`, 以每个用户的用户名为名称.

```
## 其中命令应该写绝对路径
分　 时　 日　 月　 周　 命令
```

第1列表示第几分钟(1～59), 每分钟用`*`或者`*/1`表示

第2列表示第几小时(0～23),（0表示0点）, 同理每分小时为`*`或者`*/`

第3列表示日期(1～31)

第4列表示月份(1～12)

第5列表示号星期(0～6)（0表示星期天）

第6列要运行的命令, 可以有空格

示例

```
#每晚的21:30重启lighttpd。
30 21 * * * /usr/local/etc/rc.d/lighttpd restart
#每月1、10、22日
45 4 1,10,22 * * /usr/local/etc/rc.d/lighttpd restart
#每天早上6点10分
10 6 * * * date
#每两个小时
0 */2 * * * date
#晚上11点到早上8点之间每两个小时，和早上8点...有点复杂
0 23-7/2,8 * * * date
#每个月的4号和每个礼拜的礼拜一到礼拜三的早上11点
## ...这个更复杂, 月份中的日期与星期中的日期貌似不冲突, 小时数相同时竟然可以共用
0 11 4 * 1-3 date
#1月份第天早上4点
0 4 1 * * date 
```

很多时候，我们计划任务需要精确到秒来执行，根据以下方法，可以很容易地以秒执行任务。
以下方法将每10秒执行一次

```
# crontab -e
* * * * * /bin/date >>/tmp/date.txt
* * * * * sleep 10; /bin/date >>/tmp/date.txt
* * * * * sleep 20; /bin/date >>/tmp/date.txt
* * * * * sleep 30; /bin/date >>/tmp/date.txt
* * * * * sleep 40; /bin/date >>/tmp/date.txt
* * * * * sleep 50; /bin/date >>/tmp/date.txt
```
 
注意如果用如果命令用到%的话需要用`\`转义

```
# backup mysql
00 01 * * * mysqldump -u root --password=passwd-d mustang > /root/backups/mustang_$(date +\%Y\%m\%d_\%H\%M\%S).sql
01 01 * * * mysqldump -u root --password=passwd-t mustang > /root/backups/mustang-table_$(date +\%Y\%m\%d_\%H\%M\%S).sql
```