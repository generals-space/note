# golang-Parse与ParseInLocation输出字符串格式不需要带时区而读入需要

参考文章

1. [Golang time包](https://www.jianshu.com/p/d62d605c91fa)

```go
func main() {
    // Format() 的输出默认是按照本地时区来的
	now := time.Now()
	log.Printf("time: %+v, string: %s", now, now.Format("2006-01-02 15:04:05"))

    // 但是 Parse() 默认却是使用 UTC 时间读入的...
	timeStr := "2021-10-13 11:04:37"
	startAt, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	log.Printf("time: %+v, cst time: %+v", startAt, startAt.In(time.Local))

    // 所以在读入时间字符串时, 应将 ParseInLocation() 作为首选项
	startAt2, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	log.Printf("time: %+v, cst time: %+v", startAt2, startAt2.In(time.Local))
}
```

```
2021/10/13 16:57:07 time: 2021-10-13 16:57:07.1573965 +0800 CST m=+0.003914201, string: 2021-10-13 16:57:07
2021/10/13 16:57:07 time: 2021-10-13 11:04:37 +0000 UTC, cst time: 2021-10-13 19:04:37 +0800 CST
2021/10/13 16:57:07 time: 2021-10-13 11:04:37 +0800 CST, cst time: 2021-10-13 11:04:37 +0800 CST
```
