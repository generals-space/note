# go mod GO111MODULE变量理解

参考文章

1. [离线环境中，go mod一直下载对应的依赖的解决办法](https://blog.csdn.net/qq_24271853/article/details/105820048)
    - 离线环境的 GO111MODULE 变量设置
2. [使用Go env命令设置Go的环境](https://www.cnblogs.com/fanbi/p/13649929.html)
    - go 1.13 时可使用 `go env -w`

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

当开启了`GO111MODULE`时, 在一个 golang 目录下执行`go env`, 将会看到`GOMOD=当前路径/go.mod`. 当然, 由于一个 golang 工程可能没有`go.mod`文件, 所以 vscode 可能会有提示"You are neither in a module nor in your GOPATH".

## go env -w

从 go 1.13 开始, 修改`GO111MODULE`就可以不必在"我的电脑"(win)或是".bashrc"中设置了, 可以像`sysctl -w`一样, 修改这个环境变量(全局).

但是, 通过`go env -w GO111MODULE`修改的变量并不会出现在"我的电脑"中, 如果"我的电脑"中已经存在该环境变量, 那么再使用`go env -w`修改时会提示如下错误.

```
$ go env -w GO111MODULE=off
warning: go env -w GO111MODULE=... does not override conflicting OS environment variable
```

## 关于 go.mod 中 module 名称

我一直以为 go.mod 中的`module`字段名称与工程名称是随意的...

以`kubeedge`为例, 其`module`名称为`github.com/kubeedge/kubeedge`, 那么该工程就需要放在`xxx/github.com/kubeedge/kubeedge`目录, 不管前缀是什么, 总之后面的部分要和`module`保持一致. 

这样, 在工程的根目录执行`go env`时, `GOMOD`的值就会是`xxx/github.com/kubeedge/kubeedge/go.mod`, 在开启`GO111MODULE`时, 才会将当前工程的名称当作`github.com/kubeedge/kubeedge`.

否则, 在`go run`时, 虽然能发现`go.mod`文件, 但是 go 不会把当前工程当作`github.com/kubeedge/kubeedge`, 源码中如果有引用本工程内部的文件时, 就会失败. 同时, 会重新解析`go.mod`文件, 如果在离线环境中, 无法联网, 对于`go.mod`中的各个包都会因无法解析域名而失败...

...那是不是只有在离线环境下才必须把工程放在于`module`同名的目录下? 如果有网的话, 不管放在哪个目录, 都可解析`go.mod`中的包?

