# golang-http库发送&处理form请求

参考文章

1. [golang 模拟浏览器 发送application/x-www-form-urlencoded的post请求](https://www.jianshu.com/p/88bf34b5d253)

2. [Golang发送post表单请求](https://blog.csdn.net/gophers/article/details/22870185)

与普通的json数据不同, form信息不能通过`json.Marsharl()`方法放入请求体, 只能用`url.Values`结构存放.

示例如下:

```go
	contentType := "application/x-www-form-urlencoded"

	dataURLVal := url.Values{}
	dataURLVal.Add("appKey", appKey)
	dataURLVal.Add("appSecret", appSecret)
	data := strings.NewReader(dataURLVal.Encode())
	resp, err := http.Post(conf.CameraServiceYS7Addr, contentType, data)
	if err != nil {
		l.Errorf("request ys to get access token failed: %s", err.Error())
		return
	}
```

> `net/http`包中还有一个`PostForm()`方法, 直接发送form请求, 不用设置`application/x-www-form-urlencoded`请求头.

> `url.Values`其实是一个`map`对象.

参考文章2中还写了如何在服务端处理`form`请求中的数据, 可以使用request对象的`PostFormValue`方法.