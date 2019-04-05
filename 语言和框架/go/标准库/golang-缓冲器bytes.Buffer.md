# golang-bytes.Buffer

参考文章

1. [go语言的bytes.buffer](https://my.oschina.net/u/943306/blog/127981)

2. [Golang标准库之Buffer](http://blog.51cto.com/aresy/1405184)

Buffer结构体.

```go
type Buffer struct {
	buf       []byte   // contents are the bytes buf[off : len(buf)]
	off       int      // read at &buf[off], write at &buf[len(buf)]
	bootstrap [64]byte // memory to hold first slice; helps small buffers avoid allocation.
	lastRead  readOp   // last read operation, so that Unread* can work correctly.
}
```

`bytes.Buffer`是一个`[]byte`类型的缓冲器，这个缓冲器里的数据存放着都是`[]byte`.

缓冲器是一个容器, 其中存放着以`[]byte`类型存储的数据. 它占用的空间可变, 写入数据时会使其变大, 从其中读取数据会释放空间(可以对比TCP/IP协议栈中的缓冲区). 此种方法说是读, 不如说是"取出".

...不过看完关于buffer读取与写入的代码示例后你会发现, buffer更像是一个队列(后进先出).

## 1. 初始化

Buffer类型的初始化方式有两种

1. bytes.NewBuffer(buf []byte) *Buffer

2. bytes.NewBufferString(s string) *Buffer

它们的返回值相同, 但是传入的参数不同. 

```go
// 以下三者等效
buf1 := bytes.NewBufferString("hello")
buf2 := bytes.NewBuffer([]byte("hello"))
buf3 := bytes.NewBuffer([]byte{"h","e","l","l","o"})
// 以下两者等效
buf4 := bytes.NewBufferString("")
buf5 := bytes.NewBuffer([]byte{})
```

如果参数是nil的话，意思是new一个空的缓冲器. 就算在new的时候是空的也没关系，因为可以用Write方法来写入，写在缓冲区的尾部.

------

关于输出, buffer对象有一个`String()`方法, 可以把buffer中存放的数据转换`string`类型, 便于赋值. 打印的话, buffer对象可以直接打印, 如`fmt.Println(buf1)`.

```go
buf1 := bytes.NewBufferString("hello")
fmt.Println(buf1.String())
```

------

补充: `*Buffer`对象还有一个`Reset()`方法, 用于清空缓冲区.

## 2. 写入

有关缓冲区写入的方法, 有如下几个

1. Write func(p []byte) (n int, err error)          // 按`[]byte`类型写入

2. WriteByte func(c byte) error                     // 按`byte`类型写入(每次调用只能写入一个字节(字符类型即可用做字节))

3. WriteString func(s string) (n int, err error)    // 按`string`类型写入

这几个方法功能相似, 都是将新数据写入到缓冲区尾部, 区别只在于新数据是以什么样的形式传入. 而缓冲区本身占用空间会增大.

```go
	buf := bytes.NewBufferString("hello")
	buf.WriteByte(' ')
	// 单个字符即可用做字节类型
	buf.Write([]byte("world"))

	fmt.Println(buf)
```

## 3. 读取

有关缓冲区读取的方法, 有如下几个

1. Read func(p []byte) (n int, err error)

2. ReadByte func() (byte, error) // 无需参数, 直接返回缓冲区的第一个字节.

3. ReadBytes func(delim byte) (line []byte, err error)

4. ReadString func(delim byte) (line string, err error)

前面说了, `Read*`方法是从缓冲区取出数据, 会使缓冲区本身占用空间变小. 这些方法的使用方式也相似, 取出的数据都会放在传入的参数所表示的地址中. 传入的参数一般是有 **确定大小的容器([]byte数组)**.

一定要注意, 虽然创建缓冲区可以使用`make([]byte, 0)`作为空缓冲区, 但是从缓冲区读必须指定`[]byte`数组大小, 否则读取长度只能是0.

```go
	buf := bytes.NewBufferString("hello world!")
	fmt.Println(buf)
	// 创建存储容器
	c1 := make([]byte, 5)
	buf.Read(c1)           // 取出, c1定义了5字节的长度, 则只能装5个字节. 如果定义了20个字节, buf就空了.
	fmt.Println(c1)        // 这里打印出来的是字节数组, 要转换成字符串才行
	fmt.Printf("%s\n", c1) // hello
	fmt.Println(buf)       // 空格world
```

------

`ReadByte()`无需参数, 取出的值被当作返回值返回. `ReadBytes`与`ReadByte`完全不是一回事, 而是和`ReadString`更像.

`ReadBytes`与`ReadString`都需要传入一个分隔符参数, 返回缓冲区对象中分隔符前面的内容...参数传得都一样, 区别在于返回值类型不同.

分隔符可以是单个字符, 示例如下

```go
	buf1 := bytes.NewBufferString("hello world!")
	buf2 := bytes.NewBufferString("hello world!")
	fmt.Println(buf1)
	fmt.Println(buf2)
	var delim byte = ' '

	str1, _ := buf1.ReadBytes(delim)
	fmt.Println(str1) // 这里打印出的是字节数组

	str2, _ := buf2.ReadString(delim)
	fmt.Println(str2) // hello
```

## 4. 关于WriteTo与ReadFrom

这两个方法同样可以从缓冲区写入取出数据, 区别在于写入的目标容器与取出的源容器不同, 不是普通的字节数组, 而是另一个实现了`Reader`/`Writer`接口的对象, 比如另一个缓冲区, 或是**文件**, 一般后者更常用一点.

### 4.1 缓冲区互写

```go
	buf1 := bytes.NewBufferString("hello world!")
	buf2 := bytes.NewBufferString("")

	fmt.Println("初始化...")
	fmt.Printf("buf1: %s\n", buf1) // hello world!
	fmt.Printf("buf2: %s\n", buf2) // 空

	fmt.Println("buf1写入buf2...")
	buf1.WriteTo(buf2)

	fmt.Printf("buf1: %s\n", buf1) // 空
	fmt.Printf("buf2: %s\n", buf2) // hello world!

	fmt.Println("buf1从buf2读取...")
	buf1.ReadFrom(buf2)

	fmt.Printf("buf1: %s\n", buf1) // hello world!
	fmt.Printf("buf2: %s\n", buf2) // 空
```

读写文件的方法不在本文中介绍.