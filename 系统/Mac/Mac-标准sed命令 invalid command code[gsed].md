参考文章

1. [sed: 1: “…”: invalid command code](https://www.logcg.com/archives/3142.html)
    - mac下`sed -i`的使用方法
2. [Mac 下如何使用sed -i命令](https://www.cnblogs.com/chunzhulovefeiyue/p/6561497.html)
3. [Mac下的sed命令](https://blog.csdn.net/sun_wangdong/article/details/71078083)
    - `brew install gnu-sed --with-default-names`, 但是`MacOS 10.15.4`已经失效了.
4. [[Shell]sed命令在MAC和Linux下的不同使用方式](https://www.cnblogs.com/dzblog/p/6278546.html)
    - mac下sed的详细用法
5. [How to use GNU sed on Mac OS 10.10+, 'brew install --default-names' no longer supported]()
    - `brew install gnu-sed`安装linux版的sed命令, 不过命令名变成了`gsed`

## 场景描述

在Mac下执行sed命令报错

```console
$ sed -i '0,/\<h1\>/d' ./*.html
sed: 1: ".//028.html": invalid command code .
```

按照参考文章1中, 为`sed -i`添加了额外的参数`''`, 虽然不报错了, 但是根本不生效.

## 解决办法

没必要为了一个小工具再学一遍mac下的sed命令使用方法, 使用`brew`安装linux sed命令即可.

```
brew install gnu-sed
```

安装好的命令名为`gsed`.
