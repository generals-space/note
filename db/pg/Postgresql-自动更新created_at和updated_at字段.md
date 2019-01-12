# Postgresql-自动更新created_at和updated_at字段

参考文章

1. [PostgreSQL之时间戳自动更新](https://www.cnblogs.com/MikeZhang/p/PostgreSQLRealte20171013.html)

表中拥有`created_at`与`updated_at`两个字段. 

目标: 

1. 插入新记录时自动为`created_at`赋值为当前时间(`updated_at`也是).

2. 在更新某条记录时, 自动更新`updated_at`为当时的时间.

在学校用php+mysql时, 通过简单的建表语句就可以实现. 不过现在用上了orm, 总是想把建表这种事交给orm来做.

本来是想用gorm通过`gorm`标记来做的, 在网上查到可以写成下面这种格式

```go
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
```

但是`created_at`倒是可以生效, `updated_at`字段却没什么用.

后来找到了参考文章1, 发现mysql与postgres在这方面的实现方法是不同的, 如果用的是mysql估计就成了, 但是...

好了, 首先在mysql中要达到我们的目的, sql语句应该写成

```sql
create table tbl_a (
    id int,
    created_at timestamp NOT NULL default CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP
);
```

------

不过在postgres中只能通过触发器实现.

首先创建一个函数`upd_timestamp`.

```sql
create or replace function upd_timestamp() returns trigger as
$$
begin
    new.updated_at = current_timestamp;
    return new;
end
$$
language plpgsql;
```

然后为`books`表创建触发器对象`auto_update_date_for_books`(有点长, 其实名字随便写), 调用`upd_timestamp`函数

```sql
create trigger auto_update_date_for_books before update on books for each row execute procedure upd_timestamp();
```

然后在用`\d`查看`books`表的详细信息时, 就可以看到表中存在的触发器了.

```sql
\d books
                                              Table "public.books"
          Column           |           Type           | Collation | Nullable |              Default
---------------------------+--------------------------+-----------+----------+-----------------------------------
 id                        | bigint                   |           | not null | nextval('books_id_seq'::regclass)
 created_at                | timestamp with time zone |           |          | CURRENT_TIMESTAMP
 updated_at                | timestamp with time zone |           |          | CURRENT_TIMESTAMP
Indexes:
    "books_pkey" PRIMARY KEY, btree (id)
Triggers:
    auto_update_date_for_books BEFORE UPDATE ON books FOR EACH ROW EXECUTE PROCEDURE upd_timestamp()
```