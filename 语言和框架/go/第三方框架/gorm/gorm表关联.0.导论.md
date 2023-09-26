# 表关联

参考文章

1. [gorm官方文档 - 关联](http://gorm.io/docs/belongs_to.html)

`关联`, 无非就是外键, 但是像python的sqlAlchemy可以通过外键正向查询被引用表, 而被引用表也可以反向查询谁引用了自己, 在程序里写起来是非常方便的. 这是orm本身提供的快捷方案, 使用起来很舒服.

我们来看看在gorm中如何实现. 官方文档里列出了5种关联:

1. 属于

2. 包含一个

3. 包含多个

4. 多对多

5. 多种包含

使用如下docker-compose配置部署依赖环境

```yml
version: '2'
services:
  postgres:
    image: postgres
    environment:
    - "POSTGRES_USER=gormtest"
    - "POSTGRES_PASSWORD=123456"
    - "POSTGRES_DB=gormdb"
    ports:
    - "7723:5432"
    volumes:
      - ~/Public/dbdata/gormtest:/var/lib/postgresql/data
```

> win下就不要映射目录了...