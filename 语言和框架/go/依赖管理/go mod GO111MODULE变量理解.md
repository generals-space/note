# go mod GO111MODULE变量理解

术语: module模式与gopath模式

GO111MODULE:

- `off`: 无模块支持, go 会从 GOPATH 和 vendor 目录寻找包. 
- `on`: 模块支持, go 会忽略 GOPATH 和 vendor 目录, 只根据 go.mod 下载依赖. 
- `auto`: 在 `$GOPATH/src`外且根目录有 go.mod 文件时, 开启模块支持. 

对于`auto`模式有些疑惑, 在`$GOPATH/src`外且根目录有`go.mod`文件时, 才开启module模式, 有一项不满足, 都会转换为gopath模式. 

gopath模式无法使用`go get|install path@version`语法.

> 无论module还是gopath模式, go install安装的二进制文件的目标目录都是`$GOPATH/bin`.

如果在`auto`模式出现如下问题, 最好考虑一下当前目录是否在`GOPATH`外, 所在工程是否存在`go.mod`文件, 如果是, 需要离开这个目录.

```
$ go get github.com/gorilla/mux@latest
go: cannot use path@version syntax in GOPATH mode
```

但是`go install`无法使用`@version`语法, gopath模式下的`go install`安装的是src目录的工程, 但modules模式下可能有多个版本的工程, 默认应该是安装当前pkg/mod目录下最新的工程.
