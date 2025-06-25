# Windows进程管理命令-tasklist与taskkill

参考文章

1. [Windows 进程 Tasklist查看 与 Taskkill结束](https://blog.csdn.net/wangmx1993328/article/details/80923829)

`tasklist`与`taskkill`是bat命令, 在powershell里有类似的`Get-Process`与`Stop-Process`.

## 1. tasklist

tasklist默认输出本机所有进程, 想要实现类似`grep`的操作, 需要配合`findstr`命令(bat命令). 如下

```
PS C:\Users\general> tasklist | findstr 'sublime'
sublime_text.exe              1940 Console                    1     30,656 K
```

当然也可以根据pid查询, 就不做示例了.

常用参数

`-v` 可打印出详细信息, 包括启动命令, 可执行文件的路径

```
PS C:\Users\general> tasklist /v | findstr 'sublime'
sublime_text.exe    1940 Console    1     30,660 K Running         general-PC\general 0:00:01 C:\Program Files\MxiPlayer-3.0.8\config.json ? - Sublime Text (UNREGISTE
```

`/svc` 可显示各进程启动的服务.

`/fo` format, 输出格式, 可选值为 table, list, csv, 默认为table

`/fi` tasklist内置的过滤选项, 有`eq`, `gt`, `lt`等操作符, 可过滤包括pid, 进程名, cpu时间等字段, 不常用, 具体的可查看`/?`帮助手册.

## 2. taskkill

典型用法就是`taskkill pid`, 结束指定进程.

`taskkill`有两种方式指定目标, `/pid 进程号`可以杀死指定pid的进程, 也可以使用`/im 进程名`杀死指定名称的进程.

对于`/im`选项(image, 意味进程映像), 进程名可以以`*`结尾(或者直接就是一个`*`号...咳)批量停止同名进程. 但不能像`*nginx.exe*`这种放到开头.

```ps1
PS C:\Users\general> taskkill /im sublime_text.exe
成功: 给进程 "sublime_text.exe" 发送了终止信号，进程的 PID 为 1940。
```

```ps1
PS C:\Users\general> taskkill /im *sublime*
错误: 没有找到进程 "*sublime*"。
PS C:\Users\general> taskkill /im sublime*
成功: 给进程 "sublime_text.exe" 发送了终止信号，进程的 PID 为 1956。
```

`/f` 强杀. (对于Office、WPS此类软件在打开文件的情况下，如果采用强制杀死进程的方式，则下一次再打开文件时，就很可能会提示文件错误，这就是因为强杀进程导致的，所以此时则不再建议加上`/f`参数)

`/t` 结束进程树, 就是可以停止所有子进程, 但你需要找到哪个才是父进程...