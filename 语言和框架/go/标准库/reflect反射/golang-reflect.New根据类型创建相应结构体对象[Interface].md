# golang-reflect.New根据类型创建相应结构体对象

参考文章

1. [apimachinery-v0.17.2](https://github.com/kubernetes/apimachinery/blob/v0.17.2/pkg/runtime/scheme.go#L288)
2. [Go 语言高性能编程 - Go Reflect 提高反射性能](https://geektutu.com/post/hpg-reflect.html)

```go
// New 根据传入的 GVK 对象, 为其创建一个完成 runtime.Object 对象
// 如传入 {Group: "", Version: "v1", Kind: "Pod"}, 就返回一个 Pod{} 对象.
//
// caller: pkg/runtime/codec.go -> UseOrCreateObject()
//
// New returns a new API object of the given version and name, 
// or an error if it hasn't been registered. 
// The version and kind fields must be specified.
func (s *Scheme) New(kind schema.GroupVersionKind) (Object, error) {
	if t, exists := s.gvkToType[kind]; exists {
		return reflect.New(t).Interface().(Object), nil
	}

	if t, exists := s.unversionedKinds[kind.Kind]; exists {
		return reflect.New(t).Interface().(Object), nil
	}
	return nil, NewNotRegisteredErrForKind(s.schemeName, kind)
}

```

示例代码

```go
package main

import (
    "reflect"
    "log"
)

type Student struct {
    Name string
}

func main(){
    s1 := Student{}
    typ := reflect.TypeOf(s1)
    log.Printf("%s\n", typ)
	// 好像 TypeOf() 的类型不能是指针, 不然 Interface() 后面的类型转换不好写.
    s2 := reflect.New(typ).Interface().(*Student)
    s2.Name = "general"
    log.Printf("%+v\n", s2)
}
```
