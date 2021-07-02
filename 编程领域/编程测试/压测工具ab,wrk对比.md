# 压测工具ab,wrk对比

参考文章

1. [轻量级性能测试工具ab / wrk / locust 分析 & 对比](http://www.istester.com/tester/181.html)
2. [为什么 wrk 和 ab， locust 压测的结果差异这么大？](https://www.v2ex.com/t/423435)
3. [峰云 - wrk用lua脚本构建复杂的http压力测试](http://xiaorui.cc/2018/03/14/wrk%E7%94%A8lua%E8%84%9A%E6%9C%AC%E6%9E%84%E5%BB%BA%E5%A4%8D%E6%9D%82%E7%9A%84http%E5%8E%8B%E5%8A%9B%E6%B5%8B%E8%AF%95/)
4. [【测试设计】性能测试工具选择：wrk？jmeter？locust？还是LR？](https://www.cnblogs.com/Detector/p/8684658.html)
    - 给出了不同工具的优缺点及不同场景下选择工具的建议

2019-06-07

突然又发现了一个工具: 贝吉塔[tsenart/vegeta](https://github.com/tsenart/vegeta).

golang编写, 可生成图表, 待测试.

------

其实主要看`ab`与`wrk`两个工具, `wrk`是比较新的工具.

`wrk`使用的是 HTTP/1.1, 缺省开启的是长连接, 而ab使用的是HTTP/1.0, 缺省开启的是短链接.

`ab`有一个参数-`k`, 它会增加请求头`Connection: Keep-Alive`, 相当于开启了HTTP长连接, 这样做一方面可以降低测试服务器动态端口被耗尽的风险, 另一方面也有助于给目标服务器更大的压力, 测试出更接近极限的结果.

`wrk`由于默认就开启长连接所以不需要额外指定参数, 但是为测试短连接, 可以添加`-H "Connection: Close"`. 通过参数`-H`添加自定义的请求头来关闭长链接.

参考文章3中峰云说`ab`, `webbench`已经无法满足现在的测试需求了, 那么就直接选择的`wrk`工具吧.

另外wrk貌似需要自行编译, 并且在win下还不好安装. 这里可以使用[go-wrk](https://github.com/adjust/go-wrk)工具, 用`go install`就能安装了.