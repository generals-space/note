# select..case重新执行的问题

参考文章

1. [由浅入深聊聊Golang中select的实现机制](https://blog.csdn.net/u011957758/article/details/82230316)
    - 一对关于select陷阱示例
    - select的底层运行机制, 文末的流程图值得一看

有如下两个示例

`select1.go`

```go
	t1 := time.Tick(time.Millisecond * 499)
	t2 := time.Tick(time.Millisecond * 500)
	var count int
	for {
		select {
		case <-t1:
			count++
			fmt.Printf("case 1, count: %d\n", count)
		case <-t2:
			count++
			fmt.Printf("case 2, count: %d\n", count)
		}
	}
```

`select2.go`

```go
	var count int
	for {
		select {
		case <-time.Tick(time.Millisecond * 499):
			count++
			fmt.Printf("case 1, count: %d\n", count)
		case <-time.Tick(time.Millisecond * 500):
			count++
			fmt.Printf("case 2, count: %d\n", count)
		}
	}
```

我们知道select会从case所表示的channel对象随机选择一个可用(可读, 可写, 被关闭等)的channel进行操作, 按照这个说法, `select1.go`的执行结果是符合我们的认识的.

```
$ go run select1.go
case 1, count: 1
case 2, count: 2
case 1, count: 3
case 2, count: 4
case 1, count: 5
case 2, count: 6
case 1, count: 7
case 2, count: 8
...
```

两个case语句交替执行.

但是`select2.go`的执行结果就有些奇异了, ta每次执行的都是case 1.

```
$ go run select2.go
case 1, count: 1
case 1, count: 2
case 1, count: 3
case 1, count: 4
...
```

参考文章1最后给出的解释比较拗口, 简单说来就是, 在`select2.go`的每次for循环中, `select{}`对于要执行的case都是重新计算的, `time.Tick()`499总是在500之前, 所以每次都会执行case 1, 而另一个就被丢弃.