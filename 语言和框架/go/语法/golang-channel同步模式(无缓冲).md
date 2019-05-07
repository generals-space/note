# golang-channel同步模式(无缓冲)

channel与slice不同, 如果在`make()`的时候不指定容量, 那么创建的就是同步channel.

同步channel的使用要注意, 读与写必须要成对存在, 并且由于没有缓冲区, 读和写一定会造成阻塞, 阻塞部分的代码必须写在协程中代码才能正常工作.

以下是一个错误示例

```go
package main

import "log"

func main() {
	channel := make(chan int)
	channel <- 1 // fatal error: all goroutines are asleep - deadlock!

	for i := range channel {
		log.Println(i)
		if len(channel) == 0 {
			break
		}
	}
}
```

------

正确的作法是要把可能阻塞的地方放到协程中.

```go
package main

import (
	"log"
	"time"
)

func main() {
	channel := make(chan int)
	go func() {
		i := <-channel
		log.Println(i)
	}()
	channel <- 1
	time.Sleep(time.Second * 1)
}

```

或是

```go
package main

import (
	"log"
	"time"
)

func main() {
	channel := make(chan int)
	go func() {
		channel <- 1
	}()
	i := <-channel
	log.Println(i)
	time.Sleep(time.Second * 1)
}

```