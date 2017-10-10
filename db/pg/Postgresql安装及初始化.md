# Postgresql安装及初始化

与mysql一样, 使用yum安装的pg也有客户端与服务端两个包, 分别是`postgresql`与`postgresql-server`.

初次安装, 需要进行初始化.

切换到`postgres`用户, 其home目录默认在`/var/lib/pgsql`. 执行`initdb`.

```
$ initdb -D data
```

`-D`参数指定数据库文件存放路径, 默认在`/var/lib/pgsql/data`. 这一步是必须的.

然后启动postgresql服务.

```
$ pg_ctl -l ./logs/pg.log start
```

start子命令是一个前端进程, 日志会在终端直接输出, 你需要使用`-l`指定一个日志文件, 就可以以服务的形式运行postgre了.