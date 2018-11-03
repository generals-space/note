# gorm表关联(四)-多对多

参考文章

1. [官方文档 - many to many](http://gorm.io/docs/many_to_many.html)

## 1. 模型定义

其实多对多关系的实现及中间表的建立可以参考skycmdb中sqlchemy的文档, 这里只探讨在gorm中的实现方法. 以书籍`Book`和标签`Tag`为例.

### 1.1 简单多对多

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

### 1.2 单向定义关联及自定义关联字段

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

结果和作用一样的.

同样, 由于外键默认指向的是id, 其实也可以通过这两个标记来指定非主键的字段来建立联系, 这个就不多说了.

### 1.3 自定义连接表字段名

这是官方文档中`Jointable ForeignKey`一节的内容.

默认情况下, 连接表有两个字段, 分别是两个表的主键, 命名为`book_id`和`tag_id`, 但是我们可以使用`jointable_foreignkey`和`association_jointable_foreignkey`来自定义这两个字段的名称. 前者定义的是对方表在连接表中的字段名, `association_jointable_foreignkey`则定义是己方表的字段名.

## 2. 使用示例

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
	connectStr := "host=localhost port=7723 user=gormtest dbname=gormdb sslmode=disable password=123456"
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
}
```

执行结果如下

```
2018/11/03 10:20:51 &{ID:1 Name:星辰变 Tags:[]}
2018/11/03 10:20:51 &{ID:1 Name:玄幻 Books:[]}
2018/11/03 10:20:51 &{ID:2 Name:热血 Books:[]}
```