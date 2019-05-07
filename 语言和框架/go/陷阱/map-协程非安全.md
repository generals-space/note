# map-协程非安全

我们都知道`map`是协程非安全的, 所以可能会编写自定义的结构为`map`成员的操作加锁.

但是不只是写操作, `map`成员的读操作也是要加锁的, 不然会出现`fatal error: concurrent map read and map write`

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// UserAges ...
type UserAges struct {
	ages map[string]int
	sync.Mutex
}

// Add ...
func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

// Get ...
func (ua *UserAges) Get(name string) int {
	// 如果读操作没有加锁, 则会出现`fatal error: concurrent map read and map write`
	// ua.Lock()
	// defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	counter := 10
	prefix := "general"
	userAges := &UserAges{
		ages: map[string]int{},
	}
	for i := 0; i < counter; i++ {
		go func(ord int) {
			name := fmt.Sprintf("%s-%d", prefix, ord)
			userAges.Add(name, ord)
		}(i)
	}

	for i := 0; i < counter; i++ {
		go func(ord int) {
			name := fmt.Sprintf("%s-%d", prefix, ord)
			age := userAges.Get(name)
			fmt.Printf("name: %s, age: %d\n", name, age)
		}(i)
	}
	// 需要手动ctrl-c
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-sigCh)
}

```

解开`Get()`方法中的加锁操作注释, 可以得到正确的输出

```
name: general-5, age: -1
name: general-9, age: -1
name: general-6, age: 6
name: general-7, age: 7
name: general-8, age: -1
name: general-2, age: 2
name: general-3, age: 3
name: general-4, age: 4
name: general-0, age: 0
name: general-1, age: 1
interrupt // 这里需要手动ctrl-c
```

------

突然发现这种情况用读写锁而不是互斥锁效率更高才对, 所以下面是一个读写锁的示例, 读的时候加读锁, 不影响其他读操作, 写的时候才加常规锁.

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// UserAges ...
type UserAges struct {
	ages map[string]int
	sync.RWMutex
}

// Add ...
func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

// Get ...
func (ua *UserAges) Get(name string) int {
	// 如果读操作没有加锁, 则会出现`fatal error: concurrent map read and map write`
	// ua.RLock()
	// defer ua.RUnlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	counter := 10
	prefix := "general"
	userAges := &UserAges{
		ages: map[string]int{},
	}
	for i := 0; i < counter; i++ {
		go func(ord int) {
			name := fmt.Sprintf("%s-%d", prefix, ord)
			userAges.Add(name, ord)
		}(i)
	}

	for i := 0; i < counter; i++ {
		go func(ord int) {
			name := fmt.Sprintf("%s-%d", prefix, ord)
			age := userAges.Get(name)
			fmt.Printf("name: %s, age: %d\n", name, age)
		}(i)
	}
	// 需要手动ctrl-c
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-sigCh)
}

```