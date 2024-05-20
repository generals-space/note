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

## 

linux

```bash
## x86
GOOS=linux GOARCH=amd64 go build -o xxx ./main.go
## arm
GOOS=linux GOARCH=arm64 go build -o xxx ./main.go
```

windows

```bat
set GOARCH=arm64
set GOOS=linux
go build -o xxx .\main.go
```

> `-o main`要在`main.go`的前面

> 有时可能要加上`CGO_ENABLED=0`
