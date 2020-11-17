# 安卓-scrcpy电脑控制手机.2.无线连接

参考文章

1. [如何在电脑上控制手机上所有的app软件操作？](https://www.zhihu.com/question/46795475)
2. [Genymobile/scrcpy](https://github.com/Genymobile/scrcpy/blob/master/README.zh-Hans.md)
    - 安装使用方法
3. [Mac - brew 安装缺少访问文件权限](https://www.cnblogs.com/yangtiancheng/p/9958167.html)
4. [MAC安装OpenXenManager管理Xenserver](https://blog.51cto.com/qiangsh/1731277)
5. [Developer Apple](https://blog.51cto.com/qiangsh/1731277)
    - Command Line Tools for Xcode 12

首先确认手机与电脑使用同一wifi, 在同一个局域网内.

使用 usb 将手机连接至电脑, 执行 `adb tcpip 5555`开启手机的网络adb功能, 然后断开与电脑的连接.

之后执行`adb connect 手机的IP`, 通过无线方式连接手机.

之后就可以运行 scrcpy 了.
