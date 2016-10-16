# Redis 应用场景

## 1. 多数据库

redis可以开启多个实例监听不同端口以便对不同程序服务, 有没有什么方法使不同的应用程序数据彼此分开同时又存储在相同的实例上呢? 毕竟mysql通常只是启动`3306`端口, 就可以服务多个应用程序了.

redis下, 数据库是由一个整数索引标识, 而不是一个数据库名称. 默认情况下, 一个客户端连接到数据库0. redis配置文件`redis.conf`中下面的参数来控制数据库总数:

```
databases 16
```

`redis-cli`下可以通过下面的命令来切换到不同的数据库

```
redis 127.0.0.1:6379> select 2
OK
redis 127.0.0.1:6379[2]>
```

随后, 所有的命令将使用数据库2, 直到你明确的切换到另一个数据库下.

每个数据库都有属于自己的空间, 不必担心之间的key冲突. 不同的数据库下, 相同的key取到各自的值. flushdb命令清除数据, 只会清除当前的数据库下的数据, 不会影响到其他数据库. flushall命令会清除这个实例的数据. 在执行这个命令前要格外小心.

### 1.1 各个数据库的大小

参考

[redis分好库之后怎么才能看每个库的大小呢？](https://segmentfault.com/q/1010000000665987)

redis 貌似没有提供一个可靠的方法获得每个 db 的实际占用，这主要是因为 redis 本身就没有 db 文件概念，所有 db 都是混在一个 rdb 文件里面的。

要想估算 db 的大小，需要通过 keys * 遍历 db 里所有的 key，然后用 debug object <key> 来获得 key 的内存占用，serializedlength 就是占用内存的字段长度。

根据 RDB 格式文档，可以估算出每个 key 的实际占用为：

```
key_size = strlen(key) + serializedlength + 7
```

不过这个估算极不靠谱，因为 redis 可能将 key 做压缩，此时估算出来的值可能偏大。

下面的命令可以查看 db0 的大小（key 个数），其他的以此类推。类似于mysql中`select count(*) from 表名`, 不过使用`keys *`也可以得到所有的键, 并且根据序号排列.

```
127.0.0.1:6379> select 0
OK
127.0.0.1:6379> dbsize
(integer) 473
```

或者使用 `info keyspace` 同时得到所有 db 信息。

```
127.0.0.1:6379> info keyspace
# Keyspace
db0:keys=473,expires=0,avg_ttl=0
db1:keys=3911,expires=3909,avg_ttl=0
```

## 2. 访问密码

参考文章

[redis配置认证密码](http://blog.csdn.net/zyz511919766/article/details/42268219)

### 2.1 通过配置文件进行配置

redis默认是没有密码的, 其密码设置在其配置文件中`requirepass`字段, 默认是被注释的.

```
# requirepass foobared
```

解开注释, 并将上面的'foobared'改成你自己的密码, 然后重启redis.

```
requirepass 123456
```

这个时候再次连接redis，发现可以连接上，但是执行具体命令时提示操作不允许.

```
redis-cli -h 127.0.0.1 -p 6379  
redis 127.0.0.1:6379>  
redis 127.0.0.1:6379> keys *  
(error) ERR operation not permitted  
redis 127.0.0.1:6379> select 1  
(error) ERR operation not permitted  
redis 127.0.0.1:6379[1]>
```

而尝试用密码登录并执行具体的命令, 可以看到命令成功执行.

```
redis-cli -h 127.0.0.1 -p 6379 -a 123456
redis 127.0.0.1:6379> keys *
(empty list or set)
redis 127.0.0.1:6379[1]> config get requirepass  
1) "requirepass"
2) "123456"
```

### 2.1 通过命令行进行配置

```
redis 127.0.0.1:6379> config set requirepass 654321  
OK  
redis 127.0.0.1:6379> config get requirepass  
1) "requirepass"  
2) "654321"
```

无需重启redis.

使用第一步中配置文件中配置的老密码登录redis，会发现原来的密码已不可用，也时显示操作被拒绝.

```
redis-cli -h 127.0.0.1 -p 6379 -a 123456  
redis 127.0.0.1:6379> config get requirepass  
(error) ERR operation not permitted
```

使用刚才通过命令行设置的密码登陆, 可以执行响应操作.

```
redis-cli -h 127.0.0.1 -p 6379 -a 654321  
redis 127.0.0.1:6379> config get requirepass  
1) "requirepass"  
2) "654321"
```

尝试重启一下redis，用新配置的密码登录redis执行操作，发现新的密码失效，redis重新使用了配置文件中的密码.

-----


除了在登录时通过 -a 参数制定密码外，还可以登录时不指定密码，而在执行操作前进行认证。

```
redis-cli -h 127.0.0.1 -p 6379  
redis 127.0.0.1:6379> config get requirepass  
(error) ERR operation not permitted  
redis 127.0.0.1:6379> auth 123456  
OK  
redis 127.0.0.1:6379> config get requirepass  
1) "requirepass"  
2) "123456"
```

### 2.3 集群模式下的密码配置

若master配置了密码, 则slave也要配置相应的密码参数否则无法进行正常复制。 slave中配置文件内找到如下行，移除注释，修改密码即可

```
#masterauth  master的密码
```
