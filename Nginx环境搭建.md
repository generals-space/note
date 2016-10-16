# Nginx环境搭建

## nginx源码编译

nginx源码编译依赖可以使用yum安装, 可以省不少时间

```
#!/bin/bash

## 安装依赖
yum install -y gcc gcc-c++ make glibc glibc-common
## nginx必选依赖, 不必在configure时指定
yum install -y pcre pcre-devel
## nginx 开启`--with-http_perl_module`选项
yum install -y perl perl-ExtUtils-Embed
## nginx 开启`--with-http_ssl_module`选项
yum install -y openssl openssl-devel
## nginx 开启`--with-http_image_filter_module`选项
yum install -y gd gd-devel
## nginx 开启`--with-http_image_filter_module`选项
yum install -y GeoIP GeoIP-devel
## 编译选项

./configure \
--prefix=/usr/local/nginx \
--user=nginx --group=nginx \
--with-file-aio --with-ipv6 --with-http_ssl_module \
--with-http_realip_module --with-http_addition_module \
--with-http_image_filter_module --with-http_geoip_module \
--with-http_sub_module --with-http_dav_module \
--with-http_flv_module --with-http_mp4_module \
--with-http_gunzip_module --with-http_gzip_static_module \
--with-http_random_index_module --with-http_secure_link_module \
--with-http_degradation_module --with-http_stub_status_module \
--with-http_perl_module --with-mail --with-mail_ssl_module \
--with-pcre --with-pcre-jit --with-debug

make && make install
```

## 2. 扩展

### 2.1 指定pcre

nginx的必选依赖为pcre(除非禁用nginx的rewirte模块`--without-http_rewrite_module`). 系统中不存在`pcre-devel`包而采用源码安装pcre时, 需使用`--with-pcre=pcre源码目录`. 注意: 是pcre的 **源码目录**， 不需要像apache那样预先编译pcre然后指定pcre的安装路径. 因为nginx在编译的时候会同时将pcre编译进来.

### 2.2 创建nginx用户

```
useradd -s /sbin/nologin nginx
```

在$NGINX/conf/nginx.conf文件中记得把`user`字段的值修改为nginx. 注意这里指定的用户如果不存在nginx启动会出错.

### 2.3 第三方模块

`nginx-dav-ext-module`, github可以找到
依赖 `expat-devel`
编译选项 `--with-http_dav_module --add-module=目标模块路径`
