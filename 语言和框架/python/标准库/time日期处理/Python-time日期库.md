# Python-time日期库

参考文章

1. [Python中time模块与datetime模块在使用中的不同之处](http://www.jb51.net/article/75364.htm)
2. [Python测量时间，用time.time还是time.clock](https://www.cnblogs.com/limengjie0104/archive/2018/05/06/8997466.html)

常用方法及解释

## 得到当前时间戳

`time.time()`: 当前时间戳, 单位为秒, 如`1526197995.321712`(js 里单位是毫秒, golang/java 都是秒级, 但是 python 能精确到小数点后6位, 所以还是 python 更精确)

```py
>>> time.time()
1526198052.62613
```

## 时间戳转`struct_time`对象: GMT时间与本地时间

1. `time.gmtime([secs])`: 将指定时间戳转换为`struct_time`日期对象. 如果不指定时间则默认转换当前的时间.
2. `time.localtime([secs])`: 与`gmtime()`用法类似, 不过`gmtime()`返回的总是`GMT`时区的时间, 而`localtime()`返回的是本地时区对应的时间.

```py
>>> time.gmtime()
time.struct_time(tm_year=2018, tm_mon=5, tm_mday=13, tm_hour=8, tm_min=10, tm_sec=48, tm_wday=6, tm_yday=133, tm_isdst=0)
>>> timestamp = time.time()
>>> time.gmtime(timestamp)
time.struct_time(tm_year=2018, tm_mon=5, tm_mday=13, tm_hour=8, tm_min=15, tm_sec=11, tm_wday=6, tm_yday=133, tm_isdst=0)
```

可以对这个`struct_time`对象获取年月日等值. 如下

```
>>> timeobj = time.gmtime()
>>> dir(timeobj)
['__add__', '__class__', '__contains__', '__delattr__', '__doc__', '__eq__', '__format__', '__ge__', '__getattribute__', '__getitem__', '__getslice__', '__gt__', '__hash__', '__init__', '__le__', '__len__', '__lt__', '__mul__', '__ne__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__rmul__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', 'n_fields', 'n_sequence_fields', 'n_unnamed_fields', 'tm_hour', 'tm_isdst', 'tm_mday', 'tm_min', 'tm_mon', 'tm_sec', 'tm_wday', 'tm_yday', 'tm_year']
>>> timeobj.tm_hour
8
```

- `timeobj.tm_year` 年
- `timeobj.tm_mon`  月
- `timeobj.tm_mday` 日
- `timeobj.tm_wday` 星期(0-6, 星期一是0)
- `timeobj.tm_yday` ...这个好像是指在一年中的第几天, 比如1月1日是第1天
- `timeobj.tm_hour` 时
- `timeobj.tm_min`  分
- `timeobj.tm_sec`  秒

## `struct_time`对象转换成时间戳

`time.mktime(t)`: 把目标`struct_time`对象转换成时间戳.

```
>>> timestamp = time.time()
>>> timestamp
1526200304.663764
>>> 
>>> timeobj = time.localtime(timestamp)
>>> new_timestamp = time.mktime(timeobj)
>>> new_timestamp
1526200304.0
```

不过看来小数点被省略了...

> 注意: `mktime()`所接受的`struct_time`对象必须为localtime, 否则会不准确.

```
>>> timestamp = time.time()
>>> timeobj = time.gmtime(timestamp)
>>> new_timestamp = time.mktime(timeobj)
>>> timestamp
1526200445.373669
>>> new_timestamp
1526171645.0
```

## 字符串与`struct_time`对象的相互转换

1. `time.strftime(format[, t])`: 把目标`struct_time`对象`t`转换成指定格式的字符串形式.

格式参数是必须的, `t`如果不指定的话默认使用当前时间.

1. `time.strptime(string[, format])`: 与`strftime`相反, 将指定指定格式的日期字符串转换成`struct_time`对象. 默认格式为`%a %b %d %H:%M:%S %Y`.

**时间对象转格式化字符串**

```
>>> timestamp = time.time()
>>> timeobj = time.localtime(timestamp)
time.strftime('%Y-%m-%d %H:%M:%S +0000', timeobj)
'2018-05-13 16:44:23 +0000'
```

同样注意`localtime`还是`gmtime`.

**字符串转时间对象**

```py
>>> time.strptime('2018-05-13 16:44:23', '%Y-%m-%d %H:%M:%S')
time.struct_time(tm_year=2018, tm_mon=5, tm_mday=13, tm_hour=16, tm_min=44, tm_sec=23, tm_wday=6, tm_yday=133, tm_isdst=-1)
```
------

## 时间戳直接转字符串(一般不会用到)

1. `time.ctime([secs])`: 同样可以接受时间戳作为参数, 不过转换的结果是字符串形式...但不能指定格式, 应该不会有太大用.
2. `time.asctime([t])`: 与`ctime()`用法类似, 也是转换成字符串, 不过接受的参数类型为`struct_time`对象.

----

## time计量时间: 用`time()`还是`clock()`?

先看如下一个示例

```py
>>> time.time()
1526202652.913809
>>> time.sleep(1)
>>> time.time()
1526202661.347759
```

```py
>>> time.clock()
0.173621
>>> time.sleep(1)
>>> time.clock()
0.174706
```

解释: `time.time()`计算的是绝对时间, 而`time.clock()`表示的是CPU耗费在当前线程中的时间.

我们可以看到, `time.time()`的示例中, 两个时间戳相差了将近10秒, 这是因为在输入过程中耗费了时间. 在`time.clock()`示例中, 你会发现两个时间点相差了还不到1毫秒...这是因为在`sleep`命令行空等时, CPU转而执行其他任务去了, 实际花费在**运算**上的时间很小.