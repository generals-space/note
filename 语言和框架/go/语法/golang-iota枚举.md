# golang-iota

参考文章

1. [golang 使用 iota](http://www.cnblogs.com/ghj1976/p/4224346.html)

`iota`是golang语言的常量计数器, 只能在**常量表达式(用const声明变量时, 不能用var哦)**中使用. 

`iota`在`const`关键字出现时将被重置为0(`const`内部的第一行之前), `const`中每新增一行常量声明将使`iota`计数一次(`iota`可理解为`const`语句块中的行索引). 

使用`iota`能简化定义, 在定义**枚举**时很有用. 

## 1. iota只能在常量的表达式中使用. 

```go
fmt.Println(iota)
```

编译错误: `undefined: iota`

## 2. 每次`const`出现时, 都会让`iota`初始化为0(或其他初始值).

注意以下两种声明`iota`变量的区别.

```go
const a = iota // a = 0(int类型)
const (
	b = iota // b = 0
	c        // c = 1
)
```

```go
const (
	a = iota // a = 0
	b = iota // b = 1
	c        // c = 2
)
```

## 3. 自定义类型

> 不指定枚举类型时, `iota`枚举列表默认为`int`类型.

自增长常量经常包含一个自定义枚举类型, 允许你依靠编译器完成自增设置. 

```go
type Stereotype int
const ( 
    TypicalNoob Stereotype = iota // 0 
    TypicalHipster                // 1 
    TypicalUnixWizard             // 2 
    TypicalStartupFounder         // 3 
)
```

## 4. 可跳过的值

设想你在处理消费者的音频输出. 音频可能无没有任何输出, 或者它可能是单声道, 立体声, 或是环绕立体声的. 

这可能有些潜在的逻辑定义: 

0: 没有任何输出
1: 单声道
2: 立体声

值是由通道的数量提供. 所以你给 Dolby 5.1 环绕立体声什么值? 

一方面, 它有6个通道输出, 但是另一方面, 仅仅 5 个通道是全带宽通道(因此 5.1 称号 - 其中 .1 表示的是低频效果通道). 

不管怎样, 我们不想简单的增加到 3. 

我们可以使用**下划线**跳过不想要的值. 

```go
type AudioOutput int
const ( 
    OutMute AudioOutput = iota // 0 
    OutMono                    // 1 
    OutStereo                  // 2 
    _ 
    _ 
    OutSurround                // 5 
)
```

## 5. 位掩码表达式

```go
type Allergen int
const ( 
    IgEggs Allergen = 1 << iota // 1 << 0 00000001 
    IgChocolate                 // 1 << 1 00000010 
    IgNuts                      // 1 << 2 00000100 
    IgStrawberries              // 1 << 3 00001000 
    IgShellfish                 // 1 << 4 00010000 
)
```

> 注: `Allergen`: 过敏原.

这个工作是因为当你在一个`const`组中仅仅有一个标识符在一行的时候, 它将使用增长的`iota`取得前面的表达式并且再运用它. 在 Go 语言的 spec 中,  这就是所谓的隐性重复最后一个非空的表达式列表. 

如果你对鸡蛋, 巧克力和海鲜过敏, 把这些 bits 翻转到 "on" 的位置(从左到右映射 bits). 然后你将得到一个 bit 值 00010011, 它对应十进制的 19. 

```go
fmt.Println(IgEggs | IgChocolate | IgShellfish) // 输出19
```

## 6. 定义数量级

```go
type ByteSize float64
const (
    _           = iota             // 忽略第一个值
    KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
    MB                             // 1 << (10*2)
    GB                             // 1 << (10*3)
    TB                             // 1 << (10*4)
    PB                             // 1 << (10*5)
    EB                             // 1 << (10*6)
    ZB                             // 1 << (10*7)
    YB                             // 1 << (10*8)
)
```

## 7. 定义在一行的情况

```go
const (
    Apple, Banana = iota + 1, iota + 2
    Cherimoya, Durian
    Elderberry, Fig
)
```

`iota`在下一行增长, 而不是立即取得它的引用. 

```go
	fmt.Println(Apple)      // 1
	fmt.Println(Banana)     // 2
	fmt.Println(Cherimoya)  // 2
	fmt.Println(Durian)     // 3
	fmt.Println(Elderberry) // 3
	fmt.Println(Fig)        // 4
```

> ...没看懂, 别没事找事了.

## 8. 中间插队

注意: 比较如下声明变量的方式

```go
const (
	i = iota // 0
	j = 3.14 // 3.14
	k = iota // 2    注意: 这里必须用iota表示重新"接上上面的iota"
	l        // 3
)
```

```go
const (
	i = iota // 0
	j = 3.14 // 3.14
	k        // 3.14
	l        // 3.14
)
```

...还是觉得有点不对劲, 为什么会这样?

```go
const (
	j = 3.14 // 3.14
	k        // 3.14
	l        // 3.14
)
```

`const`声明默认继承前面定义的值.


------

我在sync.Mutex的源码中发现了一个最初让我比较疑惑的iota声明.

```go
	const (
		mutexLocked      = 1 << iota // 1 1 << 0 0001
		mutexWoken                   // 2 1 << 1 0010
		mutexStarving                // 4 1 << 2 0100
		mutexWaiterShift = iota      // 3
	)
```

刚开始我怎么也想不明白为什么mutexWaiterShift的值是3, 后来发现这其实是示例5与示例8的结合. 

前三个数值是常规的位掩码定义将1分别左移0, 1, 2位, 此时iota值为2, 接着mutexWaiterShift出现, 但ta的值直接为iota, 而不再是1 << iota, 所以ta直接取了iota的下一个值3.
