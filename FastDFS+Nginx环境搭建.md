# FastDFS环境搭建

## 写在前面

FastDFS是一个开源的, 高性能的的分布式文件系统, 主要的功能包括: 文件存储, 同步和访问, 设计基于高可用和负载均衡. FastDFS非常适用于基于文件服务的站点, 例如图片分享和视频分享网站.

FastDFS有两个角色: 跟踪服务(tracker)和存储服务(storage).

- 跟踪服务: 控制调度文件以负载均衡的方式访问.

- 存储服务: 文件存储, 文件同步, 提供文件访问接口, 同时以key value的方式管理文件的元数据.

跟踪和存储服务都可以由1台或者多台服务器组成, 同时可以动态的添加, 删除跟踪和存储服务而不会对在线的服务产生影响, 在集群中, tracker服务是对等的.

fastdfs原来是`sourceforge`上的项目, 不过已经停止更新(当前最新版为1.27), 现在迁移到了github(目前最新版为5.05), [项目地址](https://github.com/happyfish100/fastdfs).

fastdfs源码编译依赖`libevent`库, 这个可以使用yum安装; 另外它还依赖`libfastcommon`, yum源中应该没有需要源码安装, 也是在github上可以找到其源码, [项目地址](https://github.com/happyfish100/libfastcommon).

nginx需要对应的fastdfs扩展, 名为`fastdfs-nginx-module`. 同样, 其在`sourceforge`上的版本已经过于陈旧, 不推荐使用. github上的[项目地址](https://github.com/happyfish100/fastdfs-nginx-module).

## 1. 环境要求

- 系统版本: CentOS6

- fastdfs: 5.05

- nginx: 1.10.1



## 2. 安装步骤

> 注意: 这一部分是`tracker`与`storage`服务器都需要安装的.

说明一点, 这个版本不需要安装`libevent`与`libevent-devel`库(之前在`sourceforge`上的`1.27`版是需要的). 所以首先是`libfastcommon`, 在其根目录下直接执行如下命令即可, 可以参见`INSTALL`文件.

```shell
cd libfastcommon
./make.sh && ./make.sh install
```

接下来是`fastdfs`本身, 这一版本的`make.sh`不需要作任何修改(网上有之前版本需要解开`HTTPD`与`LINUX SERVICE`的注释), 同上面的`libfastcommon`安装方式一样, 简单粗暴.

```shell
cd fastdfs目录
./make.sh && ./make.sh install
```

## 3. 配置方法

安装完成后fdfs会交由`service`管理, 所以可以以服务形式启动与关闭. 另外, fdfs系列的命令也已经添加到`/usr/local/bin`目录下, 可以直接执行. 且其配置文件会放在`/etc/fdfs`目录下

### 3.1 tracker

```
$ cd /etc/fdfs
## 这个版本的配置文件默认都以.sample结尾
$ ls
client.conf.sample  storage.conf.sample  tracker.conf.sample
$ mv ./tracker.conf.sample ./tracker.conf
$ vim ./tracker.conf

## 这个路径可以自定义, 但必须是已经存在的, fdfs不能自动创建
base_path=/opt/fastdfs/tracker1
## 其他的如`port`与`http.server_port`最好先保持默认, 等熟悉之后再根据需要修改
```

## FAQ

### 1.
