# python 语言级别问题处理

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


## 4.

`RuntimeError: thread.__init__() not called`

出现原因: 类中`__init__()`方法中没有初始化Thread原对象.

解决原因: 在`__init__()`方法中加入`threading.Thread.__init__(self)`即可.

## 5. 自定义模块搜索路径

### 5.1 <!已删除!>

### 5.2 

当我们试图加载一个模块时，python会在指定的路径下搜索同名的.py文件，如果找不到，就会报错：

```py
>>> import mymodule
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ImportError: No module named mymodule
```

### 5.3 

默认情况下，Python解释器会搜索当前目录，所有已安装的内置模块和第三方模块，搜索路径存放在`sys`模块的`path`变量：

```py
>>> import sys
>>> sys.path
['', '/Library/Python/2.7/site-packages/pycrypto-2.6.1-py2.7-macosx-10.9-intel.egg', '/Library/Python/2.7/site-packages/PIL-1.1.7-py2.7-macosx-10.9-intel.egg', ...]
```

### 5.4 

如果要添加自己的搜索目录，有两种方法：

一是直接修改`sys.path`，添加要搜索的目录：

```py
>>> import sys
>>> sys.path.append('/Users/michael/my_py_scripts')
```

这种方法是在运行时修改，运行结束后失效。

二是设置环境变量`PYTHONPATH`，该环境变量的内容会自动添加到模块搜索路径中。设置方式与设置PATH环境变量类似。注意：**只需要添加你自己的搜索路径，Python自己本身的搜索路径不受影响**。


## 7.

```
ValueError: Attempted relative import in non-package
```

相对导入中Attempted relative import in non-package问题, 在python2.7环境中源码安装PIL, 在`$PATHON/site-packages/PIL`目录下的一些文件中, 有类似如下方式导入PIL本身的一些变量.

```py
from . import VERSION PILLOW_VERSION
```

结果目标工程启动时要导入PIL, 就报了上述错误.

貌似是因为python版本不符所以不能用相对路径导入, 使用`from PIL import VERSION PILLOW_VERSION`就可以了.

关于这一点, 我觉得这篇文章[ValueError: Attempted relative import in non-package](http://www.cnblogs.com/DjangoBlog/p/3518887.html)讲得更加深入一点.

## 8.

```
Python 3.5 Socket TypeError: a bytes-like object is required, not 'str' 错误提示
```

参考文章

1. [Python 3.5 Socket TypeError: a bytes-like object is required, not 'str' 错误提示](https://blog.csdn.net/yexiaohhjk/article/details/68066843)