# golang-sync.atomic

参考文章

1. [golang语言中sync/atomic包的学习与使用](https://www.cnblogs.com/jkko123/p/7220654.html)
    - 简单示例
2. [GO安全并发之无锁原子操作](https://www.cnblogs.com/sunsky303/p/7453344.html)
    - 《Go并发编程实战》的样章
    - atomic包各种函数原型及参数解释
    - `AddUint`/`AddUint64`函数使用补码完成"减"操作

查看`sync.atomic`源码, 你会发现在这个包内只是定义了`AddXXX`, `SwapXXX`, `CompareAndSwapXXX`等函数类型, 并没有实现. 具体实现在`runtime/internal/atomic`包内, 使用了`asm`汇编代码.

我对原子操作的理解是, 可以实现与加锁同样的功能, 但是效率更高. 因为语言提供的锁是操作系统实现的, 而golang的原子操作是汇编代码, 直接操作硬件, 因此速度更快.

------

`atomic`包提供了底层的原子级内存操作

类型共有六种：int32, int64, uint32, uint64, uintptr, unsafe.Pinter

操作共五种：增减，比较并交换， 载入， 存储，交换

1. 增减操作

函数名以`Add`为前缀，加具体类型名, 如`AddInt32()`/`AddInt64()`. 参数一，是指针类型; 参数二，与参数一类型总是相同.

```go
	var A int32
	fmt.Println("A: ", A) // A: 0
	newA := atomic.AddInt32(&A, 3)
	fmt.Println("newA: ", newA) // newA: 3
	newA = atomic.AddInt32(&A, -2)
	fmt.Println("newA: ", newA) // newA: 1

```

1. 比较并交换

CAS(Compare And Swap)比较并交换操作, 函数名以`CompareAndSwap`为前缀，并具体类型名, 如`CompareAndSwapInt32()`/`CompareAndSwapInt64()`.

函数会先判断参数一指向的值与参数二是否相等，如果相等，则用参数三替换参数一的值。最后返回是否替换成功.

```go
	var B int32
	fmt.Println("B: ", B) // b:  0
	atomic.CompareAndSwapInt32(&B, 0, 3)
    fmt.Println("B: ", B) // b:  3

```

3. 载入操作

当我们对某个变量进行读取操作时，可能该变量正在被其他操作改变，或许我们读取的是被修改了一半的数据。所以我们通过Load这类函数来确保我们正确的读取.

函数名以`Load`为前缀，加具体类型名, 如`LoadInt32()`/`LoadInt64()`.

```go
	var C int32
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// 并发协程为变量C加1.
		// 可以适当调高并发数量, 很有可能出现修改失败的情况.
		go func() {
			defer wg.Done()
			tmp := atomic.LoadInt32(&C)
			if !atomic.CompareAndSwapInt32(&C, tmp, (tmp + 1)) {
				fmt.Println("C 修改失败")
			}
		}()
	}
	wg.Wait()
	//C的值有可能不等于100，频繁修改变量值情况下，CAS操作有可能不成功。
	fmt.Println("C: ", C)

```

上述代码同时包含了`Load`与`CAS`操作, 经过实验, 的确会出现协程内部修改C的值失败的情况.

...当然, 使用`atomic`操作整型变量的增减虽然有修改失败的情况, 但是机率比不加任何保护措施的并发写操作失败的机率要小很多(md这仍然不是使用无锁原子操作的理由好吗?).

引用参考文章2中的说法:

> 与我们前面讲到的锁相比，CAS操作有明显的不同。
> 
> CAS总是假设被操作值未曾被改变（即与旧值相等），并一旦确认这个假设的真实性就立即进行值替换。
> 
> 而使用锁则是更加谨慎的做法。我们总是先假设会有并发的操作要修改被操作值，并使用锁将相关操作放入临界区中加以保护。
> 
> 我们可以说，使用锁的做法趋于悲观，而CAS操作的做法则更加乐观(前者一定要确认成功, 而CAS不做这样的保证...?)。
> 
> CAS操作的优势是，可以在不形成临界区和创建互斥量的情况下完成并发安全的值替换操作。这可以大大的减少同步对程序性能的损耗。
> 
> 当然，CAS操作也有劣势。在被操作值被频繁变更的情况下，CAS操作并不那么容易成功。有些时候，我们可能不得不利用for循环以进行多次尝试。

之后参考文章2给出了一个使用`for`循环不断尝试修改目标变量的代码, 并且提到这种方式与"自旋锁"的行为相似.

4. 存储操作

与载入函数相对应，提供原子的存储函数. 函数名以`Store`为前缀，加具体类型名, 如`StoreInt32()`与`StoreInt64()`.

存储某个值时，任何CPU都不会都该值进行读或写操作, 存储操作总会成功，它不关心旧值是什么，与CAS不同.

```go
	var D int32
	fmt.Println("D: ", D) // D: 0
	atomic.StoreInt32(&D, 666)
	fmt.Println("D: ", D) // D: 666

```

5. 交换操作

直接设置新值，返回旧值，与CAS不同，它不关心旧值。函数名以`Swap`为前缀，加具体类型名, 如`SwapInt32()`/`SwapInt64()`.

```go
	var E int32
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// 并发协程为变量C加1.
		// 可以适当调高并发数量, 很有可能出现修改失败的情况.
		go func() {
			defer wg.Done()
			tmp := atomic.LoadInt32(&E)
			_ = atomic.SwapInt32(&E, (tmp + 1))
			// fmt.Println("oldE: ", oldE)
		}()
	}
	wg.Wait()
	fmt.Println("E: ", E) // E: 982
```