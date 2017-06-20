# Postgresql问题处理

## 1. 

```
[postgres@localhost data]$ psql
psql: could not connect to server: No such file or directory
	Is the server running locally and accepting
	connections on Unix domain socket "/opt/data/.s.PGSQL.1921"?
```

问题描述:

源码安装postgresql-9.5.4, 安装目录为`/usr/local/pgsql`, 数据目录在`/opt/data`. 使用`postgres`用户启动pg, 之后使用`psql`命令进入pg命令行时, 出现上述错误.

原因分析:

pg在启动时会生成一个`sock`文件, 用做客户端与数据库沟通的途径, 而上面说找不到`/opt/data/.s.PGSQL.1921`文件, 而实际上`/opt/data`目录下确实没有这个文件. 查看pg的配置文件, 发现有名为`unix_socket_directories`的字段, 其默认值为`/tmp`. 查看`/tmp`下的确存在此文件. 很有可能是`sock`文件路径的问题.

解决方法:

修改pg配置文件中`unix_socket_directories`字段, 然后重启pg数据库即可.

或者更简单一点, 如果不想重启数据库, 为`/tmp/.s.PGSQL.1921`文件建立软链接至`/opt/data`目录下即可.