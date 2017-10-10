# Powershell修改系统环境变量-[Environment]类使用

参考文章

1. [System.Environment类应用小技巧](http://blog.csdn.net/smeller/article/details/7059944)

2. [微软官方文档 - Creating and Modifying Environment Variables](https://technet.microsoft.com/en-us/library/ff730964.aspx)

3. [Using the System.Environment Class in PowerShell](http://www.brangle.com/wordpress/2009/08/using-the-system-environment-class-in-powershell/)

4. [微软官方文档 - Environment 方法](https://msdn.microsoft.com/zh-cn/library/system.environment_methods(v=vs.110).aspx)

这里讲的修改系统环境变量是指的操作系统中的`Path`变量, 全局生效的那种. 而不是像Shell脚本运行时在`/etc/profile`中定义的变量那样, 更狭义一点.

`[Environment]`是`[System.Environment]`的缩写, 是一个类...没错, 在`C#`语法中的确是一个类, 而且无法被继承???(谁没事想去继承它). 

```ps1
PS C:\Users\general> [System.Environment]

IsPublic IsSerial Name                                     BaseType
-------- -------- ----                                     --------
True     False    Environment                              System.Object
```

## 1. 基本语法

```ps1
PS C:\Users\general> [Environment]::UserName
general
PS C:\Users\general> [System.Environment]::GetEnvironmentVariable("Path", "User")
;D:\Microsoft VS Code\bin;C:\Users\general\AppData\Local\Microsoft\WindowsApps
```

猜测`Username`是类的静态变量, `GetEnvironmentVariable`是静态方法, 都可以直接调用.

`GetEnvironmentVariable`有一个重载方法接受两个参数, 第一个是变量名string, 第二个是其作用范围target.

`target`变量是`EnvironmentVariableTarget`枚举类型的其中一个, 有3个可能的取值: `Machine`, `Process`, `User`.