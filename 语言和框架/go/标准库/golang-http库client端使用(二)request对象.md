# golang-http库使用(二)request对象

参考文章

1. [golang中发送http请求的几种常见情况](https://studygolang.com/articles/4489)

直接调用`http.Get()/http.Post()`的确够简单方便, 但是不够灵活. 如同python urllib中的urlopen直接打开url一样.

还有另外一种方法, 借助`Reqeust`对象可以实现对请求头的自定义, 比如设置UA与代理.

先创建`http.Client` -> 再创建`http.Request` -> 之后提交请求：`client.Do(request)` -> 处理返回结果，每一步的过程都可以设置一些具体的参数，

```go
package main

import "log"
import "net/http"

import "io/ioutil"

func main() {
    client := &http.Client{}
    url := "https://www.baidu.com"
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        panic(err)
    }   
	// req.Header.Set("User-Agent", "curl/7.54.0")
        
    res, err := client.Do(req)

    result, err := ioutil.ReadAll(res.Body)
    log.Printf("%s\n", result)
}
```

`Client`结构中有一个`Transport`成员, 很强大, 但我们常用的还是对request对象的修改.

```go
NewRequest func(method, url string, body io.Reader) (*Request, error)
```

其中`method`必须为大写形式, 如`GET`, `POST`, `PUT`, `OPTION`等.

```go
req.Header.Set("User-Agent", "curl/7.54.0")
```

------

然后是post请求的实现...