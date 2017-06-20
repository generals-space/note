# powershell绕过脚本执行权限

<!tags!>: <!powershell!> <!执行权限!> <!绕过!>

## 1. 相关命令

`get-executionpolicy`: 得到当前会话的脚本执行策略

`set-executionpolicy`: 设置当前用户的执行策略.

后者可选的策略有: `Unrestricted` | `RemoteSigned` | `AllSigned` | `Restricted` | `Default` | `Bypass` | `Undefined`

windows默认无法执行ps脚本. 比如, 用户home目录下有如下脚本.

```ps
## test.ps1
write-host 'hello world'
```

在cmd或powershell控制台执行它.

```
$ .\test.ps1
.\test.ps1 : 无法加载文件 C:\Users\general\Downloads\test.ps1，因为在此系统上禁止运行脚本。有关详细信息，请参阅 http://
go.microsoft.com/fwlink/?LinkID=135170 中的 about_Execution_Policies。
所在位置 行:1 字符: 1
+ .\test.ps1
+ ~~~~~~~~~~
    + CategoryInfo          : SecurityError: (:) []，PSSecurityException
    + FullyQualifiedErrorId : UnauthorizedAccess
```

我们需要显式在powershell控制台修改这种执行策略, 而且需要在管理员运行的powershell控制台.

```
$ Set-ExecutionPolicy remotesigned

执行策略更改
执行策略可帮助你防止执行不信任的脚本。更改执行策略可能会产生安全风险，如 http://go.microsoft.com/fwlink/?LinkID=135170
中的 about_Execution_Policies 帮助主题所述。是否要更改执行策略?
[Y] 是(Y)  [A] 全是(A)  [N] 否(N)  [L] 全否(L)  [S] 暂停(S)  [?] 帮助 (默认值为“N”): a
```

再次执行它, 成功.

```
$ .\test.ps1
hello world
```

## 2. 临时绕过

在powersploit应用中, 不太可能拥有管理员权限, 而且设置全局的脚本执行策略代价太大, 容易被发现. 所以有一点临时执行脚本的技巧.

```ps
$ powershell -executionpolicy bypass -file .\test.ps1
hello world
```

> 这里的powershell, 有种类似于子shell的作用, 毕竟`-executionpolicy`选项不能通过`./`命令完成.