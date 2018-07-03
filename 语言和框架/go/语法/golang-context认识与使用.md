# golang-context认识与使用

参考文章

1. [golang中context包解读](http://www.01happy.com/golang-context-reading/)

2. [Go语言并发模型：使用 context](https://segmentfault.com/a/1190000006744213)

## 1. 认识

golang中的创建一个新的`goroutine`, 并不会返回像其他语言创建进程时返回pid, 也不像创建线程时能得到一个线程对象的引用. 

`goroutine`通过`go`关键字直接执行, 所以我们无法从外部杀死某个`goroutine`. 之前我们用 `channel ＋ select`的方式, 来解决这个问题, 但是有些场景实现起来比较麻烦. 例如由一个请求衍生出的各个`goroutine`之间需要满足一定的约束关系, 以实现一些诸如有效期, 中止routine树, 传递请求全局变量之类的功能. 

google就为我们提供一个解决方案, 开源了`context`包. 使用`context`实现上下文功能约定需要在你的方法的传入参数的第一个传入一个`context.Context`类型的变量. 

## 2. 使用方法

### 2.1 可以手动取消的Context示例

```go
package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
	go func(ctx context.Context) {
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				log.Printf("done")
				// context的Err()方法返回为什么被取消
				log.Printf("%s", ctx.Err())
				return
			default:
				log.Printf("work")
			}
		}
	}(ctx)
	//5秒后手动取消doStuff
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	log.Printf("end")
}
```

```
2018/07/02 17:09:27 work
2018/07/02 17:09:28 work
2018/07/02 17:09:29 work
2018/07/02 17:09:30 work
2018/07/02 17:09:31 done
2018/07/02 17:09:31 context canceled
2018/07/02 17:09:33 end
```

`context`包提供了多种控制相同context下的`goroutine`的方法, `WithCancel`只是其中一个. 与之类似还有`WithTimeout`, `WithDeadline`和`WithValue`.

它们都必须接受一个`context.Context`对象作为第一个参数, 而这个对象一般通过`context.Background()`得到(还有一个方法`context.TODO()`, 看`context`的源码它们两个是完全一样的)

### 2.2 为goroutine设置一个超时时间

以`WithTimeout()`为例, 我们可以为一个`goroutine`设置一个超时时间. 使用WithTimeout方法就不必手动调用cancel方法了, 但cancel方法依然是有效的.

```go
package main

import (
	"context"
	"log"
	"time"
)

func main() {
	// 使用WithTimeout方法就不必手动调用cancel方法了, 但cancel方法依然是有效的.
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go func(ctx context.Context) {
		//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
                log.Printf("done")
				// context的Err()方法返回为什么被取消
				log.Printf("%s", ctx.Err())
				return
			default:
				log.Printf("work")
			}
		}
	}(ctx)
	time.Sleep(7 * time.Second)
	log.Printf("end")
}
```

```
2018/07/02 17:11:27 work
2018/07/02 17:11:28 work
2018/07/02 17:11:29 work
2018/07/02 17:11:30 work
2018/07/02 17:11:31 done
2018/07/02 17:11:31 context deadline exceeded
2018/07/02 17:11:33 end
```