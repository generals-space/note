# Python问题 - 环境

## 1. pip镜像源

参考文章

[解决pip无法使用http的源](http://www.tuicool.com/articles/2uqEFr)

### 情景描述

在使用pip设置douban等代理地址安装python包时报错如下:

```
[root@localhost docs]# pip install -i http://pypi.douban.com/simple/ django==1.6
You are using pip version 7.1.0, however version 8.1.2 is available.
You should consider upgrading via the 'pip install --upgrade pip' command.
Collecting django==1.6
  The repository located at pypi.douban.com is not a trusted or secure host and is being ignored. If this repository is available via HTTPS it is recommended to use HTTPS instead, otherwise you may silence this warning and allow it anyways with '--trusted-host pypi.douban.com'.
  Could not find a version that satisfies the requirement django==1.6 (from versions: )
No matching distribution found for django==1.6
```

报错时, 只有下面一段是红色的, 上面的是黄色的, 原以为不重要就没有仔细看.

```
Could not find a version that satisfies the requirement django==1.6 (from versions: )
No matching distribution found for django==1.6
```

最初以为是这个镜像源没有`django`的1.6版本的原因, 换了好几个源都不行. 后来才仔细读了一下黄色的部分, 原来是因为`pip`认为这些源不可信任, 并且嫌弃它们没有`https`...

### 解决办法

修改`~/.pip/pip.conf`或是`/etc/pip.conf`, 目录或文件不存在的, 自行创建即可.

```
[global]
index-url = http://pypi.douban.com/simple
[install]
trusted-host = pypi.douban.com
```

```
[global]
index-url = http://mirrors.aliyun.com/pypi/simple/
[install]
trusted-host=mirrors.aliyun.com
```

## 2. 虚拟环境virtualenv

像python, php, apache这种环境, 权限一般掌握在root用户手中, 安装插件需要有root权限. 而软件工程一般运行在普通用户权限下(为了安全考虑), 当不同用户需要不同模块时, 就需要root用户为每一个用到这些环境的普通用户安装所有的插件.

幸而python为解决这个问题提供了一个`virtualenv`的虚拟环境, 每个用户可以拥有其本身的packages库. 多个python环境相互独立, 互不影响. 它的作用如下:

- 在没有root权限的情况下安装新模块

- 不同应用可以使用不同的模块版本

- 模块升级不影响其他应用

安装/使用方式:

```
$ pip install virtualenv
## 添加普通用户general, 创建虚拟python环境目录
$ useradd general
$ mkdir /home/general/virpython

## 使用virtualenv对虚拟环境初始化, 注意virtualenv的位置, 根据实际情况决定
$PYTHON/bin/virtualenv /home/general/virpython
## 查看virpython的目录结构
$ ls /home/general/virpython
bin  include  lib  pip-selfcheck.json
$ ls /home/general/virpython/bin
activate    activate_this.py   pip    python    python-config    activate.csh    easy_install    pip2    python2    wheel    activate.fish    easy_install-2.7  pip2.7  python2.7

## 然后将virpython目录的属主改为general
$ chown -R general:general /home/general/virpython

## 使用虚拟用户安装模块, 注意, 直接在普通用户下安装会出错
$ su - general
$ cd /home/general/virpython
## 进入虚拟环境, 不单单是普通用户环境
$ source ./bin/activate
(virpython) [general@localhost virpython] $ pip install django==1.6

## 退出虚拟环境. 没有找到deactivate的位置, 但是可以使用.
## 可以使用exit命令, 但这同时会退出普通用户的shell, 有些不方便
(virpython) [general@localhost virpython] $ deactivate
```

## 3. 

```
$ pip install pymongo==2.7.2
shell-init: error retrieving current directory: getcwd: cannot access parent directories: No such file or directory
shell-init: error retrieving current directory: getcwd: cannot access parent directories: No such file or directory
Exception:
Traceback (most recent call last):
  File "/opt/ew4-callback-server/lib/python2.7/site-packages/pip/basecommand.py", line 215, in main
    status = self.run(options, args)
  File "/opt/ew4-callback-server/lib/python2.7/site-packages/pip/commands/install.py", line 312, in run
    wheel_cache
  File "/opt/ew4-callback-server/lib/python2.7/site-packages/pip/basecommand.py", line 276, in populate_requirement_set
    wheel_cache=wheel_cache
  File "/opt/ew4-callback-server/lib/python2.7/site-packages/pip/req/req_install.py", line 187, in from_line
    path = os.path.normpath(os.path.abspath(name))
  File "/opt/ew4-callback-server/lib/python2.7/posixpath.py", line 371, in abspath
    cwd = os.getcwd()
OSError: [Errno 2] No such file or directory
```

问题描述

装了一个`pymongo`的包, 但版本太高, 将原版本`3.4.0`卸载掉准备重新装新包, `pip install pymongo==2.7.2`, 终端报错.

原因分析

卸载前进入了`site-packages/pymongo`目录, 卸载后没有退出, 所以安装其他版本的包时显示`No such file or directory`, 因为当前所处的目录已经不存在了.

解决方法

切换目录, 然后重新执行`install`操作.