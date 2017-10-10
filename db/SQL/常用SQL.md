## 1.

```
sky=> select * from idc_info into idc_info_20170830;
ERROR:  syntax error at or near "into"
LINE 1: select * from idc_info into idc_info_20170830;
                               ^
sky=> select * into idc_info_20170830 from idc_info;
SELECT 1086
```

`select 字段 into 目标表 from 源表`语句可以备份指定表(也可以指定字段), 目标表对象不必事先存在, 而且into与from不能调换位置, 否则会出错.

## 2.

```
alter table idc_info add column status int default 1 not null;
```

为目标表idc_info新增一个字段status, int类型, 默认值为1, 非空约束.