# channel 关闭 读写与panic

参考文章

1. [如何优雅地关闭Channel](https://www.jianshu.com/p/c7b25ed78b89)

go version: 1.12

1. 没有一种简单通用地方法来检测channel是否关闭而不修改channel地状态;
2. 关闭一个已关闭的channel会引起Panic, 因此如果不知道channel是否关闭, 那么关闭channel将会非常危险;
3. 将值发送到已关闭的channel会发生panic, 因此如果发送者不知道channel是否关闭, 则将值发送到channel中是危险的;
4. 关闭一个channel时, 如果该channel中还存放着数据, 这些数据并不会消失, 仍然可以被读取出来;
5. 当一个已关闭的channel中没有数据后, 仍然可以继续读取, 此时读取出的数据为该channel类型的默认值(比如0, false等, 如果是结构体, 那么会读出空的结构体);
6. 当关闭一个空的channel, 但同时有协程等待读取时, 读协程会得到channel类型的默认值, 同5;
7. 当关闭一个满的channel, 但同时有协程等待写入时, 写入协程会panic, 同3.
8. 一个只使用`var channel chan int`声明的 channel, 读写都会阻塞, close()还会panic.

示例6 

```go
	channel := make(chan int, 3)
	go func() {
		num := <-channel
		fmt.Printf("receive %d\n", num)
	}()
	time.Sleep(time.Second * 3)
	close(channel)
	time.Sleep(time.Second * 2)
```

会输出"receive 0"

示例7

```go
	channel := make(chan int, 3)
	channel <- 1
	channel <- 2
	channel <- 3
	go func() {
		channel <- 4
		fmt.Printf("complete")
	}()
	time.Sleep(time.Second * 3)
	close(channel)
	time.Sleep(time.Second * 2)
```

会出现"panic: send on closed channel"