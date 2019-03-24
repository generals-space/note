# SQL技巧 - 批量插入

参考文章

1. [SQL-批量插入和批量更新](http://blog.csdn.net/lovemenghaibin/article/details/50759003)

如果对sql语句不熟悉，大部分的插入语句是这么写的

```sql
INSERT INTO "user"(department_id, "name") VALUES(1, 'general');
INSERT INTO "user"(department_id, "name") VALUES(2, 'jiangming');
INSERT INTO "user"(department_id, "name") VALUES(2, 'lianqia');
```

这种语法的插入速度可是相当...慢的, 参考文章1中给出了一种比较"hack"的方法 - 使用`union all`语句.

```sql
insert into "user"(name)
select 'general'
union all select 'jiangming'
union all select 'lianqia';
```

参考文章1上说, 原始方法插入2000条数据要2分钟, `union all`方法只需要12秒.

示例

```sql
insert into "user"(department_id, name) 
select 1, 'general' 
union all select 2, 'jiangming' 
union all select 2, 'lianqia';
INSERT 0 3

select * from "user";
 id | department_id |   name    
----+---------------+-----------
  1 |             1 | general
  2 |             2 | jiangming
  3 |             2 | lianqia
(3 rows)
```

完美...