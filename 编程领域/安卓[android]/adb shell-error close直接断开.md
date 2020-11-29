# adb shell-error close直接断开

参考文章

1. [【Android测试】adb shell回车后出现 error closed的解决办法](https://my.oschina.net/u/199776/blog/330604)

`adb devices`查看目标设备的确在线, 然后执行`adb shell`, 就报如下错误

```
$ adb shell
error: close
```

然后再用`adb devices`查看时, 设备就`offline`了...

试了试参考文章1中说的断开重连, 执行过`adb kill-server`, 没啥用.

...重启电脑有用.
