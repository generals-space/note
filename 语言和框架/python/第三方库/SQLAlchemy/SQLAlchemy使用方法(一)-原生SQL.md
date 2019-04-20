# SQLAlchemy使用方法(一)-原生SQL

参考文章

1. [Python SqlAlchemy使用方法](http://www.cnblogs.com/Xjng/p/4902498.html)

2. [Create a Postgres Data base using python](http://jakzaprogramowac.pl/pytanie/14336,create-a-postgres-data-base-using-python)

3. [Python SQLAlchemy基本操作和常用技巧（包含大量实例,非常好）](http://www.jb51.net/article/49789.htm)

4. [flask 系列教程（ 3）—SQLAlchemy 数据库](https://www.v2ex.com/t/376799)

SQLAlchemy的依赖包

```
$ yum install MySQL-python
```

```py
#!encoding: utf-8
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

## 格式 数据库(协议)类型://用户名:密码@IP:端口/库名
## 端口如果为默认可以省略, 库名也不是必选的.
dbAddr = 'mysql://root:@localhost'
## engine = create_engine(dbAddr, echo = True)
engine = create_engine(dbAddr, echo = True, isolation_level = 'AUTOCOMMIT')
DBSessionFactory = sessionmaker(bind = engine)
session = DBSessionFactory()

## session.execute()方法可以直接执行原生sql.
session.execute('create database db_1')
print(session.execute('use db_1'))

session.execute('use db_1')
session.execute('create table person (id int, name char(50))')
session.execute('insert into person values(1, "general")')
print(session.execute('select * from person').fetchall())

```

不算debug信息, 得到如下输出.

```
[(1L, 'general')]
```


关于`create_engine`方法中的`isolation_level`参数, 貌似不写这个只能执行到创建表这一步, 没法在里面插入数据...但是有输出, 而且也没报错. 猜测可能是由于提交时机的问题, 虽然SQL执行了, 但没有实际写入数据库. `echo`参数为True时可以打印debug信息.

本来是打算用postgresql完成的, 但pg的sql有点奇怪, 没法用原生SQL实现`use database`的功能, 放弃了. 不过如果事先选中指定库, 其他的SQL还是可以执行的. 如下, 可以正常输出`idc_type`表中的数据.

```py
#!encoding: utf-8
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

dbAddr = 'postgresql://sky:@localhost/sky'
## engine = create_engine(dbAddr, echo = True, isolation_level = 'AUTOCOMMIT')
engine = create_engine(dbAddr, echo = True)
DBSessionFactory = sessionmaker(bind = engine)
session = DBSessionFactory()

print(session.execute('select * from idc_type').fetchall())
```

简单分析一下.

`create_engine()`会返回一个数据库引擎

`sessionmaker()`会生成一个数据库会话类. 这个类的**实例**可以当成一个数据库连接, 它同时还记录了一些查询的数据, 并决定什么时候执行 SQL 语句. 由于 SQLAlchemy 自己维护了一个数据库连接池（默认 5 个连接）, 因此初始化一个会话的开销并不大. 