# golang-strconv数值与[]byte互转

参考文章

1. [Go 进制转换](https://my.oschina.net/tsh/blog/1619887)

在做嵌入式相关的数据处理时, 经常遇到整型数值与[]byte数组之间相互转换的场景, 而且还会涉及到大小端序的问题.

go提供了这方面相关的工具: `encoding/binary`库.

## 1. 数值 -> []byte

如下展示了将整型转化为`[]byte`的最简方法

```go
func main() {
	buf := bytes.NewBuffer([]byte{})

	// 单字节的数据
	// 大端序, 高字节在前
	binary.Write(buf, binary.BigEndian, int8(127))
	log.Println(buf.Bytes()) // [127]
	buf.Reset()
	binary.Write(buf, binary.BigEndian, int8(-120))
    log.Println(buf.Bytes()) // [136]
    buf.Reset()
	binary.Write(buf, binary.BigEndian, uint8(255))
	log.Println(buf.Bytes()) // [255]

}
```

由于`int8`可表示有符号整数, 所以ta其中的值的范围为`127到-128`之间, 但转换成`[]byte`后最高可以表示到255.

另外, `binary.Write()`方法的第3个参数必须为`int8`, `int16`, `int32`或`int64`, 及ta们对应的无符号类型, 不接受`int`类型的参数. 我觉得ta们直接决定了添加到buf里的字节的个数, 不足就补0 (另外试了下, 貌似字符串和布尔类型的变量也是不可以的).

多字节操作如下

```go
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, int16(3600))

	log.Println(buf.Bytes()) // [14 16]
	log.Println(14<<8 + 16)  // 3600
```

## 2. []byte -> 数值

```go
	buf := bytes.NewBuffer([]byte{14, 16})
	log.Println(buf.Bytes())		// [14 16]

	var data int16
	binary.Read(buf, binary.BigEndian, &data)

	log.Println(data)			// 3600
	log.Println(buf.Bytes())	// []
```

`binary.Read()`的第3个参数为`int8`, `int16`, `int32`或`int64`, 及ta们对应的无符号类型, ...的**指针**. ta们的长度决定了从buf中读取的字节个数. 

证据如下

```go
	buf := bytes.NewBuffer([]byte{16, 14, 16})
	log.Println(buf.Bytes())		// [16 14 16]

	var data1 int8
	binary.Read(buf, binary.BigEndian, &data1)
	log.Println(data1)			// 16
	log.Println(buf.Bytes())	// [14 16]

	var data2 int16
	binary.Read(buf, binary.BigEndian, &data2)
	log.Println(data2)			// 3600
	log.Println(buf.Bytes())	// []
```