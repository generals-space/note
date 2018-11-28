# 表关联(三)-包含多个

参考文章

1. [官方文档 - has many](http://gorm.io/docs/has_many.html)

## 1. 模型定义

`has many`和`belong to`基本上是一个东西: 对于`Book`表来说, 一本书`belong to`一个用户; 而对于`User`表来说, 一个用户`has many`本书. 只不过是反过来说而已.

在`belong to`的示例中, `Book`表中有指向`User`的`UserID`字段作外键, `User`表中却没有反向的字段.

而在`has many`的示例中, `User`拥有`[]Book`类型的字段, 但`Book`表中却只保留了`UserID`字段, 不再用`User`字段. 如下

```go
// User ...
type User struct {
	ID   uint64 `gorm:"primary_key,AUTO_INCREMENT"`
    Name string `gorm:"unique"`
    Books []*Book
}

// Book ...
type Book struct {
	ID     uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Name   string
	UserID uint `gorm:"not null"`
}
```

说好听点是万变不离其宗, 说难听点是换汤不换药...

因为按照这个结构创建的表实际上还是`books.user_id`作为外键, 当然, 在数据库里根本没有外键...

## 2. 使用方法

```go
package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// User ...
type User struct {
	ID    uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Name  string `gorm:"unique"`
	Books []*Book
}

// Book ...
type Book struct {
	ID     uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Name   string
	UserID uint `gorm:"not null"`
}

func main() {
	var err error
	connectStr := "host=localhost port=7723 user=gormtest dbname=gormdb sslmode=disable password=123456"
	db, err := gorm.Open("postgres", connectStr)
	defer db.Close()

	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&User{}, &Book{})
	/*
		user := &User{
			Name: "包拯",
			Books: []*Book{
				&Book{Name: "少年包青天1"},
				&Book{Name: "少年包青天2"},
				&Book{Name: "少年包青天3"},
			},
		}
		err = db.Create(&user).Error
		if err != nil {
			log.Println(err)
		}
	*/

	u := &User{}
	bs := []*Book{}

	db.Where(&User{Name: "包拯"}).First(u)
	log.Printf("%+v\n", u)
	db.Find(&bs)
	for _, b := range bs {
		log.Printf("%+v\n", b)
	}

	// 同样可以使用Related()方法代替外键进行查询和反向引用
	log.Println("query a user's all books...")
	ubs := []*Book{}
	// 查询指定用户的所有书籍
	db.Model(u).Related(&ubs)
	for _, b := range bs {
		log.Printf("%+v\n", b)
	}

	log.Println("query which one has the book...")
	b := &Book{}
	db.Where(&Book{Name: "少年包青天1"}).First(b)
	// 查询一本书属于哪个用户?
	bu := &User{}
	db.Model(b).Related(bu)
	log.Printf("%+v\n", bu)
}
```

执行结果

```
2018/11/02 17:52:19 &{ID:1 Name:包拯 Books:[]}
2018/11/02 17:52:19 &{ID:1 Name:少年包青天1 UserID:1}
2018/11/02 17:52:19 &{ID:2 Name:少年包青天2 UserID:1}
2018/11/02 17:52:19 &{ID:3 Name:少年包青天3 UserID:1}
2018/11/02 17:52:19 query a user's all books...
2018/11/02 17:52:19 &{ID:1 Name:少年包青天1 UserID:1}
2018/11/02 17:52:19 &{ID:2 Name:少年包青天2 UserID:1}
2018/11/02 17:52:19 &{ID:3 Name:少年包青天3 UserID:1}
2018/11/02 17:52:19 query which one has the book...
2018/11/02 17:52:19 &{ID:1 Name:包拯 Books:[]}
```