# Python动态导入

参考文章

1. [python动态加载包的方法小结 - 脚本之家](http://www.jb51.net/article/82586.htm)
    - `__import__()`内置方法
    - `impl`, 貌似是个第3方包
    - `exec "import string as stringmodule"`, 直接不推荐使用吧...
2. [Python Cookbook（第2版）中文版 2.21节 - 动态地改变Python搜索路径](http://blog.csdn.net/luyafei_89430/article/details/9240603)
    - 修改`sys.path`列表
3. [python中重新导入模块](https://blog.csdn.net/fldx/article/details/89175051)
    - `importlib.reload()`
4. [Python3 动态导入模块的两种方式](https://www.cnblogs.com/bert227/p/9786784.html)
    - `__import__()`内置方法
    - `importlib.import_module()`
5. [__import__ 与动态加载 python module](http://python.jobbole.com/87492/)
    - `__import__`的底层原理
6. [__import__详解](https://www.jianshu.com/p/e7ee9b2c83b9)
    - `from module import submodule`形式如何使用`__import__`动态加载

## 1. 动态改变搜索路径

模块必须处于Python搜索路径中才能被导入, 但你不想设置个永久性的大而全的路径, 因为那样可能会影响性能, 所以可以动态地改变这个路径. 

只需简单地在Python的`sys.path`中`append`目标目录即可(如果想让目标目录处于`sys.path`处于最前, 也可以用`insert`), 即时生效, 不过要小心重复的情况.

如果要移除, 自然也可以通过列表的`pop`方法完成.

## 2. 动态加载

关于内置方法`__import__`就不介绍了, 不建议使用.

```py
import importlib
## 模块名的字符串
target_module = 'app.task' 
## 导入的就是需要导入的那个task
task = importlib.import_module(target_module) 
## Run() 为 task 模块中的方法
task.Run("Bert") 
```

## 3. 重新加载

一旦导入一个模块之后, 这个模块之后的变动都不会再被察觉了, 继续使用`importlib.import_module(target_module)`也无法改变(应该是发现`target_module`模块已经导入过, 所以不再重复导入吧), 我们需要使用`importlib.reload()`方法完成这个目的.

```py
import importlib
import target_module
## ...之后 target_module 的内容发生变动
importlib.reload(target_module)
```

需要注意的是, 貌似`importlib.reload()`与目标模块的`import`语句要在同一个上下文中, 比如如下代码就不会生效.

```py
import importlib
import target_module

def reload():
    ## 这里会报 raise TypeError("reload() argument must be a module")
    importlib.reload(target_module)
```

必须要写成

```py
import importlib

def reload():
    import target_module
    ## 这里会报 raise TypeError("reload() argument must be a module")
    importlib.reload(target_module)
```
