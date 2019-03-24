# Postgresql-常用alter语句

参考文章

[postgresql 修改字段长度](http://blog.csdn.net/baidu_18607183/article/details/78182275)

## 更改字段类型(修改varchar类型的字段长度)

```sql
ALTER TABLE 表名 alter COLUMN 列名 type character varying(3000);
```

## 更改已存在的用户密码

```
alter user postgres with password 'NewPassword';
```

## 创建列(默认值: 非空)

```sql
alter table idc_info add column status int default 1 not null;
alter table idc_info add column status int not null default 1;
```

为目标表`idc_info`新增一个字段`status`, `int`类型, 默认值为1, 非空约束.