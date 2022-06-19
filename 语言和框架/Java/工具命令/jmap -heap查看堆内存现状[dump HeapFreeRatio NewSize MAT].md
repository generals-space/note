# jmap -heap查看堆内存现状[dump HeapFreeRatio NewSize]

参考文章

1. [JVM：jmap heap 堆参数分析MinHeapFreeRatio、MaxHeapFreeRatio、MaxHeapSize、NewSize、MaxNewSize](https://blog.csdn.net/claram/article/details/104635114)
    - 单纯介绍`jmap -heap $pid`的输出选项.
    - `jmap -heap $pid`的输出结果示例
    - jdk9之前: `jmap -heap $pid`
    - jdk9之后: `jhsdb jmap --heap --pid $pid`
2. [JVM：jmap heap 堆参数分析MinHeapFreeRatio、MaxHeapFreeRatio、MaxHeapSize、NewSize、MaxNewSize](https://www.cnblogs.com/zouhong/p/14773110.html)
    - 同参考文章1(估计是同一作者)
3. [【JVM】jmap命令详解----查看JVM内存使用详情](https://www.cnblogs.com/sxdcgaq8080/p/11089664.html)
    - `-dump`生成dump文件.
4. [问题排查-jmap和MAT工具](https://blog.51cto.com/janephp/2432551)

## 查看进程的详细内存占用, 包括每个区域大小和使用大小

```
jmap -heap pid
```

## 查看存活对象大小, 单位是字节

```
jmap -histo:live 33320
```

输出可能有点多, 不方便分析...
