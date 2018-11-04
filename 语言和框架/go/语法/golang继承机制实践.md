# golang的继承机制

参考文章

1. [Golang OOP、继承、组合、接口](https://www.cnblogs.com/jasonxuli/p/6836399.html)

流传很广的OOP的三要素是：封装、继承、多态。

对象：可以看做是一些特征的集合，这些特征主要由 属性 和 方法 来体现。
封装：划定了对象的边界，也就是定义了对象。
继承：表明了子对象和父对象之间的关系，子对象是对父对象的扩展，实际上，子对象“是”父对象。相当于说“码农是人”。从特征的集合这个意义上说，子对象包含父对象，父对象有的公共特征，子对象全都有。
多态：根据继承的含义，子对象在特性上全包围了父对象，因此，在需要父对象的时候，子对象可以替代父对象。

在golang中根本没有这些概念, 虽然可以通过内嵌`struct`使用父类的方法, 但是无法实现**子类实例是父类实例**的概念. 因为在golang中这种所谓的继承只是组合而已.

```go
package main

import (
	"log"
)

// Parent ...
type Parent struct {
	Name string
}

// Greet ...
func (p *Parent) Greet() {
	log.Println("Parent name is ", p.Name)
}

// Child ...
type Child struct {
	Parent
	Name string
}

// Greet ...
// 子类同名方法会覆盖.
// 在子类没有同名方法时, 子类实例可以调用父类的方法, 但变量不通用.
func (c *Child) Greet() {
	log.Println("Child name is ", c.Name)
}

func main() {
	child := &Child{
		Parent: Parent{Name: "parentName"},
		Name:   "childName",
	}
	child.Greet()
	child.Parent.Greet()
}
```

上述代码的执行结果如下

```
2018/11/04 10:44:14 Child name is  childName
2018/11/04 10:44:14 Parent name is  parentName
```

当`Child`结构没有重新定义`Greet()`方法时, `child.Greet()`输出的是`Parent name is  parentName`...

总结一下

1. 父子类的成员变量不会相互替换, 各自调用各自的, 子类想要调用父类的成员, 只能使用`child.Parent.Name`这种方式.

2. 子类倒是可以直接调用父类方法...真的是单纯调用父类方法, 因为父子类间成员变量不互通, 无法实现像其他语言中在父类中定义方法, 然后在子类调用时可以处理子类成员变量的功能...

3. 子类重新定义父类的重名方法后, 就会将其覆盖, 要调用父类方法可以使用`child.Parent.Greet()`完成.