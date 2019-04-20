# Python-创建守护进程

参考文章

1. [守护进程 & 创建守护进程 & fork一次和fork两次的区别](http://blog.csdn.net/asd7486/article/details/51966225)

2. [Python中创建守护进程](http://www.cnblogs.com/zhiguo/p/3370599.html)

3. [Python实例浅谈之五Python守护进程和脚本单例运行](http://blog.csdn.net/taiyang1987912/article/details/44850999)

4. [Python守护进程daemon实现](https://blog.csdn.net/zhaihaifei/article/details/70213431)

## 1. 最简单的守护进程

通过`setsid()`函数, 将当前进程设置为新的session中的leader, 这样就可以摆脱当前终端的限制.

```py
#!/usr/bin/python 
#!encoding: utf-8

import sys, os

if __name__ == '__main__':
    try:
        pid = os.fork()
        if pid > 0:
            print('i am parent process')
            ## 父进程直接退出
            sys.exit(0)
    except OSError, e:
        print(e)
    
    os.setsid()
    os.execv('/bin/ping', ['ping', 'www.baidu.com'])
```

但是此进程的标准输入, 标准输出与标准错误还与当前终端有关联, 执行上述代码, 可以看到执行终端中不断有ping的日志在输出, 而由于我们已经将其设置了新的session, 无法再使用`Ctrl+C`将其终止, 关闭该执行终端也不行, 只能使用kill强制结束它. 

> ps: `ping`进程的父进程号为1

## 2. 文件描述符重定向

下面, 我们需要将创建守护进程的过程写得标准一些. 涉及到的函数有

- `os.chdir()`:  #chdir确认进程不保持任何目录于使用状态, 否则不能umount一个文件系统. 也可以改变到对于守护程序运行重要的文件所在目录. 

- `os.umask(0)`: ##重设文件创建掩码, 子进程会从父进程继承所有权限, 可以通过调用这个方法将文件创建掩码初始化成系统默认.  

- `os.setsid()`: #setsid调用成功后, 进程成为新的**会话组长**和新的**进程组长**, 并与原来的登录会话和进程组脱离. 

另外, 还需要重定向守护进程的标准文件描述符, 起码标准输入需要关闭, 标准输出/标准错误可以打开一个文件作为输出. 这时可能需要用到`os.dup2()`函数, 将守护进程的标准输出与标准错误的文件描述符修改成其他打开的文件描述符.

见下面的示例

```py
#!/usr/bin/python 
#!encoding: utf-8

import sys, os

if __name__ == '__main__':
    log_file = '/tmp/daemonit.log'
    try:
        pid = os.fork()
        if pid > 0:
            print('i am parent process')
            ## 父进程直接退出
            sys.exit(0)
    except OSError, e:
        print(e)
    
    os.chdir("/")
    os.setsid()
    os.umask(0)

    si = open('/dev/null', 'r')
    so = open(log_file, 'a+')
    se = open(log_file, 'a+', 0)
    os.dup2(si.fileno(), sys.stdin.fileno())
    os.dup2(so.fileno(), sys.stdout.fileno())
    os.dup2(se.fileno(), sys.stderr.fileno())

    os.execv('/bin/ping', ['ping', 'www.baidu.com'])
```

这下控制终端没有ping的输出了, 因为它们被重定向到`log_file`所指向的文件中了, 包括错误输出. 如下

```
$ tail -f /tmp/log_file
...
64 bytes from 115.239.210.27: icmp_seq=339 ttl=128 time=3.95 ms
64 bytes from 115.239.210.27: icmp_seq=340 ttl=128 time=6.92 ms
64 bytes from 115.239.210.27: icmp_seq=341 ttl=128 time=4.64 ms
ping: sendmsg: Network is unreachable
ping: sendmsg: Network is unreachable
ping: sendmsg: Network is unreachable
ping: sendmsg: Network is unreachable
ping: sendmsg: Network is unreachable
ping: sendmsg: Network is unreachable
From gateway (172.32.100.2) icmp_seq=358 Destination Host Unreachable
From gateway (172.32.100.2) icmp_seq=359 Destination Host Unreachable
From gateway (172.32.100.2) icmp_seq=360 Destination Host Unreachable
...
```

------

## 3. 比较完善的守护进程类

由于之前创建的进程几乎完全脱离了控制终端, 所以要停止/重启这种守护进程是非常麻烦的. 我们还希望有pid文件, 日志文件, 启动/停止/重启接口. 还有, 同一个工程只能启动一个实例, 可以通过为pid文件加锁实现.

下面是一个比较完整的python守护进程类示例代码, 基本可以满足我们的要求.

```py
#!/usr/bin/env python
#coding: utf-8

import sys, os, time, atexit, string
from signal import SIGTERM

class Daemon:
    def __init__(self, 
        pidfile, 
        stdin='/dev/null', 
        stdout='/dev/null', 
        stderr='/dev/null'
    ):
        #需要获取调试信息，改为stdin='/dev/stdin', stdout='/dev/stdout', stderr='/dev/stderr'，以root身份运行
        self.stdin = stdin
        self.stdout = stdout
        self.stderr = stderr
        self.pidfile = pidfile
    
    def daemonize(self):
        try:
            # 第一次fork，生成子进程，脱离控制终端
            pid = os.fork()
            if pid > 0:
                sys.exit(0)    #退出父进程
        except OSError, e:
            sys.stderr.write('fork #1 failed: %d (%s)\n' % (e.errno, e.strerror))
            sys.exit(1)
        
        os.chdir("/")    
        os.setsid()      
        os.umask(0)      
        
        try:
            #第二次fork，可以禁止进程再次打开终端  
            pid = os.fork()
            if pid > 0:
                sys.exit(0)
        except OSError, e:
            sys.stderr.write('fork #2 failed: %d (%s)\n' % (e.errno, e.strerror))
            sys.exit(1)
        
        #重定向文件描述符  
        sys.stdout.flush()
        sys.stderr.flush()
        si = file(self.stdin, 'r')
        so = file(self.stdout, 'a+')
        se = file(self.stderr, 'a+', 0)
        os.dup2(si.fileno(), sys.stdin.fileno())
        os.dup2(so.fileno(), sys.stdout.fileno())
        os.dup2(se.fileno(), sys.stderr.fileno())
        
        ## atexit注册退出函数，main函数执行完毕后回调.
        atexit.register(self.del_pidfile)
        pid = str(os.getpid())
        file(self.pidfile,'w+').write('%s\n' % pid)

    def get_pid(self):
        return os.get_pid()

    def del_pidfile(self):
        os.remove(self.pidfile)

    def start(self):
        #检查pid文件是否存在以探测是否存在进程  
        try:
            pf = file(self.pidfile,'r')
            pid = int(pf.read().strip())
            pf.close()
        except IOError:
            pid = None

        if pid:
            message = 'pidfile %s already exist. Daemon already running!\n'
            sys.stderr.write(message % self.pidfile)
            sys.exit(1)
        
        ## 创建守护进程, 并执行重载的run方法中的操作
        self.daemonize()
        self.run()
  
    def stop(self):
        try:
            ## 从pid文件中获取pid
            pf = file(self.pidfile,'r')
            pid = int(pf.read().strip())
            pf.close()
        except IOError:
            pid = None

        if not pid: 
            message = 'pidfile %s does not exist. Daemon not running!\n'  
            sys.stderr.write(message % self.pidfile)
            return  

        try:
            while 1:
                os.kill(pid, SIGTERM)
                time.sleep(0.1)
                #os.system('hadoop-daemon.sh stop datanode')
                #os.system('hadoop-daemon.sh stop tasktracker')
                #os.remove(self.pidfile)
        except OSError, err:
            err = str(err)
            if err.find('No such process') > 0:
                if os.path.exists(self.pidfile):
                    os.remove(self.pidfile)
            else:
                print str(err)
                sys.exit(1)

    def restart(self):
        self.stop()
        self.start()

    def run(self):
        """
        重载这个方法, 加入你自己的操作
        """

class MyApp(Daemon):
    def run(self):
        os.execv('/bin/ping', ['ping', 'www.baidu.com'])
if __name__ == '__main__':
    myApp = MyApp(
        pidfile = '/var/run/myapp.pid', 
        stdout = '/var/log/myapp.log'
    )

    if len(sys.argv) == 2:
        if 'start' == sys.argv[1]:
            myApp.start()
        elif 'stop' == sys.argv[1]:
            myApp.stop()
        elif 'restart' == sys.argv[1]:
            myApp.restart()
        else:
            print 'unknown command'
            sys.exit(2)
        sys.exit(0)
    else:
        print 'usage: %s start|stop|restart' % sys.argv[0]
        sys.exit(2)
```