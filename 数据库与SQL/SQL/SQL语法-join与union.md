# SQL语法-join与union

## 1. join表连接查询

假设存在两个表`user`和`department`, `user`表有一个字段`department_id`与`department`的`id(主键)`作外键连接. 

`user`表

|id|department_id|name|
|:-:|:-:|:-:|
|1|1|userA|
|2|2|userB|
|3|2|userC|

`department`表

|id|department|
|:-:|:-:|
|1|财务|
|2|人力|

建表语句如下

```sql
create table department(id serial primary key, department varchar(50));
CREATE TABLE
create table "user"(id serial, department_id int8 not null, name varchar(20), foreign key (id) references department);
CREATE TABLE
```

我们希望在查询`user`时通过这个外键字段同时得到`department`表中的`department`字段.

|id|department|name|
|:-:|:-:|:-:|
|1|部门1|userA|
|2|部门2|userB|
|3|部门2|userC|

最基本的写法是

```sql
select "user".*, department.department from "user", department where "user".department_id = department.id;
```

> 如果不加`where`子句, 每查出一条user记录, 所会查出所有的department记录.

有一种简化写法, 用`as`子句给表名取个别名

```sql
select a.*, b.department from "user" as a, department as b where a.department_id = b.id;
```

这个`as`可以省略, 不知道是不是pg独有的特性.

```sql
select a.*, b.department from "user" a, department b where a.department_id = b.id;
```

这其实本质上就相当于`join`表连接查询.

```sql
select a.*, b.department from "user" a join department b on a.department_id = b.id;
```

可以看出, **join查询是根据当前表已经查出的行中某个字段去另一个表查询与之匹配的行的某些字段**.

> `左连接`与`右连接`又被称为`外连接`, 而多表联合查询时默认join类型为`内连接`.

## 3. union查询

OK, 有了上面的概念作基础, 我们再来看一下`union`查询.

```sql
select id, name from "user" union select id, department from department;
 id |   name    
----+-----------
  2 | jiangming
  2 | 人力
  1 | 财务
  1 | general
  3 | lianqia
(5 rows)
```

...so, union select语句就是**查出和上一个select字段个数相同的结果直接附加到第一个select的结果后面**, 也就是说前后的select查询的字段名可以不一样, 就是为了凑行数...**但是对应字段的类型要一样!!!**

比如`user.id`, `user.department_id`都是数值类型, 那么union查询中的两个字段也必须是数值类型, 不然...如下

```sql
select id, department_id from "user" union select id, department from department;
ERROR:  UNION types integer and character varying cannot be matched
LINE 1: ...ct id, department_id from "user" union select id, department...
                                                             ^
```

## 1. 扩展 - 表连接更新

现在我们想在`user`表中创建一个新列`department`, 然后把`user`表中每个实体的`department`赋值为其department_id与department表中id值相同行的department值...就是说, 不再需要`department`这个表了, 之后也不会再需要外键字段了.

```sql
update "user" set department = (select department from department where department.id = "user".department_id);
```

还有一种写法, 看不懂

```sql
update "user" set department = a.department from department a where a.id = "user".department_id;
```