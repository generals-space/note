# SQL语法 - 关于外键

参考文章

1. [mysql中的外键foreign key](https://www.cnblogs.com/pengyin/p/6375860.html)

2. [PostgreSQL外键约束reference](http://blog.csdn.net/wujiang88/article/details/51578794)

这里不讲外键的定义方法, 不管是在建表时指定, 还是在建表后添加, 只要知道外键的基本作用即可.

使用外键的前提：

1. 表储存引擎必须是innodb，否则创建的外键无约束效果(pg和mysql的默认存储引擎都是innodb);

2. 外键的列类型必须与父表的主键类型完全一致;

3. 外键的名字不能重复(默认为`主引用表名_被引用表字段名_fkey`);

4. 已经存在数据的字段被设为外键时，必须保证字段中的数据与父表的主键数据对应起来;

5. 外键可为空, 通过设置和取消`not null`来修改, 也可以设备默认值, 但这个值必须在被引用表中能匹配到;

6. 只有唯一键可以被用作外键, 所以被引用字段必须为唯一约(unique key)束或者主键约束(primary key);

然后补充一些关于外键的认知

1. 一张表中可以有多个外键(可以是同一张被引用表的多个字段, 也可以不同被引用表的字段).


外键的默认作用有两点：

1. 对子表(外键所在的表)的作用：子表在进行写操作的时候，如果外键字段在父表中找不到对应的匹配，操作就会失败。

2. 对父表的作用：对父表的主键字段进行删和改时，如果对应的主键在子表中被引用，操作就会失败。

外键的定制作用----三种约束模式：

　　　　district：严格模式(默认), 父表不能删除或更新一个被子表引用的记录。

　　　　cascade：级联模式, 父表操作后，子表关联的数据也跟着一起操作。

　　　　set null：置空模式(前提外键字段不能为`not null`), 父表操作后，子表对应的字段被置空。
