# SQL情境分析之窗口函数

数据库: postgres

假设有`device_msgs`表, 存储的是不同设备发来的信息, 且同一设备可以存储多条信息.

```sql
# select * from device_msgs;
 device_id |       msg        |     created_at
-----------+------------------+---------------------
         1 | hi               | 2018-12-20 00:00:00
         2 | hello            | 2018-12-21 00:00:00
         1 | nice to meet you | 2018-12-21 00:00:00
         3 | good             | 2018-12-23 00:00:00
         3 | follow me        | 2018-12-24 00:00:00
(5 rows)
```

有一个这样的需求: 查询出表中所有设备最近的一条消息, 比如上表中一共有3个设备的消息, 第3条记录比第1条新, 第5条比第4条新, 结果应是2, 3, 5行.

首先建表语句如下:

```sql
create table device_msgs(device_id int, msg varchar, created_at timestamp);

insert into device_msgs(device_id, msg, created_at) values(1, 'hi', '2018-12-20');
insert into device_msgs(device_id, msg, created_at) values(2, 'hello', '2018-12-21');
insert into device_msgs(device_id, msg, created_at) values(1, 'nice to meet you', '2018-12-21');
insert into device_msgs(device_id, msg, created_at) values(3, 'good', '2018-12-23');
insert into device_msgs(device_id, msg, created_at) values(3, 'follow me', '2018-12-24');
```

## 1. 简单联合查询

首先说最终结果

```sql
select table_a.* from device_msgs as table_a inner join (select device_id, max(created_at) as created_at from device_msgs group by device_id) as table_b on table_a.device_id = table_b.device_id and table_a.created_at = table_b.created_at;
 device_id |       msg        |     created_at
-----------+------------------+---------------------
         2 | hello            | 2018-12-21 00:00:00
         1 | nice to meet you | 2018-12-21 00:00:00
         3 | follow me        | 2018-12-24 00:00:00
(3 rows)
```

首先是使用`group by`按照不同设备分组查询的子句`select device_id, max(created_at) as created_at from device_msgs group by device_id`, 其中使用了聚合函数`max`(同样的聚合函数还有`count`等).

> 注意`as created_at`.

```sql
select device_id, max(created_at) as xxxxxx from device_msgs group by device_id;
 device_id |   xxxxxx
-----------+---------------------
         3 | 2018-12-24 00:00:00
         2 | 2018-12-21 00:00:00
         1 | 2018-12-21 00:00:00
(3 rows)
```

然后将此结果作为中间表与原来`device_msgs`表进行联合查询(联合查询默认的连接方式就是`inner join`).

本来我之前也想过这个方法来着, 但是由于表中并没有唯一键, `on`子句的条件没法写, 结果群友们根据`device_id`和`created_at`两个条件一起查询, 秀了我一脸. 只要保证一个时间点内同一个设备不会有多条数据入库就可以.

## 2. 窗口函数
