# Centos7搭建NFS

参考文章

1. [centos7下NFS使用与配置](https://www.cnblogs.com/jkko123/p/6361476.html)

NFS服务端需要安装两个包: `rpcbind`, `nfs-utils`.

NFS客户端需要安装一个包: `nfs-utils`.

服务端在启动服务前需要指定共享目录, 及读写控制.

客户端使用`mount`挂载, 只有在安装`nfs-utils`后, `mount -t`才有nfs选项.

## 服务端配置

默认安装好`nfs-utils`后, 配置文件`/etc/exports`为空. 尝试写入如下配置

```
/mnt/nfsfold 192.168.1.* (rw,sync,no_all_squash)
/mnt/nfsfold 192.168.2.* (ro,sync,no_all_squash)
```

`rw`和`ro`, `sync`和`async`的区别很容易, `XXX_squash`应该是与客户端挂载后创建文件/目录的属主有关, 日后再研究.

------

2019-03-13更新

`XXX_squash`的4个配置项应该是基于为该目录设置了`rw`权限的情况. NFS服务在安装时就拥有了一个名为`nfsnobody`的系统用户和用户组.

ta们的区别应该在于创建新文件/目录时的属主为`nfsnobody`或是只有uid没有用户名的形式.