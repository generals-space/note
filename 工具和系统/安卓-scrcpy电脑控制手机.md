# 安卓-scrcpy电脑控制手机

参考文章

1. [如何在电脑上控制手机上所有的app软件操作？](https://www.zhihu.com/question/46795475)
2. [Genymobile/scrcpy](https://github.com/Genymobile/scrcpy/blob/master/README.zh-Hans.md)
    - 安装使用方法
3. [Mac - brew 安装缺少访问文件权限](https://www.cnblogs.com/yangtiancheng/p/9958167.html)
4. [MAC安装OpenXenManager管理Xenserver](https://blog.51cto.com/qiangsh/1731277)
5. [Developer Apple](https://blog.51cto.com/qiangsh/1731277)
    - Command Line Tools for Xcode 12

本来想 root 以后试着编译下 vnc server 的, 后来找到了参考文章1...

`scrcpy`是 genymonbile 开源的工具, 无需 root, 无需手机上安装使用 app, 只在电脑上安装`scrcpy`并拥有`adb`工具即可.

按照参考文章2中的文档, 通过`brew`安装`scrcpy`, 结果出现如下问题.

```console
$ brew install scrcpy
Updating Homebrew...
Error: The following directories are not writable by your user:
/usr/local/lib/pkgconfig
/usr/local/share/man/man8

You should change the ownership of these directories to your user.
  sudo chown -R $(whoami) /usr/local/lib/pkgconfig /usr/local/share/man/man8

And make sure that your user has write permission.
  chmod u+w /usr/local/lib/pkgconfig /usr/local/share/man/man8
```

找到参考文章3, 其实这个问题很好解决, 上面报错的两个目录`/usr/local/lib/pkgconfig`, `/usr/local/share/man/man8`的属主为`root`, 改成登录用户就行了.

> 我有试过使用`sudo`执行`brew`, 但被提示不安全而拒绝执行, 所以改目标目录的权限是最方便的.

------

然后继续执行`brew`安装, 结果又出了如下错误

```
$ brew install scrcpy
Updating Homebrew...
Error: The following formula
  [#<Dependency: "python" []>, #<Options: []>]
cannot be installed as binary package and must be built from source.
Install the Command Line Tools:
  xcode-select --install

```

参考文章4中的情况与我遇到的一样, 我执行`xcode-select --install`也失败了, 所以就去苹果官网上再下载的 xcode dmg安装包安装的.

之后再执行`brew install`就可以了, 花的时间还挺久的.

------

关于使用, 手机连上电脑, 打开 usb 调试, 然后命令行执行`scrcpy`, 就能看到手机的界面了, 非常方便.

鼠标右键是返回, 中键是桌面(窗口上没有按键, 所以最好还是有个鼠标来操作)

