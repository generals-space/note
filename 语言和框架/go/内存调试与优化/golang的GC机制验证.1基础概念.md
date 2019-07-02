# golang的GC机制验证

参考文章

1. [Golang源码探索(三) GC的实现原理](https://www.cnblogs.com/zkweb/p/7880099.html)

## 内存结构

go在程序启动时会分配一块虚拟内存地址是连续的内存, 结构如下:

![](https://gitee.com/generals-space/gitimg/raw/master/9d728b7e8e69179ba70c6b47c93d6723.png)

这一块内存分为了3个区域, 在X64上大小分别是512M, 16G和512G.

### arena

> arena: 圆形运动场;圆形剧场;斗争场所;竞争舞台;活动场所

arena区域就是我们通常说的heap, go从heap分配的内存都在这个区域中.

### bitmap

bitmap区域中每2个bit对应arena区域中一个指针大小(8 byte)的内存(1:32), 用于表示

bitmap区域用于表示arena区域中哪些地址保存了对象, 并且对象中哪些地址包含了指针.

bitmap区域中每个byte(8 bit)对应了arena区域中的四个指针大小的内存, 也就是2 bit对应一个指针大小的内存.



所以bitmap区域的大小是 512GB / 指针大小(8 byte) / 4 = 16GB.

bitmap区域中的一个byte对应arena区域的四个指针大小的内存的结构如下, 每一个指针大小的内存都会有两个bit分别表示是否应该继续扫描和是否包含指针:

### spans

## 什么时候从Heap分配对象

## GC Bitmap

## Span

## Span的类型

## Span的位置
