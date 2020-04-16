# golang-os文件操作

参考文章

1. [GO语言常用的文件读取方式](https://www.jb51.net/article/58147.htm)

2. [Golang读写文件的几种方式](https://www.jianshu.com/p/7790ca1bc8f6)

参考文章1中给出了3种读取文件的方式: 

1. (小文件)一次性读取
2. (大文件)分块读取
3. (文本文件)逐行读取

参考文章2中则是给出了`os`和`bufio`结合使用的方式...目前暂不深究.

> `加密与认证机制`分类下的[计算大文件md5及sha1值(golang版)]()在给出计算文件md5值的示例的同时也是读取文件(一次性和分块读取)的示例.

我们知道C和python中, 打开文件的内置函数都是叫`open()`, 可以传入路径和打开模式(只读, 读写等)

在golang中不是这样...

golang的os包中有`Open()`和`OpenFile()`两个函数, `Open()`的函数定义如下

```go
func Open(name string) (*File, error) {
	return OpenFile(name, O_RDONLY, 0)
}
```

就是只读模式打开, 而能够使用各种Flag等参数的却是`OpenFile()`函数.
