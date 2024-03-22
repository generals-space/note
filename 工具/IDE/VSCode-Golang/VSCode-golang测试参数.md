# VSCode-golang测试参数

参考文章

1. [How to disable (run test) Cached](https://github.com/Microsoft/vscode-go/issues/1597)

2. [Run tests with verbose output](https://github.com/Microsoft/vscode-go/issues/1377)

vscode在测试golang代码时, 发现测试结果都是旧的. 接口连接了数据库并写入一条数据, 但是只有第一条能写入, 之后的测试结果就显示为`cached`, 数据库中没有出现新数据. 如下

![](https://gitee.com/generals-space/gitimg/raw/master/3b9b23c5a06cb9598f6a7ae4da090ba7.jpg)

查看参考文章1, 会发现测试参数是由golang本身提供的, vscode只是代为执行`go test`命令而已. 我们需要在测试命令中加入`-count=1`选项.

正好在参考文章2中有提到`go test`测试过程中加入`-v`选项, 是在vscode的自定义配置文件中添加的, 顺便也把`count`选项加上.

最后的vscode配置文件为

```json
{
    "go.testFlags": ["-v", "-count=1"],
}
```

再次执行测试命令, 不再有缓存, 输出结果也变得丰富起来.

![](https://gitee.com/generals-space/gitimg/raw/master/7fbce72b95507bcf763d034db15e0249.jpg)
