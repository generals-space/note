# adb安装apk包

参考文章

1. [如何使用adb命令安装APK](https://www.jianshu.com/p/a64f6a476184)
2. [installing APK via ADB hangs the emulator](https://stackoverflow.com/questions/21826527/installing-apk-via-adb-hangs-the-emulator/21827103)
    - Android Virtual Device Manager-->select emulator-->edit--> Internal Storage-->increase size to >512-->OK
3. [Disable Streamed Installation in Android Studio](https://androidforums.com/threads/disable-streamed-installation-in-android-studio.1318205/)
    - `adb install -r -t --no-streaming app-debug.apk`

查看系统中已经安装了的包

```console
$ adb shell pm list packages
package:com.android.email
package:com.android.music
package:com.android.phone
package:com.android.shell
```

网上大部分文章直接使用`install`子命令就可以, 顶多再加个`-r`, 但我在实际测试时, `install`会卡住.

```console
$ adb install ./com.quark.browser_V3.8.1.125.apk
Performing Streamed Install
^C
```

等了很久都不成功, 也不出错.

参考文章2中的问题与我这个很相似, 高票回答说是因为内置存储(internal storage)太小, 使用`avdmanager`调整到512M就可以正常安装了.

但是我使用`avdmanager`从命令行创建的虚拟机默认就是512M的, SDcard, 应该算是内置存储, 而且`avdmanager create`只有`--sdcard`参数, 没有什么内置存储.

后来找到了参考文章3, 需要加一个`--no-streaming`参数.

```console
$ adb install -r -t --no-streaming com.quark.browser_V3.8.1.125.apk
Performing Push Install
com.quark.browser_V3.8.1.125.apk: 1 file pushed, 0 skipped. 123.3 MB/s (48769324 bytes in 0.377s)
Failure [INSTALL_FAILED_NO_MATCHING_ABIS: Failed to extract native libraries, res=-113]
```

算是能安装了, 但是安装出错了.
