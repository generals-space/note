# powershell获取对象的所有属性

参考文章

1. [如何列出Powershell对象的所有属性？](https://codeday.me/bug/20170711/41352.html)

为什么有这样的需求? 

举例来说, 下面的语句只显示了如下几条结果, 我们只能看到6条属性.

```console
$ Get-WmiObject -Class Win32_OperatingSystem

SystemDirectory : C:\WINDOWS\system32
Organization    : Razer
BuildNumber     : 18362
RegisteredUser  : generals.space@gmail.com
SerialNumber    : 00331-20300-00000-AA015
Version         : 10.0.18362
```

但是某些文章有如下用法.

```ps1
$ Get-WmiObject -Class Win32_OperatingSystem | Select-Object Caption

Caption
-------
Microsoft Windows 10 专业版
```

...`Caption`属性哪里来的?

其实可以说大部分`WMI`相关函数都只返回一小部分的属性列表, 为了获取我们想要的信息, 需要额外做一些操作.

## 常用示例

下面列出几个常用的方法.

1. `Get-Member`: 可以得到目标对象所有的属性及ta们的类型

```ps1
$ Get-WmiObject -Class Win32_OperatingSystem | Get-Member

   TypeName:System.Management.ManagementObject#root\cimv2\Win32_OperatingSystem

Name                                      MemberType    Definition
----                                      ----------    ----------
PSComputerName                            AliasProperty PSComputerName = __SERVER
Reboot                                    Method        System.Management.ManagementBaseObject Reboot()
...省略
```

2. `Format-List`: 可以得到目标对象的所有属性及属性值列表.

```ps1
$ Get-WmiObject -Class Win32_OperatingSystem | Format-List *

PSComputerName                            : LAPTOP-S9NMLS7S
Status                                    : OK
Name                                      : Microsoft Windows 10 专业版|C:\WINDOWS|\Device\Harddisk0\Partition4
...省略
```

3. `Select-Object key`: 可以得到目标对象指定属性的属性值

```ps1
 $ Get-WmiObject -Class Win32_OperatingSystem | Select-Object Caption

Caption
-------
Microsoft Windows 10 专业版
```
