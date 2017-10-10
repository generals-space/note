## 无法执行`ps1`脚本
```ps1
PS D:\> .\pshell.ps1
无法加载文件 D:\pshell.ps1，因为在此系统中禁止执行脚本。有关详细信息，请参阅 "get-help about_signing"。
所在位置 行:1 字符: 13
+ .\pshell.ps1 <<<<
    + CategoryInfo          : NotSpecified: (:) [], PSSecurityException  
    + FullyQualifiedErrorId : RuntimeException  
```

原因是操作系统默认禁止执行脚本，执行一次`set-executionpolicy remotesigned`后脚本顺利执行(需要在管理员启动的powershell中执行)

## 设置别名出错 

```
$ set-alias cp copy-item
set-alias : 无法从别名“cp”中删除 AllScope 选项。
所C:\Users\general\Documents\WindowsPowerShell\profile.ps1:2在位置 C:\Users\general\Documents\WindowsPowerShell\profile.ps1:2 字符: 1
+ set-alias cp copy-item
+ ~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : WriteError: (cp:String) [Set-Alias], SessionStateUnauthorizedAccessExcepti
   on
    + FullyQualifiedErrorId : AliasAllScopeOptionCannotBeRemoved,Microsoft.PowerShell.Commands.SetAliasC
   ommand
```