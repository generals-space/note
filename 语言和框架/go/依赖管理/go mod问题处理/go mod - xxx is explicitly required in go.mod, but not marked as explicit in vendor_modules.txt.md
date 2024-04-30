# go mod - xxx is explicitly required in go.mod, but not marked as explicit in vendor_modules.txt

## 问题描述

go: 1.22.2

一般来说, 这种场景只出现在, 开发者以 go module 形式, 使用`go get`在工程目录中下载了xxx依赖, 并更新了`go.mod`文件.

但是在没有执行`go mod vendor`的时候(此时xxx还没有被拷贝到 vendor 目录), 就使用`go build -mod=vendor`直接以`vendor`目录为标准去编译了.

`-mod=vendor`会将`vendor/modules.txt`与`go.mod`做对比, 如果发现上述情况就会报错了.
