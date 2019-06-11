# wrk error-connect,read,write

情景描述

win10系统, 安装docker桌面版2.0.0.3(build: 8858db3, docker engine: 18.09.2)

启动两个容器

1. 172.17.0.2: centos7容器, 运行nginx
2. 172.17.0.3: alpine容器, 运行wrk进行压测

nginx容器只提供`index`静态页面服务, 端口映射`3001:3001`.

在alpine容器中执行wrk命令

```
$ wrk -t 16 -c 10000 -d 30 --timeout 15 --latency http://192.168.0.8:3001
Running 30s test @ http://192.168.0.8:3001
  16 threads and 10000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.00us    0.00us   0.00us    -nan%
    Req/Sec     0.00      0.00     0.00      -nan%
  Latency Distribution
     50%    0.00us
     75%    0.00us
     90%    0.00us
     99%    0.00us
  0 requests in 30.10s, 0.00B read
  Socket errors: connect 0, read 117, write 94405, timeout 0
Requests/sec:      0.00
Transfer/sec:       0.00B
```

如上述结果所示, 1w个连接发起请求, 竟然没有一个成功的...但是nginx日志有输出, 说明的确有在处理请求.

其实`c2000`的结果还行, `c5000`时结果就不可接受了, 因为错误数量太多了, 且write错误要占大多数.

docker桌面版对每个容器的的资源限制都是8核2G, 但是实际上nginx的消耗远远到不了这个水平, cpu大概能跑满, 但是内存只用了不到30M.

于是我就去查wrk的write error是由哪些原因引起的. 有说是因为系统的socket的写缓冲区不足的, 也有说在写数据时连接已经断开的(报broken pipe那种). 但感觉都不对, 差点要去改注册表中的TCP相关配置了.

后来突然意识到可能是由于经过了`192.168.0.8`这个宿主机地址, 这个是电脑的无线网卡获取的IP地址, 我们在容器内访问宿主机地址, 再通过端口映射访问另一个容器, 多了一层转发, 造成了这个结果.

事实证明我的猜想是正确的, 如果在wrk容器中直接请求nginx所在容器的地址, 错误率会低很多.

```
$ wrk -t 8 -c 10000 -d 30 --timeout 15 --latency http://172.17.0.2:3001
Running 30s test @ http://172.17.0.2:3001
  8 threads and 10000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   127.72ms  376.00ms  14.05s    98.43%
    Req/Sec     5.14k     1.50k   10.14k    69.21%
  Latency Distribution
     50%   84.15ms
     75%  146.92ms
     90%  201.81ms
     99%    1.01s
  1205871 requests in 30.14s, 4.42GB read
  Socket errors: connect 0, read 1094, write 0, timeout 6
Requests/sec:  40004.56
Transfer/sec:    150.28MB
```

不过奇怪的是, 虽然请求量上去了, 但是nginx所在容器的establish连接数倒是少得很, 不到7500, 更多的时候是2000多. 应该是由于瞬时处理速度很快, 没有给10000连接同时处理的机会.

这里先不考虑为什么经过了win的物理网卡导致write error的问题, 估计要做win下的socket调优.
