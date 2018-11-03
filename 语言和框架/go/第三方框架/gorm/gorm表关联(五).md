# gorm表关联(五)

参考文章

1. [官方文档 - Associations](http://gorm.io/docs/associations.html)

通过前面几节我们可以了解到, gorm在模拟外键实现上封装了一些操作, 使创建引用等常规操作都是通过model对象的成员字段来完成的. gorm会在这些操作中自动在数据库层面加上对关联表的操作, 以保证关联的完整性.

这一节的官方文档中主要讲解了如何去除这些自动的操作, 但我感觉十分没有必要...不然要你干啥?

不过文章最后`Association Mode`这一小节倒是又给出了另外一种关联查询的方法 - `Association()`函数. 可以实现与`Related()`函数相似的功能.

在使用`Related()`方法时, 我们需要先将关联的记录查出来, 才能操作; 而`Association()`则可以直接对关联记录进行操作. 

比如在`many to many`和`have many`关联关系中, 可以直接为主引用表添加新的引用记录.

```go
db.Model(&user).Association("Book").Append([]Book{book4, book4})
```

也可以在`has one`关联关系中, 更改一对一关系

```go
db.Model(&user).Association("Card").Append(newCard)
```

还有`Replace`, `Delete`, `Clear`和`Count`函数, 应该比较好理解. 就不具体实验了.