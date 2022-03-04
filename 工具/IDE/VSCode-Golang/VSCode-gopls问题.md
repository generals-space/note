参考文章

1. [The gopls server crashed 5 times in the last 3 minutes. The server will not be restarted](https://github.com/microsoft/vscode-go/issues/3026)
2. [The gopls server crashed 5 times in the last 3 minutes. ](https://github.com/microsoft/vscode-go/issues/2740)

本来在Mac下vmware的CentOS7中远程编写go代码时, gopls工作正常, 能够正常提供代码提示, 跳转的功能. 

但是某一个pause了CentOS7虚拟机并重启了mac, 再启动虚拟机时, 修改kubernetes代码, gopls在做预热的时候就崩溃出错了, vscode直接弹窗提示如下错误

```
The gopls server crashed 5 times in the last 3 minutes. The server will not be restarted
```

官方issue都说要把gopls更新到最新版`GO111MODULE=on go get golang.org/x/tools/gopls@latest`, 但是根本没用.

后来发狠把gopath下的bin和pkg目录都删了, 重新打开工程, 把所有代码工具都重新装了一遍, 可以了.

...真不让人消停.
