# avdmanager创建模拟器

参考文章

1. [AVD Manager 模拟器使用](https://www.cnblogs.com/guo2733/p/10558462.html)
    - win下 AVD Manager 的使用方法, 参数设置, 很详细
2. [Android命令行启动模拟器](https://blog.csdn.net/u010359739/article/details/54708960)
    - `android`命令没见过, 创建`AVD`虚拟机需要使用`avdmanager`

## 创建

与`sdkmanager`一样, MacOS的`avdmanager`在同一目录. 

而且win下的`avdmanager.exe`是GUI程序, MacOS下也是只有命令行界面, 不过没什么大不了, 我更喜欢命令行.

```console
$ sdkmanager "system-images;android-27;default;x86_64"
[=======================================] 100% Unzipping... x86_64/kernel-ranchu
```

```console
$ avdmanager list
Available Android Virtual Devices:
Parsing legacy package: /Users/general/Public/android/cmdline-tools/toolsParsing /Users/general/Public/android/platform-tools/package.xmlParsing /Users/general/Public/android/platforms/android-30/package.xmlAvailable devices definitions:
id: 0 or "tv_1080p"
    Name: Android TV (1080p)
    OEM : Google
    Tag : android-tv
---------
id: 1 or "tv_720p"
    Name: Android TV (720p)
    OEM : Google
    Tag : android-tv
---------
id: 2 or "wear_round"
    Name: Android Wear Round
    OEM : Google
    Tag : android-wear
## ...省略
---------
id: 21 or "pixel_3"
    Name: Pixel 3
    OEM : Google
```

创建 AVD 虚拟机文件需要事先安装`emulator`, 否则会出错

```console
$ avdmanager create avd -n android-01 -d 21 -k 'system-images;android-27;default;x86_64'
Auto-selecting single ABI x86_64========] 100% Fetch remote repository...
Error: "emulator" package must be installed!
null
```

使用如下命令安装`emulator`工具.

```console
$ sdkmanager "emulator"
[===                                    ] 10% Downloading emulator-darwin-646632
```

安装完成后, 会在`ANDROID_HOME`下创建一个`emulator`子目录, 下面有`emulator`可执行文件.

然后再创建就可以了.

```console
$ avdmanager create avd -n android-01 -d 21 -k 'system-images;android-27;default;x86_64'
Auto-selecting single ABI x86_64========] 100% Fetch remote repository...
~/Public/android/platform-tools
$ avdmanager list avd
Available Android Virtual Devices:
    Name: android-01
  Device: pixel_3 (Google)
    Path: /Users/general/.android/avd/android-01.avd
  Target: Default Android System Image
          Based on: Android 8.1 (Oreo) Tag/ABI: default/x86_64
  Sdcard: 512 MB
```

之后需要使用`emulator`命令启动 AVD 虚拟机, ta只有一个参数, 就是`AVD`虚拟机的名称(感觉`emulator`就像`vmware player`一样, 只能用来启动, 不能用来创建)

```console
$ emulator -avd android-01
Your emulator is out of date, please update by launching Android Studio:
 - Start Android Studio
 - Select menu "Tools > Android > SDK Manager"
 - Click "SDK Tools" tab
 - Check "Android Emulator" checkbox
 - Click "OK"

emulator: INFO: boot completed
emulator: INFO: boot time 26826 ms
emulator: Increasing screen off timeout, logcat buffer size to 2M.
emulator: Revoking microphone permissions for Google App.

```

同时桌面上会弹出模拟器界面, 如下

![](https://gitee.com/generals-space/gitimg/raw/master/fa36095a78ce79fa9840f730a19a0c97.png)

然后`adb shell`可以进入模拟器终端, 完成!

## 

使用`-list-avds`子命令可以查看当前已启动的虚拟机.

```console
$ emulator -list-avds
android-01
```
