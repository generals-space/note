# VSCode remote连接穿透后的win10主机

参考文章

1. [Windows Server SSH Remoting Fails if Server has Win32_OpenSSH 7.9 or newer installed](https://github.com/microsoft/vscode-remote-release/issues/2475)
2. [Failed to set remote SSH tunnel](https://github.com/microsoft/vscode-remote-release/issues/75)

VSCode: 1.46.1
Remote SSH: v0.51.0
Win10: 教育版 1903(build number: 18362)

win10电脑使用开源 openssh 服务做内网穿透, 使用另一台 mac 远程连接暴露出来的端口时失败, 具体内容如下.

```
[17:55:01.210] stderr> OpenSSH_for_Windows_8.1p1, LibreSSL 2.9.2
[17:55:02.159] stderr> debug1: Server host key: ecdsa-sha2-nistp256 SHA256:ubsZl1Q9SjOlzifmMqGd/w8m95rTpUtsSeRL2FJgMh4
[17:55:03.703] stderr> Authenticated to 0.tcp.ngrok.io ([3.19.114.185]:19084).
[17:55:04.411] stderr> shell request failed on channel 2
[17:55:04.423] > local-server> ssh child died, shutting down
[17:55:04.433] Local server exit: 0
[17:55:04.434] Received install output: OpenSSH_for_Windows_8.1p1, LibreSSL 2.9.2
debug1: Server host key: ecdsa-sha2-nistp256 SHA256:ubsZl1Q9SjOlzifmMqGd/w8m95rTpUtsSeRL2FJgMh4
Authenticated to 0.tcp.ngrok.io ([3.19.114.185]:19084).
shell request failed on channel 2

[17:55:04.434] Stopped parsing output early. Remaining text: OpenSSH_for_Windows_8.1p1, LibreSSL 2.9.2debug1: Server host key: ecdsa-sha2-nistp256 SHA256:ubsZl1Q9SjOlzifmMqGd/w8m95rTpUtsSeRL2FJgMh4Authenticated to 0.tcp.ngrok.io ([3.19.114.185]:19084).shell request failed on channel 2
[17:55:04.435] Failed to parse remote port from server output
[17:55:04.436] Resolver error: 
[17:55:04.438] ------
```

按照参考文章1的说法, 是因为win10上的 openssh 版本太新了, 高于7.9的 openssh 都会出现这个问题(我的是8.1). 

后来将其降为 7.7, 使用 mac 重新连接没有问题, 但是在测试时使用 win10 通过穿透暴露出来的公网端口连接ta本身仍然报错, 显示`vscode-server start failed`, 相关日志中有如下内容.

```
*
* Reminder: You may only use this software with Visual Studio family products,
* as described in the license (https://go.microsoft.com/fwlink/?linkid=2077057)
*
```

具体报错内容见参考文章2, 同时ta也给出了解决方法, 设置`remote.SSH.showLoginTerminal`为`true`即可, 也无需重启.

不过使用 mac 连接却完全没问题, 不需要设置这个参数...
