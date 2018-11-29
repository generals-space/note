# 计算大文件md5及sha1值(golang版)

参考文章

```go
package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"io"
)

// MD5FileByPartsWithPath ...
func MD5FileByPartsWithPath(filepath string) (md5code string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("open file failed: %s", err.Error())
	}
	defer file.Close()
	
	md5obj := md5.New()
	for {
		buf := make([]byte, 1024)
		length, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("end of file")
				err = nil
				break
			}else {
				log.Printf("read file error: %s\n", err.Error())
				return md5code, err
			}
		}
		md5obj.Write(buf[:length])
	}
	md5byte := md5obj.Sum([]byte(""))
	// md5byte := md5obj.Sum(nil)
	md5code = hex.EncodeToString(md5byte)
	return 
}

// MD5FileWithPath ...
func MD5FileWithPath(filepath string)(md5code string, err error){
	cnt, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("read file error: %s\n", err)
		return
	}

	md5obj := md5.New()
	md5obj.Write(cnt)
	md5byte := md5obj.Sum([]byte(""))
	// md5byte := md5obj.Sum(nil)
	md5code = hex.EncodeToString(md5byte)
	return 
}

func main() {
	filepath := "/Users/general/Movies/Lost Ark - Open Beta CG Trailer - PC - F2P - KR.mp4"

	md5code1, err := MD5FileWithPath(filepath)
	if err != nil {
		log.Printf("MD5FileWithPath() error: %s", err.Error())
	}
	log.Println(md5code1) // 63a47d8af1ecd7457a61614a16c61d6c
	md5code2, err := MD5FileByPartsWithPath(filepath)
	if err != nil {
		log.Printf("MD5FileByPartsWithPath() error: %s", err.Error())
	}
	log.Println(md5code2) // 63a47d8af1ecd7457a61614a16c61d6c
}
```

执行结果

```
2018/11/29 22:26:35 63a47d8af1ecd7457a61614a16c61d6c
2018/11/29 22:26:35 end of file
2018/11/29 22:26:35 63a47d8af1ecd7457a61614a16c61d6c
```