# Postgresql建表语句 - 主键&自增序列

参考文章

1. [postgresql 9.4是取消了serial类型吗？](https://www.zhihu.com/question/31845424)

```sql
create table "user"(id serial primary key, name varchar(20));
skycmdb=> insert into "user"(name) values ('general');
INSERT 0 1
skycmdb=> insert into "user"(name) values ('jiangming');
INSERT 0 1
skycmdb=> select * from "user";
 id |   name    
----+-----------
  1 | general
  2 | jiangming
(2 rows)
```

## 对`serial`的认识

关于`serial`, serial 并不是psql的类型，只是一个宏.

```sql
create table 表名(列名 serial);
```

等价于

```sql
CREATE SEQUENCE 表名_列名_seq;
CREATE TABLE 表名 (
    列名 integer NOT NULL DEFAULT nextval('表名_列名_seq')
);
ALTER SEQUENCE 表名_列名_seq OWNED BY 表名.列名;
```

即, 实际上需要事先创建一个`SEQUENCE`序列对象, 然后将主键的默认值设置为这个序列的名称.

> 不过序列名应该是可以自定义的, `表名_列名_seq`这种命名格式只是为了区分序列对象是用于哪张表的.

示例

```sql
skycmdb=> create sequence department_id_seq;
CREATE SEQUENCE
skycmdb=> create table department(id int8 primary key default nextval('department_id_seq'), department varchar(50));
CREATE TABLE
skycmdb=> insert into department(department) values('财务');
INSERT 0 1
skycmdb=> insert into department(department) values('人力');
INSERT 0 1
skycmdb=> select * from department;
 id | department 
----+------------
  1 | 财务
  2 | 人力
(2 rows)
```

> 但是有些客户端没有`serial`这个宏, 所以根本没有办法在字段类型中选择`serial`, 我们只能先创建一个索引, 然后将它的默认值赋给自增字段.

## 对`sequence`对象进行更精确的定义.

```sql
CREATE SEQUENCE 表名_列名_seq increment by 1 minvalue 1 no maxvalue start with 1;
```

- `increment by`: 指定了自增的步长;

- `start with`: 指定了初始值;

- `minvalue`与`maxvalue`划分了计数区间;