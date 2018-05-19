# Nginx + uWSGI + Flask部署

## 1. CGI的一点理解

`uWSGI`与`FastCGI`, 在一定程度上也与JavaWeb中的`servlet`类似, 是`http`服务器与应用程序通信的桥梁. 

`http`服务接收到系统协议栈中传来的`tcp`底层包, 可以对其以`http`协议进行解读, 比如可以获取http请求头的协议类型(http/https), 目标路径, 访问端口, 请求参数等信息, 然后生成http响应.

在更后端的应用程序层面, 一般程序无法直接解析http协议, 而是在http服务器与应用程序中间部置了`CGI`(Common Gateway Interface). 于是, 构建在http协议之上, 应用程序又拥有了更为方便的操作cookie/session, 读取表单数据等接口. 比如php通过`$_POST['表单name名称']`就可以获取前端页面传来的数据, `python`/`servlet`中存在可读的`request`对象, 后者甚至可以直接操作`response`对象(我想这正是为什么Java Web比诸多脚本语言强大的表现之一吧), 就是CGI的功劳.

php/python/java各自拥有更适合自己的CGI, 于是衍生出`uWSGI`与`FastCGI`这种独立的`CGI`类型, 但是他们做的事情是一样的.

还有一句, 上面的都是瞎掰的, 有什么疑问去百度找找.

## 2. 环境搭建

uwsgi与nginx都使用源码安装, 这样配置为方便一点, 目录也可以自定义.

nginx的安装就不多说了, 在另一篇文章中可以找到, 安装目录在`/usr/local/nginx`, 目录属主为nginx, 并且以`nginx`用户运行.

[uwsgi的源码地址](https://github.com/unbit/uwsgi), 本文所使用的是tag中`2.0.14`版本.

解压后直接执行`make`命令, 会在当前目录生成`uwsgi`可执行文件, 目前对我们有用的就只是这个可执行文件...其他的还不会用, 有待探索.

手动在`/usr/local`下创建`uwsgi`目录, 并在其中创建如下目录结构, 这种结构借鉴了nginx的安装目录, 没什么特殊的优点, 个人喜好而已...

```
uwsgi/
├── bin
│   └── uwsgi
├── etc
└── var
    ├── log 
    └── run
```

将上面生成的`uwsgi`可执行文件放在`bin`子目录, `etc`下用来放置`uwsgi`的配置文件, `var/log`存放的是日志文件, `var/run`存放的是pid文件以及sock文件, sock文件的作用在下面会有讲到.

**将`uwsgi`目录的属主也修改成`nginx`, 方便与nginx进程"沟通".**

## 3. 配置

我们使用`uwsgi`的emperor模式运行. 类似于nginx的vhost(虚拟主机), 即一个主进程负责请求分发及子进程管理, 子进程负责响应与执行.

在`etc`目录下创建如下结构, `emperor.ini`文件就是主进程即管理进程的配置, `vassals`目录下就是各子进程的配置, 这里我们只配置一个网站, 所以只有一个`skycmdb.ini`.

```
etc
├── emperor.ini
└── vassals
    └── skycmdb.ini
```

`emperor.ini`文件内容如下

```ini
[uwsgi]
## 子进程配置目录
emperor = /usr/local/uwsgi/etc/vassals

## 与nginx使用同样的用户运行, 不推荐使用root.
uid = nginx
gid = nginx

enable-threads = ture
master = ture
workers = 8

## 指定pid文件与日志文件路径
pidfile = /usr/local/uwsgi/var/run/uwsgi.pid
daemonize = /usr/local/uwsgi/var/log/uwsgi.log
```

`vassals/skycmdb.ini`文件内容如下

```ini
[uwsgi]
## socket = 127.0.0.1:3001
socket = /usr/local/uwsgi/var/run/skycmdb.sock
## 目标工程目录, 如果不指定则该目录中定义的模块无法加载, 运行可能会出错
chdir = /var/www/html/skycmdb

uid = nginx
gid = nginx

logto = /usr/local/uwsgi/var/log/skycmdb.log
## 工程入口文件
wsgi-file = SKY.wsgi
```

> 来解释一下`skycmdb.sock`文件的作用, linux系统中使用`socket`系统API可以方便地进行跨主机的进程间通信, 也可以让同一主机上的进程间进行通信. 作用类似于系统提供的`共享内存`, `PIPE`, `信号量`等.

> `socket`API实现起来有两种方法, 一种就是为人熟知的`IP:Port`方式, 学过网络编程应该都不会陌生. 另一种就是使用`sock`文件, 后者适合于同主机的进程间通信. 与第一种方式相比, 第二种方式不会再将数据打包发到tcp层然后IP层, 然后再从底层将数据解析出来再传递给目标进程, 也就是说, `sock`文件进行的"网络"通信不会再经过内核协议栈, 而第一种至少还要过一遍回环网卡.

wsgi协议需要程序的入口文件主函数名为`application`...希望我没有理解错, 编写完flask工程后, 将主类以模块形式导入到这个`SKY.wsgi`文件中, 并通过它进入整个工程.

`SKY.wsgi`文件内容为

```python
# -*- coding: utf-8 -*-
from SkyFlask import sky as application


if __name__ == '__main__':

```

然后, `uwsgi`的启动方式, 启动时指定加载`emperor`配置文件即可.

```
$ /usr/local/uwsgi/bin/uwsgi /usr/local/uwsgi/etc/emperor.ini
```

可以看到`var`目录下有如下文件生成

```
var
├── log 
│   ├── skycmdb.log
│   └── uwsgi.log
└── run
    ├── skycmdb.sock
    └── uwsgi.pid
```

相应的, `uwsgi`的重启与关闭命令如下

```
$ /usr/local/uwsgi/bin/uwsgi --reload /usr/local/uwsgi/var/run/uwsgi.pid
$ /usr/local/uwsgi/bin/uwsgi --stop /usr/local/uwsgi/var/run/uwsgi.pid
```

------

nginx配置如下

```conf
location / {
        include uwsgi_params;
        ## uwsgi以--http参数启动时可以使用fastcgi_pass连接      
        ## fastcgi_pass 127.0.0.1:5000;
        ## uwsgi以--http参数启动时也可以使用proxy_pass连接      
        ## proxy_pass http://127.0.0.1:5000;
        ## uwsgi使用--socket启动, 并且是指定的IP:端口的形式     
        ## uwsgi_pass 127.0.0.1:5000;
        ## uwsgi使用--socket启动, 但指定的是sock文件名          
        uwsgi_pass unix:/usr/local/uwsgi/var/run/skycmdb.sock;
}
```

启动nginx, 完成.

## 4. 完善

### 4.1 便捷启动

**1. 使用命令别名**

```
alias uwsgi-start='/usr/local/uwsgi/bin/uwsgi /usr/local/uwsgi/etc/emperor.ini'
alias uwsgi-reload='/usr/local/uwsgi/bin/uwsgi --reload /usr/local/uwsgi/var/run/uwsgi.pid'
alias uwsgi-stop='/usr/local/uwsgi/bin/uwsgi --stop /usr/local/uwsgi/var/run/uwsgi.pid'
```

可追加在`/root/.bashrc`, 供root用户使用(root启动uwsgi时进程权限依然是`uid`中指定的用户), 也可以放置在`/etc/profile`中, 全局调用.

**2. CentOS7服务脚本**

```
[Unit]
Description=uWSGI Emperor
After=syslog.target

[Service]
Type=forking
ExecStart=/usr/local/uwsgi/bin/uwsgi --ini /usr/local/uwsgi/etc/emperor.ini
PIDFile=/usr/local/uwsgi/var/run/uwsgi.pid
Restart=always
KillSignal=SIGQUIT
ExecReload=/usr/local/uwsgi/bin/uwsgi --reload $MAINPID

[Install]
WantedBy=multi-user.target
```

将其命名为`uwsgi.service`, 放在`/usr/lib/systemd/system`目录下, 赋予其执行权限755.

然后执行`systemctl daemon-reload`重新加载服务脚本列表.

```
## 启动
$ systemctl start uwsgi
## 重启, 注意是restart不是reload
$ systemctl restart uwsgi
## 关闭
$ systemctl stop uwsgi
```

## 5. 注意点

`emperor`与`vassals`都配置了`uid`与`gid`为普通用户, 但`emperor`生成的日志文件属主依然是root, 但emperor进程的启动用户为uid字段指定用户, 并且vassals生成的日志文件与sock文件属主是uid指定的用户, 所以要注意目录及权限的配置问题. 比较凌乱.

因为`uwsgi`生成的`skycmdb.sock`文件需要让nginx进程也可以读写, 所以其属主必须也设置为`nginx`, 否则uwsgi日志中会出现`Permission Denied`.
