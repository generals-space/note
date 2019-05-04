# golang-各种类型的引用传递和值传递

```go
package main

import "log"

// User ...
type User struct {
	Name string
	Age  int
}

func main() {
	log.Println("======= struct 值传递")
	user1 := User{Name: "user1", Age: 21}
	user2 := user1
	user2.Name = "user2"
	log.Printf("%+v\n", user1) // {Name:user1 Age:21}
	log.Printf("%+v\n", user2) // {Name:user2 Age:21}

	log.Println("======= slice 值传递")
	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice2 = append(slice2, 4)
	log.Printf("%+v\n", slice1) // [1 2 3]
	log.Printf("%+v\n", slice2) // [1 2 3 4]

	log.Println("======= array 值传递")
	array1 := [3]int{1, 2, 3}
	array2 := array1
	array2[0] = 0
	log.Printf("%+v\n", array1) // [1 2 3]
	log.Printf("%+v\n", array2) // [0 2 3]

	log.Println("======= map 引用传递")
	map1 := map[string]interface{}{
		"Name": "general",
		"Age":  21,
	}
	map2 := map1
	map2["Name"] = "longbei"

	log.Printf("%+v\n", map1) // map[Name:longbei Age:21]
	log.Printf("%+v\n", map2) // map[Name:longbei Age:21]

	log.Println("======= channel 引用传递")
	channel1 := make(chan int, 10)
	channel2 := channel1
	channel1 <- 1

	log.Printf("len %d\n", len(channel1)) // len 1
	log.Printf("len %d\n", len(channel2)) // len 1
}
```

我觉得唯一需要注意的就是slice了, 本来以为切片是引用传递的.
