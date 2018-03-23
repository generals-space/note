# GO环境编译安装

参考文章

1. [Golang新起航！（编译安装go）](https://studygolang.com/articles/9999)

源码地址(不必下载github源码版本)

1. [官方下载地址](https://golang.org/dl/)

2. [中文镜像地址](https://dl.gocn.io/)

## 1. GO语言环境搭建

按照[官方文档](https://golang.org/doc/install/source#go14)的说法, 编译1.5版本之后的GO环境需要使用GO1.4事先编译GO的工具链. 因为高版本的GO编译器是由GO本身写的, GO1.4当初还是只是用C编译的, 可以先通过GO1.4编译得到初始的GO编译器, 然后再编译高版本的GO.

### 1.1 编译go1.4.3

下载GO1.4源码, 解压. 其实解压开的目录名称为`go`, 这里我们将其重命名成`go1.4`.

```
$ cd $GO1.4/src
$ ./make.bash
---
Installed Go for linux/amd64 in /root/go1.4
Installed commands in /root/go1.4/bin
```

GO1.4编译完成后会在其根目录下出现`bin`子目录, 其中有如下可执行文件生成

```
$ pwd
/root/go1.4/bin
$ ls
go  gofmt
```

然后定义环境变量`GOROOT_BOOTSTRAP`, 设置为`go1.4`的绝对路径, 其默认值为`$HOME/go1.4`.

```
$ echo 'export GOROOT_BOOTSTRAP=/root/go1.4' >> /etc/profile
```

### 1.2 编译go1.8.3

下载高版本的go源码, 这里我们选择1.8.3. 解压后目录名称还是`go`, 所以上一步重命名是很有必要的, 否则会被覆盖掉.

```
$ cd $GO1.8.3/src
$ ./all.bash
##### Building Go bootstrap tool.
cmd/dist

##### Building Go toolchain using /root/go1.4.
...
---
Installed Go for linux/amd64 in /root/go
Installed commands in /root/go/bin
*** You need to add /root/go/bin to your PATH.
```

编译完成后, `$GO1.8.3/bin`目录下会出现几个可执行文件

```
$ pwd
/root/go/bin
$ ll
total 28104
-rwxr-xr-x. 1 root root 10068959 Jun  8 18:38 go
-rwxr-xr-x. 1 root root 15226597 May 25 02:16 godoc
-rwxr-xr-x. 1 root root  3477458 Jun  8 18:38 gofmt
```

我们需要设置3个环境变量.

- GOROOT: 用于放置go的标准库和工具链, 使用yum安装时这个值默认为`/usr/lib/golang`

- PATH: 可执行程序`go`与`gofmt`等路径.

- GOPATH: 自定义工程路径, 自定义的工程都放在这个目录下.

把刚才编译的go1.8.3源码码包拷贝到`/usr/local/go`.

`/etc/profile`添加了这几行

```bash
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/root/go_pro
```

使用`go env`, 可以查看与go相关的环境变量.

```
$ go env
GOARCH="amd64"
GOBIN=""
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH=""
GORACE=""
GOROOT="/usr/local/go"
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GO15VENDOREXPERIMENT="1"
CC="gcc"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0"
CXX="g++"
CGO_ENABLED="1"
```

最后, go1.4不想要就删掉吧.

## 2. go工程编译

### 2.1 hello world

先来一行'hello world'压压惊...

在`/root/GO_PRO`目录下创建`gtest.go`文件, 写入如下代码.

```go
package main
import "fmt"
func main(){
    fmt.PrintLn("hello world")
}
```

编译执行

```
$ pwd
/root/go_pro
$ go run gtest.go 
hello world
```

成功

### 2.2 简单工程编译

上面一节是简单的直接运行, 但有很多情况是我们需要编译生成可执行文件, 并且工程规模比较大, 目录比较复杂. 我们需要了解go工程的基本目录结构, 以及第三方库的添加方法. 

```makefile
GOPATH := ${PWD}
PKG_CONFIG_PATH :=/usr/local/lib/pkgconfig
export GOPATH
export PKG_CONFIG_PATH


## default: build

build:
	go build -v -x -o ./bin/ew4-chat-service ./chat.go
	go build -v -x -o ./bin/ew4-chat-broker ./broker.go

dep:
	go install -a chat
	rm -rf pkg/linux_amd64/chat.a

clean:
	rm -rf ./bin/ew4-*
	rm -rf ./bin
	rm -rf ./pkg
all:
	clean
	dep
	build
```

## FAQ

### 1.

```
# cmd/pprof
/root/go1.4/pkg/linux_amd64/runtime/cgo.a(_all.o): unknown relocation type 42; compiled without -fpic?
/root/go1.4/pkg/linux_amd64/runtime/cgo.a(_all.o): unknown relocation type 42; compiled without -fpic?
runtime/cgo(.text): unexpected relocation type 298
runtime/cgo(.text): unexpected relocation type 298
# cmd/go
/root/go1.4/pkg/linux_amd64/runtime/cgo.a(_all.o): unknown relocation type 42; compiled without -fpic?
/root/go1.4/pkg/linux_amd64/runtime/cgo.a(_all.o): unknown relocation type 42; compiled without -fpic?
runtime/cgo(.text): unexpected relocation type 298
runtime/cgo(.text): unexpected relocation type 298
```

情景描述

fedora下编译go1.4时报上述错误.

解决方法(设置环境变量貌似没用)

```
$ env CGO_ENABLED=0 ./make.bash
```

### 2. 

```
package context: unrecognized import path "context" (import path does not begin with hostname
```

go get一个go工程时出现这个问题. 

原因貌似是由于当前go版本过低(当时是1.6.3), 将其升级到1.8.3问题解决.