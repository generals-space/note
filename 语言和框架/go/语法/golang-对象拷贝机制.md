# golang-对象拷贝机制

```go
package main

import "log"

// Test ...
type Test struct {
	Name string
	Age  int
}

func main() {
	log.Println("========================")
	test1 := Test{
		Name: "general",
		Age:  21,
	}

	test2 := test1
	test2.Name = "jiangming"

	log.Printf("%+v\n", test1) // {Name:general Age:21}
	log.Printf("%+v\n", test2) // {Name:jiangming Age:21}

	log.Println("========================")
	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice2 = append(slice2, 4)
	log.Printf("%+v\n", slice1) // [1 2 3]
	log.Printf("%+v\n", slice2) // [1 2 3 4]

	log.Println("========================")
	map1 := map[string]interface{}{
		"Name": "general",
		"Age":  21,
	}
	map2 := map1
	map2["Name"] = "longbei"

	log.Printf("%+v\n", map1) // map[Name:longbei Age:21]
	log.Printf("%+v\n", map2) // map[Name:longbei Age:21]

	log.Println("========================")
	channel1 := make(chan int, 10)
	channel2 := channel1
	channel1 <- 1

	log.Printf("len %d\n", len(channel1)) // len 1
	log.Printf("len %d\n", len(channel2)) // len 1
}

```

结构体实例(非指针)和切片的复制是真的复制.

结构体指针类型, map和chan的复制其实是引用复制.