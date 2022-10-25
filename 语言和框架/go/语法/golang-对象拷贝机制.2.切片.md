# golang-对象拷贝机制

本来以为切片用等号赋值是值拷贝, 对新切片的修改不会影响到原来的切片. 因为我的确实验过如下代码

```go
	slice1 := []int{1, 2, 3}
	slice2 := slice1
	slice2 = append(slice2, 4)
	log.Printf("%+v\n", slice1) // [1 2 3]
	log.Printf("%+v\n", slice2) // [1 2 3 4]

	map1 := map[string]interface{}{
		"Name": "general",
		"Age":  21,
	}
	map2 := map1
	map2["Name"] = "longbei"

	log.Printf("%+v\n", map1) // map[Name:longbei Age:21]
	log.Printf("%+v\n", map2) // map[Name:longbei Age:21]

```

但是没想到修改和新增操作对原切片的影响是不同的...

```go
	var slice1 []int

	slice1 = []int{1, 2, 3}

	slice2 := slice1
	slice2 = append(slice2, 4)
	// 这里可以看到slice2新增成员对slice1是没有影响的
	log.Printf("%+v\n", slice1) // [1 2 3]
	log.Printf("%+v\n", slice2) // [1 2 3 4]

	slice3 := slice1
	slice3[0] = 0
	// 但是修改slice3中的成员却是会影响到slice1的.
	log.Printf("%+v\n", slice1) // [0 2 3]
	log.Printf("%+v\n", slice3) // [0 2 3]

	log.Println("========================")
	// 以下是对通过[:]赋值的操作的检验, 没有什么区别
	slice1 = []int{1, 2, 3}

	slice4 := slice1[:]
	slice4 = append(slice4, 4)
	log.Printf("%+v\n", slice1) // [1 2 3]
	log.Printf("%+v\n", slice2) // [1 2 3 4]

	slice5 := slice1[:]
	slice5[0] = 0
	// 但是修改slice5中的成员却是会影响到slice1的.
	log.Printf("%+v\n", slice1) // [0 2 3]
	log.Printf("%+v\n", slice5) // [0 2 3]
```


对于这种差异情况, 我的理解是: 通过等号赋值的操作的确是值拷贝, 但是要记得切片在底层存储的内存布局. 

切片其实是一种对象, 包含`pointer`, `len`和`cap`3个成员, 其中`pointer`指向一段连续的空间, 就是C语言中常规的数组.

通过等号赋值, 复制得到新的切片对象, 于是`pointer`成员的值也是相同的, 所以指向同一块数组空间.

当对slice2进行append操作时, 会造成数组空间的扩容, golang会重新开辟一段大小是原来2倍的空间, 并将原来的成员拷贝过去, 这样的结果就是...slice2中`pointer`所指向的数组空间地址变掉了, 于是不会影响到slice1原来的值.

而slice3通过索引修改切片成员的值则不会有这样的后果, 因而同时影响到了slice1的成员.

------

我重新尝试了下map和channel两种类型, 并没有这种问题. 因为map和channel的底层都指向堆中的地址, ta们是特殊的数据结构, 所以map/channel变量其实可以看作是指针.
