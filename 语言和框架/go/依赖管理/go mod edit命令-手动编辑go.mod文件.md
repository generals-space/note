# go mod edit命令-手动编辑go.mod文件

golang: 1.13.14.

有时在go mod 工程目录下执行"go get -v path@version", 并不能真的将该依赖库写入 go.mod 文件(有的时候写入了, 但是vscode一保存, 就又变回来了, 手动修改 go.mod 文件也是这样), 尤其是对一个迭代过多次的 GOPATH 工程, 最容易出现这样的问题, 搞得不知道怎么让工程使用我们指定版本的依赖库...

这种情况可以使用"go mod edit"命令, 如下

```
go mod edit -require=k8s.io/client-go@v0.17.2 go.mod
```

直接写入 go.mod 文件, 不再变回去了.

可以写多个.

```
go mod edit -require=k8s.io/api@v0.17.2 -require=k8s.io/apimachinery@v0.17.2 go.mod
```


```
go mod edit -replace=k8s.io/kubernetes@v0.17.2=./vendor/k8s.io/kubernetes
```

> 需要事先创建"./vendor/k8s.io/kubernetes"本地目录.

