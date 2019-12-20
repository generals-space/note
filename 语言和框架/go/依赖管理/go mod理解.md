# go mod理解

参考文章

1. [跳出Go module的泥潭](https://colobu.com/2018/08/27/learn-go-module/)

2. [Golang官方包依赖管理工具 go mod 简明教程](https://ieevee.com/tech/2018/08/28/go-modules.html)
    - vendor, node_modules, 和maven的对比: go mod实现的是类似于maven的中心缓存, 而不是像node那样的局部缓存


`go mod download`可以下载所需要的依赖，但是依赖并不是下载到`$GOPATH`中，而是`$GOPATH/pkg/mod`中，多个项目可以共享缓存的module. 同时改写go.mod文件, 添加上下载的pkg信息.

`go mod vendor` 会复制modules下载到vendor中, 貌似只会下载你代码中引用的库，而不是go.mod中定义全部的module. 

`go get|test|list|build`都会修改`go.mod`文件

------

go1.11后开始, 貌似go module模式下, go get的行为也会像go mod download那样了.

关于`indirect`标记

```
require k8s.io/client-go v0.0.0-20190918160344-1fbdaa4c8d90 // indirect
```

有可能是在工程目录下使用`go get`或`go mod download`手动下载了不相关的包, 并不是工程本身需要的, 这种情况下, 可能需要手动删除此行记录.
