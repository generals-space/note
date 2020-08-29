# Linux 源码安装 Python 环境

<!-- 
<!tags!>: <!源码安装!> 
<!keys!>: eSy'39Nmarrypjlo 
-->

参考文章

1. [urllib2.URLError: < urlopen error unknown url type: https >](http://blog.csdn.net/hewy0526/article/details/9202523)

2. [get-pip源码](https://bootstrap.pypa.io/get-pip.py)

3. [关于在centos下安装python3.7.0以上版本时报错ModuleNotFoundError: No module named '_ctypes'的解决办法](https://blog.csdn.net/qq_36416904/article/details/79316972)

实验环境:

- 系统版本: docker的CentOS6镜像...够精简了吧, 其自带的python版本为`2.6.6`

- 目标版本: python官网最新的2.x为`2.7.12`

> 经实验, `2.7`与`3.5`, `3.6`都可以使用这篇文章中的流程进行编译.

## 1. python编译, 安装及环境配置

### 1.1 安装依赖包

django可能会用到sqlite库, 其实sqlite库已经集成到python里, 但是如果不安装`sqlite-devel`就无法使用. 

`openssl-devel`则是为了使用python的httplib库去获取`https`类型的url的内容.

```
$ yum install -y python-devel openssl-devel sqlite-devel
```

### 1.2 配置并编译

```shell
wget https://www.python.org/ftp/python/2.7.12/Python-2.7.12.tgz
tar -zxf ./Python-2.7.12.tgz
cd Python-2.7.12
```

在配置编译选项之前, 先修改一个地方. 解开对SSL这一段的注释.

```shell
vim Python-2.7.12/Modules/Setup.dist

# Socket module helper for SSL support; you must comment out the other
# socket line above, and possibly edit the SSL variable:
SSL=/usr/local/ssl
_ssl _ssl.c \
        -DUSE_SSL -I$(SSL)/include -I$(SSL)/include/openssl \
        -L$(SSL)/lib -lssl -lcrypto

```

如果不进行此操作的话, 之后有用到`urllib2`的时候可能会报错如下, 无法解析`https`类型的url

```
urllib2.URLError: <urlopen error unknown url type: https>
```

------

接下来配置编译选项并进行安装.

```
cd Python-2.7.12
## 用到的编译选项很少, 与php不同. 缺少的模块可以后期通过python的包管理工具进行安装.
./configure --prefix=/usr/local/python2.7
make && make install
```

将`/usr/local/python2.7/bin`加入到环境变量.

### 3. 替换旧版本(可选)

> 注意: 确认`/usr/bin`下面python, python2.6同时存在. 否则不要直接rm!!!

```
ll /usr/bin | grep python
lrwxrwxrwx. 1 root root      34 Jul 14 11:09 python
lrwxrwxrwx. 1 root root       6 Jul  1 20:21 python2 -> python
-rwxr-xr-x. 1 root root    4864 Jul 23  2015 python2.6

## 取代原有的python版本
rm /usr/bin/python
ln -s /usr/local/python2.7/bin/python2.7 /usr/bin/python

## 修改yum命令所使用的python版本
## 将原本的python改为python2.6, 这样yum还可以继续使用.
vim /usr/bin/yum
#!/usr/bin/python2.6
```

## 2. pip

源码安装python完成后, 不能再使用yum安装`pip`工具了, 会出错(不过如果像修改`yum`命令那样修改`pip`理论上应该也可以的, 不过那样做就没有意义了). 我们想要以后通过pip安装的python包都放在自定义的路径`/usr/local/python2.7`下, 所以还是需要使用当前的python安装.

~~点击python官网顶部的[PyPI](https://pypi.python.org/pypi)链接, 搜索`pip`, 下载其源码(当前最新版本为`8.1.2`)~~

~~按照pip源码目录下的`README.rst`中`Installation`提示的[网址](https://pip.pypa.io/en/stable/installing/), 找到了安装pip的方法. 虽然这个文档里说明了python2.7.9以上的版本内置了pip, 但我在生成的`/usr/local/python2.7`目录下并没有找到~~. 所以还是按照文档里的`get-pip.py`脚本手动安装.

下载[get-pip.py](https://bootstrap.pypa.io/get-pip.py)文件到本地, 然后执行`/usr/local/python2.7/bin/python2.7 get-pip.py`.

通过这种指定新版python绝对路径安装pip的方法, pip会出现在新版python路径下的bin目录中. 之后指定使用这个pip安装的python包都会出现在新版python目录下. 与之前旧版本python环境不会冲突, 不过使用的时候就需要指定可执行文件的绝对路径了, 加入PATH环境变量是个不错的选择.

pip在`/usr/local/python2.7/bin`目录下可能是`pip2`(使用python3执行`get-pip.py`时就是pip3了). 你可能需要建个软链接或者重命一下名...

然后设置pip的镜像源(位置在`~/.pip/pip.conf`), 用以加速国内包的下载速度.

```ini
[global]
trusted-host = pypi.douban.com
index-url = http://pypi.douban.com/simple/
```

完成.

## 3. FAQ

### 3.1 关于在centos下安装python3.7.0以上版本时报错ModuleNotFoundError: No module named '_ctypes'的解决办法

3.7版本需要一个新的包libffi-devel，安装此包之后再次进行编译安装即可。

```
$ yum install libffi-devel -y
$ make && make install
```