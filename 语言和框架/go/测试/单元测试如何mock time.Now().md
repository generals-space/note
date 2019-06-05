# 单元测试如何mock time.Now()

参考文章

1. [如何在单元测试过程中模拟日期和时间](https://www.cnblogs.com/redmoon/p/4433609.html)

2. [Is there an easy way to stub out time.Now() globally in golang during test?](https://stackoverflow.com/questions/18970265/is-there-an-easy-way-to-stub-out-time-now-globally-in-golang-during-test)

代码中有按照请求触发时的当前时间进行数据库查询, 如何插入有效的数据库记录?

在google中搜索`golang mock time.Now()`会有很多结果, 模拟当前时间是一个非常常见的测试需求, 各种语言都有.

参考文章2中第二高票回答提到的`bou.ke/monkey`很有意思, 他做了和python中gevent()类似的事, 可以为指定函数打补丁, 使之行为模式变成我们可控的模样.

```go
package main

import (
	"fmt"
	"os"

	"bou.ke/monkey"
)

func main() {
	patches := monkey.Patch(fmt.Printf, func(format string, a ...interface{}) (n int, err error) {
		fmt.Println("monkeyed")
		// fmt.Printf()内部调用了Fprintf()实现, 这里不能再使用fmt.Printf()了, 会死循环的.
		return fmt.Fprintf(os.Stdout, format, a...)
	})
	fmt.Printf("hello world: %d\n", 100)
	patches.Unpatch()
	fmt.Printf("hello world: %d\n", 100)
}

```

运行输出为

```
monkeyed
hello world: 100
hello world: 100
```

貌似monkey的license许可有问题, 不能在生产环境使用...不过这个问题不在我们考虑范围...?

但是这样打的补丁函数不能随意修改函数签名, 否则代码中使用的函数无法执行. 要根据某个变量控制行为, 就需要通过monkey本身的闭包实现了.

```go
package main

import (
	"fmt"
	"os"

	"bou.ke/monkey"
)

func main() {
	var monkeyIt bool
	patches := monkey.Patch(fmt.Printf, func(format string, a ...interface{}) (n int, err error) {
		if monkeyIt == true {
			fmt.Println("monkeyed")
		}
		// fmt.Printf()内部调用了Fprintf()
		return fmt.Fprintf(os.Stdout, format, a...)
	})
	fmt.Printf("hello world: %d\n", 100)
	monkeyIt = true
	fmt.Printf("hello world: %d\n", 100)
	patches.Unpatch()
	fmt.Printf("hello world: %d\n", 100)
}
```

执行输出为

```
hello world: 100
monkeyed
hello world: 100
hello world: 100
```

这样我们就可以通过局部变量对标准库函数打补丁了...`Time.Now()`自然也不在话下.

```go
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fakeNow, _ := time.ParseInLocation("2006-01-02 15:04:05", "2019-01-01 12:00:00", loc)
	timenowPatcher := monkey.Patch(time.Now, func() time.Time {
		return fakeNow
	})
	defer timenowPatcher.Unpatch()
```
