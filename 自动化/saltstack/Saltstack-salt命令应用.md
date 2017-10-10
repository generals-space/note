参考文章

1. [Saltstack快速入门简单汇总](http://www.jb51.net/article/80291.htm)

指定多台minion节点, 通过`-L`(List)参数, 多个minion用逗号隔开.

```
$ salt -L 'sn192-168-176-54,sn192-168-176-55' test.ping
sn192-168-176-54:
    True
sn192-168-176-55:
    True
```