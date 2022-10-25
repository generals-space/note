# golang-strconv FormatInt与ParseInt[进制转换]

## 10进制 -> 其他进制

```go
	// 转换为16进制和2进制, FormatUint也是同样的用法
	str2 := strconv.FormatInt(int64(123456789), 2)
	log.Println(str2) // 111010110111100110100010101
	str16 := strconv.FormatInt(int64(123456789), 16)
	log.Println(str16) // 75bcd15

```

注意: 这里得到的结果为 string 类型.

## 其他进制 -> 10进制

```go
    num2, _ := strconv.ParseInt(str2, 2, 64)
	log.Printf("%d\n", num2) // 123456789

    num16, _ := strconv.ParseInt(str16, 16, 64)
	log.Printf("%d\n", num16) // 123456789
```

这里`ParseInt()`读入的是 string 类型, 第3个参数为最终转换出来的数值的位数上限???(32位/64位). 当然, 返回值结果是确定的`int64`, 但是如果指定为32, 则可能会舍弃前面的32位, 只保留后32位.
