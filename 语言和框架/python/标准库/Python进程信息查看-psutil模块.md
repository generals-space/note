# Python进程信息查看-psutil模块

参考文章

1. [python笔记之psutil模块 - 博客园](http://www.cnblogs.com/liu-yao/p/5678157.html)

2. [psutil documentation - pypi官方文档](http://pythonhosted.org/psutil/#psutil.Process.create_time)

3. [【记录】折腾Python中的psutil：一个用于获得处理器和系统相关信息的模块](https://www.crifan.com/try_python_psutil/)

## 1. ps命令 - 进程信息

```py
import psutil

pids = psutil.pids()  ## 得到系统当前所有进程的pid列表
## [1, 2, 3, 7, 8, 9, 10, 11, 12, ... 60360, 60419]
p = psutil.Process(1)   ## 得到指定pid的进程对象, Process()的参数必须为整型

# 进程名, bash, python, ruby这类依赖'执行环境'的进程, 可以直接得到可执行文件的名称, 但java进程只能得到java
p.name()                    
# 进程的可执行文件路径, 这个更不准确, shell脚本或python进程等只会得到bash, python的路径而不是对应的启动脚本路径, java更是只能得到java的路径, 只有nginx, httpd这种C进程可以
p.exe()
p.cwd()                     # 进程的工作目录绝对路径, 一般没什么用
p.cmdline()                 # 是ps -ef得到的最后一个字段的值, 不过是list类型
p.status()                  # 进程状态, sleeping...字符串类型
p.create_time()             # 进程创建时间, float类型, 时间戳
import datetime
# 得到'2017-09-16 13:54:53', 字符串类型
## 其中fromtimestamp()可以将时间戳转化为通用datetime对象, 可用于数据库存储
datetime.datetime.fromtimestamp(p.create_time()).strftime("%Y-%m-%d %H:%M:%S")
p.uids()                    # 得到puids对象, 进程uid信息
p.gids()                    # 得到pgids对象, 进程的gid信息
p.pid()                     # 当前进程的进程号
p.ppid()                    # 当前进程的父进程号
p.cpu_times()               # 进程的cpu时间信息,包括user,system两个cpu信息
p.memory_info()             # 得到pmem对象, 包括进程内存rss,vms信息
p.io_counters()             # 得到pio对象, 进程的IO信息,包括读写IO数字及参数
p.connections()             # 得到pconn对象, 与当前进程建立连接的socket信息
p.children()                # 得到列表结果, 成员为Process对象, 是当前进程的**直接**子进程, 当然还有parent方法
p.num_threads()             # 进程开启的线程数
p.username()                # 进程的启动用户

p.kill()                    # kill掉当前进程, 还有其他操作
```

至于其他的可用方法, 可以通过`dir(进程对象)`查看.

### 异常处理

```py
    pids = psutil.pids()
    for pid in pids:
        try:
            p = psutil.Process(pid)
        except psutil.NoSuchProcess, pid:
            print "no process found with pid = %s" % (pid)
```

## 2. top, free - 系统信息相关