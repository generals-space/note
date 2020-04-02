# golang-格式化输出

参考文章

1. [Go 字符串格式化](https://studygolang.com/articles/1915)

2. [基础知识 - Golang 中的格式化输入输出](https://www.cnblogs.com/golove/p/3284304.html)

```go
type point struct {x, y int}
var p = point{1, 2}
```

- `%v`: 
    - ``: `

```go
// %v 可以打印结构体的对象的值...只有值
fmt.Printf("%v\n", p)                           // {1 2}

// %+v 可以打印结构体的键和值
fmt.Printf("%+v\n", p)                          // {x:1 y:2}

// %#v 可以打印一个值的Go语法表示方式. 
fmt.Printf("%#v\n", p)                          // main.point{x:1, y:2}

// 使用`%p`输出一个指针的值
fmt.Printf("%p\n", &p)                          // 0xc0000140b0

// %T 打印一个值的数据类型(int, string, bool, map[string][string]等)
fmt.Printf("%T\n", p)                           // main.point

// %t 打印布尔变量值
fmt.Printf("%t\n", true)                        // true
// %s 打印基本的字符串
fmt.Printf("%s\n", "\"string\"")                // "string"
// %d 打印整型
fmt.Printf("%d\n", 123)                         // 123

// %c 打印单个字符, 也可以打印ascii码中某个整型对应的字符
fmt.Printf("%c\n", 'a')                         // a
fmt.Printf("%c\n", 97)                          // a

// %b 打印整型数值的二进制表示方式
fmt.Printf("%b\n", 14)                          // 1110
// %x 输出一个值的16进制表示方式
fmt.Printf("%x\n", 456)                         // 1c8
// %x 也可以输出字符串的16进制表示, 每个字符串的字节用两个字符输出
fmt.Printf("%x\n", "hex this")                  // 6865782074686973

// %f 浮点型数值最基本的一种格式化方法(默认精度为6). 
fmt.Printf("%f\n", 78.9)                        // 78.900000

// %e 和 %E 使用科学计数法来输出整型(好像没什么区别)
fmt.Printf("%e\n", 123400000.0)                 // 1.234000e+08
fmt.Printf("%E\n", 123400000.0)                 // 1.234000E+08

// %q 打印像Go源码中那样带双引号的字符串(???)
fmt.Printf("%q\n", "\"string\"")                // "\"string\""

///////////////////////////////////////////////////////////////////////
// 输出的宽度和精度. 
// 可以使用一个位于%后面的数字来控制输出的宽度, 默认
// 情况下输出是右对齐的, 左边加上空格
fmt.Printf("|%6d|%6d|\n", 12, 345)              // |    12|   345|
// `-` 符号 左对齐(默认情况下, 输出是右对齐的)
fmt.Printf("|%-6d|%-6d|\n", 12, 345)            // |12    |345   |
// 宽度前加0可以补全(其他字符都不好使)
// 仅限左侧补0, 因为右侧补0相当于改变了数值大小.
fmt.Printf("|%06d|%6d|\n", 12, 345)             // |000012|   345|

// 可以指定浮点数的输出宽度与精度
fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)        // |  1.20|  3.45|
fmt.Printf("|%06.2f|%6.2f|\n", 1.2, 3.45)       // |001.20|  3.45|
// 同样可以使用`-`符号
fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)      // |1.20  |3.45  |

// 指定字符串的宽度(同样默认右对齐). 
fmt.Printf("|%6s|%6s|\n", "foo", "b")           // |   foo|     b|
// 同样可以使用`-`符号
fmt.Printf("|%-6s|%-6s|\n", "foo", "b")         // |foo   |b     |
// 字符串是没有补0这一说的
fmt.Printf("|%-06s|%-6s|\n", "foo", "b")        // |foo   |b     |
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

不过如果你想读取`0xff`这样的输入的话, 使用`%#x`是不行的.

```go
	var age int
	length, err := fmt.Scanf("%#x", &age)
```

```
bad verb '%#' for integer
```

还是直接用`0x%x`取巧一点吧.
