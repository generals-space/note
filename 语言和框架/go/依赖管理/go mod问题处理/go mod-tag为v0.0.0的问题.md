# go mod-tag为v0.0.0的问题

参考文章

学 kuber 源码的时候, 想在本地用 [go-restful](https://github.com/emicklei/go-restful) 写个小 demo, 使用 go mod 初始化一个仓库, 然后 go get 下载 2.9.5 版本(kuber v1.17.3 使用的就是这个版本), 但是却报如下错误

```console
$ go get -v github.com/emicklei/go-restful@v2.9.5
go: finding github.com v2.9.5
go: finding github.com/emicklei/go-restful v2.9.5
go: finding github.com/emicklei v2.9.5
go get github.com/emicklei/go-restful@v2.9.5: github.com/emicklei/go-restful@v2.9.5: reading https://goproxy.cn/github.com/emicklei/go-restful/@v/v2.9.5.info: 404 Not Found
```

此时的 go proxy 为 `https://goproxy.cn`, 于是换了个 `https://mirrors.aliyun.com/goproxy/`.

这回倒是能下载了, 但是 go get 完成后, go.mod 的内容成了这样

```go
require (
	github.com/emicklei/go-restful v0.0.0-20190516080722-b993709ae1a4 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
)
```

明明 github.com 上有明确的版本号 v2.9.5, 到了 go.mod 里就成了`v0.0.0-20190516080722-b993709ae1a4`, 这样真的很难辨别啊, 换了其他版本也一样, 后面只显示 commit id, 不显示版本号.

后来又将 go proxy 换成了`https://goproxy.io`, 再执行 go get , go.mod 的信息就正常了.

```go
require (
	github.com/emicklei/go-restful v2.9.5+incompatible // indirect
	github.com/json-iterator/go v1.1.10 // indirect
)

```
