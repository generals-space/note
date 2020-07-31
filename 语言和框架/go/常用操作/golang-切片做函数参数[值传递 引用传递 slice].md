# golang-切片做函数参数[值传递 引用传递 slice]

参考文章

1. [go语言切片作为函数参数](https://www.cnblogs.com/MyUniverse/p/12653686.html)
    - 切片作为函数参数时是值传递, 通过下标修改成员时是修改底层的数组, 但是增加和删除时会将指针指向其他地方.
2. [go语言切片作为函数参数的研究](https://www.cnblogs.com/endurance9/p/10347336.html)
    - 切片对象取指针作为函数参数, 可以在被调函数中对切片成员进行增删, 值得一看.

go 1.12.15

我最初就是希望将一个切片对象传入一个函数, 在此函数中向这个切片中追加如下, 最终主调函数中可以得到拥有成员的切片. 但我似乎想错了.

## 1. 

```go
func makeSlice(slice []int) {
	slice = append(slice, 1)
	fmt.Printf("%+v\n", slice)	// [1]
}

func main() {
	slice := make([]int, 0)
	makeSlice(slice)
	fmt.Printf("%+v\n", slice)	// []
}

```

## 2. 

作为函数参数是值拷贝, 在被调函数中通过下标修改切片成员, 是通过slice中保存的地址对底层数组进行修改.

作为函数参数, 当在函数中使用append增加切片元素的时候, 就相当于创建一个新的变量(追加是在容量 cap 之后的部分追加).

```go
func makeSlice(slice []int) {
	slice[0] = 1
	slice = append(slice, 1)
	fmt.Printf("%+v\n", slice) // [1 1]
}

func main() {
	slice := make([]int, 1)
	slice[0] = 0
	fmt.Printf("%+v\n", slice) // [0]
	makeSlice(slice)
	fmt.Printf("%+v\n", slice) // [1] 对下标的修改成功了, 但是 append 操作的结果仍然没有反映到主调函数.
}

```

## 3. 

之后猜测是不是需要把容量加够, 之后的`append`就不会修改原来的切片地址了(毕竟超出一次容量, 就会引起重新申请空间).

但是我又错了.

```go
func makeSlice(slice []int) {
	slice[0] = 1
	slice = append(slice, 1)
	fmt.Printf("%+v\n", slice) // [1 0 0 0 0 1]
}

func main() {
	slice := make([]int, 5)
	slice[0] = 9
	fmt.Printf("%+v\n", slice) // [9 0 0 0 0]
	makeSlice(slice)
	fmt.Printf("%+v\n", slice) // [1 0 0 0 0]
}

```

...得, `append`是在 cap 容量之外追加的(实际是 len 之外, 上面`make`的第2个参数设为5, 那么 len 和 cap 都会是5, 而且会把 slice 填满.)

## 4.

```go
func makeSlice(slice *[]int) {
	*slice = append(*slice, 1)
	fmt.Printf("%+v\n", slice) // &[1]
}

func main() {
	slice := make([]int, 0)
	makeSlice(&slice)
	fmt.Printf("%+v\n", slice) // [1]
}

```
