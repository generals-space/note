# golang-文件读写(二)-os

参考文章

1. [Go实战--golang中读写文件的几种方式](https://blog.csdn.net/wangshubo1989/article/details/74777112)

2. [Golang读写文件操作](https://my.oschina.net/xxbAndy/blog/1594259)

用os完成读写就比较复杂了, 要先打开文件(还要注意打开的模式), 读取的字节数, 偏移量等...

首先要注意的就是两个打开文件的方法.

## 1. 打开文件

1. `Open func(name string) (*File, error)`
2. `OpenFile func(name string, flag int, perm FileMode) (*File, error)`
3. `Create func(name string) (*File, error)`

`os.Open()`调用了`OpenFile`...如下

```go
func Open(name string) (*File, error) {
	return OpenFile(name, O_RDONLY, 0)
}
```

所以两者是同一个东西, `Open`是只读的.
