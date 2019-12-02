# SQL情境分析之分组后按照分组规则合并字符串

参考文章

1. [sql： 分组后按照分组规则拼接字符串 -- group by与 group_concat()](https://blog.csdn.net/weixin_42845682/article/details/81317392)

2. [Postgresql GROUP_CONCAT equivalent?](https://stackoverflow.com/questions/2560946/postgresql-group-concat-equivalent/53089719)

有如下表, 表示学生的科目和得分情况, 每行记录表示一个科目.

```sql
create table student(
    id int,
    name varchar(50),
    class varchar(50),
    score int
);
```

有如下数据

```sql
insert into student values(1,'张三','高数',32);
insert into student values(1,'张三','大学物理',27);
insert into student values(1,'张三','计量经济学',38);
insert into student values(2,'李四','高数',29);
insert into student values(2,'李四','计量经济学',41);
insert into student values(3,'王五','高数',34);
insert into student values(3,'王五','大学物理',33);
```

得到如下内容

```sql
 select * from student;
+------+--------+-----------------+-------+
| id   | name   | class           | score |
+------+--------+-----------------+-------+
|    1 | 张三   | 高数            |    32 |
|    1 | 张三   | 大学物理        |    27 |
|    1 | 张三   | 计量经济学      |    38 |
|    2 | 李四   | 高数            |    29 |
|    2 | 李四   | 计量经济学      |    41 |
|    3 | 王五   | 高数            |    34 |
|    3 | 王五   | 大学物理        |    33 |
+------+--------+-----------------+-------+
```

我们希望按照学生姓名分组, 得到这个学生所有科目的总分, 并且科目名称要按照指定规则拼接在一起, 用单行记录表示.

## 1. mysql的实现方式

启动容器

```
docker run -d --name mysql -e LANG=C.UTF-8 -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=mydb -e MYSQL_USER=mydb -e MYSQL_PASSWORD=123456 -p 3306:3306 mysql
```

进入容器

```
docker exec -it mysql mysql -u root -p
Enter password: ## 输入123456
```

选择数据库, 创建表, 插入数据后进行如下查询.

```sql
mysql> SELECT id, name, group_concat(class separator '-') class, sum(score) score FROM student GROUP BY id, name;
+------+--------+-------------------------------------+-------+
| id   | name   | class                               | score |
+------+--------+-------------------------------------+-------+
|    1 | 张三   | 高数-大学物理-计量经济学            |    97 |
|    2 | 李四   | 高数-计量经济学                     |    70 |
|    3 | 王五   | 高数-大学物理                       |    67 |
+------+--------+-------------------------------------+-------+
3 rows in set (0.00 sec)

```

方便.

## 2. postgresql的实现方式

```sql
skycmdb=# select id, name, string_agg(class, '-'), sum(student.score) from student group by id, name;
 id | name |        string_agg        | sum
----+------+--------------------------+-----
  3 | 王五 | 高数-大学物理            |  67
  1 | 张三 | 高数-大学物理-计量经济学 |  97
  2 | 李四 | 高数-计量经济学          |  70
(3 rows)
```

> 注意: `sum`聚合函数要加上`student.`作为前缀.

------

以下是postgres实现方式的扩展.

有一个`array_agg()`函数会把聚合后的字段变成数组类型.

```sql
skycmdb=#  select id, name, array_agg(class), sum(student.score) from student group by id, name;
 id | name |         array_agg          | sum
----+------+----------------------------+-----
  3 | 王五 | {高数,大学物理}            |  67
  1 | 张三 | {高数,大学物理,计量经济学} |  97
  2 | 李四 | {高数,计量经济学}          |  70
(3 rows)
```

可以使用`::text`将其转换为字符串, 但是这样就没办法自定义分隔符了.

```sql
skycmdb=#  select id, name, array_agg(class)::text, sum(student.score) from student group by id, name;
 id | name |         array_agg          | sum
----+------+----------------------------+-----
  3 | 王五 | {高数,大学物理}            |  67
  1 | 张三 | {高数,大学物理,计量经济学} |  97
  2 | 李四 | {高数,计量经济学}          |  70
(3 rows)
```

不过正好有一个`array_to_string()`函数, 类似程序语言中数组对象的`join`方法, 用于将数组中和成员拼接成字符串.

```sql
skycmdb=#  select id, name, array_to_string(array_agg(class), '-'), sum(student.score) from student group by id, name;
 id | name |     array_to_string      | sum
----+------+--------------------------+-----
  3 | 王五 | 高数-大学物理            |  67
  1 | 张三 | 高数-大学物理-计量经济学 |  97
  2 | 李四 | 高数-计量经济学          |  70
(3 rows)
```

完美.