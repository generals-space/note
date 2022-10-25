# golang-string与byte数组互转[类型转换 Buffer]

参考文章

1. [go语言int类型转化成string类型的方式](https://blog.csdn.net/love_se/article/details/7947511)
2. [如何在Golang中将类型从字符串转换为 float64 i？](https://cloud.tencent.com/developer/ask/44574/answer/70011)

在Go当中string底层是用`[]byte`存的, 并且是不可以改变的(如果要修改string内容需要将string转换为`[]byte`或`[]rune`, 并且修改后的string内容是重新分配的).

`string` -> `[]byte`: `[]byte(str)`

`string` -> `[]byte`: `[]byte{'h', 'e', 'l', 'l', 'o'}`

`[]byte` -> `string`: `string(byte变量)`

## 使用 bytes.Buffer

`string` -> `[]byte`

```go
	buf := bytes.NewBufferString("hello")
	buf.WriteByte(' ')
	// 单个字符即可用做字节类型
	buf.Write([]byte("world"))
	log.Printf("%d\n", buf.Bytes())
```

`[]byte` -> `string`

```go
	buf := bytes.NewBuffer([]byte{"h","e","l","l","o"})
	buf.WriteByte(' world')
	log.Printf("%s\n", buf.String())
```
