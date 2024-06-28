原文链接: [Go 性能调优](https://sammyne.github.io/go-profiling/#runtime-pprof)

pprof 是 go 官方提供的性能测评工具, 包含在`net/http/pprof`和`runtime/pprof`两个包, 分别用于不同场景:

1. `runtime/pprof`主要用于可结束的代码块, 如一次编解码操作等;
2. `net/http/pprof`是对`runtime/pprof`的二次封装, 主要用于不可结束的代码块, 如 web 应用等;

pprof 开启后，每隔一段时间（10ms）就会收集下当前的堆栈信息，获取各个函数占用的 CPU 以及内存资源，最后通过对这些采样数据进行分析，形成一个性能分析报告。

## runtime/pprof

我们先看看如何利用 runtime/pprof 进行性能测评。

下列代码循环向一个列表尾部添加元素。导入`runtime/pprof`并添加两段测评代码（补充具体行号）就可以实现 CPU 和内存的性能评测。

```go
// counter_v1.go
package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

var (
	cpu string
	mem string
)

func init() {
	flag.StringVar(&cpu, "cpu", "", "write cpu profile to file")
	flag.StringVar(&mem, "mem", "", "write mem profile to file")
}

func main() {
	flag.Parse()

	//采样 CPU 运行状态
	if cpu != "" {
		f, err := os.Create(cpu)
		if err != nil {
			log.Fatal(err)
		}
		_ = pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go workOnce(&wg)
	}
	wg.Wait()

	//采样内存状态
	if mem != "" {
		f, err := os.Create(mem)
		if err != nil {
			log.Fatal(err)
		}
		_ = pprof.WriteHeapProfile(f)
		f.Close()
	}
}

func counter() {
	slice := make([]int, 0)
	var c int
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
	_ = slice
}

func workOnce(wg *sync.WaitGroup) {
	counter()
	wg.Done()
}
```

编译并执行获得 pprof 的采样数据，然后利用相关工具进行分析。

```log
$ go build -o app counter_v1.go
./app --cpu=cpu.pprof
./app --mem=mem.pprof
```

至此就可以获得 cpu.pprof 和 mem.pprof 两个采样文件，然后利用 go tool pprof 工具进行分析如下。

```log
$ go tool pprof cpu.pprof
Type: cpu
Time: Apr 12, 2021 at 10:42pm (CST)
Duration: 201.51ms, Total samples = 310ms (153.84%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top                             ## 交互式命令
Showing nodes accounting for 310ms, 100% of 310ms total
Showing top 10 nodes out of 58
      flat  flat%   sum%        cum   cum%
     150ms 48.39% 48.39%      150ms 48.39%  runtime.memmove
      50ms 16.13% 64.52%       50ms 16.13%  runtime.memclrNoHeapPointers
      50ms 16.13% 80.65%       50ms 16.13%  runtime.usleep
      10ms  3.23% 83.87%       10ms  3.23%  runtime.madvise
      10ms  3.23% 87.10%       10ms  3.23%  runtime.markBits.isMarked
      10ms  3.23% 90.32%       10ms  3.23%  runtime.procyield
      10ms  3.23% 93.55%       10ms  3.23%  runtime.pthread_cond_wait
      10ms  3.23% 96.77%       10ms  3.23%  runtime.pthread_kill
      10ms  3.23%   100%       10ms  3.23%  runtime.scanobject
         0     0%   100%      210ms 67.74%  main.counter
(pprof)
```

相关字段如下（Type 和 Time 字段就不过多解释了）：

- `Duration`:   程序执行时间。本例中 go 自动分配任务给多个核执行程序，总计耗时 201.51ms，而采样时间为 310ms；也就是说假设有 10 核执行程序，平均每个核采样 31ms 的数据;
- `top`:        pprof 的指令之一，显示 pprof 文件的前 10 项数据，可以通过 top 20 等方式显示前 20 行数据。pprof 还有很多指令，例如 list、pdf、eog 等等;
- `flat/flat%`: 分别表示在当前层级的 CPU 占用时间和百分比。例如 runtime.memmove 在当前层级占用 CPU 时间 150ms，占比本次采集时间的 48.39%;
- `cum/cum%`:   分别表示截止到当前层级累积的 CPU 时间和占比。例如 main.counter 累积占用时间 210ms，占本次采集时间的 67.74%;
- `sum%`:       所有层级的 CPU 时间累积占用，从小到大一直累积到 100%，即 310ms;

由上图的 cum 数据可以看到，counter 函数的 CPU 占用时间最多。接下来可利用 list 命令查看占用的主要因素如下

```log
(pprof) list main.counter
Total: 310ms
ROUTINE ======================== main.counter in /Users/sammy/Workspaces/github.com/sammyne/go-profiling/code/counter_v1.go
         0      210ms (flat, cum) 67.74% of Total
         .          .     52:func counter() {
         .          .     53:   slice := make([]int, 0)
         .          .     54:   var c int
         .          .     55:   for i := 0; i < 100000; i++ {
         .          .     56:           c = i + 1 + 2 + 3 + 4 + 5
         .      210ms     57:           slice = append(slice, c)
         .          .     58:   }
         .          .     59:   _ = slice
         .          .     60:}
         .          .     61:
         .          .     62:func workOnce(wg *sync.WaitGroup) {
(pprof)
```

可见，程序的 57 行分别占用 210ms，这就是优化的主要方向。通过分析程序发现，由于切片的初始容量为 0，导致循环 append 时触发多次扩容。切片的扩容方式是：申请 2 倍或者 1.25 倍的原来长度的新切片，再将原来的切片数据拷贝进去。

仔细一看还会发现：runtime.usleep 占用 CPU 时间将近 16.13%，但是程序明明没有任何 sleep 相关的代码，却为什么会出现，并且还占用这么高呢？大家可以先思考一下，后文将揭晓。

当然，也可以使用 web 命令获得更加直观的信息。macOS 通过如下命令安装渲染工具 graphviz。

安装完成后，在 pprof 的命令行输入 svg 即可生成一个 svg 格式的文件，将其用浏览器打开即可看到[此图](https://sammyne.github.io/go-profiling/images/counter1.svg)。

由于文件过大，此处只截取部分重要内容如下。

![](https://gitee.com/generals-space/gitimg/raw/master/2024/a2df8d7b01337804eef4e69bd1618770.png)

可以看出其基本信息和命令行下的信息相同，但是可以明显看出 runtime.memmove 耗时 380ms。由图逆向推断 main.counter 是主要的优化方向。**图中各个方块的大小也代表 CPU 占用的情况，方块越大说明占用 CPU 时间越长。**

同理可以分析 mem.pprof 文件，从而得出内存消耗的主要原因进一步进行改进。

上述 main.counter 占用 CPU 时间过多的问题，实际上是 append 函数重新分配内存造成的。那简单的做法就是事先申请一个大的内存，避免频繁的进行内存分配。所以将 counter 函数进行改造：

```go
func counter() {
	var slice [100000]int
	var c int
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice[i] = c
	}
	_ = slice
}
```

通过编译、运行、采集 pprof 信息后如下图所示。

```log
➜  code git:(main) ✗ go build -o app counter_v2.go
➜  code git:(main) ✗ ./app --cpu=cpu.pprof
➜  code git:(main) ✗ go tool pprof cpu.pprof
Type: cpu
Time: Apr 14, 2021 at 10:04pm (CST)
Duration: 200.52ms, Total samples = 0
No samples were found with the default sample value type.
Try "sample_index" command to analyze different sample values.
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
```

可见，已经采集不到占用 CPU 比较多的函数，即已经完成优化。同学们可以试试如果往 counter 添加一个 fmt.Println 函数后，对 CPU 占用会有什么影响呢？

## net/http/pprof

针对后台服务型应用，服务一般不能停止，这时需要使用 net/http/pprof 包。类似上述代码，编写如下代码：

```go
// +build ignore

package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

var hello []int

func counter() {
	slice := make([]int, 0)
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
	// 个人实验时，加上这行貌似可以防止 slice 被优化掉，导致无法统计到期望的内存概况
	hello = slice
}

func workForever() {
	for {
		go counter()
		time.Sleep(1 * time.Second)
	}
}

func httpGet(w http.ResponseWriter, r *http.Request) {
	counter()
}

func main() {
	go workForever()
	http.HandleFunc("/get", httpGet)
	http.ListenAndServe("localhost:8000", nil)
}
```

首先导入 net/http/pprof 包。注意该包利用下划线 _ 导入，意味着只需要该包运行其 init() 函数即可。这样之后，该包将自动完成信息采集并保存到内存。所以服务上线时需要将 net/http/pprof 包移除，避免其影响服务的性能，更重要的是防止其造成内存的不断上涨。

编译并运行依赖，便可以访问：http://localhost:8000/debug/pprof/ 查看服务的运行情况。本文实验得出如下示例，大家可以自行探究查看。不断刷新网页可以发现采样结果也在不断更新中。

![](https://gitee.com/generals-space/gitimg/raw/master/2024/9eacab5bb57595ac710833cb5730e122.png)

当然也可以网页形式查看。现在以查看内存为例，在服务程序运行时，执行下列命令采集内存信息。

采集完成后调用 svg 命令得到如下 svg 文件.

```log
## main 表示编译生成的可执行文件
$ go tool pprof main http://localhost:8000/debug/pprof/heap
Fetching profile over HTTP from http://localhost:8000/debug/pprof/heap
Saved profile in /root/pprof/pprof.main.alloc_objects.alloc_space.inuse_objects.inuse_space.005.pb.gz
File: main
Type: inuse_space
Time: Apr 14, 2021 at 2:44pm (UTC)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) svg                             ## 生成 svc 文件
Generating report in profile001.svg
(pprof)
```

![](https://gitee.com/generals-space/gitimg/raw/master/2024/5cb2c2ffd6844fadf229bfba796b9215.png)

该图表明所有的堆空间均由 counter 产生，同理可以生成 CPU 的 svg 文件用于同步进行分析优化方向。

上述方法在工具型应用可以使用，然而在服务型应用时，仅仅只是采样了部分代码段；而只有当有大量请求时才能看到应用服务的主要优化信息。

另外，Uber 开源的火焰图工具[go-torch](https://github.com/uber-archive/go-torch)也能辅助我们直观地完成测评。感兴趣的话，请自行学习。

注意: pprof 也会使用堆空间，所以在服务上线时应该将 pprof 关闭。

