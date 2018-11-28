# golang-判断map中是否存在某key

从一个map中取一个不存在的key时, 总返回0值.

```go
package main

import (
	"fmt"
)
func main(){
    theMap := map[string] string {
        "one":"1", 
        "two":"2",
    }
    keyList := []string{"one", "two", "three"}
    for _, key := range keyList {
        value, ok := theMap[key]; 
        if !ok {
            fmt.Printf("There is not the key: %s\n", key)
        }else {
            fmt.Printf("%s: %s\n", key, value)
        }
    }
}
```

输出为

```
one: 1
two: 2
There is not the key: three
```