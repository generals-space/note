# adb shell

`adb shell xxx`可以在安卓命令行执行 linux 命令, 如下

```
adb shell ls /
```

不过不是交互式的, 交互式命令行还是直接执行`adb shell`吧.

该命令行权限有限, 用的是`shell`命令.

```bash
$ adb shell
127|chiron:/ $ whoami
shell
```
