# wrk使用

参考文章

1. [wrk 压力测试 http benchmark POST接口](https://www.cnblogs.com/felixzh/p/8400729.html)

```
$ wrk -t 5 -c 5 -d 1 https://www.baidu.com
Running 1s test @ https://www.baidu.com
  5 threads and 5 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    42.04ms   30.82ms 116.08ms   77.78%
    Req/Sec    22.48     12.40    50.00     44.00%
  116 requests in 1.10s, 1.73MB read
Requests/sec:    105.51
Transfer/sec:      1.57MB
```

`-t`指定线程数, `-c`指定并发连接数. `-d`指定压测时间.

在wrk命令中, 数值可以使用(1k, 1M, 1G)表示, 时间也可以使用(2s, 2m, 2h)表示.

wrk的输出结果, 以线程为单位进行统计(Thread Stats), 只有两个指标

1. `Latency`: 延迟, 可以是响应时间.
2. `Req/Sec`: 每个线程每秒的请求数量.

横坐标为统计维度

- `Avg`: 平均值
- `Stdev`: 标准偏差(结果的离散程度，越高说明越不稳定.)
- `Max`: 最大值
- `+/- Stdev`: 正负一个标准差占比(结果的离散程度，越大越不稳定). 
