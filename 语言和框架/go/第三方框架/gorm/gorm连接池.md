# gorm连接池

参考文章

1. [Pooling/Concurrency](https://github.com/jinzhu/gorm/issues/246)

gorm是协程安全的, 默认支持连接池, 可以通过如下语句设置连接池属性.

```go
db.DB().SetMaxIdleConns(100)
```