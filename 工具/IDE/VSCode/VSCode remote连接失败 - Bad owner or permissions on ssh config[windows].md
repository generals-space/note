# VSCode remote连接失败 - Bad owner or permissions on ssh config[windows]

参考文章

1. [Bad owner or permissions on C:\\Users\\user/.ssh/config ＞ 过程试图写入的管道不存在。](https://blog.csdn.net/qq_44944580/article/details/137611875)

## 问题描述

windows 上使用 vscode 远程连接 linux 主机, 配置 ssh/config 主机别名, 但显示该文件权限不正确.

...windows上文件权限太烦, 按照参考文章1中所说, 使用 git 工具带的 ssh.exe 就可以.

VSCODE设置

File -> 偏好Preference -> Setting
搜索`remote.ssh.path`

设置`D:\Program Files\Git\usr\bin\ssh.exe`或者`C:\Program Files\Git\usr\bin\ssh.exe`

然后重启vscode再连就行了.
