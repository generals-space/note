# adb命令使用

参考文章

1. [android SDK安装和环境变量配置和adb工具的使用](https://blog.csdn.net/g695144224/article/details/50844446)
2. [adb push 失败提示 ‘Read-only file system’](https://www.jianshu.com/p/eca9a8ad4996)
3. [Appium的使用方法](https://www.cnblogs.com/weizhibin1996/p/9254261.html)
    - 获取系统版本: `adb shell getprop ro.build.version.release`
    - 获取系统api版本: `adb shell getprop ro.build.version.sdk`

MacOS: 10.15.4 (19E287)
手机: Vivo Y85A, 安卓版本: 8.1.0

```console
$ adb version
Android Debug Bridge version 1.0.41
Version 30.0.4-6686687
Installed as /Users/general/Public/android//platform-tools/adb
```

## 

`adb`是类似`kubectl`一样的客户端工具, 可以用来调试android模拟器和真姬. 

在使用`sdkmanager`安装`platform`相关包之后, 在`$ANDROID_HOME/platform-tools`会出现`adb`命令, 将此目录添加到`PATH`环境变量中即可.

真姬的话, 需要手机开启"开发者选项", 同时开启"USB调试". usb线就相当于网线, 网上也有不通过usb线, 而是使用 wifi 进行调试的方法.

开发者选项的开启方式在不同品牌的手机是不同的, 可以在网上搜索相关方法. Vivo的手机是"设置" -> "更多设置" -> "关于手机" -> "软件版本号", 连续点击"软件版本号"7次, 就可以了. "开发者选项"会出现在"更多设置"下, 注意同时开启"开发者选项"中的"USB 调试"即可.

然后, 在命令行执行`adb devices`

```console
$ adb devices
List of devices attached
1891dc13	device
```

参考文章3给出了两条命令用来查看手机的系统信息, 很有用.

```console
$ adb shell getprop ro.build.version.release
8.1.0
$ adb shell getprop ro.build.version.sdk
27
```

## push/pull

`adb`有`push`和`pull`两个命令可以在电脑和手机之间互传文件.

```console
$ adb -s 1891dc13 pull /storage/emulated/0/Download/xxx ./
/storage/emulated/0/Download/xxx: 1 file pulled, 0 skipped. 0.0 MB/s (80 bytes in 0.012s)
$ adb -s 1891dc13 push ssh.log /storage/emulated/0/Download/
ssh.log: 1 file pushed, 0 skipped. 9.0 MB/s (20961 bytes in 0.002s)
```

## adb shell 进入命令行

原来只要`adb devices`中有设备显示, `adb shell`就真的能进入ta的命令行...

```console
$ adb shell
PD1730:/ $ ls
ls: ./verity_key: Permission denied
## ...省略
ls: ./donuts_key: Permission denied
acct   bt_firmware cache   config data         dev etc      mnt persist res  sbin   storage system     vendor
athena bugreports  charger d      default.prop dsp firmware oem proc    root sdcard sys     tombstones
```

如果进入`/storage/emulated/0`, 就会得到和使用系统内置的文件浏览器一样的内容了.

```console
PD1730:/storage/emulated/0 $ ls
Alarms   Mob           Podcasts  UCDownloads attribution  cache             libs       qmt      system        xianyu
Android  Movies        QQBrowser Xiaomi      backup       com.tencent.tim   mipush     setup    tbs           阅图锁屏
DCIM     Music         Quark     alipay      backups      huanju            msc        sitemp   tencent
Download Notifications Ringtones amap        baidu        i\ Music          netease    snowball tencentmapsdk
MQ       Pictures      TurboNet  at          baiduDuerTTS internetComponent pluginInfo sogou    vipc

PD1730:/storage/emulated/0 $ ls -al
total 324
drwxrwx--x 74 root sdcard_rw 4096 2020-08-17 15:49 .
drwx--x--x  4 root sdcard_rw 4096 1972-04-29 11:54 ..
drwxrwx--x  2 root sdcard_rw 4096 2020-08-01 11:11 Alarms
drwxrwx--x  5 root sdcard_rw 4096 2020-08-17 14:01 Android
drwxrwx--x  7 root sdcard_rw 4096 2020-08-03 19:36 DCIM
drwxrwx--x  3 root sdcard_rw 4096 2020-08-17 22:10 Download
```

可以看到, 文件属主其实还是`root`, 不过所属组`sdcard_rw`还是有权限读写的.

我试了试, 在`shell`命令行里还是可以写入文件信息的, 很方便.

