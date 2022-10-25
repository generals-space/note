# golang-text编码转换

参考文章

1. [Golang的字符编码介绍](https://www.cnblogs.com/yinzhengjie/p/7956689.html)
2. [Golang GBK To Utf-8](https://blog.csdn.net/a99361481/article/details/83273053)
3. [goquery 增加GBK支持](https://blog.csdn.net/jrainbow/article/details/52712685)

[iconv-go](github.com/djimenez/iconv-go)使用了cgo, 需要gcc环境, 在win下无法通过`go get`安装. 一般情况下

```go
package main

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
)

// CharsetMap 字符集映射
var CharsetMap = map[string]encoding.Encoding{
	"utf-8":   unicode.UTF8,
	"gbk":     simplifiedchinese.GBK,
	"gb2312":  simplifiedchinese.GB18030,
	"gb18030": simplifiedchinese.GB18030,
	"big5":    traditionalchinese.Big5,
}
```

指定编码读写的方式有两种, 分别是各字符集对象的`NewEncoder()`和`NewDecoder()`方法, 和`x/text/transform`包中的`NewReader()`方法.

下面分别用这两种方式实现了`DecodeToUTF8()`和`EncodeFromUTF8()`函数. 

## 1. transform包

`transform`包是专门用于进行字符集转换的, 使用起来非常简单.

```go
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// DecodeToUTF8 从输入的byte数组中按照指定的字符集解析出对应的utf8格式的内容并返回.
func DecodeToUTF8(input []byte, charset encoding.Encoding) (output []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(input), charset.NewDecoder())
	output, err = ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	return
}

// EncodeFromUTF8 将输入的utf-8格式的byte数组中按照指定的字符集编码并返回
func EncodeFromUTF8(input []byte, charset encoding.Encoding) (output []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(input), charset.NewEncoder())
	output, err = ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	return
}

func main() {
	charset := simplifiedchinese.GB18030
	sourceBytes := []byte("这段内容是要被编码转换")

	// 编码操作
	encodedBytes, _ := EncodeFromUTF8(sourceBytes, charset)
	log.Printf("编码后的内容: %s", encodedBytes) // 编码后的内容 ����������Ҫ������ת��

	// 以指定编码写入文件, 打开文件需要指定gbk编码, 否则看到的是乱码.
	file, err := os.Create("file.txt")
	defer file.Close()
	_, err = file.Write(encodedBytes)
	if err != nil {
		log.Printf("写入文件失败: %s", err.Error())
	}

	// 解码操作
	decodedBytes, _ := DecodeToUTF8(encodedBytes, charset)
	log.Printf("编码前的内容: %s", decodedBytes) // 编码前的内容: 这段内容是要被编码转换
}

```

## 2. 字符集对象接口

手写的话就比较麻烦了, 不推荐使用这种.

```go
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// DecodeToUTF8 从输入的byte数组中按照指定的字符集解析出对应的utf8格式的内容并返回.
func DecodeToUTF8(input []byte, charset encoding.Encoding) (output []byte, err error) {
	reader := bytes.NewReader(input)
	// 包装reader, 返回的新reader
	charsetReader := charset.NewDecoder().Reader(reader)
	// 如果你知道reader中内容的长度, 比如`reader.Len()`可以得到,
	// 可以直接使用charsetReader.Read(buf)读取到[]byte数组buf中.
	// 如果不知道, 就使用ioutil.ReadAll(), 反正charsetReader本身也是一个reader可读对象.
	output, err = ioutil.ReadAll(charsetReader)
	if err != nil {
		log.Printf("读取内容失败: %s", err.Error())
		return
	}
	return
}

// EncodeFromUTF8 将输入的utf-8格式的byte数组中按照指定的字符集编码并返回
func EncodeFromUTF8(input []byte, charset encoding.Encoding) (output []byte, err error) {
	// 创建空缓冲区, 以下两种方式等效.
	// buffer := bytes.NewBuffer([]byte{})
	buffer := bytes.NewBuffer(make([]byte, 0))
	encoder := charset.NewEncoder()
	// 包装writer, 返回的新writer可以在写入数据前对其进行编码
	charsetWriter := encoder.Writer(buffer)
	// 待编码内容必须要在用字符集对象的Writer方法包装后写入才有效, 原buffer中的内容将不会被编码.
	// 实际的写入对象仍是buffer, 可以通过`buffer.Bytes()`取得其中的内容.
	charsetWriter.Write(input)
	output = buffer.Bytes()
	return
}

func main() {
	charset := simplifiedchinese.GB18030
	sourceBytes := []byte("这段内容是要被编码转换")

	// 编码操作
	encodedBytes, _ := EncodeFromUTF8(sourceBytes, charset)
	log.Printf("编码后的内容: %s", encodedBytes) // 编码后的内容 ����������Ҫ������ת��

	// 以指定编码写入文件, 打开文件需要指定gbk编码, 否则看到的是乱码.
	file, err := os.Create("file.txt")
	defer file.Close()
	_, err = file.Write(encodedBytes)
	if err != nil {
		log.Printf("写入文件失败: %s", err.Error())
	}

	// 解码操作
	decodedBytes, _ := DecodeToUTF8(encodedBytes, charset)
	log.Printf("编码前的内容: %s", decodedBytes) // 编码前的内容: 这段内容是要被编码转换
}

```
