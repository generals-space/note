# python 语言级别问题

## 1.

```
python /home/paramiko.py
Traceback (most recent call last):
  File "/home/paramiko.py", line 11, in <module>
    import paramiko
  File "/home/paramiko.py", line 18, in <module>
    ssh = paramiko.SSHClient()
AttributeError: 'module' object has no attribute 'SSHClient'
```

场景描述:

home目录下写了一个调用paramiko模块的小程序, python执行的时候报上述错误.

问题分析:

自定义的程序文件名不能与代码中import语句中导入的模块名相同. 代码里`import paramiko`, 那程序文件就不能再叫`paramiko.py`了.

解决方法:

将程序文件改个名字即可.

## 2. 关于三引号"""

```python
>>> print("""This string has three quotes!
... Look at what it can do!""")
This string has three quotes!
Look at what it can do!
```

使用三引号实现换行.

## 3. `__name__`变量

python中, 每个py文件都是一个模块, 也同时是一个可执行文件(即包含main方法). 因此, 对每个py文件, 可以单独运行, 也可以`import`它给其他程序使用, 这两种情况不一样. 为了区分这两种情况, 可以使用`__name__`属性.

当py文件是直接运行时, `__name__ = "__main__"`,

当此文件被当作模块导入时, `__name__ = 其本身模块名`

## 4.

`RuntimeError: thread.__init__() not called`

出现原因: 类中`__init__()`方法中没有初始化Thread原对象.

解决原因: 在`__init__()`方法中加入`threading.Thread.__init__(self)`即可.

## 5. 自定义模块搜索路径

### 5.1 引用其他文件中的函数

```py
from wsgiref.simple_server import make_server
```

其中`wsgiref.simple_server`是`/usr/lib/python2.7/wsgiref/simple_server.py`文件，`make_server`是该文件中的一个函数。

### 5.2 

当我们试图加载一个模块时，python会在指定的路径下搜索同名的.py文件，如果找不到，就会报错：

```py
>>> import mymodule
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ImportError: No module named mymodule
```

### 5.3 

默认情况下，Python解释器会搜索当前目录，所有已安装的内置模块和第三方模块，搜索路径存放在sys模块的path变量：

```py
>>> import sys
>>> sys.path
['', '/Library/Python/2.7/site-packages/pycrypto-2.6.1-py2.7-macosx-10.9-intel.egg', '/Library/Python/2.7/site-packages/PIL-1.1.7-py2.7-macosx-10.9-intel.egg', ...]
```

### 5.4 

如果要添加自己的搜索目录，有两种方法：

一是直接修改sys.path，添加要搜索的目录：

```py
>>> import sys
>>> sys.path.append('/Users/michael/my_py_scripts')
```

这种方法是在运行时修改，运行结束后失效。

二是设置环境变量`PYTHONPATH`，该环境变量的内容会自动添加到模块搜索路径中。设置方式与设置PATH环境变量类似。注意：**只需要添加你自己的搜索路径，Python自己本身的搜索路径不受影响**。

## 6.

参考文章

[解决Python2.7的UnicodeEncodeError: ‘ascii’ codec can’t encode异常错误](http://wangye.org/blog/archives/629/)

在打包python程序时, 报如下错误

```
$ python setup.py sdist
Traceback (most recent call last):
...
UnicodeEncodeError: 'ascii' codec can't encode characters in position 0-78: ordinal not in range(128)
```

本来`setup.py`与出错文件都注释了`#!encoding: utf-8`, 所以错误不可能是这个. 貌似是因为python在处理工程源文件时使用的是`ascii`编码, 与utf-8不兼容, 所以需要设置python编码环境默认为utf-8才行.

可以通过在命令行里执行如下代码得到python的默认编码, 一般是ascii.

```python
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

```python
import sys
reload(sys)
sys.setdefaultencoding('utf-8')
```

一般执行哪个文件报出的`UnicodeEncodeError`, 就在哪个文件开头加上这三行, 就可以解决问题了.