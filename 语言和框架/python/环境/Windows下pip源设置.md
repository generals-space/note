# Windows下pip源设置

参考文章

1. [Win下安装pip并更新源](http://blog.csdn.net/zhanghe2775115/article/details/51540584)

路径: `C:\Users\general\AppData\Local\pip`

创建文件: `pip.ini`

内容

```ini
[global]
index-url = http://mirrors.aliyun.com/pypi/simple/
[install]
trusted-host=mirrors.aliyun.com
```

立即生效.