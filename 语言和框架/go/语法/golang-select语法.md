# golang-select语法

参考文章

1. [Golang select的使用及典型用法](https://blog.csdn.net/zhaominpro/article/details/77570290)

2. [Go 语言 select 语句](https://wizardforcel.gitbooks.io/w3school-go/content/13.html)

3. [由浅入深聊聊Golang中select的实现机制](https://blog.csdn.net/u011957758/article/details/82230316)
    - 一对select的陷阱示例
    - select在应用中的运行机制
    - select底层源码分析, 文末的流程图值得一看

## 1. 基本用法 

`select`是Go中的一个控制结构，类似于`switch`语句，用于处理异步IO操作。`select`会监听`case`语句中`channel`的读写操作，当`case`中`channel`读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。

> `select`中的`case`语句必须是一个`channel`操作

> `select`中的`default`子句总是可运行的

如果有多个`case`都可以运行，select会**随机公平**地选出一个执行，其他不会执行。

如果没有可运行的`case`语句，且有`default`语句，那么就会执行`default`的动作。

如果没有可运行的`case`语句，且没有`default`语句，`select`将阻塞，直到某个`case`通信可以运行

1. select+case是用于阻塞监听goroutine的，如果没有case，就单单一个select{}，则为监听当前程序中的goroutine，此时注意，需要有真实的goroutine在跑，否则select{}会报panic
2. select底下有多个可执行的case，则随机执行一个。
3. select常配合for循环来监听channel有没有故事发生。需要注意的是在这个场景下，break只是退出当前select而不会退出for，需要用break TIP / goto的方式。
4. 无缓冲的通道，则传值后立马close，则会在close之前阻塞，有缓冲的通道则即使close了也会继续让接收后面的值
5. 同个通道多个goroutine进行关闭，可用recover panic的方式来判断通道关闭问题
