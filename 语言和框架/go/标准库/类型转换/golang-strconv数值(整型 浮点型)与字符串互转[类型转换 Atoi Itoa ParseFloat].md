# golang-strconv数值(整型 浮点型)与字符串互转[类型转换 Atoi Itoa ParseFloat]

参考文章

1. [go语言int类型转化成string类型的方式](https://blog.csdn.net/love_se/article/details/7947511)
2. [如何在Golang中将类型从字符串转换为 float64 i？](https://cloud.tencent.com/developer/ask/44574/answer/70011)

## 整型 <-> 字符串

整型转字符串只需要用`Sprintf()`方法就可以.

而字符串转整型需要借助`strconv`标准库.

```go
	strA := "123"
	intA, _ := strconv.Atoi(strA)
	log.Printf("%d\n", intA)
```

`strconv.Atoi func(s string) (int, error)`: ASCII -> Int, 字符串转整型

`strconv.Itoa func(i int) string`: Int -> ASCII, 整型转字符串

## 浮点型 <-> 字符串

```go
a := "77.9285731709928"
b := "250"

a1, _ := strconv.ParseFloat(a, 64)
b1, _ := strconv.ParseFloat(b, 64)

fmt.Printf("%f\n", a1)	// 77.928573
fmt.Printf("%f\n", b1)	// 250.000000
```

`strconv`库还有`ParseBool`, `ParseInt`, `ParseUint`等方法, 记录一下.
