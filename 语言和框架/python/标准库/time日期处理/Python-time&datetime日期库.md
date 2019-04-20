# Python-time&datetime的区别和联系

参考文章

1. [Python中time模块与datetime模块在使用中的不同之处](http://www.jb51.net/article/75364.htm)

Python 中提供了对时间日期的多种多样的处理方式，主要是在 time 和 datetime 这两个模块里。今天稍微梳理一下这两个模块在使用上的一些区别和联系。

## 1. time

在 Python 文档里，time是归类在`Generic Operating System Services`中，换句话说， 它提供的功能是更加接近于操作系统层面的。通读文档可知，time 模块是围绕着 Unix Timestamp 进行的。

该模块主要包括一个类`struct_time`，另外其他几个函数及相关常量。 需要注意的是在该模块中的大多数函数是调用了所在平台C library的同名函数， 所以要特别注意有些函数是**平台相关**的，可能会在不同的平台有不同的效果。另外一点是，由于是基于Unix Timestamp，所以其所能表述的日期范围被限定在`1970 - 2038`之间，如果你写的代码需要处理在前面所述范围之外的日期，那可能需要考虑使用datetime模块更好。

> Unix Timestamp是从1970年开始计时的秒数, 32位系统中最大的整数值所表示的秒数只能计数到2038年, 而且在mysql中, 时间戳类型的字段所占空间为4个字节, 也是如此.

## 2. datetime

`datetime`比`time`高级了不少，可以理解为`datetime`基于`time`进行了封装，提供了更多实用的函数。在datetime 模块中包含了几个类，具体关系如下:

- `timedelta`  # 主要用于计算时间跨度
- `tzinfo`         # 时区相关
- `time`          # 只关注时间
- `date`          # 只关注日期
- `datetime`  # 同时有时间和日期

名称比较绕口，在实际实用中，用得比较多的是`datetime.datetime`和`datetime.timedelta`，另外两个 datetime.date 和 datetime.time 实际使用和 datetime.datetime 并无太大差别。

