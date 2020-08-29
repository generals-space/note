
参考文章

1. [python学习笔记-zipimport.ZipImportError: can't decompress data; zlib not available](https://www.iteye.com/blog/zhoujianghai-1521993)
2. [python3.7移植到ARM开发板](https://blog.csdn.net/u012230668/article/details/89206857)
    - `apt-get install libffi-dev`

安装依赖

```
apt-get install libssl-dev libffi-dev
```

## FAQ

### 

```
zipimport.ZipImportError: can't decompress data; zlib not available
```

`./configure`与`make`执行没有出现问题, 但是在`make install`的时候报这个错误而退出了.

按照参考文章1, 执行如下命令安装`zlib`相关依赖.

```
apt-get install zlibc zlib1g-dev
```

貌似无需重新`make`, 直接`make install`即可成功, 且最终`python`也可以执行成功.

### 

```
Failed to build these modules:
_ctypes
```

参考文章2.

```
apt-get install libffi-dev
```
