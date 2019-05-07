# golang环境变量-`GOTRACEBACK`异常信息等级

参考文章

1. [Go 语言运行时环境变量快速导览](https://blog.csdn.net/htyu_0203_39/article/details/50852856)

`GOTRACEBACK`用于控制当异常发生时, 系统提供信息的详细程度.

------

在go 1.5, `GOTRACEBACK`有4个值. 

`0`: 只输出panic异常信息. 
`1`: 此为go的默认设置值, 输出所有goroutine的`stack traces`, 但不包括与go runtime相关的`stack frames`.
`2` 在1的基础上, 还输出与go runtime相关的`stack frames`, 从而了解哪些goroutines是由go runtime启动运行的. 
`crash`: 在2的基础上, go runtime触发进程segfault错误, 从而生成`core dump`, 当然要操作系统允许的情况下, 而不是调用`os.Exit`. 

`GOTRACEBACK` 在go 1.6中的变化

- `none`: 只输出panic异常信息. 
- `single`: 只输出被认为引发panic异常的那个goroutine的相关信息. 
- `all`: 输出所有goroutines的相关信息, 除去与go runtime相关的stack frames.
- `system`: 输出所有goroutines的相关信息, 包括与go runtime相关的stack frames,从而得知哪些goroutine是go runtime启动运行的. 
- `crash`: 与go 1.5相同, 未变化. 

为了与go 1.5兼容, `0`对应`none`, `1`对应`all`, 以及`2`对应`system`.

注意: 在go 1.6中, 默认只输出引发panci异常的goroutine的stack trace.

------

以如下代码为例, golang版本1.11.1

```go
package main

import "fmt"

func layer2() {
	fmt.Println("layer 2")
	panic("kerboom")
}

func layer1() {
	fmt.Println("layer 1")
	layer2()
}

func main() {
	layer1()
}

```

常规运行结果:

```
layer 1
layer 2
panic: kerboom

goroutine 1 [running]:
main.layer2()
        /root/main.go:7 +0x79
main.layer1()
        /root/main.go:12 +0x62
main.main()
        /root/main.go:16 +0x20
exit status 2
```

使用`GOTRACEBACK=0`或`GOTRACEBACK=none`的结果如下

```
$ GOTRACEBACK=0 go run main.go
layer 1
layer 2
panic: kerboom
exit status 2
$ GOTRACEBACK=none go run main.go
layer 1
layer 2
panic: kerboom
exit status 2
```

其余的数值与等级的对应关系与上述所说一致, 目前还不清楚`single`与`all`的区别.

至于`crash`标记, 通过如下命令开启`core dump`记录后, 也的确有效. 会在当前目录生成一个名为`core`的核心转储文件.

```
ulimit -c unlimited    # 设置core大小为无限制
ulimit unlimited    # 设置文件大小为无限制(可以不执行这句, 防止core文件过大)
```