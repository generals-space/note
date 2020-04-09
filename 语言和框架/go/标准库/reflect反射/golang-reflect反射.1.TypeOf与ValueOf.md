# golang-reflect反射.1.TypeOf与ValueOf

参考文章

1. [Go语言_反射篇](https://www.cnblogs.com/yjf512/archive/2012/06/10/2544391.html)
2. [Golang通过反射实现结构体转成JSON数据](http://blog.cyeam.com/golang/2014/08/11/go_json)

反射的作用有什么呢?

1. 判断数据类型, 类似于python的`type()`函数, js的`instanceof`操作符等
2. 以字符串为键, 获取结构体中的指定成员, 类似于`obj['item']`, 或`obj.__dict__['item']`

## 1. Type与Value

`reflect`包定义了两种数据类型: `Type`与`Value`.

- `Type`表示目标对象所属的类型
- `Value`表示目标对象的值

以如下一个简单示例来解释.

```go
str := "this is string"
fmt.Println(reflect.TypeOf(str))        // string
fmt.Println(reflect.ValueOf(str))       // this is string

var x float64 = 3.4
fmt.Println(reflect.TypeOf(x))          // float64
fmt.Println(reflect.ValueOf(x))         // 3.4
```

当然, 自定义结构体类型也是可以的识别的. 定义如下结构体, 本文之后的测试代码都基于这个`User`结构体.

```go
type User struct {
	Name string
}
func (this *User) GetName() string {
	return this.Name
}
func (this *User) SetName(name string) {
	this.Name = name
}
```

```go
user := &User{Name: "general"}
theType := reflect.TypeOf(user)
theVal := reflect.ValueOf(user)
fmt.Println(theType)        // *main.User
fmt.Printf("%T\n", theType) // 类型: *reflect.rtype
fmt.Println(theVal)         // &{general}
fmt.Printf("%T\n", theVal)  // 类型: reflect.Value
```

## 2. Type类型的可用方法

### `Elem()`

这个方法可以返回目标对象的引用类型(而不是指针类型), 举个栗子.

```go
user1 := &User{Name: "general"}                 // 指针
user2 := User{Name: "jiangming"}                // 引用

fmt.Println(reflect.TypeOf(user1))              // *main.User 指针类型
fmt.Println(reflect.TypeOf(user2))              // main.User 引用类型

fmt.Println(reflect.TypeOf(user1).Elem())       // main.User 这下也变成了引用类型
```

之所以先介绍这个方法, 是因为`reflect`有一些方法对**指针**和**引用**类型的执行结果是不同的.

> 注意: 已经是引用类型的`Type`对象再次调用`Elem()`会出错的(golang又没有异常捕获机制, 感觉这很坑啊)

```go
	userType1 := reflect.TypeOf(user1)
	userType2 := reflect.TypeOf(user2)
```

### `NumField()`与`NumMethod()`

这两个方法分别可以得到目标对象所属类型的**成员属性个数**和**成员方法的个数**.

但是!!!要注意!!!

只有引用类型才拥有成员属性, 而成员方法则是分别要看`receiver`的类型是`T`还是`*T`.

```go
fmt.Println(userType2.NumField())       // 1 表示1个属性, Name
fmt.Println(userType1.NumMethod())      // 2 表示*User作receiver拥有2个方法
fmt.Println(userType2.NumMethod())      // 0 表示User作receiver没有方法
```

### `Field()`与`FieldByName()`

- `Field func(i int) StructField`: 接受一个整型变量作为参数`i`, 其值为按照`NumField()`返回的属性数量范围内的索引, 即按照属性定义的顺序的序号, ta返回`StructField`结构体, 表示目标字段的相应信息.
- `FieldByName()`可以直接通过属性名得到该属性成员的相关信息.

```go
// {Name:Name PkgPath: Type:string Tag: Offset:0 Index:[0] Anonymous:false}
// Name, 即属性名
fmt.Printf("%+v\n", userType2.Field(0))
field, ok := userType2.FieldByName("Name")
if ok {
	// {Name:Name PkgPath: Type:string Tag: Offset:0 Index:[0] Anonymous:false}
	fmt.Printf("%+v\n", field)
}
```

只要知道`StructField`结构体中`Name`表示字段名, `Type`表示字段类型就行了, 其他的也没必要知道. (以后也许会需要用到`Offset`字段, 与字段在结构体中的顺序有关).

## 3. Value类型的可用方法

`Type`类型的相关方法只是为了获取目标类型的成员属性(或方法)的名称, 类型等信息, 而如果想要通过字符串变量获取指定成员字段的值的话, 就需要使用`Value`类型了.

与`Type`类型中的`Elem()`方法用法相似, `Value`类型也可以使用.

```go
user1 := &User{Name: "general"} // 指针
// user2 := User{Name: "jiangming"}                // 引用
theVal := reflect.ValueOf(user1).Elem()
fmt.Println(theVal.FieldByName("Name"))				// 输出general
```
