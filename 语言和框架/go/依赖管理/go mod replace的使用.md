# go mod replace的使用

参考文章

1. [从gopath到go mod的一次尝试](https://blog.csdn.net/qq_33296108/article/details/88184060)
2. [justwatchcom/elasticsearch_exporter](https://github.com/justwatchcom/elasticsearch_exporter/tree/v1.1.0rc1)

最近在看 elasticsearch 的 exporter 源码, 见参考文章2.

这个工程下只有 vendor, 没有 go.mod, 说明是没有使用 go.mod 管理依赖的, 最好将`GO111MODULE`关掉.

不过我最初没有注意到, 直接执行 build 构建, 结果出现如下报错

```console
$ GOOS=linux GOARCH=amd64 go build -o elasticsearch_exporter ./*.go
build command-line-arguments: cannot load gopkg.in/alecthomas/kingpin.v2: cannot find module providing package gopkg.in/alecthomas/kingpin.v2
```

我换了几个 goproxy 镜像源地址, 也试过关掉 cgo, 但是都不行.

后来找到参考文章1, 提到可以使用`replace`, 我访问了一下`gopkg.in/alecthomas/kingpin.v2`, 是直接跳转到`github.com/alecthomas/kingpin`的, 于是在`go.mod`下添加如下语句

```go mod
replace gopkg.in/alecthomas/kingpin.v2 => github.com/alecthomas/kingpin v0.0.0-20171217180838-947dcec5ba9c
```

重新编译就可以了.

由此认识到`replace`标识的使用场景.

------

重新测试, 发现不用修改任何东西, `vendor/vendor.json`会自动转换成`go.mod`...

```
$ GOOS=linux GOARCH=amd64 go build -o elasticsearch_exporter ./*.go
go: creating new go.mod: module github.com/justwatchcom/elasticsearch_exporter
go: copying requirements from vendor/vendor.json
go: updates to go.mod needed, but contents have changed
```

md, 那我纠结了半天搞了点啥...
