# golang环境变量-GOOS与GOARCH跨平台编译

参考文章

1. [Golang 交叉编译与选择性编译](https://blog.csdn.net/dengming0922/article/details/82217929)
2. [Golang交叉编译各个平台的二进制文件](https://www.jianshu.com/p/efaef7940207)

交叉编译主要是由如下两个编译环境参数决定:

- `GOOS`: 目标平台的系统类型
- `GOARC`: 目标平台的处理器体系结构

`GOOS`可选值如下:

- `windows`(win系统)
- `darwin`(mac系统)
- `linux`(linux系统)
- ...

`GOARCH`可选值如下:

- `amd64`(常用64位系统架构))
- `386` 
- `arm`

在linux系统下编译运行在windows下的exe程序

```
GOOS=windows GOARCH=amd64 go build -o main.exe main.go
```

windows下编译运行在linux下的可执行程序

```
set GOARCH=386; set GOOS=linux; go build -o main main.go
```

> `-o main`要在`main.go`的前面

> 有时可能要加上`CGO_ENABLED=0`

...windows还是有问题, 在powerhshell执行上面的命令创建的arm64程序, 还是无法运行.

因为在ps中执行`set GOARCH=arm64`后, 再执行`go env`, 看到的`GOARCH`变量仍是原来的`amd64`. 后来尝试在cmd中执行, 结果`go env`没法执行了, 因为当前架构已经被修改, go程序没法运行了.

所以上面的命令应该在cmd里执行, 像这样.

```
set GOARCH=386
set GOOS=linux
go build -o main main.go
```

## 

linux

```
## x86
GOOS=linux GOARCH=amd64 go build -o xxx ./main.go
## arm
GOOS=linux GOARCH=arm64 go build -o xxx ./main.go
```

windows

```
set GOARCH=arm64
set GOOS=linux
go build -o xxx .\main.go
```
