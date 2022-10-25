# golang-url操作

参考文章

1. [golang url解析](http://www.cnblogs.com/benlightning/articles/4441027.html)
2. [Combine URL paths with path.Join()](https://stackoverflow.com/questions/34668012/combine-url-paths-with-path-join)
	- url拼接

## 1. 解析

`net/url`包中关于解析的`Parse`函数有3个:

1. `Parse()`
2. `ParseQuery()`
3. `ParseRequestURI()`

这3个看名字就可以明白其作用. `Parse()`的示例如下.

```go
func main() {
	// 这个URL包含了一个 scheme, 认证信息, 主机名, 端口, 路径, 查询参数和片段.
	urlStr := "https://user:pass@www.baidu.com:8080/path?key1=val1&key2=val2#frag"
	urlObj, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}

	log.Println(urlObj.Scheme) // https

	// User属性包含了所有的认证信息, 这里调用Username()和Password()方法来获取独立值.
	user := urlObj.User
	username := user.Username()
	passwd, _ := user.Password()
	log.Println(user)     // user:pass
	log.Println(username) // user
	log.Println(passwd)   // pass

	// Host 同时包括主机名和端口信息, 如端口存在的话, 使用 strings.Split() 从 Host 中手动提取端口.
	log.Println(urlObj.Host) // www.baidu.com:8080
	hostArray := strings.Split(urlObj.Host, ":")
	log.Println(hostArray[0])      // www.baidu.com
	log.Println(hostArray[1])      // 8080
	log.Println(urlObj.Hostname()) // www.baidu.com
	log.Println(urlObj.Port())     // 8080, 字符串类型
	log.Println(urlObj.Path)       // /path
	log.Println(urlObj.Fragment)   // frag

	// RawQuery()可以得到a=1&b=2原始查询字符串.
	log.Println(urlObj.RawQuery) // key1=val1&key2=val2
	// Query()方法得到查询参数map, 格式见示例.
	// queryMap为url.Values对象, 而Values为map[string][]string别名.
	queryMap := urlObj.Query()
	// url.ParseQuery(queryString) 可以解析指定的查询参数字符串, 常用.
	// queryMap, _ := url.ParseQuery(urlObj.RawQuery)
	log.Println(queryMap)            // map[key1:[val1] key2:[val2]]
	log.Println(queryMap["key1"][0]) // val1
	log.Println(queryMap["key2"][0]) // val2
}
```

## 2. 拼接

拼接可用url对象的`ResolveReference()`方法, 使用示例如下.

```go
func main() {
	urlStr := "https://www.baidu.com/path1/path2?key=val#frag"
	urlObj, _ := url.Parse(urlStr)

	subURL := "path3?key=subquery#subfrag"
	subURLObj, _ := url.Parse(subURL)
	fullURLObj := urlObj.ResolveReference(subURLObj)
	log.Println(fullURLObj.String()) // https://www.baidu.com/path1/path3?key=subquery#subfrag
}
```

## 3. 编解码

1. `urlObj.EscapedPath()`
2. `url.QueryEscape(str)`
3. `url.PathEscape(str)`

