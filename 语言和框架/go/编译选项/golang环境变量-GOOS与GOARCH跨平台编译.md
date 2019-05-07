# golang环境变量-GOOS与GOARCH跨平台编译

参考文章

1. [Golang 交叉编译与选择性编译](https://blog.csdn.net/dengming0922/article/details/82217929)

交叉编译主要是由如下两个编译环境参数决定:

- `GOOS`: 目标平台的系统类型
- `GOARC`: 目标平台的处理器体系结构

`GOOS`可选值如下:

- `windows`(win系统)
- `darwin`(mac系统)
- `linux`(linux系统)
- `dragonfly` 
- `freebsd` 
- `netbsd` 
- `openbsd` 
- `plan9` 
- `solaris` 

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
set GOARCH=386; set GOOS=windows; go build -o main main.go
```

> 有时可能要加上`CGO_ENABLED=0`