参考文章

1. [解决libc.so.6: version `GLIBC_2.14' not found问题](http://blog.csdn.net/cpplang/article/details/8462768/)

2. [Linux误删C基本运行库libc.so.6急救方法](https://www.cnblogs.com/fjping0606/p/4551475.html)

3. [libc官方网站](http://www.gnu.org/software/libc/)

以[2.26](http://ftp.gnu.org/gnu/libc/glibc-2.26.tar.gz)版本为例.

查看是否真的没有目标版本的接口

```
$ strings /lib64/libc.so.6 | grep GLIBC
```

## 重新编译`libc`

升级这个库的风险太大, glibc是gnu发布的libc库, 即c运行库. glibc是linux系统中最底层的api，几乎其它任何运行库都会依赖于glibc. glibc除了封装linux操作系统所提供的系统服务外，它本身也提供了许多其它一些必要功能服务的实现…

总的来说，不说运行在linux上的一些应用，或者你之前部署过的产品，就是很多linux的基本命令，比如`cp`, `rm`, `ls`之类，都得依赖于它.

网上很多人有惨痛教训，甚至升级失败后系统退出后无法重新进入了.

------

编译的过程其实挺简单, 不过有两个注意点.

```
tar -zxf glibc-2.26.tar.gz && cd glibc-2.26
mkdir build && cd build
../glibc-2.15/configure  --prefix=/usr/local/libc
make
```

`configure`操作不能在源码包根目录执行, 再好先创建一个`build`子目录, 进入`build`目录, `configure`会在当前目录生成`Makefile`.

`prefix`参数也是要加的, 不然会提示不能与现有的`libc`库冲突.

编译完成后, 会在当前目录生成`libc-2.26.so`文件.

最有风险的就是libc库的替换了, 因为为了达到目的, 肯定要先删除`libc`的软链接, 但一旦删除, 几乎所有命令都失效了.

解决的办法是, 重定义`LD_PRELOAD`环境变量.

```
$ export LD_PRELOAD=/lib64/libc-2.26.so
$ ln -s /lib64/libc.so.6 /lib64/libc-2.26.so
```

幸亏`export`不依赖`libc`...

也有说写在同一行的.

```
$ LD_PRELOAD=/lib64/libc-2.26.so ln -s /lib64/libc-2.26.so  lib64/libc.so.6 
```

都可以.