# golang-异常处理机制panic, defer与recover的应用

参考文章

1. [golang异常panic和恢复recover用法](https://blog.csdn.net/ghost911_slb/article/details/7831574)

2. [关于golang的panic recover异常错误处理](http://xiaorui.cc/2016/03/09/%E5%85%B3%E4%BA%8Egolang%E7%9A%84panic-recover%E5%BC%82%E5%B8%B8%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86/)

3. [Go的异常处理 defer, panic, recover<转载>](https://blog.csdn.net/newsyoung1/article/details/39369667)

golang 官方不推荐统一处理异常(如`try..catch..`), 你需要手动处理每一个返回的错误. 这个虽然有争议，但是支持的人也是很多的, 牺牲了代码的简洁性但是增加了可维护性.

但本文讨论的异常与`error`这种处理的返回值异常不太相符, 通过`panic`报出的异常那就是真的异常了.

参考文章1的解释简洁明了, 示例也是相当简单;

参考文章2中, 峰云对go的异常处理机制的判断很客观, 也讲解了几点需要注意的地方, 很实际;

参考文章3的示例最详尽, 分别对`defer`, `panic`, `recover`进行了解释;

> go中可以抛出一个`panic`的异常，然后在`defer`中通过`recover`捕获这个异常，然后正常处理.

> 在一个主进程，多个go程处理逻辑的结构中，这个很重要，如果不用recover捕获panic异常，会导致整个进程出错中断.

## 1. 一个简单示例

关于异常处理程序编写, 要注意以下几点

1. `defer`需要放在`panic`之前定义, 另外`recover`只有写在`defer`调用的函数中才有效.

2. `recover`处理异常后，逻辑并不会恢复到`panic`那个点去，而是会跑到所属函数的外层.

```go
package main

import "fmt"

func main() {
	defer func() { //必须要先声明defer, 否则不能捕获到panic异常
		fmt.Println("b")
		if err := recover(); err != nil {
			//这里的err其实就是panic传入的内容, 可以是任意值, 任意类型
			fmt.Println(err)
		}
		fmt.Println("c")
	}()
	fmt.Println("a")
	panic("panic")
	fmt.Println("d")
}
```

上述代码的执行结果为

```
a
b
panic
c
```

我们可以看到, 字符串`d`没有打印出来. 因为`defer`是在函数退出前执行的动作, 所以它处理完函数也就执行完了, 控制权会交给外层函数(如果是在`main()`里其实就已经退出了).

再一个易懂的例子

```go
package main

import "fmt"

func myfunc() {
	defer func() { //必须要先声明defer, 否则不能捕获到panic异常
		fmt.Println("b")
		if err := recover(); err != nil {
			//这里的err其实就是panic传入的内容, 可以是任意值, 任意类型
			fmt.Println(err)
		}
		fmt.Println("c")
	}()
	fmt.Println("a")
	panic("panic")
	fmt.Println("d")
}

func main() {
	fmt.Println("before myfunc")
	myfunc()
	fmt.Println("after myfunc")
}
```

这段代码的执行结果为

```
before myfunc
a
b
panic
c                   // myfunc()中的defer执行完毕
after myfunc        // 这里跳出了myfunc()
```

ok, 那么实际应用场景是怎样的? 

可以参考一下`gin`框架的方式, `gin`提供了`gin.Recovery()`中间件, 在映射路由前`Use`这个中间件, 如下

```go
ret := gin.New()
ret.Use(gin.Recovery())
ret.GET("/", controller.IndexAction)
ret.POST("/", controller.IndexAction)
```

查看`gin.Recovery`的源码一部分如下

```go
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				if logger != nil {
					stack := stack(3)
					httprequest, _ := httputil.DumpRequest(c.Request, false)
					logger.Printf("[Recovery] panic recovered:\n%s\n%s\n%s%s", string(httprequest), err, stack, reset)
				}
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
```

按照`defer` + `recover`只能捕捉之后的异常, 并且无法恢复到异常代码之后的代码的这种特性, 优先加载`Recovery`中间件后, `c.Next()`中, 也就是控制器, 实际的业务逻辑中出现的异常, 都会被`Recovery`捕捉, 不会造成http主服务崩溃.