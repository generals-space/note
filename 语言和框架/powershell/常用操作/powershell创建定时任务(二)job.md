# powershell创建定时任务(二)job

参考文章

1. [PowerShell 计划工作（ScheduledJob）](https://www.pstips.net/about-scheduledjob.html)

2. [官方文档 New-JobTrigger](https://docs.microsoft.com/zh-cn/powershell/module/PSScheduledJob/New-JobTrigger?view=powershell-5.1&redirectedfrom=MSDN)

网上关于windows下的定时任务讲述的都是win下的计划任务, 与linux下的crontab有些出入. 大部分文章讲的都是一次性任务, 或是从某一时间开始执行的任务, 而不是我常用的循环任务.

创建定时任务要分为两个部分, 一个是任务本身, 一个是间隔时间的设置. 这其实和crontab很像, 因为crontab的规则也是两部分, 前半段为间隔时间, 后半段为执行的命令.

创建计划任务可以使用`Register-ScheduledJob`函数(对应的是`Unregister-ScheduledJob`), ta的用法与`Start-Job`相似.

```ps1
Register-ScheduledJob -Name myJob -ScriptBlock {
    ls ~ >> ~/job.log
}
```

然后就可以使用`Get-ScheduledJob`查看计划任务列表. 

```ps1
 gener@WORKGROUP  ~ $ Get-ScheduledJob

Id         Name            JobTriggers     Command                                  Enabled
--         ----            -----------     -------                                  -------
1          myJob           0               ...                                      True
```

> 同样, `Get-ScheduledJob`也有`-Id`和`-Name`过滤选项, 和`Get-Job`非常像.

然后通过`New-JobTrigger`创建间隔设置, 

```ps1
$myTrigger = New-JobTrigger -Once -RepeatIndefinitely -At (Get-Date) -RepetitionInterval (New-TimeSpan -Seconds 60)
```

注意:

1. `cronjob`形式的触发器需要同时指定`-Once`和`-RepeatIndefinitely`两个选项. 前者表示单次执行, 貌似是一个必选项, 如果要无限循环执行, 则需要加上后者;
1. `-At`貌似是一个必选项, 表示从某刻开始执行, `(Get-Date)`表示从现在开始. 注意需要用括号包裹, 因为`Get-Date`只是一个函数, `(Get-Date)`则表示其执行结果;
2. `-RepetitionInterval`时间间隔最短是1分钟, 如果小于1分钟(如5s), 就会报错 `New-JobTrigger : RepetitionInterval 参数值必须大于 1 分钟。`;

最后还需要将任务与触发器绑定.

```
Add-JobTrigger -Name myJob -Trigger $myTrigger
```

> 也可以在创建`ScheduledJob`对象时直接指定`-Trigger`参数, 如`Register-ScheduledJob -Name myJob -ScriptBlock {} -Trigger $trigger`.

绑定完成后任务即开始执行. 每到一个间隔的时间点, `ScheduledJob`就会创建一个`Job`对象, `Id`会递增, 但`Name`将会与`ScheduledJob`保持一致, 最终结果如下

```
 gener@WORKGROUP  ~ $ Get-Job

Id     Name            PSJobTypeName   State         HasMoreData     Location             Command
--     ----            -------------   -----         -----------     --------             -------
1      myJob           PSScheduledJob  Completed     True            localhost            ...
2      myJob           PSScheduledJob  Completed     True            localhost            ...
3      myJob           PSScheduledJob  Completed     True            localhost            ...
```

## 其他关于任务管理的操作.

### 查询JobTrigger

要明白, `JobTrigger`无法独立存在, 查询`JobTrigger`的时候实际需要的过滤参数是其绑定的`ScheduledJob`选项, 如下

```ps1
 gener@WORKGROUP  ~ $ Get-JobTrigger -Name myJob

Id         Frequency       Time                   DaysOfWeek              Enabled
--         ---------       ----                   ----------              -------
1          Once            2019/11/16 21:44:22                            True
```

这是查看单个trigger, 如果要查看所有存在的trigger, 可以尝试如下命令.

```ps1
Get-ScheduledJob | Get-JobTrigger
```

### 其他

- `Disable-ScheduledJob -Id 1`: 暂停定时任务
- `Enable-ScheduledJob -Id 1`: 重启定时任务
- `Unregister-ScheduledJob -Id 1`: 删除定时任务, 同时也会删除绑定的`JobTrigger`和衍生的`Job`对象.
- `Remove-JobTrigger -Name myJob -Id 1`: 移除指定定时任务上绑定的触发器, 不指定`-Id`参数的话将移除目标任务上的所有触发器.
- `Set-JobTrigger`: 修改目标trigger的内容, 对于某一定时任务中的trigger, 可以先获取到trigger, 再用管道的方式传入
    - `Get-JobTrigger -Name myJob -TriggerID 1 | Set-JobTrigger -Weekly -WeeksInterval 4 -DaysOfWeek Monday -At "12:00 AM"`

