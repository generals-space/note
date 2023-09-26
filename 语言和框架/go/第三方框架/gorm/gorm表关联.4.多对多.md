# gorm表关联(四)-多对多

参考文章

1. [官方文档 - many to many](http://gorm.io/docs/many_to_many.html)

## 1. 模型定义

其实多对多关系的实现及中间表的建立可以参考skycmdb中sqlchemy的文档, 这里只探讨在gorm中的实现方法. 以书籍`Book`和标签`Tag`为例.

```go
// Tag ...
type Tag struct {
	ID    uint64  `gorm:"primary_key,AUTO_INCREMENT"`
	Name  string  `gorm:"unique"`
	Books []*Book `gorm:"many2many:tag_book;"`
}

// Book ...
type Book struct {
	ID   uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Name string
	Tags []*Tag `gorm:"many2many:tag_book;"`
}
```

gorm会为我们隐式创建一个中间表, 就是`many2many`标签指定的表`tag_book`, 而`tag_book`表有两个字段, 就是`tags`表和`books`表的主键, 分别命名为`tag_id`和`book_id`.

```
gormdb=# \dt
          List of relations
 Schema |   Name   | Type  |  Owner
--------+----------+-------+----------
 public | books    | table | gormtest
 public | tag_book | table | gormtest
 public | tags     | table | gormtest
(3 rows)

gormdb=# select * from tag_book;
 tag_id | book_id
--------+---------
(0 rows)
```

代码示例

```go
package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Tag ...
type Tag struct {
	ID    uint64  `gorm:"primary_key,AUTO_INCREMENT"`
	Name  string  `gorm:"unique"`
	Books []*Book `gorm:"many2many:tag_book;"`
}

// Book ...
type Book struct {
	ID   uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Name string
	Tags []*Tag `gorm:"many2many:tag_book;"`
}

func main() {
	var err error
	// 两种连接方式都可以
	// connectStr := "host=localhost port=7723 user=wuhou password=123456 dbname=wuhoudb sslmode=disable"
	connectStr := "postgresql://gormtest:123456@localhost:7723/gormdb?sslmode=disable"
	db, err := gorm.Open("postgres", connectStr)
	defer db.Close()

	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&Tag{}, &Book{})

	/*
		tag1 := &Tag{Name: "玄幻"}
		tag2 := &Tag{Name: "热血"}
		tag3 := &Tag{Name: "斗气"}
		db.Create(tag1)
		db.Create(tag2)
		db.Create(tag3)

		book1 := &Book{
			Name: "星辰变",
			Tags: []*Tag{tag1, tag2},
		}
		book2 := &Book{
			Name: "斗破苍穹",
			Tags: []*Tag{tag2, tag3},
		}
		db.Create(book1)
		db.Create(book2)
	*/

	book := &Book{}
	// 同样, 单纯Find/First查询出来的记录, Tags字段是不会有值的
	db.Where(&Book{Name: "星辰变"}).Find(book)
	log.Printf("%+v\n", book)
	// 要获取这本书所属的Tag, 还是需要使用Related()方法
	// 但是多对多关系中, Related()函数需要第二个参数, 貌似是目标表名(大小写无关)
	bTags := []*Tag{}
	db.Model(book).Related(&bTags, "tags")
	for _, t := range bTags {
		log.Printf("%+v\n", t)
	}

	log.Println("========================")
	tag := &Tag{}
	db.Where(&Tag{Name: "热血"}).Find(tag)
	log.Printf("%+v\n", tag)
	tBooks := []*Book{}
	db.Model(tag).Related(&tBooks, "books")
	for _, b := range tBooks {
		log.Printf("%+v\n", b)
	}
}
```

执行结果如下

```
2018/11/09 22:13:26 &{ID:1 Name:星辰变 Tags:[]}
2018/11/09 22:13:26 &{ID:1 Name:玄幻 Books:[]}
2018/11/09 22:13:26 &{ID:2 Name:热血 Books:[]}
2018/11/09 22:13:26 ========================
2018/11/09 22:13:26 &{ID:2 Name:热血 Books:[]}
2018/11/09 22:13:26 &{ID:1 Name:星辰变 Tags:[]}
2018/11/09 22:13:26 &{ID:2 Name:斗破苍穹 Tags:[]}
```

> 注意, 多对多关联查询中`Related()`的第二个参数为目标关系的表名, 不是结构体名, 且大小写无关, 不过最好直接写小写的表名, 与数据库中的表保持一致.

## 2. 单向定义关联及自定义关联字段

我们知道`association_foreignkey`和`foreignkey`分别可以作为外键引用和反向引用. 而在gorm里, model结构体内包含另一个model类型其实除了表示连接关系外是没什么用的. 上面的模型定义中多对多关系是双方同时定义的, 我们也可以在单一model中通过同时指定这两个标记来完成这种关系的建立(本来默认的就是使用双方的id主键建立的关联, 手动建立问题不大)).

```go
// Tag ...
type Tag struct {
	ID   uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Name string `gorm:"unique"`
}

// Book ...
type Book struct {
	ID   uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Name string
	Tags []*Tag `gorm:"many2many:tag_book;association_foreignkey:id;foreignkey:id"`
}
```

~~结果和作用一样的.~~

同样, 由于外键默认指向的是id, 其实也可以通过这两个标记来指定非主键的字段来建立联系, 这个就不多说了.

------

我错了, 这种定义方式和上面的是有区别的. 使用同样的代码, 程序输出如下

```
2018/11/09 22:19:18 &{ID:1 Name:星辰变 Tags:[]}
2018/11/09 22:19:18 &{ID:1 Name:玄幻}
2018/11/09 22:19:18 &{ID:2 Name:热血}
2018/11/09 22:19:18 ========================
2018/11/09 22:19:18 &{ID:2 Name:热血}
/////////////////////windows下有乱码, 不要在乎
?[35m(invalid association [books])?[0m
?[33m[2018-11-09 22:19:18]?[0m ?[31;1m ?[0m
```

由于`Book`结构体中定义了`Tags`成员, 所以正向关联查询是可以的, 但是没有办法通过tag实例反向查询book了...

## 3 自定义连接表字段名

这是官方文档中`Jointable ForeignKey`一节的内容.

默认情况下, 连接表有两个字段, 分别是两个表的主键, 命名为`book_id`和`tag_id`, 但是我们可以使用`jointable_foreignkey`和`association_jointable_foreignkey`来自定义这两个字段的名称. 前者定义的是对方表在连接表中的字段名, `association_jointable_foreignkey`则定义是己方表的字段名.

没做实验, 暂不作深入讨论.

------

注意:

多对多关联中, 使用gorm删除一方的记录, 中间表中不会删除此条记录, 因为被删除的记录id以后不会再出现, 只是留着关联记录. 不过使用另一方对方进行关联查询时是不会得到已经删除了的关系的.

但是!!! 关联查询貌似是没有办法与`Limit()`, `Order()`等函数一起用的, 也就是说, `Related()`函数查出来的是所有关联数据...md我一个标签下有10000本书要全都查出来???

所以只能手动从中间表中得到关联关系, 然后再从对方表中继续查询才行(此时你需要事先定义好中间表的Model模型结构, 比如叫`TagBook`)...

`db.Table("tag_book").Limit(20).Find(&tagBooks)`

那么问题来了, 已经删除了的记录是不会在中间表中删除的, 那么中间表中limit(20)得到的20条记录可能有失效记录...

md所以在删除一方记录后还需要手动删除中间表中记录, fuck...

...另外!!!

关于更新关联对象...md我本来准备通过`book.Tags = []*Tag{newTag}`为一个book对象**重设**标签关联的, 然而...这个操作只是新增的操作, 结果这个book就有了3个标签. 

我擦, 所以更新的时候也要手动操作中间表呗?

...好在不用, 我试了下高级的`Association()`函数, `db.Model(book).Association("tags").Clear()`这一句可以直接删除中间表中此book对象的所有tag关联记录, 所以在删除和更新时可以用`Association()`做一个预处理.

神经病...