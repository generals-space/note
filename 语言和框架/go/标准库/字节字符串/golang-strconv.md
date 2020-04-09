# golang-strconv进制转换

<!tags!>: <!进制转换!> <!类型转换!>

参考文章

1. [Go 进制转换](https://my.oschina.net/tsh/blog/1619887)

2. [Go语言---strconv包](https://blog.csdn.net/li_101357/article/details/80252653)

引用参考文章1中的说法

`strconv`包括四类函数

1. `Format`类, 例如`FormatBool(b bool) string`, 将`bool`, `float`, `int`, `uint`类型的转换为`string`, **`FormatInt`的缩写为`Itoa`**;

2. `Parse`类, 例如`ParseBool(str string)(value bool, err error)`将字符串转换为`bool`, `float`, `int`, `uint`类型的值, `err`指定是否转换成功, **`ParseInt`的缩写是`Atoi`**;

3. `Append`类, 例如`AppendBool(dst []byte, b bool) []byte`, 将值转化后添加到`[]byte`的末尾;

4. `Quote`类, 对字符串的 双引号 单引号 反单引号 的操作;

## 1. `Format`与`Parse`类函数

这两组函数的用法比较简单, 也容易理解.

### 1.1 `Format`

`Format`族可以把多个类型的变量转换成string类型, 在转换过程中可以设置进制, 浮点精度, 小数位数等.

```go
	// 转换为16进制和2进制, FormatUint也是同样的用法
	str2 := strconv.FormatInt(int64(123456789), 2)
	log.Println(str2) // 111010110111100110100010101
	str16 := strconv.FormatInt(int64(123456789), 16)
	log.Println(str16) // 75bcd15

	log.Println(strconv.FormatBool(true))	// "true"
    log.Println(strconv.FormatBool(false))	// "false"
    
	oriFloat := float64(123.456)
	log.Println(strconv.FormatFloat(oriFloat, 'f', 6, 32)) // 123.456001
	log.Println(strconv.FormatFloat(oriFloat, 'e', 6, 32)) // 1.234560e+02
```


其中`FormatInt`与`FormatUint`, `FormatBool`都比较好理解, 在这些应用场景中, `fmt.Sprintf()`格式化输出其实也能达到目的.

较为复杂的是`FormatFloat`函数.

```go
// FormatFloat 将浮点数 f 转换为字符串形式
// f：要转换的浮点数
// fmt：格式标记（b、e、E、f、g、G）
// prec：精度（小数部分的长度，不包括指数部分）
// bitSize：指定浮点类型（32:float32、64:float64），结果会据此进行舍入。
//
// 格式标记：
// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)    // 嗯, 大E与小e的区别只在于指数前的字母e的大小写, 没有其他区别.
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)
//
// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）

// 参考格式化输入输出中的旗标和精度说明
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
```

### 1.2 `Parse`

## 2. `Append`

这个函数可以把各种类型的变量添加到`[]byte`数组中, 但是ta转换的方式让人哭笑不得, 个人感觉没什么实际价值.

```go
	intArray := strconv.AppendInt(nil, int64(20), 16)
	log.Println([]byte("14")) // 输出: [49 52]
	log.Println(intArray)     // 输出: [49 52]

	boolArray := strconv.AppendBool(nil, true)
	log.Println([]byte("true")) // 输出: [116 114 117 101]
	log.Println(boolArray)      // 输出: [116 114 117 101]

	floatArray := strconv.AppendFloat(nil, float64(123.456), 'f', 6, 64)
	log.Println(floatArray)	// [49 50 51 46 52 53 54 48 48 48]
	log.Println(string(floatArray))	// 123.456000

	quoteArray := strconv.AppendQuote(nil, "hello")
	log.Println([]byte("hello")) // 输出: [104 101 108 108 111]
	log.Println(quoteArray)      // 输出: [34 104 101 108 108 111 34]
```

没错, `Append`的函数都是对`Format`函数做了个包装, 而且很过分的把结果按照字符逐一转换成ASCII字符...真是简直了...

还有`AppendQuote`这个函数...给目标字符串加上首尾两个双引号, 好有意义啊...

## 3. `Quote`

这类函数用得不多, 没仔细看, 感觉很鸡肋...具体示例可以见参考文章2.

```go
	log.Println("hello world")                // hello world
	log.Println(strconv.Quote("hello world")) // "hello world"
	log.Println("世界你好")                       // 世界你好
	log.Println(strconv.Quote("世界你好"))        // "世界你好"

	log.Println(strconv.QuoteToASCII("hello world")) // "hello world"
	log.Println(strconv.QuoteToASCII("世界你好"))        // "\u4e16\u754c\u4f60\u597d"

	log.Println('好')                           // 22909
	log.Println("好")                           // 好
	log.Println(strconv.QuoteRune('好'))        // '好'
	log.Println(strconv.QuoteRuneToASCII('好')) // '\u597d'

	str := "\"hello world\""
	log.Println(str) // "hello world"
	uq, _ := strconv.Unquote(str)
	log.Println(uq) // hello world
```
