# pip镜像源设置

<!-- <!tags!>: <!pip!> <!镜像源!> -->

`easy_install`和`pip`安装第三方库都是从Python的官方源`pypi.python.org/pypi` 下载到本地,然后解包安装.

[国内可用镜像源列表](https://www.pypi-mirrors.org/)

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
