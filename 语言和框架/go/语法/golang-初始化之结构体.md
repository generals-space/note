# golang-初始化之结构体

参考文章

1. [go语言初始化内部结构体3中方式](https://studygolang.com/articles/3085)

2. [golang 嵌套struct 如何直接初始化？](https://www.zhihu.com/question/22746100)

## 1. 普通结构体

```go
package main

import "log"

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	userA := User{
		Id:   1,
		Name: "general",
		// 最后一个成员的必须要有一个逗号
		Age: 21,
	}
	// 输出{Id:1 Name:general Age:21}
	log.Printf("%+v", userA)

	// 可以按照成员排列顺序进行初始化, 但是不够灵活
	userB := User{2, "jiangming", 24}
	// {Id:2 Name:jiangming Age:24}
	log.Printf("%+v", userB)
}
```

## 2. 嵌套结构体

```go
package main

import "log"

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
    User
    Employer []User
    Title string
    Company []string
}

func main() {
	mA := Manager {
        User: User{
            Id:   1,
            Name: "general",
            // 最后一个成员的必须要有一个逗号
            Age: 21,
        },
        Title: "CEO",
        // 结构体内嵌基本类型切片
        Company: []string{"百度", "阿里", "腾讯"},
        // 结构体内嵌结构体切片
        Employer: []User{
            User{
                Id: 2,
                Name: "李彦宏",
                Age: 42,
            },
            User{
                Id: 3,
                Name: "马云",
                Age: 48,
            },
            User{
                Id: 4,
                Name: "马化腾",
                Age: 38,
            },
        },
	}
	// 输出{User:{Id:1 Name:general Age:21} Employer:[{Id:2 Name:李彦宏 Age:42} {Id:3 Name:马云 Age:48} {Id:4 Name:马化腾 Age:38}] Title:CEO Company:[百度 阿里 腾讯]}
	log.Printf("%+v", mA)
}
```

## 3. 嵌套 + 匿名结构体

这个例子是参考文章2中提出的...感觉很有意义.

```go
package main

import "log"

type User struct {
	Id      uint32
	Name    string
	Company struct {
		Owner   string
		Address string
	}
}

func main() {
	userA := User{
		Id:   1,
		Name: "general",
		// 匿名结构体初始化要给出具体结构
		Company: struct {
			Owner   string
			Address string
		}{
			Owner:   "general",
			Address: "杭州",
		},
	}
	// {Id:1 Name:general Company:{Owner:general Address:杭州}}
	log.Printf("%+v", userA)

	// 不过不建议使用上述方法初始化
	// new方法返回的是指针变量
	userB := new(User)
	userB.Id = 2
	userB.Name = "jiangming"
	userB.Company.Owner = "jiangming"
	userB.Company.Address = "北京"
	// &{Id:2 Name:jiangming Company:{Owner:jiangming Address:北京}}
	log.Printf("%+v", userB)
}
```