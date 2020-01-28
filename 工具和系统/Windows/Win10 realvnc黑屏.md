# Win10 realvnc黑屏

参考文章

1. [Display issues when connecting to VNC Server running on Windows 10](https://help.realvnc.com/hc/en-us/articles/360004012211-Display-issues-when-connecting-to-VNC-Server-running-on-Windows-10)

在win10上安装了realvnc server破解版 6.1.1, 然后在MacOS上用vnc viewer连接, 使用了一段时间好好的, 突然有一天连接就显示黑屏了.

鼠标指针变成了小方块, 虽然是黑屏, 但还是可以点击, 只是viewer界面上无法显示了.

后来在参考文章1中找到了答案, 是因为server找不到monitor设备了(正好之前更新了一下显卡驱动, 可能是那时出的问题). 右击任务栏图标 -> Options -> Expert -> CaptureMethod, 修改为1(默认为0), 然后重启server, 再次连接正常.

这里的0, 1应该是指笔记本原屏蔽与扩展屏之间的数字的, 不过在"显示"设置中, 是用1和2区分的.
