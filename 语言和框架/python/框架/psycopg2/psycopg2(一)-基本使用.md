# psycopg2(一)-基本使用

参考文章

1. [psycopg2笔记](https://www.cnblogs.com/hao-ming/p/7215050.html)

2. [cursor() — 数据库连接操作 python](https://www.cnblogs.com/qixu/p/6133429.html)

    - 这里介绍的是mysql的python连接库, 但是用法和`psycopg2`几乎相同.

3. [数据库的Connection、Cursor两大对象](https://www.cnblogs.com/zhouziyuan/p/10155612.html)

## 1. 连接方式

```py
import psycopg2

db_config = {
    'host': 'localhost',
    'port': '5432',
    'database': 'wuhoudb',
    'user': 'wuhou',
    'password': '123456',
}

db_conn = psycopg2.connect(**db_config)
## 关闭事务自动提交(事务在执行第一次execute时自动开启, 手动调用db_conn.commit()提交)
db_conn.set_isolation_level(1)
db_curr = db_conn.cursor()

db_curr.execute('select id, name from books limit 2')
row = db_curr.fetchall()
print(row) ## [(1004, '大主宰'), (1005, '东山再起[娱乐圈]')]

db_curr.close()
db_conn.close()
```

`connect()`: 返回数据库连接对象. 数据库连接对象可以用来开启事务, 提交或是回滚.

`cursor()`: 从连接对象获取游标对象, 游标对象可以用来执行sql语句. 关于`cursor`对象的解释, 可以见参考文章3, 很形象.

> `set_isolation_level(1)`: 关闭自动提交

## 2. 执行sql

### 2.1 select查询

```py
sql_str = 'select id, name from books limit 2'

db_curr.execute(sql_str)
row = db_curr.fetchall()
print(row) ## [(1004, '大主宰'), (1005, '东山再起[娱乐圈]')]

db_curr.execute(sql_str)
row = db_curr.fetchone()
print(row) ## (1004, '大主宰')

db_curr.execute(sql_str)
row = db_curr.fetchmany(size = 1)
print(row) ## [(1004, '大主宰')]
```

首先需要调用`execute()`执行sql语句, 然后需要通过`fetchXXX`方法来取得执行结果. 

可用的`fetch`方法有: 

1. `fetchone`: 返回元组, 元组成员顺序为sql语句中`select xxx, yyy`的顺序. 空结果为`None`

2. `fetchall`: 返回列表, 成员为元组. 

3. `fetchmany(size = None)`: `size`为数值, 可以从sql执行结果中返回`size`个行, 默认按1处理, 空结果为`[]`. `size`的作用类似于socket编程中`recv(buf)`中的buf差不多, 先规定一个预期值, 多了不会超, 少了要看实际值. 

需要注意的是, `execute()`的执行结果只能被1个`fetchXXX()`取出, 之后再接的`fetchXXX()`是不会有结果的, 所以上述示例中写了3个`execute()`.

### 2.2 insert新增

