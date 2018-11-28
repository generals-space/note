# golang-json解析

参考文章

1. [go json解析Marshal和Unmarshal](http://www.baiyuxiong.com/?p=923&utm_source=tuicool&utm_medium=referral)

2. [JSON与Go](http://rgyq.blog.163.com/blog/static/316125382013934153244/)

3. [Golang 解析 json 数据](https://blog.tanteng.me/2017/07/golang-decode-json/)

json字符串的解析结果一般对应语言中的`dict(python)`, `map(java)`, `object(js)`等相似结构, 在go中为结构体struct.

在go中, json的序列化与反序列化主要涉及结构体与`[]byte`类型的相互转换.

参考文章1中提供了简单易读的示例可以看一下.

> 注意: 结构体的成员首字母大写, 不然`Marshal`得不到其中的成员.

参考文章3中给出了除了标准json字符串转结构体, 还可以把切片数组与字符串相互转换.

## 1. 字符串 < -- > 结构体

```go

```

## 2. 字符串 < -- > 切片数组

简单示例

```go
	jsonData := []byte(`["a","b","c","d","e"]`)
	var a []string
	json.Unmarshal(jsonData, &a)
	log.Printf("%+v\n", a) // 输出: [a b c d e]
```

复杂示例

```go
package main

import "log"
import "encoding/json"

type User struct {
	Name string
	Age  int
}

func main() {
	// 1. 切片转字符串
	users := []User{
		User{Name: "general", Age: 21},
		User{Name: "jiangming", Age: 24},
	}
	str, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	// 输出: [{"Name":"general","Age":21},{"Name":"jiangming","Age":24}]
	log.Printf("%s\n", str)

	// 2. 字符串转切片数组
	oriStr := []byte(`
		[{"Name":"general","Age":21},{"Name":"jiangming","Age":24}]
	`)
	var newUsers []User
	// newUsers := []User{}
	// 注意Unmarshal参数2的取指符号&
	err = json.Unmarshal(oriStr, &newUsers)
	// 输出: [{Name:general Age:21} {Name:jiangming Age:24}]
	log.Printf("%+v\n", newUsers)
}
```