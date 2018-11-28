# golang-函数作为参数的使用方法

参考文章

1. [go语言将函数作为参数传递](https://blog.csdn.net/eclipser1987/article/details/11772539)

```go
package main

import "fmt"

func goFunc(f interface{}, args ...interface{}) {
	if len(args) > 1 {
		f.(func(...interface{}))(args)
	} else if len(args) == 1 {
		f.(func(interface{}))(args[0])
	} else {
		f.(func())()
	}
}

func f1() {
	fmt.Println("f1 done")
}

func f2(i interface{}) {
	fmt.Println("f2 done", i)
}

func f3(args ...interface{}) {
	fmt.Printf("%T\n", args) // []interface{}
	fmt.Println("f3 done", args)
}

func main() {
	// 通过goFunc传入目标函数代为执行
	goFunc(f1)
	goFunc(f2, "hello")
	goFunc(f3, "hello world", 3.14)
}
```