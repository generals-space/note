参考文章

1. [C? Go? Cgo!](https://go.dev/blog/cgo)
2. [golang-wiki/cgo.md](https://github.com/zchee/golang-wiki/blob/master/cgo.md)
3. [cgo快速入门之golang调用C语言](https://zhuanlan.zhihu.com/p/116749102)
4. [第二章 CGO编程](https://www.cntofu.com/book/73/ch2-cgo/readme.md)

## 最简示例

```go
package main

//#include <stdio.h>
import "C"

func main() {
	C.puts(C.CString("Hello, World\n"))
}
```

golang+clang 混合编程的格式在于, clang 部分在注释中写明, 然后紧跟`import "C"`, 中间不可以有空行.

golang 要 import 其他包, 需要单独写`import()`块.
