# golang-sync之WaitGroup

参考文章

1. [golang语言异步通信之WaitGroup](https://www.jianshu.com/p/a89beb04ef15)

golang的`sync.WaitGroup`作用类似于Python多线程中的`join`, 用于等待全部子线程子线程执行结束的.

来一个最简示例.

```go
package main

import (
	"log"
	"net/http"
	"sync"
)

const aioURL = "http://note.generals.space:3000/aio"

func getURL(wg *sync.WaitGroup, i uint64) {
	defer wg.Done()

	resp, err := http.Get(aioURL)
	if err != nil {
		log.Printf("error in %d: %s", i, err.Error())
		return
	}
	defer resp.Body.Close()

	var buf = make([]byte, 512)
	length, _ := resp.Body.Read(buf)
	log.Printf("order %d: %s", i, buf[:length])
}

func main() {
	wg := sync.WaitGroup{}
	var i uint64
	for i = 0; i < 10; i++ {
        wg.Add(1)
        // WaitGroup对象不是一个引用类型, 这里必须要传递指针
		go getURL(&wg, i)
	}
	wg.Wait()
}
```

需要注意的几点:

1. `WaitGroup`的值不能是负值, 当`Done()`操作多于`Add()`操作时会出现这个问题.

2. 当`Add()`操作多于`Done()`时, 并且已经没有goroutine在运行的时候, 会出现死锁, 程序会崩溃.

具体可以见参考文章1, 介绍的比较清楚.