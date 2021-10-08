# golang-http库server端使用.2.静态文件服务

参考文章

1. [golang http 服务器编程](https://juejin.im/post/58cffa535c497d0057cfcdfe)
2. [golang http.FileServer 遇到的坑](https://blog.csdn.net/liangguangchuan/article/details/60326495)
3. [Golang1.8标准库http.Fileserver跟http.ServerFile小例子](https://blog.csdn.net/fyxichen/article/details/60570484)

按照参考文章1中所说, 大部分的服务器逻辑都需要使用者编写对应的 Handler, 不过有些 Handler 使用频繁, 因此 `net/http` 提供了它们的实现. 比如负责静态文件的 `FileServer`、负责 404 的`NotFoundHandler` 和 负责重定向的`RedirectHandler`. 

```go
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/tmp"))))
	http.HandleFunc("/get", getHandler)
    http.HandleFunc("/post", postHandler)
}
```

这样, 在浏览器中访问`localhost:8079/static/`就可以看到`/tmp`目录下的文件列表了.

------

网上大多文章都直接使用`http.FileServer()`就能处理静态文件请求, 如下

```go
	http.Handle("/static/", http.FileServer(http.Dir("/tmp")))
```

但我在实验的时候却得到了404. 后来按照参考文章1中, 加上了`http.StripPrefix()`方法, 才成功处理静态请求.

那`StripPrefix()`方法是做什么的呢? 官方文档如下

```
func StripPrefix(prefix string, h Handler) Handler

StripPrefix返回一个Handler, 该Handler会将请求的URL.Path字段中给定前缀prefix去除后再交由h处理. StripPrefix会向URL.Path字段中没有给定前缀的请求回复404 page not found.
```

没懂.

说的明白点, 在`http.Handle("/static/", http.FileServer(http.Dir("/tmp")))`这句中, `"/static/"`相当于nginx中的`location`值, 而`http.Dir("/tmp")`类似于nginx中的`root`字段值. 当用户访问`/static/`时, 其实是查询`/tmp/static/`目录下的文件. 所以我尝试在`/tmp`目录下新建了一个`static`子目录, `touch`一些文件, 然后不加`StripPrefix()`处理, 但这次可以访问到了.

> 同样注意路由路径字符串中尾部的斜线`/`, 如果是目录最好加上.

## 高级处理

理论上nginx能做的事件, `net/http`也能做. 如果限制只有登录用户可以访问静态服务, 或者限制`referer`头(尤其做反爬, 防盗链时), 该怎么做?

咳, 当然用nginx更方便, 没必要在程序里做这些. 这里只是展示一下可能的方式.

```go
var staticfs = http.FileServer(http.Dir("/tmp"))

//这里可以自行定义安全策略
func static(resp http.ResponseWriter, req *http.Request) {
	log.Printf("访问静态文件: %s\n", req.URL.Path)
	if strings.HasSuffix(req.URL.Path, ".ini") {
		resp.Write([]byte("禁止访问!"))
		return
	}
	staticfs.ServeHTTP(resp, req)
}

func main() {
	http.HandleFunc("/static/", static)
	log.Println("http server started...")
	http.ListenAndServe(":8079", nil)
}
```