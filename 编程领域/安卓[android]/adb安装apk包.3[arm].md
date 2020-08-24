# adb安装apk包

参考文章

1. [在 Android 模拟器上运行 ARM 应用](https://zhuanlan.zhihu.com/p/127016204)
2. [Android Studio 模拟器版本说明](https://developer.android.google.cn/studio/releases/emulator#30-0-0)
    - Android 11 系统映像能够在不影响整个系统的前提下, 直接将 ARM 指令转换成 x86 指令. 
    - 开发者无需搭建高负载的 ARM 环境即可执行 ARM 二进制文件并进行测试. 
    - x86：包括 x86 和 ARMv7 ABI。
    - x86_64：包括 x86、x86_64、ARMv7 和 ARM64 ABI。
3. [Android 设备的CPU类型(通常称为”ABIs”)](https://www.cnblogs.com/janehlp/p/7473240.html)
    - armeabi-v7a: 第7代及以上的 ARM 处理器。2011年15月以后的生产的大部分Android设备都使用它.
    - arm64-v8a: 第8代、64位ARM处理器，很少设备，三星 Galaxy S6是其中之一。
    - armeabi: 第5代、第6代的ARM处理器，早期的手机用的比较多。
    - x86: 平板、模拟器用得比较多。
    - x86_64: 64位的平板。
4. [如何查看Android手机CPU类型是armeabi，armeabi-v7a，还是arm64-v8a](https://blog.csdn.net/qq_36317441/article/details/89494686)
    - `adb shell getprop ro.product.cpu.abi`
5. [Android中的ABI](https://www.jianshu.com/p/170f65439844)
    - Android目前支持以下七种ABI：armeabi、armeabi-v7a、arm64-v8a、x86、x86_64、mips、mips64。
    - Android目前有以下七种cpu架构：ARMv5、ARMv7、ARMv8、x86、x86_64、MIPS和MIPS64。
    - 每种CPU架构都有其自己支持的ABIs
6. [获取 Android 11](https://developer.android.google.cn/preview/get)
7. [avd manger创建的虚拟机启动不起来，或者启动起来后黑屏](https://www.cnblogs.com/sy_test/p/12056040.html)
    - `intelhaxm-android.exe`
8. [Mac搭建Android环境](https://www.jianshu.com/p/25897462c090)
    - `Hardware_Accelerated_Execution_Manager`
9. [Mac OS X 10.9 使用 Hardware Accelerated Execution 之后死机问题](https://www.mobibrw.com/2013/756)

## 1. step 1

参考文章2说 Android 11 的任意版本, 以及 Android 9 的 x86 版本都可以安装并运行针对 Arm 平台编译的 APK 包. 

但我在实验的时候, 虽然可以安装, 但是一打开就闪退, 为此尝试了很多种方法.

首先是, `avdmanager create avd`时有一个参数为`--abi`, 可以为模拟器指定`abi`版本, 按照参考文章2中所说, `system-images;android-30;google_apis;x86_64`除了包含`x86`和`x86_64`的`abi`, 还包括`ARMv7`和`ARM64`, 所以我尝试在创建`avd`时指定这个参数为`armv7`, but...

```console
$ avdmanager create avd -f -n android-05 -d 9 -b 'ARMv7' -k 'system-images;android-30;google_apis;x86_64'
Valid ABIs: google_apis/x86_64==========] 100% Fetch remote repository...
Error: Invalid --abi ARMv7 for the selected package.
null
```

这个镜像里没有`ARMv7`的`abi`...

最先想到的是, 网上查找关于`INSTALL_FAILED_NO_MATCHING_ABIS`的问题时, 很多人提到了`armeabi-v7a`, 可能不应该写作`ARMv7`?

```console
$ avdmanager create avd -f -n android-05 -d 9 -b 'armeabi-v7a' -k 'system-images;android-30;google_apis;x86_64'
Valid ABIs: google_apis/x86_64==========] 100% Fetch remote repository...
Error: Invalid --abi armeabi-v7a for the selected package.
null
```

我把`x86`和`x86_64`的镜像按`armeabi-v7a`, `arm64-v8a`都试了一遍, apk还是没能正常启动.

然后我又尝试了如下

```console
$ avdmanager create avd -f -n android-05 -d 9 -b 'x86' -k 'system-images;android-30;google_apis;x86_64'
Valid ABIs: google_apis/x86_64==========] 100% Fetch remote repository...
Error: Invalid --abi x86 for the selected package.
null
```

`x86_64`中连`x86`的`abi`都不支持, 更别说`armv7`和`armv8`了.

我意识到, 参考文章2中所说的, `x86`镜像包括 x86 和 ARMv7 ABI, `x86_64`包括x86、x86_64、ARMv7 和 ARM64 ABI. 应该是底层能够支持, 但并不是通过在`--abi`参数中指定来使用的.

## step 2

然后考虑使用原生的支持 ARM 的 avd 镜像, 但是这些镜像启动实在在慢了, 一直卡在`Android` logo 那里. 参考文章7中有加速器可以安装, `sdkmanager`中其实也有相关工具.

```console
$ sdkmanager --list | grep -i Hardware
  extras;intel;Hardware_Accelerated_Execution_Manager | 7.5.1   | Intel x86 Emulator Accelerator (HAXM installer) | extras/intel/Hardware_Accelerated_Execution_Manager/
  extras;intel;Hardware_Accelerated_Execution_Manager                                      | 7.5.1        | Intel x86 Emulator Accelerator (HAXM installer)
```

`$ANDROID_SDK/extras/intel/Hardware_Accelerated_Execution_Manager`目录下有个`slient_install.sh`, 不过这东西貌似挺有风险, 没敢安装.

换用 genymotion 吧, 原生的放弃了.

...当然后来 genymotion 也放弃了, 换真姬吧.
