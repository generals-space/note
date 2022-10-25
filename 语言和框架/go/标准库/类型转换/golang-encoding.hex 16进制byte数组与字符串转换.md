# golang-16进制与字符串转换

参考文章

1. [16进制到文本字符串的转换，在线实时转换](https://www.bejson.com/convert/ox2str/)
2. [Golang中的[]byte与16进制(String)之间的转换](https://blog.csdn.net/jason_cuijiahui/article/details/79418557)

本文涉及到的字符串, 16进制变量, 16进制字符串和`[]byte`数组转换.

- `hello world`
- `[104 101 108 108 111 32 119 111 114 108 100]`
- `68656c6c6f20776f726c64`
- `[]byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}`

------

首先明确一点, 单个`int`类型与单个16进制的变量都是一个字节, 不过单个`int`值与单个16进制变量的上限都是一样的(65535吧).

```go
a := 10
fmt.Printf("%d\n", unsafe.Sizeof(a))    // 8
b := 0xA
fmt.Printf("%d\n", unsafe.Sizeof(b))    // 8
c := 0xAAAA
fmt.Printf("%d\n", unsafe.Sizeof(c))    // 8
```

各种转换方法如下示例.

```go
	// str也可以是中文
	str := "hello world"
	// string转[]byte数组和16进制字符串
	fmt.Printf("%d\n", []byte(str)) // [104 101 108 108 111 32 119 111 114 108 100]
	fmt.Printf("%x\n", str)         // 68656c6c6f20776f726c64

	// []byte数组转字符串, 这个很简单.
	// 要注意这个数组的定义方式, 可不是[]byte{1, 2, 3...}哦, 而是[]byte{104, 101, 108...}
	array := []byte{0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}
	fmt.Printf("%s\n", array) // hello world

	// 16进制字符串转[]byte数组, 然后可以通过[]byte数组转原始字符串
	byteArray, _ := hex.DecodeString("68656c6c6f20776f726c64")
	fmt.Printf("%d\n", byteArray) // [104 101 108 108 111 32 119 111 114 108 100]
	for _, byt := range byteArray {
        fmt.Printf("0x%x, ", byt)
    }		// 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 
    fmt.Println()
	fmt.Printf("%s\n", byteArray) // hello world

	// asicc字符串转16进制字符串
	hexStr := fmt.Sprintf("%x", str)
	fmt.Printf("%s\n", hexStr) // 68656c6c6f20776f726c64
	// hex包里自带的EncodeToString()方法其实和上面的Sprintf()作用相同
	hexStr2 := hex.EncodeToString([]byte(str))
	fmt.Printf("%s\n", hexStr2) // 68656c6c6f20776f726c64

```

`strconv`包的`Itoa`与`Atoi`方法其实不适合16进制的转换, 因为它们面向的对象是`int`类型, 而16进制变量有`ABCDEF`字符, 会出错的.
