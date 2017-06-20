# Python包管理工具setuptools详解

> <!tags!>: <!setuptools!> <!setup.py!>

参考文章

1. [Python包管理工具setuptools详解及entry point](http://blog.csdn.net/pfm685757/article/details/48651389#t3)

2. [python项目打包发布总结](http://blog.csdn.net/tw18761720160/article/details/52826450)

3. [一文教会你正确打包Python程序](http://www.tuicool.com/articles/Ivuaaq)

4. [关于python中的setup.py](http://blog.csdn.net/xluren/article/details/41114779)

5. [PyPI打包和分发文档](https://packaging.python.org/distributing/)

打包的目的一般是, 要么是想提供类似`os`, `sys`, `getopt`这种类库, 可在程序中调用, 或者像是`supervisord`一样有可执行文件供用户直接使用.

python自带一个基本的安装工具`distutils`, 适用于非常简单的应用场景使用, 不支持依赖包的安装. 通过`distutils`来打包，生成/安装python包等工作，需要编写名为`setup.py`的python脚本文件。

`setuptools`是对`distutils`的增强, 其使用方法与`distutils`相似, 本文不介绍`distutils`的使用方法.

## 1. 最简示例 - 单文件发布

1. `hello.py`

```py
#!/usr/bin/env python
#!encoding: utf-8

print('hello world')
```

2. `setup.py`

```py
#!/usr/bin/env python
#coding: utf-8

from setuptools import setup
setup(
    name = 'hello',          ## 这个是最终打包的文件名
    version = '1.0.0',
    py_modules = ['hello'], ## 要打包哪些.py文件
)
```

目录下只有这两个文件

```
$ ls
hello.py  setup.py
```

然后开始打包

```
$ python setup.py sdist
```

之后的目录结构为

```
$ tree
.
├── dist
│   └── hello-1.0.0.tar.gz
├── hello.egg-info
│   ├── dependency_links.txt
│   ├── PKG-INFO
│   ├── SOURCES.txt
│   └── top_level.txt
├── hello.py
└── setup.py

2 directories, 7 files
```

`dist`目录下就是可以使用`pip`安装的包了. `pip install hello-1.0.0.tar.gz`后, 可以在python的lib目录下找到这个py文件.

```
$ pip install hello-1.0.0.tar.gz
$ ls /usr/lib/python2.7/site-packages | grep hello
hello-1.0.0.dist-info
hello.py
hello.pyc
```

然后就可以在程序里通过`import hello`引用这个包了.

## 2. 进阶示例 - 包目录发布

工程较大时, 结构也会比较复杂, 会有许多源文件. 我们需要打出一个目录的工程包来. 首先看一下目录结构.

```
$ tree 
.
├── person
│   ├── greet.py
│   ├── hello.py
│   └── __init__.py
└── setup.py

1 directory, 5 files
```

它们的内容分别为

1. `person/hello.py`

```py
#!/usr/bin/env python
#!encoding: utf-8

def say_hello():
    print('hello world')
```

2. `person/greet.py`

```py
#!/usr/bin/env python
#!encoding: utf-8

from hello import say_hello

if __name__ == '__main__':
    say_hello()
```

3. `person/__init__.py`是一个空文件.

4. `setup.py`和之前的有一点不同

```py
#!/usr/bin/env python
#coding: utf-8

from setuptools import setup
setup(
    name = 'person',          ## 这个是最终打包的文件名
    version = '1.0.0',
    packages = ['person']
)
```

这里`setup()`中的`packages`参数, 是指定要打入包中的目录名. 同样使用`python setup.py sdist`然后用`pip`安装后, 可以按照如下方式使用.

```
[root@localhost dist]# python 
Python 2.7.5 (default, Aug 18 2016, 15:58:25) 
[GCC 4.8.5 20150623 (Red Hat 4.8.5-4)] on linux2
Type "help", "copyright", "credits" or "license" for more information.
>>> import person
>>> from person.hello import say_hello
>>> say_hello()
hello world
>>> 
```

**注意:**

`setup.py`文件中`setup()`方法的`py_modules`与`packages`并不冲突, 可以同时使用, 这样可以同时发布包目录与单文件. 但在官方文档中没有看到`py_modules`这个参数, 但是单文件发布只能使用这个, 不知道怎么回事.

------

`packages`参数的值可以不手动填写, 使用`setuptools`模块中的`find_packages()`方法, 可以自动搜索当前目录下的package并发布.

## 3. 官方文档解读

### 3.1 setup()方法参数

`package_data`: 默认工程目录中只有`.py`的文件会被打到发布包中, 但有时还需要加入一些别的文件, 比如`version.txt`文件等, 就需要把这种文件指定到`package_data`中.

```py
package_data={
    'sample': ['package_data.dat'],
},
```

> `package_data`字典成员的键名并没有什么特殊作用, 真的只是一个名称而已, 值所表示的静态文件将被打包到工程的原来的路径下. 不过貌似`package_data`这个键并不是很有效, 大部分人都使用`MANIFEST.in`文件作为替代了.

`data_files`: 除了将额外的文件打到发布包里面, 有时候pip安装时还要将工程的某些文件放到其他指定目录, 比如示例配置文件, 一般拷贝到`/etc`目录下.

```py
data_files=[
    ('config', ['conf/config_file'])
],
```

> 注意: `data_files`参数貌似不支持将元组的第1个成员写作绝对路径. 静态文件在`install`操作中将被安装到以`python/site-packages`为基准的相对路径下, 如果目标目录(如`config`)不存在, 会自动创建. 可以参考stackoverflow的[这篇问答](http://stackoverflow.com/questions/40588634/how-to-install-data-files-to-absolute-path)

## 4. 二进制发布

`python setup.py sdist`发布的是源码包, 当我们需要打成打包rpm包呢? 如果要发布成`exe`程序呢?

下面的命令可以查看`sdist`子命令支持什么样的发布格式, 其中`wininst`是exe, `rpm`自然就是rpm格式了.

```
$ python setup.py bdist --help-formats
List of available distribution formats:
  --formats=rpm      RPM distribution
  --formats=gztar    gzip'ed tar file
  --formats=bztar    bzip2'ed tar file
  --formats=ztar     compressed tar file
  --formats=tar      tar file
  --formats=wininst  Windows executable installer
  --formats=zip      ZIP file
  --formats=msi      Microsoft Installer
  --formats=egg      Python .egg file
```