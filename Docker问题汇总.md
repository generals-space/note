# Docker问题汇总

## 1. Docker容器bash中输入中文

**参考文章**

[linux终端不能输入中文解决方法 ](http://blog.sina.com.cn/s/blog_5c4dd3330100cpmm.html)

[在Docker容器bash中输入中文](http://blog.shiqichan.com/Input-Chinese-character-in-docker-bash/)

------

docker容器内的bash无论如何都无法输入中文, 不管是在启动容器时打开bash, 还是以服务形式启动容器后再通过 `nsenter`工具进入容器之后显示的bash. 不管是在什么情况下输入甚至粘贴, 不是出现乱码, 回车无反应甚至根本无法上屏. 而且输出时中文全都是乱码.

尝试在容器内 `/root`家目录新建 `.inputrc`文件, 添加以下内容

```shell
set meta-flag on
set convert-meta off
set input-meta on
set output-meta on
```

重启容器发现可以在bash命令行上输入中文, 但是回车发现与预期结果不同, 而且输出时中文依然是乱码. 尝试设置`locale`, 不管将环境变量LANG设置为 `LANG=en_US.UTF-8`还是 `LANG=zh_CN.UTF-8`都不起作用.

------

真正的解决方法是, 在启动容器时传入 `env`参数

```shell
docker run -i -t ubuntu env LANG=C.UTF-8 /bin/bash
```

或是在Dockerfile文件中写入如下行

```shell
ENV LANG=C.UTF-8
```

## 2. CentOS:7的docker镜像使用systemctl

### 问题描述

使用`docker.io/centos:7`的docker官方镜像, 配置`php-fpm`的服务脚本完成. 但是使用`systemctl`命令操作`php-fpm`时出现如下问题. (服务脚本的建立参考[这里](http://www.centoscn.com/CentOS/config/2015/0507/5374.html))

```shell
[root@e13c3d3802d0 /]# systemctl start php-fpm
Failed to get D-Bus connection: Operation not permitted
```

### 原因分析及解决方法

参考下面两篇文章的解释.

[docker Failed to get D-Bus connection 报错](http://welcomeweb.blog.51cto.com/10487763/1735251)

[如何在Docker CentOS容器中使用Systemd](http://www.codesec.net/view/434721.html)

这是CentOS:7镜像中的一个bug, 目前无法修复, 只能等到7.2版镜像. 这个bug的原因是因为`dbus-daemon`没能启动. 其实`systemctl`并不是不可以使用. 将你dockerfile的`CMD`或者`ENTRYPOINT`设置为`/usr/sbin/init`即可, 容器会在运行时将`dbus`等服务启动起来. 然后再执行`systemctl`命令即可运行正常.

```shell
docker run --privileged  -e "container=docker"  -v /sys/fs/cgroup:/sys/fs/cgroup -d docker.io/centos:7  /usr/sbin/init
```

> 注意: 上面的命令要完全执行...少一个都会被坑的很惨.

`--privileged`: systemd 依赖于`CAP_SYS_ADMIN capability`. 意味着运行Docker容器需要获得 `privileged`(这不利于一个base image);

`-v`选项: systemd 依赖于访问`cgroups filesystem`; systemd 有很多并不重要的文件存放在一个docker容器中, 如果不删除它们会产生一些错误; 原来以为`-v`选项只是一个共享目录而已, 然而我擅自去掉这个选项后, `systemctl start php-fpm`执行时卡住了, 而且并没有执行成功. 加上这个选项才可以.

另外, 最好加上`-d`选项, 这个是我自己加上的. 交互时启动时竟然提示输入容器密码...而且尝试了很多遍都不对(也不是宿主机的密码)...这真是个奇妙的经历.

```shell
...
CentOS Linux 7 (Core)
Kernel 4.5.5-300.fc24.x86_64 on an x86_64

325606d855b9 login: root
Password:
Login incorrect

325606d855b9 login:
Password:
Login incorrect
...
```

## 3. docker容器无法作为服务启动

有些容器如果不使用`-it`选项并搭配执行`/bin/bash`命令无法以服务形式保持启动状态, 都是立刻结束, 使用`docker logs 容器ID`也没有错误日志.

因为容器是否长久运行, 与`docker run`指定的命令有关, 与`-d`参数无关.

### 解决办法

使用`-it`选项并执行`/bin/bash`命令, 进入容器shell. 然后退出, 此时容器将会停止. 使用`docker ps -a`查看刚才运行的容器ID, 再使用`docker start 容器ID`将会使其进入服务状态, `docker ps`可以看到它依然在运行, 而且命令还是`/bin/bash`. 然后就可以通过`nsenter`等工具进入容器了.

## 4. Ubuntu14.04安装nsenter

ubuntu14的util-linux版本为2.20, 但想要进入docker容器, 不能低于2.24. 需要手动编译安装. 安装命令如下, 注意要首先安装依赖包.

```shell
sudo apt-get install autopoint autoconf libtool automake
wget https://www.kernel.org/pub/linux/utils/util-linux/v2.24/util-linux-2.24.tar.gz
tar xzvf util-linux-2.24.tar.gz
cd util-linux-2.24
./configure --without-ncurses
make && make install
```

## 4. mysql容器数据丢失

参考文章

[mysql的docker镜像中如何创建数据库](http://dockone.io/question/887)

问题描述:

使用mysql官方docker镜像, 新建一个容器A后在其中创建数据库并存储数据. 然后使用`docker commit`将容器A保存为一个镜像B. 以镜像B为基础启动容器C, 但是C中没有保存之前在A中创建的数据.

原因分析:

mysql-server的[Dockerfile](https://hub.docker.com/r/mysql/mysql-server/~/dockerfile/)有这样的一行

```
VOLUME /var/lib/mysql
```

意味着使用这个镜像时, 容器的`/var/lib/mysql`目录会被映射到宿主机上docker工作目录(默认为`/var/lib/docker`)下的某个目录. 具体映射到哪个目录, 可以通过 `docker inspect containerID` 查看.

```
$ docker inspect 原mysql容器ID | grep -i volume
"VolumeDriver": "",
 "VolumesFrom": null,
     "Source": "/var/lib/docker/volumes/023bb9aa4c35bd12625f89768fbfd86b73f0c8286fbc2d504e921872886b0e70/_data",
 "Volumes": {
$ cd /var/lib/docker/volumes/023bb9aa4c35bd12625f89768fbfd86b73f0c8286fbc2d504e921872886b0e70/
$ ls
_data
```

如果不做更改, 那么也就意味着你写入的数据会被直接写入到宿主机的该目录中, 而且不随容器的销毁而销毁。同样, commit的时候, 该目录的内容也不会被加入到镜像中。所以使用commit出来的镜像, 你就无法看到先前的数据了, 因为commit命令不会将挂载卷中的数据commit到镜像中.

解决方法:

将原容器的`_data`目录复制到目标容器的挂载卷下, 重启容器即可.

建议读一下[mysql-server](https://hub.docker.com/r/mysql/mysql-server/)中`Where to Store Data`一节，会对你了解如何存储mysql数据有所帮助.
