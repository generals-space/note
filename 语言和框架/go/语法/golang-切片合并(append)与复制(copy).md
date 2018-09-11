# golang-切片合并

参考文章

1. [golang如何优雅的合并切片？](https://segmentfault.com/q/1010000011354818)

## 1. 合并

php里面有`array_meage()`, js有`concat()`, go的切片操作中是否有类似函数呢? 还是只能手动遍历然后一个一个的append?

`append()`函数只能将单个元素添加至目标切片末尾? 不是的.

```go
arr3 := append(arr1, arr2...)
```

`arr1`和`arr2`都是切片类型, 尤其注意最后那三个点`...`, 三个点是解构的意思.

## 2. 拷贝

内置函数`copy(dst, src)`, 参数都必需是切片类型, 而不能是数组. 将第二个slice里的元素拷贝到第一个slice里, 拷贝的长度为两个slice中长度较小的长度值.

```go
	old := []byte{1, 2, 3, 4}
	new := [2]byte{}
	copy(new[:], old)
	log.Println(new)	// [1 2]
```
