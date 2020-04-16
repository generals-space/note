# golang-context认识与使用

参考文章

1. [golang中context包解读](http://www.01happy.com/golang-context-reading/)

2. [Go语言并发模型：使用 context](https://segmentfault.com/a/1190000006744213)

3. [Golang context初探](https://www.jianshu.com/p/0dc7596ba90a)
	- A -> B -> C调用实例
	- context的使用规范

## 1. 认识

golang中的创建一个新的`goroutine`, 并不会返回像其他语言创建进程时返回pid, 也不像创建线程时能得到一个线程对象的引用. 

`goroutine`通过`go`关键字直接执行, 所以我们无法从外部杀死某个`goroutine`. 之前我们用 `channel ＋ select`的方式, 来解决这个问题, 但是有些场景实现起来比较麻烦. 例如由一个请求衍生出的各个`goroutine`之间需要满足一定的约束关系, 以实现一些诸如有效期, 中止routine树, 传递请求全局变量之类的功能. 

google就为我们提供一个解决方案, 开源了`context`包. 使用`context`实现上下文功能约定需要在你的方法的传入参数的第一个传入一个`context.Context`类型的变量. 

## 2. 使用方法

网上大部分文章的获取`context.Context`对象时都会用到`context.Background()`或`context.TODO()`函数, 查看源码你会发现, 这两个返回的都是空context: `emptyCtx`对象(其实是`int`的别名).

### 2.1 `WithCancel()`可以手动取消的Context示例

```go
	ctx, cancel := context.WithCancel(context.Background())
	//每1秒work一下, 同时会判断ctx是否被取消了, 如果是就退出
	go func(ctx context.Context) {
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				log.Printf("done")
				// context的Err()方法返回为什么被取消
				log.Printf("%s", ctx.Err())
				return
			default:
				log.Printf("work")
			}
		}
	}(ctx)
	//5秒后手动取消
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	log.Printf("end")
```

```
2018/07/02 17:09:27 work
2018/07/02 17:09:28 work
2018/07/02 17:09:29 work
2018/07/02 17:09:30 work
2018/07/02 17:09:31 done
2018/07/02 17:09:31 context canceled
2018/07/02 17:09:33 end
```

`context`包提供了多种控制相同context下的`goroutine`的方法, `WithCancel`只是其中一个. 与之类似还有`WithTimeout`, `WithDeadline`和`WithValue`.

### 2.2 `WithTimeout`为goroutine设置一个超时时间

以`WithTimeout()`为例, 我们可以为一个`goroutine`设置一个超时时间. 使用WithTimeout方法就不必手动调用cancel方法了, 但cancel方法依然是有效的.

```go
	// 使用WithTimeout方法就不必手动调用cancel方法了, 但cancel方法依然是有效的.
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go func(ctx context.Context) {
		//每1秒work一下, 同时会判断ctx是否被取消了, 如果是就退出
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
                log.Printf("done")
				// context的Err()方法返回为什么被取消
				log.Printf("%s", ctx.Err())
				return
			default:
				log.Printf("work")
			}
		}
	}(ctx)
	time.Sleep(7 * time.Second)
	log.Printf("end")
```

```
2018/07/02 17:11:27 work
2018/07/02 17:11:28 work
2018/07/02 17:11:29 work
2018/07/02 17:11:30 work
2018/07/02 17:11:31 done
2018/07/02 17:11:31 context deadline exceeded
2018/07/02 17:11:33 end
```

### 2.3 `WithDeadline`设置deadline时间

```go
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))

	go func() {
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				log.Println("done")
				// context的Err()方法返回为什么被取消
				log.Printf("%s", ctx.Err())
				return
			default:
				log.Println("work")
			}
		}
	}()

	time.Sleep(time.Second * 7)
	log.Printf("end")
```

```
2019/05/29 11:32:46 work
2019/05/29 11:32:47 work
2019/05/29 11:32:48 work
2019/05/29 11:32:49 work
2019/05/29 11:32:50 done
2019/05/29 11:32:50 context deadline exceeded
2019/05/29 11:32:52 end
```

可以看到, `WithDeadline`能达到的效果其实与`WithTimeout`是相同的, 只不过参数类型不同而已.

阅读一下源码你会发现, `WithTimeout`实际上就是调用了`WithDeadline`而已.

### 2.4 `WithValue`传参

`func WithValue(parent Context, key, val interface{}) Context`

`WithValue`在传入一个Context对象的同时, 还可传入一对键值, 可在内层函数中取用.

注意, `WithValue`只返回一个`Context`对象, 没有`cancel`方法, 所以一般`WithValue`是与其他`WithXXX`函数一起使用的(不然根本没法控制, 没有意义). 

不过传入的键值只有一对, 感觉...挺鸡肋的.

如下是我想到的一种可行场景.

```go
	ctx := context.WithValue(context.Background(), "session", "BAIDUID=90173631B0")
	ctx, _ = context.WithTimeout(ctx, time.Second*5)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				log.Println("done")
				// context的Err()方法返回为什么被取消
				log.Printf("%s", ctx.Err())
				log.Printf("当前session: %s 已失效", ctx.Value("session").(string))
				return
			default:
				log.Println("work")
			}
		}
	}()

	time.Sleep(time.Second * 7)
	log.Printf("end")
```

```
2019/05/29 12:48:17 work
2019/05/29 12:48:18 work
2019/05/29 12:48:19 work
2019/05/29 12:48:20 work
2019/05/29 12:48:21 done
2019/05/29 12:48:21 context deadline exceeded
2019/05/29 12:48:21 当前session: BAIDUID=90173631B0 已失效
2019/05/29 12:48:23 end
```

虽然上述代码可运行成功, 但是`WithValue`在函数注释中表示key需要是可比较类型, 且不应该是内置类型(为了达到这样的要求, 可以使用`type type1 int`创建自定义的类型).

## 3. Context对象方法

### 3.1 Deadline()

### 3.2 Done()

一般情况下返回一个channel, 在手动取消或是超时情况下可读, 表示任务结束.

如果context不可被取消, 则返回nil.

### 3.3 Err()

在`Done()`返回的channel没有关闭的时候, `Err()`的值为nil.

在`Done()`返回的channel关闭之后, `Err()`会返回一个error类型的值解释channel关闭的原因:

1. 被取消
2. 超时

### 3.4 Value()

## 4. 总结 - Context使用规范

- `context`是协程程安全的
- 不要把`context`存储在结构体中, 而是要显式地进行传递
- 把`context`作为第一个参数, 并且一般都把变量命名为`ctx`
- 就算是程序允许, 也不要传入一个`nil`的`context`, 如果不知道是否要用`context`的话, 用`context.TODO()`来替代
- `context.WithValue()`只用来传递请求范围的值, 不要用它来传递可选参数
- `WithTimeout`可设置超时时间, 但是如果子协程在超时之前完成, 需要手动`cancel()`以释放资源.

> 这些在`context`源码中都有注释说明.
