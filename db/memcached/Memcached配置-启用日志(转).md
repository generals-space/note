# Linux配置Memcached日志(转)

[原文链接](http://chenzhou123520.iteye.com/blog/1925209)

`memcached`在linux上安装时并不支持显式地配置服务日志, 我们如果想要把memcached服务日志保存到日志文件中, 则需要在启动参数中进行配置. 

安装好`memcached`后, 我们可以通过`-h`选项查看`memcached`支持的参数: 

```
[chenzhou@localhost ~]$ /usr/local/memcached/bin/memcached -h  
...
...  
-v            verbose (print errors/warnings while in event loop)  
-vv           very verbose (also print client commands/reponses)  
-vvv          extremely verbose (also print internal state transitions)  
...
...
```

从上面可以看到, 启动memcached时有3个参数是和日志信息相关的: 

`-v`代表打印普通的错误或者警告类型的日志信息

`-vv`比`-v`打印的日志更详细, 包含了客户端命令和server端的响应信息

`-vvv`则是最详尽的, 甚至包含了内部的状态信息打印

你可以根据你的实际需要来选择对应的参数, 我这里使用`-vv`就OK了. 

由于我们需要把日志信息保存在文件中, 而不是在控制台输出, 而-vv等参数只能把日志信息输出在控制台. 所以我们需要对-vv参数的输出进行数据流重定向, 关于重定向的知识在这里就不细述了, 有兴趣的可以查下资料了解一下. 

综上, 启动memcached的命令如下. 

```shell
memcached -d  -u memcached -vv >> /tmp/memcached.log 2>&1
```

重点在最后的: `-vv >> /tmp/memcached.log 2>&1`.

`-vv >> /tmp/memcached.log`代表把`-vv`的输出重定向到`/tmp/memcached.log`文件中;

`2>&1`的意思是把错误日志也一起写入到该文件中;

启动成功后我们可以测试一下, 首先起两个terminal, terminal1用来查看日志信息, terminal2进行client操作.

terminal1: 启动memcached后默认的日志信息如下:

````
[chenzhou@localhost ~]$ tail -f /tmp/memcached.log   
<31 send buffer was 110592, now 268435456  
<30 server listening (udp)  
<31 server listening (udp)  
<30 server listening (udp)  
<31 server listening (udp)  
<30 server listening (udp)  
<30 server listening (udp)  
<31 server listening (udp)  
<31 server listening (udp)  
<32 new auto-negotiating client connection
```

terminal2: 使用telnet连接, 并向memcached里存入一个数据

```
[root@localhost bin]# telnet localhost 11211  
Trying 127.0.0.1...  
Connected to localhost.localdomain (127.0.0.1).  
Escape character is '^]'.  
set name 0 60 5 chenzhou
```

如上所示: 使用`set`命令存入`key`为name,` value`为chenzhou的键值对.

terminal1日志记录: 

```
32: Client using the ascii protocol  
<32 set name 0 60 5 chenzhou
```

这样, 我们的配置就生效了. 