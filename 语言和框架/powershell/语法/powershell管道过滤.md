# powershell管道

参考文章

1. [Powershell 语句块](http://www.pstips.net/powershell-using-scriptblocks.html)

## 1. 过滤

首先要明白, powershell的命令返回值大部分是对象类型, 而不是像bash shell那种纯粹的字符串. 这些对象拥有属性和值, 所以过滤它们的返回值时, 需要指定目标属性, 或者加上期望值. 比较常用的是如下命令.

1. `Select-Object`

2. `Where-Object`

3. `ForEach-Object`

### 索引过滤

`select-object`可以选择出单个或多个对象的属性, 还能像`head`, `tail`一样取出结果的开头/结尾的n行. 

如下, `get-service`返回系统所有服务(不管有没有正在运行), 使用`select-object`取出开头10个和结尾5个结果. 看清楚了, 它其实比较`head`, `tail`还要强, 因为它可以用单条命令完成.

```ps
$ get-service | select-object -first 10 -last 5

Status   Name               DisplayName
------   ----               -----------
Stopped  AJRouter           AllJoyn Router Service
Stopped  ALG                Application Layer Gateway Service
Stopped  AppIDSvc           Application Identity
Running  Appinfo            Application Information
Stopped  AppMgmt            Application Management
Stopped  AppReadiness       App Readiness
Stopped  AppVClient         Microsoft App-V Client
Stopped  AppXSvc            AppX Deployment Service (AppXSVC)
Running  AudioEndpointBu... Windows Audio Endpoint Builder
Running  Audiosrv           Windows Audio
Stopped  WwanSvc            WWAN AutoConfig
Stopped  XblAuthManager     Xbox Live 身份验证管理器
Stopped  XblGameSave        Xbox Live 游戏保存
Stopped  XboxNetApiSvc      Xbox Live 网络服务
Stopped  ZhuDongFangYu      主动防御
```

`select-object`还可以使用`-index`选项指定显示第n行(多个索引用逗号隔开).

```ps
$ get-service | select -index 5,6,7

Status   Name               DisplayName
------   ----               -----------
Stopped  AppReadiness       App Readiness
Stopped  AppVClient         Microsoft App-V Client
Stopped  AppXSvc            AppX Deployment Service (AppXSVC)
```

### 列过滤

`select-object`可以选择显示的属性, 在一组结果对象中, 就相当于选择列的显示. 语法就是, 直接指定属性名. 例如上例中可以选择`status`, `name`和`displayname`. 选择多个属性时使用逗号`,`隔开.

```ps
$ get-service | select-object status,name  -first 5

 Status Name
 ------ ----
Stopped AJRouter
Stopped ALG
Stopped AppIDSvc
Running Appinfo
Stopped AppMgmt
```

**注意**

当然, `get-service`的输出应该不只只有`status`, `name`, `displayname`3个属性, 要查看某个对象的全部属性, 可以使用`get-member`命令.

```ps
$ get-service |  get-member

   TypeName:System.ServiceProcess.ServiceController

Name                      MemberType    Definition
----                      ----------    ----------
Name                      AliasProperty Name = ServiceName
RequiredServices          AliasProperty RequiredServices = ServicesDependedOn
Disposed                  Event         System.EventHandler Disposed(System.Object, System.EventArgs)
Close                     Method        void Close()
Continue                  Method        void Continue()
...
```

## 行选择

行选择类似于`grep`操作, 需要从管道中传入的值中过滤信息. 需要用到`where-object`.

`where-object`可以有多种比较方法, 可以通过指定关键字, 正则表达式, 或是数值甚至数值范围来过滤.

如下, 常用的还是通过字符串过滤. 从`get-service`输出的系统所有服务中过滤, 正在运行的服务.

```ps
$ get-service | where-object {$_.status -ccontains 'running'}

Status   Name               DisplayName
------   ----               -----------
Running  Appinfo            Application Information
Running  AudioEndpointBu... Windows Audio Endpoint Builder
Running  Audiosrv           Windows Audio
Running  BFE                Base Filtering Engine
Running  BITS               Background Intelligent Transfer Ser...
Running  BrokerInfrastru... Background Tasks Infrastructure Ser...
Running  CDPUserSvc_2146b9  CDPUserSvc_2146b9
Running  COMSysApp          COM+ System Application
...省略
```

> 过滤条件需要使用`{}`包裹起来.

> `$_`默认表示每个传入`where-object`的行对象.

`where-object`的过滤方式很多, 常用的有

- `-contains 字符串`: 包含目标子字符串;

- `-notcontains 字符串`: 不包含目标子字符串;

- `-match 模式`: 符合目标正则模式;

- `-eq|ge|gt|le|lt -value 数值`: 数值比较;

- `like *字符串*`: 类似于sql的like查询;

## 2. 遍历

遍历方法用到`ForEach-Object`命令, 它有几个别名: `%`, `foreach`.

比如使用`1..10`创建一个数字序列, 可以遍历其中每一个成员

```ps1
PS C:\Users\general> 1..10 | % {echo $_}
1
2
3
4
5
6
7
8
9
10
```

其中`%`就是foreach-object函数, 而`$_`则表示当前项.

