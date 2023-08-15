# go -mod选项[vscode launch.json]

该特性应该是在 go v1.15 前后添加的.

- go build -mod=mod -o main main.go
- go build -mod=vendor -o main main.go

`-mod=mod`表示常规的 go module 项目, 依赖包都在 GOPATH/pkg/mod/ 目录下, 而`-mod=vendor`则可以将依赖包指向工程目录下的 vendor 目录.

```json
{
    "go.buildFlags": [
        "-mod=vendor"
    ]
}
```
