# golang-类型转换

参考文章

1. [go语言int类型转化成string类型的方式](https://blog.csdn.net/love_se/article/details/7947511)

2. [如何在Golang中将类型从字符串转换为 float64 i？](https://cloud.tencent.com/developer/ask/44574/answer/70011)

主要还是`[]byte`和`字符串`, `缓冲区buffer`之间的转换.

在Go当中string底层是用[]byte存的, 并且是不可以改变的(如果要修改string内容需要将string转换为[]byte或[]rune，并且修改后的string内容是重新分配的).

一般来说, buffer类型变量都实现了`io.Reader/Writer`接口.


string -> `[]byte`: `[]byte(str)`

string -> `[]byte`: `[]byte{'h', 'e', 'l', 'l', 'o'}`

`[]byte` -> string: `string(byte变量)`

------

buffer与string之间貌似没有直接转换的关系, 因为buffer其实是一个结构体, 并且可能由于拥有`Read`或`Write`方法而实现了`Reader`或`Writer`接口. 而buffer结构体中实现存储数据的成员是`[]byte`...

所以buffer的各种操作都是基于字节数组的.

buffer结构在`bytes`标准库中定义. 

------

数值与字符串之间互转

数值转字符串只需要用`Sprintf()`方法就可以.

而字符串转数值需要借助`strconv`标准库.

```go
	strA := "123"
	intA, _ := strconv.Atoi(strA)
	log.Printf("%d\n", intA)
```

`strconv.Atoi func(s string) (int, error)`: ASCII -> Int, 字符串转整型

`strconv.Itoa func(i int) string`: Int -> ASCII, 整型转字符串

------

浮点型与字符串之间互转

```go
a := "77.9285731709928"
b := "250"

a1, _ := strconv.ParseFloat(a, 64)
b1, _ := strconv.ParseFloat(b, 64)

fmt.Printf("%f\n", a1)	// 77.928573
fmt.Printf("%f\n", b1)	// 250.000000
```

`strconv`库还有`ParseBool`, `ParseInt`, `ParseUint`等方法, 记录一下.