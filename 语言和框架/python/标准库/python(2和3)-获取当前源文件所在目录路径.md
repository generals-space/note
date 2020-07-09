# python(2和3)-获取当前源文件所在目录路径

参考文章

1. [python获取当前文件路径以及父文件路径](https://www.cnblogs.com/yajing-zh/p/6807968.html)
    - python2

python2

```py
import os

curr_dir = os.path.abspath(os.path.dirname(__file__))
target_file = 
```

python3

```py
from pathlib import Path

currDir = Path.cwd().joinpath(Path(__file__).parent)
target_file = 
```

> 获取用户执行脚本时所在的目录可以使用`os.getcwd()`函数, 具体可见[note-cloud]()仓库的`k2file.py`脚本.

