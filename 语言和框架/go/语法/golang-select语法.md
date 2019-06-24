# golang-select语法

参考文章

1. [Golang select的使用及典型用法](https://blog.csdn.net/zhaominpro/article/details/77570290)

2. [Go 语言 select 语句](https://wizardforcel.gitbooks.io/w3school-go/content/13.html)

3. [由浅入深聊聊Golang中select的实现机制](https://blog.csdn.net/u011957758/article/details/82230316)
    - 一对select的陷阱示例
    - select在应用中的运行机制
    - select底层源码分析, 文末的流程图值得一看

## 1. 基本用法 

`select`是Go中的一个控制结构, 类似于`switch`语句, 用于处理异步IO操作. `select`会监听`case`语句中`channel`的读写操作, 当`case`中`channel`读写操作为非阻塞状态（即能读写）时, 将会触发相应的动作. 

> `select`中的`case`语句必须是一个`channel`操作

> `select`中的`default`子句总是可运行的

如果有多个`case`都可以运行, select会**随机公平**地选出一个执行, 其他不会执行. 

如果没有可运行的`case`语句, 且有`default`语句, 那么就会执行`default`的动作. 

如果没有可运行的`case`语句, 且没有`default`语句, `select`将阻塞, 直到某个`case`通信可以运行

1. select+case是用于阻塞监听goroutine的, 如果没有case, 就单单一个select{}, 则为监听当前程序中的goroutine, 此时注意, 需要有真实的goroutine在跑, 否则select{}会报panic
2. select底下没有可执行的case时, 走default, 没有default会阻塞(注意: 必须要有某个协程对case中的channel进行读/写, 以保证整个select流程不是死锁. 你可以让ta们sleep等待, 但必须存在, 否则仍然会报panic). 有多个可执行的case时, 则随机执行一个. 
3. select常配合for循环来监听channel有没有故事发生. 需要注意的是在这个场景下, break只是退出当前select而不会退出for, 需要用break TIP / goto的方式. 
4. 无缓冲的通道, 则传值后立马close, 则会在close之前阻塞, 有缓冲的通道则即使close了也会继续让接收后面的值
5. 同个通道多个goroutine进行关闭, 可用recover panic的方式来判断通道关闭问题

## 1. case可读可写, 可为同一channel的不同操作

case可读可写可关闭.

```go
	ch := make(chan int, 1)
	var count int
	for {
		select {
		case ch <- count:
			count++
			fmt.Printf("case 1, count: %d\n", count)
		case <-ch:
			count++
			fmt.Printf("case 2, count: %d\n", count)
		}
		time.Sleep(time.Second * 1)
	}
```

```
case 1, count: 1
case 2, count: 2
case 1, count: 3
case 2, count: 4
...
```

## 2. case可为同一channel的相同操作

```go
	ch := make(chan int)
	go func() {
		var counter int
		for {
			counter++
			ch <- counter
		}
	}()

	for {
		select {
		case c := <-ch:
			fmt.Printf("case 1, counter: %d\n", c)
		case c := <-ch:
			fmt.Printf("case 2, counter: %d\n", c)
		}
		time.Sleep(time.Second)
	}
```

执行结果

```
case 2, counter: 1
case 1, counter: 2
case 2, counter: 3
case 1, counter: 4
case 1, counter: 5
case 1, counter: 6
case 2, counter: 7
case 1, counter: 8
...
```

可以看到就算case相同, 每次执行的选择都是随机且不是固定平均的.

## 3. 空语言`select{}`作为阻塞主协程的手段, wait子协程

```go
	go func() {
		var counter int
		for {
			counter++
			fmt.Println(counter)
			time.Sleep(time.Second)
		}
	}()

	select {}
```

```
1
2
3
4
5
...
```

注意: `select{}`可以用来阻塞主协程, 等待子协程的执行, 但是必须要保证真的有子协程在运行, 否则会报panic

```
fatal error: all goroutines are asleep - deadlock!
```

在本例中, 协程中是一个无限循环, 如果该循环有限, 在for循环结束, 子协程退出时, 主线程也会报这个错.

## 4. 关于channel对象close事件的通知

```go
	ch := make(chan int, 1)
	close(ch)

	for {
		select {
		case <-ch:
			fmt.Println("case 1")
		}
	}
```

```
case 1
case 1
case 1
case 1
...
```

注意: 当case为读取被关闭的channel时, 总是会成功. 这是由于channel的特性, 向一个缓冲通道写入值后将其关闭, 读操作依然可以进行, 将数据读完后还是可以进行读操作, 只是读不出数据了. 但是对于一个被关闭了的channel进行写操作会panic.

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
