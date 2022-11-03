# pip镜像源设置

1. [mac 修改pip镜像为国内镜像](https://blog.csdn.net/liushimiao0104/article/details/96475364)
2. [MAC的Python PIP源设置成阿里云](https://blog.csdn.net/u012865381/article/details/80390258)
3. [windows及linux环境下永久修改pip镜像源的方法](https://www.jb51.net/article/98401.htm)
4. [国内可用镜像源列表](https://www.pypi-mirrors.org/)

`easy_install`和`pip`安装第三方库都是从Python的官方源`pypi.python.org/pypi` 下载到本地,然后解包安装.

以阿里云镜像源为例`http://mirrors.aliyun.com/pypi/simple/`

使用方式

```
easy_install -i http://mirrors.aliyun.com/pypi/simple/ saltTesting
pip install -i http://mirrors.aliyun.com/pypi/simple/ saltTesting
```

建议**写入配置文件**, 默认使用镜像源

在`~/.pip/pip.conf` 中添加

```ini
[global]
trusted-host = mirrors.aliyun.com
index-url = http://mirrors.aliyun.com/pypi/simple/
```

> PS: pip的官方文档:`http://pip.readthedocs.org/`, 在`configuration`一节可详细查询其使用方式

------

2016-08-08更新:

pip官方文档介绍`$HOME/.pip/pip.conf`只是针对单用户的配置, 全局配置在`/etc/pip.conf`, 这个也是需要自行创建的.

## macos

mac下的pip配置文件路径也在`~/.pip/pip.conf`.

## windows

路径: `C:\Users\general\AppData\Roaming\pip`

> 实际上是在windows文件管理器中,输入`%APPDATA%`显示的路径

创建文件: `pip.ini`

```ini
[global]
index-url = http://mirrors.aliyun.com/pypi/simple/
[install]
trusted-host=mirrors.aliyun.com
```

立即生效.
