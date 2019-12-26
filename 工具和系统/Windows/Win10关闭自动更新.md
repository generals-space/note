# Win10关闭自动更新

参考文章

1. [为什么很多人要禁止 Windows 10 自动更新？](https://www.zhihu.com/question/271414438?sort=created)
    - 有意思的话题
1. [[Windows] win10 跟自动更新说再见，只需要几秒钟](https://www.52pojie.cn/thread-1067072-1-1.html)
    - `windows update blocker`工具, 谷歌可搜索
3. [Windows update禁用后自动重启的解决方法](https://jingyan.baidu.com/article/fcb5aff766a4f0edaa4a710c.html)
    - 主要还是关闭`Windows Update Medic Service`服务的方法, 图文并茂
4. [Windows Update Medic Service 拒绝访问](https://www.cnblogs.com/feipeng8848/p/10096151.html)
    - 关闭`Windows Update Medic Service`的方法-修改注册表, 这个方法在参考文章3中也能找到.
    - 不过只需要修改`Start`字段即可, 不需要修改`FailureActions`, (而且其中的选择排列可能不一致, 容易出错).
5. [windows update medic service不能设置为自动，显示拒绝访问导致系统更新不了](https://answers.microsoft.com/zh-hans/windows/forum/windows_10-performance/windows-update-medic/c0cc9dcc-b5fd-47c9-9c9b-b28cb6d2f32c)
    - MS官方的问答网站, 除了参考文章3和4通过设置注册表, 启用`Administrator`用户修改, 不过没用...
6. [如何彻底禁止Windows 10自动更新？ - 不名的回答 - 知乎](https://www.zhihu.com/question/287260272/answer/483721867)
    - `usosvc.dll`, `wuaueng.dll`, `WaaSMedicSvc.dll`
    - 编辑权限时可能复选框是灰色的, 无法点击, 可以使用`takeown /f usosvc.dll`先获取文件权限.

win10 2019经常半夜自动更新, 更新完自动重启, 我运行的虚拟机集群全被强杀了. 

参考文章1中有提到win10的自动更新非常不人性化, 因为虽然可以设置更新时间, 设置推迟时间, 但也在很多场合机器是24小时不能关机的. 而且win10判断系统空闲的算法也有些脱俗, 貌似只检测用户键鼠活动, 也不管用户是否有后台任务在进行就强制重启了...另外, 单纯禁用`windows update`服务已经无效了, 自动更新设计得极为阴险.

20191226更新

失败了, 修改注册表也被windows绕过了...

后来找到参考文章6, 给出的方式是修改3个系统文件的权限.

按照其中的方法修改之后, `windows update`服务无法打开(会弹出错误). 但是打开右下角"所有设置"时会变得很慢, 而且设置中打开"更新和安全"会变得超级慢...

然后把`usosvc.dll`和`wuaueng.dll`的权限恢复, 只修改`WaasMedicSvc.dll`, 这样在"更新和安全"中的系统更新也会失败, 且`Windows Update Medic Service`在"服务"中消失了.

先这样看看吧...

------

除了需要关闭`windows update`服务, 还需要关闭`Windows Update Medic Service`服务, 因为后者会尝试重新启动前者. 

![](https://gitee.com/generals-space/gitimg/raw/master/29AF5ECF67C49F324E2EF693C68D596A.jpg)

不过禁用后者, 同时设置不自动启动前者时, 点击"应用"会显示访问拒绝.

![](https://gitee.com/generals-space/gitimg/raw/master/C0AF9543A2739487FE5C85BAD2323B03.png)

按照参考文章3和4的描述, 需要修改注册表字段, 将`HEKY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\WaaSMedicSvc` 中`Start`的值改为`4`, 然后再点击应用, 虽然还是会显示访问拒绝, 但是服务已经被禁用了. 而且在"Windows更新"页面会显示更新遇到错误.

![](https://gitee.com/generals-space/gitimg/raw/master/82611AACEBB50C6651792834E2D9E9C7.jpg)
