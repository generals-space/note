# Mac-Finder强制使用多标签代替多窗口

参考文章

1. [将所有Finder窗口合并到Mac OS X中的单个标签窗口中](https://mos86.com/20543.html)
    - Finder -> 窗口 -> 合并所有窗口, 可以将多个窗口的 finder 合并到单窗口下, 但治标不治本.
2. [Force Finder to use only a single window?](https://apple.stackexchange.com/questions/342798/force-finder-to-use-only-a-single-window)

在Mac中, 经常需要在各种应用中打开一个已知文件所在的目录, 比如 vscode 的 `Reveal in Finder`, 钉钉/微信打开已下载文件的目录等.

![](https://gitee.com/generals-space/gitimg/raw/master/30f7333405f473cf211a0903f62a7a4d.png)

![](https://gitee.com/generals-space/gitimg/raw/master/41b883ab262ace65743de19c9353cfab.png)

但是这些操作总是打开一个新的 Finder 窗口, 而不是在一个窗口的多个标签中打开, 使用时总是需要使用**Command + 反引号**进行切换, 很麻烦.

之前找到一种方法, 就是在 Finder 的偏好设置中, 在"通用"选项卡下, 勾选"在标签页(而不是新窗口)中打开文件夹", 但是收效甚微...

![](https://gitee.com/generals-space/gitimg/raw/master/9d8013e7942e91536983d379e1ad02a9.png)

后来找到了参考文章2, 打开 系统偏好设置 -> 程序坞(Dock) -> 打开文稿时首先标签页, 该下拉菜单默认选项为"仅在全屏幕视图下", 将其修改成"始终".

![](https://gitee.com/generals-space/gitimg/raw/master/e82cb98d52a3582cfd5e8d7c1f5de791.png)

现在上述的那些操作都会在新的标签页中打开了.
