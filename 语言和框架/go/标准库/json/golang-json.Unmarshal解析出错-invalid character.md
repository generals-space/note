# json相关

## 1. 

```
“invalid character '\x00' after top-level value”
```

情境描述

在使用`json.Unmarshal()`解析json字节数组为struct时报上述错误.

原因在于, 读取json文件时, 预先`make`了一个1024个字节`[]byte`数组, 如果读取的json文件小于1024个字节, 那么`[]byte`数组就没有被填满, `Unmarshal`解析时就会报错.

解决方法为, `[]byte`数组实际占用长度为多少, 就解析到哪里.

```go
fileCnt := make([]byte, 1024)
var device Device
num, err := _file.Read(fileCnt)
if err != nil {
    return
}
// 注意这里
err = json.Unmarshal(fileCnt[:num], &device)
```
