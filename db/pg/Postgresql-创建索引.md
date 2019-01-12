# Postgresql-创建索引

创建索引最大的好处就是大数据量下, 根据指定字段的查询速度可以明显提高.

首先某个表的索引是该表的某个字段, 这样在用`select * from table_xxx where 索引字段 = ?`查询时, 速度会快很多.

```sql
# create index on books(info_addr);
CREATE INDEX
```

这样可以在`books`表上的`info_addr`列创建索引. 

完成后, 在`\d books`时可以看到有新的索引出现.

```sql
wuhoudb=# \d books
                                              Table "public.books"
          Column           |           Type           | Collation | Nullable |              Default              
---------------------------+--------------------------+-----------+----------+-----------------------------------
 id                        | bigint                   |           | not null | nextval('books_id_seq'::regclass)
 info_addr                 | character varying(500)   |           |          | 
Indexes:
    "books_pkey" PRIMARY KEY, btree (id)
    "books_info_addr_idx" btree (info_addr)
```

默认一张表中的主键字段就会创建索引.

需要注意的是, 上述创建索引的sql是阻塞的. 如果有其他插入/更新操作时(查询可以), 会被挂起, 直到索引创建完成.

不过postgres还提供了一个异步(并行)创建索引的参数`concurrently`, 不会阻塞插入/更新操作.

```sql
create index concurrently on books(info_addr);
```

------

创建索引操作可能花较长的时间, 视服务器配置和表记录数量而定, 创建操作是可以中途取消的, 不会有影响.

```sql
# create index on book_chapters(info_addr);
^CCancel request sent
ERROR:  canceling statement due to user request
```