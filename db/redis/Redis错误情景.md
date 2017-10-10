# redis 错误情景

## 1.

参考文件

[强制关闭Redis快照导致不能持久化](http://www.cnblogs.com/anny-1980/p/4582674.html)

### 问题描述

Redis被配置为保存数据库快照，但它目前不能持久化到硬盘。用来修改集合数据的命令不能用。请查看Redis日志的详细错误信息。

```
(error) MISCONF Redis is configured to save RDB snapshots, but is currently not able to persist on disk. Commands that may modify the data set are disabled. Please check Redis logs for details about the error.
```

### 原因分析

强制关闭Redis快照导致不能持久化。

### 解决方法

`redis-cli`下运行`config set stop-writes-on-bgsave-error no`命令后，关闭配置文件中`stop-writes-on-bgsave-error`项解决该问题。

```
$ /usr/local/redis/src/redis-cli
127.0.0.1:6379> config set stop-writes-on-bgsave-error no
OK
127.0.0.1:6379> lpush myColour "red"
(integer) 1
```

## 2.

### 问题描述

redis安装完成, 但启动报错, redis版本: 3.0.7

```
redis-server /usr/local/etc/redis.conf
*** FATAL CONFIG FILE ERROR ***
Reading the configuration file, at line 54
>>> 'tcp-backlog 511'
Bad directive or wrong number of arguments```
```

### 原因分析

redis的配置文件与安装的redis程序不是同一个版本, 有可能是之前安装过redis, 此次启动读取的是之前的配置文件. 顺便可以查看时不是已经有一个redis进程在运行了.

### 解决办法

记得停止正在运行的redis进程, 删除可执行文件, 配置文件等内容, 然后再次尝试启动新的redis.

## 3. 主从设置不生效

参考文章

[Redis Slave Master connections fails Slave logs show: Unable to connect to MASTER: Permission denied](http://stackoverflow.com/questions/34906127/redis-slave-master-connections-fails-slave-logs-show-unable-to-connect-to-maste)

[Redis Slave Master connections fails Slave logs show: Unable to connect to MASTER: Permission denied](http://stackoverflow.com/questions/34906127/redis-slave-master-connections-fails-slave-logs-show-unable-to-connect-to-maste)

### 问题描述

在`cli`中执行`slaveof 172.16.1.100 6379`, 希望其将`172.16.1.100`中的数据同步到本地. 但是不生效, 而且再执行`set key1 'val1'`, 竟然返回`OK`. 要知道作为从节点的服务是不可写的.

查看日志, 显示

```
[20671] 12 Jan 15:48:02.369 * Connecting to MASTER 172.16.1.100:6379 [20671] 12 Jan 15:48:02.369 # Unable to connect to MASTER: Permission denied
```

### 原因分析及解决

SELinux的问题, 需要将远程redis所再主机连同本地主机的SELinux都关闭.
