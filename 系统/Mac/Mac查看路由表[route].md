# Mac查看路由表[route]

参考文章

1. [Mac OS X 查看路由表](http://wp.panwenbin.com/2013/04/08/mac-os-x-%E6%9F%A5%E7%9C%8B%E8%B7%AF%E7%94%B1%E8%A1%A8/)

想用`route`查一下当前网络的网关的, 结果怎么用都不正确, 看了`man`手册也是一脸萌...

```console
$ route get
route: writing to routing socket: Invalid argument
$ route -n get
route: writing to routing socket: Invalid argument
$ route -n get -net
route: writing to routing socket: Invalid argument
You have new mail in /var/mail/general
$ route -n get -hosts
route: bad keyword: hosts
usage: route [-dnqtv] command [[modifiers] args]
$ route -n get -host
route: writing to routing socket: Invalid argument
```

后来找到参考文章1, 可以使用`netstat -r`来代替`route`.

```console
$ netstat -r
Routing tables

Internet:
Destination        Gateway            Flags        Netif Expire
default            192.168.2.1        UGSc           en0
10.37.129/24       link#15            UC           vnic1      !
10.211.55/24       link#14            UC           vnic0      !
127                localhost          UCS            lo0
localhost          localhost          UH             lo0
...
```

显示得极慢, 跟windows有一拼, 差劲...不好总好过没有, 反正也不经常用.
