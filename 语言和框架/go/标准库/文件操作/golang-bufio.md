# golang-bufio

参考文章

1. [Golang学习 - bufio 包](http://www.cnblogs.com/golove/p/3282667.html)

网上关于`bufio`的说明一般都是 **bufio 包实现了带缓存的 I/O 操作**, 说实话我没太懂`ioutil`, `bytes`和`bufio`在读取方法上的差异.

其实查看一下`bufio`的源码就会发现, ta的重点在于后半句: `provides buffering and some help for textual I/O`.

也就是说, 在缓冲区IO方法, ta和`bytes`库相似, 都对普通Reader进行的包装. ta更重要的是实现了一些对文本信息的处理函数, 这使得我们在处理文本文件(文本字符串貌似没什么帮助?)时, 更加方便.

## Peek()

`Peek(n)`返回缓存的一个切片, 该切片引用缓存中前`n`个字节的数据, 该操作不会将数据读出, 只是引用, 引用的数据在下一次读取操作之前是有效的. 如果切片长度小于`n`, 则返回一个错误信息说明原因(比如EOF).

```go
	strReader := strings.NewReader("hello world")
	reader := bufio.NewReader(strReader)

	byteArray, err := reader.Peek(100)
	if err != nil {
		log.Println(err) // EOF
	}
	log.Printf("%s", byteArray) // hello world
```

## Reader

bufio中文本读取常用的函数有

1. `ReadLine()`
2. `ReadBytes()`
3. `ReadString()`
4. `ReadSlice()`

如果以普通文本文件的读取操作为例, `ReadLine()`基本上等同于`ReadBytes('\n')`, `ReadString('\n')`和`ReadSlice('\n')`.

但是`ReadLine`不会读取行尾的换行符`\n`, 另外3个会. 并且如果使用另外3个函数, 文本不以空行结尾, 则最后一行的数据读取不出来.

```go
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		log.Printf("%s\n", line)
	}
```

`Buffered()`不会用, 每次都返回0, 不知道什么情况下才有效.
