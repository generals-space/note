# Mac命令行操作parallel虚拟机[pd]

查看列表

```console
$ prlctl list
UUID                                    STATUS       IP_ADDR         NAME
{e1746e62-d8ae-4352-a715-f4c3d23d2778}  running      -               Windows 10
```

停止指定实例

```console
$ prlctl stop e1746e62-d8ae-4352-a715-f4c3d23d2778
Stopping the VM...
The VM has been successfully stopped.

$ prlctl list
UUID                                    STATUS       IP_ADDR         NAME
```
