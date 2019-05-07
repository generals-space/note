# Golang OOP, 继承, 组合, 接口(转)

原文链接

[Golang OOP、继承、组合、接口](https://www.cnblogs.com/jasonxuli/p/6836399.html)

## 1. 传统 OOP 概念

OOP（面向对象编程）是对真实世界的一种抽象思维方式, 可以在更高的层次上对所涉及到的实体和实体之间的关系进行更好的管理. 
 
流传很广的OOP的三要素是: 封装、继承、多态. 

- 对象: 可以看做是一些特征的集合, 这些特征主要由 属性 和 方法 来体现. 
- 封装: 划定了对象的边界, 也就是定义了对象. 
- 继承: 表明了子对象和父对象之间的关系, 子对象是对父对象的扩展, 实际上, 子对象“是”父对象. 相当于说“码农是人”. 从特征的集合这个意义上说, 子对象包含父对象, 父对象有的公共特征, 子对象全都有. 
- 多态: 根据继承的含义, 子对象在特性上全包围了父对象, 因此, 在需要父对象的时候, 子对象可以替代父对象. 
 
传统的 OOP 语言, 例如 Java, C++, C#, 对OOP 的实现也各不相同. 以 Java 为例:  Java 支持 extends, 也支持 interface. 这是两种不同的抽象方式. 

extends 就是继承, A extends B, 表明 A 是 B 的一种, 是概念上的抽象关系. 

```java
Class Human {
    name:string
    age:int
    function eat(){}
    function speak(){}
}
 
Class Man extends Human {
    function fish(){}
    function drink(){}
}
```

## 2. Golang 的 OOP

回到 Golang. Golang 并没有 extends, 它类似的方式是 `Embedding`. 这种方式并不能实现 `is-a` 这种定义上的抽象关系, 因此 Golang 并没有传统意义上的多态. 

注意下面代码中的注释, 把 Student 当做 Human 会报错. 

```go
package main
 
import "fmt"
 
func main(){
    var h Human
 
    s := Student{Grade: 1, Major: "English", Human: Human{Name: "Jason", Age: 12, Being: Being{IsLive: true}}}
    fmt.Println("student:", s)
    fmt.Println("student:", s.Name, ", isLive:", s.IsLive, ", age:", s.Age, ", grade:", s.Grade, ", major:", s.Major)
 
    //h = s // cannot use s (type Student) as type Human in assignment
    fmt.Println(h)
 
    //Heal(s) // cannot use s (type Student) as type Being in argument to Heal
    Heal(s.Human.Being) // true
 
    s.Drink()
    s.Eat()
}
 
 
type Car struct {
    Color string
    SeatCount int
}
 
type Being struct {
    IsLive bool
}
 
type Human struct {
    Being
    Name string
    Age int
}
 
func (h Human) Eat(){
    fmt.Println("human eating...")
    h.Drink()
}
 
func (h Human) Drink(){
    fmt.Println("human drinking...")
}
 
func (h Human) Move(){
    fmt.Println("human moving...")
}
 
type Student struct {
    Human
    Grade int
    Major string
}
 
func (s Student) Drink(){
    fmt.Println("student drinking...")
}
 
type Teacher struct {
    Human
    School string
    Major string
    Grade int
    Salary int
}
 
func (s Teacher) Drink(){
    fmt.Println("teacher drinking...")
}
 
 
type IEat interface {
    Eat()
}
 
type IMove interface {
    Move()
}
 
type IDrink interface {
    Drink()
}
 
 
func Heal(b Being){
    fmt.Println(b.IsLive)
}
```

输出结果

```go
student: {{{true} Jason 12} 1 English}
student: Jason , isLive: true , age: 12 , grade: 1 , major: English
{{false}  0}
true
student drinking...
human eating...
human drinking...
```

这里有一点需要注意, `Student` 实现了 `Drink` 方法, 覆盖了 `Human` 的 `Drink`, 但是没有实现 `Eat` 方法. 因此, `Student` 在调用 `Eat` 方法时, 调用的是 `Human` 的 `Eat()`；而 `Human` 的 `Eat()` 调用了 `Human` 的 `Drink()`, 于是我们看到结果中输出的是 human drinking... . 这既不同于 Java 类语言的行为, 也不同于 prototype 链式继承的行为, Golang 叫做 `Embedding`, 这像是一种**寄生**关系: `Human` 寄生在 `Student` 中, 但仍保持一定程度的独立. 

## 3. Golang 的接口

我们从接口产生的原因来考虑. 

代码处理的是各种数据. 对于强类型语言来说, 非常希望一批数据都是单一类型的, 这样它们的行为完全一致. 但世界是复杂的, 很多时候数据可能包含不同的类型, 却有一个或多个共同点. 这些共同点就是抽象的基础. 单一继承关系解决了 is-a 也就是定义问题, 因此可以把子类当做父类来对待. 但对于父类不同但又具有某些共同行为的数据, 单一继承就不能解决了. 单一继承构造的是树状结构, 而现实世界中更常见的是网状结构. 

于是有了接口. 接口是在某一个方面的抽象, 也可以看做具有某些相同行为的事物的标签. 

但不同于继承, 接口是松散的结构, 它不和定义绑定. 从这一点上来说, Duck Type 相比传统的 extends 是更加松耦合的方式, 可以同时从多个维度对数据进行抽象, 找出它们的共同点, 使用同一套逻辑来处理. 

Java 中的接口方式是先声明后实现的强制模式, 比如, 你要告诉大家你会英语, 并且要会听说读写, 你才具有英语这项技能. 

```java
interface IEnglishSpeaker {
    ListenEnglish()
    ReadEnglish()
    SpeakEnglish()
    WriteEnglish()
}
```

Golang 不同, 你不需要声明你会英语, 只要你会听说读写了, 你就会英语了. 也就是实现决定了概念: 如果一个人在学校（有School、Grade、Class 这些属性）, 还会学习（有Study()方法）, 那么这个人就是个学生. 

Duck Type 更符合人类对现实世界的认知过程: 我们总是通过认识不同的个体来进行总结归纳, 然后抽象出概念和定义. 这基本上就是在软件开发的前期工作, 抽象建模. 

相比较而言,  Java 的方式是先定义了关系（接口）, 然后去实现, 这更像是从上帝视角先规划概念产生定义, 然后进行造物. 
 
因为 interface 和 object 之间的松耦合, Golang 有 `type assertion` 这样的方式来判断一个接口是不是某个类型: 
`value, b := interface.(Type)`, value 是 Type 的默认实例；b 是 bool 类型, 表明断言是否成立. 

```go
// 接上面的例子
 
v1, b := interface{}(s).(Car)
fmt.Println(v1, b)
 
v2, b := interface{}(s).(Being)
fmt.Println(v2, b)
 
v3, b := interface{}(s).(Human)
fmt.Println(v3, b)
 
v4, b := interface{}(s).(Student)
fmt.Println(v4, b)
 
v5, b := interface{}(s).(IDrink)
fmt.Println(v5, b)
 
v6, b := interface{}(s).(IEat)
fmt.Println(v6, b)
 
v7, b := interface{}(s).(IMove)
fmt.Println(v7, b)
 
v8, b := interface{}(s).(int)
fmt.Println(v8, b)
```

输出结果: 

```
{ 0} false
{false} false
{{false}  0} false
{{{true} Jason 12}  1 English} true
{{{true} Jason 12}  1 English} true
{{{true} Jason 12}  1 English} true
<nil> false
false
```

上面的代码中, 使用空接口 interface{} 对 s 进行了类型转换, 因为 s 是 struct, 不是 interface, 而类型断言表达式要求点号左边必须为接口. 

常用的方式应该是类似泛型的使用方式: 

```go
s1 := Student{Grade: 1, Major: "English", Human: Human{Name: "Jason", Age: 12, Being: Being{IsLive: true}}}
s2 := Student{Grade: 1, Major: "English", Human: Human{Name: "Tom", Age: 13, Being: Being{IsLive: true}}}
s3 := Student{Grade: 1, Major: "English", Human: Human{Name: "Mike", Age: 14, Being: Being{IsLive: true}}}
t1 := Teacher{Grade: 1, Major: "English", Salary: 2000, Human: Human{Name: "Michael", Age: 34, Being: Being{IsLive: true}}}
t2 := Teacher{Grade: 1, Major: "English", Salary: 3000, Human: Human{Name: "Tony", Age: 31, Being: Being{IsLive: true}}}
t3 := Teacher{Grade: 1, Major: "English", Salary: 4000, Human: Human{Name: "Ivy", Age: 40, Being: Being{IsLive: true}}}
drinkers := []IDrink{s1, s2, s3, t1, t2, t3}
 
for _, v := range drinkers {
    switch t := v.(type) {
    case Student:
        fmt.Println(t.Name, "is a Student, he/she needs more homework.")
    case Teacher:
        fmt.Println(t.Name, "is a Teacher, he/she needs more jobs.")
    default:
        fmt.Println("Invalid Human being:", t)
    }
}
```

输出结果: 

```
Jason is a Student, he/she needs more homework.
Tom is a Student, he/she needs more homework.
Mike is a Student, he/she needs more homework.
Michael is a Teacher, he/she needs more jobs.
Tony is a Teacher, he/she needs more jobs.
Ivy is a Teacher, he/she needs more jobs.
```

这段代码中使用了 Type Switch, 这种 switch 判断的目标是类型. 

## 4. Golang: 接口为重

了解了 Golang 的 OOP 相关的基本知识后, 难免会有疑问, 为什么 Golang 要用这种“非主流”的方式呢? 
 
Java 之父 James Gosling 在某次会议上有过这样一次问答: 

``` 
I once attended a Java user group meeting where James Gosling (Java’s inventor) was the featured speaker.
During the memorable Q&A session, someone asked him: “If you could do Java over again, what would you change?”  
“I’d leave out classes,” he replied. After the laughter died down, he explained that the real problem wasn’t classes per se, but rather implementation inheritance (the extends relationship). Interface inheritance (the implements relationship) is preferable. You should avoid implementation inheritance whenever possible.
```

大意是: 
问: 如果你重新做 Java, 有什么是你想改变的? 

答: 我会把类（class）丢掉. 真正的问题不在于类本身, 而在于基于实现的继承（the extends relationship）. 基于接口的继承（the implements relationship）是更好的选择, 你应该在任何可能的时候避免使用实现继承. 

 
我的理解是: 实现之间应该少用继承式的强关联, 多用接口这种弱关联. 接口已经可以在很多方面替代继承的作用, 比如多态和泛型. 而且接口的关系松散、随意, 可以有更高的自由度、更多的抽象角度. 
 
以继承为特点的 OOP 只是编程世界的一种抽象方式, 在 Golang 的世界里没有继承, 只有组合和接口, 这看起来更符合 Gosling 的设想. 借用那位老人的话: 黑猫白猫, 捉住老鼠就是好猫. 让我来继续探索吧. 

注: 刚刚学习 Golang 不久, 后面可能会发现也许某些理解是错误的. 随时修正. 