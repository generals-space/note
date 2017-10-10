# powershell远程连接

<!tags!>: <!远程连接!>

参考文章

1. [Powershell远程管理服务器&客户端（非mstsc远程连接）](http://www.2cto.com/net/201607/529452.html)

这里的远程连接不是指远程桌面的那种连接, 而是相当于linux下`ssh`远程控制. windows中有一个`WinRM`(Web Services for Management). 它使用了一个新的通信协议, 有可能是`http`或`https`.

## 1. 开启远程连接服务

管理员权限运行Powershell, 输入`Enable-PsRemoting`开启Powershell远程管理, **远程端和被远程端都需要启用**. 另外说明一下, `WinRM`也就是Powershell远程管理时使用的端口http: 5985; https: 5986.

当然也可以修改默认的端口号, 但是这么做的话每次进行远程操作时需要制定端口号进行连接, 具体方法见help文档.

```ps
$ get-service | ? name -Contains 'winrm'

Status   Name               DisplayName
------   ----               -----------
Stopped  WinRM              Windows Remote Management (WS-Manag...

$ Enable-PSRemoting
WinRM 已更新为接收请求。
成功更改 WinRM 服务类型。
已启动 WinRM 服务。

WinRM 已经进行了更新, 以用于远程管理。
WinRM 防火墙异常已启用。
已配置 LocalAccountTokenFilterPolicy 以远程向本地用户授予管理权限。

$ get-service | ? name -Contains 'winrm'

Status   Name               DisplayName
------   ----               -----------
Running  WinRM              Windows Remote Management (WS-Manag...
```

## 2. 远程连接

这种方式其实一般用于域中计算机的远程连接, 听说未在同一个域中远程连接比较鸡肋.

连接方法是使用`enter-pssession`命令...当然, 失败了...

```ps
PS C:\WINDOWS\system32> Enter-PSSession -ComputerName 172.32.100.142
Enter-PSSession : 连接到远程服务器 172.32.100.142 失败，并显示以下错误消息: WinRM 客户端无法处理该请求。如果身份验证方
案与 Kerberos 不同，或者客户端计算机未加入到域中， 则必须使用 HTTPS 传输或者必须将目标计算机添加到 TrustedHosts 配置设
置。 使用 winrm.cmd 配置 TrustedHosts。请注意，TrustedHosts 列表中的计算机可能未经过身份验证。 通过运行以下命令可获得有
关此内容的更多信息: winrm help config。 有关详细信息，请参阅 about_Remote_Troubleshooting 帮助主题。
所在位置 行:1 字符: 1
+ Enter-PSSession -ComputerName 172.32.100.142
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidArgument: (172.32.100.142:String) [Enter-PSSession]，PSRemotingTransportException
    + FullyQualifiedErrorId : CreateRemoteRunspaceFailed

```

> `Kerberos`貌似是域内计算机连接的认证方式.

要连接非域内pc, 需要首先将对方加入到本地的可信任主机列表中. 如下

```ps
PS C:\WINDOWS\system32> set-item WSMan:\localhost\Client\TrustedHosts -value 172.32.100.142

WinRM 安全配置。
此命令修改 WinRM 客户端的 TrustedHosts 列表。TrustedHosts
列表中的计算机可能不会经过身份验证。该客户端可能会向这些计算机发送凭据信息。是否确实要修改此列表?
[Y] 是(Y)  [N] 否(N)  [S] 暂停(S)  [?] 帮助 (默认值为“Y”): y
```

然后需要与目标pc建立连接(需要认证)

```ps
PS C:\WINDOWS\system32> $session = Get-Credential

位于命令管道位置 1 的 cmdlet Get-Credential
请为以下参数提供值:
Credential
```

这里会弹出一个窗口, 提示输入目标pc的用户名和密码, 认证成功后, 这个会话变量`$session`会被保留, 作为`enter-pssession`的参数.

```
PS C:\WINDOWS\system32> Enter-PSSession -ComputerName 172.32.100.142 -Credential $session
[172.32.100.142]: PS C:\Users\general\Documents>
```

连接成功.