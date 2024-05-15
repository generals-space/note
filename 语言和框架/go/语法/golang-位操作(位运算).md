# golang-位操作

参考文章

1. [golang 位运算](https://studygolang.com/articles/6337)
2. [运算操作符](https://gfw.go101.org/article/operators.html)
	- Go语言101
3. [Operators](https://go.dev/ref/spec#Operators)
	- 官方文档

- 按位与`&`
- 按位或`|`
- 按位异或`^`
- 按位取反`^`
- 左移`<<`
- 右移`>>`

```go
package main

import (
	"log"
)

func main() {
	byteArray := []byte{0x01, 0x02}
	// Warning: byteArray[0] (8 bits) too small for shift of 8
	result1 := byteArray[0]<<8 + byteArray[1]
	log.Println(result1)

	intArray := []int{1, 2}
	result2 := intArray[0]<<8 + intArray[1]
	log.Println(result2)
}
```
