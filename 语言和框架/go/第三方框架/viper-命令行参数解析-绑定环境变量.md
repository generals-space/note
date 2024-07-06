# viper-绑定环境变量

```go
package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("BACKEND")
	// 合法的环境变量只能包含下划线_, 不能包含中横线或点号
	// replacer用于将目标key转换成合法的环境变量字符串格式
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetDefault("ttn-manager", ":7301")
	viper.SetDefault("ttn.manager", ":7301")
	fmt.Println(viper.GetString("ttn-manager"))
	fmt.Println(viper.GetString("ttn.manager"))
}
```

直接运行, 将得到

```
:7301
:7301
```

设置了对应的环境变量后, 如`export BACKEND_TTN_MANAGER=xxx`, 将得到

```
xxx
xxx
```

`BACKEND`为环境变量的前缀, `SetEnvKeyReplacer()`将程序中变量的分隔符转换成下划线`_`去匹配环境变量, 所以变量不好用驼峰形式书写, 最好写成中横线或点号连接.

另外, 我尝试了下移动`AutomaticEnv()`和`SetEnvKeyReplacer()`的位置, 事实证明只要在`viper.GetString()`前调用这两个函数就可以生效, 它们与设置变量默认值的先后顺序没有影响.