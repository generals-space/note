# adb安装apk包

参考文章

1. [INSTALL_FAILED_NO_MATCHING_ABIS解决办法](https://www.jianshu.com/p/a781ad09b092)
    - `build.gradle`
2. [Android模拟器下安装APP报INSTALL_FAILED_NO_MATCHING_ABIS错误解决方案](https://blog.csdn.net/stemq/article/details/51502759)
    - 使用Arm镜像
3. [Android：INSTALL_FAILED_NO_MATCHING_ABIS](https://vcoo.cc/blog/1223/)
4. [AVD 模拟 arm 那么卡你们都怎么开发的](https://www.v2ex.com/t/475597)
5. [在 Android 模拟器上运行 ARM 应用](https://blog.csdn.net/jILRvRTrc/article/details/105383037)
6. [Android Studio 模拟器版本说明](https://developer.android.google.cn/studio/releases/emulator#30-0-0)
    - Android 11 系统映像能够在不影响整个系统的前提下, 直接将 ARM 指令转换成 x86 指令. 
    - 开发者无需搭建高负载的 ARM 环境即可执行 ARM 二进制文件并进行测试. 

上一篇文章中尝试安装"夸克浏览器"到安卓模拟器失败, 报错如下

```console
$ adb install -r -t --no-streaming com.quark.browser_V3.8.1.125.apk
Performing Push Install
com.quark.browser_V3.8.1.125.apk: 1 file pushed, 0 skipped. 123.3 MB/s (48769324 bytes in 0.377s)
Failure [INSTALL_FAILED_NO_MATCHING_ABIS: Failed to extract native libraries, res=-113]
```

按照网上的文章分析来看, 基本可以确定是因为传统`apk`是针对手机平台ARM架构编译的, 而我用的模拟器镜像为`system-images;android-27;default;x86_64`, 所以是不兼容的.

关于这个问题, 我首先找到了参考文章1(有一大票文章是与这个相似的). 就是在自己做开发的时候, 编译选项里同时添加上对`x86`, `x86_64`和`arm`架构的支持.

这个我没实验过, 毕竟我不是要做安卓开发. 这种场景明显只能安装自己开发的apk包, 不能安装从应用商店下载下来的包.

然后找到参考文章2, 这个的建议就简单多了, 下载一个支持ARM架构的镜像就可以了.

但是我在实验时还是遇到问题, 我使用的镜像是`system-images;android-25;google_apis;arm64-v8a`, 但是启动了20多分钟, 都没能开机, 一直是Android Logo闪烁的界面.

放弃.

后来我还找到了参考文章4, 原来Arm镜像启动本来就是慢的, 运行也慢. 在想测从应用商店下载的包时, 需要使用真姬, 或是用`Genymotion + Genymotion-ARM-Translation`模拟器.

最终我又发现了参考文章5, 新版的安卓镜像已经支持直接在x86平台上直接运行arm程序.

于是我下载了`system-images;android-30;google_apis;x86_64`, 并使用ta创建了一个新的AVD.

```console
$ adb install --no-streaming ./com.quark.browser_V3.8.1.125.apk
Performing Push Install
./com.quark.browser_V3.8.1.125.apk: 1 file pushed, 0 skipped. 13.3 MB/s (48769324 bytes in 3.497s)
Success
```

这次果然安装成功了.

![](https://gitee.com/generals-space/gitimg/raw/master/0b39f5821f1f17ac9bc7242ac200b91f.png)
