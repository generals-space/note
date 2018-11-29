# 计算字符串及文件的md5及sha1等值

参考文章

1. [golang 中的md5 、hmac、sha1算法的简单实现](https://blog.csdn.net/yue7603835/article/details/73497034)

2. [[Golang]计算一个文件的MD5值](https://blog.csdn.net/bobodem/article/details/80227029)

```go
package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Md5 ...
func Md5(data string) string {
	md5Obj := md5.New()
	md5Obj.Write([]byte(data))
	// return hex.EncodeToString(md5Obj.Sum(nil))
	return hex.EncodeToString(md5Obj.Sum([]byte("")))
}

// Sha1 ...
func Sha1(data string) string {
	sha1Obj := sha1.New()
	sha1Obj.Write([]byte(data))
	return hex.EncodeToString(sha1Obj.Sum([]byte("")))
}

// Sha256 ...
func Sha256(data string) string {
	sha256Obj := sha256.New()
	sha256Obj.Write([]byte(data))
	return hex.EncodeToString(sha256Obj.Sum([]byte("")))
}

// Hmac ...
func Hmac(key, data string) string {
	hmacObj := hmac.New(md5.New, []byte(key))
	hmacObj.Write([]byte(data))
	return hex.EncodeToString(hmacObj.Sum([]byte("")))
}

func main() {
	fmt.Println(Md5("hello"))          // 5d41402abc4b2a76b9719d911017c592
	fmt.Println(Sha1("hello"))         // aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d
	fmt.Println(Sha256("hello"))       // 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
	fmt.Println(Hmac("key2", "hello")) // f1b90b4efd0e5c7db52dfa0efd6521a3
}
```

对比md5, sha1算法的python标准库中的使用方式, 可以看出, 

`Write()` == `update()`, 向md5对象中推入新的字节内容, 不断更新.

`Sum()` == `digest()/hexdigest()`, 生成最终的结果, `Sum()`接受一个参数, 表示向当前md5对象追回最后一次字节内容, 然后生成最终结果. 不过我搜到的网上的教程都是填入空字节数组的, 不清楚ta的实际意义所在, 暂不深究<???>
