# Python-datetime日期库

参考文章

1. [Python中time模块与datetime模块在使用中的不同之处](http://www.jb51.net/article/75364.htm)

## 1. datetime对象

使用`datetime.datetime.now()`可以获得当前时刻的`datetime.datetime`实例。 对于一个 `datetime.datetime`实例，主要会有以下属性及常用方法，看名称就能理解，应该没有太大问题：

### 实例属性

- `datetime.year`

- `datetime.month`

- `datetime.day`

- `datetime.hour`

- `datetime.minute`

- `datetime.second`

- `datetime.microsecond`

- `datetime.tzinfo`

### 实例方法

- datetime.date() # 返回 date 对象(只有datetime的date(年月日)部分)

- datetime.time() # 返回 time 对象(只有时分秒部分)

- datetime.replace(name = value) # 前面所述各项属性是 read-only 的，需要此方法才可更改

- dattime.strftime(format)  # 将datetime对象转换成指定格式的字符串.

- datetime.timetuple() # 返回time.struct_time 对象

```py
>>> from datetime import datetime
>>> now = datetime.now()
>>> now
datetime.datetime(2018, 5, 13, 17, 37, 21, 805738)
>>> now.date()
datetime.date(2018, 5, 13)
>>> now.time()
datetime.time(17, 37, 21, 805738)
>>> now.strftime('%Y-%m-%d %H:%M:%S')
'2018-05-13 17:37:21'
```

## datetime类方法

除了实例本身具有的方法,类本身也提供了很多好用的方法：

datetime.today()  # 返回当前时间, localtime, 和datetime.now(tz = None)相同.
datetime.now([tz]) # 返回当前时间, 默认localtime
datetime.utcnow()  # UTC 时间
datetime.fromtimestamp(timestamp[, tz]) # 由 Unix Timestamp 构建datetime对象
datetime.strptime(date_string, format)  # 给定时间格式解析字符串构建datetime对象

...貌似没有datetime对象直接得到对应时间戳的方法, 不过倒是有从时间戳到datetime的方法.

`datetime`对象要得到时间戳, 需要调用`timetuple`先转成`struct_time`对象, 然后再用`time`包中的`mktime`方法得到时间戳.

## timedelta时间间隔

两个`datetime`对象直接相减就能获得一个`timedelta`对象.

`datetime`对象 +/- `timedelta`对象可以得到一个新的`datetime`对象.

两个`timedelta`对象相加减可以得到一个新的`datetime`对象.

来看看如何手动创建一个`timedelta`对象.

`datetime.timedelta(days=0, seconds=0, microseconds=0, milliseconds=0, minutes=0, hours=0, weeks=0)`

`timedelta`实例拥有如下属性

- `days`: 包含几天(不够一天的话为0)

- `seconds`: 按秒计算的总量

```
>>> a = timedelta(hours=25)
>>> a
datetime.timedelta(1, 3600)
>>> print(a)
1 day, 1:00:00
>>> a.days
1
>>> a.seconds
3600
```

