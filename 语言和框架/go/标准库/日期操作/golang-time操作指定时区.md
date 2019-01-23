# golang-time操作指定时区

参考文章

1. [Golang time.Parse和time.Format的时区问题](https://www.jianshu.com/p/92d9344425a7)

2. [Golang时区设置](https://my.oschina.net/jsk/blog/1817395)

## 1. `time.ParseInLocation()`和`time.Parse()`

## 2. `time.LoadLocation("UTC")`和`In()`

```go
loc, _ := time.LoadLocation("Asia/Shanghai")
for _, CreatedAt := range CreatedAts {
    log.Printf("%s\n", CreatedAt.In(loc).Format("2006-01-02 15:04:05-07"))
}
```
