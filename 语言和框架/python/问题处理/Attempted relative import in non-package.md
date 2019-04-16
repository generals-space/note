# ValueError: Attempted relative import in non-package

相对导入中Attempted relative import in non-package问题, 在python2.7环境中源码安装PIL, 在`$PATHON/site-packages/PIL`目录下的一些文件中, 有类似如下方式导入PIL本身的一些变量.

```py
from . import VERSION PILLOW_VERSION
```

结果目标工程启动时要导入PIL, 就报了上述错误.

貌似是因为python版本不符所以不能用相对路径导入, 使用`from PIL import VERSION PILLOW_VERSION`就可以了.

关于这一点, 我觉得这篇文章[ValueError: Attempted relative import in non-package](http://www.cnblogs.com/DjangoBlog/p/3518887.html)讲得更加深入一点.
