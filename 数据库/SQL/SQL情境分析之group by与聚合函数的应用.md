# SQL情境分析之group by与聚合函数的应用

参考文章

1. [GROUP BY与COUNT用法详解](https://blog.csdn.net/liu_shi_jun/article/details/51329472)

2. [SQL 连接(JOIN)](http://www.runoob.com/sql/sql-join.html)

3. [SQL QUERY using LEFT JOIN and CASE Statement](https://stackoverflow.com/questions/14732500/sql-query-using-left-join-and-case-statement)

必须要用到group by的场景是, 有两张表`table_a`, `table_b`, 其中表b中有一个列级外键字段, 关联表的a主键id.

```sql
CREATE TABLE table_a
(
	id        SERIAL PRIMARY KEY,
	name        varchar(50) not null unique
);

CREATE TABLE table_b
(
	id        SERIAL PRIMARY KEY,
	name        varchar(50) not null unique,
	foreign_id int references table_a(id) not null
);

insert into table_a(name) values('组1'), ('组2'), ('组3'), ('组4');
insert into table_b(name, foreign_id) values
('成员1', 1),
('成员2', 1),
('成员3', 4),
('成员4', 2),
('成员5', 2),
('成员6', 2)
;
```

我的目标是查询所有表a中的记录, 并且为每条记录都添加一个字段, 表示在表b中以此条记录为外键的行的总量, 最终结果类似如下(主要是`count`列);

```
 id | name | foreign_id | count
----+------+------------+-------
  1 | 组1  |          1 |     2
  2 | 组2  |          2 |     3
  3 | 组3  |            |     0
  4 | 组4  |          4 |     1
(4 rows)
```

我们想得到所有的分组, 及每个分组的成员人数. 为了达到我们的目的需要使用`group by`语句, `group by`必须与**聚合函数**配合使用.

> 聚合函数， 例如`SUM`, `COUNT`, `MAX`, `AVG`等, 这些函数和其它函数的根本区别就是它们一般作用在**多条记录**上.

举个最简单的例子. 

```sql
/*group不能这么使用*/
select * from table_b group by foreign_id;
ERROR:  column "table_b.id" must appear in the GROUP BY clause or be used in an aggregate function
LINE 1: select * from table_b group by foreign_id;
               ^
select count(*) from table_b group by foreign_id;
 count
----
  3
  2
  1
(3 rows)
```

由于使用了`group by`, 那我们得到的必然不会是全部的记录, 而是不同组的记录的总体信息(或者说共性). 在同一分组中, 被group by的字段必然是相同的, `select`只能查询这些共性字段, 其他任何一个单独的字段都是不可能的.

好了, 我们先按照常规的连接查询

```sql
select table_a.*, table_b.foreign_id from table_a, table_b where table_a.id = table_b.foreign_id;
 id | name | foreign_id
----+------+------------
  1 | 组1  |          1
  1 | 组1  |          1
  4 | 组4  |          4
  2 | 组2  |          2
  2 | 组2  |          2
  2 | 组2  |          2
(6 rows)
```

我们用`group by`对上面的表进行一下加工.

```sql
select table_a.*, count(table_b.foreign_id) as sum from table_a, table_b where table_a.id = table_b.foreign_id group by table_b.foreign_id;
ERROR:  column "table_a.id" must appear in the GROUP BY clause or be used in an aggregate function
LINE 1: select table_a.*, count(table_b.foreign_id) as sum from table_a, ta...
               ^
```

直觉告诉我这种方式行不通, 因为我们无法在`group by`表b字段的同时还能把表a的所有字段拿出来. (然而我打脸了, 这个之后说, 先看看我自己摸索的思路)

## 1. 换一种思路

```sql
select foreign_id, count(*) as sum from table_b group by foreign_id;
 foreign_id | sum
------------+-----
          4 |   1
          2 |   3
          1 |   2
(3 rows)
```

把这个结果当作一张新表和表a做连接.

```sql
select * from table_a, (select foreign_id, count(*) as sum from table_b group by foreign_id) as table_c where table_a.id = table_c.foreign_id;
 id | name | foreign_id | sum
----+------+------------+-----
  1 | 组1  |          1 |   2
  2 | 组2  |          2 |   3
  4 | 组4  |          4 |   1
(3 rows)
```

好像对了...但是!!!!!!!!!!!!!!!!!!

表b中没有对应的外键指向`组3`, 所以结果中没有这一行的数量. 我希望至少能显示`sum`为0吧.

在我找方法时, 了解到了sql语句中的`join`种类: 

1. `左连接(left join)`

2. `右连接(right join)`

3. `内连接(inner join)`

4. `全连接(full join)`. 

其中`左连接`与`右连接`又被称为`外连接`, 而多表联合查询时默认join类型为`内连接`. 其中具体的区别可以查看参考文章2, 讲解得还算比较清晰.

我们更新下上面的sql

```sql
select * from table_a left join (select foreign_id, count(*) as sum from table_b group by foreign_id) as table_c on table_a.id = table_c.foreign_id;
 id | name | foreign_id | sum
----+------+------------+-----
  1 | 组1  |          1 |   2
  2 | 组2  |          2 |   3
  3 | 组3  |            |
  4 | 组4  |          4 |   1
(4 rows)
```

有门! 使用左连接时, 不管处于右侧的表中有没有与左侧的表相匹配(就是`on`子句中的条件), 左侧表中的字段都会出现, 右侧表中的字段显示为null.

但是我想让为null的行`sum`为0, 需要用到`case...when...`子句. 表连接 + `case when`的语法我参照了参考文章3, 需要再增加一个字段.

```sql
select *, case when table_c.count is null then 0 else table_c.count end as sum from table_a left join (select foreign_id, count(*) from table_b group by foreign_id) as table_c on table_a.id = table_c.foreign_id;
 id | name | foreign_id | count | sum
----+------+------------+-------+-----
  1 | 组1  |          1 |     2 |   2
  2 | 组2  |          2 |     3 |   3
  3 | 组3  |            |       |   0
  4 | 组4  |          4 |     1 |   1
(4 rows)
```

这里把count列作为了中间字段, 但也没有办法去除(除非再在外层套一个select).

> `case when`子句要放在`select`附近.

在我对自己的sql水平沾沾自喜的时, 之前求助过的dba终于回复我了, 随手丢过来一条sql来吊打...

## group by多个字段

我们在换思路之前试验过一条语句, 但是失败了, 原因是`无法在group by表b字段的同时还能把表a的所有字段拿出来`.

而dba的sql是这样的.

```sql
select table_a.id,table_a.name,table_b.foreign_id,count(table_b.foreign_id) from table_a left join table_b on table_a.id=table_b.foreign_id group by 1,2,3;
 id | name | foreign_id | count
----+------+------------+-------
  1 | 组1  |          1 |     2
  2 | 组2  |          2 |     3
  3 | 组3  |            |     0
  4 | 组4  |          4 |     1
(4 rows)
```

不用生成中间表`table_c`, 没有中间变量, 短很多(可能这样看不大出来, 但在为表取别名之后会短一大截)...dba看了我的sql表示不屑.

因为他给的sql是按多个列分组的. 语句等同如下

```sql
select table_a.*, count(table_b.foreign_id) from table_a left join table_b on table_a.id=table_b.foreign_id group by table_a.id, table_a.name, table_b.foreign_id;
```

因为仔细观察你会发现, 使用`id`, `name`, `foreign_id`3列同时分组与单纯使用`foreign_id`分组结果是相同的, ta们一一对应. 只要再用`count`计算`foreign_id`的值就可以了.

...这段sql花了我不少时间, 但也学到了很多东西. 值了!