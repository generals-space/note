# go build gcflags参数-m打印编译器优化描述

参考文章

1. [从golang的垃圾回收说起（下篇） 5.1节](https://sq.163yun.com/blog/article/192800393799778304)

2. [参考golang官网(FAQ) How do I know whether a variable is allocated on the heap or the stack?](https://golang.org/doc/faq#stack_or_heap)

Q: 
我怎样知道一个变量被分配到堆上还是栈上?

A: 
站在正确性的立场上, 开发者并不需要知道. 在go中, 每个变量只要被引用就会一直存在, 而实现存储位置的选择方式与语言本身的语义无关.

当然变量存储位置对编写高效程序来说的确会有影响. 在可能的情况下, GO编译器会将函数内的局部变量分配到函数的栈帧. 但是当编译器确认变量在函数返回后仍被引用时, 会将变量分配到可被GC的堆上以保证不会出现空指针的错误. 另外, 如果一个局部变量非常大, 那ta很有可能被分配在堆上而不是栈上.

在现在的编译器中, 如果一个变量被取了地址, 那ta就有可能被分配到堆上. 不过一个基本的`逃逸分析(escape analysis)`会在某些情况下(比如这些被取地址的变量在函数返回后就不会再被引用)将其分配到栈上.

参考文章1给出了一个示例, 来展示不同对象的分配情况, 使用`-gcflags`选项的`-m`参数可以清楚地看到编译器对变量的处理过程.

go版1.11.1

```go
package main

func foo() *int {
	// 被取地址的局部变量会发生escape逃逸被分配到堆上
	var x int
	return &x
}

func bar() int {
	// 这里却被分配在了栈上, 并没有发生逃逸
	// ...难道因为是直接创建的指针(本质上是一个整数)?
	x := new(int)
	*x = 1
	return *x
}

func big() {
	len := 10
	// 较小局部变量被分配到栈中
	x := make([]int, 0, 20)
	// 较大局部变量被分配到堆上
	y := make([]int, 0, 20000)
	// 这里为什么被分配到了堆上?
	z := make([]int, 0, len)
}

func main() {
	// 就算没有运行, 也会打印出编译器的优化过程
}
```

编译过程如下

```
$ go run -gcflags='-l -m -v' .\main.go
# command-line-arguments
.\main.go:6:9: &x escapes to heap
.\main.go:5:6: moved to heap: x
.\main.go:12:10: bar new(int) does not escape
.\main.go:22:11: make([]int, 0, 20000) escapes to heap
.\main.go:24:11: make([]int, 0, len) escapes to heap
.\main.go:20:11: big make([]int, 0, 20) does not escape
.\main.go:20:2: x declared and not used
.\main.go:22:2: y declared and not used
.\main.go:24:2: z declared and not used
```
