# Mac调整Launchpad(启动台)图标大小

参考文章

1. [Mac Launchpad(启动台）图标大小调整](https://www.jianshu.com/p/4c3d11eebb6d)

MacOS: Mojave 10.14.5 (18F132)

命令行下执行.

1、调整每一列显示图标数量，10表示每一列显示10个，比较不错，可根据个人喜好进行设置。 

```
defaults write com.apple.dock springboard-rows -int 10
```

2、调整多少行显示图标数量，这里我用的是8 

```
defaults write com.apple.dock springboard-rows -int 8
```

3、重置Launchpad

```
defaults write com.apple.dock ResetLaunchPad -bool TRUE
```

4、重启Dock

```
killall Dock
```

注意

1. 原文中的`killall Dock`好像有不可见的非法字符, 执行时会出现`killall Dock: command not found`的错误.
2. 这个调整操作貌似会使原本Launchpad的文件夹中的东西都跑出来.


