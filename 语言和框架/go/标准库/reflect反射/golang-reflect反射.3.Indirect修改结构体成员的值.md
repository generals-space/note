# golang-reflect反射.3.Indirect修改结构体成员的值

参考文章

1. [golang-利用反射给结构体赋值](https://www.cnblogs.com/fwdqxl/p/7789162.html)
2. [golang中的reflect包用法](https://www.cnblogs.com/andyidea/p/6193606.html)
	- 通过`Indirect()`获取`Value`对象的指针对象
	- `SetXXX()`设置目标类型变量的值

本文中的做法比前文要简单很多.

```go
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
}

func (this *User) GetName() string {
	return this.Name
}
func (this *User) SetName(name string) {
	this.Name = name
}

func main() {
	user1 := &User{Name: "general"} // 指针
	// theVal := reflect.ValueOf(user1).Elem()
	theVal := reflect.Indirect(reflect.ValueOf(user1))
	fmt.Printf("%+v\n", theVal) // {Name:general}
	theVal.FieldByName("Name").SetString("jiangming")
	fmt.Printf("%+v\n", user1) // &{Name:jiangming}
}
```

至于应用场景, 以[controller-runtime](https://github.com/kubernetes-sigs/controller-runtime)为例, ta有一些api用于获取对象, 这个对象并不是通过函数返回的, 而是先构造一个对象当作参数传入, 然后由函数将内容填进去的.

```go
pod, err := client.CoreV1().Pods("ns").Get("name")
```

```go
pod := &corev1.Pod{}
client.Get(ns/name, pod)
```

在`Get()`方法里, pod是一个指针, 不能直接对一个指针类型的参数赋值, 那样是没办法把内容传到外面的, 只能修改 pod 参数的各种字段.

但是pod类型的字段太多了, 而且也不一定是`Pod{}`类型, 可能是`Node{}`或是`Service{}`等, 这样就需要用到反射.

```go
	outVal := reflect.ValueOf(out)
	objVal := reflect.ValueOf(obj)
	if !objVal.Type().AssignableTo(outVal.Type()) {
		return fmt.Errorf("cache had type %s, but %s was asked for", objVal.Type(), outVal.Type())
	}
	// 将 objVal 的内容赋值给 outVal
	reflect.Indirect(outVal).Set(reflect.Indirect(objVal))
	out.GetObjectKind().SetGroupVersionKind(c.groupVersionKind)
```
