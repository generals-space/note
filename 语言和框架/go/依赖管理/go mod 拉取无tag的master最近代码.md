# go mod 拉取无tag的master最近代码

```
require (
    github.com/generals-space/crd-ipkeeper
)
```

```
$ go mod tidy
go: errors parsing go.mod:
/home/cni-terway/go.mod:9: usage: require module/path v1.2.3
```

出现这个问题是因为自己的仓库没写版本号, 给仓库加个tag就可以了.

但是在实验过程中的仓库不好加版本号, 一直推在master分支上, 此时不好直接在go.mod中写版本号, 需要在项目目录下使用go get命令先下载, 完成后会自动修改go.mod文件

```
$ go get github.com/generals-space/crd-ipkeeper@master
go: finding github.com/generals-space/crd-ipkeeper master
go: downloading github.com/generals-space/crd-ipkeeper v0.0.0-20200321193328-62c4ac27beb2
go: extracting github.com/generals-space/crd-ipkeeper v0.0.0-20200321193328-62c4ac27beb2
```

> 呃, 已经存在go.mod中的`v0.0.0-20200321193328-62c4ac27beb2`记录, 在使用`go get`时好像不会更新, 还是把这种记录删除后再`go get`吧.

