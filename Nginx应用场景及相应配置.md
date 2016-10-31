# nginx应用场景及相应配置

## 1 泛解析情况下非法子域名返回404

### 1.1 情景描述

当前有顶级域名`example.com`, `www`与`bbs`为合法子域名, 可分别访问主页和论坛, `nginx`配置如下. 现要求任何非法子域名(如`jkl.example.com`等**未由`nginx`的`server`块显示配置的子域名**)由`nginx`返回404.

咳, 这个要求看起来有点搞笑, 禁止非法子域名访问只要不开启域名泛解析就行了, 何必要由nginx返回404呢.

```
http {
    ...
    server {
        listen 80;
        server_name www.example.com;
        ...
    }
    server {
        listen 80 default_server;
        server_name bbs.example.com;
        ...
    }
}
```

### 1.2 解决方法

#### 1.2.1 打开泛解析

首先在域名服务器打开域名泛解析, 即`*.example.com`要全部指向目标服务器IP, 在`nginx`未作任何其他配置时, 访问其他非法子域名如`jkl.example.com`, 会指向`nginx`配置文件中带有**`default_server`**标记的`server`块所声明的域名. 即上面的`bbs.example.com`.

另外, 如果整篇配置文件中都没有`default_server`标记, `nginx`将会返回配置文件中出现的第1个的`server`块(**不管是不是相同的顶级域名**). 比如移除上面配置中`bbs`子域名的`default_server`标记, 再次访问`jkl.example.com`将会返回`www.example.com`.

#### 1.2.2 nginx配置

配置`nginx`, 让所有其他非法域名都返回404. 有三种方法, 就是上面说到的`default_server`标记和首部插入`server`块, 还有就是`server_name`域名通配.

**`default_server`标记**

在配置文件`http`块中任何位置添加如下内容. 因为在`nginx`中`default_server`标识唯一, 所以使用这种方法时需要将其他`default_server`标记删除, 这对域名指向的**正确性**没有影响.

```
server {
    listen 80 default_server;
    server_name _;
    return 404;
}
```

其实, 有了`default_server`标记, `server_name`字段实际上也不需要了. 但是在一些情况下, `server_name`其实还是有用的, 比如多个顶级域名情况下, 需要只对其中某一个顶级域名的非法子域名返回404时.

**插入`server`块**

在`http`块中添加如下内容(相比以上只是少了`default_server`标记), 要确保`nginx`配置文件中没有其他`default_server`块并且一定要是`http`中第1个`server`块. 当然, 这种方法也不需要`server_name`字段. 不过如果是多个顶级域名时, 可能会出现访问`jkl.example2.com`而得到`www.example.com`的情况. 第1个`server`块嘛, 都跨域名匹配了...

```
server {
    listen 80;
    server_name _;
    return 404;
}
```

**`server_name`域名通配**

这个感觉比较好懂, 就是将`server_name`设置为`*.example.com`然后`return 404`就可以了. 这也算是显示的指明了子域名了.

------

这3种方式都是将由DNS服务器指来的访问域名地址, 与`nginx`配置文件中有**显示声明**的`server`块进行匹配, 所有其他未显式在`server`块中声明的子域名(如果有多个顶级域名, 那就是全部顶级域名未显式声明的非法子域名), 都会被这3种方式捕捉到.

不过直觉上来说, 感觉第2种不太好啊.

### 1.3 扩展

再考虑一种情况, 所有`example.com`子域名都监听`80`端口, 另外还有`example2.com`监听`8080`, 它们两个都分别有`www`与`bbs`子域名. 怎样设置`nginx`对这两个顶级域名的非法子域名都返回404?

上面说过, `default_server`标记在配置文件中是唯一的, 但我尝试了在`www.example.com`与`www.example2.com`所在的`server`块中`listen`字段都加了这个标记, 没有发生错误, 同时生效, 可以说**default标记对同一端口才是唯一的**.

`nginx`会根据端口值准备好为请求服务的`server`块, 不是对应端口的绝不处理. 比如当前这个例子, 设置`nginx`泛解析方面配置如下. 整个配置文件中只有这一个`default_server`标记, 没有对80端口做泛解析. 你说`jlk.example.com`能不能被下面这个**default_server**捕捉到? 这个请求是80端口的哦.

```
server {
    listen 8080 default_server;
    server_name _;
    ...
}
```

答案是不会, 谁让请求不能跨端口呢? `example.com`由于dns开了泛解析, 在没有80端口的`default_server`情况下, 只能寻找配置文件中的第一个监听80端口的`server`块了.

我都不知道这样是好还是不好了, 看起来nginx处理请求是很有原则的, 值得肯定.

但如果我不只有两个顶级域名呢? 如果我有20个顶级域名(分别监听不同端口), 它们都需要对非法子域名返回404呢? (虽然这个问题真的很扯--开泛解析是为了什么!?) 我是需要为每个正在监听的端口都再添加一个带`default_server`标记的`server`块呢, 还是在`http`块开头部分为它们添加第1个匹配的`server`块呢, 还是为它们分别添加顶级域名通配呢? 这3个不管哪种, 都需要添加20个`server`块呀...

>PS: **nginx的域名匹配流程**--**`server_name`的下划线**

>它有什么作用呢? 它可以作为未显式在配置文件中声明的域名匹配, 但匹配流程次序的干扰太多了.

>首先它不能跨端口匹配域名请求, 所以下面都假设是相同端口. 在设置其他`server`块为`default_server`时, 非法子域名也只会匹配到`default_server`所在的`server`块, 然后下面假设不存在`default_server`标记. 就算不存在`defalut_server`, 它的优先级也不如`*.example.com`高, 因为后者相当于显示设置了域名匹配. 但就算没有`*.example.com`这种形式, 非法子域名匹配的也不是下划线`server`块, 而是`http`中第1个`server`块. 这样看来, 它的使用情况实在有限啊.

## 2. 使用nginx搭建http代理服务器

这里要求的代理服务器不是简单的访问`www.abc.com?key=value`由nginx实际取回`www.xyz.com?key=value`这样类似于重定向的功能, 而是那种可以在chrome中的http代理服务器配置中填写nginx所在服务器IP及nginx监听端口, 每次访问网络时由代理服务器代为获取的形式.

```shell
server {
    resolver 8.8.8.8;
    resolver_timeout 5s;
    listen       0.0.0.0:80;

    access_log  /var/log/nginx/access_proxy.log  main;
    error_log  /var/log/nginx/error_proxy.log;

    location / {
            proxy_pass $scheme://$host$request_uri;
            proxy_set_header Host $http_host;

            proxy_buffers 256 4k;
            proxy_max_temp_file_size 0;
            proxy_connect_timeout 30;

            proxy_cache_valid 200 302 10m;
            proxy_cache_valid 301 1h;
            proxy_cache_valid any 1m;
    }
}
```

关于`location`块中的`proxy_*`不再多介绍, `proxy_set_header`可以决定是否将客户端IP对目标服务器隐藏, 也许这就是高匿代理的秘密所在. 开头的`resolver`字段是必须的, 因为用户在设置代理服务器之后访问网络不会再自行进行DNS解析, 而是将目标网址直接发送给代理服务器, 此时需要代理服务器自己去解析, 所以只能设置DNS地址.

另外, nginx貌似不能作为`https`代理, 因为nginx不支持CONNECT，所以如果访问Https网站, 比如: `https://www.google.com`, nginx的access.log 日志如下:

```shell
"CONNECT www.google.com:443 HTTP/1.1" 400
```

## 3. Nginx作反向代理处理后端服务器的301/302问题

### 3.1 场景描述

Nginx作反向代理服务器, 用户`U`通过nginx服务器`N`访问后端服务器`S`, 由`N`取得`S`的数据并将结果返回给用户`U`. 如果后端服务器返回`301/302`, 那N也将这`301/302`直接返回给用户.

现在的要求是当后端服务器`S`返回`301/302`时, 由`N`再向`S`发送一次请求, 直到取得的结果不是`301/302`, 才将结果返回给用户.

首先要说明的是, `301/302`不是出错, 但如果后端服务器`S`重定向的地址是`U`无法直接访问时(这种情况应该比较普便, 可以将`N`看作是前端防火墙), 用户`U`就会得到错误. 另外可能由于重定向, 重写后的`url`无法正确找到, 也会得到错误.

### 3.2 场景再现

有时由于后端服务器的`301/302`响应, 用户无法得到正确的结果, 举个例子来说明.

- 反向代理服务器`N`的IP为`192.168.1.100`;

- 后端服务器`S`的IP为`192.168.1.200`;

`N`的配置如下

```
server {
    listen       9100;

    location /some-prefix/ {
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        ##  proxy_set_header Via "nginx";
            proxy_pass http://192.168.1.200:9200;
    }

}

```

`S`也用nginx作服务器(这个选择可以多样, 因为当前要解决问题的重点在于前端服务器), 其配置如下

```
server {
    listen      9200;
    location /some-prefix/ {
        rewrite ^/some-prefix/(.*) /some-prefix/some-suffix/$1 rewrite;
    }
    location /some-prefix/some-suffix/ {
        root /usr/share/html;
        index index.html index.php;
    }
}
```

对于这种应用场景的解释, `S`的`/usr/share/html`为网站根目录, 原来Web程序在网站根目录的`/some-prefix`目录下存放, 后来迁移到`/some-prefix/some-suffix`目录下, 于是使用`rewrite`进行重定向. 出于某种原因(比如后端服务器不是nginx, 是由Web程序自定义的重定向时), 无法让后端服务器在`rewrite`之后再次匹配本机的`location`字段的其他配置, 所以只能将`rewrite`之后的url返回并加上`302`码.

使用`N`作反向代理, 用户访问`192.168.1.100:9100/some-prefix`, 结果怎样呢? 浏览器地址栏显示`192.168.1.100:9200/some-prefix/some-suffix/`, 这是`N`的IP加上`S`的端口的组合, 很是奇异. 用户浏览器访问这个地址自然无法得到正确结果.

分析一下这种情况的原因, `rewrite`命令处理相对路径的url(以`/`为根, 不涉及`http://IP`或`http://域名`)时, `http://IP`或是`http://域名`与来源url保持一致, 所以返回给`N`与`U`的`host`地址是不变的, 依然是`192.168.1.100`. 不过端口貌似就是自己被访问到的端口了, 而这个组合返回给`N`是默认不会被处理直接转发给用户`U`的.

### 3.3 解决方案

解决方案是使用`proxy_*`系中的`proxy_redirect`指令, 它可以将`N`返回给用户的url进行改写...嗯, 好像与预想中的解决方案不同. 原来想使用反向代理`N`对后端`S`进行二次访问再返回给用户的方法没找到, 只能退而求其次了.

话说回来, `proxy_redirect`指令可以修改返回给用户的`Location`与`Refresh`字段(注意, 由nginx向后端服务器发送请求时, nginx本身也可看作是 **用户**). `301/302`重定向时都会带一个地址, 引导用户...的浏览器再次访问, 这个地址就存储在http协议的`Location`字段里.

它的使用方法与`rewrite`类似, 不过还不清楚是否支持正则, 应该是不支持的. `proxy_redirect Location字段中的一部分 想要替换成的部分`. 比如上面的出错示例中, 用户浏览器得到的`Location`为`http://192.168.1.100:9200/some-prefix/some-suffix/`, 我们只需要把`9200`替换成`9100`就行了, 这样就可以写成`proxy_redirect 9200 9100`. 咳, 开个玩笑, 事实上需要写为 **`proxy_redirect http://192.168.1.100:9200/ http://192.168.1.100:9100/`**. 这样, 用户得到的`Location`字段就会变成`http://192.168.1.100:9200/some-prefix/some-suffix/`了.

`N`的配置变为如下(就多了一行)

```
server {
    listen       9100;

    location /some-prefix/ {
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        ##  proxy_set_header Via "nginx";
            proxy_redirect http://192.168.1.100:9200/ http://192.168.1.100:9100/;
            proxy_pass http://192.168.1.200:9200;
    }

}
```

实际情况中可能需求多种多样, 也许上面的方面并不适用, 但是只要知道用户得到的`Location`字段, 根据实际要访问的url地址进行比对, 重写一下还是不难的.

不过这样的场景真是...让人无语

## 4. Nginx端口相关

## 5. Nginx不缓存

在开发调试web的时候, 经常因为浏览器缓存(cache)去清空缓存或者强制刷新...现在的问题是, 微信嵌入网页开发, 没找到禁止缓存的选项, 只能让服务器端配合了.

解决办法: Nginx配置文件里, url对应的location中加入`Cache-Control`字段

```shell
    location / {
        add_header Cache-Control no-store; ##注意add_header中间是下划线
        ...
    }
```

另外`add_header`这个指令还可以加自定义的响应头字段. 见下一小节.

## 6. Nginx自定义响应头

[参考文章](http://sumsung753.blog.163.com/blog/static/14636450120133794814985/)

`add_header`可以添加自定义响应头字段, 可以在浏览器响应头里面看到.

`add_header`指令的使用方式为`add_header key value`.

- `key`可以自定义, 不加引号, 作为键
- `value`可以为字符串, Nginx内置变量, 也可以是自定义变量

```shell
    location / {
        add_header hello 'world';               ##字符串
        add_header request_uri $request_uri;    ##Nginx内部变量
        set $abc 'cool';                        ##创建自定义变量, 以'$'起始, 类似于php
        add_header abc $abc;                    ##输出自定义变量
        ...
    }
```

我把这种作为一种**日志查看方式**, 尤其是对请求的location正则匹配与rewrite操作的结果, 可以不必去翻看日志转而将需要的信息以键值对的形式返回到浏览器响应头. 使用curl的`-I`选项访问目标路径打印响应头信息十分方便.

```shell
#curl -I localhost
...
request_uri: /
hello: world
abc: cool
...
```

**注意: 当为一个add_header指令设置的响应头的值为空时, 将没有办法看到结果.** 如果遇到设置了响应头为某变量, 但是访问结果中并没有输出, 可能是应为此变量的值为空了.

## 7. Nginx自定义状态码

来一个有意思的东西, 关于Nginx的`return`指令. 玩玩而已, 没有什么特别的功能.

我们可以手动指定某`location`或`server`字段等的`4xx`状态码, 并指定根据此状态码的值返回的页面(`errorpage`的作用); 或者对于`3xx`是通过`rewrite`指定代为完成的, 其中`permanent`与`redirect`分别代表`301`与`302`.

当我们手动返回一个状态码时, 比如

```shell
    server {
        ...
        ##errorpage 404 /404.html;
        location / {
            ...
            return 333;
        }
    }
```

这样, 在浏览器的响应头中会出现`333`代码, 一般接在`HTTP/1.1`后面, 像`HTTP/1.1 333`这样. 不过这样可能无法得到页面的正常显示, 想看到响应结果, 只能`curl -I 目标URL`.

不过, `return`的返回值不可以是字符串, 像`return 'hehe'`. 不能像`add_header`那么方便啊.

## 8. nginx与php-fpm分离

### 8.1 场景分析

非生产环境下, nginx, php-fpm及php工程代码都是放在同一台主机上. 其工作方式一般为: nginx接收客户端请求, 如果是静态文件, 则由其本身直接返回给客户端; 如果是php脚本文件, 则将其通过`fastcgi_*`系列指令交由`php-fpm`(而且一般是9000端口)进程处理, 然后将其处理的结果返回. 这种形式最容易实现.

现在有如下场景: 有A, B, C, D共4台主机, X, Y两个php工程. 其中A, B共同作为前端nginx转发(可以认为前面由F5或另一台Nginx将请求转发至A和B), 根据 `location`前缀判断客户端访问的是X还是Y工程, 如`/X/index.php`即转向C或D的X目录并通过`proxy_*`系列命令与C和D进行数据传输; C与D上各自部署X, Y两个工程, 并按照上面的方式运行了nginx + php-fpm; 另外,X的数据库在C机器上, Y的数据库在D机器上. 总体架构如下图.

可以看到, C和D上面运行的程序过多, 而且相互依赖, 耦合严重, 并且nginx配置修改时很可能要涉及A, B, C, D四台机器, 配置烦琐, 尤其业务较重时横向扩展相当复杂. 现在考虑, 将运行单元分离, nginx, php工程与数据库分别放在单独的服务器上, 各司其职. 基本架构如下.

------

### 8.2 工程路径

目前一个问题是, **php工程代码应该放在哪里?**

有两种可能的猜测: 1. 工程目录放在php-fpm主机中, nginx 接收前端请求, 将`root`指令设置为工程目录在fpm主机上的路径, 而本身不保留工程代码; 2. 工程目录放置在nginx主机中, nginx接收前端请求, 在本地寻找`root`指令所指相应的本地目录中的php脚本, 将其作为数据流传到php-fpm主机相应端口, 由php-fpm解析并执行. 仔细想一下, 其实第2种情况的可能性不大, 毕竟即使nginx与php-fpm主机同属内网环境, 传输php脚本数据流的代价还是太大了.

实际测试一下. 测试环境为:

docker CentOS7镜像下源码安装php与nginx, 和官方mysql镜像(如果依赖于`systemctl`的还是暂时放弃吧, CentOS目前开启`systemctl`相当麻烦, 要等到7.2时才能修复).

nginx 配置大体如下:

```shell

http {
...
  log_format main '$remote_addr - $remote_user "$request" '
  '$status $body_bytes_sent '
  '"$http_user_agent" "$http_x_forwarded_for" $document_root $fastcgi_script_name';

  access_log logs/access.log main;
..
  server {
    listen 80;
    server_name localhost;
    location / {
      root /var/www/html;
      fastcgi_pass php-fpm主机IP:9000;
      fastcgi_index index.php;
      fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
      include fastcgi_params;
    }
  }
}
```

注意`server`块下的`root`字段值, 在php-fpm所在主机的`/var/www/html`目录下建立一个php文件, 这里命名为`here.php`.

```php
<?php
  echo "i am here";
?>
```

使用curl命令(容器中一般不带有这个命令, 需要自行安装)访问nginx所在主机IP`curl NginxIP/hehe.php`, 可以得到`i am here`的输出. 说明nginx传给`php-fpm`进程的, 不是php文件流, 而是脚本所在路径, `php-fpm`进程会根据这个路径自行寻找并解析.

------

### 8.3 动静分离

然而这个问题解决不意味着一切都结束了, 在这种情况下, **nginx的动静分离怎么实现?**

先问问php-fpm是怎么看的...在php-fpm所在主机(容器)的`/var/www/html`目录下放一张图片`img.jpg`, 然后访问`curl NginxIP/img.jpg`. 命令行下得到

```shell
Access denied
```

查看nginx的错误日志, 有如下输出

```
2016/07/03 08:56:47 [error] 36#0: *1 FastCGI sent in stderr: "Access to the script '/var/www/html/img.jpg' has been denied (see security.limit_extensions)" while reading response header from upstream, client: 127.0.0.1, server: localhost, request: "GET /img.jpg HTTP/1.1", upstream: "fastcgi://172.17.0.2:9000", host: "localhost"
```

...so, php-fpm是没法自己处理静态文件的, 还是要由nginx自己来. 那么问题来了, nginx与php-fpm所在主机都要保留一份工程代码, 而且路径还必需相同以便于同时正常响应静态文件请求与动态请求...那么这种nginx与php-fpm的分离还有必要吗?

另外, 就算nginx按照这样的方式实现了动态分离, 如果用户有 **上传文件** 的需要, 怎么办? 上传的文件是在nginx这边还是php-fpm那里? 猜测是php-fpm那里, 因为上传文件流是需要由php代码捕获并存储的.

以`wordpress`为例, 在nginx与php-fpm所在主机的`/var/www/html`部署wordpress工程, nginx自行处理静态文件请求. nginx大致配置如下

```shell
server {
  listen 80;
  server_name localhost;
  root /var/www/html;
  index index.php;
  location / {
    try_files $uri $uri/ /index.php?$args;
  }
  location ~ \.php$ {
    fastcgi_pass php-fpm主机IP:9000;
    fastcgi_index index.php;
    fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    fastcgi_param QUERY_STRING $query_string;
    include fastcgi_params;
  }
  location ~ .*\.(gif|jpg|jpeg|png|bmp|swf)$ {
    expires 30d;
  }
}
```

安装配置成功, 然后在wordpress管理后台写一篇文章, 并在文章中插入一张图片并发布. 其实不用等到发布了, 上传图片时后台会提示uploads目录权限不足, 于是在php-fpm所在主机的wordpress/wp-content目录下建立uploads目录并设置其权限为777, 然后上传成功...但是没有办法看到, 图片预览失败. 文章发布后, 图片依然是破碎的状态.

在php-fpm主机的`wordpress/wp-content/uploads/2016/07`目录下可以找到上传的图片, 然而nginx所在主机的相应目录下并没有这个图片, 这就是看不到上传的图片的原因-静态文件请求不会到达后端的程序服务器.

这种情况下, 解决办法可以是使用`rsync`进行数据同步, 包括更新的代码文件, 用户上传的静态文件等.

但还是那个问题, 这样做nginx与php-fpm的分离到底是不是显得有点得不偿失了?

------

### 8.4 扩展

以下这些情况下, 进行业务分离会比较方便.

- 后台程序(不只是php)比如`django`,`dotcms`等, 会将静态文件路径隐藏, 所有静态请求都需要由程序本身去寻找并响应;

- 完全没有用户上传行为的需要;

- php工程目录其实是以nfs形式挂载在共享网络卷上;

如果是在小型业务下, 额外搭建同步服务器还是很费时间的.

## 9. Nginx防盗链实现

参考文章

[nginx实现图片防盗链(referer指令)](http://www.ttlsa.com/nginx/nginx-referer/)

首先看一个例子

```
location ~* \.(jpg|png|gif){
  valid_referers none blocked *.test.com;
    if ($invalid_referer) {
        return 403;
    }
}
```

`valid_referers`指令指定了合法的`Referer`字段信息, 除了`*.test.com`是一个域名示例外, `none`表示请求头中`Referer`字段为空的情况, 例如从浏览器种直接访问该资源; `blocked`表示请求头`Referer`字段不为空, 但是里面的值被代理或者防火墙删除了，这些值都不以`http://`或者`https://`开头.

如果来源请求头部的`Referer`字段信息不在此列表中, 则将`$invalid_referer`变量的值置为1(默认貌似为空), 之后使用`if`语句定义这种情况下的处理方式.

`valid_referers`指令来自nginx的`ngx_http_referer_module`模块, 通常用于阻挡来源非法的域名请求. 但是, 伪装Referer头部是非常简单的事情，所以这个模块只能用于阻止大部分非法请求. 若有特殊要求可以使用第三方模块`ngx_http_accesskey_module`来实现公用key的防盗链，迅雷都可以防的哦亲...

## 10. 隐藏响应头中的nginx版本信息

由Nginx服务的请求在其响应的响应头中会有一个`Server`字段, 其值类似于`nginx/1.10.1`, 如果不希望客户端看到Nginx的版本信息, 可以在nginx配置文件里设置`server_tokens off;` 这样响应头的`Server`字段就只有`nginx`而没有版本号了.

## 11. Nginx访问控制/限制IP

使用nginx的`deny`/`allow`指令. 使用规则为

location / {
    deny  192.168.1.1;         ## 单一IP
    allow 192.168.1.0/24;    ## IP段
    allow 2001:0db8::/32;    ## IPv6
    deny  all;
}
类似于`iptables`, 匹配到就退出, 否则继续向下执行.

有效作用域为`http`, `location`, `server`.

但如果配置了CDN, 用户请求是通过CDN服务器转发的, 屏蔽用户来源IP是无效的. 此时需要通过用`if`语句判断`$http_x_forwarded_for`条件. 举例如下

```conf
if ( $http_x_forwarded_for ~* '^180\.97\.106\.' ) {
    return 403;
}
```

关于这个访问控制的最佳实践, 创建一个`block.ip`文件, 在主配置文件中调用, 之后添加新的访问控制时单独在这个文件中添加即可.

## 12. 开启浏览器端压缩传输

```conf
## 开启压缩传输
gzip on;
## 大于1K的才压缩，一般不用改
gzip_min_length 1k;

gzip_buffers 4 16k;

gzip_http_version 1.1;
## 压缩级别，1-10，数字越大压缩的越好文件越小，压缩时间也越长
gzip_comp_level 2;
## 指定需要压缩的文件类型. 可在浏览器响应头的`Content-Type`字段中查看(不同的教程好像文件meta类型也不一样, 根据实际情况来吧)
gzip_types text/plain application/javascript text/css application/xml text/javascript application/x-httpd-php
## 听说图片类型默认是压缩过的, 所以其实不用两次压缩
## image/jpeg image/gif image/png;
gzip_vary off;
## IE6不支持Gzip，不对它Gzip了
gzip_disable "MSIE [1-6]\.";
```

被压缩过的资源可以在浏览器响应头中查看, `Content-Encoding:gzip`, 且其值`Content-Length`.

如果想要压缩的数据没有被压缩, **请确认`gzip_types`中是否有指定正确的文件类型(`Content-Type`)**.

## 13. rewrite时expires指令生效的作用域

```
location / {
    root html;
    try_files $uri $uri/index.html /static;
}

location /static/ {
    root html/static;
}
```

`rewrite`指令或`try_files`发起的内部跳转行为, 还是比较独立的. 比如, 如果在第1个`location`块中加入`expires`指令, 那么当访问的uri被重写到`/static/`的块的时候, `/static`下返回的资源是不会被缓存的.