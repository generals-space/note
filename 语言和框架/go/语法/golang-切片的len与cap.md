# golang-切片的len与cap

参考文章

1. [从golang的垃圾回收说起（下篇） 5.3节](https://sq.163yun.com/blog/article/192800393799778304)

```go
package main

import "fmt"

func main() {
	// 创建切片, len=0, cap=10, 底层实际上分配了10个元素大小的空间
	slice1 := make([]int, 0, 10)
	// 此时切片为空, 长度为0, 无法通过slice1[0]访问内容, 会越界
	fmt.Printf("slice1: %+v\n", slice1)         // slice1: []
	fmt.Printf("slice1 len: %d\n", len(slice1)) // slice1 len: 0
	fmt.Printf("slice1 cap: %d\n", cap(slice1)) // slice1 cap: 10
	fmt.Printf("slice1 address: %p\n", &slice1) // slice1 address: 0xc000004440
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Printf("slice1: %+v\n", slice1)         // slice1: [0 1 2 3 4 5 6 7 8 9]
	fmt.Printf("slice1 len: %d\n", len(slice1)) // slice1 len: 10
	fmt.Printf("slice1 cap: %d\n", cap(slice1)) // slice1 cap: 10
	fmt.Printf("slice1 address: %p\n", &slice1) // slice1 address: 0xc000004440

	// 当cap容量已满时再追加成员
	// 原则是重新申请cap*2的空间, 将数据复制过去.
	// 原来的切片对象可能被GC回收
	// 所以这一步slice1的地址可能会发生变化, 但也不一定.
	slice1 = append(slice1, 10)
	fmt.Printf("slice1: %+v\n", slice1)         // slice1: [0 1 2 3 4 5 6 7 8 9 10]
	fmt.Printf("slice1 len: %d\n", len(slice1)) // slice1 len: 11
	fmt.Printf("slice1 cap: %d\n", cap(slice1)) // slice1 cap: 20
	fmt.Printf("slice1 address: %p\n", &slice1) // slice1 address: 0xc000004440

	fmt.Println("///////////////////////////////////////////")
	// 创建切片, len=5, cap=10, 底层实际上分配了10个元素大小的空间
	slice2 := make([]int, 5, 10)
	// 此时切片不为空, 长度为5, 可以通过slice2[n]访问内容, n取0-4
	fmt.Printf("slice2: %+v\n", slice2)         // slice2: [0 0 0 0 0]
	fmt.Printf("slice2 len: %d\n", len(slice2)) // slice2 len: 5
	fmt.Printf("slice2 cap: %d\n", cap(slice2)) // slice2 cap: 10
	fmt.Printf("slice2 address: %p\n", &slice2) // slice2 address: 0xc0000044e0
	// 从第5个成员处开始追加
	for i := 0; i < 5; i++ {
		slice2 = append(slice2, i)
	}

	fmt.Printf("slice2: %+v\n", slice2)         // slice2: [0 0 0 0 0 0 1 2 3 4]
	fmt.Printf("slice2 len: %d\n", len(slice2)) // slice2 len: 10
	fmt.Printf("slice2 cap: %d\n", cap(slice2)) // slice2 cap: 10
	fmt.Printf("slice2 address: %p\n", &slice2) // slice2 address: 0xc0000044e0

	fmt.Println("///////////////////////////////////////////")

	// 如果没有指定capacity，那么cap与len相等
	slice3 := make([]int, 10)
	fmt.Printf("slice3 len: %d\n", len(slice3)) // slice3 len: 10
	fmt.Printf("slice3 cap: %d\n", cap(slice3)) // slice3 cap: 10
	// make方式必须指定len参数, 平常经常使用如下方式创建slice, len和cap默认都是0
	slice4 := []int{}
	fmt.Printf("slice4 len: %d\n", len(slice4)) // slice4 len: 0
	fmt.Printf("slice4 cap: %d\n", cap(slice4)) // slice4 cap: 0
}
```