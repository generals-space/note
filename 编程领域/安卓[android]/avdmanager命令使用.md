删除一个已经存在的AVD设备

```console
$ avdmanager delete avd -n android-05
Deleting file /Users/general/.android/avd/android-05.ini
Deleting folder /Users/general/.android/avd/android-05.avd

AVD 'android-05' deleted.
```

`avdmanager list`相关.

- `list`              : Lists existing targets or virtual devices.
- `list avd`          : 查看已经存在的AVD设备(通过`create avd`子命令创建)
- `list target`       : 这个就是本地已经存在的`platform`列表, 比如通过`sdkmanager`安装了`platforms;android-27`和`platforms;android-30`, 这个命令就会显示出这两条.
    - 不过`create avd`时好像也没有参数可以指定`target`啊...
- `list device`       : 设备类型就是可选的机型, 像`Nexus 4, 5, 6, 7`, 或是`pixel 1, 2, 3`等, 也有TV, 平板等宽屏设备等.

