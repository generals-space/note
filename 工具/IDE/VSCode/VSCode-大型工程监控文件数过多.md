# VSCode-大型工程监控文件数过多

参考文章

1. [Visual Studio Code is unable to watch for file changes in this large workspace" (error ENOSPC)](https://code.visualstudio.com/docs/setup/linux#_visual-studio-code-is-unable-to-watch-for-file-changes-in-this-large-workspace-error-enospc)

使用vscode打开kubernetes工程, 保存文件时右下角总是跳出提示框, 显示监控的文件数太多, 点进去看会跳转到官方文档, 写的不错.

最主要是文档中还提到了每个监控文件消耗的内存占用, 可以计算自己的系统配置承受的上限, 对于认识sysctl也有帮助.

原本系统中此配置为8192, 太小了.

```
# sysctl -a | grep max_user
fs.epoll.max_user_watches = 784117
fs.inotify.max_user_watches = 8192
```
