# golang-interface{}与任意类型互转

参考文章

1. [go语言将函数作为参数传递](https://blog.csdn.net/eclipser1987/article/details/11772539)

```go
package main

import "fmt"

func main() {
	var name string
	var age int
	user := map[string]interface{}{
		"name": "general",
		"age":  21,
		"sayHello": func(name string) string {
			fmt.Printf("My name is %s\n", name)
			return name
		},
	}
	// 这种interface成员不能直接赋值, 因为类型不匹配
	// cannot use user["name"] (type interface {}) as type string in assignment:
	// need type assertion
	// name = user["name"]
	// 注意下面这种转换方法
	name = user["name"].(string)
	age = user["age"].(int)

	fmt.Printf("%s\n", name) // general
	fmt.Printf("%d\n", age)  // 21
	// interface转函数的方法, 目标函数的类型要有参数类型和返回值类型作为函数的类型标志(应该叫函数签名)
	user["sayHello"].(func(string) string)("newName") // My name is newName
}
```