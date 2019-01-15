# Python-postgres连接库psycopg2(二)-进阶使用

1. [psycopg2笔记](https://www.cnblogs.com/hao-ming/p/7215050.html)

## 1. 动态sql-传入变量

示例1(以元组形式传入变量)

```py
sql_str_select = 'select id from books where info_addr = %s'
db_curr.execute(sql_str_select, (url, ))
```

注意:

1. `execute()`中传入变量是用逗号分隔, 用`%`会报sql语句中的单引号错误(不过作用和`%`一样)

2. 第二个参数是元组类型, 且**以逗号结尾**, 否则会报`TypeError: not all arguments converted during string formatting`的错误.

示例2(以对象形式传入变量)

```py
book_info = {
    'name': '斗罗大陆',
    'key': 'douluodalu',
    'author': '唐家三少',
    'cover': 'https://www.baidu.com',
    'summary': '有一天',
    'category_id': 1,
    'status': 1,
    'info_addr': 'https://www.baidu.com',
}
fields = 'name, key, author, summary, book_category_id, status, cover_addr, info_addr'
values = '%(name)s, %(key)s, %(author)s, %(summary)s, %(category_id)s, %(status)s, %(cover_addr)s, %(info_addr)s'
sql_str_insert = 'insert into books(%s) values(%s)' % (fields, values)
db_curr.execute(sql_str_insert, book_info)
```

注意:

1. 在`execute()`中传入的变量必须全是`%s`, 就算是数值也一样(比如上面的`category_id`).

2. `sql_str_insert`把`fields`和`values`的值填入, 但`%(name)s`这些仍然有百分号, 这些是要填入`book_info`的字段.

## 2. 插入/更新操作返回值

在使用orm时我们可以定义一个Model对象, 将其存入数据库后, orm会自动把得到的id等值将这个对象补全. 但是使用`psycopg2`执行原生sql如何实现这种操作呢? 

原生sql提供了`returning`子句, 可以在执行`insert`或`update`操作时依然可以得到匹配行中字段.

```sql
kanjula=# select * from device_msgs;
 device_id |       msg        |     created_at
-----------+------------------+---------------------
         1 | hi               | 2018-12-20 00:00:00
(5 rows)

kanjula=# insert into device_msgs(device_id, msg) values(5, 'perfect') returning device_id, msg;
 device_id |   msg
-----------+---------
         5 | perfect
(1 row)

INSERT 0 1
kanjula=# select * from device_msgs;
 device_id |       msg        |     created_at
-----------+------------------+---------------------
         1 | hi               | 2018-12-20 00:00:00
         5 | perfect          |
(6 rows)

kanjula=# update device_msgs set created_at = '2019-01-14' where device_id = 5 returning msg;
   msg
---------
 perfect
(1 row)

UPDATE 1
```

同样, 在`psycopg2`中可以使用如下方法.

```py
fields = 'name, key, author, summary, book_category_id, status, cover_addr, info_addr'
values = '%(name)s, %(key)s, %(author)s, %(summary)s, %(category_id)s, %(status)s, %(cover_addr)s, %(info_addr)s'
sql_str_insert = 'insert into books(%s) values(%s) returning id' % (fields, values)
db_curr.execute(sql_str_insert, book_info)

row = db_curr.fetchone()
self.db.commit()
```

我们在`insert`语句中写了`returning`语句, 在执行`execute()`后, **`commit()`之前**, 可以通过`fetchone()`直接获取到`returning`返回的字段.