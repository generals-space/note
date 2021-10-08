# golang-http库client端使用.1.get与post方法示例

参考文章

1. [golang HTTP客户端](https://studygolang.com/articles/253)
2. [golang中发送http请求的几种常见情况](https://studygolang.com/articles/4489)
3. [Golang Web编程的Get和Post请求发送与解析](https://blog.csdn.net/typ2004/article/details/38669949)

值得一看的是, 参考文章2中`Post`方法发送json数据时, 如何构造json请求体.

## 1. Get方法获取网页并读取内容

```go
package main

import "log"
import "net/http"

// import "io/ioutil"

func main() {
    addr := "https://www.baidu.com"
    // 直接发送Get请求, res为Response类型
	res, err := http.Get(addr)
	if err != nil {
		panic(err)
	}
	// 这一句貌似是必须的
	defer res.Body.Close()
	/*
	 * header是Header类型, 而Header本身的定义为
	 * type Header map[string][]string
	 * 所以可以用很常规的方法获取到响应头信息
	 */
	header := res.Header
	for k, v := range header {
		log.Printf("key: %s, value: %s", k, v)
	}
	log.Println("============================================")
	// Body是一个类似于Buffer的类型, 有两个方法可以读取其中的内容.
	// 1. 如果使用res.Body自己的Read方法, 就必须创建一个容器来承接
	// 但实际上不一定读取1024个字节, 有可能存在多于或少于的情况.
	bodyCnt := make([]byte, 1024)
	length, err := res.Body.Read(bodyCnt)
	log.Printf("length: %d\n", length)
	log.Printf("%s\n", bodyCnt)

	// 2. 也可以用ioutil库来读取内容
	// body, err := ioutil.ReadAll(res.Body)
	// log.Printf("%s\n", body)
}
```

## 2. Post方法发送json数据

`http`库能实现`Post`请求的函数有两个

1. `Post func(url string, contentType string, body io.Reader) (resp *Response, err error)`
2. `PostForm func(url string, data url.Values) (resp *Response, err error)`

前者适合发送二进制数据, 如文件, 图片, json数据等. 后者发送的是表单格式`application/x-www-form-urlencoded`的数据.

不过json数据需要构造, 尤其特定结构的数据需要从结构体序列化成json字符串.

```go
package main

import "log"
import "encoding/json"
import "bytes"
import "net/http"

import "io/ioutil"

type User struct {
	Name string
	Age  int
}

func main() {
	addr := "http://localhost:8080"
	// jsonData := url.Values{
	// 	"name": {"general"},
	// 	"age":  {"21"},
	// }
	// res, err := http.PostForm(addr, jsonData)

	user := User{
		Name: "general",
		Age:  21,
	}
	data, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	jsonData := bytes.NewBuffer(data)
	res, err := http.Post(addr, "application/json", jsonData)
	if err != nil {
		panic(err)
	}
	// 这一句貌似是必须的
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	log.Printf("%s\n", body)
}
```