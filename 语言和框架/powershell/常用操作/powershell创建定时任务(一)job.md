# powershell创建定时任务(一)job

参考文章

1. [PowerShell实战指南 Chapter 13-17](https://wohin.me/powershell/2018/03/30/psPractice-chp13_17.html)
    - [Chapter 15 多任务后台作业]小节

2. [Start-Job](https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/start-job?view=powershell-6)

## 后台任务

创建新Job

```ps1
 gener@WORKGROUP  ~ $ Start-Job -ScriptBlock {ls}

Id     Name            PSJobTypeName   State         HasMoreData     Location             Command
--     ----            -------------   -----         -----------     --------             -------
1      Job1            BackgroundJob   Running       True            localhost            ls
```

- `-ScriptBlock {}`应该是必选项, 定义Job中要执行的命令;
- `-Name Job名称`定义Job名称;

查看Job列表(获取指定Job可以添加``作为过滤选项)

```ps1
 gener@WORKGROUP  ~ $ Get-Job

Id     Name            PSJobTypeName   State         HasMoreData     Location             Command
--     ----            -------------   -----         -----------     --------             -------
1      Job1            BackgroundJob   Completed     True            localhost            ls
```

`Get-Job`可用过滤选项:

1. `-Id 1`
2. `-Name Job名称`


其他管理命令

- `Remove-Job`
- `Stop-Job`
- `Wait-Job`

## 获取Job的输出结果

> To view the job's output, use the Receive-Job cmdlet. For example, `Receive-Job -Id 1`. --参考文章2

```ps1
 gener@WORKGROUP  ~ $ Receive-Job -id 1


    目录: C:\Users\gener\Documents


Mode                LastWriteTime         Length Name
----                -------------         ------ ----
d-----        2019/11/8     15:16                Tencent Files
d-----        2019/11/5     16:22                WindowsPowerShell
-a----        2019/9/23      9:54      140527620 test.pdf
```

## 向`-ScriptBlock`块中传入参数
