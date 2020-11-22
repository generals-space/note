# Java jvm 参数 -Xms -Xmx -Xss

参考文章

1. [java jvm 参数 -Xms -Xmx -Xmn -Xss 调优总结](https://www.cnblogs.com/jpfss/p/8618297.html)
2. [java的Xmx是设置什么的？](https://www.cnblogs.com/yjf512/p/7770968.html)
    - 使用`java -X`可以看到`-X`系列的参数
3. [JAVA_OPTS参数-Xms和-Xmx的作用](https://www.cnblogs.com/zxp_9527/archive/2008/12/24/1361911.html)
    - `java.lang.Runtime`类`freeMemory()`, `totalMemory()`, `maxMemory()`

使用`java -X`可以看到`-X`系列的参数.

```console
$ java -X
    ## 省略
    -Xms<size>        设置初始 Java 堆大小
    -Xmx<size>        设置最大 Java 堆大小
    -Xss<size>        设置 Java 线程堆栈大小
    ## 省略
```

`Xmx`和`Xms`是相对应的, 一个是`memory max(Xmx)`, 代表程序最大可以从操作系统中获取的内存数量, 一个是`memory start(Xms)`, 代表程序启动的时候从操作系统中获取的内存数量. 

比如`java -cp . -Xms80m -Xmx256m`指定此程序启动的时候使用`80m`的内存, 最多可以从操作系统中获取`256m`的内存. 

> 与 k8s 中`requests`与`limits`的概念很像.

