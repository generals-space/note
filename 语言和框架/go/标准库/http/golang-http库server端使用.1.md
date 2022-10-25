# golang-http库server端使用.1

参考文章

1. [golang http 服务器编程](https://juejin.im/post/58cffa535c497d0057cfcdfe)
2. [golang http.FileServer 遇到的坑](https://blog.csdn.net/liangguangchuan/article/details/60326495)

## 1. 获取GET/POST参数

```go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// getHandler: 获取get请求参数
func getHandler(resp http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query["name"][0]
	title := query["title"]
	log.Println(name)
	log.Println(title)

	data := map[string]interface{}{
		"name":  name,
		"title": title,
	}
	result, _ := json.Marshal(data)
	resp.Write([]byte(result))
}

// postHandler: 获取post请求json数据
func postHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	byteData, _ := ioutil.ReadAll(req.Body)
	log.Printf("%s\n", byteData)
	// var mapData map[string]interface{}
	mapData := map[string]interface{}{}
	err := json.Unmarshal(byteData, &mapData)
	if err != nil {
		log.Println(err)
	}

	mapResult := map[string]interface{}{
		"username": mapData["username"],
		"password": mapData["password"],
	}
	result, _ := json.Marshal(mapResult)
	resp.Write(result)
}

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	log.Println("http server started...")
	http.ListenAndServe(":8079", nil)
}
```

测试结果如下

```console
$ curl 'localhost:8079/get?name=general&title=ceo&title=manager'
{"name":"general","title":["ceo","manager"]}
$ curl -X POST -d '{"username": "general", "password": "123456"}' localhost:8079/post
{"password":"123456","username":"general"}
```

关于POST请求的参数的获取, 虽然有人说在手动调用`req.ParseForm()`方法后可以通过`req.Form`变量得到POST的参数, 但是得到的是一个map类型, 而且内容非常不合理(上面POST函数测试时得到了一个`map[{"username": "general", "password": "123456"}:[]]`...这什么东西?). 还是用解析`Body`中的数据更简单方便.

------

咳, 打脸了.

通过解析`req.Body`中的内容获取的信息是json格式的数据, 如ajax, postman测试等.

但是当使用`form`标签元素, 通过submit类型的按钮点击直接提交时, Body里是没有内容的...

此时我们只能先调用`req.ParseForm()`, 然后`req.Form`就表示了提交来的数据, 类型为`map[string]interface{}`, 如`map[username:[admin] password:[123456]]`. 

注意: 同一个name可能有多个值, 所以取值时需要这样`req.Form["username"][0]`, 或者`req.Form.Get("username")`.

## 理解: http标准库的缺陷及使用第三方http server的必要性

虽然 net/http 提供的各种功能已经满足基本需求了, 但是很多时候还不够方便, 比如: 

1. 不支持 URL 匹配, 所有的路径必须完全匹配, 不能捕获 URL 中的变量(像`/id/1/name/general`这种restful形式的接口), 不够灵活.

2. 不支持 HTTP 方法匹配. 没有地方可以指定`GET`, `POST`等请求类型, 只能在函数内部通过判断`Method`分别处理.

3. 不支持扩展和嵌套, URL 处理都在都一个 ServeMux 变量中(类似flask中的蓝图, 前缀匹配, 方便模块化).
