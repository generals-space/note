# jstat查看gc耗时

参考文章

1. [jstat查看gc情况](https://blog.csdn.net/wangshuminjava/article/details/107041189)
2. [ jstat命令查看jvm的GC情况 （以Linux为例）](https://www.cnblogs.com/zhangfengshi/p/11342212.html)

## -gcutil

```
jstat -gcutil PID INTERVAL
```

- INTERVAL: 采样间隔, 单位为毫秒

```console
$ jstat -gcutil 13 1000 
S0    S1    E     O    M     CCS   YGC  YGCT   FGC FGCT  CGC CGCT GCT 
37.54 0.00  33.52 1.66 81.71 75.42 2330 53.596 7   2.062 -   -    55.658 
0.00  29.76 45.69 1.66 81.71 75.42 2331 53.617 7   2.062 -   -    55.680 
35.93 0.00  57.84 1.66 81.71 75.42 2332 53.642 7   2.062 -   -    55.704
```

- `YGC`: yong gc(年轻代gc)的次数
- `YGCT`: yong gc time(年轻代gc)所花费的时间, 单位为秒
- `FGC`: full gc(全局gc, 将会引发STW)
- `FGCT`: full gc time
- `GCT`: 其值为`YGCT+FGCT`的和

上面的5个字段都是单调递增. 

以`YGC`为例, 2秒之内其值从2330增加到了2332, 说明在这2秒内, 又进行了2次 yong gc, 基本上每秒1次. 

同时, `YGCT`也在增加, 在这2秒之内, 新增的2次gc行为分别耗时0.021s和0.025s.

------

关于 full gc

由于 java 的 gc 特性, yong gc 的确是比较频繁的, 而 full gc 就不那么频繁了(有些jvm调优手段是尽量减少full gc频率的).

full gc 并不是定时执行的, 而是 jvm heap 堆内存到达一定阈值时自动触发的.
