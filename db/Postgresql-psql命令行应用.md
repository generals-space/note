# PostgreSQL-psql命令行应用
```
ostgres@dev-cmdb-> psql -U sky
psql (9.5.2)
Type "help" for help.

sky=> \l            ## 查看数据库
sky=> \c sky        ## 选择数据库
You are now connected to database "sky" as user "sky".
sky=> ls
sky-> \dt           ## 查看表
sky-> 

```

`\?`: 查看psql风格的命令帮助列表

`\d 表|视图|索引|序列`: 查看表结构