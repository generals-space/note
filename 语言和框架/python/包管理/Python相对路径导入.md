# Python相对路径导入

参考文章

1. [相对导入中Attempted relative import in non-package问题](https://www.cnblogs.com/qingyuanjushi/p/6400067.html)

2. ~~[Python：Relative import 相对路径 ValueError: Attempted relative import in non-package](http://blog.csdn.net/chinaren0001/article/details/7338041)~~
    - 参考文章2中提到的几种方法都是错误的.

关于`Attempted relative import in non-package`这个错误, 以下是一个简单示例.

```
.
├── config
│   ├── __init__.py
│   ├── list.txt
│   └── settings.py
├── main.py
└── others
    ├── __init__.py
    └── show.py
```

整个工程的执行流程为:

`python main.py`导入config包中的settings文件, 定义其中的一些变量作为配置, 然后调用others包中的show.py打印这些变量. 我们其实希望在show.py文件中导入settings.py模块的, 但是这就出现了上述错误.

详细代码

`main.py`

```py
from pathlib import Path
## 这里有两种相对路径的导入方法
from config.settings import config
from others import show

config['name'] = 'jiangming'
config['file'] = open(str(Path.cwd().joinpath(Path(__file__).parent)) + '/config/list.txt')

show.getContent()
```

`config/settings.py`

```py
#!/usr/bin/python3.6
#!encoding:utf-8

config = {
    'name': 'general',
    'file': None,
}
```

...`list.txt`这个文件的内容随便了.

`others/show.py`

```py
## from ..config import settings
## from ..config.settings import config
## 上面的两种导入方式都会报如下错误
## ValueError: attempted relative import beyond top-level package
def getContent():
    cnt = config['file'].read()
    print(cnt)
```

执行main.py时, show.py中的导入操作行就会报错, 原因在于, `from ..`涉及到了顶层模块`main.py`所在的路径, 而顶层模块不允许被子模块导入(不管是当作包还是当作模块, 都不允许), 就算在路径中也不允许包含顶层模块所在路径, 就算顶层模块中包含`__init__.py`文件也不行.

> 当然, 顶层模块中的文件之间是可以相互导入的.

那么, 我们的想法还有没有办法实现呢? 

当然是可以的, 只要我们调整一个目录结构, 让子模块之间的引用不再涉及到顶层模块所在路径即可.

如下:

```
.
├── main.py
└── sub
    ├── __init__.py
    ├── config
    │   ├── __init__.py
    │   ├── list.txt
    │   └── settings.py
    └── others
        ├── __init__.py
        └── show.py
```

这样, 在`show.py`中导入`settings`文件时, 将不经过顶层模块所在目录, 就可以了.