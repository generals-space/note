# Python-UnicodeEncodeError问题

参考文章

1. [解决Python2.7的UnicodeEncodeError: ‘ascii’ codec can’t encode异常错误](http://wangye.org/blog/archives/629/)
2. [Python3出现UnicodeEncodeError: &apos;ascii&apos; codec can&apos;t encode characters问题](https://www.aliyun.com/jiaocheng/436946.html)
3. [Linux下locale: Cannot set LC_CTYPE to default locale: No such file or directory警告](https://yuchi.blog.csdn.net/article/details/108501289)
4. [英文Ubuntu安装中文包（locale）的方法](https://www.cnblogs.com/jefffyang/archive/2013/01/17/2864160.html)

## 1. python2.7

在打包python程序时, 报如下错误

```
$ python setup.py sdist
Traceback (most recent call last):
...
UnicodeEncodeError: 'ascii' codec can't encode characters in position 0-78: ordinal not in range(128)
```

本来`setup.py`与出错文件都注释了`#!encoding: utf-8`, 所以错误不可能是这个. 貌似是因为python在处理工程源文件时使用的是`ascii`编码, 与utf-8不兼容, 所以需要设置python编码环境默认为utf-8才行.

可以通过在命令行里执行如下代码得到python的默认编码, 一般是ascii.

```py
import sys
print sys.getdefaultencoding()
# 'ascii'
```

sys模块提供了一个`setdefaultencoding()`函数设置这个编码环境, 但是在命令行中执行`sys.setdefaultencoding('utf-8')`时会报如下错误

```
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
AttributeError: 'module' object has no attribute 'setdefaultencoding'
```

解决办法是, 在setdefaultencoding之前先reload()一遍sys模块, 完整的代码为

```py
import sys
reload(sys)
sys.setdefaultencoding('utf-8')
```

一般执行哪个文件(程序入口文件)报出的`UnicodeEncodeError`, 就在哪个文件开头加上这三行, 就可以解决问题了.

## 2. python3

上面说了在python2中的解决办法, 但是在python3.6环境下, 使用上述方法时却报了`reload`方法不存在的错误.

###  修改前

```
[root@efd527db107f ~]# python server.py 
Traceback (most recent call last):
  File "server.py", line 24, in <module>
    print(u'\u5f00\u59cb\u76d1\u542c...')
UnicodeEncodeError: 'ascii' codec can't encode characters in position 0-3: ordinal not in range(128)
```

### 修改后

```
[root@efd527db107f ~]# python server.py 
Traceback (most recent call last):
  File "server.py", line 6, in <module>
    reload(sys)
NameError: name 'reload' is not defined
```

环境是在docker容器中, 因为可能是国外的镜像, 所以很大可能出现编码问题. 后来按照参考文章2检查了`locale`的值.

```
[root@efd527db107f ~]# locale
LANG=
LC_CTYPE="POSIX"
LC_NUMERIC="POSIX"
LC_TIME="POSIX"
LC_COLLATE="POSIX"
LC_MONETARY="POSIX"
LC_MESSAGES="POSIX"
LC_PAPER="POSIX"
LC_NAME="POSIX"
LC_ADDRESS="POSIX"
LC_TELEPHONE="POSIX"
LC_MEASUREMENT="POSIX"
LC_IDENTIFICATION="POSIX"
LC_ALL=
```

然后`locale -a`可以查看所有可选的语系. 用`echo 'export LANG=en_US.UTF-8' >> ~/.bashrc`.

重新登录bash, 再次执行程序, 成功!

> 注意此时源码中已经不需要`reload()`这3行了.

------

20201124 更新

在某个 docker 容器(ubuntu)内又经历了这个问题, 这次修改了`.bashrc`后再进入 bash, 执行 python 程序还是出错.

我执行了一下`locale`, 发现了如下错误

```
locale: Cannot set LC_CTYPE to default locale: No such file or directory
locale: Cannot set LC_MESSAGES to default locale: No such file or directory
locale: Cannot set LC_ALL to default locale: No such file or directory
```

于是我找到了参考文章3, 按照ta的说法, 应该执行`locale-gen en_US.UTF-8`安装该编码类型.

但是容器里没有`locale-gen`, 于是我又搜索`ubuntu locale-gen 安装`, 于是找到了参考文章4. ta提到了如下两句命令, 主要就是没安装对应的语言包

```
cd /usr/share/locales
sudo ./install-language-pack zh_CN
```

但是`/usr/share/locales`这个目录根本不存在, 别说`install-language-pack`命令了. 

不过这给了我启发, 于是我搜索了一下

```console
$ apt-cache search language-pack-en
language-pack-en - translation updates for language English
language-pack-en-base - translations for language English
```

看到有`language-pack-en`就装了一下, 然后再执行`locale`就不报错了, 退出重新进入 bash, 执行 python 程序也可以了.
