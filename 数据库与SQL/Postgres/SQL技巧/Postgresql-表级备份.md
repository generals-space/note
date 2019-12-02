## Postgresql-表级备份

```sql
sky=> select * from idc_info into idc_info_20170830;
ERROR:  syntax error at or near "into"
LINE 1: select * from idc_info into idc_info_20170830;
                               ^
sky=> select * into idc_info_20170830 from idc_info;
SELECT 1086
```

`select 字段 into 目标表 from 源表`语句可以备份指定表(也可以指定字段), 目标表对象不必事先存在, 而且`into`与`from`不能调换位置, 否则会出错.
