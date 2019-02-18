# golang-interface{}与[]interface{}

参考文章

1. [golang []interface{} 的数组如何转换为 []string 的数组](https://segmentfault.com/q/1010000003505053)

有json字符串如下

```json
{
    "list": [12, 23.4, 45]
}
```

想要将其反序列化到一个类型为`map[string]interface{}`的对象中.

