# golang-reflect反射.2.通过反射修改结构体成员的值

参考文章

1. [golang-利用反射给结构体赋值](https://www.cnblogs.com/fwdqxl/p/7789162.html)

在golang中, 结构体的成员的取值一般为`对象名.字段名`, 而字段名是没有引号的(像`user["name"]`这种). 所以当要修改的字段不确定时, 我们没有办法像在python中使用`__dict__`一样方便地修改结构体成员的值. 

比如一个结构体对象

```go
type MyStruct struct{
	Attr01 string
	Attr02 string
	Attr03 string
	...
}
```

而我们有一个map的键值对

```go
map[string]string{
	"Attr01": "Value01",
	"Attr02": "Value02",
	"Attr03": "Value03",
	...
}
```

如何将map中的值更新到结构体对象中?(除了先`Marshal(map)`, 再`Unmarshal(struct)`到结构体.)

这就需要用到反射了. 参考文章3中给出了详细的做法, 这里贴一个简短的示例.

```go
package main
import (
	"fmt"
	"reflect"
	"unsafe"
)
// User ...
type User struct {
	Name string
	Age int
}
func main(){
	user := &User{
		Name: "general",
		Age: 21,
	}
	fmt.Printf("%+v\n", user)

	nameField := reflect.ValueOf(user).Elem().FieldByName("Name")
	addrOfName := nameField.Addr().Pointer()	// 这里是 uintptr 类型
	// 先用 Pointer() 将 uintptr 转换成通用指针, 再用 *string 将其按照字符串指针对待,
	// 再加一个*取值再赋值.
	*(*string)(unsafe.Pointer(addrOfName)) = "jiangming"
	
	ageField := reflect.ValueOf(user).Elem().FieldByName("Age")
	addrOfAge := ageField.Addr().Pointer()
	*(*int)(unsafe.Pointer(addrOfAge)) = 26

	fmt.Printf("%+v\n", user)
}
```
