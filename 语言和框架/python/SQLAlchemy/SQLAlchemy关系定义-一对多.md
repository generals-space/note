# SQLAlchemy关系定义-一对多

<!--

<!tags!>: <!sqlalchemy!> <!设计原则!>

-->

参考文章

1. [使用SQLAlchemy](http://www.liaoxuefeng.com/wiki/001374738125095c955c1e6d8bb493182103fac9270762a000/0014021031294178f993c85204e4d1b81ab032070641ce5000)

2. [SQLAlchemy_定义(一对一/一对多/多对多)关系](http://blog.csdn.net/Jmilk/article/details/52445093)


在表示数据库一对多的关系时，一般通过在子表类中使用外键约束(`foreign key`)引用父表类。

在SQLAlchemy中, 还需要在父表类中通过`relationship()`方法引用子表类.

举个例子, 如果一个User拥有多个Book，就可以定义一对多关系如下.

```py
class User(Base):
    __tablename__ = 'user'

    id = Column(String(20), primary_key=True)
    name = Column(String(20))
    # 一对多:
    books = relationship('Book')

class Book(Base):
    __tablename__ = 'book'

    id = Column(String(20), primary_key=True)
    name = Column(String(20))
    # “多”的一方的book表是通过外键关联到user表的:
    user_id = Column(String(20), ForeignKey('user.id'))
```

在这个例子中, User为父表类, Book为子表类, Book表中每一个`user_id`都必须是User表中已经存在的行的id. 

这样就已经实现了User与Book的**一对多**关系映射.

SQLAlchemy还要求在`User`类中定义`relationship()`方法以建立双向关系.

> 在纯数据库层面的设计中, 外键只是一种约束(另外三种包括: 唯一, 非空, 范围), 并不是用来方便进行多表查询的. 多表查询需要使用连接语句`join`. 虽然SQLAlchemy中`relationship()`语句多与Foreign Key同时出现, 前者也的确实现了连接查询, 但是要明确, **不要滥用外键约束**.