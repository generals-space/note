# Mac-Finder显示隐藏文件

参考文章

1. [如何查看Mac上的隐藏文件和文件夹？](https://www.macbl.com/article/tips/1843)
2. [MacBook Pro 设置Finder显示隐藏文件](https://blog.csdn.net/qq_35624642/article/details/82764545)
3. [在 Mac OS X 上，如何隐藏文件和查看隐藏文件](https://www.helplib.com/Mac/article_11284)
    - Finder界面 `Command + Shift + .` 切换显示/隐藏模式
    - 命令行输入 `chflag hidden|nohidden 文件路径` 设置文件显示与隐藏属性

很多时候在终端能看到的文件, 在finder里没有显示. 比如在docker中用wget下载的文件, 或是从windows通过scp推过来的文件, 都是这样.

但是终端显示文件的权限并没有什么不同.

解决方法是开启finder显示隐藏的文件.

终端执行如下命令

```
defaults write com.apple.finder AppleShowAllFiles TRUE
```

然后执行`killall`关闭finder(GUI上好像没有让finder退出的按钮啊...)

```
killall Finder
```

然后再打开finder, 就可以显示隐藏文件了.

下面的命令可以再将隐藏文件的显示关闭, 当然也需要重启finder

```
defaults write com.apple.finder AppleShowAllFiles FALSE 
```
