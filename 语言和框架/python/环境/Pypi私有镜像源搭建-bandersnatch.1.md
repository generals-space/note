# Pypi私有镜像源搭建-bandersnatch

参考文章

1. [bandersnatch 2.0.0](https://pypi.python.org/pypi/bandersnatch)
2. [bandersnatch源码](https://bitbucket.org/pypa/bandersnatch)
3. [使用bandersnatch搭建pypi 本地镜像源（20211026）](https://blog.csdn.net/liurizhou/article/details/120980817)
    - `download-mirror = https://pypi.doubanio.com/`: 从官网获取元数据信息, 获取到以后再换成国内源进行下载
    - 不过这个选项貌似是 5.0.0+ 版本以后才支持的, 那就需要 python3.8+

python官网倒是的确收集了很多私有镜像的[解决方案](https://wiki.python.org/moin/PyPiImplementations), 不过呵呵, 年久失修, 很多都找不到了. `bandersnatch`是无意间发现的, 貌似是一个不错的选择.

## 1. 安装

> bandersnatch依赖于python3.5+

首先安装好`python3.5+`与`pip`环境, 然后安装`2.0.0`版本的bandersnatch(目前最新的).

(md, pypi官网和bandersnatch官网都说要下一个requirements.txt中的依赖包...然而那是`1.11`版本的, 坑啊X﹏X)

```
pip install bandersnatch==2.0.0
```

## 2. 配置

第一次执行`bandersnatch mirror`, 将会创建一个新的配置文件`/etc/bandersnatch.conf`(如果配置文件已经存在就开始同步了).

```log
$ bandersnatch mirror
2017-06-20 11:28:52,473 WARNING: Config file '/etc/bandersnatch.conf' missing, creating default config.
2017-06-20 11:28:52,473 WARNING: Please review the config file, then run 'bandersnatch' again.
```

然后开始定义我们所需要的选项.

- `directory`: 同步数据存放的位置, 如果目标路径不存在, 则bandersnatch会自动创建.
- `master`: 源站地址, 默认为`https://pypi.python.org`. 此项貌似必须是`https`协议.
- `timeout`: 我的理解是, 从源站下载的tcp连接超时时间. 超过这个时间就重开一个进程而不是被一个连接卡死, 默认值为10(s), 较低会好一点.
- `workers`: 下载资源的进程数, 默认为3, 不可以超过10(貌似是因为官网源站的限制, 所以bandersnatch不会提供超过10个进程去下载)

配置完成, 启动, 再次执行`bandersnatch mirror`.

...貌似是前端启动啊(⊙﹏⊙)

好吧, `bandersnatch mirror`会一次性将源站数据全部down下来. 如果我们需要定时操作的话, 需要写crontab定时任务. 这样也避免了终端日志输出的问题.

```
## 4小时同步一次
* */4 * * * bandersnatch的绝对路径 mirror
```

于是, 在同步路径下会出现如下目录结构

```
.
├── generation
├── todo
└── web
    ├── local-stats
    │   └── days
    ├── packages
    │   └── ec
    │       └── 65
    │           └── c0116953c9a3f47de89e71964d6c7b0c783b01f29fa3390584dbf3046b4d
    │               └── 0805nexter-1.1.0.zip
    └── simple
        └── 1020-nester
            └── index.html
```

等着吧, bandersnatch会慢慢同步下来的.

...ok, 现在我们需要配置web服务器了. 以nginx为例, 端口域名同步路径记得自定义啊.

```conf
    server {
        listen 127.0.0.1:80;
        server_name pypi.sky-mobi.com;
        autoindex on;
        charset utf-8;
        root 同步数据路径/web;
    }
```

然后客户端`/etc/pip.conf`只需要配置成如下形式即可

```ini
[global]
trusted-host = pypi.sky-mobi.com
index-url = http://pypi.sky-mobi.com/simple
```

完成.

------

2017-06-22新增

参考[pip 更换软件镜像源](http://www.jianshu.com/p/785bb1f4700d)

镜像源方式只能提高下载速率, 无法加速search操作. 因为pip的search行为使用的是xmlrpc协议, 默认直接与`https://pypi.python.org/pypi`连接, 内网环境无法完成这个操作. 索性内网服务器一般不会进行search, 问题也不是很大.
