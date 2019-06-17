# go-redis与redigo

参考文章

1. [Go实战--golang中使用redis(redigo和go-redis/redis)](https://blog.csdn.net/wangshubo1989/article/details/75050024)
    - [go-redis/redis](https://github.com/go-redis/redis)
    - [gomodule/redigo](https://github.com/gomodule/redigo)

参考文章1详细介绍了`redigo`与`go-redis`的使用方法, 通篇看下来, 我觉得两者最大的区别在于

`redigo`提供的`Do()`方法类似于`redis-cli`命令行工具, ta单纯把你写的命令发送给`redis`实例, 然后取回结果, 所有的命令都通过`Do()`来封装, 没有其他的方法.

而`go-redis`则提供了`Set`, `Get`, `Del`等命令(包括对列表, 哈希的操作), 当然ta也提供了`Do()`方法, 可以接受原生`redis`命令.

```go
result1 := client.Keys("*")
fmt.Printf("%+v\n", result1)  // keys *: [name]

result2 := client.Keys("*").Val()
fmt.Printf("%+v\n", result2) // [name]

result3 := client.Do("keys", "*")
fmt.Println(result3) // &{{[keys *] <nil> <nil>} [name]}

result4 := client.Do("keys", "*").Val()
fmt.Println(result4) // [name]
```

> 按照**峰云**的说法, go redis库, 推荐使用`go-redis`. `redigo`那个库连个cluster都不整, 作者把cluster实现的机会让他其他人了. redigo的连接池实现有点操.
