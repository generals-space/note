# Linux服务脚本编写

参考文章

[linux centos service 参数详解](http://www.cnblogs.com/cosiray/p/5112809.html)

[Systemd (简体中文)](https://wiki.archlinux.org/index.php/Systemd_(%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87))

[Systemd 入门教程：实战篇](http://www.ruanyifeng.com/blog/2016/03/systemd-tutorial-part-two.html)

以`php-fpm`服务脚本为例

CentOS6环境下的脚本位置在`/etc/init.d/php-fpm`, 其内容如下

```bash
#! /bin/sh
#
# chkconfig: - 84 16
# description:	PHP FastCGI Process Manager
# processname: php-fpm
# config: /etc/php-fpm.conf
# config: /etc/sysconfig/php-fpm
# pidfile: /var/run/php-fpm/php-fpm.pid
#
### BEGIN INIT INFO
# Provides: php-fpm
# Required-Start: $local_fs $remote_fs $network $named
# Required-Stop: $local_fs $remote_fs $network
# Short-Description: start and stop PHP FPM
# Description: PHP FastCGI Process Manager
### END INIT INFO

# Standard LSB functions
#. /lib/lsb/init-functions

# source命令加载函数库
. /etc/init.d/functions

# Check that networking is up.
. /etc/sysconfig/network

# 额外的环境变量
if [ -f /etc/sysconfig/php-fpm ]; then
      . /etc/sysconfig/php-fpm
fi

if [ "$NETWORKING" = "no" ]
then
	exit 0
fi

RETVAL=0
prog="php-fpm"
pidfile=${PIDFILE-/var/run/php-fpm/php-fpm.pid}
lockfile=${LOCKFILE-/var/lock/subsys/php-fpm}

start () {
	echo -n $"Starting $prog: "
	dir=$(dirname ${pidfile})
	[ -d $dir ] || mkdir $dir
    ## daemon命令不是bash命令, 而是function函数库中的函数
	daemon --pidfile ${pidfile} /usr/sbin/php-fpm --daemonize
	RETVAL=$?
	echo
	[ $RETVAL -eq 0 ] && touch ${lockfile}
}
stop () {
	echo -n $"Stopping $prog: "
	killproc -p ${pidfile} php-fpm
	RETVAL=$?
	echo
	if [ $RETVAL -eq 0 ] ; then
		rm -f ${lockfile} ${pidfile}
	fi
}

restart () {
        stop
        start
}

reload () {
	echo -n $"Reloading $prog: "
	if ! /usr/sbin/php-fpm --test ; then
	        RETVAL=6
	        echo $"not reloading due to configuration syntax error"
	        failure $"not reloading $prog due to configuration syntax error"
	else
		killproc -p ${pidfile} php-fpm -USR2
		RETVAL=$?
	fi
	echo
}


# 可用方法
case "$1" in
  start)
	start
	;;
  stop)
	stop
	;;
  status)
	status -p ${pidfile} php-fpm
	RETVAL=$?
	;;
  restart)
	restart
	;;
  reload|force-reload)
	reload
	;;
  configtest)
 	/usr/sbin/php-fpm --test
	RETVAL=$?
	;;
  condrestart|try-restart)
	[ -f ${lockfile} ] && restart || :
	;;
  *)
	echo $"Usage: $0 {start|stop|status|restart|reload|force-reload|condrestart|try-restart|configtest}"
	RETVAL=2
        ;;
esac

exit $RETVAL
```

CentOS7环境下的脚本位置在`/usr/lib/systemd/system/php-fpm.service`, 其内容为

```
[Unit]
Description=The PHP FastCGI Process Manager                            
After=syslog.target network.target                                     

[Service]
Type=notify
PIDFile=/run/php-fpm/php-fpm.pid
EnvironmentFile=/etc/sysconfig/php-fpm
ExecStart=/usr/sbin/php-fpm daemonize
ExecReload=/bin/kill -USR2 $MAINPID
PrivateTmp=true

[Install]
WantedBy=multi-user.target
```

## 源码分析



## 变量定义

### `systemd`风格的变量

在`.service`文件中在`[service]`块中使用`Environment`字段可以定义单个变量, 格式如下

```
Environment=变量名=变量值
```

也可以使用`EnvironmentFile`字段指定一个变量定义文件, 然后把所以的变量写在这个文件里, 格式如下

```
EnvironmentFile=变量文件绝对路径
```

变量文件中的也是`变量名=变量值`的格式, 一行一条.

#### 1.

使用`Environment`定义的变量, **变量值中不可以有空格**, 我曾经尝试过使用单双引号将包含空格的变量值包裹起来, 但会报错(引用这个变量的时候才会).

```
## 正确
Environment=OPTS=--pid=/tmp/fpm.pid
## 错误, 不能有空格, 并且不能写多条, 长选项可以使用'='连接, 短选项可以合并到一起写
Environment=OPTS=--pid /tmp/fpm.pid
```

使用`EnvironmentFile`在变量定义文件中, 变量值可以有空格, 并且可以多个选项写在同一行, 如

```
OPTS=--pid /tmp/hehe.pid --daemonize
```

则`.service`文件中可以这样使用

```
ExecStart=/usr/local/php/sbin/php-fpm $OPTS
```

#### 2.

引用一个不存在的变量不会报错, 但引用一个不存在的变量配置文件会报`Failed to load environment files: No such file or directory`. 有时不确定是不是存在这样的变量配置文件, 需要在`EnvironmentFile`的值前面加上`-`中划线**抑制错误**, 就算不存在这个文件也不会出错停止.

```
EnvironmentFile=-变量文件绝对路径
```

#### 3.

`.service`文件中类似于`ExecStart`的字段(其实应该是大多数字段)第1个参数应该是一个可执行文件的绝对路径, **必须是绝对路径, 存在于环境变量也不行, 并且不能包含变量**.

如果在环境变量文件中有如下定义

```
PREFIX=/usr/local/php
```

而在`.service`文件中希望这样使用

```
ExecStart=$PREFIX/sbin/php-fpm --nodaemonize
```

是不可以的, `systemd`无法解析出这个变量, 将会启动失败.

另外, 环境变量配置文件中也无法引用其他变量, 例如

```
PID=/tmp/php-fpm.pid
OPTS=--daemonize --pid $PID
```

在`.service`文件打算这样使用

```
ExecStart=/usr/local/php/sbin/php-fpm $OPTS
```

并不能达到预期的效果, `systemd`没有办法解析`$OPTS`, 它将把'$OPTS'字符串当作PID文件名, 然后在默认路径下创建名为'$OPTS'的PID文件...

所以在变量配置文件中也不能随意进行变量引用, 还不如不用, 只能写全路径.