# Python-os模块listdir与walk遍历目录

参考文章

1. [python使用os.listdir和os.walk获得文件的路径](https://www.cnblogs.com/jiaxin359/p/7324077.html)

`os.listdir(path)`: 返回一个列表, 成员为目标目录下的子目录和文件. 不会遍历子目录, 并且由于成员都是字符串类型, 无法判断是目录还是文件. 如果参数不是目录, 会抛出`NotADirectoryError`异常.

`os.walk(path)`: 返回的是一个生成器, 使用`list()`列表化后会发现ta的成员是一个三元元组, 分别为`dirpath`, `dirnames`和`filenames`. 如下

```
[('/root/sites', ['testdir'], ['ckck.tv', '97daimeng.com']), ('/root/sites/testdir', [], ['testfile'])]
```

对应`/root/sites`目录结构如下

```
.
├── sites
    ├── 97daimeng.com
    ├── ckck.tv
    └── testdir
        └── testfile

2 directories, 4 files
```

可以看到, `os.walk()`还可以遍历子目录内容, 三元组内容分别表示当前目录路径, 当前目录下的目录列表和文件列表.

如下代码可以获得目录目录下的所有文件的绝对路径, 包括子目录.

```py
for dirpath, dirnames, filenames in os.walk(path):
    for filename in filenames:
        print(pathlib.Path.join(dirpath, filename))
```