# Win10做ssh隧道穿透

首先在win上装ssh服务, 专业版可以在控制面板的"可选功能"中添加`openssh`服务, 如果是教育版/家庭版, 则需要自行下载安装`openssh`服务.

然后在`~/.ssh`目录下放置密钥对, 用于无密码登录中转机.

然后添加`config`服务器别名列表.

```
Host forwarder
    ## HostName可以是域名也可以是IP
    HostName forwarder-IP
    Port 22
    User jiale.huang
    ServerAliveInterval 60

Host forward-ssh
    HostName forwarder-IP
    Port 22
    User jiale.huang
    ServerAliveInterval 60
    ## 不开启tty, 与ssh的`-T`选项作用相同
    RequestTTY no
    ## RemoteForward 0.0.0.0:10001 127.0.0.1:22
    RemoteForward 172.17.0.6:10001 127.0.0.1:22

Host forward-rdp
    HostName forwarder-IP
    Port 22
    User jiale.huang
    ServerAliveInterval 60
    ## 不开启tty, 与ssh的`-T`选项作用相同
    RequestTTY no
    ## RemoteForward 0.0.0.0:10002 127.0.0.1:3389
    RemoteForward 172.17.0.6:10002 127.0.0.1:3389

Host forward-vnc
    HostName forwarder-IP
    Port 22
    User jiale.huang
    ServerAliveInterval 60
    ## 不开启tty, 与ssh的`-T`选项作用相同
    RequestTTY no
    ## RemoteForward 0.0.0.0:10003 127.0.0.1:5900
    RemoteForward 172.17.0.6:10003 127.0.0.1:5900

```

将[ps-libs](https://gitee.com/generals-space/ps-libs)项目clone到`~/Documents/WindowsPowerShell`目录.

```
git clone https://gitee.com/generals-space/ps-libs.git ./WindowsPowerShell
```

`ps-libs`项目中提供了通过`ssh`命令连接中转机做穿透的函数.

然后注册定时任务, 这段命令在`ps-libs`中ssh相关的脚本中的注释部分.

```ps1
Register-ScheduledJob -Name cronJob -FilePath ~\Documents\WindowsPowerShell\ssh-reconnect.ps1
$cronTrigger = New-JobTrigger -Once -RepeatIndefinitely -At (Get-Date) -RepetitionInterval (New-TimeSpan -Seconds 60)
Add-JobTrigger -Name cronJob -Trigger $cronTrigger
```

> 注意: `ps-libs`中`ssh-connect.ps1`脚本文件中的`get_ssh_proc forward-XXX`指令要与`~/.ssh/config`中的ssh别名对应.

再然后是`vim`设置, 上面的形式连接windows的话进入的是windows的powershell命令行. 如果需要在命令行完成编辑任务, 可以使用`git for windows`中的bash, 把`bash.exe`路径添加到环境变量就可以了. 

