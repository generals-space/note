# golang-reflect.Value对象的Type与Kind方法

参考文章

1. [golang中的reflect包用法](https://www.cnblogs.com/andyidea/p/6193606.html)
	- `reflect.Value`类型的`Type()`和`Kind()`的联系和区别
	- `reflect.Kind`列举出`Kind()`可以得到的所有原生类型
2. [apimachinery-v0.17.2](https://github.com/kubernetes/apimachinery/blob/v0.17.2/pkg/conversion/helper.go#L27)

```go
// EnforcePtr 确认 obj 对象为指针类型, 然后返回该指针的引用类型.
//
// EnforcePtr ensures that obj is a pointer of some sort. 
// Returns a reflect.Value of the dereferenced pointer, 
// ensuring that it is settable/addressable.
//
// Returns an error if this is not possible.
func EnforcePtr(obj interface{}) (reflect.Value, error) {
	// v 是一个 reflect.Value{} 对象
	v := reflect.ValueOf(obj) 
	// 如果 obj 是一个指针类型变量, 那么 v.Kind() 必然是 reflect.Ptr.
	if v.Kind() != reflect.Ptr {
		if v.Kind() == reflect.Invalid {
			return reflect.Value{}, fmt.Errorf("expected pointer, but got invalid kind")
		}
		return reflect.Value{}, fmt.Errorf("expected pointer, but got %v type", v.Type())
	}
	// nil 不是指针, 但也没办法取指针.
	if v.IsNil() {
		return reflect.Value{}, fmt.Errorf("expected pointer, but got nil")
	}
	// 运行到这里, 说明 obj 是一个指针对象, Elem() 方法对其取指针, 得到其引用类型
	return v.Elem(), nil
}

```
