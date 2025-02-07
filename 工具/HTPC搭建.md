参考文章

1. [Google浏览器开机自启,并自动打开全屏模式（windows）](https://blog.csdn.net/qq_38992048/article/details/129068177)
2. [Chrome浏览器全屏打开指定网页以及开机自启](https://blog.csdn.net/sunddy_x/article/details/122695593)
    - `Win + R`唤起"运行"，输入`shell:startup`打开程序启动项的目录.
    - 全屏打开指定网页有两种模式：`--kiosk`模式是chrome的演示模式，无法退出全屏并且禁用右键功能，`--start-fullscreen`模式是普通全屏模式，F11F12右键菜单均可使用。
3. [Kiosk Mode not available in Windows 11](https://superuser.com/questions/1749385/kiosk-mode-not-available-in-windows-11)
    - windows 10/11家庭版不支持`kiosk`特性.

[零刻EQ12小主机两天折腾之路](https://post.smzdm.com/p/a5o9q63x/)
[【播放/下载/远控】关于使用低功耗小主机打造家庭全能HTPC高清播放器的可行性](https://post.smzdm.com/p/a8x8rkx6/)
[HTPC/迷你主机遥控器操纵方案](https://zhuanlan.zhihu.com/p/606007339)


[给客厅小主机配个遥控器](https://post.smzdm.com/p/az6zk6v0/)
[【HTPC】几款蓝牙遥控器的体验点评、比较好用的遥控器推荐](https://zhuanlan.zhihu.com/p/604569395)
[【完美版】大屏幕+Windows系统+手柄遥控，完美版方案演示+分享~（派大星精选系列）](https://zhuanlan.zhihu.com/p/558765678)


## windows

关闭自动更新

底部菜单栏, 图标居中.

## chrome

### 插件与基本设置

adblock

插件Tab Wrangler: 自动关闭多余标签页. 设置为只保留一个标签页(当前页), 其他全部关闭. 超时时间为1秒, 即一跳转到新页面, 就关闭旧页面.

itab: 导航页/主页.

设置 -> 外观 -> 展示"主页"按钮
设置 -> 启动时 -> 主页

### 开机启动并自动全屏

#### 1. 开机启动

`Win + R`唤起"运行"，输入`shell:startup`打开程序启动项的目录, 一般为`C:\Users\general\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`

为 chrome 创建快捷方式, 并将快捷方式移动到该目录下.

#### 2. 自动全屏

右键 chrome 快捷方式, 点击"属性", 停在"快捷方式"标签页.

"目标(T)"的内容一般为`"C:\Program Files (x86)\Google\Chrome\Application\chrome.exe"`.

我们需要在后面追加全屏启动的参数, 这里选用`--kiosk`演示模式, 完整内容如下

`"C:\Program Files (x86)\Google\Chrome\Application\chrome.exe" --kiosk`

点击"确定".

> `kiosk`是windows的特性, "家庭版"是不支持的, 未激活状态的"专业版"也不支持, chrome启动后无法自动进入全屏状态.

### 自定义插件

可参考

1. 小舒同学 - 基于书签的新标签页
2. iLaunch: 管理书签、历史记录、标签、扩展 & 快捷命令启动工具

### 窗口置顶

chrome打开后自动置顶, 要看到其他窗口, 只能先关闭.

使用DeskPins-1.32工具, 添加自动置顶规则

- Description: `Chrome`
- Title: `*Google Chrome*`
- Class: `Chrome_WidgetWin_1`

每次chrome启动, 就会自动识别并置顶.

调整"图钉"标识跟随窗口的频率，太高会占用CPU，太低的话在移动窗口的时候图钉标识容易产生残影. 由于主要是全屏钉住, 因此可以调低一点, 设置为"200ms"就可以了.

### ~~右键映射"主页"按钮: AutoControl: Keyboard shortcut, Mouse gesture~~(飞鼠有主页功能)

配置选项

Action -> Trigger: "Right Button"
          |
          Conditions: Whole URL | doesn't | start with | chrome-extension://
          Conditions: Whole URL | doesn't | start with | chrome://
          |
          Action: Open New Tab

### 关闭自动更新

## Firefox

### 强制单标签页

[Open links in the same tab?](https://support.mozilla.org/en-US/questions/970999)

1. about:config 
2. browser.link.open_newwindow.restriction: 0
3. browser.link.open_newwindow: 1

## 无线摇控

remote mouse:

