# 并发fmt

参考文章

1. [internal/poll: inconsistent poll.fdMutex while println string](https://github.com/golang/go/issues/25558)

2. [internal/poll: better panic message for lock overflow](https://go-review.googlesource.com/c/go/+/119956)

情景描述

go 1.10.3

```
internal/poll: inconsistent poll.fdMutex while println string
```

遍历一个128位的大数, 为每个值进行一些计算操作. 每个值的计算的过程都丢到goroutine里, goroutine中用`fmt`打印当前处理的值. 在不限协程数量(遍历时不断把任务开协程执行)运行了几分钟后, 报上述错误.

按照参考文章1的说法, 是因为fmt在打印到终端时貌似会加锁, 而出现这个问题的原因在于, 锁的数量是有限的, 最多为`1048575`. 当所有协程同时向终端打印时, 就出现了这个问题...

不过对fmt还没听过怎么手动给ta加锁, 所以解决办法, 要么限制协程数量(用协程池), 要么把`fmt`语句从协程里删掉, 就可以了.
