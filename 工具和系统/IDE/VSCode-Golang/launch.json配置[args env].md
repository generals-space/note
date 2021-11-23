# launch.json配置[args env]

参考文章

1. [【VSCode】golang的调试配置launch.json](https://www.jianshu.com/p/e4cca4fe6478)

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "golang",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            //当运行单个文件时{workspaceFolder}可改为{file}
            "program": "${workspaceFolder}",
            "env": {},
            "args": []
        }
    ]
}
```

- `name`: 取值随意, 一般用当前工程名称. 可以在左侧菜单栏中"Run and Debug"中看到该值.
- `request`: 可选值: "launch", "attach", 前者是自己启动调试进行, 后者则可以链接到一个正在运行的进程.
- `program`: 工程路径
    - ${workspaceFolder}
    - ${fileDirname}
    - ${file}: 运行单个文件时可使用
- `args`: 可以添加启动参数.

