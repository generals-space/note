# Pypi私有镜像源搭建-pypiserver

参考文章

1. [CentOS7 搭建 python pypi 私有源 ](https://www.cnblogs.com/panhongyin/p/7065830.html)
2. [pypi 无法通过rsync进行同步，我使用rsync无法从清华站点下载pypi的目录，提示Unknown module 'pypi' #775](https://github.com/tuna/issues/issues/775)
3. [能在rsync服务中加入pypi吗](https://github.com/tuna/issues/issues/517)

听说可以用 rsync 直接从国内的 yum 镜像源搭建本地的私有源, 不用使用额外服务, 就想试试 pypi 可不可以, 于是找到了参考文章1.

不过在同步的时候出错了.

```console
$ pypi_site="rsync://rsync.mirrors.ustc.edu.cn/pypi/web/"
$ dest_dir="/root/pypi"
$ log_file="/var/log/pypi-$(date "+%Y%m%d").log"

$ /usr/bin/rsync -avrtH --delete --log-file=$log_file  $pypi_site  $dest_dir
Served by rsync-proxy (https://github.com/ustclug/rsync-proxy)

unknown module: pypi
```

可以看到, 同步的方式比 bandersnatch 简单太多了.

本来想换个其他的镜像源再试试的, 但是没找到支持 rsync 协议的源.

暂时放弃吧.
