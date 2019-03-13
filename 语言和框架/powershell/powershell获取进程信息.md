# powershell获取进程信息

参考文章

1. [PowerShell写守护进程](https://blog.csdn.net/weixin_36485376/article/details/83210591)

2. [PowerShell实现获取进程所有者](https://www.jb51.net/article/62740.htm)

3. [powershell – 通过文件名杀死进程](https://codeday.me/bug/20181102/351444.html)

4. [题 如何在PowerShell或C＃中获取进程的命令行信息](http://landcareweb.com/questions/8616/ru-he-zai-powershellhuo-c-zhong-huo-qu-jin-cheng-de-ming-ling-xing-xin-xi)

`Get-Process`可以获取系统中正在运行的进程列表, 类似于任务管理器中的数据.

`Get-Process node | Get-Member`: 可以查看一个进程对象的所有属性.

但是`Get-Process`能获取的信息有限, 比如`node index.js`, ta只能获取到`node.exe`和node本身所在的路径, 无法获得`index.js`的路径.

另外ta也不能得到进程的启动用户信息.

为了解决这个问题, 可以见参考文章3和4, 通过`Get-WmiObject`函数.

`Get-WmiObject Win32_Process -Filter "name = 'node.exe'"`得到更详细的数据, ta所返回的对象中有一个成员为`CommandLine`, 值为`node index.js`, 正好就是启动行启动的参数.