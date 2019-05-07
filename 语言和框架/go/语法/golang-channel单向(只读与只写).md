# golang-channel单向(只读与只写)

参考文章

1. [Go学习笔记 第4版 7.2.1节]()

单向channel一般是从正常的channel中通过类型声明转换而来(不然只读通道的数据从哪来, 只写通道的数据要给谁?).

```go
package main

import "log"

func main() {
	channel := make(chan int, 10)
	// 将rchan声明为只读通道, wchan声明为只写通道
	var rchan <-chan int = channel
	var wchan chan<- int = channel

	for i := 0; i < 10; i++ {
		wchan <- i
	}

	for i := range rchan {
		log.Println(i)
		if len(rchan) == 0 {
			break
		}
	}
}
```

输出为

```
2019/05/05 13:54:40 0
2019/05/05 13:54:40 1
2019/05/05 13:54:40 2
2019/05/05 13:54:40 3
2019/05/05 13:54:40 4
2019/05/05 13:54:40 5
2019/05/05 13:54:40 6
2019/05/05 13:54:40 7
2019/05/05 13:54:40 8
2019/05/05 13:54:40 9
```

我们可以在程序里把`rchan`和`wchan`传递给某些方法使用, 以使得读写操作不会混乱.

> 无法再将单向channel再转换成普通双向channel.