# golang-切片合并

<!key!>: {0a187fb9-dc95-410a-b32a-9fc8108655c1}

参考文章

1. [golang如何优雅的合并切片？](https://segmentfault.com/q/1010000011354818)

## 1. 合并

php里面有`array_meage()`, js有`concat()`, go的切片操作中是否有类似函数呢? 还是只能手动遍历然后一个一个的append?

`append()`函数只能将单个元素添加至目标切片末尾? 不是的.

```go
arr3 := append(arr1, arr2...)
```

`arr1`和`arr2`都是切片类型, 尤其注意最后那三个点`...`, 三个点是解构的意思.

如果想把一个切片(或是一个指定元素)放到另一个切片中间(而不是末尾), 要怎么做?

我本来以为可以用下面的代码解决

```go
	a := []int{2, 3, 6, 9, 1, 5,}
	b := 0
	leftList := a[:3]
	rightList := a[3:]
	// append 虽然能使用 ... 对切片解构, 但是不能显式使用3个参数, 只能分开设置
	// too many arguments to append
	// newList := append(leftList, b, rightList...)
	newList := append(leftList, b)
	newList = append(newList, rightList...)
	fmt.Printf("%+v\n", newList)	// [2 3 6 0 0 1 5]
```

但是这样的结果是错的, 因为我忘记了slice本质就是 **数组 + 游标**.

在`append(leftList, b)`时, 实际上也修改了原切片`a`的第3个成员9, 使之变成了0. 不过这样的替换没有影响到`a`后面的成员, 此时`a`的值为`[2 3 6 0 1 5]`.

```go
	a := []int{2, 3, 6, 9, 1, 5,}
	b := 0
	leftList := a[:3]
	rightList := a[3:]
	newList := append(leftList, b) 	// [2 3 6 0]
	fmt.Printf("%+v\n", a)			// [2 3 6 0 1 5]
	fmt.Printf("%+v\n", rightList)	// [0 1 5]

```

如果`b`也是一个切片, 且比`rightList`长会怎么样呢? 会把`a`的后半部分全部替换掉吗? 

我以为会的, 但我又错了.

```go
	a := []int{2, 3, 6, 9, 1, 5}
	b := []int{7,7,7,7,7}
	leftList := a[:3]
	rightList := a[3:]
	newList := append(leftList, b...) 	// [2 3 6 7 7 7 7 7]
	fmt.Printf("a %+v\n", a)			// [2 3 6 9 1 5]
```

竟然对`a`本身没有任何影响...这应该涉及到golang底层了, 应该是`append()`对单个变量与切片的处理方式不同.

至少目前应该确认, **在使用切片的合并操作时, 要保持都使用切片, 单个成员也要构造成切片再操作, 这样可以避免出错**.

我以为这就是终点了, 结果我又双叒叕错了. ๐·°(৹˃̵﹏˂̵৹)°·๐

```go
	a := []int{2, 3, 6, 9, 1, 5}
	b := []int{7,7,7,7,7}
	leftList := a[:0]
	rightList := a[0:]
	newList := append(leftList, b...)
	// fmt.Printf("a %+v\n", a)					// [7 7 7 7 7 5]
	// fmt.Printf("rightList %+v\n", rightList)	// [7 7 7 7 7 5]
	// fmt.Printf("newList %+v\n", newList)		// [7 7 7 7 7]
	newList = append(newList, rightList...)
	// fmt.Printf("newList %+v\n", newList)		// [7 7 7 7 7 7 7 7 7 7 5]
```

当切片游标为0时(leftList实际上为一个空数组`[]`), 结果一切都不一样了. 

在向`leftList`中追加`b`切片时, 修改了原本的`a`, 也许是因为`a[0]`的地址与`a`底层的数组地址相同, 总之这仍然是底层实现的问题. 

如果不考虑这些东西, 应该怎么解决?

那就是将`newList`声明为一个空切片, 按顺序追加`leftList`, `b`和`rightList`数组.

```go
	a := []int{2, 3, 6, 9, 1, 5}
	b := []int{7,7,7,7,7}
	leftList := a[:0]
	rightList := a[0:]
	newList := []int{}
	newList = append(newList, leftList...)
	newList = append(newList, b...)
	// fmt.Printf("a %+v\n", a)					// [2 3 6 9 1 5]
	// fmt.Printf("rightList %+v\n", rightList)	// [2 3 6 9 1 5]
	// fmt.Printf("newList %+v\n", newList)		// [7 7 7 7 7]
	newList = append(newList, rightList...)
	// fmt.Printf("newList %+v\n", newList)		// [7 7 7 7 7 2 3 6 9 1 5]
```

唉, 给我这顿折腾.

## 2. 拷贝

内置函数`copy(dst, src)`, 参数都必需是切片类型, 而不能是数组. 将第二个slice里的元素拷贝到第一个slice里, 拷贝的长度为两个slice中长度较小的长度值.

```go
	old := []byte{1, 2, 3, 4}
	new := [2]byte{}
	copy(new[:], old)
	log.Println(new)	// [1 2]
```
