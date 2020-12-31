# golang的继承机制

参考文章

1. [Golang OOP、继承、组合、接口](https://www.cnblogs.com/jasonxuli/p/6836399.html)

流传很广的OOP的三要素是：封装、继承、多态。

在golang中根本没有这些概念, 虽然可以通过内嵌`struct`使用父类的方法, 但是无法实现 **子类实例是父类实例**的概念. 因为在golang中这种所谓的继承只是 **组合**而已.

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

当`Child`结构没有重新定义`Greet()`方法时, `child.Greet()`输出的是`Parent name is parentName`...

总结一下

1. 父子类的成员变量不会相互替换, 各自调用各自的, 子类想要调用父类的成员, 只能使用`child.Parent.Name`这种方式.
2. 子类没有重新定义父类的方法时, 倒是可以直接调用父类方法...真的是单纯调用父类方法, 因为父子类间成员变量不互通, 无法实现像其他语言中在父类中定义方法, 然后在子类调用时可以处理子类成员变量的功能...
3. 子类重新定义父类的重名方法后, 就会将其覆盖, 要调用父类方法可以使用`child.Parent.Greet()`完成.

所以看一下下面这个示例.

```go
package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB() // 这里只能调用本身(People)的ShowB()方法.
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
```

输出为

```
showA
showB
```

## 多重"继承"

考虑如果父类有多个子类成员, 且每个子类成员拥有同名方法时, 在父类实例上调用这个方法, 会怎样?

```go
package main

import (
	"log"
)

// StepParent ...
type StepParent struct {
	Name string
}

// Greet ...
func (sp *StepParent) Greet() {
	log.Println("Step Parent name is ", sp.Name)
}

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
	StepParent
	Name string
}

/*
// Greet ...
// 子类同名方法会覆盖.
// 在子类没有同名方法时, 子类实例可以调用父类的方法, 但变量不通用.
func (c *Child) Greet() {
	log.Println("Child name is ", c.Name)
}
*/

func main() {
	child := &Child{
		Parent:     Parent{Name: "parentName"},
		StepParent: StepParent{Name: "stepParentName"},
		Name:       "childName",
	}
	child.Greet()
	child.Parent.Greet()
	child.StepParent.Greet()
}
```

但是编译会报错, 报错行为`child.Greet()`

```
# command-line-arguments
./assemble.go:49:7: ambiguous selector child.Greet
```

因为在父类`Child`没有定义`Greet()`方法时, 直接调用`child.Greet()`会让编译器疑惑, 不知道该调用`Parent`成员的`Greet()`还是`StepParent`的`Greet()`方法. 不像python, 继承时谁在前面调用谁的.
