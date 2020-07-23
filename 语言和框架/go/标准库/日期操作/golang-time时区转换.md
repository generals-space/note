# golang-time时区转换

参考文章

1. [golang时区转换](https://blog.csdn.net/hehexiaoxia/article/details/71629225)
2. [golang time包使用时注意时区](https://blog.csdn.net/su_sai/article/details/52913820)

```go
	now := time.Now()
	DateFormat1 := "2006-01-02T15:04:05"

	DateFormat2 := "2006-01-02T15:04:05-07:00"
	DateFormat3 := "2006-01-02T15:04:05 -07:00"
	DateFormat4 := "2006-01-02T15:04:05 -0700"
	DateFormat5 := "2006-01-02T15:04:05Z07:00"
	DateFormat6 := "2006-01-02T15:04:05 Z07:00"
	DateFormat7 := "2006-01-02T15:04:05 Z0700"

	nowStr1 := now.Format(DateFormat1)
	nowStr2 := now.Format(DateFormat2)
	nowStr3 := now.Format(DateFormat3)
	nowStr4 := now.Format(DateFormat4)
	nowStr5 := now.Format(DateFormat5)
	nowStr6 := now.Format(DateFormat6)
	nowStr7 := now.Format(DateFormat7)

	fmt.Println(nowStr1) // 2018-06-15T17:02:08
	fmt.Println(nowStr2) // 2018-06-15T17:02:08+08:00
	fmt.Println(nowStr3) // 2018-06-15T17:02:08 +08:00
	fmt.Println(nowStr4) // 2018-06-15T17:02:08 +0800
	fmt.Println(nowStr5) // 2018-06-15T17:02:08+08:00
	fmt.Println(nowStr6) // 2018-06-15T17:02:08 +08:00
	fmt.Println(nowStr7) // 2018-06-15T17:02:08 +0800
```

在golang中, 除了`2006-01-02 15:04:05`这些特殊的占位符外, 还有`-07:00`或`Z07:00`(或`-0700`, `Z0700`), 它表示时区.
