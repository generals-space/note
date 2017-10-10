# PHP语言级别问题处理

## 1. PHP连接MySQL报错＂No such file or directory＂

参考文章

1. [mysql_connect报告”No such file or directory”错误的解决方法](http://www.cnblogs.com/AloneSword/p/4137730.html)

2. [Mac下PHP连接MySQL报错＂No such file or directory＂的解决办法](http://www.linuxidc.com/Linux/2012-12/76150.htm)

在LNMP环境中, php测试数据库连接的脚本一直显示连接失败. 但用户名密码绝对是正确的, nginx日志中显示如下

```
"PHP message: PHP Warning: mysql_connect(): No such file or directory in /usr/local/nginx/html/test.php on line 6"
```

脚本内容为

```php
<?php 
/**
* 测试php与mysql连接
* 编辑：www.jb51.net
*/
$link=mysql_connect('localhost','root','83337879'); 
if(!$link) 
    echo "FAILD!连接错误，用户名密码不对"; 
else 
    echo "OK!可以连接"; 
?>
```

直接使用`php`命令执行脚本, 得到的也是`No such file or directory`.

------

原因分析

可能是由于mysql安装选项的差异, 导致php无法找到mysql的本地socket文件.

解决方法, 进入mysql命令行(注意: 必须是以root身份), 执行`status`命令, 如下

```
mysql> status
--------------
mysql  Ver 14.14 Distrib 5.6.31, for Linux (x86_64) using  EditLine wrapper

Connection id:		104438
Current database:	
Current user:		root@localhost
SSL:			Not in use
Current pager:		stdout
Using outfile:		''
Using delimiter:	;
Server version:		5.6.31-log MySQL Community Server (GPL)
Protocol version:	10
Connection:		Localhost via UNIX socket
Server characterset:	utf8
Db     characterset:	utf8
Client characterset:	utf8
Conn.  characterset:	utf8
UNIX socket:		/opt/mysqldata/mysql.sock
Uptime:			137 days 20 hours 53 min 32 sec

Threads: 2  Questions: 3105297  Slow queries: 939757  Opens: 447  Flush tables: 1  Open tables: 203  Queries per second avg: 0.260
--------------

```

查看其中`UNIX socket`值, 这是一个不太常见的路径, 一般应该是`/tmp/mysql.sock`或是在`/var/lib/mysql`中.

打开`php.ini`, 找到其中`mysql.default_socket`、`mysqli.default_socket`、`pdo_mysql.default_socket`三个字段, 将其值修改为上述`UNIX socket`的路径. 如下

```ini
mysql.default_socket=/opt/mysqldata/mysql.sock
mysqli.default_socket=/opt/mysqldata/mysql.sock
pdo_mysql.default_socket=/opt/mysqldata/mysql.sock
```

最后, **如果是`LNMP`要重启`php-fpm`容器, 如果是`LAMP`则要重启`apache`, 否则就算直接使用`php`命令也无法正常连接mysql**.