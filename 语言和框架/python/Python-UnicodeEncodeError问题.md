# Python-UnicodeEncodeError问题


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

一般执行哪个文件(程序入口文件)报出的`UnicodeEncodeError`, 就在哪个文件开头加上这三行, 就可以解决问题了.
