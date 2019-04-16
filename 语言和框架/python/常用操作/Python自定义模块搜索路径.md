# 自定义模块搜索路径

当我们试图加载一个模块时，python会在指定的路径下搜索同名的.py文件，如果找不到，就会报错：

```py
>>> import mymodule
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ImportError: No module named mymodule
```

默认情况下，Python解释器会搜索当前目录，所有已安装的内置模块和第三方模块，搜索路径存放在`sys`模块的`path`变量：

```py
>>> import sys
>>> sys.path
['', '/Library/Python/2.7/site-packages/pycrypto-2.6.1-py2.7-macosx-10.9-intel.egg', '/Library/Python/2.7/site-packages/PIL-1.1.7-py2.7-macosx-10.9-intel.egg', ...]
```

如果要添加自己的搜索目录，有两种方法：

一是直接修改`sys.path`，添加要搜索的目录：

```py
>>> import sys
>>> sys.path.append('/Users/michael/my_py_scripts')
```

这种方法是在运行时修改，运行结束后失效。

二是设置环境变量`PYTHONPATH`，该环境变量的内容会自动添加到模块搜索路径中。设置方式与设置PATH环境变量类似。注意：**只需要添加你自己的搜索路径，Python自己本身的搜索路径不受影响**。
