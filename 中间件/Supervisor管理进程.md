# Supervisor管理进程

参考文章

[使用 supervisor 管理进程](http://www.ttlsa.com/linux/using-supervisor-control-program/)

[supervisord监控详解](http://www.2cto.com/os/201406/306622.html)

[Linux 进程管理与监控（supervisor and monit）](http://tchuairen.blog.51cto.com/3848118/1827716)

[supervisor(一)基础篇](http://lixcto.blog.51cto.com/4834175/1539136)及其后续

[Supervisor](http://supervisord.org)是一个用 Python 写的进程管理工具，可以很方便的用来启动、重启、关闭进程（不仅仅是 Python 进程）。除了对单个进程的控制，还可以同时启动、关闭多个进程，比如很不幸的服务器出问题导致所有应用程序都被杀死，此时可以用 supervisor 同时启动所有应用
程序而不是一个一个地敲命令启动。

`supervisord`管理的进程必须由supervisord来启动，并且管理的程序必须要是非Daemon程序，Supervisor会帮你把它转化为Daemon程序. 比如想要使用Supervisor来管理Nginx进程，就必须在Nginx配置文件中加入 daemon off让Nginx以非Daemon方式运行。

## 1. 安装

Supervisor 可以运行在 Linux、Mac OS X 上。如前所述，supervisor 是 Python 编写的，所以安装起来也很方便，可以直接用pip或是用yum安装, 如果是 Ubuntu 系统，还可以使用 `apt-get` 安装:

```
$ pip install supervisor
$ yum install supervisor
$ sudo apt-get install supervisor
```

## 2. 配置

Supervisor 相当强大，提供了很丰富的功能，不过大部分情况下只需要用到其中一小部分。安装完成之后，可以编写配置文件，来满足自己的需求。为了方便，一般把配置分成两部分：supervisord本身的配置(也相当于全局配置)和待管理的应用程序自己的配置。

> supervisor 是一个C/S模型的程序，`supervisord`是server端，对应的有client端`supervisorctl`.

### 2.1 supervisord的配置及启动

首先来看 supervisord 的配置文件。安装完supervisor之后，可以运行`echo_supervisord_conf` 命令输出默认的配置项，也可以将这些输出重定向到一个配置文件里作为模板：

去除里面大部分注释和“不相关”的部分，我们可以先看这些配置：

```conf
[unix_http_server]
file=/tmp/supervisor.sock   ; UNIX socket 文件，supervisorctl 会使用
;chmod=0700                 ; socket 文件的 mode，默认是 0700
;chown=nobody:nogroup       ; socket 文件的 owner，格式： uid:gid

;[inet_http_server]         ; HTTP 服务器，提供 web 管理界面
;port=127.0.0.1:9001        ; Web 管理后台运行的 IP 和端口，如果开放到公网，需要注意安全性
;username=user              ; 登录管理后台的用户名
;password=123               ; 登录管理后台的密码

[supervisord]
logfile=/tmp/supervisord.log ; 日志文件，默认是 $CWD/supervisord.log
logfile_maxbytes=50MB        ; 日志文件大小，超出会 rotate，默认 50MB
logfile_backups=10           ; 日志文件保留备份数量默认 10
loglevel=info                ; 日志级别，默认 info，其它: debug,warn,trace
pidfile=/tmp/supervisord.pid ; pid 文件
nodaemon=false               ; 是否在前台启动，默认是 false，即以 daemon 的方式启动
minfds=1024                  ; 可以打开的文件描述符的最小值，默认 1024
minprocs=200                 ; 可以打开的进程数的最小值，默认 200

; the below section must remain in the config file for RPC
; (supervisorctl/web interface) to work, additional interfaces may be
; added by defining them in separate rpcinterface: sections
[rpcinterface:supervisor]
supervisor.rpcinterface_factory = supervisor.rpcinterface:make_main_rpcinterface

[supervisorctl]
; 通过 UNIX socket 连接 supervisord，路径与 unix_http_server 部分的 file 一致
serverurl=unix:///tmp/supervisor.sock 
通过 HTTP 的方式连接 supervisord
;serverurl=http://127.0.0.1:9001 ; 

; 包含其他的配置文件, 即待管理的应用程序各自的配置, 可以是对这个文件而言的相对路径.
[include]
; 可以是 *.conf 或 *.ini
files = relative/directory/*.ini    
```

我们把上面这部分配置保存到`/etc/supervisord.conf`（或其他任意有权限访问的文件），然后启动 `supervisord`（通过`-c`选项指定配置文件路径，如果不指定会按照这个顺序查找配置文件：`$CWD/supervisord.conf, $CWD/etc/supervisord.conf, /etc/supervisord.conf`）：

```
$ supervisord -c /etc/supervisord.conf
```

### 2.2 应用程序配置

上面我们已经把 supervisrod 运行起来了，但是还没有被管理的进程. 现在可以添加我们要管理的进程的配置文件。可以把所有配置项都写到 `/etc/supervisord.conf` 文件里，但并不推荐这样做，而是通过 include 的方式把不同的程序（组）写到不同的配置文件里。

为了举例，我们新建一个目录`/etc/supervisor.d/`用于存放这些配置文件，相应的，把`/etc/supervisord.conf`里`include`部分的的配置修改一下：

```conf
[include]
files = /etc/supervisor/*.conf
```

假设有个用 Python 和 Flask 框架编写的web应用，取名`usercenter`，用[gunicorn](http://gunicorn.org/)做web服务器。工程目录位于`/home/leon/projects/usercenter`，`gunicorn`配置文件为`gunicorn.py`，`WSGI callable`是`wsgi.py`里的app属性。所以直接在命令行启动该web应用的方式可能是这样的：

```
$ cd /home/leon/projects/usercenter
$ gunicorn -c gunicorn.py wsgi:app
```

现在编写一份配置文件, 让`supervisord`管理这个进程（需要注意：用 supervisord 管理时，`gunicorn`自己的`daemon`选项需要设置为`False`）：

```conf
[program:usercenter]
; 程序的启动目录, 某些应用程序必需要进入到工程目录启动才可以, 因为某些模块是工程自定义的, 并未加入到系统中的模块搜索路径中.
directory = /home/leon/projects/usercenter 
; 启动命令，可以看出与手动在命令行启动的命令是一样的
command = gunicorn -c gunicorn.py wsgi:app  
autostart = true     ; 在 supervisord 启动的时候此web应用也自动启动
startsecs = 5        ; 启动 5 秒后没有异常退出，就当作已经正常启动了
autorestart = true   ; 程序异常退出后自动重启
startretries = 3     ; 启动失败自动重试次数，默认是 3
user = leon          ; 用哪个用户启动
redirect_stderr = true  ; 把 stderr 重定向到 stdout，默认 false
stdout_logfile_maxbytes = 20MB  ; stdout 日志文件大小，默认 50MB
stdout_logfile_backups = 20     ; stdout 日志文件备份数
; stdout 日志文件，需要注意当指定目录不存在时无法正常启动，所以需要手动创建目录（supervisord 会自动创建日志文件）
stdout_logfile = /var/log/usercenter.log

; 可以通过 environment 来添加需要的环境变量，一种常见的用法是修改 PYTHONPATH
; environment=PYTHONPATH=$PYTHONPATH:/path/to/somewhere
```

一份`supervisord`需要的配置文件至少需要一个 `[program:x]` 部分的配置，来告诉`supervisord`需要管理那个进程。`[program:x]` 块中的`x`表示进程名称, 可以自定义, 这个值会在客户端(`supervisorctl`或web界面)显示，在`supervisorctl`中可以通过这个值来对程序进行`start`、`restart`、`stop`等操作。

## 3. 客户端操作-supervisorctl

`supervisorctl`是`supervisord`的一个命令行客户端工具，用以查看被管理的应用程序列表, 状态, 及对其执行操作等. 执行此命令时需要指定与`supervisord`使用同一份配置文件，否则与`supervisord`一样按照顺序查找配置文件。

```
$ supervisorctl -c /etc/supervisord.conf
```

上面这个命令会进入supervisorctl的shell界面，然后可以执行不同的命令了：

```
status    # 查看程序状态
stop usercenter   # 关闭 usercenter 程序
start usercenter  # 启动 usercenter 程序
restart usercenter    # 重启 usercenter 程序
reread    ＃ 读取有更新（增加）的配置文件，不会启动新添加的程序
update    ＃ 重启配置文件修改过的程序
```

上面这些子命令都有相应的输出，除了进入`supervisorctl`的shell界面，也可以直接在终端运行：

```
## 初始启动Supervisord，启动、管理配置中设置的进程。
$ supervisord
## 停止某一个进程(programxxx)，programxxx为[program:chatdemon]里配置的值，这个示例就是chatdemon。
$ supervisorctl stop programxxx
## 启动某个进程
$ supervisorctl start programxxx
## 重启某个进程
$ supervisorctl restart programxxx
## 停止全部进程，注：start、restart、stop都不会载入最新的配置文件。
$ supervisorctl stop all
## 载入最新的配置文件，停止原有进程并按新的配置启动、管理所有进程。
$ supervisorctl reload
## 根据最新的配置文件，启动新配置或有改动的进程，配置没有改动的进程不会受影响而重启。
$ supervisorctl update

## 重启所有属于名为groupworker这个分组的进程(start,restart同理)
$ supervisorctl stop groupworker
```

## 4. 其他

除了`supervisorctl`之外，还可以配置`supervisrod`启动web界面执行管理操作，这个web后台使用`Basic Auth`的方式进行身份认证。

除了单个进程的控制，还可以配置group，进行分组管理。

经常查看日志文件，包括`supervisord`的日志和各个`pragram`的日志文件，程序crash或抛出异常的信息一半会输出到stderr，可以查看相应的日志文件来查找问题。

`supervisor`有很丰富的功能，还有其他很多项配置，可以在[官方文档](http://supervisord.org/index.html)获取更多信息.

## 5. 附录

```conf
; Notes:
;  - 不支持'~'或是"$HOME"这种变量形式, 用户主目录的环境变量可以使用"%(ENV_HOME)s"表示
;  - 分号表示注释, 但是行内注释需要分号前预留一个空格. 比如'a=b;这里是注释'是错误的, 'a=b ;这里是注释'才是正确的
;  - echo_supervisord_conf命令可以得到最原始的, 未经修改的配置文件模板, 可以当作参考

; web管理界面配置, 通过sock文件与http服务器通信
[unix_http_server]
file=/var/run/supervisor.sock   ; (the path to the socket file)
chmod=0700                      ; socket file mode (default 0700)
chown=root:root                 ; socket file uid:gid owner
username=devops                 ; (default is no username (open server))
password=hh1q2w3edd4r5t6y       ; (default is no password (open server))

; web管理界面配置, 通过ip:port与http服务器通信, 也可以直接对外服务, 是远程`supervisorctl`工具与web server的管理接口.
[inet_http_server]              ; inet (TCP) server disabled by default
port=*:19001                    ; (ip_address:port specifier, *:port for all iface)
username=devops                 ; (default is no username (open server))
password=hh1q2w3edd4r5t6y       ; (default is no password (open server))

[supervisord]
logfile=/var/log/supervisord.log ; (main log file;default $CWD/supervisord.log)
logfile_maxbytes=50MB        ; (max main logfile bytes b4 rotation;default 50MB)
logfile_backups=10           ; (num of main logfile rotation backups;default 10)
loglevel=info                ; (log level;default info; others: debug,warn,trace)
pidfile=/var/run/supervisord.pid       ; (supervisord pidfile;default supervisord.pid)
nodaemon=false               ; (start in foreground if true;default false)
minfds=1024                  ; 这个是最少系统空闲的文件描述符，低于这个值supervisor将不会启动。默认1024
minprocs=200                 ; 最小可用的进程描述符，低于这个值supervisor也将不会正常启动。默认200
;umask=022                   ; (process file creation umask;default 022)
user=root                 ; (default is current user, required if root)
;identifier=supervisor       ; (supervisord identifier, default is 'supervisor')
;directory=/tmp              ; 当supervisord作为守护进程运行的时候，启动supervisord进程之前，会先切换到这个目录
nocleanup=true              ; 这个参数当为false的时候，会在supervisord进程启动的时候，把以前子进程产生的日志文件(路径为AUTO的情况下)清除掉。想要看历史日志，当可以设置为true. 默认为false
;childlogdir=/tmp            ; 当子进程日志路径为AUTO的时候，子进程日志文件的存放路径。默认路径为$TEMP, 可以通过`python -c "import tempfile;print tempfile.gettempdir()"`命令查看
;environment=KEY="value"     ; (key value pairs to add to environment)
;strip_ansi=false            ; 这个选项如果设置为true，会清除子进程日志中的所有ANSI 序列, 即\n,\t这些字符。默认为false

; rpc配置块是supervisor远程操作的接口, ctl命令与web管理都需要这个接口
[rpcinterface:supervisor]
supervisor.rpcinterface_factory = supervisor.rpcinterface:make_main_rpcinterface

; supervisorctl管理接口, 也是supervisorctl工具需要加载的配置, 通过.sock文件或是端口任意一种方式通信
; 只要这里配置的username与password与上面unix|inet形式的http server的认证口令一致, 就可以管理supervisord(当然也可以在命令行中指定)
; 通过ip:port通信的方式, supervisorctl可以管理远程的supervisord服务
[supervisorctl]
;serverurl=unix:///var/run/supervisor.sock ; use a unix:// URL  for a unix socket
serverurl=http://127.0.0.1:19001 ; use an http:// url to specify an inet socket
username=devops              ; should be same as http_username if set
password=hh1q2w3edd4r5t6y                ; should be same as http_password if set
;prompt=mysupervisor         ; cmd line prompt (default "supervisor")
;history_file=~/.sc_history  ; use readline history if available

; 被管理进程的配置块, 建议写在include块中, 各个子服务分开
;[program:theprogramname]
;command=/bin/cat              ; the program (relative uses PATH, can take args)
;process_name=%(program_name)s ; process_name expr (default %(program_name)s)
;numprocs=1                    ; number of processes copies to start (def 1)
;directory=/tmp                ; directory to cwd to before exec (def no cwd)
;umask=022                     ; umask for process (default None)
;priority=999                  ; the relative start priority (default 999)
;autostart=true                ; start at supervisord start (default: true)
;startsecs=1                   ; # of secs prog must stay up to be running (def. 1)
;startretries=3                ; max # of serial start failures when starting (default 3)
;autorestart=unexpected        ; when to restart if exited after running (def: unexpected)
;exitcodes=0,2                 ; 'expected' exit codes used with autorestart (default 0,2)
;stopsignal=QUIT               ; signal used to kill process (default TERM)
;stopwaitsecs=10               ; max num secs to wait b4 SIGKILL (default 10)
;stopasgroup=false             ; send stop signal to the UNIX process group (default false)
;killasgroup=false             ; SIGKILL the UNIX process group (def false)
;user=chrism                   ; setuid to this UNIX account to run the program
;redirect_stderr=true          ; redirect proc stderr to stdout (default false)
;stdout_logfile=/a/path        ; stdout log path, NONE for none; default AUTO
;stdout_logfile_maxbytes=1MB   ; max # logfile bytes b4 rotation (default 50MB)
;stdout_logfile_backups=10     ; # of stdout logfile backups (default 10)
;stdout_capture_maxbytes=1MB   ; number of bytes in 'capturemode' (default 0)
;stdout_events_enabled=false   ; emit events on stdout writes (default false)
;stderr_logfile=/a/path        ; stderr log path, NONE for none; default AUTO
;stderr_logfile_maxbytes=1MB   ; max # logfile bytes b4 rotation (default 50MB)
;stderr_logfile_backups=10     ; # of stderr logfile backups (default 10)
;stderr_capture_maxbytes=1MB   ; number of bytes in 'capturemode' (default 0)
;stderr_events_enabled=false   ; emit events on stderr writes (default false)
;environment=A="1",B="2"       ; process environment additions (def no adds)
;serverurl=AUTO                ; override serverurl computation (childutils)

; 事件监听器配置
;[eventlistener:theeventlistenername]
;command=/bin/eventlistener    ; the program (relative uses PATH, can take args)
;process_name=%(program_name)s ; process_name expr (default %(program_name)s)
;numprocs=1                    ; number of processes copies to start (def 1)
;events=EVENT                  ; event notif. types to subscribe to (req'd)
;buffer_size=10                ; event buffer queue size (default 10)
;directory=/tmp                ; directory to cwd to before exec (def no cwd)
;umask=022                     ; umask for process (default None)
;priority=-1                   ; the relative start priority (default -1)
;autostart=true                ; start at supervisord start (default: true)
;startsecs=1                   ; # of secs prog must stay up to be running (def. 1)
;startretries=3                ; max # of serial start failures when starting (default 3)
;autorestart=unexpected        ; autorestart if exited after running (def: unexpected)
;exitcodes=0,2                 ; 'expected' exit codes used with autorestart (default 0,2)
;stopsignal=QUIT               ; signal used to kill process (default TERM)
;stopwaitsecs=10               ; max num secs to wait b4 SIGKILL (default 10)
;stopasgroup=false             ; send stop signal to the UNIX process group (default false)
;killasgroup=false             ; SIGKILL the UNIX process group (def false)
;user=chrism                   ; setuid to this UNIX account to run the program
;redirect_stderr=false         ; redirect_stderr=true is not allowed for eventlisteners
;stdout_logfile=/a/path        ; stdout log path, NONE for none; default AUTO
;stdout_logfile_maxbytes=1MB   ; max # logfile bytes b4 rotation (default 50MB)
;stdout_logfile_backups=10     ; # of stdout logfile backups (default 10)
;stdout_events_enabled=false   ; emit events on stdout writes (default false)
;stderr_logfile=/a/path        ; stderr log path, NONE for none; default AUTO
;stderr_logfile_maxbytes=1MB   ; max # logfile bytes b4 rotation (default 50MB)
;stderr_logfile_backups=10     ; # of stderr logfile backups (default 10)
;stderr_events_enabled=false   ; emit events on stderr writes (default false)
;environment=A="1",B="2"       ; process environment additions
;serverurl=AUTO                ; override serverurl computation (childutils)

; 进程组配置块, 多个相关进程可归为一组, 方便管理
;[group:thegroupname]
;programs=progname1,progname2  ; each refers to 'x' in [program:x] definitions
;priority=999                  ; the relative start priority (default 999)

; include配置块可以指定子配置文件, 
; 支持通配符. 多个文件可以使用空格或换行.
; 支持以此配置文件为基准的相对路径
;[include]
files = /etc/nginx.ini
```