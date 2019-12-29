## 1.

```
$ go mod download
warning: pattern "all" matched no module dependencies
```

问题分析: 也许当前工程目录放在了GOPATH路径中, 而go mod工程必须放在GOPATH路径之外.

## 2.

```
~/Code/playground/go-mod $ go mod init
go: cannot determine module path for source directory /Users/general/Code/playground/go-mod (outside GOPATH, no import comments)
```

问题分析: `go mod init`后需要一个参数指定包名, 比如`github.com/generals-space/test`
