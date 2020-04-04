# golang-string, byte与rune的关系与区别

参考文章

1. [string rune byte 的关系](https://www.golangtc.com/t/528cc004320b52227200000f)

2. [golang byte和rune的区别](https://blog.csdn.net/cscrazybing/article/details/79107412)

20200403 更新

今天做了一道题, 需要遍历一个`string`, 最开始写的如下代码

```go
	// 其他语言中的单个字符 char, 在 golang 只能用 int 表示.
	theMap := make(map[int]int)
	for _, char := range inputStr {
		fmt.Println(i, char)
		// cannot use char (type rune) as type int in map index
		if _, ok := theMap[char]; !ok {
			// 还没有此字母的键
			// cannot use char (type rune) as type int in map index
			theMap[char] = 1
		}
	}
```

时间久了都忘了这`rune`这事了, 由于输入的都是英文字母, 不需要考虑汉字的问题, 所以可以不用这么写.

```go
	// 其他语言中的单个字符 char, 在 golang 只能用 byte 表示...
	theMap := make(map[byte]int)
	for i := 0; i < len(inputStr); i++ {
		// fmt.Println(inputStr[i])
		if _, ok := theMap[inputStr[i]]; !ok {
			// 还没有此字母的键
			theMap[inputStr[i]] = 1
		} else {
			theMap[inputStr[i]] = theMap[inputStr[i]] + 1
		}
	}
```

------

在Go当中`string`底层是用`[]byte`存的, 并且是不可以改变的(如果要修改string内容需要将string转换为[]byte或[]rune，并且修改后的string内容是重新分配的).

```go
s := "Go编程"
fmt.Println(len(s))
```

输出结果应该是8, 因为中文字符是用3个字节存的. 

如果想要获得我们想要的情况的话，需要先转换为rune切片再使用内置的len函数. 在使用方法上, `byte`与`rune`没有任何区别.

```go
fmt.Println(len([]rune(s)))
```

结果就是4了。

所以用string存储unicode的话，如果有中文，按下标是访问不到的，因为你只能得到一个byte。 要想访问中文的话，还是要用rune切片，这样就能按下表访问。

参考文章2中介绍了`byte`与`rune`的区别: `byte`是`uint8`的别名, 而`rune`是`uint32`的别名.

所以一般在涉及到包含中文字符串的切片操作时, 尽量先把字符串类型转换成`[]rune`数组来做.