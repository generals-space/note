# golang-time与duration时间加减操作

参考文章

1. [golang time and duration](https://studygolang.com/articles/5016)

2. [golang time 时间的加减法](https://studygolang.com/articles/8919)

```go
package main

import "log"
import "time"

func main(){
    now := time.Now()

    log.Printf("%s\n", now)

    // 创建一个Date时间对象
    then, _ := time.Parse("2006-01-02 15:04:05", "2012-01-01 00:00:00")
    log.Printf("%s\n", then)

    // Before, After, Equal比较两个时间对象, 在秒级单位判断谁先谁后
    log.Printf("then < now ? %t\n", then.Before(now))
    log.Printf("then > now ? %t\n", then.After(now))
    log.Printf("then = now ? %t\n", then.Equal(now))

    // 两个Date对象相减得到Duration对象
    delta := now.Sub(then)
    log.Printf("%T\n", delta)
    // 一个Date对象加上一个Duration对象得到一个新的Date对象
    future := now.Add(delta)
    log.Printf("%T\n", future)
    log.Printf("%s\n", future)

    // 也可以用一个负号-表示减去Duration
    then2 := now.Add(-delta)
    log.Printf("%T\n", then2)
    log.Printf("%s\n", then2)

    // 我们可以获取一个Duration对象一共有几个小时, 或几分钟, 几秒
    log.Printf("Hours: %f\n", delta.Hours())
    log.Printf("Hours: %f\n", delta.Minutes())
    log.Printf("Hours: %f\n", delta.Seconds())
}
```

------

有的时候, 我们知道一个时间点, 只想直接获取到它的前/后的一小时, 一天或一周的时间点, 这种民情况下, 我们需要直接构建出`duration`对象(再找两个`time`对象相减就不太好了吧).

`time`包里构建`duration`对象可以用`time.ParseDuration()`函数, 使用示例如下

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%s\n", now.Format(time.RFC3339))			// 2018-06-06T18:21:32+08:00

	// 可用单位只有 "ns", "us" (or "µs"), "ms", "s", "m", "h"
	d1, _ := time.ParseDuration("-24h")

	oneHourAgo := now.Add(d1)
	fmt.Printf("%s\n", oneHourAgo.Format(time.RFC3339))		// 2018-06-05T18:21:32+08:00

	oneYearAgo := now.Add(d1 * 365)
	fmt.Printf("%s\n", oneYearAgo.Format(time.RFC3339))		// 2017-06-06T18:21:32+08:00
}
```