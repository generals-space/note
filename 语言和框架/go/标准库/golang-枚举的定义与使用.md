# golang-枚举的定义与使用

参考文章

1. [go 枚举](https://studygolang.com/articles/5386)

2. [golang枚举类型 - iota用法拾遗](https://studygolang.com/articles/7897)

## 1. 认识

参考文章2中第一节对枚举类型的解释比较容易理解. 

golang中没有枚举类型.

枚举类型在golang中既不是基本类型也不是复合类型, 客观地说, 它应该是一组某种类型的常量集合.

```go
package main

import "fmt"

type State int

// iota 初始化后会自动递增
const (
	Running    State = iota // value --> 0
	Stopped                 // value --> 1
	Rebooting               // value --> 2
	Terminated              // value --> 3
)

func main() {
	state := Stopped
	fmt.Println(state)             // 1
	fmt.Printf("%T\n", Terminated) // main.State
}
```

```go
package main
import "fmt"
type State int

// iota 初始化后会自动递增
const (
    Running State = iota // value --> 0
    Stopped              // value --> 1
    Rebooting            // value --> 2
    Terminated           // value --> 3
)

// 重载String函数
func (this State) String() string {
    switch this {
        case Running:
        return "Running"
    case Stopped:
        return "Stopped"
    default:
        return "Unknow"
    }
}

func main() {
    state1 := Stopped
    // 输出 state Running
    // 没有重载String函数的情况下则输出 state 0
    fmt.Println(state1)  // 0
    var state2 int
    state2 = 3
    fmt.Printf("%T", Terminated)  // 0
}
```

`Running State = iota`定义了之后的枚举类型都是`State`并递增, 而`State`是`int`的别名. 在不指定`State`类型时, 枚举类型的默认取值类型也是`int`.

在使用时, 可以直接使用`Running`, `Stopped`这种