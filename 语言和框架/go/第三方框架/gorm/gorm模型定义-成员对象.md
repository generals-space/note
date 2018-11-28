# gorm模型定义-成员对象

gorm不支持直接将结构体成员对象按照json形式存储到数据库中.

```go
package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Card ...
type Card struct {
	SerialNumber string `json:"serialNumber" gorm:"unique"`
}

// User ...
type User struct {
	ID   uint64 `json:"gorm" gorm:"primary_key,AUTO_INCREMENT"`
	Name string `json:"name"`
	Card *Card  `json:"card"`
}

func main() {
	var err error
	connectStr := "host=localhost port=7723 user=gormtest dbname=gormdb sslmode=disable password=123456"
	db, err := gorm.Open("postgres", connectStr)
	defer db.Close()

	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&User{})

	card := &Card{
		SerialNumber: "1111111111111111",
	}
	user := &User{
		Name: "general",
		Card: card,
	}
	db.Create(user)
}
```

...跟想像中的不太一样, 上述代码在运行后并不会在`users`表中创建`book`字段. 更不要说希望能存储`[]*Card`类型的数据了.

```
gormdb=# select * from users;
 id |  name
----+---------
  1 | general
(1 row)
```

## 2. 如何实现?

```go
package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
	"unsafe"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Card ...
type Card struct {
	SerialNumber string `json:"serialNumber" gorm:"unique"`
}

// UserCards ...
type UserCards []*Card

// User ...
type User struct {
	ID    uint64     `json:"gorm" gorm:"primary_key,AUTO_INCREMENT"`
	Name  string     `json:"name" gorm:"unique"`
	Cards *UserCards `json:"card" gorm:"type:text"`
}

// 以下是`UserCards`类型的成员方法. 有了这两个方法就能自动将结构体对象按照json格式存入数据库.

// Scan 解析前操作
func (u *UserCards) Scan(data interface{}) (err error) {
	if data == nil {
		return
	}

	var byteData []byte
	switch values := data.(type) {
	case []byte:
		byteData = values
	case string:
		byteData = []byte(values)
	default:
		err = errors.New("unsupported driver")
		return
	}
	err = json.Unmarshal(byteData, u)
	return
}

// Value 存储前操作
func (u *UserCards) Value() (driver.Value, error) {
	result, err := json.Marshal(u)
	return string(result), err
}

func main() {
	var err error
	connectStr := "host=localhost port=7723 user=gormtest dbname=gormdb sslmode=disable password=123456"
	db, err := gorm.Open("postgres", connectStr)
	defer db.Close()

	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&User{})

	/*
		card := &Card{
			SerialNumber: "1111111111111111",
		}
		user := &User{
			Name:  "general",
			Cards: &UserCards{card},
		}
		db.Create(user)
	*/

	myUser := &User{}
	db.Where(&User{Name: "general"}).First(myUser)
	log.Printf("%+v\n", myUser)

	userStr, _ := json.Marshal(myUser)
    log.Printf("%s\n", userStr)
    
	// 切片类型由别名`UserCard`转换回来还是比较难的, 不如直接先转成json再转成`[]*Card`
	cards := *(*[]*Card)(unsafe.Pointer(myUser.Cards))
	for _, c := range cards {
		log.Printf("%+v\n", c)
	}
}
```

上述代码在创建`uses`表时会自动创建`cards`列, 且类型为`text`, 存储的是json形式的字符串.

```
gormdb=# select * from users;
 id |  name   |                 cards
----+---------+---------------------------------------
  1 | general | [{"serialNumber":"1111111111111111"}]
(1 row)
```

而上述程序的查询输出如下

```
generals-MacBook-Pro:gormtest general$ go run main.go
2018/11/12 18:30:08 &{ID:1 Name:general Cards:0xc00013f2e0}
2018/11/12 18:30:08 {"gorm":1,"name":"general","card":[{"serialNumber":"1111111111111111"}]}
2018/11/12 18:30:08 &{SerialNumber:1111111111111111}
```

很完美.

> 注意: 成员对象的标记中必须要用`gorm:"type:text"`来存储文本模式的内容.