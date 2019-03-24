# Postgresql建表语句 - 外键约束

参考文章

1. [PostgreSQL 9.4官方文档5.3. Constraints](https://www.postgresql.org/docs/9.4/static/ddl-constraints.html)

2. [PostgreSQL外键约束reference](http://blog.csdn.net/wujiang88/article/details/51578794)

我们希望创建两张表, `user`与`department`, 其中`user`表的`department_id`字段表示用户所属的部门的id, 即外键.

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

建立外键关系需要先创建被引用表, 再创建主引用表, 否则会报错, 比如先创建user表时, 报错如下.

```sql
create table "user"(id serial primary key, department_id integer not null, name varchar(20), foreign key (id) references department(id));
ERROR:  relation "department" does not exist
```

所以要先创建`department`表.

```sql
create table department(id serial primary key, department varchar(50));
CREATE TABLE
create table "user"(id serial, department_id integer not null, name varchar(20), foreign key (id) references department(id));
CREATE TABLE
```

------

**列约束与表约束**

可以看到, 我们将`foreign key`的创建放在与其他列字段同等级的水平, 这种创建[外键]约束的方法叫作`表约束`.

还有一种方法叫做`列约束`, 是把`foreign key`与上面的`primary key`同等级的水平, 即作为列字段的一个修饰.

```sql
create table department(id serial primary key, department varchar(50));
CREATE TABLE
create table "user"(id serial primary key, department_id integer references department(id), name varchar(20));
CREATE TABLE
```

...就是移除了`foreign key`而已, 但是感觉比`表约束`少了一个显式的外键字段, 更清晰了.

但是, 约束只能定义单一字段, 而表约束则可以组合多个列. 也就是说多个外键字段时, 列约束要写很多个, 而表约束则可以`references 被引用表(被引用字段1, 被引用字段2...)`

> pg里默认引用products表的主键字段，但最好不要这样用, 要指定被引用字段的名称.

```sql
CREATE TABLE IF NOT EXISTS "user"
(
    id        SERIAL PRIMARY KEY,
	name        varchar(50) not null unique
);

CREATE TABLE IF NOT EXISTS book
(
	id        SERIAL PRIMARY KEY,
	name        varchar(50) not null unique,
	user_id int references "user"(id) not null
);

CREATE TABLE IF NOT EXISTS book
(
	id        SERIAL PRIMARY KEY,
	name        varchar(50) not null unique,
	user_id int not null,
    CONSTRAINT book_user_id_fkey FOREIGN KEY (user_id) REFERENCES "user"(id)
);
```

...实际上根本不存在什么表约束还是列约束, 因为这两种方式创建的表是完全一样, 只是写法上不一样而已, 误导人了.