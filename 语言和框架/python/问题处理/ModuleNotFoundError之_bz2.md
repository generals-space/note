# ModuleNotFoundError之_bz2

参考文章

1. [missing python bz2 module](https://stackoverflow.com/questions/12806122/missing-python-bz2-module/12806325)

2. [python3添加对zlib和bz2的支持](http://www.pythontip.com/blog/post/4533/)

环境: 源码安装的python3.7

在安装celery时, 执行示例代码报错: `ModuleNotFoundError: No module named '_bz2'`

在python交互式命令行中import bz2库也会出错.

```py
Python 3.7.1 (default, Dec 30 2018, 12:41:46)
[GCC 4.8.5 20150623 (Red Hat 4.8.5-28)] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> import bz2
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
  File "/usr/local/python3.7/lib/python3.7/bz2.py", line 19, in <module>
    from _bz2 import BZ2Compressor, BZ2Decompressor
ModuleNotFoundError: No module named '_bz2'
```

然后按照参考文章1和2, 新安装了`zlib zlib-devel`和`bzip2 bzip2-devel`, 然后重新**编译python源码**.