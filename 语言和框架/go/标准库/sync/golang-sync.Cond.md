# golang-sync.Cond

参考文章

1. [golang的sync.cond使用](https://blog.csdn.net/jinglexy/article/details/80516788)

python的标准库`threading`中有一个`Condition`类, 与其`Event`类实现了高于互斥锁与可重入锁的锁控制机制, 和golang的`sync.Cond`的原理与使用场景很相似.

在参考文章1中, python中的实现的场景比较复杂, 但也比较实际. 而下面的示例是基于参考文章1的, 比较简单, 基本没有什么控制流程. 两者可以对比一下, 互相移植.

```go
package main

import (
	"log"
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})

func test(x int) {
	cond.L.Lock()
	defer cond.L.Unlock()

	log.Println("获取锁, 等待通知: ", x)
	cond.Wait()
	log.Println("收到通知, 继续执行: ", x)
}

func main() {
	for i := 0; i < 5; i++ {
		// 多次运行, 但是同时得到锁, 说明其内部锁为可重入锁...?
		// 错了, cond.Wait()方法内部有Unlock()的操作, 所以协程中才能打印"获取锁"这一句
		go test(i)
	}
	time.Sleep(time.Second * 1)
	log.Println("start all")

	// 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	log.Println("signal")
	cond.Signal()

	// 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 5)
	log.Println("signal")
	cond.Signal()

	// 3秒之后 下发广播给所有等待的goroutine
	time.Sleep(time.Second * 7)
	log.Println("broadcast")
	cond.Broadcast()

	time.Sleep(time.Second * 5)
	log.Println("finish all")

}

```

执行输出如下

```
$ go run .\main.go
2019/05/10 00:08:54 获取锁, 等待通知:  0
2019/05/10 00:08:54 获取锁, 等待通知:  2
2019/05/10 00:08:54 获取锁, 等待通知:  1
2019/05/10 00:08:54 获取锁, 等待通知:  4
2019/05/10 00:08:54 获取锁, 等待通知:  3
2019/05/10 00:08:55 start all
2019/05/10 00:08:58 signal
2019/05/10 00:08:58 收到通知, 继续执行:  0
2019/05/10 00:09:03 signal
2019/05/10 00:09:03 收到通知, 继续执行:  2
2019/05/10 00:09:10 broadcast
2019/05/10 00:09:10 收到通知, 继续执行:  1
2019/05/10 00:09:10 收到通知, 继续执行:  4
2019/05/10 00:09:10 收到通知, 继续执行:  3
2019/05/10 00:09:15 finish all
```