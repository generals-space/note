# genymotion安装arm应用APK

1. [官方文档 Deploying an application](https://docs.genymotion.com/desktop/3.0/03_Virtual_devices/032_Deploying_an_application.html#deploying-an-application)
    - `Genymotion Desktop`创建的模拟器CPU架构为x86, 要运行 ARMv7 的应用需要安装`ARM translation tool`, 且必须与模拟器中系统的版本相匹配.
2. [m9rco/Genymotion_ARM_Translation](https://github.com/m9rco/Genymotion_ARM_Translation)
    - 支持的 Android 版本: 4.3, 4.4, 5.1, 6.0, 7.x, 8.0
3. [How to unroot a Genymotion virtual device?](https://support.genymotion.com/hc/en-us/articles/360003125397-How-to-unroot-a-Genymotion-virtual-device-)
    - genymotion 提供的模拟器 4.2+ 以后默认都是 root 过的, 而且没有办法关闭.

genymotion 需要安装 VirtualBox, 如果没有安装, 启动时会提示.

genymotion 个人使用免费, 但是需要登录.

genymotion 提供的安卓镜像与`Android Studio`不通用, 需要重新下载.

安装 APK, 把 zip 包拖到模拟器中即可. 

由于一般 APK 是针对 ARM 平台编译的, 所以直接拖进去可能会弹出如下提示.

![](https://gitee.com/generals-space/gitimg/raw/master/0ee80549894cad79aebfab143471abad.png)

安装 Arm 转换器, 只要把 zip 包拖到模拟器中即可, 会提示重启.

![](https://gitee.com/generals-space/gitimg/raw/master/2bb17bb3e6b1393ca9889f74d08a82d9.png)

![](https://gitee.com/generals-space/gitimg/raw/master/3ccd12abb109748b9fcb0b452f24ca94.png)

目前来说, 安卓 6.0 的镜像是兼容性最好的, 但仍然不够稳定, 连夸克和中移掌上营业厅都装不了. 算了, 换用真姬了.
