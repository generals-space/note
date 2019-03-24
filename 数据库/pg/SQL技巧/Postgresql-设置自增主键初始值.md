# Postgresql应用技巧-设置自增主键初始值

参考文章

1. [postgresql自增字段初始值的设定](https://my.oschina.net/justdo/blog/125042)

假设有如下表, 主键为`id`, 类型为自增序列.

```sql
\d books
                  Table "public.books"
    Column    |      Type     | Collation | Nullable |              Default
--------------+---------------+-----------+----------+-----------------------------------
 id           | bigint        |           | not null | nextval('books_id_seq'::regclass)
```

我们想设置其初始值为1000, 可以执行如下语句.

```sql
select setval('testtable_id_seq', 1000, false);
```

貌似setval的第2, 3个参数还有各自的含义, 这里不再深究.