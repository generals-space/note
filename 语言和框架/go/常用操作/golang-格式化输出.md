# golang-格式化输出

参考文章

1. [Go 字符串格式化](https://studygolang.com/articles/1915)

2. [基础知识 - Golang 中的格式化输入输出](https://www.cnblogs.com/golove/p/3284304.html)

```go
package main
import "fmt"
type point struct {x, y int}

func main() {
    p := point{1, 2}

    // %v可以打印结构体的对象的值...只有值
    fmt.Printf("%v\n", p)

    // `%+v`可以打印结构体的键和值
    fmt.Printf("%+v\n", p)

    // `%#v`格式化输出将输出一个值的Go语法表示方式。
    fmt.Printf("%#v\n", p)

    // `%T`输出一个值的数据类型
    fmt.Printf("%T\n", p)

    // `%t`格式化布尔型变量
    fmt.Printf("%t\n", true)

    // 格式化整型
    fmt.Printf("%d\n", 123)

    // 这种方式输出整型的二进制表示方式
    fmt.Printf("%b\n", 14)

    // %c打印单个字符, 也可以打印ascii码中某个整型对应的字符
    fmt.Printf("%c\n", 'a')
    fmt.Printf("%c\n", 97)

    // 使用`%x`输出一个值的16进制表示方式
    fmt.Printf("%x\n", 456)

    // 浮点型数值也有几种格式化方法。最基本的一种是`%f`
    fmt.Printf("%f\n", 78.9)

    // `%e`和`%E`使用科学计数法来输出整型
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)

    // 使用`%s`输出基本的字符串
    fmt.Printf("%s\n", "\"string\"")

    // 输出像Go源码中那样带双引号的字符串，需使用`%q`
    fmt.Printf("%q\n", "\"string\"")

    // `%x`以16进制输出字符串，每个字符串的字节用两个字符输出
    fmt.Printf("%x\n", "hex this")

    // 使用`%p`输出一个指针的值
    fmt.Printf("%p\n", &p)

    // 当输出数字的时候，经常需要去控制输出的宽度和精度。
    // 可以使用一个位于%后面的数字来控制输出的宽度，默认
    // 情况下输出是右对齐的，左边加上空格
    fmt.Printf("|%6d|%6d|\n", 12, 345)

    // 你也可以指定浮点数的输出宽度，同时你还可以指定浮点数
    // 的输出精度
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    // To left-justify, use the `-` flag.
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

    // 你也可以指定输出字符串的宽度来保证它们输出对齐。默认
    // 情况下，输出是右对齐的
    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    // 为了使用左对齐你可以在宽度之前加上`-`号
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
}
```

## 关于16进制的输出

使用`#`标记可以输出`0x`字样

```go
fmt.Printf("%#x\n", 1)      // 0x1
fmt.Printf("%#x\n", 255)    // 0xff
fmt.Printf("%#x\n", 256)    // 0x100
```

使用`0n`(`n`为大于0的整数)可以定义每一位16进制的宽度, 用于保持格式的一致

```go
fmt.Printf("%02x\n", 1)     // 01
fmt.Printf("%02x\n", 255)   // ff
fmt.Printf("%02x\n", 256)   // 100 ...这个好像不是02的锅, 因为单个16进制数最大值为255
fmt.Printf("%02x\n", []byte{1, 2, 254, 255}) // 0102feff
```

当然, 两种标记可以配合使用

```go
fmt.Printf("%#02x\n", []int{1, 2, 254, 255}) // [0x01 0x02 0xfe 0xff]
```