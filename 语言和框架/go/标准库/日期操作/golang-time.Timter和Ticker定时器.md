# golang-time.Timter和Ticker定时器

参考文章

1. [go timer 和 ticker 的区别 ](https://learnku.com/articles/23578/the-difference-between-go-timer-and-ticker)

- `time.NewTimer(d Duration) *Timer`
- `time.NewTicker(d Duration) *Ticker`

`ticker`类似于js中的`setInterval`, 是循环调用; 而`timer`则类似于`setTimeout`, 是单次调用.

`ticker`对象只有一个方法: `Stop()`, `timer`对象则有两个方法: `Reset()`与`Stop()`.

## 1. 简单使用方法

与`setTimeout`相似, 如果要实现类似`setInterval`循环执行函数的行为, 需要在函数中重新设置`setTimeout`. `timer`可以通过`Reset()`方法重置`timer`对象, 然后下一次到达指定时间又可以执行.

timer的简单使用方法如下

```go
	log.Println("start")
	timer := time.NewTimer(time.Second * 2)
	<-timer.C
	log.Println("2s后")
	// 此时如果再次 <-timer.C会阻塞, 需要reset重新设置定时器
	for {
		timer.Reset(time.Second * 2)
		<-timer.C
		log.Println("work")
	}
	// Ctrl-C结束
```

执行结果

```
2019/06/01 15:02:30 start
2019/06/01 15:02:32 2s后
2019/06/01 15:02:34 work
2019/06/01 15:02:36 work
2019/06/01 15:02:38 work
exit status 2
```

ticker的使用方法如下

```go
	log.Println("start")
	ticker := time.NewTicker(time.Second * 2)
	for {
		<-ticker.C
		log.Println("work")
	}
	// Ctrl-C结束
```

```
2019/06/01 15:01:12 start
2019/06/01 15:01:14 work
2019/06/01 15:01:16 work
2019/06/01 15:01:18 work
exit status 2
```

## 2. time包中其他类似方法

### 2.1 After()

`After func(d Duration) <-chan Time`

`After`返回一个channel, 在指定时间d后可读, 类似于timer对象中的C成员.

```go
	log.Println("start")
	tChannel := time.After(time.Second * 2)
	t := <-tChannel
    log.Println("2s后")
	log.Println(t)
```

```
2019/06/01 15:09:20 start
2019/06/01 15:09:22 2s后
2019/06/01 15:09:22 2019-06-01 15:09:22.7600229 +0800 CST m=+2.035148601
```

### 2.2 AfterFunc()

`AfterFunc func(d Duration, f func()) *Timer`

```go
	log.Println("start")
	time.AfterFunc(time.Second*2, func() {
		log.Println("work")
	})
	time.Sleep(time.Second * 3)
	log.Println("stop")
```

```
2019/06/01 15:00:21 start
2019/06/01 15:00:23 work
2019/06/01 15:00:24 stop
```

### 2.3 Tick()

`Tick func(d Duration) <-chan Time`

与`After`相似, ta们的区别就如同timer与ticker的区别, `Tick()`返回的channel可以循环读取, 而`After()`的只能读取一次.

```go
	log.Println("start")
	tChannel := time.Tick(time.Second * 2)
	for {
		t := <-tChannel
		log.Println("work")
		log.Println(t)
	}
	// Ctrl-C结束
```

```
2019/06/01 15:17:32 start
2019/06/01 15:17:34 work
2019/06/01 15:17:34 2019-06-01 15:17:34.238464 +0800 CST m=+2.034506401
2019/06/01 15:17:36 work
2019/06/01 15:17:36 2019-06-01 15:17:36.2380019 +0800 CST m=+4.034044301
2019/06/01 15:17:38 work
2019/06/01 15:17:38 2019-06-01 15:17:38.2384128 +0800 CST m=+6.034455201
exit status 2
```