# Postgresql查看正在运行的sql及处理方法

参考文章

1. [postgreSQL查询正在运行的sql和kill杀掉的处理方法](http://blog.sina.com.cn/s/blog_e964c8bd0102w1sc.html)

查看正在运行的sql, 语句为`SELECT * FROM pg_stat_activity;`. 结果如下

```
sky=> SELECT * FROM pg_stat_activity;
 datid | datname |  pid  | usesysid | usename | application_name | client_addr | client_hostname | client_port |         backend_start         |          xact_start           |          query_start          |         state_change          | waiting | state  | backend_xid
 | backend_xmin |              query               
-------+---------+-------+----------+---------+------------------+-------------+-----------------+-------------+-------------------------------+-------------------------------+-------------------------------+-------------------------------+---------+--------+------------
-+--------------+----------------------------------
 16386 | sky     | 30882 |    16384 | sky     |                  | 172.16.3.50 |                 |       58763 | 2017-12-08 14:15:08.044085+08 |                               | 2017-12-08 14:15:24.901902+08 | 2017-12-08 14:15:24.902068+08 | f       | idle   |            
 |              | ROLLBACK
 16386 | sky     | 30651 |    16384 | sky     | psql             |             |                 |          -1 | 2017-12-08 14:11:55.59858+08  | 2017-12-08 14:18:57.166021+08 | 2017-12-08 14:18:57.166021+08 | 2017-12-08 14:18:57.166026+08 | f       | active |            
 |       220915 | SELECT * FROM pg_stat_activity;
(2 rows)
```

各字段含义:

`datname`: 正在执行的sql的目标库名.

`pid`: sql表示的进程id.

`usename`: 执行者(数据库用户名)