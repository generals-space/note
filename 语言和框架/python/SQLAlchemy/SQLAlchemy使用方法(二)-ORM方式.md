# SQLAlchemy使用方法(二)-ORM方式

## 1. 表的创建与删除

```py
#!encoding: utf-8
from sqlalchemy import create_engine
from sqlalchemy import Column
from sqlalchemy.types import CHAR, Integer, String
from sqlalchemy.ext.declarative import declarative_base

dbAddr = 'postgres://sky:@localhost/sky'
engine = create_engine(dbAddr, echo = True)

BaseModel = declarative_base()

def init_db():
    BaseModel.metadata.create_all(engine)
def drop_db():
    BaseModel.metadata.drop_all(engine)

class User(BaseModel):
    __tablename__ = 'user'
    id = Column(Integer, primary_key = True)
    name = Column(CHAR(30)) # or Column(String(30))

init_db()
```

上述代码会以`sky`用户身份在`sky`库中创建名为`user`的表.

`declarative_base()` 创建了一个 `BaseModel` 类，这个类的**子类**可以自动与一个表关联。以`User`子类为例，它的`__tablename__` 属性就是数据库中该表的名称，它有 id 和 name 这两个字段，分别为整型和 30 个定长字符。Column 还有一些其他的参数，我就不解释了。

最后，`BaseModel.metadata.create_all(engine)` 会找到 BaseModel 的所有子类，并在数据库中建立这些表；`drop_all()` 则是删除这些表。

## 2. 表操作-增删改查

```py
#!encoding: utf-8
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

from sqlalchemy import Column
from sqlalchemy.types import CHAR, Integer, String
from sqlalchemy.ext.declarative import declarative_base

dbAddr = 'postgres://sky:@localhost/sky'
engine = create_engine(dbAddr, echo = True)
DBSessionFactory = sessionmaker(bind = engine)
session = DBSessionFactory()

BaseModel = declarative_base()

class User(BaseModel):
    __tablename__ = 'user'
    id = Column(Integer, primary_key = True)
    name = Column(CHAR(30)) # or Column(String(30))


## 新增
user = User(id = 1, name = 'abcd')
session.add(user)
user = User(id = 2, name = '1234')
session.add(user)
session.commit()                            ## 多次add但只有一次提交

## 查询
query = session.query(User)
users = query.all()

#### all()方法返回的是一个列表对象
print(type(users))                          ## <List>
#### all()方法返回的第一个列表成员对象即为first()方法获得的对象
print(users[0] == query.first())            ## True

for i in users:
    print(i)                                ## i并不是一个简单dict对象, 它应该是一个类的实例
    print(i.name)                           ## 但依然可以通过点号'.'取得对象属性


## 删除
from sqlalchemy import or_
query.filter(or_(User.name == 'abcd', User.name == '1234')).delete()
session.commit()
```


需要注意的是, 不是所有query的filter()方法后都能直接跟`delete()`

比如`filter(User.id.in_(1, 2)).delete()`就不行, `filter(User.name.contains('abc')).delete()`也不行

这种情况下的`delete()`只能这么干

```py
userCollection = query.filter(User.name.contains('abc'))
session.delete(userCollection)
session.commit()
```

------

另外还有更新的语法

```py
query.filter(User.id == 2).update({'name': '1111'})
query.filter(User.id == 9).update({'name': '0000'})
session.commit()
```

还有一点需要注意, `filter()`语句查出的结果有可能为空, 但后面的接的`.delete()`和`.update()`不会报错. 所以我认为有必要在delete或update之前先判断目标对象是否存在.