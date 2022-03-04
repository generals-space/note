# Win10 启用管理员账户Administrator

以管理员身份运行 powershell 输入如下命令.

```console
$ net user administrator /active:yes
命令成功完成。

$ net user administrator 123456
命令成功完成。
```

第一条命令执行后, 点击"开始" -> 账户头像, 就会出现`Administrator`的选项, 点击即可切换.
