# golang-struct&tag

<!keys!>: <!反射!> <!结构体!> <!标签!>

参考文章

1. [Go struct tag介绍](http://www.baiyuxiong.com/?p=1165)

2. [Golang通过反射实现结构体转成JSON数据](http://blog.cyeam.com/golang/2014/08/11/go_json)

1. 什么是struct的tag?

```go
type User struct {
    Name   string `user name`
    Passwd string `user password`
}
```

上面代码里反引号\`\`中的部分就是tag。

2. tag能用来干什么？

tag一般用于表示一个映射关系，最常见的是json解析中：

```go
type User struct {
    Name   string `json:"name"`
    Passwd string `json:"password"`
}
```

这个代码里，解析时可以把json中”name”解析成struct中的”username”.

3. tag定义必须用键盘ESC键下面的那个吗？

不是，用双引号也可以：

```go
type User struct {
    Name string "user name"
    Passwd string "user passsword"
}
```

4. 怎么获取struct的tag?

用反射：

```go
package main

import "fmt"
import "reflect"

type User struct {
	Name   string `json:"name"`
	Passwd string `json:"password"`
}

func main() {
	user := &User{"chronos", "pass"}
	//通过反射获取type定义
	theType := reflect.TypeOf(user).Elem()
	// 遍历此类型的成员字段
	for i := 0; i < theType.NumField(); i++ {
		// 将tag输出出来(Field()函数看来是按索引值来获取成员字段的)
		_tag := theType.Field(i).Tag
		fmt.Println(_tag)
		_tagVal := _tag.Get("json") // 获取该字段json标签下的值
		fmt.Println(_tagVal)
	}
}
```

执行结果为

```
json:"name"
name
json:"password"
password
```