# Pypi-使用nginx搭建内网代理缓存服务器

参考文章

1. [使用nginx做pypi的反向代理，搭建pypi官方镜像](https://haoluobo.com/2012/02/nginxpypi/)
2. [利用Proxy Cache使Nginx对静态资源进行缓存](http://blog.csdn.net/czp11210/article/details/28596649)
3. [Simple PyPi Caching Proxy](https://gist.github.com/dctrwatson/5785638#file-nginx-conf)
    - md, 一点都不simple

我们知道, 受天朝网络限制, 使用pip/easy_install工具直接从python官网下载安装包是多么痛苦的一件事, 不稳定还经常抽风(*￣︿￣). 

## 1. 适合个人开发者的镜像源方式

很多时候, 个人开发者经常按照如下方法使用国内的pypi镜像源.

1. 在命令行中使用

```
pip install -i http://mirrors.aliyun.com/pypi flask
pip install --index-url http://mirrors.aliyun.com/pypi flask
```

2. 写在配置文件中作为全局配置(/etc/pip.conf文件, 如果没有可手动创建)

```ini
[global]
trusted-host = mirrors.aliyun.com
index-url = http://mirrors.aliyun.com/pypi
```

然后直接`pip install flask`. 这种方法更方便.

## 2. 企业的选择

但对于企业来说, 内部网络环境与公网是隔绝的, 只通过特殊的入口进行转发, 无法直接使用公共镜像源. 这时有两种选择, 一种是搭建供内部使用的私有镜像源, 定时同步官方数据. 另一种是在网络选择一台拥有外部访问权限的服务器, 使用nginx/haproxy等中间件作为跳转.

第一种做法在网络上有很多相关教程, pypi库中也有很多或强大或简便的工具, 在这篇文章里不作介绍.我们主要尝试第二种方法.

第二种方法感觉有点low, 但却是最简单方便还节省资源的方式. 

我们选择使用nginx作为正向代理(没错, 我就是觉得**参考文章1**中的反向代理的说法不正确, 这种用法明显是正向的嘛╮（╯＿╰）╭). 而且, 我们还可以启用nginx的缓存功能, 把下载过的包缓存到本地...想想还有些小激动(￣︶￣)↗

------

按照个人镜像源的设置, 本来想直接使用如下配置, 但是pip设置代理后下载失败.

```conf
    location /pypi {
        proxy_pass http://mirrors.aliyun.com/pypi/simple/;
    }
```

错误的配置不再细说, 我们需要分析一下pip安装流程. 

仔细想想, 其实pip安装与yum安装有相似的地方, 一定有一个索引地址, 然后指向一个实际下载地址. 上面的配置中我们只指定了索引地址, 所以下载的时候找不到文件.

我们访问`http://mirrors.aliyun.com/pypi/simple/`, 可以看到页面中列出的是无数包的超链接, 这便是索引地址, 也就是pip中`--index-url`的设置项. 

点进`flask`, 进入`http://mirrors.aliyun.com/pypi/simple/flask/`页面, 这里面列出了`pip install flask`时需要下载的文件(还可以指定版本, 因为列表里的链接都包含了版本号), 例如`http://mirrors.aliyun.com/pypi/packages/f3/46/53d83cbdb79b27678c7b032d5deaa556655dd034cc747ee609b3e3cbf95b/Flask-0.10.tar.gz#md5=92bc6b6ebd37d3120c235430a0491a15`. 

所以除了索引页的跳转, 我们还需要转发对下载包路径`packages`的请求.

需要注意的是, 在下载页, 下载链接是以相对路径存在的(使用chrome F12开发者工具可以查看), 并且pip没有chrome辣么聪明, 后者会把`http://www.baidu.com/../package/`解释成`http://www.baidu.com/package/`再访问, pip貌似不会, 所以它访问这种类型的链接会得到404.

因此不要图简单, location最好按照镜像源的设置来, 长一点就长一点吧.

下面是nginx示例配置

```conf
    ## 缓存目录(/usr/local/nginx/cache/)不用手动创建, nginx重启会自动生成
    ## keys_zone=pypi:100m 表示这个zone名称为pypi，分配的内存大小为100MB
    ## levels=1:2 表示缓存目录的第一级目录是1个字符，第二级目录是2个字符，即/usr/local/nginx/cache/pypi/a/1b这种形式
    ## inactive=1d 表示这个zone中的缓存文件如果在1天内都没有被访问，那么文件会被cache manager进程删除掉
    ## max_size=10g 表示这个zone的硬盘容量为10GB
    proxy_cache_path  /usr/local/nginx/cache/pypi levels=1:2 inactive=1d keys_zone=pypi:100m max_size=10g;
    server {
        listen 80;
        gzip off;
        server_name pypi.sky-mobi.com;
        access_log logs/pypi.log main;
        error_log logs/pypi.err.log;
        location /pypi/simple {
            proxy_pass http://mirrors.aliyun.com/pypi/simple/;
            proxy_cache pypi;                                   ##缓存zone
            proxy_cache_valid  any 2400h;
        }
        location /pypi/packages {
            proxy_pass http://mirrors.aliyun.com/pypi/packages;
            proxy_cache cache;
            proxy_cache_valid  any 24000h;
        }
    }
```

对应pip.conf的配置为

```ini
## /etc/pip.conf
[global]
trusted-host = pypi.sky-mobi.com
index-url = http://pypi.sky-mobi.com/pypi/simple
```

当然还有`pypi.sky-mobi.com`的DNS解析.

完成!
