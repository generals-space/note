# 压测工具Jmeter

参考文章

1. [使用Jmeter进行http接口性能测试](https://www.cnblogs.com/star91/p/5059222.html)
    - Jmeter的安装启动及简单的的使用方法, 生成表格及图表报告.

Jmeter的安装启动可以见参考文章1. 本文只介绍Jmeter常用的参数设置

## 线程组

线程数: 该线程组有多个个线程数, 每个线程表示一个客户端连接
`Ramp-Up Period(单位: 秒)`: 设置多少秒内启动全部线程, 如果为0, 则并发启动所有线程.
`Loop Count`循环次数: 执行次数.

真正的请求总数为: 线程数 * 循环次数, 没有限制时间.
