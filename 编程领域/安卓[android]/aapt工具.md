# aapt工具

参考文章

1. [Python+appium自动化实例（一）：使用Python3+appium实现自动收取支付宝蚂蚁森林能量](https://www.cnblogs.com/deliaries/p/12410835.html)
    - 本文示例主要是以真姬为主, 虽然要安装安卓模拟器, 也只需要`platform-tools`即可(主要是`adb`工具)
    - `aapt`工具在`build-tools`包下(需要使用`sdkmanager`安装)

```
$ aapt dump badging XXX.apk
package: name='com.unionpay' versionCode='233' versionName='8.0.3' platformBuildVersionName='8.0.3' compileSdkVersion='28' compileSdkVersionCodename='9'
sdkVersion:'16'
targetSdkVersion:'28'
uses-permission: name='android.permission.INTERNET'
uses-permission: name='android.permission.USE_BIOMETRIC'
uses-library-not-required:'org.simalliance.openmobileapi'
uses-library-not-required:'org.apache.http.legacy'
launchable-activity: name='com.unionpay.activity.UPActivityWelcome'  label='' icon=''
feature-group: label=''
  uses-feature-not-required: name='android.hardware.camera'
  uses-feature-not-required: name='android.hardware.camera.autofocus'
  uses-feature-not-required: name='android.hardware.nfc.hce'
  uses-feature: name='android.hardware.bluetooth'
  uses-implied-feature: name='android.hardware.bluetooth' reason='requested android.permission.BLUETOOTH permission, requested android.permission.BLUETOOTH_ADMIN permission, and targetSdkVersion > 4'
  uses-feature: name='android.hardware.wifi'
  uses-implied-feature: name='android.hardware.wifi' reason='requested android.permission.ACCESS_WIFI_STATE permission, and requested android.permission.CHANGE_WIFI_STATE permission'
provides-component:'payment'
main
other-activities
other-receivers
other-services
supports-screens: 'small' 'normal' 'large' 'xlarge'
supports-any-density: 'true'
densities: '160' '240' '320' '480' '640' '65534'
native-code: 'armeabi-v7a'
```

- `package: name`: `appium`中, 可用于`capabilities`的`appPackage`字段.
- `launchable-activity: name`: `appium`中, 可用于`capabilities`的`appActivity`字段.
- `native-code`: 该 apk 编译时针对的CPU架构 
- `targetSdkVersion`: 编写代码时使用的API版本.???
