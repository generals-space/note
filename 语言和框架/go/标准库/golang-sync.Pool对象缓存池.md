# golang-sync.Pool对象缓存池

参考文章

1. [go语言的官方包sync.Pool的实现原理和适用场景](https://blog.csdn.net/yongjian_lian/article/details/42058893)

2. [go的临时对象池--sync.Pool](https://www.jianshu.com/p/2bd41a8f2254)

3. [Golang 优化之路——临时对象池](https://blog.cyeam.com/golang/2017/02/08/go-optimize-slice-pool)
	- 介绍了golang在内存分配机制, 尤其是对象分配时的栈/堆选择, 并且给出了对应的解决方法-`sync.Pool`, 以及一个基于`sync.Pool`的第三方对象池的优化示例.

...我还以为`sync.Pool`是协程池呢, 原来并不是. 当了解golang的gc机制后再来看`sync.Pool`临时对象池就更容易理解了, 可以见参考文章3.

Pool用于存储那些被分配了但是没有被使用, 而未来可能会使用的值, 以减小垃圾回收的压力. 我们先看一个简单示例, 再分析场景.

```go
package main

import (
	"log"
	"sync"
)

func main() {
	p := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}
	a := p.Get().(int)
	log.Print(a) // 0

	p.Put(1)
	b := p.Get().(int)
	log.Print(b) // 1
}

```

`Pool`对象`p`只有两个方法: `Get()`, `Put()`, 还有一个成员属性`New`. 创建新的`Pool`对象没有办法指定大小, 所以不能限制其资源消耗.

参考文章2中给出了一个使用`sync.Pool`的确能够使性能提升的示例.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// BytePool ...
var BytePool = &sync.Pool{
	New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	},
}

func main() {
	sum := 1000000000
	time1 := time.Now().Unix()

	for i := 0; i < sum; i++ {
		// 这一句比下面一句要快上3s左右, 应该是变量赋值导致的开销更大一些吧.
		// _ = make([]byte, 1024)
		b := make([]byte, 1024)
		_ = b
	}
	time2 := time.Now().Unix()

	for i := 0; i < sum; i++ {
		b := BytePool.Get().(*[]byte)
		BytePool.Put(b)
	}
	time3 := time.Now().Unix()

	fmt.Printf("without pool %ds\n", time2-time1) // 33s
	fmt.Printf("with    pool %ds\n", time3-time2) // 25s
}
```

不过要使用`-gcflags='-l -N'`编译选项, 因为在参考文章2的评论区, 有人说使用了`sync.Pool`反而变慢了, 尤其是第1个for循环竟然只用了1s甚至0s, 这简直是不可能的. 使用这个编译选项是为了禁用编译器优化, 这样才能看出差别.

```
$ go run -gcflags='-l -N' .\main.go
without pool 36s
with    pool 29s
```

嗯, 这个`[]byte`对象的示例应用场景倒是不难想. 在编写socket服务器时, 或是频繁发起http请求时, 对于响应结果的读取总是需要用`[]byte`放数据的.

...但是从缓存池里拿出的`[]byte`对象不会有可能是存放过数据的对象吗? 使用时还要手动清空一下吗? 难道就是节省了分配/回收内存时的开销?

参考文章1中提出的几个很重要的问题还没看懂...日后再说吧.
