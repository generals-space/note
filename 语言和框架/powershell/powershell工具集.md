# powershell工具集

## 1. powercat

powershell可实现同linux下`netcat`相同功能的工具, 叫做[powercat](https://github.com/besimorhino/powercat)

使用方法同标准nc.

下载源码后命令行中执行`. .\powercat.ps1 选项参数`.

或者直接执行.

```ps1
IEX (New-Object System.Net.Webclient).DownloadString('https://raw.githubusercontent.com/besimorhino/powercat/master/powercat.ps1')
powercat -l -p 8000
```

在当前命令行会话中有效, 可以直接进行端口监听.