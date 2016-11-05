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

## 3. "403 Forbidden"错误

文章翻译自: 

[Resolving "403 Forbidden" error](http://nginxlibrary.com/403-forbidden-error/)

### 3.1 引言

`403 Forbidden`错误是Nginx在告诉你, "你请求了一个资源(这个请求我已经收到了), 但我不能给你". 技术上讲, `403 Forbidden`并不是一种错误, 而是HTTP响应的一个状态码. 403响应会在某些情况下**有意**地传回, 比如:

1. 用户不允许访问该页面或资源, 或者整个网站;

2. 用户尝试访问一个目录, 然而`autoindex`被设置为`off`;

3. 用户试图访问一个只能从内部访问的文件;

这些是出现403响应的一些可能的情况, 但这里我们将要讨论的不是服务器有意响应403或者说, 我们并不希望看到403的情况, 这一般是服务器端配置错误导致的.

### 3.2 权限配置不正确

这是发生此类错误较为普遍的原因. 这里的权限, 我并不是单单指被访问文件的权限. 为了向用户提供一个文件响应, Nginx需要拥有对该文件的`read(r)`权限, 并且还需要有该文件各级父目录的`excute(x)`权限. 举个例子, 为了访问这样一个文件:

```
/usr/share/myfiles/image.jpg
```

Nginx需要拥有这个文件的`r`权限, 还需要拥有对`/`, `/usr`, `/usr/share`, `/usr/share/myfiles`这些目录的`x`权限. 如果你设置这些目录的权限为标准的`755`, 该文件的权限为`644(umask: 022)`, 就不会出现这个403错误.

为了检查该路径上各级的属主和权限, 我们可以使用`namei`具, 像这样:

```
$ namei -l /var/www/vhosts/example.com

f: /var/www/vhosts/example.com
drwxr-xr-x root     root     /
drwxr-xr-x root     root     var
drwxr-xr-x www-data www-data www
drwxr-xr-x www-data www-data vhosts
drwxr-xr-x clara    clara    example.com
```

### 3.3 目录的index选项未被正确定义

> PS:译者个人是因为index指令未配置

有些时候, `index`指令并没有包含我们希望的目录中的默认索引(index). 举例来说, 提供PHP程序响应的标准index指令应该设置为:

```
index index.html index.htm index.php;
```

在这个例子中, 当用户**直接访问一个目录时**, Nginx首先尝试响应该目录下的`index.html`, 如果不存在就去找`index.htm`, 然后是`index.php`. 如果都没有找到Nginx就会返回403的响应头. 如果index.php没有在`index`(原文中是`root`指令, 感觉...不太对)指令中定义, Nginx就不去查找`index.php`是否存在而直接返回`403`.

类似的, 如果是Python的服务, 你需要指定`index.py`为目录的默认索引.

These are the most common causes of undesired 403 responses. Feel free to leave a comment if you are still getting 403s.

If you find this article helpful, please consider making a donation.

## 4. 502/504错误

参考文章

[Nginx 502错误原因和解决方法总结](http://www.server110.com/nginx/201312/4409.html)

在普通的`Linux+Nginx+PHP(php-fpm)+MySQL`的架构中, 未根据服务器自身性能进行优化时, 当网站访问量增加, 系统负载变大时, 很容易出现502/504错误. 大多数与Nginx本身无关.

> `Nginx 502 Bad Gateway`的含义是请求的`PHP-CGI`已经执行, 但是由于某种原因(一般是读取资源的问题)没有执行完毕而导致PHP-CGI进程终止.

> `Nginx 504 Gateway Time-out`的含义是所请求的网关没有请求到, 简单来说就是没有请求到可以执行的PHP-CGI.

> 解决这两个问题其实是需要综合思考的, 一般来说`Nginx 502 Bad Gateway`和`php-fpm.conf`的设置有关, 而`Nginx 504 Gateway Time-out`则是与`nginx.conf`的设置有关.

> 而正确的设置需要考虑服务器自身的性能和访客的数量等多重因素.

### 4.1 问题分析

#### 4.1.1 Nginx 502 Bad Gateway

咳, 首先一点, 确定`php-fpm`进程已经启动, 且`proxy_pass/fastcgi_pass`配置正确(这种一般出现在初次配置的时候).

既然是请求已经被执行, 就说明问题存在于**PHP代码执行期间**, 或是**php-fpm将执行结果返回**这两个阶段.

**1. php-fpm执行时间过长**

可能是由于php代码本身有问题(如陷入死循环), 或者数据库读写时间过长(数据库本身也有锁机制), 从而长时间无法得到执行结果.

**2. fastcgi请求, 发送, 读取超时**

可能是系统繁忙, 与FastCGI沟通时间过长

**3. fastcgi响应读取失败或超时**

一般是由于Nginx中的buffer不足, FastCGI的响应被缓冲到磁盘, 减慢了读取速度.

#### 4.1.2 Nginx 504 Gateway Time-out

这个错误的含义是**没有请求到可以使用的CGI**. 那就可能是

**1. php-fpm进程数量不足**

可能是max_children设置值较小, 其余都在"忙"

**2. php-fpm进程占用时间过长**

可能是php代码有bug, 也可能是数据库读取过慢. 总之占用时间过长, 导致该进程迟迟无法生成响应, 也不能接受其他请求.

### 4.2 相关配置

#### 4.2.1 Nginx方面

Nginx处于等待与接收结果的角色, 所以一般502错误在Nginx端解决.

针对产生`502`原因的第1条和第2条, 可以规定PHP-CGI的连接, 发送和读取的时间, 300秒足够用了.

```conf
fastcgi_connect_timeout 300s;
fastcgi_send_timeout 300s;
fastcgi_read_timeout 300s;
```

针对产生`502`原因的第3条, 可以调整nginx缓冲区大小, 相关配置如下

```conf
fastcgi_buffer_size 128k;
fastcgi_buffers 8 128k;
fastcgi_busy_buffers_size 256k;
fastcgi_temp_file_write_size 256k;
```

#### 4.2.2 php-fpm方面

php-fpm处于处理请求/响应结果的角色, 504错误在php-fpm端解决.

对504错误的第1项, 可以调整`max_children`选项, 设置合理的子进程值.

对504错误的第2项, 可以设置如下选项

```conf
request_terminate_timeout 60s ##默认为0s
```

这是为了防止php代码本身存在bug, 使php-fpm进程无法得到释放. 默认为0s是指可以让php-fpm进程一值执行, 没有时间限制

### 4.3. 调整原则

**正确的设置需要考虑服务器自身的性能和访客的数量等多重因素**.

理论上max_children选项值越大越好, 过小会造成排队, 而且php-fpm进程处理的也会很慢

但实际上要根据服务器的内存配置设置, 一个php-fpm进程大概占用20M的内存,  这个值过大的话反而会让服务器因内存不足崩溃或将进程杀死.

缓冲区的设置也是如此, 这个值也要根据响应页面的大小设置, 过小会导致缓冲到硬盘, 读取速度降低; 过大则会造成资源浪费.

## 5. 499错误

情境描述:

前端发起`ajax`请求, nginx将请求转发到后端tomcat服务器, 10s后请求显示失败.

在没有显式设置`timeout`时间的nginx中, 其实这种情况一般不会是nginx引起的. 因为nginx的各种`timeout`时间大多是`60s`, 而`10s`的等待时间的确是太短了, 不应该是nginx主动断开.

499错误是什么？让我们看看NGINX的源码中的定义：

```c
ngx_string(ngx_http_error_495_page), /* 495, https certificate error */
ngx_string(ngx_http_error_496_page), /* 496, https no certificate */
ngx_string(ngx_http_error_497_page), /* 497, http to https */
ngx_string(ngx_http_error_404_page), /* 498, canceled */
ngx_null_string,                    /* 499, client has closed connection */
```

可以看到，499对应的是`client has closed connection`。这很有可能是因为服务器端处理的时间过长，客户端"不耐烦"了。

也有可能是前端对`ajax`请求设置了超时时间, 超时后断开连接并触发`abort`方法.