# "etcdctl del" clear all my data with the key range

参考文章

1. [ETCD:etcdctl](https://www.codenong.com/p11934614/)
2. [官方仓库readme etcd-io/etcd/etcdctl](https://github.com/etcd-io/etcd/tree/master/etcdctl)
3. [官方文档 etcd3 API](https://etcd.io/docs/v3.3.12/learning/api/)

etcdctl version: 3.2.24

## 

最开始怀疑是由于上面的`put`操作引起了类似 redis **缓存雪崩**的情况, 因为etcd容器的时区与宿主机节点相关8小时, 写入一个新于已有数据的键, 导致原有的键瞬间过期失效而被删除.

```
GET [options] <key> [range_end]
DEL [options] <key> [range_end]
```

将`put`改为`del`时并没有那么谨慎, 后面跟了一个串字符串也没怎么在意. 因为最开始我并不知道`range`这个概念的存在, 另外, etcd 作为一个键值存储库, 理论上各种操作都是单键的, 就和 redis 一样, 批量删除是需要做额外的选项甚至要使用脚本完成.

```bash
export ETCDCTL_API=3
export ETCDCTL_ENDPOINTS=https://127.0.0.1:2379
export ETCDCTL_CACERT=/etc/kubernetes/pki/etcd/ca.crt
export ETCDCTL_CERT=/etc/kubernetes/pki/etcd/server.crt
export ETCDCTL_KEY=/etc/kubernetes/pki/etcd/server.key
```

```
etcdctl put foo bar     # OK
etcdctl put foo1 bar1   # OK
etcdctl put foo2 bar2   # OK
etcdctl put foo3 bar3   # OK
```

单键查询

```console
[root@k8s-master-01 ~]# etcdctl get foo1
foo1
bar1
```

`--from-key KEY`, 获取从目标`KEY`开始, 之后的所有键.

```
[root@k8s-master-01 ~]# etcdctl get --from-key foo1
foo1
bar1
foo2
bar2
foo3
bar3
```

如果`KEY`为空'', 可以获取etcd中的所有键.

然后就是`range_end`所表示的结尾键了.

```
[root@k8s-master-01 ~]# etcdctl get foo1 foo3
foo1
bar1
foo2
bar2
```

> `range_end`是闭区间, 不包含在结果中.

参考文章3中讲解的太生硬, 不太容易理解, 其实本质上就是一个字典排序的问题. 在etcd中, 数据的存储顺序并不与插入顺序相关, 而是全部将key以字典序排列(与python3中的`dict`相似.), 遍历时会以字典序输出.

在这样的前提下, `range_end`才有意义.

那么ta与前缀匹配有什么不同呢? 下面以kuber在 etcd 中存储的部分数据为例

```
/registry/ranges/serviceips
/registry/ranges/servicenodeports
/registry/secrets/kube-system/service-account-controller-token-kmblf
/registry/secrets/kube-system/service-controller-token-lmpbh
/registry/serviceaccounts/default/default
/registry/serviceaccounts/kube-node-lease/default
/registry/serviceaccounts/kube-public/default
/registry/serviceaccounts/kube-system/attachdetach-controller
```

如果我想删除(或查询)`ranges`和`secrets` ~~目录下(抱歉`ETCD3`没有目录的概念了)~~ 为前缀的键, 如果用前缀匹配, 需要执行两次, 如果有多个与ta们平级的前缀单词, 操作会更多.

在上面key为字典序的条件下, 我们能确定`range_end`为`serviceaccounts`, 那么只要删除`[/registry/ranges, /registry/serviceaccounts)`区域内的key就可以了.

```
etcd del /registry/ranges /registry/serviceaccounts
```

参考文章3中用了加法来表示`key`与`range_end`的关系, 像`[a, a+1) (e.g., [‘a’, ‘a\x00’) looks up ‘a’) `, 又或者`“aa"+1 == “ab”, “a\xff"+1 == “b”`, 简直离谱... 其实就像这样啦.

![](https://gitee.com/generals-space/gitimg/raw/master/7E069E0635F84D1BD4634B00F56440BF.png)
