# golang-内置函数copy切片深拷贝

参考文章

1. [golang 切片copy复制和等号复制的区别](https://blog.csdn.net/qq_36520153/article/details/85788763)
    - 等号复制与copy复制在效率上的差异
    - 等号浅拷贝与copy深拷贝特性

`copy(dst, src)`

其中`dst`和`src`必须是切片类型, 不可以是定长数组.

## 1. 使用方法

下面看一下`copy`的用法

```go
	slice1 := []int{1, 2, 3}

	slice2 := []int{}
    copy(slice2, slice1)
    // slice2没有内容, 是因为没有分配内存空间
	log.Println(slice2) // []

	slice3 := make([]int, 5)
    copy(slice3, slice1)
    // 
	log.Println(slice3) // [1 2 3 0 0]

	slice4 := [5]int{}
    copy(slice4[2:], slice1)
    // 本次结果可以证明, copy可以通过指定切片索引决定内容拷贝到哪里
	log.Println(slice4) // [0 0 1 2 3]
```

上述代码可以得出结论: copy需要在拷贝前为目标切片申请空间.

## 2. 等号浅拷贝 vs copy深拷贝

```go
	slice1 := []int{1, 2, 3}

	slice2 := make([]int, 3)
	copy(slice2, slice1)
	log.Println(slice2) // [1 2 3]

	slice2[0] = 0
	log.Printf("%p %p %+v", &slice1, &slice1[0], slice1) // 0xc000004080 0xc00000e300 [1 2 3]
	log.Printf("%p %p %+v", &slice2, &slice2[0], slice2) // 0xc0000040a0 0xc00000e320 [0 2 3]

	// 此时slice1仍未发生变化
	slice3 := slice1
	slice3[0] = 0
	log.Printf("%p %p %+v", &slice1, &slice1[0], slice1) // 0xc000004080 0xc00000e300 [0 2 3]
	log.Printf("%p %p %+v", &slice3, &slice3[0], slice3) // 0xc0000042a0 0xc00000e300 [0 2 3]
```

copy得到的新切片更改后, 不会影响原切片, 可以猜测copy不只拷贝了切片对象的`pointer`, `len`, `cap`3个成员, 还拷贝了底层的数组空间...

> 这里要注意一点, `&slice`得到的只是切片对象本身地址, `&slice[0]`才是`pointer`指向的底层数组空间的起始地址.

