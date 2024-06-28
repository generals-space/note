原文链接: [Go 性能调优](https://sammyne.github.io/go-profiling/#runtime-pprof)

trace 工具也是 go 工具之一，能够辅助我们跟踪程序的执行情况，进一步方便我们排查问题，往往配合 pprof 使用。trace 的使用和 pprof 类似，为了简化分析，我们首先利用下列代码进行讲解，只是用单核核运行程序：

```go
// +build ignore

package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

var hello []int

func counter(wg *sync.WaitGroup) {
	defer wg.Done()

	slice := []int{0}
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
	hello = slice
}

func main() {
	runtime.GOMAXPROCS(1)

	var traceProfile = flag.String("traceprofile", "trace.pprof", "write trace profile to file")
	flag.Parse()

	if *traceProfile != "" {
		f, err := os.Create(*traceProfile)
		if err != nil {
			log.Fatal(err)
		}
		trace.Start(f)
		defer f.Close()
		defer trace.Stop()
	}

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go counter(&wg)
	}
	wg.Wait()
}
```

同样，通过编译、执行和如下指令得到 trace 图：

```
go tool trace -http=127.0.0.1:8000 trace.pprof
```

![](https://gitee.com/generals-space/gitimg/raw/master/2024/d93aa4d53745f9680539b52dec9b6e64.png)

图像的查看依赖 Chrome 浏览器。如果依旧无法查看改图像，MacOS 请按照下述方法进行操作

1. 登录 google 账号，访问 https://developers.chrome.com/origintrials/#/register_trial/2431943798780067841，其中 web Origin 字段为此后你需要访问的 web 网址，例如我使用的 127.0.0.1:8000。如此你将获得一个 Active Token 并复制下来。
2. 在 go 的安装目录 $GOROOT/src/cmd/trace/trace.go 文件找到元素范围，并添加
3. 在该目录下分别执行 go build 和 go install，此后重启 Chrome 浏览器即可查看上图。

上图有几个关键字段，分别介绍如下：

- `Goroutines`: 运行中的协程数量；通过点击图中颜色标识可查看相关信息，可以看到在大部分情况下可执行的协程会很多，但是运行中的只有 0 个或 1 个，因为我们只用了 1 核。
- `Heap`: 运行中使用的总堆内存；因为此段代码是有内存分配缺陷的，所以 heap 字段的颜色标识显示堆内存在不断增长中。
- `Threads`: 运行中系统进程数量；很显然只有 1 个。
- `GC`: 在程序的末端开始回收资源。
- `Syscalls`: 系统调用；由上图看到在 GC 开始只有很微少的一段。
- `Proc0`: 系统进程，与使用的处理器的核数有关，1 个。

另外由图可知，程序的总运行时间约 4ms。在 Proc0 轨道上，不同颜色代表不同协程，各个协程都是串行的，执行 counter 函数的有 G6、G7 和 G8 协程，同时 Goroutines 轨道上的协程数量也相应在减少。伴随着协程的结束，GC 也会将内存回收，另外 GC 过程出现了 STW（stop the world）过程，这对程序的执行效率会有极大的影响。STW 过程会将整个程序通过 sleep 停止下来，所以前文出现的 runtime.usleep 就是此时由 GC 调用的。

下面我们使用多个核来运行，只需要改动 GOMAXPROCS 即可，例如修改成 5 并获得 trace 图：

```
runtime.GOMAXPROCS(5)
```

TODO（结果无法复现）：从上图可以看到，3 个 counter 协程在每个核上都有执行，同时程序的运行时间为 TODO（目前实际时间变化不大，甚至增加），运行时间大大降低，可见提高 cpu 核数是可以提高效率的，但是也不是所有场景都适合提高核数，还是需要具体分析。同时为了减少内存的扩容，同样可以预先分配内存，获得 trace 图如下所示：

![](https://gitee.com/generals-space/gitimg/raw/master/2024/950a9476730a34827c14a61472649fdc.png)

由上图看到，由于我们提前分配好足够的内存，系统不需要进行多次扩容，进而进一步减小开销。从 slice 的源码中看到其实现中包含指针，即其内存是堆内存，而不是 C/C++ 中类似数组的栈空间分配方式。另外也能看到程序的运行时间为 0.18ms，进一步提高运行速度。

另外，trace 图还有很多功能，例如查看事件的关联信息等等，通过点击`Flow events/All`即可生成箭头表示相互关系，大家可以自己探究一下其他功能。

![](https://gitee.com/generals-space/gitimg/raw/master/2024/210393fae2c03243c40c7ffd1c985576.png)

如果我们对 counter 函数的循环中加上锁会发生什么呢？

```go
func counter(wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()

	slice := make([]int, 0, 100000)
	c := 1
	for i := 0; i < 100000; i++ {
		mtx.Lock()
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
		mtx.Unlock()
	}
	hello = slice
}
```

生成 trace 图如下：

![](https://gitee.com/generals-space/gitimg/raw/master/2024/5968e42b4acc4541a624b3451f2645da.png)

可以看到程序运行的时间又增加了，主要是由于加/放锁使得 counter 协程的执行时间变长。~~但是并没有看到不同协程对 CPU 占有权的切换呀？这是为什么呢？主要是这个协程运行时间太短，而相对而言采样的频率低、粒度大，导致采样数据比较少。~~ 如果在程序中人为 sleep 一段时间，提高采样数量可以更加真实反映 CPU 占有权的切换。

如果对 go 协程加锁呢？

```go
for i := 0; i < 3; i++ {
	mtx.Lock()
	go counter(&wg)
	time.Sleep(time.Millisecond)
	mtx.Unlock()
}
```

从得到的 trace 图可以看出，CPU 主要时间都是在睡眠等待中，所以在程序中应该减少此类 sleep 操作。

![](https://gitee.com/generals-space/gitimg/raw/master/2024/5557eb91e8f4658fa3d63252075e2559.png)

trace 图可以非常完整地跟踪程序的整个执行周期，所以大家可以从整体到局部分析优化程序。

我们可以先使用 pprof 完成初步的检查和优化，主要是 CPU 和内存，而 trace 主要是用于分析各个协程的执行关系，从而优化结构。

本文主要讲解了一些性能评测和 trace 的方法，仍然比较浅显，更多用法大家可以自己去探索。

