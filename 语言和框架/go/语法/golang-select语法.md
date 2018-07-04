# golang-select语法

参考文章

1. [Golang select的使用及典型用法](https://blog.csdn.net/zhaominpro/article/details/77570290)

2. [Go 语言 select 语句](https://wizardforcel.gitbooks.io/w3school-go/content/13.html)

## 1. 基本用法 

`select`是Go中的一个控制结构，类似于`switch`语句，用于处理异步IO操作。`select`会监听`case`语句中`channel`的读写操作，当`case`中`channel`读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。

> `select`中的`case`语句必须是一个`channel`操作

> `select`中的`default`子句总是可运行的

如果有多个`case`都可以运行，select会**随机公平**地选出一个执行，其他不会执行。

如果没有可运行的`case`语句，且有`default`语句，那么就会执行`default`的动作。

如果没有可运行的`case`语句，且没有`default`语句，`select`将阻塞，直到某个`case`通信可以运行