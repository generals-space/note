# golang-gob序列化与反序列化

<!--
<!key!>: {c2f8cf90-5387-11e9-ae66-aaaa0008a014}
<!link!>: {664fdd98-537c-11e9-b398-aaaa0008a014}
-->

参考文章

1. [golang encoding/gob包使用demo](https://blog.csdn.net/qq_21816375/article/details/80022298)
2. [golang利用gob序列化struct对象保存到本地](https://www.cnblogs.com/reflectsky/p/golang-gob-struct.html)
3. [Golang Gob编码](http://www.cnblogs.com/yjf512/archive/2012/08/24/2653697.html)
4. [使用golang gin框架sessions时碰到的gob问题](https://my.oschina.net/sannychan/blog/1840048)
5. [Package gob](https://golang.org/pkg/encoding/gob/#GobEncoder)

gob序列化也是在使用[faygo](https://github.com/henrylee2cn/faygo)框架时, redis作session分布式缓存时遇到的, 存入redis的自定义对象取出时为空, 排查问题时发现是在服务启动前需要加一句`gob.Register()`. 

## 1. Encode与Decode

golang 的 gob 与 python 的 pickle 作用相同. 

简单示例如下

```go
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// P ...
type P struct {
	Name string
	Age  int
}

// Q ...
type Q struct {
	Name string
	Age  *int
}

func main() {
	// 作为发送方, 将struct对象序列化为bytes类型
	var buf bytes.Buffer
	sender := P{
		Name: "general",
		Age:  21,
	}
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(sender)
	if err != nil {
		log.Fatal("encode error:", err)
	}
    // 接收方从bytes对象中反序列化为指定类型的对象.
    // 要成功反序列化本地保存的对象, 前提是要知道本地保存的struct的结构
	var receiver Q
	decoder := gob.NewDecoder(&buf)
	err = decoder.Decode(&receiver)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%+v\n", receiver)                                   // {Name:general Age:0xc0000623f8}
	fmt.Printf("name: %s, age: %d\n", receiver.Name, *receiver.Age) // name: general, age: 21
}
```

> 注意: gob序列化struct实例与json的`Unmarshal()`相似, 导出字段需要大写字母开头.

另外, golang中的struct无非是由基本类型组成的结构, 指针可作为普通数值看待, 又无引用类型影响. 所以解析手段都有迹可循.

参考文章3中对此作了详细的讲解.

## 2. Register和RegisterName

这两个函数, 虽然参考文章3给出了使用方法, 但并没有实际的使用场景. 更真实的场景可以见参考文章4.

实际上我也曾想在session中存储自定义struct对象(还是指针), 遇到过无法从session获得相应对象的问题. 

```go
	u := &types.User{}
	if u.Uid != "" {
		session := sessions.Default(c)
		session.Set("login_user", u)
		session.Save()
		return u, nil
	}
```

所以在服务启动前, 需要事先使用`Register()`注册可能的对象类型.

```go
gob.Register(&types.User{})
```

这句话告诉系统: 所有的不可知类型(序列化操作应该是按`interface{}`作为通用类型)是有可能为`&types.User{}`结构的. 

注意: `session.Set()`如果指定了指针类型, 那么对应的, `gob.Register()`注册的也应该是指针.

------

不过如何序列化`channel`, `func`, 倒是没什么头绪, 网上也没有相关资料. 参考文章5中对`GobEncoder`接口函数的介绍, 不过没看太懂.
