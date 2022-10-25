# golang-json空字段报错校验[tag required nonempty]

参考文章

1. [How not to allow empty fields when Unmarshalling json to a struct](https://stackoverflow.com/questions/56895966/how-not-to-allow-empty-fields-when-unmarshalling-json-to-a-struct)
2. [Go 中使用 JSON 时，如何区分空字段和未设置字段](https://zhuanlan.zhihu.com/p/347574574)

golang 原生的 json tag 并没有空字段校验的能力(类似于`required`标记), 只能借助第3方库实现.

参考文章1, 2中介绍了一种方法, 将字段设置为指针类型.

```go
type DictInfo struct {
    ClusterName string `json:"clusterName"`
    TimeStamp string `json:"timeStamp"`
}
```

```go
type DictInfo struct {
    ClusterName *string `json:"clusterName,omitempty"`
    TimeStamp *string `json:"timeStamp,omitempty"`
}
```

不过好像是用来鞭策开发者强制手动校验的功能, 而不是在`json.UnMarshal()`的时候自动校验.

