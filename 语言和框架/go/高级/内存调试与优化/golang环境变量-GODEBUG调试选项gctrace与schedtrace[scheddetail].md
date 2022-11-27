# golang环境变量-GODEBUG调试选项gctrace与schedtrace

参考文章

1. [Go 语言运行时环境变量快速导览](https://blog.csdn.net/htyu_0203_39/article/details/50852856)
2. [golang开启GODEBUG gctrace =1 显示信息的含义](https://my.oschina.net/u/2374678/blog/799477)
    - 解释了gctrace的输出各字段含义, golang版本应该>=1.6
3. [go内存泄露case](https://blog.csdn.net/chosen0ne/article/details/46939259)
    - 第1小节解释了gctrace的输出各字段含义, golang版本应该<=1.5
4. [Debugging performance issues in Go programs](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs)
5. [Go 调度器跟踪](https://colobu.com/2016/04/19/Scheduler-Tracing-In-Go/)

`GODEBUG`的值被解释为一个个的`name=value`对, 每一对间由逗号分割, 每一对用于控制go runtime调试工具设施, 例如: 

```
GODEBUG="gctrace=1,schedtrace=1000" godoc -http=:6060
```

上面这条命令用于运行godoc程序时开启`GC tracing and schedule tracing`.

常用的调试工具也就是这两个了: `gctrace`与`schedtrace`

```c++
// Holds variables parsed from GODEBUG env var.
struct DebugVars
{
	int32	gctrace;
	int32	schedtrace;
	int32	scheddetail;
};
```

## gctrace

```
GODEBUG="gctrace=1" godoc -http=:6060
```

可选值为0或1, 一般都使用1表示打开gc调试功能.

`gctrace`信息的输出格式随golang的不同版本会发生变化(见参考文章2和3), 但总是能发现共性的东西. 如: 每一GC 阶段所花费的时间量, heap size 的变化量, 也包括每一GC阶段完成时间, 相对于程序启动时的时间, 当然老版本go可能省略一些信息.

参考文章1有一部分没有翻译, 我自己理解了一下.

`gctrace`的值大于1时会导致gc在每个周期内执行2次, 可能会造成程序的某些需要2次gc来完成的部分提前完成. 开发者不应该使用这样的选项去做程序最终版的性能调试, 可能会在实际场景中造成一些不必要的麻烦.

### The heap scavenger()

> scavenger: 食腐动物, 拾荒者

到目前为止, `gctrace`给出的最有用的信息就是`the heap scavenger`的输出.

```
scvg0: inuse: 22, idle: 40, sys: 63, released: 0, consumed: 63 (MB)
...
scvg143: inuse: 8, idle: 104, sys: 113, released: 104, consumed: 8 (MB)
```

> 我观察到`idle`与`released`的值一般相同, 应该表示空闲的内存分页都应该被释放; 另外`inuse+idle`的值等于`sys`减1, 应该表示向操作系统申请的`sys`个内存页被划分为`inuse`和`idle`两类, 多余的1应该可以不关心.

`scavenger`的工作就是周期性地打扫heap中无用的操作系统内存分页, 它会向操作系统发出请求回收无用内存页.

当然并不能强迫操作系统立刻就去做回收处理, 操作系统可以忽略此建义, 或是延迟回收, 比如直到可分配的空闲内存不够的时候. 

`scavenger`输出的信息是我们了解go程序虚拟内存空间使用情况的最好方式, 当然你也可以通过其它工具, 如`free`, `top`来获到这些信息, 不过你应该信任`scavenger`.

## schedtrace

因为go runtime管理着大量的goroutine, 并调度goroutine在操作系统线程集上运行. 这个操作系统线程集, 其实是就是线程池, 所以从外部考察go程序的性能我们不能获取足够的细节信息(`ps`只能看到线程级别的信息, 不能查看语言层面的协程信息), 更谈不上准确分析程序性能. 

故此我们需要通过`schedtrace`直接了解go runtime scheduler的每一个操作.

`schedtrace`的值为协程信息的打印频率, 单位是毫秒. 如果值太小, 比如1, 控制台会频繁打印, 可能会导致无法接受到ctrl-c信号而卡住, 一定要当心!

```console
$ GODEBUG=schedtrace=1000 godoc -http=:6060
SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=6 spinningthreads=1 idlethreads=0 runqueue=0 [0 0 0 0 0 0 0 0]
SCHED 1008ms: gomaxprocs=8 idleprocs=8 threads=20 spinningthreads=0 idlethreads=15 runqueue=0 [0 0 0 0 0 0 0 0]
SCHED 2016ms: gomaxprocs=8 idleprocs=8 threads=20 spinningthreads=0 idlethreads=15 runqueue=0 [0 0 0 0 0 0 0 0]
SCHED 3023ms: gomaxprocs=8 idleprocs=8 threads=20 spinningthreads=0 idlethreads=15 runqueue=0 [0 0 0 0 0 0 0 0]
```

### scheddetail

如果同时使用`scheddetail=1`将使go runtime输出总结性信息时, 一并输出每一个goroutine的状态信息.

```console
$ GOMAXPROCS=2 GODEBUG="schedtrace=1000,scheddetail=1" godoc -http=:6060
SCHED 0ms: gomaxprocs=2 idleprocs=0 threads=5 spinningthreads=0 idlethreads=0 runqueue=0 gcwaiting=0 nmidlelocked=1 stopwait=0 sysmonwait=0
  P0: status=0 schedtick=0 syscalltick=0 m=-1 runqsize=0 gfreecnt=0
  P1: status=1 schedtick=2 syscalltick=0 m=3 runqsize=0 gfreecnt=0
  M4: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
  M3: p=1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
  M2: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=1 dying=0 helpgc=0 spinning=false blocked=false lockedg=-1
  M1: p=-1 curg=17 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=false lockedg=17
  M0: p=-1 curg=-1 mallocing=0 throwing=0 preemptoff= locks=0 dying=0 helpgc=0 spinning=false blocked=true lockedg=1
  G1: status=1(chan receive) m=-1 lockedm=0
  G17: status=6() m=1 lockedm=1
  G2: status=4(force gc (idle)) m=-1 lockedm=-1
  G3: status=4(GC sweep wait) m=-1 lockedm=-1
```

...8核的话协程信息太多了, 使用`GOMAXPROCS`做了一下限制.

这个输出对于调试`goroutines leaking`很有帮助, 不过其它工具, 诸如: `net/http/pprof`好像更有用一些. 
 
深入阅读请看godoc for the runtime package.
