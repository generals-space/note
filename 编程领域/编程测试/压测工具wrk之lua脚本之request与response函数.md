# 压测工具wrk之lua脚本

参考文章

1. [使用 WRK 压力测试工具对 ASP.NET Core 的接口进行压力测试](https://www.cnblogs.com/myzony/p/9798116.html)
    - wrk调用lua的3个阶段: setup, running, done, 有图示, 各个函数也有注释, 比较清晰.
2. [Http压测工具wrk使用指南](https://www.cnblogs.com/xinzhao/p/6233009.html)
    - wrk调用lua的一些常用场景, 比如request方法, delay延迟, 访问认证等.
3. [wrk用lua脚本构建复杂的http压力测试](http://xiaorui.cc/2018/03/14/wrk%E7%94%A8lua%E8%84%9A%E6%9C%AC%E6%9E%84%E5%BB%BA%E5%A4%8D%E6%9D%82%E7%9A%84http%E5%8E%8B%E5%8A%9B%E6%B5%8B%E8%AF%95/)
    - request与response的复杂示例

关于wrk的生命周期各个钩子函数不再具体介绍, 这里以实现为主讲述ta们的用法. 

首先就是常用的request与response.

## response

最简单的示例就是直接打印了. 

```lua
function response(status, headers, body)
    print(status, headers, body)
end
```

在wrk调用脚本时, 压测过程中会输出每次请求的结果.

```
$ wrk -t 1 -c 1 -d 1 -s post.lua http://192.168.0.8:7777/uplink
Running 1s test @ http://192.168.0.8:7777/uplink
  1 threads and 1 connections
200     table: 0x40c6c410       {
  "code": 0,
  "msg": "",
  "data": null
}
200     table: 0x40c6c550       {
  "code": 0,
  "msg": "",
  "data": null
}
...省略
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    33.33ms    8.19ms  57.97ms   65.52%
    Req/Sec    29.00      3.16    30.00     90.00%
  29 requests in 1.00s, 4.73KB read
Requests/sec:     28.96
Transfer/sec:      4.72KB
```

## request

关于request, 其实request函数的返回值就某个具体的请求, 所以request必须要有返回值. 

你可以在函数体内做一些处理, 动态修改要发送的请求(一般是请求体或是路径参数等信息), 然后返回请求字符串对象.

```lua
function request()
    data = [[{
        "id": 1,
        "name": "general"
    }]]
    wrk.method = "POST"
    wrk.body = string.format(data)
    wrk.headers["Content-Type"] = "application/json"
    return wrk.format()
end
```

上述代码同样可以发送post请求, 与之前写的一篇文章有些不同.

------

上面的两个例子都太过简单了, 而且好像没什么实用价值. 参考文章3最后一个实例我觉得是个不错的场景.

```lua
token = nil
path  = "/auth"
 
request = function()
    return wrk.format("GET", path)
end
 
response = function(status, headers, body)
    if not token and status == 200 then
        token = headers["X-Token"]
        path  = "/resource"
        wrk.headers["X-Token"] = token
    end
end
```

如果一个接口需要先认证, 获取token后才能访问, 那么token初始值先设置为nil, 请求的路径path也为`/auth`. 在第一次请求完成后拿到响应头中的`X-Token`字段并赋值给token变量, 并且修改path路径. 之后的请求就可以按照普通的流程进行了.
