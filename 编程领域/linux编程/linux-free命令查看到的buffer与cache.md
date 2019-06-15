# linux-free命令查看到的buffer与cache

参考文章

1. [Linux 内存中的 Cache 真的能被回收么？](https://linux.cn/article-7310-1.html)
    - 人们对于free命令所展示数据认知的3个层次...很有意思
    - buffer与cache的含义与区别
    - cache的回收机制以及手动触发回收的方法
    - cache不可被回收的几种情况: tmpfs, shmXXX, mmap
