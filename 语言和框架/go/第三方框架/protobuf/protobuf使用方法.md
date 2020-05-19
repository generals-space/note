# protobuf使用方法

参考文章

1. [Protocol Buffer Basics: Go](https://developers.google.com/protocol-buffers/docs/gotutorial)
    - 使用`go install`安装`protoc-gen-go`工具
2. [Protocol Buffers Release页](https://github.com/protocolbuffers/protobuf/releases)
    - 下载`protoc`编译器
3. [Switch from --go_out=plugins to -go-grpc_out PATH problem](https://stackoverflow.com/questions/61044883/switch-from-go-out-plugins-to-go-grpc-out-path-problem)
4. [protoc-gen-go-grpc: program not found or is not executable](https://stackoverflow.com/questions/60578892/protoc-gen-go-grpc-program-not-found-or-is-not-executable)

## 1. 安装

分两个步骤, 一个是使用`go install`安装`protoc-gen-go`, 一个是从参考文章2中下载`protoc`(参考文章2中有`protobuf`各种语言的二进制程序, 目前不清楚是用来做什么的.)

```
go install google.golang.org/protobuf/cmd/protoc-gen-go
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.12.0/protoc-3.12.0-linux-x86_64.zip
```

> `google.golang.org/protobuf/cmd/protoc-gen-go`这个库可能会有问题(下面会讲到), 最好换成`github.com/golang/protobuf/protoc-gen-go`.

## 2. 使用

示例`.proto`文件见当前目录中的`test.proto`

`protobuf`文件编译命令

```
protoc --go_out=plugins=grpc:. ./test.proto
```

> `protoc`的`--go_out`选项会调用`protoc-gen-go`工具, 所以`$GOPATH/bin`也需要添加到`$PATH`路径中.

...出错了?(之前一直这么用来着, 没出过错)

```console
$ protoc --go_out=plugins=grpc:. ./test.proto
--go_out: protoc-gen-go: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC
```

按照参考文章3和4中所说, 是因为`google.golang.org/protobuf`库中的`protoc-gen-go`不再支持`grpc`插件, 而且也找不早期的release版本了.

于是换成如下命令

```
$ protoc --go-grpc_out=. ./test.proto
protoc-gen-go-grpc: program not found or is not executable
```

...还是按照参考文章3和4, 能够生成`grpc`格式的插件`proto-gen-go-grpc`好像还没发布?

好在还有另外一个可用的库.

```
go install github.com/golang/protobuf/protoc-gen-go
```

同时换回使用

```
protoc --go_out=plugins=grpc:. ./test.proto
```
