# settings.json配置

## 自定义后缀文件语法关联

在阅读 golang 1.2 版本源码时, 存在后缀为`.goc`的源文件, 不过其内容是C写的, 可以通过如下配置让vscode将这种后缀的文件视为C语言源文件.

```json
{
    "files.associations": {
        "*.goc": "c"
    }
}
```
