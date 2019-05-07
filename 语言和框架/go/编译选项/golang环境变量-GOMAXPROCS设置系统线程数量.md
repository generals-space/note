# golang环境变量-GOMAXPROCS设置系统线程数量

参考文章

1. [Go 语言运行时环境变量快速导览](https://blog.csdn.net/htyu_0203_39/article/details/50852856)

2. [Go语言GOMAXPROCS（调整并发的运行性能）](http://c.biancheng.net/view/94.html)

在程序启动前设置`GOMAXPROCS`的值与在程序中使用`runtime`库设置的效果是相同的.

以如下代码为例

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
    // 设置只使用1个CPU核心
	// runtime.GOMAXPROCS(1)
	counter := runtime.NumCPU()
	waitG := &sync.WaitGroup{}

	fmt.Printf("CPU core %d\n", counter)

	startT := time.Now().Unix()
	waitG.Add(counter)
	for i := 0; i < counter; i++ {
		// waitG.Add(i)
		go func(waitG *sync.WaitGroup, ord int) {
			// 每个协程空等5s...好像在golang里time.Sleep()是异步操作, 不会浪费CPU
			// time.Sleep(time.Second * 5)
			var x int
			for x = 0; x < 10000000000; x++ {
			}
			waitG.Done()
			fmt.Printf("goroutine %d complete, x: %d\n", ord, x)
		}(waitG, i)
	}
	waitG.Wait()
	endT := time.Now().Unix()
	fmt.Printf("cost time %ds\n", endT-startT)

}
```

测试电脑为8核, linux系统常规执行输出如下

```
$ go run main.go
CPU core 8
goroutine 7 complete, x: 10000000000
goroutine 2 complete, x: 10000000000
goroutine 4 complete, x: 10000000000
goroutine 5 complete, x: 10000000000
goroutine 6 complete, x: 10000000000
goroutine 1 complete, x: 10000000000
goroutine 3 complete, x: 10000000000
goroutine 0 complete, x: 10000000000
cost time 9s
```

执行5次, cost time的值分别为8, 8, 10, 9, 9, 平均`8.8s`

如果使用`GOMAXPROCS=1 go run main.go`, 或者解开上面代码中的`runtime.GOMAXPROCS(1)`, 得到的cost time分别为

1. 34, 34, 35, 35, 34, 平均`34.4s`
2. 35, 35, 35, 37, 36, 平均`35.6s`

可以看到两者的效果是相同的.

至于8线程与单线程的执行效率相差只有4倍的问题...管ta呢, 本文的主题并不是这个.
