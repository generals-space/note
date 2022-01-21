# golang-reflect.New根据类型创建相应结构体对象

参考文章

1. [apimachinery-v0.17.2](https://github.com/kubernetes/apimachinery/blob/v0.17.2/pkg/runtime/scheme.go#L288)

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
