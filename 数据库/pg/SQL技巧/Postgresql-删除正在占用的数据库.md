# Postgresql-删除正在占用的数据库

参考文章

1. [RDS for PostgreSQL 删除数据库时提示 There are 2 other sessions using the database.](https://help.aliyun.com/knowledge_detail/41763.html)

尝试删除数据库, 但是报错说有客户端(桌面客户端, 程序连接等)正在连接, 不能删除.

```sql
postgres=# drop database kanjula;
ERROR:  database "kanjula" is being accessed by other users
DETAIL:  There are 5 other sessions using the database.
```

使用如下语句强制删除连接.

```sql
postgres=# select pg_terminate_backend(pid) from (select pid from pg_stat_activity where datname = 'kanjula') as a;
 pg_terminate_backend
----------------------
 t
 t
 t
 t
 t
(5 rows)
```

之后就可以删除数据库了.

```sql
postgres=# drop database kanjula;
DROP DATABASE
```