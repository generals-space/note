# golang-判断变量类型

参考文章

1. [Go语言中怎样判断数据类型](http://blog.sina.com.cn/s/blog_487109d101013g2p.html)
2. [Go 字符串格式化](https://studygolang.com/articles/1915)

## 1. 反射

根据参考文章1的示例, 标准库`reflect`用起来很方便.

```go
package main

import (
    "log"
    "reflect"
)
type point struct {
    x, y int 
}

func main() {
    var x float64 = 3.4 
    var p = point{x: 1, y: 3}
    log.Println("type of x: ", reflect.TypeOf(x))
    log.Println("type of p: ", reflect.TypeOf(p))
}
```

输出为

```log
2018/04/26 10:30:08 type of x:  float64
2018/04/26 10:30:08 type of p:  main.point
```

## 2. 格式化输出

按照参考文章2的示例, golang格式化输出函数`Printf`可以使用`%T`占位符直接输出变量类型(结构体类型也可以哦).

```go
package main
import "log"

type point struct {
    x, y int
}

func main() {
    x := 3.4
    p := point{1, 3}

    log.Printf("type of x: %T", x)
    log.Printf("type of p: %T", p)
}
```

输出同上, 感觉第2种更方便...
