## 1.

```
$ go mod download
warning: pattern "all" matched no module dependencies
```

问题分析: 也许当前工程目录放在了GOPATH路径中, 而go mod工程必须放在GOPATH路径之外.

## 2.

```
~/Code/playground/go-mod $ go mod init
go: cannot determine module path for source directory /Users/general/Code/playground/go-mod (outside GOPATH, no import comments)
```

问题分析: `go mod init`后需要一个参数指定包名, 比如`github.com/generals-space/test`

## 

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