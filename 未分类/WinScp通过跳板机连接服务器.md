# WinScp通过跳板机连接服务器

参考文章

1. [winScp如何通过隧道代理进行远程连接](https://www.cnblogs.com/wangkongming/p/4124945.html)

2. [如何远程使用跳板机连接设备传输文件](http://www.wkgoto.com/i/216984.html)

3. [SSHTunnelTroubleshooting](http://wiki.metawerx.net/wiki/SSHTunnelTroubleshooting)

我们已经知道XShell, secureCRT可以通过配置**登录脚本**实现跳板机登录, 但是有时希望ftp工具也能实现类似的功能(有些开发人员不懂linux或vim, 对于线上修改配置, 上传文件等, 使用ftp工具更方便).

本来更喜欢用filezilla, 但是网上说它没法实现, 倒是WinScp可以. 下面以WinScp为例.

跳板机地址: 192.168.101.65

目标服务器: 192.168.163.52

用户可以用密钥以`log.game`用户登录跳板机, 然后从跳板机以`log`用户登录目标机器.

新建会话

![](https://gitee.com/generals-space/gitimg/raw/master/fd1541258db2ca7b1526c25939ea4af7.png)

右侧框中在选择密钥时默认寻找后缀为`ppk`的文件, 这是putty使用的密钥格式. 我自己的密钥是XShell所用的格式, 分为公钥和私钥. 

初次选择时, 可以显示所有文件, 然后选中私钥, WinScp会弹出如下提示.

![](https://gitee.com/generals-space/gitimg/raw/master/16132f9a1ce88e84c22f6affb18c3fb5.png)

按照它的提示就可以把原来的私钥转换成`ppk`的格式了.

然后继续

![](https://gitee.com/generals-space/gitimg/raw/master/2ceaec4491c1e9858ece6c850affc304.png)

上图中右侧框的密钥文件指的是从跳板机上用于`log`用户登录目标机器的.

确定, 保存, 连接.

## FAQ

### 1. 连接失败

按照上面的步骤完成操作, 点击连接却报如下错误.

![](https://gitee.com/generals-space/gitimg/raw/master/98b0a112c33025b42f2be1fa04a8cd00.png)

跳板机上的日志如下.

```
Apr 12 16:45:32 192_168_101_65 sshd[20092]: Accepted publickey for log.game from 10.96.0.46 port 57155 ssh2
Apr 12 16:45:32 192_168_101_65 sshd[20092]: pam_unix(sshd:session): session opened for user log.game by (uid=0)
Apr 12 16:45:32 192_168_101_65 sshd[20095]: Received request to connect to host 192.168.163.52 port 22, but the request was denied.
Apr 12 16:45:32 192_168_101_65 sshd[20092]: pam_unix(sshd:session): session closed for user log.game
```

可以看到`log.game`登录跳板机已经成功, 但是登录目标机器就不行了.

按照参考文章3中对**Forwarded connection refused by server: Administratively prohibited [open failed], or channel N: open failed: administratively prohibited: open failed**的讨论, 查出是因为跳板机的ssh不支持端口转发, `/etc/ssh/sshd_config`中`AllowTcpForwarding`字段的值为`no`.

将跳板机上这个值改成`yes`, 重启跳板机上的sshd, 再次连接就成功了.