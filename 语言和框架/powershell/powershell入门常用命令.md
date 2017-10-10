# powershell入门常用命令

`write-host $psversiontable.psversion`: 查看当前powershell版本

```
$ write-host $psversiontable.psversion
5.1.14393.206
$ $psversiontable.psversion

Major  Minor  Build  Revision
-----  -----  -----  --------
5      1      14393  206
```

## 1. 帮助文档

`get-help 命令名`: 得到指定命令的帮助文档, 相当于man手册.

`get-help 命令名 -online`: 调用浏览器打开目标命令名在msdn上的帮助页面.

`get-help 命令名 -example|examples`: 查看目标命令的使用示例.

`update-help`: 更新所有命令的帮助文档(需要以管理员身份运行powershell)

------

`write-host 字符串/变量`: 控制台输出, 字符串或变量

`get-command 命令名`: 获取指定命令名的版本, 类型, 位置等信息, 还可以通过使用通配符查找包含指定字符串的命令. 示例如下

```ps1
$ get-command 'get-command'

CommandType     Name                                               Version    Source
-----------     ----                                               -------    ------
Cmdlet          Get-Command                                        3.0.0.0    Microsoft.PowerShell.Core

$ Get-Command '*alias*'

CommandType     Name                                               Version    Source
-----------     ----                                               -------    ------
Cmdlet          Export-Alias                                       3.1.0.0    Microsoft.PowerShell.Utility
Cmdlet          Get-Alias                                          3.1.0.0    Microsoft.PowerShell.Utility
Cmdlet          Import-Alias                                       3.1.0.0    Microsoft.PowerShell.Utility
Cmdlet          New-Alias                                          3.1.0.0    Microsoft.PowerShell.Utility
Cmdlet          Set-Alias                                          3.1.0.0    Microsoft.PowerShell.Utility
```

`alias`: 查看系统命令的所有简写版

`get-content 文件名/文件路径`: 查看目录文件内容, 相当于`cat`, 这是powershell默认的alias.

`set-alias 新字符串 旧字符串`: 相当于`alias`

`copy-item src dst`: 拷贝文件, 相当于`cp`

`remove-item 目标文件`: 删除目标文件, 也可以是目录.

`Set-Content test.txt -Value "i love light so much"`: 设置文本内容

`Add-Content test.txt -Value "but i love you more"`: 追加内容

`Clear-Content test.txt`: 清除内容

`Get-service`: 查看运行在机器上的所有服务

------

$profile: 有效的powershell配置文件路径.

`env`变量

`$env:temp`: 系统缓存目录

`$env:username`

`$env:userdomain`: 类似于hostname, 而实际上

`get-childitem env:`: 可以查看所有env内的变量, 这里其实把env当做了一个磁盘驱动器, `get-childitem`用以获得其子目录. 同理, 可以使用`ls env:*`来查看.


还有一个域, `variable`, 里面存放着显式定义过的变量的信息, 可以使用`ls variable:*`查看其中的变量(不加`*`号也可以).