# 表关联(一)-属于

1. [gorm官方文档 - belong to](http://gorm.io/docs/belongs_to.html)

## 1. 模型定义

基本逻辑: 每个用户可以拥有多本书, 每本书都必须有一个主人

```go
// User ...
type User struct {
	ID   uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Name string `gorm:"unique"`
}

// Book ...
type Book struct {
	ID     uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Name   string
	User   *User
	UserID uint `gorm:"not null"`
}
```

以这种方式建表, 在数据库里`users`与`books`不会有任何外键关系, 在`books`表中, 会出现一个`user_id`字段, 没有`user`字段(不管`Book`结构体中的User字段是`User`类型还是`*User`类型).

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
	ID   uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Name string `gorm:"unique"`
}

// Book ...
type Book struct {
	ID     uint64 `gorm:"primary_key,AUTO_INCREMENT"`
	Name   string
	User   *User
	UserID uint `gorm:"not null"`
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

	db.AutoMigrate(&User{}, &Book{})
	/*
		user := &User{
			Name: "包拯",
		}
		err = db.Create(&user).Error
		if err != nil {
			log.Println(err)
		}
		books := []*Book{
			&Book{
				Name: "少年包青天1",
				User: user,
			},
			&Book{
				Name: "少年包青天2",
				User: user,
			},
			&Book{
				Name: "少年包青天3",
				User: user,
			},
		}
		// 在插入`book`记录时, 给其中的`User`字段赋值为一个已经存在的`user`记录, 就完成了外键的引用, 在数据库中这条记录的`user_id`的值就会成为`user.id`.
		for _, b := range books {
			db.Create(&b)
		}
	*/
	log.Println("getting books...")
	// 注意: 在查询book表时, 得到的结果中book记录的`User`字段总是为`nil`的
	bs := []*Book{}
	db.Find(&bs)
	for _, b := range bs {
		log.Printf("%+v\n", b)
	}
	// 如果想得到一个book记录实际引用的user记录, 可以使用`Related()`方法再查一次
	// 根据主引用表记录查询被引用表的记录
	bu := &User{}
	b := &Book{}
	db.Where(&Book{Name: "少年包青天1"}).First(b)
	db.Model(b).Related(bu)
	log.Printf("%+v\n", bu)

	/////////////////////////////////////////////////////////////////
	log.Println("getting users...")
	us := []*User{}
	db.Find(&us)
	for _, u := range us {
		log.Printf("%+v\n", u)
	}

	log.Println("related query...")
	// 我想查出一个用户名下有多少本书? 同样可以使用`Related()`方法
	// 根据被引用表记录查询主引用表的记录.
	u := &User{}
	ubs := []*Book{}
	db.Where(&User{Name: "包拯"}).First(u)
	log.Printf("%+v\n", u)
	db.Model(u).Related(&ubs)
	for _, ub := range ubs {
		log.Printf("%+v\n", ub)
	}
}
```

以上代码的输出为

```
2018/11/02 15:41:02 getting books...
2018/11/02 15:41:02 &{ID:1 Name:少年包青天1 User:<nil> UserID:1}
2018/11/02 15:41:02 &{ID:2 Name:少年包青天2 User:<nil> UserID:1}
2018/11/02 15:41:02 &{ID:3 Name:少年包青天3 User:<nil> UserID:1}
2018/11/02 15:41:02 &{ID:1 Name:包拯}
2018/11/02 15:41:02 getting users...
2018/11/02 15:41:02 &{ID:1 Name:包拯}
2018/11/02 15:41:02 related query...
2018/11/02 15:41:02 &{ID:1 Name:包拯}
2018/11/02 15:41:02 &{ID:1 Name:少年包青天1 User:<nil> UserID:1}
2018/11/02 15:41:02 &{ID:2 Name:少年包青天2 User:<nil> UserID:1}
2018/11/02 15:41:02 &{ID:3 Name:少年包青天3 User:<nil> UserID:1}
```

> `Related()`不如SqlAlchemy智能, 方便...垃圾.

`UserID`是默认的外键形式, 可以通过`gorm:"foreignkey:UserRefer"`自定义外键名称.

另外, 还有一个`association_foreignkey`的标记, 这个代表什么呢?

默认情况下, 主引用表的外键字段连接的是被引用表的主键, 即ID. 使用`association_foreignkey`可以指定连接目标表中任意字段.