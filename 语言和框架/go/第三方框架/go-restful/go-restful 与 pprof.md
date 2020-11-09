# go-restful 与 pprof

参考文章

1. [kuber v1.17.3 源码](https://github.com/kubernetes/kubernetes/blob/v1.17.13/pkg/kubelet/server/server.go#L478)

`net/http/pprof`下的各 handler 都实现了 golang 内置的 `http.HandlerFunc`, 有一些第三方的 http 框架自身都兼容内置框架, 直接接受`http.HandlerFunc`类型的处理函数, 如[httprouter](https://github.com/julienschmidt/httprouter)

```go
router := httprouter.New()
router.HandlerFunc("GET", "/debug/pprof/", pprof.Index)
http.ListenAndServe(":8080", router)
```

但是`go-restful`不行, 如果想要使用`net/http/pprof`, 需要对其做一层包装. 

```go
	// 一个 WebService 表示一个前缀对象, 比如 /user, 通过 Path() 方法设置.
	// 之后可以通过 ws.Route() 在这个前缀下添加各种操作.
	ws := new(restful.WebService)
	pprofBasePath := "/debug"
	// 设置 pprof 的基本路径前缀
	ws = new(restful.WebService).Path(pprofBasePath)

	handlePprofEndpoint := func(req *restful.Request, resp *restful.Response) {
		name := strings.TrimPrefix(req.Request.URL.Path, pprofBasePath)
		switch name {
		case "profile":
			pprof.Profile(resp, req.Request)
		case "symbol":
			pprof.Symbol(resp, req.Request)
		case "cmdline":
			pprof.Cmdline(resp, req.Request)
		case "trace":
			pprof.Trace(resp, req.Request)
		default:
			pprof.Index(resp, req.Request)
		}
	}

	ws.Route(ws.GET("/{subpath:*}").To(func(req *restful.Request, resp *restful.Response) {
		handlePprofEndpoint(req, resp)
	})).Doc("pprof endpoint")

	restful.Add(ws)

	http.ListenAndServe(":8080", nil)
```

上述代码来自参考文章1, kubernetes 的 kubelet 源码, 对应的 go.mod 内容如下

```go
module test

go 1.13

require (
	github.com/emicklei/go-restful v2.9.5+incompatible
	github.com/json-iterator/go v1.1.10 // indirect
)
```

通过`http://172.16.91.10:8080/debug/profile`, 可以访问`profile`页面, 其他同理.
