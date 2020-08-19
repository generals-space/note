# sdkmanager --list各包的作用

参考文章

1. [Android SDK Command line tools运行sdkmanager报告Warning: Could not create settings错误信息解决方法](https://blog.csdn.net/zhufu86/article/details/106747556)
    - `sdkmanager --sdk_root=c:\androidsdk "cmdline-tools;latest"`, 其实没有必要

```
$ sdkmanager --list
[=======================================] 100% Computing updates...
Available Packages:
  Path                                                                                     | Version      | Description
  -------                                                                                  | -------      | -------
  add-ons;addon-google_apis-google-24                                                      | 1            | Google APIs

  build-tools;27.0.0                                                                       | 27.0.0       | Android SDK Build-Tools 27
  build-tools;27.0.1                                                                       | 27.0.1       | Android SDK Build-Tools 27.0.1
  build-tools;27.0.2                                                                       | 27.0.2       | Android SDK Build-Tools 27.0.2
  build-tools;27.0.3                                                                       | 27.0.3       | Android SDK Build-Tools 27.0.3

  cmake;3.10.2.4988404                                                                     | 3.10.2       | CMake 3.10.2.4988404
  cmake;3.6.4111459                                                                        | 3.6.4111459  | CMake 3.6.4111459
  cmdline-tools;1.0                                                                        | 1.0          | Android SDK Command-line Tools
  cmdline-tools;2.1                                                                        | 2.1          | Android SDK Command-line Tools
  cmdline-tools;latest                                                                     | 2.1          | Android SDK Command-line Tools (latest)
  emulator                                                                                 | 30.0.12      | Android Emulator
  extras;android;m2repository                                                              | 47.0.0       | Android Support Repository
  extras;google;auto                                                                       | 1.1          | Android Auto Desktop Head Unit emulator
  extras;google;google_play_services                                                       | 49           | Google Play services
  extras;google;instantapps                                                                | 1.9.0        | Google Play Instant Development SDK
  extras;google;m2repository                                                               | 58           | Google Repository
  extras;google;market_apk_expansion                                                       | 1            | Google Play APK Expansion library
  extras;google;market_licensing                                                           | 1            | Google Play Licensing Library
  extras;google;simulators                                                                 | 1            | Android Auto API Simulators
  extras;google;webdriver                                                                  | 2            | Google Web Driver
  extras;intel;Hardware_Accelerated_Execution_Manager                                      | 7.5.1        | Intel x86 Emulator Accelerator (HAXM installer)
  extras;m2repository;com;android;support;constraint;constraint-layout-solver;1.0.0        | 1            | Solver for ConstraintLayout 1.0.0
  extras;m2repository;com;android;support;constraint;constraint-layout-solver;1.0.1        | 1            | Solver for ConstraintLayout 1.0.1
  extras;m2repository;com;android;support;constraint;constraint-layout-solver;1.0.2        | 1            | Solver for ConstraintLayout 1.0.2
  extras;m2repository;com;android;support;constraint;constraint-layout;1.0.0               | 1            | ConstraintLayout for Android 1.0.0
  extras;m2repository;com;android;support;constraint;constraint-layout;1.0.1               | 1            | ConstraintLayout for Android 1.0.1
  extras;m2repository;com;android;support;constraint;constraint-layout;1.0.2               | 1            | ConstraintLayout for Android 1.0.2
  ndk-bundle                                                                               | 21.3.6528147 | NDK
  ndk;16.1.4479499                                                                         | 16.1.4479499 | NDK (Side by side) 16.1.4479499
  ndk;17.2.4988734                                                                         | 17.2.4988734 | NDK (Side by side) 17.2.4988734
  ndk;18.1.5063045                                                                         | 18.1.5063045 | NDK (Side by side) 18.1.5063045
  ndk;19.2.5345600                                                                         | 19.2.5345600 | NDK (Side by side) 19.2.5345600
  ndk;20.0.5594570                                                                         | 20.0.5594570 | NDK (Side by side) 20.0.5594570
  ndk;20.1.5948944                                                                         | 20.1.5948944 | NDK (Side by side) 20.1.5948944
  ndk;21.0.6113669                                                                         | 21.0.6113669 | NDK (Side by side) 21.0.6113669
  ndk;21.1.6352462                                                                         | 21.1.6352462 | NDK (Side by side) 21.1.6352462
  ndk;21.2.6472646                                                                         | 21.2.6472646 | NDK (Side by side) 21.2.6472646
  ndk;21.3.6528147                                                                         | 21.3.6528147 | NDK (Side by side) 21.3.6528147
  patcher;v4                                                                               | 1            | SDK Patch Applier v4
  platform-tools                                                                           | 30.0.4       | Android SDK Platform-Tools

  platforms;android-27                                                                     | 3            | Android SDK Platform 27

  skiaparser;1                                                                             | 3            | Layout Inspector image server for API 29-30
  
  sources;android-27                                                                       | 1            | Sources for Android 27
  
  system-images;android-27;android-tv;x86                                                  | 8            | Android TV Intel x86 Atom System Image
  system-images;android-27;default;x86                                                     | 1            | Intel x86 Atom System Image
  system-images;android-27;default;x86_64                                                  | 1            | Intel x86 Atom_64 System Image
  system-images;android-27;google_apis;x86                                                 | 10           | Google APIs Intel x86 Atom System Image
  system-images;android-27;google_apis_playstore;x86                                       | 3            | Google Play Intel x86 Atom System Image
  system-images;android-Q;android-tv;x86                                                   | 1            | Android TV Intel x86 Atom System Image
```

假设最初使用的是从`android studio`下载的`sdkmanager`包, 并使用`sdkmanager`这个命令下载上面的包.

- `cmdline-tools`部分, 如果安装这个包, 会在`ANDROID_HOME`目录下创建`cmdline-tools`子目录. 
    - 参考文章1中有先下载`sdkmanager`, 将其随便放在一个位置, 使用`sdkmanager`下载`cmdline-tools`包的方法. 其实只要把从`android studio`下载的`sdkmanager`包(解压开是`tools`目录)放入`ANDROID_HOME/cmdline-tools/tools`就可以了, 可见ta们两个作用是相同的.
- `platform-tools`部分, 会在`ANDROID_HOME`目录下创建`platform-tools`子目录, 包含`adb`工具.
- `system-images;`部分, 这里是创建安卓模拟器需要使用的ISO镜像, 按照类型可以分为"安卓TV", "default"版, "google_apis"版
    - 估计华为只能使用"default"版, 不能使用谷歌服务.
    - 由于是在PC上的模拟器, 所以都是"Intel x86"的架构, 不过也有"ARM 64 v8a"的镜像, 不知道在PC模拟器能不能启动, 理论上不行...
    - 27版本的还少了一个"android-wear"类型, 可穿戴设备, 应该是谷歌眼镜, 小天才电话手表之类的使用的系统.
