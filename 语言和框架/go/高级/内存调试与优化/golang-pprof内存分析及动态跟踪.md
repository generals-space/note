# golang-pprof内存分析及动态跟踪

参考文章

1. [Golang 大杀器之性能剖析 PProf](https://github.com/EDDYCJY/blog/blob/master/golang/2018-09-15-Golang%20%E5%A4%A7%E6%9D%80%E5%99%A8%E4%B9%8B%E6%80%A7%E8%83%BD%E5%89%96%E6%9E%90%20PProf.md)
    - 先介绍了使用`net/http/pprof`的使用示例以及对应的web界面的访问方法
    - 然后介绍了使用`go tool pprof`命令行工具查看更直观的显示, 尤其是`top`命令过滤操作及其打印结果的列字段介绍(如`flat`与`cum`)
2. [golang 内存分析/动态追踪](https://lrita.github.io/2017/05/26/golang-memory-pprof/)
    - 详细介绍了`net/http/pprof`开启的web采样信息界面内容(主要是heap部分)各字段的含义
3. [记一次Golang内存分析——基于go pprof](https://yq.aliyun.com/articles/573743)
    - 介绍`GODEBUG='gctrace=1'`输出各字段所表示的含义, 以及`pprof/heap`界面的输出解释, 命令行调用`pprof`的`-alloc_space`, `-inuse_space`选项.

网上关于golang pprof如何使用的文章不少, 但大部分都是告诉你引用`runtime/pprof`, `net/http/pprof`, 做一些操作, 得到一些数据, 生成一些漂亮的图...但是完全不明白这些数据表示什么意思. 找了半天才找到几个写的比较细致的文章, 记录一下性能分析的方法.

## web界面heap的输出

```
heap profile: 3190: 77516056 [54762: 612664248] @ heap/1048576
1: 29081600 [1: 29081600] @ 0x89368e 0x894cd9 0x8a5a9d 0x8a9b7c 0x8af578 0x8b4441 0x8b4c6d 0x8b8504 0x8b2bc3 0x45b1c1
#    0x89368d    github.com/syndtr/goleveldb/leveldb/memdb.(*DB).Put+0x59d
#    0x894cd8    xxxxx/storage/internal/memtable.(*MemTable).Set+0x88
#    0x8a5a9c    xxxxx/storage.(*snapshotter).AppendCommitLog+0x1cc

......

# runtime.MemStats
# Alloc = 2463648064
# TotalAlloc = 31707239480
# Sys = 4831318840
# HeapAlloc = 2463648064
# HeapSys = 3877830656
# HeapIdle = 854990848
# HeapInuse = 3022839808
# HeapObjects = 11908336
# NumGC = 31
# DebugGC = false
```

其中显示的内容会比较多, 但是主体分为2个部分: 第一个部分打印为通过`runtime.MemProfile()`获取的`runtime.MemProfileRecord`记录. 其含义为: 

```
heap profile: 3190(inused objects): 77516056(inused bytes) [54762(alloc objects): 612664248(alloc bytes)] @ heap/1048576(2*MemProfileRate)
1: 29081600 [1: 29081600] (前面4个数跟第一行的一样, 此行以后是每次记录的, 后面的地址是记录中的栈指针)@ 0x89368e 0x894cd9 0x8a5a9d 0x8a9b7c 0x8af578 0x8b4441 0x8b4c6d 0x8b8504 0x8b2bc3 0x45b1c1
#    0x89368d    github.com/syndtr/goleveldb/leveldb/memdb.(*DB).Put+0x59d 栈信息
```

第二部分就比较好理解, 打印的是通过`runtime.ReadMemStats()`读取的`runtime.MemStats`信息. 我们可以重点关注一下

- `Sys`: 进程从系统获得的内存空间, 虚拟地址空间.
- `HeapSys`: 进程从系统获得的堆内存, 因为golang底层使用TCmalloc机制, 会缓存一部分堆内存, 虚拟地址空间.
- `HeapAlloc`: 进程堆内存分配使用的空间, 通常是用户new出来的堆对象, 包含未被gc掉的.
- `PauseNs`: 记录每次gc暂停的时间(纳秒), 最多记录256个最新记录.
- `NumGC`: 记录gc发生的次数.

`Sys` > `HeapSys` > `HeapAlloc`

## 命令行heap

```
[root@b055c5c3461b /]# go tool pprof --seconds 30 http://localhost:6060/debug/pprof/heap
Fetching profile over HTTP from http://localhost:6060/debug/pprof/heap?seconds=30
Please wait... (30s)
Saved profile in /root/pprof/pprof.main.alloc_objects.alloc_space.inuse_objects.inuse_space.003.pb.gz
File: main
Build ID: ee7a3cbdbe73a8bd5f5c0899ae75b1a7631d62b0
Type: inuse_space
Time: May 9, 2019 at 5:31pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 4198.26kB, 100% of 4198.26kB total
Showing top 10 nodes out of 21
      flat  flat%   sum%        cum   cum%
 2612.50kB 62.23% 62.23%  2612.50kB 62.23%  regexp.(*bitState).reset
  561.50kB 13.37% 75.60%   561.50kB 13.37%  golang.org/x/net/html.init
  512.14kB 12.20% 87.80%   512.14kB 12.20%  regexp.progMachine
  512.12kB 12.20%   100%   512.12kB 12.20%  github.com/jinzhu/gorm.(*Scope).GetModelStruct
         0     0%   100%  3636.76kB 86.63%  gitee.com/generals-space/site-mirror-go.git/crawler.NewCrawler
         0     0%   100%   561.50kB 13.37%  gitee.com/generals-space/site-mirror-go.git/crawler.init
         0     0%   100%  3636.76kB 86.63%  gitee.com/generals-space/site-mirror-go.git/model.GetDB
         0     0%   100%   561.50kB 13.37%  github.com/PuerkitoBio/goquery.init
         0     0%   100%  3636.76kB 86.63%  github.com/jinzhu/gorm.(*DB).AutoMigrate
         0     0%   100%  3124.64kB 74.43%  github.com/jinzhu/gorm.(*ModelStruct).TableName
(pprof)
```

- `flat`: 给定函数上运行耗时
- `flat%`: 同上的 CPU 运行耗时总比例
- `sum%`: 给定函数累积使用 CPU 总比例
- `cum`: 当前函数加上它之上的调用运行总耗时
- `cum%`: 同上的 CPU 运行耗时总比例
