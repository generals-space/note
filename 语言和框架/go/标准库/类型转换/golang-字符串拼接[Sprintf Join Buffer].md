# golang-字符串拼接

参考文章

1. [[笔记]Go语言的字符串拼装方式性能对比](https://studygolang.com/articles/2507)
    - 还指明了这几种字符串拼接方式的速度水平, 可以看一下

## 1. `+`

首先就是`字符串A + 字符串B`, 这个最简单了.

```go
package main

import "log"

func main() {
    a := "hello"
    b := "world"
    c := a + " " + b 
    // 输出`hello world`
    log.Printf("%s", c)
}
```

## 2. `fmt.SPrintf()`

```go
c := fmt.Sprintf("%s %s", a, b)
```

这个不多说, 太麻烦

## 3. `strings.Join()`

这个函数的原型如下

```go
Join func(a []string, sep string) string
```

可以看到和其他语言的`join`函数类似, 把字符串数组按照指定的分隔字符拼接起来, 需要引入`strings`包.

使用方法如下

```go
package main

import "strings"
import "log"

func main() {
    a := "hello"
    b := "world"
    c := strings.Join([]string{a, b}, " ")
    log.Printf("%s", c)
}
```

## 4. 借助`bytes.Buffer`结构体

`bytes.Buffer`结构体的成员不需要清楚, 只需要知道这是一个类`[]byte`的缓冲区就可以了, 我们可以往里以字符串的形式写入, 之后把这个结构体中的内容导出为字符串.

```go
package main

import "bytes"
import "log"

func main() {
	a := "hello"
	b := "world"
	var c bytes.Buffer
	c.WriteString(a)
	c.WriteString(" ")
	c.WriteString(b)

	s := c.String()
	log.Printf("%s", s)
}
```
