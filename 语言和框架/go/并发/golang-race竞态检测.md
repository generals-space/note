# golang-race竞态检测

参考文章

1. [介绍Go竞争检测器](https://blog.csdn.net/fighterlyt/article/details/13994887)
    - 参考文章2的译文
2. [Introducing the Go Race Detector](https://blog.golang.org/race-detector)
    - 参考文章1的原文
3. [golang中的race检测](https://www.cnblogs.com/yjf512/p/5144211.html)
    - 对参考文章1中第一个示例给出了一个更简单明了的示例, 并解释地很好.

我们用参考文章3中的示例来说明

```go
func main() {
    a := 1
    go func(){
        a = 2
    }()
    a = 3
    fmt.Println("a is ", a)

    time.Sleep(2 * time.Second)
}
```

直接用`go run main.go`完全没问题, 

```
$ go run -race main.go
a is  3
```

但是如果加上`-race`选项, 就会报错.

```
$ go run -race main.go
a is  3
==================
WARNING: DATA RACE
Write at 0x00c0000bc000 by goroutine 6:
  main.main.func1()
      /Users/general/Code/playground/go-race/main.go:11 +0x38

Previous write at 0x00c0000bc000 by main goroutine:
  main.main()
      /Users/general/Code/playground/go-race/main.go:13 +0x88

Goroutine 6 (running) created at:
  main.main()
      /Users/general/Code/playground/go-race/main.go:10 +0x7a
==================
Found 1 data race(s)
exit status 66
```

上面的代码, 错就错在没有对共享变量`a`加锁保护, 或者用channel通道做通信也行.

```go
func main() {
	m := &sync.Mutex{}
	a := 1

	go func() {
		m.Lock()
		a = 2
		m.Unlock()
	}()

	m.Lock()
	a = 3
	fmt.Println("a is ", a)
	m.Unlock()

	time.Sleep(2 * time.Second)
}

```

这样运行就没错了.

参考文章1和2的第2个示例我没看懂, 应该是golang早期的问题, 我阅读最早的golang源码是go1.2, 已经找不到相关的问题代码了, `blackHole`也使用了`sync.Pool`对象缓存, 应该是已经修复过的.

这个问题先保留.

## 2. 如何实现

从上面的示例可以看到, 使用`-race`选项并不会让进程退出, ta只是打印了哪里可能存在问题, 只是warning级别.
