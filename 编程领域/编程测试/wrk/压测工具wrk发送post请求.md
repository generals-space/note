# wrk发送post请求

参考文章

1. [wrk issue - JSON Post](https://github.com/wg/wrk/issues/267)
2. [wrk post.lua](https://github.com/wg/wrk/blob/master/scripts/post.lua)
3. [wrk用lua脚本构建复杂的http压力测试](http://xiaorui.cc/2018/03/14/wrk%E7%94%A8lua%E8%84%9A%E6%9C%AC%E6%9E%84%E5%BB%BA%E5%A4%8D%E6%9D%82%E7%9A%84http%E5%8E%8B%E5%8A%9B%E6%B5%8B%E8%AF%95/)

没找到行内编写post请求的方法, 这一点比curl还不方便, 只能编写lua脚本.

`post.lua`

```lua
wrk.method = "POST"
wrk.body   = '{"id": "1","name": "general"}'
wrk.headers["Content-Type"] = "application/json"
function request()
    return wrk.format('POST', nil, nil, body)
end
```

```
$ wrk -t 5 -c 5 -d 1 -s post.lua https://localhost/login
```

网上有找到上述代码, 但是服务端根本接收不到发送的数据. 

但是如果把`Content-Type`指定为`application/x-www-form-urlencoded`, 且body的格式改为`id=1&name=general`, 就可以接收到. 这也是参考文章1提到的问题.

实际上, 官方给出的post示例发送的数据就是`x-www-form-urlencoded`格式, 见参考文章2. 网上有很多`json`格式的示例, 不知道是怎么得来的.

wrk版本试验过`3.8.0`和`4.1.0`两个版本, 都是这种情况.

然后发现了参考文章3, ta给出了post发送json的示例, 使用`string.format()`编码对象.

```lua
data = [[{
    "id": 1,
    "name": "general"
}]]
wrk.method = "POST"
wrk.body = string.format(data)
wrk.headers["Content-Type"] = "application/json"
```

这种方法证明有效, `3.8.0`和`4.1.0`两个版本都有效.
