# golang-init与main函数的区别

参考文章

1. [main() ,init()方法是go中默认的两个方法, 两个保留的关键字.](https://studygolang.com/articles/5036)

## 1. 引言

1. `main()`方法只能用在`package main`中; `init()`方法是在任何`package`中都可以出现, 但是建议每个`package`中只包含一个`init()`函数比较好, 容易理解. 

2. Go程序执行时会自动调用`init()`和`main()`, 所以你不需要在任何地方调用这两个函数. 每个`package`中的`init()`函数都是可选的, 但`package main`就**必须**包含一个`main()`函数. 

**go程序的执行流程**

程序的初始化和执行都起始于`main`包. 如果`main`包还导入了其它的包, 那么就会在编译时将它们依次导入. 有时一个包会被多个包同时导入, 那么它只会被导入一次（例如很多包可能都会用到fmt包, 但它只会被导入一次, 因为没有必要导入多次）. 

当一个包被导入时, 如果该包还导入了其它的包, 那么会先将其它包导入进来, 然后再对这些包中的 **包级常量和变量** 进行初始化, 接着执行`init()`函数（如果有的话）, 依次类推. 等所有被导入的包都加载完毕了, 就会开始对`main`包中的包级常量和变量进行初始化, 然后执行`main`包中的`init()`函数（如果存在的话）, 最后执行`main()`函数

## 2. 简单示例

`test.go`文件

```go
package main
import "log"

func init(){
    log.Println("in init...")
}
func main(){
    log.Println("in main...")
}
```

执行结果如下

```
$ go run test.go 
2018/05/15 22:08:16 in init...
2018/05/15 22:08:16 in main...
```

可以看到, 同一个包中的`init()`与`main()`, `init()`优先执行.

## 3. 复杂示例

编写如下结构的代码

1. `$GOPATH/mypkg/main.go`

2. `test.go`

`$GOPATH/mypkg/main.go`文件

```go
package mypkg
import "log"

var Mytxt string
func init(){
    Mytxt = "hello world"
    log.Println("init in mypkg")
}

func main(){
    log.Println("main in mypkg")
}
```

`test.go`文件

```go
package main

import "log"
import "mypkg"
func init(){
    log.Println("init in main...")
    log.Println(mypkg.Mytxt)    
}
func main(){
    log.Println("main in main...")
    log.Println(mypkg.Mytxt)
}
```

执行结果如下

```
2018/05/15 22:32:16 init in mypkg
2018/05/15 22:32:16 init in main...
2018/05/15 22:32:16 hello world
2018/05/15 22:32:16 main in main...
2018/05/15 22:32:16 hello world
```

所以一般`init()`是作为包级变量和常量初始化(但不方便暴露给外部)的方法使用的.