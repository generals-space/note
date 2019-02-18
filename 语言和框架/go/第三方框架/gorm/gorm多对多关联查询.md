# gorm多对多关联查询

参考文章

1. [Many-To-Many relationships?](https://github.com/jinzhu/gorm/issues/168)

多对多关联的的简查查询可以用`Related()`或`Association()`实现, 但是如果想在多对多查询同时进行过滤, 分页和排序等操作时如何实现? 是在查询中间表的时候手动进行分页这些操作得到另一张表的id集合, 然后再去用`in`操作符查询相应的记录吗?