# 用gdb替换dlv[delve gdb dlv]

参考文章

1. [go语言有哪些好的debug方法？ - 幂恩的回答 - 知乎](https://www.zhihu.com/question/40980436/answer/767289819)
    - MemStats & GC
    - PProf
    - Trace
    - GDB & Delve
2. [golang调试工具Delve](https://www.cnblogs.com/li-peng/p/8522592.html)

这里的调试不是指`pprof`那种监控, 而是断点, 单步调试的工具.

按照参考文章1, delve更适合golang调试, gdb虽然可以用, 但并不是最佳方案.

`dlv`貌似是随vscode的golang插件一起安装的, 不需要额外动手...

但我在使用delve时, 发现~~ta只能调试源码~~, 就是只能`dlv main.go xxx`.

> dlv(1.7.1)也可以调试二进制文件, 使用`dlv exec runc`即可.

```log
$ dlv debug runc run mycontainer01
can't load package: package runc: cannot find package "runc" in any of:
	/usr/local/go/src/runc (from $GOROOT)
	/usr/local/gopath/src/runc (from $GOPATH)
...
exit status 1
```

我要调试的工程是`runc`, ta要求在执行`runc run 容器名称`时, 所在目录下有待创建的容器的配置文件和文件系统目录, 这些放在源码目录不合适. 所以我需要将代码编译出来, 然后再在执行命令时调试, 所以只能用gdb了.

好在golang的二进制文件在用gdb调试时, 和c程序没什么差别, 还挺方便.
