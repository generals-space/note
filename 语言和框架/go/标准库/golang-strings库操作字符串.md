# golang-strings库操作字符串

参考文章

1. [golang package整理(strings)](https://studygolang.com/articles/4588)

2. [Go语言中字符串的查找方法小结](http://www.jb51.net/article/73981.htm)

## 1. 判断包含

```go
//判断字符串 s 是否以 prefix 开头
strings.HasPrefix(s, prefix string) bool
//判断字符串 s 是否以 suffix 结尾
strings.HasSuffix(s, suffix string) bool
//判断字符串 s 是否包含 substr
strings.Contains(s, substr string) bool
```

示例代码

```go
package main

import "strings"
import "log"

func main() {
	a := "hello world!"
	b := "hello"

	// 结果为布尔值
	if strings.HasPrefix(a, b) {
		log.Println("yes")
	}
}
```

## 2. 查找索引

```go
// 查找str在s中的索引（str 的第一个字符的索引），-1表示未找到
strings.Index(s, str string) int
// indexAny貌似与index用法完全一样...
strings.IndexAny(s, chars string) int

// 反向查找str在字符串s中出现位置的索引（str 的第一个字符的索引），-1表示未找到
strings.LastIndex(s, str string) int
```

## 3. 替换

```go
// 将str中的前n个字符串old替换为字符串new，并返回一个新的字符串，如果n = -1则替换所有
strings.Replace(str, old, new, n) string
```

示例代码

```go
package main

import "strings"
import "log"

func main() {
	a := "hello world!"

	str1 := strings.Replace(a, "l", "x", 2)
	str2 := strings.Replace(a, "l", "x", -1)
	log.Printf("%s\n", str1)
	log.Printf("%s\n", str2)
}
```

## 4. 分割

```go
// 将s以空白符号分隔, 返回切片对象. 如果s只包含空白符号，则返回一个长度为 0 的 slice
strings.Fields(s)
// 可以自定义分割符号来对指定字符串进行分割，同样返回 slice. 
strings.Split(s, sep)
//Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
Strings.Join(sl []string, sep string)
```

```go
package main

import "strings"
import "log"

func main() {
	a := "a b  c   d  "

	str1 := strings.Fields(a)
	log.Printf("长度: %d\n", len(str1))
	for _, v := range str1 {
		log.Println(v)
    }
    
    str2 := strings.Split(a, "c")
    log.Printf("长度: %d\n", len(str2))
	for _, v := range str2 {
		log.Println(v)
    }
}
```

## 5. 其他

```go
//ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符
strings.ToLower(s) string
//ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符
strings.ToUpper(s) string
// 用于计算字符串 str 在字符串 s 中出现的非重叠次数
strings.Count(s, str string) int
//Repeat 用于重复 count 次字符串 s 并返回一个新的字符串
strings.Repeat(s, count int) string
// 去除字符串s开头和结尾的空白符号
strings.TrimSpace(s)
// 将开头和结尾的 cut 去除掉, 该函数的第二个参数可以包含任何字符，如果你只想剔除开头或者结尾的字符串，则可以使用`TrimLeft`或者`TrimRight`来实现
strings.Trim(s, "cut")
```