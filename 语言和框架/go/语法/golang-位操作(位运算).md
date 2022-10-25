# golang-位操作

参考文章

1. [golang 位运算](https://studygolang.com/articles/6337)

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
