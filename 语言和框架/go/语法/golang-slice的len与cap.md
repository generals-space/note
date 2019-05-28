# golang-slice的len与cap

参考文章

1. [从golang的垃圾回收说起（下篇） 5.3节](https://sq.163yun.com/blog/article/192800393799778304)

## 1. `make([]int, 0, 10)`

```go
	// 创建切片, len=0, cap=10, 底层实际上分配了10个元素大小的空间
	s1 := make([]int, 0, 10)
	// 此时切片为空, 长度为0, 无法通过s1[0]访问内容, 会越界
	fmt.Printf("s1 len: %d, cap: %d, %+v\n", len(s1), cap(s1), s1) // s1 len: 0, cap: 10, []
	fmt.Printf("s1 address: %p\n", &s1)                            // s1 address: 0xc000004440
	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
	}
	fmt.Printf("s1 len: %d, cap: %d, %+v\n", len(s1), cap(s1), s1) // s1 len: 10, cap: 10, [0 1 2 3 4 5 6 7 8 9]
	fmt.Printf("s1 address: %p\n", &s1)                            // s1 address: 0xc000004440

	// 当cap容量已满时再追加成员. 原则是重新申请cap*2的空间, 将数据复制过去.
	// 原来的切片对象可能被GC回收, 所以这一步s1的地址可能会发生变化, 但也不一定.
	s1 = append(s1, 10)
	fmt.Printf("s1 len: %d, cap: %d, %+v\n", len(s1), cap(s1), s1) // s1 len: 11, cap: 20, [0 1 2 3 4 5 6 7 8 9 10]
	fmt.Printf("s1 address: %p\n", &s1)                            // s1 address: 0xc000004440

```

## 2. `make([]int, 5, 10)`

```go
	// 创建切片, len=5, cap=10, 底层实际上分配了10个元素大小的空间
	s2 := make([]int, 5, 10)
	// 此时切片不为空, 长度为5, 可以通过s2[n]访问内容, n取0-4
	fmt.Printf("s2 len: %d, cap: %d, %+v\n", len(s2), cap(s2), s2) // s2 len: 5, cap: 10, [0 0 0 0 0]
	fmt.Printf("s2 address: %p\n", &s2)                            // s2 address: 0xc0000044e0
	// 从第5个成员处开始追加
	for i := 0; i < 5; i++ {
		s2 = append(s2, i)
	}

	fmt.Printf("s2 len: %d, cap: %d, %+v\n", len(s2), cap(s2), s2) // s2 len: 10, cap: 10, [0 0 0 0 0 0 1 2 3 4]
	fmt.Printf("s2 address: %p\n", &s2)                            // s2 address: 0xc0000044e0

```

## 3. `make([]int, 10)`不指定cap值

```go
	// 如果没有指定capacity，那么cap与len相等
	s3 := make([]int, 10)
	fmt.Printf("s3 len: %d, cap: %d\n", len(s3), cap(s3)) // s3 len: 10, cap: 10
```

## 4. `[]int{...}`

```go
	// make方式必须指定len参数, 平常经常使用如下方式创建slice, len和cap默认都是0
	// 其实等同于make([]int, 0)
	s4 := []int{}
	fmt.Printf("s4 len: %d, cap: %d\n", len(s4), cap(s4)) // s4 len: 0, cap: 0
	s5 := []int{1, 2, 3, 4, 5}
	fmt.Printf("s5 len: %d, cap: %d\n", len(s5), cap(s5)) // s5 len: 5, cap: 5
```

## 5. 定长数组的len与cap相等, 且不可变

```go
	s6 := [3]int{}
	fmt.Println(len(s6), cap(s6), s6) // 3 3 [0 0 0]
```

## 6. append(slice1, slice2...)时的len与cap

在我以为我已经完全弄懂了golang中切片的len与cap时, 如下一个示例就把我打蒙了.

```go
	s7 := make([]int, 5)
	fmt.Println(len(s7), cap(s7), s7) // 5 5 [0 0 0 0 0]
	s8 := []int{11, 22, 33, 44, 55, 66}
	s7 = append(s7, s8...)
	fmt.Println(len(s7), cap(s7)) // 11 12

```

为什么`s7`的cap不是20???

我做了如下实验, 得到几个不太确定的结论. 之所以说不太确定, 是因为至今还未找到相关的文章印证我的这种猜测.

```go
	s7 := make([]int, 5)
	var sx, result []int
	fmt.Println(len(s7), cap(s7)) // 5 5

	// 通过append(slice1, slice2...)连续追元素造成cap空间的扩展, 最初的确会是slice1的2倍
	sx = []int{1, 2}
	result = append(s7, sx...)
	fmt.Println(len(result), cap(result)) // 7 10

	// 但是当slice2的元素个数与slice1相加后超过了slice1的2倍时, 之后扩展的值就是以2为单位了.
	// 比如如下示例中, 元素6已经超过10, 于是cap空间扩展为12, 长度为11.
	sx = []int{1, 2, 3, 4, 5, 6}
	result = append(s7, sx...)
	fmt.Println(len(result), cap(result)) // 11 12

	// 下面的2个示例同样印证了上述假设.
	sx = []int{1, 2, 3, 4, 5, 6, 7}
	result = append(s7, sx...)
	fmt.Println(len(result), cap(result)) // 12 12

	sx = []int{1, 2, 3, 4, 5, 6, 7, 8}
	result = append(s7, sx...)
	fmt.Println(len(result), cap(result)) // 13 14

	// 当跳出通过append以及...连续追加元素时, 新的cap空间又会成为之前的2倍了.
	sx = []int{1, 2}
	// 注意这里是对result的扩展而不再是s7
	result = append(result, sx...)
	fmt.Println(len(result), cap(result)) // 15 28

```

> 最初以为可能是编译器优化的原因, 但是使用`go run -gcflags='-N -l' .\main.go`执行代码并没有区别, 所以并不是这个原因, 应该就是golang的内部机制.

