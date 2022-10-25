# golang-fmt交互式读入 bufio os.Stdin

参考文章

1. [Golang的交互模式进阶-读取用户的输入](https://www.cnblogs.com/yinzhengjie/p/7798498.html)
	- `fmt.Scan()`, `bufio.NewReader(os.Stdin)`, 
2. [golang 读取输入的几种写法](https://blog.csdn.net/wangxinwen/article/details/85040126)
	- `bufio.NewScanner(os.Stdin)`, `bufio.NewReader(os.Stdin)`, `os.Stdin.Read()`

`fmt.Scan`系就和C语言中的`scan/scanf`函数一样.

## fmt.Scan()/fmt.Scanf()

```go
	var name string
	var age int
	// Scanln 扫描来自标准输入的文本, 将空格分隔的值依次存放到后续的参数内, 直到碰到换行.
	// 注意需要是指针类型, 否则会报错: type not a pointer: string
	length, err := fmt.Scanln(&name, &age) 		// 输入: general 23
	if err != nil {
		fmt.Println(err)
		return
	}
	pattern := "read length: %d, name: %s, age: %d\n"
	// 注意这里的length值, 看起来是填充了length个参数的意思
	fmt.Printf(pattern, length, name, age) 		// read length: 2, name: general, age: 23
```

```go
	// length, err := fmt.Scanln(&name, &age)
	// 基本等价
	length, err := fmt.Scanf("%s %d", &name, &age)
```

> 如果不输入数据, 直接输入回车, 会得到`unexpected newline`错误, 可用此作为多次输入的结束标识.

## fmt.Sscan()/fmt.Sscanf()

ta们与`Scan()/Scanf()`的区别就如同`Sprintln()/Sprintf()`与`Println()/Printf()`的区别一样, 就是从已知的字符串变量中读取数据.

```go
	var name string
	var age int
	var knownString = "general 23 中国"
	// Scanln 扫描已知字符串, 将空格分隔的值依次存放到后续的参数内, 直到碰到换行. 多余的内容会忽略.
	// 注意需要是指针类型, 否则会报错: type not a pointer: string
	length, err := fmt.Sscan(knownString, &name, &age)
	if err != nil {
		fmt.Println(err)
		return
	}
	pattern := "read length: %d, name: %s, age: %d\n"
	fmt.Printf(pattern, length, name, age) 		// read length: 2, name: general, age: 23
```

```go
	// length, err := fmt.Sscan(knownString, &name, &age)
	// 基本等价
	length, err := fmt.Sscanf(knownString, "%s %d", &name, &age)
```

## bufio.NewReader(os.Stdin)

```go
	// 创建一个读取器, 并将其与标准输入绑定
	inputReader := bufio.NewReader(os.Stdin)
	// 从输入中读取内容, 直到碰到 delim 字符, 然后将读取到的内容连同 delim 字符**一起放到缓冲区**
	str, err := inputReader.ReadString('\n')	// 输入: hello general
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)							// hello general 注意这个字符串里包含了换行符
```

由于`bufio.NewReader()`无法实现格式化读入, 只能直接读取一整行, 所以不建议使用.

## bufio.NewScanner(os.Stdin)

```go
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please input: ")
	scanner.Scan()											// 输入: general 23
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner failed: %s", err)
	}
	fmt.Printf("bufio.NewScanner: %s\n", scanner.Text())	// general 23
```

同样无法指定读入格式, 不建议使用.

## os.Stdin.Read()

```go
	buf := make([]byte, 512)
	length, err := os.Stdin.Read(buf)	// 输入: general 23
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf[0:length]) 			// [103 101 110 101 114 97 108 32 50 51 10]
	fmt.Printf("%s", buf[0:length]) 	// general 23
```

同上.
