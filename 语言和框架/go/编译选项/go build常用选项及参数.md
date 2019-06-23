# go build常用选项及参数

参考文章

1. [关于Go tools的比较有用的flags](https://gocn.vip/article/6)

使用`go help build`可以查看所有可用的编译选项, 常用的有`-gcflags`, `-ldflags`等. 这些编译选项被`build`, `clean`, `get`, `install`, `list`, `run`, `test`子命令所共用.

为了禁止编译器优化和内联, 你可以使用`gcfalgs`:

```
$ go build -gcflags='-N -l'
```

在go build编译选项中, `-asmflags`, `-gccgoflags`, `-gcflags`与`-ldflags`接受的参数格式相同, 都是用(单/双)引号包裹的, 以空格分隔的参数列表. 这些参数会在`build`期间传递给底层的go tools.

使用`go tool compile --help`可以查看所有可用的编译参数.

常用参数

`-N`: 禁用编译器优化
`-l`: 禁用内联
`-m`: 打印编译器优化的详细描述
`-race`: 开启竞态检测
