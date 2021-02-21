# where..in..子句

参考文章

1. [select from sqlite table where rowid in list using python sqlite3 — DB-API 2.0](https://stackoverflow.com/questions/5766230/select-from-sqlite-table-where-rowid-in-list-using-python-sqlite3-db-api-2-0)

对于字段类型为字符串的列, 我们可以使用如下语句

```py
device_list = ['192.168.31.1:5555', '192.168.31.2:5555']
sql_str = "delete from devices where device_id in (?)"
cursor.execute(sql_str, (','.join(device_list), ))
```

但是如果字段为数值类型, 那么用`','join()`方法就没有用了, 直接传列表对象更是不行

```py
ids = [1, 2]
## 这种方式是错的
sql_str = "delete from devices where device_id in (?)"
cursor.execute(sql_str, (ids, ))
```

按照参考文章1中所说, 没有直接的办法, 只能手动拼接 sql 语句, 不过拼接得比较巧妙.

```py
ids = [1, 2]
## 这种方式是错的
sql_str = "delete from devices where device_id in (%s)"
## 将 {list} 的内容拼接成 n 个问号
seq=','.join(['?'] * len(ids))
sql_str = sql_str % seq
## 注意这里, 要用*号解开列表
cursor.execute(sql_str, (*ids, ))
```
