# linux-oom事件(Out of memory)

参考文章

1. [记一次 Linux OOM-killer 分析过程](https://pylixm.cc/posts/2018-11-28-Linux-oom-killer.html)
    - 记录了oom事件发生的过程与原理, 并提出了解决方法, 很实用

系统发生oom事件一般是因为物理内存占用较高的原因, 与虚拟内存大小没有关系.

docker容器在不指定cpu, memory限制条件的情况下默认没有限制, 容易被宿主机的oom机制kill掉.
