参考文章

1. [VSCode debug error : Version of Delve is too old for this version of Go](https://github.com/go-delve/delve/issues/1974)
    - "dlvFlags": ["--check-go-version=false"] 
2. [vscode利用delve调试go1.12代码_呀一1不小心的博客-程序员宝宝](https://www.cxybb.com/article/qq_28382661/118703683)
    - delve 的 [CHANGELOG](https://github.com/go-delve/delve/blob/master/CHANGELOG.md) 可以检索到各版本golang语言的支持记录，如delve是在v1.3.0添加go1.12的支持的。

参考文章1中说, 可以在 launch.json 中添加`dlvFlags`参数, 让 delve 不再进行版本检查.

需要更新 go 插件, Go 0.14.4 是不支持`dlvFlags`参数的.

Property dlvFlags is not allowed.
