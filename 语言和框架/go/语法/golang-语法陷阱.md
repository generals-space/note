# golang-语法陷阱

参考文章

1. [Golang开发新手常犯的50个错误](https://blog.csdn.net/gezhonglei2007/article/details/52237582)

## 1. 不允许出现未使用的import

解决方法：使用`_`作为引入包别名

```go
package main

import (  
    _ "fmt"         // 指定别名为`_`, 就算程序中没有用到这个包, 也不会出错
    "log"
    "time"
)

var _ = log.Println // 变量名为`_`

func main() {  
    _ = time.Now
}
```

## 2. 数组用于函数传参时是值复制

方法法或函数调用时，传入参数都是值复制（跟赋值一致）, 除非是`map`、`slice`、`channel`、`指针`类型这些特殊类型是引用传递.

```go
x := [3]int{1,2,3}

// 数组在函数中传参是值复制
func(arr [3]int) {
    arr[0] = 7
    fmt.Println(arr) //prints [7 2 3]
}(x)
fmt.Println(x)       //prints [1 2 3] (not ok if you need [7 2 3])

// 使用数组指针实现引用传参
func(arr *[3]int) {
    (*arr)[0] = 7
    fmt.Println(arr) //prints &[7 2 3]
}(&x)
fmt.Println(x)       //prints [7 2 3]
```
