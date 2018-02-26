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

