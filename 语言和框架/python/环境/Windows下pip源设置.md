# Windows下pip源设置

参考文章

1. [windows及linux环境下永久修改pip镜像源的方法](https://www.jb51.net/article/98401.htm)

路径: `C:\Users\general\AppData\Roaming\pip`

> 实际上是在windows文件管理器中,输入`%APPDATA%`显示的路径

创建文件: `pip.ini`

内容

```ini
[global]
index-url = http://mirrors.aliyun.com/pypi/simple/
[install]
trusted-host=mirrors.aliyun.com
```

立即生效.