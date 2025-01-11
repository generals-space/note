# golang是否可以控制底层thread系统线程的数量

参考文章

1. [runtime: limit number of operating system threads](https://github.com/golang/go/issues/4056)
2. [Number of threads used by Go runtime](https://stackoverflow.com/questions/39245660/number-of-threads-used-by-go-runtime)
    - GOMAXPROCS 限制的是 P 的数量, 与 golang runtime 使用的系统线程数没有关系, 无法通过该值指定.
3. [runtime: Should respect/understand the process limit when managing threads](https://github.com/golang/go/issues/14835)
4. [runtime/debug: SetMaxThreads(maxInt) crashes program with "-1-thread limit"](https://github.com/golang/go/issues/16076)
    - runtime/debug 包下提供了`SetMaxThreads()`, 可用来限制 golang 程序底层创建的系统线程数量.
5. [Why my program of golang create so many threads?](https://stackoverflow.com/questions/27600587/why-my-program-of-golang-create-so-many-threads)
    - 什么情况下, golang runtime 会创建许多系统线程(threads)
    - golang 中 net 相关的系统调用是非阻塞的, 因为使用了 epoll 多路复用. 不过磁盘IO则是阻塞的.

可以, 但不是通过 `GOMAXPROCS`环境变量, 而是通过`runtime/debug.SetMaxThreads()`

```go
package main

import (
	"time"
	"runtime/debug"
)

func main() {
	debug.SetMaxThreads(4)
	time.Sleep(time.Second * 1000)
}

```

在 go 1.2 版本中, thread 值最小为 4, 否则"go run main.go"执行时, 会直接报错.

```log
$ go run main.go
runtime: program exceeds 3-thread limit
fatal error: thread exhaustion

goroutine 1 [running]:
runtime.throw(0x587921)
	/root/go/src/pkg/runtime/panic.c:464 +0x69 fp=0x7fe91cf1aef0
checkmcount()
	/root/go/src/pkg/runtime/proc.c:412 +0x50 fp=0x7fe91cf1af08
runtime/debug.setMaxThreads(0x3, 0x2710)
	/root/go/src/pkg/runtime/proc.c:4131 +0x49 fp=0x7fe91cf1af18
runtime/debug.SetMaxThreads(0x3, 0x400c82)
	/root/go/src/pkg/runtime/debug/garbage.go:134 +0x27 fp=0x7fe91cf1af30
main.main()
	/root/go/gopath/src/test/main.go:9 +0x26 fp=0x7fe91cf1af48
runtime.main()
	/root/go/src/pkg/runtime/proc.c:322 +0x11f fp=0x7fe91cf1afa0
runtime.goexit()
	/root/go/src/pkg/runtime/proc.c:1959 fp=0x7fe91cf1afa8
exit status 2
```
