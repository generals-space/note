# mac下usb真机调试 adb devices 获取不到机器

参考文章

1. [mac 电脑 adb devices获取不到机器](https://www.jianshu.com/p/75aec58c5512)
2. [解决mac下adb devices命令找不到设备](https://blog.csdn.net/linhunshi/article/details/72866345)

MacOS: 10.15.4 (19E287)
手机: Vivo Y85A, 安卓版本: 8.1.0

我将手机的"开发者选项"打开, usb连接Mac, 但是使用`adb devices`检测不到机器.

```console
$ adb devices
List of devices attached
```

参考文章1和2都有说了大致相同的方法, 不过我的情况是...忘了打开USB调试...

