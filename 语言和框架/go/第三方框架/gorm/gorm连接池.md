# gorm连接池

参考文章

1. [Pooling/Concurrency](https://github.com/jinzhu/gorm/issues/246)

2. [Documentation on how the database pool works?](https://github.com/jinzhu/gorm/issues/1334)

gorm是协程安全的, 默认支持连接池, 可以通过如下语句设置连接池属性.

```go
db.DB().SetMaxOpenConns(100) // 默认不限制
db.DB().SetMaxIdleConns(50) // 默认为2
```

另外, gorm使用原生标准库的连接池, 如果使用`SetMaxOpenConns`设置连接池数量为10, 而同时进入了100个请求, 那么其表现将与标准库所定义的行为一样, 直接拒绝.
