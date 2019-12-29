# etcd简单认识

参考文章

1. [etcd的初步使用](https://www.jianshu.com/p/b788c3271846)
    - 常用命令put/get/del示例, 尝试解释了目录与键值的关系(虽然感觉不怎么清晰)
2. [etcd键值存储系统的介绍和使用](https://blog.csdn.net/u010424605/article/details/44592533)
    - etcd restful接口的使用示例(curl)
3. [Ubuntu 搭建etcd](https://www.cnblogs.com/xiao987334176/articles/9942195.html)
    - etcdctl各子命令的使用示例

> etcd是一个golang编写的分布式、高可用的一致性键值存储系统，用于提供可靠的分布式键值(key-value)存储、配置共享和服务发现等功能。

本文示例中使用的是`ETCDCTL_API=2`.

**etcd与redis的区别**

1. redis中可存储众多数据类型, 如string, list, set, map等, 而etcd中只能存储简单的string;
2. redis中所有的key都存放在同一层级, 而etcd可以将key按照目录结构存储;

在etcd中存在着目录与键值两种概念, etcdctl命令中也存在get/set/rm这种对key的操作和mkdir/setdir/rmdir/ls这种对目录的操作, 一直傻傻分不清楚.

究竟是什么意思呢?

以`/`为根目录, 使用`mkdir`创建一个目录`/dir1`, 然后使用`set`创建一个kv对`/dir1/key1`-> `val1`, 那么使用`ls`查看目录结构时, 就会发现`key1`这个键位于`/dir1`这个目录下.

```bash
## 初始根目录为空, ls -r表示递归显示目录中的内容
$ etcdctl ls -r 
## 创建dir1目录
$ etcdctl mkdir /dir1
$ etcdctl ls -r
/dir1
## 创建key1键
$ etcdctl set /dir1/key1 val1
val1
## 查看dir1下的内容
$ etcdctl ls /dir1
/dir1/key1
$ etcdctl ls -r
/dir1
/dir1/key1
```

可以说

- key则可以理解为文件系统中的"文件", value则是文件内容;
- 目录其实是key-value的集合, 不可以直接存储value;

需要注意的是, 不管是dir还是key, 在etcd中的存储都是有序的. 继续上面的实验

```bash
$ etcdctl mkdir /dir0
$ etcdctl mkdir /dir1/key0 val0
$ etcdctl ls -r
/dir1
/dir1/key1
/dir1/key0
/dir0
```

可以看到, 如果按照字母排序, `dir0`应该在`dir1`的前面, `key0`也应该在`key1`的前面, 实际上并不是这样.
