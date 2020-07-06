# "etcdctl del" clear all my data with the key range

参考文章

1. [ETCD:etcdctl](https://www.codenong.com/p11934614/)
2. [官方文档 etcd-io/etcd/etcdctl](https://github.com/etcd-io/etcd/tree/master/etcdctl)

etcdctl version: 3.2.24

## 

最开始怀疑是由于上面的`put`操作引起了类似 redis **缓存雪崩**的情况, 因为etcd容器的时区与宿主机节点相关8小时, 写入一个新于已有数据的键, 导致原有的键瞬间过期失效而被删除.

```
GET [options] <key> [range_end]
DEL [options] <key> [range_end]
```

将`put`改为`del`时并没有那么谨慎, 后面跟了一个串字符串也没怎么在意. 因为最开始我并不知道`range`这个概念的存在, 另外, etcd 作为一个键值存储库, 理论上各种操作都是单键的, 就和 redis 一样, 批量删除是需要做额外的选项甚至要使用脚本完成.

```
./etcdctl put foo bar
# OK
./etcdctl put foo1 bar1
# OK
./etcdctl put foo2 bar2
# OK
./etcdctl put foo3 bar3
# OK
```

