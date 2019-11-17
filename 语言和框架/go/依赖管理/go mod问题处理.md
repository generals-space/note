```
$ go mod download
warning: pattern "all" matched no module dependencies
```

问题分析: 也许当前工程目录放在了GOPATH路径中, 而go mod工程必须放在GOPATH路径之外.
