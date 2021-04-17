
参考文章

1. [change import for github.com/googleapis/gnostic/OpenAPIv2](https://github.com/kubernetes/client-go/issues/743)

我所遇到的问题与参考文章1中所说的相同, `go mod tidy`总会报错说

```
k8s.io/client-go/discovery/discovery_client.go:30:2: cannot find package "github.com/googleapis/gnostic/OpenAPIv2"
The above import is renamed to
"github.com/googleapis/gnostic/openapiv2"
```

其实是因为 client-go 仓库中并没有规定`googleapis/gnostic`的版本号, 这导致我的工程在使用`client-go`时为ta自动下载了最新版的`googleapis/gnostic`, 但是最新版的库中的目录路径已经变了, 没有`OpenAPIv2`这个目录了.

解决方法就是, 在我的工程里手动用`go get -v github.com/googleapis/gnostic@v0.4.0`下载旧版本, 这样在`go.mod`中会出现这条记录, 并且会包含`indirect`标记, 因为我的工程本身并没有直接用到这个库.
