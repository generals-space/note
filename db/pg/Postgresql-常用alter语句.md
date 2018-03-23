1. 修改varchar类型的字段长度.

参考文章

[postgresql 修改字段长度](http://blog.csdn.net/baidu_18607183/article/details/78182275)

```sql
ALTER TABLE 表名 alter COLUMN 列名 type character varying(3000);
```