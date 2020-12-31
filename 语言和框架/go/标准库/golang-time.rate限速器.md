# golang-time.rate限速器

<!key!>: {5E8AD061-7897-4D0F-A0DA-232CD84367FE}

<!link!>: {A03F8671-228A-4BE2-821F-05E983245D7B}

参考文章

1. [令牌桶算法_golang_rate 对速度进行限制 ](http://www.liuvv.com/2018/05/19/%E4%BB%A4%E7%89%8C%E6%A1%B6%E7%AE%97%E6%B3%95_golang_rate%E5%AF%B9%E9%80%9F%E5%BA%A6%E8%BF%9B%E8%A1%8C%E9%99%90%E5%88%B6/)
    - github: unix2dos.github.io
    - 介绍了两种限流算法: `漏桶算法`与`令牌桶算法`, 十分生动, 易于理解.
    - 好吧, 后来发现关于漏桶与令牌桶算法是经典算法, 网上的资料大同小异, 没有必要专门看这篇.
    - 不过文章中关于`golang.org/x/time/rate`的代码应该是原创的, 可以参考.
2. [Golang实现请求限流的几种办法](https://blog.csdn.net/micl200110041/article/details/82013032)
    - 介绍了两种通过channel和计时器完成的限速器设计及代码
    - 介绍了基于IP地址或是接口作为标识符实现的独立限速器, 相对于全局限速器更加灵活.
3. [Golang time/rate 限速器](https://www.jianshu.com/p/1ecb513f7632)
4. [golang服务器开发频率限制 golang.org/x/time/rate 使用说明](https://www.jianshu.com/p/4ce68a31a71d)

关于限速器, 最初接触这个概念是因为在微服务架构中需要实现请求限流的操作. 限流的应用场景是, 当我们的服务处理请求时需要消耗大量的CPU/内存资源, 请求数量过多会导致服务卡顿或崩溃. 限流目的就是为我们的服务设置一个阈值, 超过这个阈值的请求不做处理直接返回, 以保证正在处理的请求能够得以快速响应, 同时也为服务本身能够正常工作提供了保障(听起来像是为了提高响应速度牺牲了一部分吞吐量啊...).

关于微服务中的限流算法一般是"令牌桶算法", 但实际上限流算法不只一种, 应用场景也不只微服务架构. 参考文章1中介绍了两种限流算法: `漏桶算法`与`令牌桶算法`, 可以借鉴一下.

本文的重点不在于讨论限速算法的优劣, 而是golang中已经实现的限速算法.

目前我所知道的在golang中实现的令牌算法有3种:

1. [time/rate](https://golang.org/x/time/rate)
2. [juju/ratelimit](https://github.com/juju/ratelimit)
3. [didip/tollbooth](https://github.com/didip/tollbooth)

ta们应该是都实现了令牌桶限速工具库, 使用方法大同小异, 本文介绍的是`time/rate`包的使用方法.

该包提供`rate.Limiter`对象, 实例化方法如下.

```go
func NewLimiter(r Limit, b int) *Limiter {
	return &Limiter{
		limit: r,
		burst: b,
	}
}
```

解释一下, 令牌桶容量为`b`, 最开始装满令牌, 然后每秒往里面填充`r`个令牌(多余的令牌会被丢弃, 桶中令牌总量不会超过`b`).

在这里, 取得一个令牌的请求被称为"事件". 由于令牌池中最多有`b`个令牌, 所以一次最多只能允许`b`个事件发生, 一个事件花费掉一个令牌. 

```go
//第一个参数为每秒发生多少次事件, 第二个参数是最大可运行多少个事件
l := rate.NewLimiter(1, 3) 
```

得到的`rate.Limiter`提供了三类方法用来限速:

1. `Wait`/`WaitN` 当没有可用或足够的事件时, 将阻塞等待. **推荐实际程序中使用这个方法**
2. `Allow`/`AllowN` 当没有可用或足够的事件时, 返回false
3. `Reserve`/`ReserveN` 当没有可用或足够的事件时, 返回 Reservation, 和要等待多久才能获得足够的事件. 

看一看源码你就会发现, `Wait() <=> WaitN(time.Now(), 1)`, `Allow`与`Reserve`也是同理.

既然推荐在程序中使用`Wait`/`WaitN`, 这里我们就使用`Wait(N)`来验证ta提供的令牌桶的运行机制. 实际上也比`Allow(N)`, `Reserve(N)`来认识令牌桶更容易理解.

## 1. Wait(N)

`WaitN`阻塞当前直到`limit`允许`n`个事件的发生(从桶中取出n个令牌). 

如果 n 超过了令牌池的容量大小则报错. 
如果 Context 被取消了则报错. 
如果 limit 的等待时间超过了 Context 的超时时间则报错. 

如果你需要在事件超出频率的时候丢弃或跳过事件, 就使用`AllowN`, 否则使用`Reserve`或`Wait`(就是说如果此时桶中令牌为空, `AllowN`会直接返回`false`, 而`Reserve`/`Wait`会阻塞以等待可用令牌).

```go
package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// 最初桶中有3个令牌, 每秒向其中新增1个令牌.
	limiter := rate.NewLimiter(1, 3)
	ctx, cancel := context.WithCancel(context.TODO()) // 好像不能用WithTimeout()...<???>
	defer cancel()

	log.Printf("令牌桶容量burst: %d, 令牌更新速率limit: %+v", limiter.Burst(), limiter.Limit())
	// 程序运行时间
	duration := time.Second * 5
	// 计时开始
	startT := time.Now()
	for {
		// Wait是从桶中取1个令牌, WaitN可以取N个令牌.
		limiter.Wait(ctx)
		// limiter.WaitN(ctx, 1)
		log.Println("执行")

		if time.Now().Sub(startT) >= duration {
			log.Println("结束")
			break
		}
	}
}
```

执行结果如下

```
2019/05/18 11:57:01 令牌桶容量burst: 3, 令牌更新速率limit: 1
2019/05/18 11:57:01 执行
2019/05/18 11:57:01 执行
2019/05/18 11:57:01 执行
2019/05/18 11:57:02 执行
2019/05/18 11:57:03 执行
2019/05/18 11:57:04 执行
2019/05/18 11:57:05 执行
2019/05/18 11:57:06 执行
2019/05/18 11:57:06 结束
```

可以看到在程序在第一秒直接循环了3次, 取出了3个令牌, 之后就会阻塞等待每秒放入桶中的令牌.

## 2. Reserve(N)

当没有可用事件时返回`Reservation`对象, 标识调用者需要等多久才能等到 n 个事件发生 (意思就是等多久令牌池中至少含有 n 个令牌). 

如果 ReserveN 传入的 n 大于令牌池的容量 b, 那么返回 false.

如果希望根据频率限制等待和降低事件发生的速度而不丢掉事件, 就使用这个方法. 

我认为这里要表达的意思就是如果事件发生的频率是可以由调用者控制的话, 可以用`ReserveN`来控制事件发生的速度而不丢掉事件. 如果要使用 context 的超时控制或`cancel`方法的话, 使用`WaitN`. 

```go
package main

import (
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(1, 3)
	log.Printf("令牌桶容量burst: %d, 令牌更新速率limit: %+v", limiter.Burst(), limiter.Limit())
	// 程序运行时间
	duration := time.Second * 5
	// 计时开始
	startT := time.Now()

	for {
		r := limiter.Reserve()
		// r := limiter.ReserveN(time.Now(), 1)
		s := r.Delay()
		time.Sleep(s)
		log.Println("执行")

		if time.Now().Sub(startT) >= duration {
			log.Println("结束")
			break
		}
	}
}
```

执行结果如下

```
2019/05/18 12:34:35 令牌桶容量burst: 3, 令牌更新速率limit: 1
2019/05/18 12:34:35 执行
2019/05/18 12:34:35 执行
2019/05/18 12:34:35 执行
2019/05/18 12:34:36 执行
2019/05/18 12:34:37 执行
2019/05/18 12:34:38 执行
2019/05/18 12:34:39 执行
2019/05/18 12:34:40 执行
2019/05/18 12:34:40 结束
```

没有使用context, 程序表现与上例中`Wait(N)`相同.

## 3. Allow(N)

AllowN 标识在时间 now 的时候, n 个事件是否可以同时发生 (也意思就是 now 的时候是否可以从令牌池中取 n 个令牌). 如果你需要在事件超出频率的时候丢弃或跳过事件, 就使用 AllowN, 否则使用 Reserve 或 Wait.

```go
package main

import (
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limiter := rate.NewLimiter(1, 3)
	log.Printf("令牌桶容量burst: %d, 令牌更新速率limit: %+v", limiter.Burst(), limiter.Limit())

	// 程序运行时间
	duration := time.Second * 5
	// 计时开始
	startT := time.Now()

	for {
		if limiter.Allow() {
			log.Println("task done")
		} else {
			log.Println("task died")
		}
		// 每秒5个请求
		time.Sleep(time.Millisecond * 200)

		if time.Now().Sub(startT) >= duration {
			log.Println("结束")
			break
		}
	}
}

```

执行结果如下

```
2019/05/18 12:49:29 令牌桶容量burst: 3, 令牌更新速率limit: 1
2019/05/18 12:49:29 task done
2019/05/18 12:49:29 task done
2019/05/18 12:49:29 task done
2019/05/18 12:49:30 task died
2019/05/18 12:49:30 task died
2019/05/18 12:49:30 task done
2019/05/18 12:49:30 task died
2019/05/18 12:49:30 task died
2019/05/18 12:49:31 task died
2019/05/18 12:49:31 task died
2019/05/18 12:49:31 task done
2019/05/18 12:49:31 task died
2019/05/18 12:49:31 task died
2019/05/18 12:49:32 task died
2019/05/18 12:49:32 task died
2019/05/18 12:49:32 task done
2019/05/18 12:49:32 task died
2019/05/18 12:49:32 task died
2019/05/18 12:49:33 task died
2019/05/18 12:49:33 task died
2019/05/18 12:49:33 task done
2019/05/18 12:49:33 task died
2019/05/18 12:49:33 task died
2019/05/18 12:49:34 task died
2019/05/18 12:49:34 task died
2019/05/18 12:49:34 结束
```

仔细分析就会发现, 只有第一秒中的前3个请求成功了, 之后的每一秒内, 5个请求只有1个是成功的.

## 补充

`time/rate`包只提供了两个方法: `NewLimiter`与`Every`.

关于`Every()`, 在上述创建Limiter对象时, 向桶中添加令牌的速率单位是秒, 如果我想以分钟/小时为单位添加呢? 

这就需要使用`Every()`方法了. `Every()`返回Limit对象, 正是`NewLimiter`所需的第1个参数. 所以自定义令牌刷新频率一般要这样写

```go
limiter := rate.NewLimiter(rate.Every(time.Second*1), 5)
limiter := rate.NewLimiter(rate.Every(time.Minute*1), 5)
limiter := rate.NewLimiter(rate.Every(time.Hour*1), 5)
```

但是`Limit`(不是`Limiter`)对象其实就是float64的别名, `Every()`也只是做了个数学计算而已.
