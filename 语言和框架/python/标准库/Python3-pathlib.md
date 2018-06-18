# Python3-pathlib

目前我对目录相关操作一般是查找与当前文件所在路径相对的另一个文件的路径, 然后去加载或执行. python2中的abs等操作比较麻烦, 但能满足我的需求.

```py
from pathlib import Path
```

`Path.cwd()`: 类似shell的`pwd`命令, 使用`python xxx.py`执行一个脚本时, `Path.cwd()`总是返回在终端命令行中所在的目录路径, 不管有没有`python ./abc/def/xxx.py`这样的相对路径的前缀.

`Path(__file__)`可以得到执行`python ./abc/def/xxx.py`时的相对路径(包括文件名, 如`./abc/def/xxx.py`), 它的值等同于`sys.argv[0]`

`Path(__file__).parent`可以得到`Path(__file__)`文件的所在目录, 这个值是不会包含文件名的.

`currDir = Path.cwd().joinpath(Path(__file__).parent)`