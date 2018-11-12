# gorm模型定义-数组类型

参考文章

1. [invalid sql type StringArray (slice) for postgres](https://github.com/jinzhu/gorm/issues/1248)

## 1. 常规类型

如下代码中`Titles`和`Ages`分别是字符串和整型的数组.

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
	ID     uint64   `json:"gorm" gorm:"primary_key,AUTO_INCREMENT"`
	Name   string   `json:"name"`
	Titles []string `json:"titles"`
	Ages   []int    `json:"ages"`
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

	u := &User{
		Name:   "general",
		Titles: []string{"CEO", "CTO", "CFO"},
		Ages:   []int{12, 3, 56},
	}
	db.Create(u)
}

```

在运行上述代码时, `AutoMigrate()`操作会报`panic: invalid sql type  (slice) for postgres`错误.

根据参考文章1, 了解到gorm并不支持通过golang切片直接创建pg的数组类型. 想要使用pg的数组类型需要引入`lib/pq`库. 如下

```go
import(
	"github.com/lib/pq"
)

// User ...
type User struct {
	ID     uint64         `json:"gorm" gorm:"primary_key,AUTO_INCREMENT"`
	Name   string         `json:"name"`
	Titles pq.StringArray `json:"titles" gorm:"type:varchar(64)[]"`
	Ages   pq.Int64Array  `json:"ages" gorm:"type:int[]"`
}
```

## 2. 对象数组