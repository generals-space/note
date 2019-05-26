# golang-switch与type类型断言

```go
package main

import "fmt"

func main() {
	// i必须为interface{}类型, 无法是具体类型.
	// 否则switch i.(type)会出错.
	// var i interface{} = 1
	var i interface{} = map[string]string{}

	// 无法在switch之外使用.(type)断言
	// use of .(type) outside type switch
	// fmt.Println(i.(type))

	switch i.(type) {
	case int:
		fmt.Println(i, "is an int value.")
	case string:
		fmt.Println(i, "is a string value.")
	case int64:
		fmt.Println(i, "is an int64 value.")
	case bool:
		fmt.Println(i, "is an bool value.")
	case interface{}:
		fmt.Println(i, "is an interface{} value.")
	default:
		fmt.Println(i, "is an unknown type.")
	}
}
```
