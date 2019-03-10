# golang-sql数据库操作(坑)

参考文章

1. [Go实战--go语言操作PostgreSQL数据库(github.com/lib/pq)](https://blog.csdn.net/wangshubo1989/article/details/77478838?locationNum=3&fps=1)

2. [go语言数据库查询后对结果的处理方法的探讨](https://blog.csdn.net/westhod/article/details/80799266)

关于标准库`database/sql`的使用, 参考文章写的比较清楚了. 下面先给出连接的示例代码.

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbHost   = "localhost"
	dbPort   = 7723
	dbUser   = "gormtest"
	dbPasswd = "123456"
	dbName   = "gormdb"
)

func main() {
	connectStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPasswd, dbName)
	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
```

参考文章1中有插入和查询的示例, 插入操作比较符合常规, 变量可以用`$1`, `$2`这种占位符表示. 

但是查询, 简直是噩梦. 查询的结果需要通过`Scan`操作将各字段的值填写到指定好的变量中. 又不能像js那样直接从object里取属性, 或是像python那样按照select中字段的顺序通过索引取成员, 于是在go里, 查询语句应如下.

```go
	// obj := make(map[string]interface{}) ...妄图用map对象将所有字段取出的行为是愚蠢的.
	// 结尾可以不加分号
	queryStr := "select id, name, leading_role from books"
	rows, err := db.Query(queryStr)
	if err != nil {
		panic(err)
    }
    // for循环依次取出
	for rows.Next() {
		var id uint64
		var name string
		var leadingRole string
		err = rows.Scan(&id, &name, &leadingRole)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name, leadingRole)
	}
```

如果你想得到一个结果列表, 就只能用for循环一个一个`Scan`再`append`了.

参考文章2中提出了一种映射思路, `select *`时通过反映依次为每个字段赋值...嗯, 想想skycmdb项目里model定义中`get_values()`方法, 应该是同样的作用.

...但是还是拒绝使用这个库了, 人生苦短, 我用gorm.