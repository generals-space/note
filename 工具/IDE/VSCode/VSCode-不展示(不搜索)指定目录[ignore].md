# VSCode-不展示(不搜索)指定目录[ignore]

参考文章

1. [](https://stackoverflow.com/questions/30140112/how-to-hide-specified-files-directories-e-g-git-in-the-sidebar-vscode)

```json
    // 左侧浏览器中不展示如下路径的内容
    "files.exclude": {
        "**/.git": true,
    },
    // 全局搜索不搜索如下目录内容
    "search.exclude": {
        "**/node_modules": true,
    }
```
