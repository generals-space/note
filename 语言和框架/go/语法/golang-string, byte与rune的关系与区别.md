# golang-string, byte与rune的关系与区别

参考文章

1. [string rune byte 的关系](https://www.golangtc.com/t/528cc004320b52227200000f)

2. [golang byte和rune的区别](https://blog.csdn.net/cscrazybing/article/details/79107412)

在Go当中string底层是用`[]byte`存的, 并且是不可以改变的(如果要修改string内容需要将string转换为[]byte或[]rune，并且修改后的string内容是重新分配的).

```go
s := "Go编程"
fmt.Println(len(s))
```

输出结果应该是8, 因为中文字符是用3个字节存的. 

如果想要获得我们想要的情况的话，需要先转换为rune切片再使用内置的len函数. 在使用方法上, `byte`与`rune`没有任何区别.

```go
fmt.Println(len([]rune(s)))
```

结果就是4了。

所以用string存储unicode的话，如果有中文，按下标是访问不到的，因为你只能得到一个byte。 要想访问中文的话，还是要用rune切片，这样就能按下表访问。

参考文章2中介绍了`byte`与`rune`的区别: `byte`是`uint8`的别名, 而`rune`是`uint32`的别名.

所以一般在涉及到包含中文字符串的切片操作时, 尽量先把字符串类型转换成`[]rune`数组来做.