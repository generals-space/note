# Saltstack常用命令

参考文章

1. [Saltstack系列3：Saltstack常用模块及API](http://www.cnblogs.com/MacoLee/p/5753640.html)

查看所有可用模块

```
$ salt '*' sys.list_modules
172.32.100.233:
    - acl
    - aliases
    - alternatives
...省略
```

查看模块下所有可用方法

```
$ salt '*' sys.list_functions sys
172.32.100.233:
    - sys.argspec
    - sys.doc
    - sys.list_functions
    - sys.list_modules
...省略
```
