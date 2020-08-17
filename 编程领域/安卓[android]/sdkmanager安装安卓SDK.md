# sdkmanager安装安卓SDK

参考文章

1. [`android studio`官方网站](https://developer.android.com/studio)
    - , `sdkmanager`在页面末尾部分, 可以单独下载.
2. [Android SDK Command line tools运行sdkmanager报告Warning: Could not create settings错误信息解决方法](https://blog.csdn.net/zhufu86/article/details/106747556)
    - `sdkmanager --sdk_root=c:\androidsdk "cmdline-tools;latest"`, 其实没有必要
3. [Android Command line tools sdkmanager always shows: Warning: Could not create settings](https://stackoverflow.com/questions/60440509/android-command-line-tools-sdkmanager-always-shows-warning-could-not-create-se)
    - `tools` -> `$ANDROID_SDK_ROOT/cmdline-tools/tools`
4. [如何使用Android SDK Manager下载 SDK](https://www.cnblogs.com/Caiyilong/p/8559394.html)
5. [`android studio`官方网站 - sdkmanager](https://developer.android.com/studio/command-line/sdkmanager)
6. [mac android sdk manager](https://www.jianshu.com/p/37732d00c115)
    - win的`sdkmanager.exe`有GUI, mac下的`sdkmanager`只有命令行.
    - `mac android sdk manager`的下载及使用方法

最近想测试一下使用`appium`做安卓的自动化脚本, 安装`appium-desktop`后不够, 还要安装`jdk`, `android sdk`, 这两都是比较庞大的家伙, 我是很不情愿的.

我记得当初在大学时写安卓大作业的时候就费了老大劲才把环境搭起来.

现在2020年, 网上的文章仍然都是通过`android studio`, `eclipse`等IDE内置的插件去下载`android sdk`, `android studio`的插件是`sdk manager`, `eclipse`的叫`ADT`.

不过我只想要`android sdk`而已, 并想安装什么IDE. 毕竟我就想写个python脚本, 为什么要下载`Pycharm`呢?

## 安装 sdkmansger

网上一些提供 sdk 合集的中文网站都已经很旧了, 如下

1. [android dev tools](https://www.androiddevtools.cn/)
2. [](http://tools.android-studio.org/index.php/sdk/)

所以还是使用官方网站吧.

参考文章1的`Command line tools only`部分可以单独下载`sdkmanager`工具. 解压后, `tools/bin/sdkmanager`为目标可执行文件.

需要注意的是, 解压出来的`tools`目录不能随便放, 必须放在`$ANDROID_HOME/cmdline-tools/tools/`目录, 否则执行`sdkmanager --version`时会出错.

```
$ sdkmanager --version
Warning: Could not create settings
java.lang.IllegalArgumentException
    at com.android.sdklib.tool.sdkmanager.SdkManagerCliSettings.<init>(SdkManagerCliSettings.java:428)
    at com.android.sdklib.tool.sdkmanager.SdkManagerCliSettings.createSettings(SdkManagerCliSettings.java:152)
    at com.android.sdklib.tool.sdkmanager.SdkManagerCliSettings.createSettings(SdkManagerCliSettings.java:134)
    at com.android.sdklib.tool.sdkmanager.SdkManagerCli.main(SdkManagerCli.java:57)
    at com.android.sdklib.tool.sdkmanager.SdkManagerCli.main(SdkManagerCli.java:48)
```

参考文章2给出了解决的过程, 但是稍微有些繁琐. ta主要想先下载`sdkmanager`, 然后用`sdkmanager --sdk_root=c:\androidsdk "cmdline-tools;latest"`重新构建`ANDROID_HOME`的目录.

参考文章2借鉴了参考文章3, 就是把`tools`放在`$ANDROID_HOME/cmdline-tools/tools/`就可以解决问题了.

## 安装 adb

然后用`sdkmanager`安装`android sdk`.

windows下的`sdkmanager.exe`是一个GUI界面, 打勾选择就可以了, 但是mac下就只有命令行.

执行如下命令, 查看可以安装的包.

```
$ sdkmanager --list
[=======================================] 100% Computing updates...
Available Packages:
  Path                                                                                     | Version      | Description
  -------                                                                                  | -------      | -------
  add-ons;addon-google_apis-google-15                                                      | 3            | Google APIs
  add-ons;addon-google_apis-google-16                                                      | 4            | Google APIs
  add-ons;addon-google_apis-google-17                                                      | 4            | Google APIs
  add-ons;addon-google_apis-google-18                                                      | 4            | Google APIs
  add-ons;addon-google_apis-google-19                                                      | 20           | Google APIs
  add-ons;addon-google_apis-google-21                                                      | 1            | Google APIs
  add-ons;addon-google_apis-google-22                                                      | 1            | Google APIs
  add-ons;addon-google_apis-google-23                                                      | 1            | Google APIs
  add-ons;addon-google_apis-google-24                                                      | 1            | Google APIs
  build-tools;19.1.0                                                                       | 19.1.0       | Android SDK Build-Tools 19.1
  build-tools;20.0.0                                                                       | 20.0.0       | Android SDK Build-Tools 20
  build-tools;21.1.2                                                                       | 21.1.2       | Android SDK Build-Tools 21.1.2
  build-tools;22.0.1                                                                       | 22.0.1       | Android SDK Build-Tools 22.0.1
  build-tools;23.0.1                                                                       | 23.0.1       | Android SDK Build-Tools 23.0.1
  build-tools;23.0.2                                                                       | 23.0.2       | Android SDK Build-Tools 23.0.2
  build-tools;23.0.3                                                                       | 23.0.3       | Android SDK Build-Tools 23.0.3
```

按照参考文章4中的勾选项, 目前就只装了一个`platforms;android-30`, 参考文章5中有安装命令, 如下.

```
sdkmanager "platform-tools" "platforms;android-30"
```

这会在`$ANDROID_HOME`目录下生成`platforms`和`platform-tools`子目录, `platforms`目录有`android-30`, `platform-tools`下有`adb`命令.

