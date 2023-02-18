# cgo无法打印字符串[vscode printf stdout]

参考文章

1. [golang-wiki/cgo.md](https://github.com/zchee/golang-wiki/blob/master/cgo.md)
2. [Using cgo, why does C output not 'survive' piping when golang's does?](https://stackoverflow.com/questions/42634640/using-cgo-why-does-c-output-not-survive-piping-when-golangs-does)

## 问题描述

golang: 1.16.15

最初学习使用 cgo 时, 使用如下代码, 通过 vscode 进行调试.

```go
package main

//#include <stdio.h>
import "C"

func main() {
	C.puts(C.CString("Hello, World"))
}
```

vscode 的 launch.json 配置如下

```json
{
	"name": "cgotest",
	"type": "go",
	"request": "launch",
	"mode": "auto",
	"env": {
		"GO111MODULE": "auto",
		"CGO_ENABLED": "1",
	},
	"program": "${workspaceFolder}"
}
```

但是点击调试时, 啥也没输出.

------

然后修改代码如下

```go
package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import (
	"log"
)

func main() {
	log.Printf("begin")
	cs := C.CString("Hello world")
	C.myprint(cs)
	log.Printf("end")
}
```

此时的输出如下

```log
2023/01/07 18:55:50 begin
2023/01/07 18:55:50 end
```

还是没输出

## 解决方案

按照参考文章2所说, 需要在print语句后, 加一句flush刷新缓冲区.

```go
	C.puts(C.CString("Hello, World\n"))
	C.fflush(C.stdout)
```

```go
	C.myprint(cs)
	C.fflush(C.stdout)
```

可以了.

------

不过后来发现, 这并不是 golang 本身的问题, 因为在命令行执行`go run main.go`, 不需要`C.fflush(C.stdout)`也能正常输出.
