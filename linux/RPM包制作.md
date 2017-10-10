# RPM包制作

参考文章

1. [ 一堂课玩转rpm包的制作](http://blog.chinaunix.net/uid-23069658-id-3944462.html)

2. [一步步制作RPM包](http://laoguang.blog.51cto.com/6013350/1103628)

## 1. 声明

1. rpm包不解决依赖, 安装之前需要手动安装依赖包

2. 制作rpm也需要事先解决依赖

3. 安装rpm实际上是把rpm包中预先编译好的文件直接拷贝到指定目录(系统约定, 无需手动编写), 安装者需要事先解决其中可执行文件动态链接库的依赖

4. 所以rpm包的制作过程实际上是一次源码编译过程

安装打包工具

```
yum install rpmdevtools
```

```
rpmdev-setuptree
[root@localhost rpmbuild]# cd /root/rpmbuild/
[root@localhost rpmbuild]# tree
.
├── BUILD
├── RPMS
├── SOURCES
├── SPECS
└── SRPMS

5 directories, 0 files

```

```
$ rpmdev-newspec -o salt-20161105.spec
```

会在当前目录下创建`salt-20161105.spec`文件, 所以要事先进入到`SPECS`目录下.