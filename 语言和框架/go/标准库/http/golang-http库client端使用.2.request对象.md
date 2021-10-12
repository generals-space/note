# golang-http库client端使用.2.request对象

参考文章

1. [golang中发送http请求的几种常见情况](https://www.cnblogs.com/Goden/p/4658287.html)

直接调用`http.Get()/http.Post()`的确够简单方便, 但是不够灵活. 如同python urllib中的urlopen直接打开url一样.

还有另外一种方法, 借助`Reqeust`对象可以实现对请求头的自定义, 比如设置UA与代理.

先创建`http.Client` -> 再创建`http.Request` -> 之后提交请求：`client.Do(request)` -> 处理返回结果, 每一步的过程都可以设置一些具体的参数.

```go
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

如下示例实现了, post请求, 带json参数, 修改请求头字体, 解析响应体, 够详细了.

```go
// LoginInfo ...
type LoginInfo struct {
	Action string `json:"action"`
	Params struct {
		LoginID string `json:"login_id"`
		Passwd  string `json:"passwd"`
		Type    int    `json:"type"`
	} `json:"params"`
}
// LoginRespPayload ...
type LoginRespPayload struct{
	Token string	`json:"token"`
}
// LoginResp ...
type LoginResp struct {
	Token  string `json:"token"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Total  int `json:"total"`
	Result LoginRespPayload `json:"result"`
}

func main() {
	serverAddr := "http://sapigs.vipvoice.link/user/checkin.boxuds"
	client := &http.Client{}

	// loginInfo := make(map[string]interface{})
	// loginInfo["action"] = "user/checkin"
	// loginInfo["params"] = map[string]interface{}{
	// 	"login_id": "13073197649",
	// 	"passwd": "123456",
	// 	"type": 0,
	// }
	loginInfo := &LoginInfo{
		Action: "user/checkin",
	}
	loginInfo.Params.LoginID = "13073197649"
	loginInfo.Params.Passwd = "123456"
	loginInfo.Params.Type = 0

	bytesData, err := json.Marshal(loginInfo)
	if err != nil {
		log.Fatal(err)
		return
	}
	// log.Println(string(bytesData))
	// return
	jsonData := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", serverAddr, jsonData)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "http://ip.vipvoice.link/login?redirect=%2F")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	rsp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer rsp.Body.Close()

	rspData, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(rspData))
	loginRsp := &LoginResp{}
	json.Unmarshal(rspData, loginRsp)
	log.Printf("%+v", loginRsp)
}
```
