# Python源码安装[ubuntu] arm

参考文章

1. [python3.7移植到ARM开发板](https://blog.csdn.net/u012230668/article/details/89206857)

python: 3.7.8
ubuntu: docker 版本 20.04

本来我的常用系统是 CentOS, 但是 CentOS 上的交叉编译文档太少了, 大多数人用的都是 Ubuntu.

centos 的 arm 编译器是`arm-linux-gnu-gcc`和`arm-linux-gnu-g++`. 

```
./configure --prefix=/opt/python3.7 \
CC=arm-linux-gnu-gcc \
CXX=CC=arm-linux-gnu-g++ \
--build=x86_64-linux-gnu \
--host=aarch64-linux-gnu \
--target=aarch64-linux-gnu
```

但是`configure`失败了, 没找到解决办法, 换用 Ubuntu 试试.

不过毕竟是两个系统, 就算是x86的平台, CentOS 编译出来的 Python, Ubuntu 也不能用, 找不到共享库, 更别说 arm 平台了. 所以在 Ubuntu 交叉编译的结果最终应该也只能在 Ubuntu 下运行, 很心碎.

```
apt-get install gcc-arm-linux-gnueabi
```

```
./configure --prefix=/opt/python3.7 CC=arm-linux-gnueabi-gcc --build=x86_64-linux-gnu --host=arm-linux-gnueabi --target=arm-linux-gnueabi --disable-ipv6
```

不同平台下的`--host`和`--target`值很难选, 貌似是跟编译器绑定的, 格式就像`CPU平台(arm或aarch64)-编译器名称(linux-gun或linux-gnueabi)`一样.

> `--disable-ipv6`是必须要加的, 否则`configure`会出错.

上面的命令还是会出错.

```
configure: error: set ac_cv_file__dev_ptmx to yes/no in your CONFIG_SITE file when cross compiling
```

要按参考文章1中所说, 再加两个参数.

./configure --prefix=/opt/python3.7 CC=arm-linux-gnueabi-gcc --build=x86_64-linux-gnu --host=arm-linux-gnueabi --target=arm-linux-gnueabi --disable-ipv6 ac_cv_file__dev_ptmx=yes ac_cv_file__dev_ptc=yes

这次出错的更快了...

```
 ./configure --prefix=/opt/python3.7 CC=arm-linux-gnueabi-gcc --build=x86_64-linux-gnu --host=arm-linux-gnueabi --target=arm-linux-gnueabi --disable-ipv6 ac_cv_file__dev_ptmx=yes ac_cv_file__dev_ptc=yes
checking build system type... x86_64-pc-linux-gnu
checking host system type... arm-unknown-linux-gnueabi
checking for python3.7... no
checking for python3... no
checking for python... python
checking for python interpreter for cross build... configure: error: python3.7 interpreter not found
```

最开始非常不理解原因, 参考文章1提到了这种情况, 但是没有说过解决方法. 后来才发现, 参考文章1一开始先编译了x86平台的 Python 然后`make install`了, 这样在`configure` arm 平台时才不会出错, 于是我先按照平常的方法先编译了一遍 x86 的, 最后`ln -s /usr/local/python3.7/bin/python3 /usr/local/bin/python3`就可以了.

要注意x86平台`configure`时的`--prefix`与arm平台的不要相同, 免得覆盖.

