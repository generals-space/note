# Pypi私有镜像源搭建-pypiserver

<!date!>: 2017-06-20

参考文章

1. [如何搭建自己的pypi私有源服务器](http://python.jobbole.com/86585/)
2. [自建内部pypi镜像](http://help.angzhou.com/2016/09/13/internal-pypi-server/)
    - 配置解释详细
3. [官方启动脚本](https://pypi.python.org/pypi/pypiserver#managing-automated-startup)

**注意: 这个服务(pypiserver-1.2.0)只能作为缓存存在, 无法实现定时同步...好坑啊.**

首先你得有python和pip环境...

然后直接pip安装

```
pip install pypiserver
```

安装的过程中会会出现BeautifulSoup版本不匹配的报错

```log
Downloading/unpacking BeautifulSoup<=3.0.9999 (from pypimirror)
  Could not find a version that satisfies the requirement BeautifulSoup<=3.0.9999 (from pypimirror) (from versions: 3.2.0, 3.2.1)
Cleaning up...
No distributions matching the version for BeautifulSoup<=3.0.9999 (from pypimirror)
Storing debug log for failure in /home/topgear/.pip/pip.log

```

将`PYTHON路径/dist-packages/pypimirror-1.0.16-py2.7.egg/EGG-INFO/requires.txt`文件修改一下即可.

其中有一行是: `BeautifulSoup<=3.0.9999`, 将版本号去掉，即改成: `BeautifulSoup`. 改完之后再重新执行上面的安装命令即可.

> 实在不行, 直接去[官网](https://pypi.python.org/pypi/pypiserver/1.2.0)下载安装包, 本地安装也是一样的.

安装完成后, 首先查看其帮助文档

```
pypi-server --help
```

如果安装成功, 但是报`-bash: pypi-server: command not found`错误的, 创建`/usr/bin/pypi-server`文件, 写入如下内容

```py
#!/usr/bin/python

# -*- coding: utf-8 -*-
import re
import sys

from pypiserver.__main__ import main

if __name__ == '__main__':
    sys.argv[0] = re.sub(r'(-script\.pyw?|\.exe)?$', '', sys.argv[0])
    sys.exit(main())

```

------

在启动之前, 创建服务启动配置文件`/usr/lib/systemd/system/pypiserver.service`.(**未验证...**)

```ini
[Unit]
Description=Pypi Mirror
[Service]
Environment=PACKAGES_DIRECTORY=/opt/pypi
ExecStart=/usr/bin/pypi-server \
    --port 8080 \
    --log-file /var/log/pypi.log \
    PACKAGES_DIRECTORY
Restart=on-failure
[Install]
WantedBy=multi-user.target
```

`PACKAGES_DIRECTORY`: 最重要, 是下载的包要放置的位置, 这里指定`/opt/pypi`, 如果不存在最好手动创建.

`--port`指定pypi服务监听的端口.

`--fallback-url`: 指定如果私有源中没有目标安装包时去哪里下载. 默认是官网源`http://pypi.python.org/simple`, 也可以指定国内的镜像源`http://mirrors.aliyun.com/pypi/simple/`.

...不想写了, 没意思