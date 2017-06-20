# Nginx错误-配置,日志等各方面

## 1. nginx: [emerg] getpwnam("www") failed

[参考文章]

[nginx安装 nginx: \[emerg\] getpwnam(“www”) failed 错误](http://my.oschina.net/u/1036767/blog/210443)

### 问题描述

nginx源码安装完成, 第一次启动出现这个错误

### 原因分析

编译选项中指定了nginx的执行用户为`xxx`, 但是这个用户没有被创建, nginx.conf中的`user`指令也没有修改.

### 解决方法

创建指定用户, 将nginx.conf中的 `user`指令改成此用户.

```shell
useradd -M nginx -s /sbin/nologin
```

## 2. [emerg] invalid host in upstream

### 问题描述

配置完`upstream`块, 启动或是检查时报错如下

```
$ nginx -t
nginx: [emerg] invalid host in upstream "172.16.3.132:8080/" in /usr/local/nginx/conf/nginx.conf:79
nginx: configuration file /usr/local/nginx/conf/nginx.conf test failed
```

### 原因分析

注意: upstream的标准写法是

```
upstream pool名称 {
  server IP:端口 参数;
}
```

其中`IP`前不可以加`http(s)://`前缀, 端口后不可以加`/`和任何后缀.

