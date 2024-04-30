# go run -mod选项[vscode launch.json]

参考文章

1. [Build commands](https://go.dev/ref/mod#build-commands)
    - 官方文档

该特性应该是在 go v1.15 前后添加的, 如

```
go build -mod=mod -o main main.go
```

`-mod`参数的可选项有: `mod`, `vendor`, `readonly`, 默认为`vendor`.

`-mod=vendor`表示直接使用工程目录下`vendor`子目录中的依赖包, 且不再从`GOPATH/pkg/mod`下的本地module缓存中寻找, 也不会从网上下载了. 

相当于**离线模式**, 但并不纯粹, ta会将`vendor/modules.txt`与`go.mod`中的依赖包版本做对比, 如果不一致则可能会报如下错误

```log
$go build -ldflags="-s -w" -mod=vendor -a -o backend cmd/backend/app/main.go
go: inconsistent vendoring in D:\Code\unifiedportal\unified-oam:
        stellaris: is replaced by ./../../stellaris in go.mod, but marked as replaced by unifiedportal/stellaris.git@v1.2.0-alpha2 in vendor/modules.txt

        To ignore the vendor directory, use -mod=readonly or -mod=mod.
        To sync the vendor directory, run:
                go mod vendor

```

`-mod=mod`表示常规的 go module 项目, 依赖包都在`GOPATH/pkg/mod/`目录下, 执行时也会自动将工程中用到但没记录的依赖包, 写入`go.mod`文件.

`-mod=readonly`可以看作是`-mod=mod`的特殊形式, ta仍然是 go module 形式(忽略 vendor 目录), 但是不再自动变更`go.mod`文件, 如果有哪些工程中用到但没出现在`go.mod`文件中的话, 就报错.

------

```json
{
    "go.buildFlags": [
        "-mod=vendor"
    ]
}
```
