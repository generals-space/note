# Python动态导入

参考文章

1. [python动态加载包的方法小结 - 脚本之家](http://www.jb51.net/article/82586.htm)

2. [Python Cookbook（第2版）中文版 2.21节 - 动态地改变Python搜索路径](http://blog.csdn.net/luyafei_89430/article/details/9240603)

3. [__import__ 与动态加载 python module](http://python.jobbole.com/87492/)

## 1. 动态改变搜索路径

模块必须处于Python搜索路径中才能被导入, 但你不想设置个永久性的大而全的路径, 因为那样可能会影响性能, 所以, 你希望能够动态地改变这个路径. 

解决方案

只需简单地在Python的`sys.path`中`append`目标目录即可(如果想让目标目录处于`sys.path`处于最前, 也可以用`insert`), 即时生效, 不过要小心重复的情况.

如果要移除, 自然也可以通过列表的`pop`方法完成.

## 2. 动态加载模块

3种方法

### 2.1 系统函数`__import__()`

```py
stringmodule = __import__('string')
```

### 2.2 imp模块

```py
import imp 
stringmodule = imp.load_module('string',*imp.find_module('string'))
imp.load_source("TYACMgrHandler_"+app.upper(), filepath)
```

### 2.3 exec函数

```py
import_string = "import string as stringmodule"
exec import_string
```

`exec`还是非常不提倡的...