# channel的实现原理

参考文章

1. [Golang-Channel原理解析](https://blog.csdn.net/u010853261/article/details/85231944)
    - 讲解channel的底层数据结构`hchan`, 图形化展示存取数据, 及数据满等情况下channel的工作流程
    - channel源码分析
2. [【我的架构师之路】- golang源码分析之channel的底层实现](https://blog.csdn.net/qq_25870633/article/details/83388952)
    - channel源码分析
    - send/recv/close的执行流程图, 很清晰
3. [深入理解channel：设计+源码](https://www.jianshu.com/p/afe41fe1f672)
    - PPT资料 [Understanding Channels](https://speakerdeck.com/kavya719/understanding-channels)


我本来以为chan, slice, map这种内置类型是在语言底层实现的, 但原来是在标准库中实现的. chan的源码在`runtime/chan.go`, 在`runtime`目录中还包含`slice.go`和`map.go`.

channel是协程安全的, 因为其内部使用了mutex. 

作为协程间通信的通道, 每个协程在读写channel对象时都会与其绑定, 在close()操作时, 内部会遍历所有读写的协程, 依次解除联系.

来一个问题, 一个有数据缓冲的channel, 在关闭后, 其中的数据还能被读取吗?

答案是可以的, 向一个缓冲通道写入值后将其关闭, 读操作依然可以进行, 将数据读完后还是可以进行读操作, 只是读不出数据了. 但是对于一个被关闭了的channel进行写操作会panic.

```go
	ch := make(chan int, 1)
	ch <- 1
	close(ch)

	val1, ok1 := <-ch
	fmt.Println(val1) // 1
	fmt.Println(ok1)  // true

	val2, ok2 := <-ch
	fmt.Println(val2) // 0
	fmt.Println(ok2)  // false
```
