# Mac开启ssh服务

参考文章

1. [一条命令使Mac开启ssh服务](https://baijiahao.baidu.com/s?id=1600436641208190448&wfr=spider&for=pc)

Mac本身安装了ssh服务，但是默认情况下不会开机自启，因此当我们需要用到ssh相关的功能时，只需以下一条命令即可。

1. 启动sshd服务

```
sudo launchctl load -w /System/Library/LaunchDaemons/ssh.plist
```

2. 查看是否启动

```
sudo launchctl list | grep ssh
```

如果看到下面的输出表示成功启动了: 

```
-	0	com.openssh.sshd
```

3. 停止sshd服务

```
sudo launchctl unload -w /System/Library/LaunchDaemons/ssh.plist
```
