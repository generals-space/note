# golang-interface接口与receiver对象

参考文章

1. [go语言面试题 第10题](https://my.oschina.net/qiangmzsx/blog/1478739)

2. [请教类型T与*T的方法集的问题](https://www.golangtc.com/t/54d5c7fb421aa93482000089)

```go
package main

import (
	"fmt"
)

// People ...
type People interface {
	Speak(string)
}

// Student ...
type Student struct{}

// Speak ...
func (stu *Student) Speak(think string) {
	fmt.Println(think)
}

func main() {
	// cannot use Student literal (type Student) as type People in assignment:
	// Student does not implement People (Speak method has pointer receiver)
	var peo People = Student{}
	think := "hello world"
	peo.Speak(think)
}

```

上面的代码中有错误, 编译通不过, IDE也会用红色波浪线提示. 实例化的`Student{}`对象并没有实现`People`接口, 因为`Speak()`方法是挂在`*Student`这个receiver上的, `Student{}`对象不能直接使用.

如果是`var peo People = &Student{}`就可以运行, 另外如果`Speak()`方法挂在`Student`上, `Student{}`和`&Student{}`都可以编译通过.

按照参考文章1中的总结, **golang的方法集仅仅影响接口实现和方法表达式转化, 与通过实例或者指针调用方法无关**. 

关于这一句总结, 通过实例/指针调用方法基本是无影响的(我们平时也是`T`和`*T`都可以混用), `T`与`*T`对接口实现的影响从上例已经可以看到, 不过 **方法表达式转化**还是一头雾水. 暂不考虑<???>