# grpc连接机制

参考文章

1. [grpc 超时和重连](https://www.jianshu.com/p/a5dec04d042b)

其实在使用时就知道grpc有重连机制了, 连接断线会在日志中不断打印警告信息, 说明grpc客户端SDK本身通过心跳维护着连接. 就像常用数据库的连接驱动一样. 但是与数据库连接不同的是, 数据库在服务启动时, 如果连接失败是不会创建连接对象的, 而是直接退出, 把这当作是严重错误. 而grpc不会, 如果在`Dial()`参数中不手动指定超时选项的话, 会直接返回一个连接对象, 之后不断尝试连接...

官方文档貌似没有详细介绍这个话题, 需要查看源码.

```go
	customerConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Errorf("connect customer service failed: %s", err.Error())
		return
	}
	fmt.Printf("%+v\n", customerConn.GetState()) // 直接输出IDLE
	customerCli := protos.NewCustomerServiceClient(customerConn)
```

上述代码连接名为`customer`的服务, 在`customer`根本没开启的情况下运行时依然没出错, `err`为`nil`, 此时连接状态为`IDLE`, 表示连接空闲, 其实相当于异步了.
