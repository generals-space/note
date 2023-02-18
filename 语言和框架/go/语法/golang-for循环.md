# golang-for循环

参考文章

1. [golang：for循环](https://www.jianshu.com/p/1de5050cc88c)
2. [【GoLang】GoLang for 中有多个循环变量怎么处理？](https://www.cnblogs.com/junneyang/p/6072680.html)

最常用的是`for..range..`循环, 类似于其他语言中的`foreach`循环.

```go
for i in range slice1 {

}
```

然后是条件循环, 使用方法与其他语言中几乎相同.

```go
for i := 0; i < 10; i++ {

}
```

再然后是只有一个限制条件, 无初始化, 类似于其他语言中的`while..do..`循环, 两种表示方式.

```go
for i < 10 {
    // statement
}
for ; i < 10; {
    // statement
}
```

最后是无限循环, 需要手动使用`break`跳出.

```go
for {

}
```

------

多条件for循环的使用方法.

```go
for i, j := 0, 0; i <= 5 && j <= 5; i, j = i+1, j-1 {

}
```

下面是错误的

```
for i, j := 1, 10; i < j; i++, j++ {
    fmt.Println(i)
}
```
