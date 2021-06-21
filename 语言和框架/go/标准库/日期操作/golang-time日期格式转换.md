# golang-time日期格式转换

参考文章

1. [golang -- 时间日期总结](https://studygolang.com/articles/669)

## 1. golang日期对象转时间戳和格式化字符串

```go
	now := time.Now() // Time对象

	timeStamp := now.Unix() // int64时间戳
	log.Printf("%d\n", timeStamp)

	/*
	 * 这种格式化真的是奇葩, 其他语言中的"%Y-%M-%d %H:%m:%S"
	 * 在go语言中对应的是"2006-01-02 15:04:05"
	 * 这个时间点是go诞生之日, 记忆方法:6-1-2-3-4-5.
	 * 你可以把2006当作%Y来用, 比如"2006xx2006"会输出:
	 * 当前年份xx当前年份...
	 */
	// timeStr := now.Format("2006-01-02 15:04:05")
	timeStr := now.Format("2006xx2006")
	log.Printf("%s\n", timeStr)
```

## 2. 反向转换, 时间戳/字符串转换为日期对象

```go
	// 时间戳转Time对象
	var timeStamp int64 = 1525788126
	// Unix(sec, nsec), sec为秒级整型, nsec为纳秒级整型
	timeObj := time.Unix(timeStamp, 0)
	timeStr := timeObj.Format("2006-01-02 15:04:05")
	log.Printf("%s\n", timeStr) // 2018-05-08 22:02:06

	// 字符串转Time对象
	timeStr = "2016-01-01 12:00:00"
	timeObj, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		panic(err)
	}
	timeStamp = timeObj.Unix()
	log.Printf("%d\n", timeStamp)
```

------

日期可用函数:

- `timeObj.Year()`
- `timeObj.Month()`: 结果为`Month`类型, 其实为`int`类型, 不过它有一个`String()`方法, 可以转换为英文单词的形式(如`January`).
- `timeObj.Day()`
- `timeObj.Hour()`
- `timeObj.Minute()`
- `timeObj.Second()`
- `timeObj.Location()`: 所在时区, Location类型, 如UTC, CST等.

