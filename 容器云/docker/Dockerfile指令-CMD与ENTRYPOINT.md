# Dockerfile指令-CMD与ENTRYPOINT

参考文章

1. []

`CMD`与`ENTRYPOINT`两者有相似之处, 都是镜像在实例化成容器时执行的命令.

首先要明白, `-d`(服务模式)和`-it`(交互模式)两者后面都可以指定要执行的命令.

以服务模式启动容器, 容器中运行`tail -f`命令.

```
$ docker run -d  centos:6 tail -f /etc/yum.conf
```

> 虽然看起来`/bin/bash`也是长驻进程, 但它不能运行在`-d`模式下. 若要运行, 需得使用`-dt`. tty模式.

------

进入正题, 如果一个镜像在封装之时没有通过`CMD`或是`ENTRYPOINT`指令指定在启动时执行的命令, 就只能在命令行启动时手动指定.

即, `CMD`与`ENTRYPOINT`的目的就是可以在容器启动时不必手动执行这一过程.

这也要看镜像的, 像centos, ubuntu这种系统镜像是不会为你指定启动命令的, 我们得自己定义. 而其他与nginx, mysql这种应用级镜像, 则一般会带有启动命令, 目的是简化操作, 直接使用.

`CMD`与`ENTRYPOINT`的区别在于, 使用`CMD`创建的镜像, 如果在启动容器时手动指定了命令, 会覆盖`CMD`指定的命令. 而使用`ENTRYPOINT`创建的镜像, 会同时执行通过`ENTRYPOINT`指定与通过手动指定的命令, 不会发生覆盖.

## 指令格式

```
CMD ping www.baidu.com 
CMD ["/bin/ping","www.baidu.com"]
CMD ["参数1","参数2"]
```

前两种是通过CMD指定在容器启动时执行指令的两种格式, 第3种是将CMD指定的列表当作`ENTRYPOINT`的参数, 不能单独使用.

同样, `ENTRYPOINT`也有几种可用的格式.

```
ENTRYPOINT ping www.baidu.com 
ENTRYPOINT ["/bin/ping","www.baidu.com"]
```

> md, 列表里必须用双引号...

> 注意1: `ENTRYPOINT ping www.baidu.com`这种格式依然会导致被命令行指定的命令覆盖, 第2种列表格式则不会.

> 注意2: 做这种命令覆盖的实验时, 不建议使用`echo`这种命令, 因为容器内外标准输出可能被阻断, 尽量写入到文件.

其实`ENTRYPOINT`这种不会覆盖的命令很鸡肋. 以如下dockerfile为例

```
FROM centos:6
ENTRYPOINT ["tail", "-f", "/etc/passwd"]
```

容器在启动时执行`tail -f /etc/passwd`, 这种情况下在`docker run`时就最好使用`-d`选项了. 并且我在末尾又通过指定`tail -f /etc/yum.conf`命令. 结果等容器启动后, 进入容器查看进程发现了如下情况

```
[root@deabb9078762 ~]# ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 07:16 ?        00:00:00 tail -f /etc/passwd tail -f /etc/yum.conf
root         7     0  0 07:16 ?        00:00:00 su - root
root         8     7  0 07:16 ?        00:00:00 -bash
root        19     8  0 07:16 ?        00:00:00 ps -ef
```

这很让人无语, 因为这相当于在命令行指定的命令直接附加到`ENTRYPOINT`所指定的命令后面, 完全没有预想的那么神.

------

CMD可以为ENTRYPOINT提供参数，ENTRYPOINT本身也可以包含参数，但是你可以把那些可能需要变动的参数写到CMD里而把那些不需要变动的参数写到ENTRYPOINT里面例如：

```
FROM centos:6
ENTRYPOINT ["top", "-b"]   
CMD ["-c"]  
```

把可能需要变动的参数写到CMD里面。然后你可以在docker run里指定参数，这样CMD里的参数(这里是-c)就会被覆盖掉而ENTRYPOINT里的不被覆盖。