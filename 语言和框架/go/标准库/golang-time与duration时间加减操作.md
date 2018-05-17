# golang-time与duration时间加减操作

参考文章

1. [golang time and duration](https://studygolang.com/articles/5016)

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