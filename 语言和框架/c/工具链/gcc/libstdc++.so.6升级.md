# libstdc++.so.6升级

参考文章

1. [ 通用方法 解决/usr/lib64/libstdc++.so.6: version `CXXABI_1.3.8' not found的问题](http://blog.csdn.net/u012811841/article/details/77854581)

2. [CentOS 6.6 升级GCC G++ (当前最新版本为v6.1.0) (完整)](http://www.cnblogs.com/lzpong/p/5755678.html)

3. [Index of /gnu/gcc - gcc官方源码列表](http://ftp.gnu.org/gnu/gcc/)

## 1. /usr/lib64/libstdc++.so.6: version `CXXABI_1.3.8' not found问题

在运行某些软件的时候, 可能会得到标题中提到的错误, 这是因为系统的`libstdc++`共享库版本太低的缘故.

通过系统命令`strings`命令可以查看这些动态链接库包含哪些接口.

```
$ strings /usr/lib64/libstdc++.so.6.0.22 | grep CXXABI
CXXABI_1.3
CXXABI_1.3.1
CXXABI_1.3.2
...
CXXABI_1.3.10
CXXABI_1.3.3
```

如果报错中显示的共享库中真的没有需要的接口, 就只能重新编译这些库了, 因为一般系统原装的库不会提供更新版本...

## 2. 编译GCC

[GNU官方网站](http://ftp.gnu.org/gnu/gcc)里面有所有的gcc版本供下载, 这里以[6.4.0](http://ftp.gnu.org/gnu/gcc/gcc-6.4.0/gcc-6.4.0.tar.gz)为例.

```
tar -zxf gcc-6.4.0.tar.gz
cd gcc-6.4.0
```

### 下载供编译需求的依赖项

`gcc-6.4.0/contrib/download_prerequisites`这个神奇的脚本文件会帮我们下载、配置、安装依赖库, 可以节约我们大量的时间和精力, 直接执行即可.

注意: 这个工具不能帮我们安装`gcc`和`gcc-c++`工具, 还是需要手动安装的.

### 建立一个目录供编译出的文件存放

```
$ mkdir gcc-build-6.4.0 && cd gcc-build-6.4.0
```

### 生成Makefile文件

```
$ ../configure -enable-checking=release -enable-languages=c,c++ -disable-multilib
```

将在当前目录下生成`Makefile`文件, 之后编译后的结果也在这个目录下.

### 开始编译

```
$ make -j4
```

注意: ...非常耗时

### 安装

编译完成后不要急着`make install`, 一般我们是为升级而来的, `make install`的结果有可能影响较大, 所以只需要替换一些指定的文件即可.

一般是`/usr/lib64/libstdc++.so.6.xxxx`或`/usr/lib/libstdc++.so.6.xxxx`这两个地方的库.

编译完成的文件在`x86_64-pc-linux-gnu`目录下, 用`find`命令搜索一下即可找到.

## FAQ

### 1. `make[1]: *** [stage1-bubble] Error 2`

编译libstdc++时, `make`操作执行不久报错如下.

```
configure: error: error verifying int64_t uses long long
make[2]: *** [configure-stage1-gcc] Error 1
make[2]: Leaving directory `/home/sjarvis/dev/gcc/srcdir'
make[1]: *** [stage1-bubble] Error 2
make[1]: Leaving directory `/home/sjarvis/dev/gcc/srcdir'
make: *** [all] Error 2
```

按照这篇文章[gcc 5.2.0编译错误](http://blog.csdn.net/u012509728/article/details/49923995)中的解决方法, 安装`gcc`和`gcc-c++`后解决.