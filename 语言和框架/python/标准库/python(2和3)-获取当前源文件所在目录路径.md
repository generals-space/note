# python(2和3)-获取当前源文件所在目录路径

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

