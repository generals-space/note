# powershell获取进程信息

参考文章

1. [PowerShell写守护进程](https://blog.csdn.net/weixin_36485376/article/details/83210591)

2. [PowerShell实现获取进程所有者](https://www.jb51.net/article/62740.htm)

3. [powershell – 通过文件名杀死进程](https://codeday.me/bug/20181102/351444.html)

4. [题 如何在PowerShell或C＃中获取进程的命令行信息](http://landcareweb.com/questions/8616/ru-he-zai-powershellhuo-c-zhong-huo-qu-jin-cheng-de-ming-ling-xing-xin-xi)

5. [Get-WmiObject - Use OR-Operator in -filter?](https://stackoverflow.com/questions/36861216/get-wmiobject-use-or-operator-in-filter)

## `Get-Process`

`Get-Process`可以获取系统中正在运行的进程列表, 类似于任务管理器中的数据.

`Get-Process node | Get-Member`: 可以查看一个进程对象的所有属性.

但是`Get-Process`能获取的信息有限, 比如`node index.js`, ta只能获取到`node.exe`和node本身所在的路径, 无法获得`index.js`的路径.

另外ta也不能得到进程的启动用户信息.

## `Get-WmiObject Win32_Process`

为了解决这个问题, 可以见参考文章3和4, 通过`Get-WmiObject`函数.

`Get-WmiObject Win32_Process -Filter "name = 'node.exe'"`得到更详细的数据, ta所返回的对象中有一个成员为`CommandLine`, 值为`node index.js`, 正好就是启动行启动的参数.

`Get-WmiObject`可所以查询的类型有如下几种:

- `Win32_Process`: 进程信息
- `Win32_LogicalDisk`: 本地逻辑卷
- `win32_service`: 服务信息

`filter`可用语法见参考文章5, 可使用类似sql过滤的查询方法. 如

```ps1
get-wmiobject win32_process -filter "name like '%python%'"
```

还可以使用or, and等操作符.

```ps1
get-wmiobject win32_process -filter "name = 'python.exe' and commandline like '%main.py%'"
```

> 使用`=`, `like`操作符时, 后面的字段串要用引号包裹.

查询当前系统中正在运行的python程序, 显示pid和启动命令

```ps1
get-wmiobject win32_process -filter "name like '%python%'" | select-object processid, commandline
```

**注意:**

虽然`get-wmiobject win32_process`获取的进程信息有很多字段, 但不是所有字段都可以作为`-Filter`的过滤选项. 比如`ProcessName`

```ps1
 $ get-wmiobject win32_process -filter "ProcessName = 'ssh.exe'"
get-wmiobject : 无效查询 “select * from win32_process where ProcessName = 'ssh.exe'”
```

