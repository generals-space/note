# golang-big大数运算

参考文章

1. [9.4 精密计算和 big 包](http://wiki.jikexueyuan.com/project/the-way-to-go/09.4.html)

现代操作系统中, 整型的最大位数为64位, 相当于8个`byte`.

在进行`lorawan`开发时, 设备key为`[16]byte`的16进制数值, 已经超出了系统所能表示的极限, 所以找到了这个库.

按照参考文章1中所说, big库其实是用于高精度计算的, 目前没有找到相关示例. 这篇文章简单介绍大数计算.

`big`库在`math/big`包中.

## 1. 定义big int

首先定义一个简单的大数.

```go
	simpleBig := big.NewInt(1000)
	fmt.Println(simpleBig)          // 1000
	fmt.Printf("%T\n", simpleBig)   // *big.Int
```

...好像没什么特别的啊?

像上面提到的, 如果我想表示一个值为`FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF`的16进制整型, 怎么办? 如下

```go
	var a int64
	// constant 340282366920938463463374607431768211455 overflows int64
	a = 340282366920938463463374607431768211455
```

单纯使用字符串表示数值是不行的, 因为之后还会涉及计算. `big`库提供了解决办法. 因为大的数据可以通过字符串, 或是数组表示, `big`允许从这些结构中读入, 然后构造成大数类型. 如下

```go
    // 我们需要得到一个big int的结构.
	// 以下两种初始化方式都可以
	baseBig := new(big.Int)
    // baseBig := big.NewInt(0)
    
    // SetString和SetBytes可以为big int赋值
	big1, _ := baseBig.SetString("340282366920938463463374607431768211455", 10)
	fmt.Println(big1)
	big2, _ := baseBig.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
    fmt.Println(big2)
    	big3 := baseBig.SetBytes([]byte{
		0xFF, 0xFF, 0xFF, 0xFF, 
		0xFF, 0xFF, 0xFF, 0xFF, 
		0xFF, 0xFF, 0xFF, 0xFF, 
		0xFF, 0xFF, 0xFF, 0xFF,
	})
	fmt.Println(big3)
```

> 注意: 上述代码中`big1`, `big2`, `big3`的地址相同, 因为值也是相同的, 只不过为了区分不同的`Set`方式才定义为不同的变量.

## 2. 计算

big int类型的数据无法直接通过`+`, `-`, `*`, `/`进行计算, 而是提供类似于`time.Duration`的计算方式.

计算的方法包括

- `Add func(x, y *big.Int) *big.Int`: 加

- `Sub func(x, y *big.Int) *big.Int`: 减

- `Mul func(x, y *big.Int) *big.Int`: 乘

- `Div func(x, y *big.Int) *big.Int`: 除

...等等

这些方法都把计算双方都作为参数传入, 本来这样的安排应该把方法作为库的静态方法暴露出来才对, 但是golang却把ta们放到了`big.Int`的成员方法中...所以在使用这个时, 你要先创建一个`big.Int`对象.

```go
	baseBig := big.NewInt(0)
	bigInt, _ := baseBig.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
    fmt.Println(bigInt) // 340282366920938463463374607431768211455
    // 减1操作
	result1 := big.NewInt(0).Sub(bigInt, big.NewInt(1))
    fmt.Println(result1) // 340282366920938463463374607431768211454
    result2 := big.NewInt(0).Sub(bigInt, big.NewInt(10))
	fmt.Println(result2) // 340282366920938463463374607431768211445
```

## 3. 结果输出

`big.Int`类型的值可以转换为其他类型, 比如`bytes(数组)`, `int64(不能超过上限)`, 或是`string`.

```go
	baseBig := big.NewInt(0)
	bigInt, _ := baseBig.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
	fmt.Println(bigInt)             // 340282366920938463463374607431768211455
	result := big.NewInt(0).Sub(bigInt, big.NewInt(1))
	fmt.Println(result)             // 340282366920938463463374607431768211454
	fmt.Println(result.String())    // 340282366920938463463374607431768211454
	fmt.Println(result.Bytes())     // [255 255 255 255 255 255 255 255 255 255 255 255 255 255 255 254]

	fmt.Println(result.Int64())     // -2
```

## 4. 比较

数值比较可以使用`Cmp()`方法.

```go
	baseBig1 := big.NewInt(0)
	big1, _ := baseBig1.SetString("340282366920938463463374607431768211455", 10)

	baseBig2 := big.NewInt(0)
	big2, _ := baseBig2.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)

	// 本来两者是相等的
    fmt.Println(big1.Cmp(big2)) // big1 == big2 0

	big1 = big.NewInt(0).Sub(big1, big.NewInt(1))

	fmt.Println(big1.Cmp(big2)) // big1 < big2 -1
	fmt.Println(big2.Cmp(big1))	// big2 > big1 1
```