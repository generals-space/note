参考文章

1. [【golang】unsafe.Sizeof浅析](https://blog.csdn.net/HaoDaWang/article/details/80005072)
2. [【golang】unsafe.Sizeof浅析](https://www.cnblogs.com/lovezbs/p/13127478.html)
    - 参考文章1的转换文章

unsafe.Sizeof() 取的是类型的大小, 而非实例的大小.

```go
str := "hello"
fmt.Println(unsafe.Sizeof(str)) //16
```

不论字符串的len有多大, sizeof始终返回16.

实际上字符串类型对应一个结构体, 该结构体有两个域, 第一个域是指向该字符串的指针, 第二个域是字符串的长度, 每个域占8个字节, 但是并不包含指针指向的字符串的内容, 这也就是为什么 Sizeof() 始终返回的是16

struct{} 类型情况也差不多, 空结构体占用为0. 增加成员字段, Sizeof() 结果也会相应增大.
