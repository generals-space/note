# celery+redis存储内容

参考文章

1. [Celery + Redis 的探究](https://www.jianshu.com/p/52552c075bc0)

celery在以redis为broker时, 会在redis中建立几个键, 记录队列, 任务, worker信息. 一般会有如下几种

```
127.0.0.1:6379> keys *
1) "_kombu.binding.celery.pidbox"
2) "celery"
3) "unacked_mutex"
4) "unacked"
5) "_kombu.binding.celery"
6) "_kombu.binding.celeryev"
7) "unacked_index"
```

1. `celery`: 任务队列, list类型, 存储**待执行**的任务对象.

2. `_kombu.binding.celery.pidbox`: worker集合(不管是否在线), set类型.

3. `unacked`: 已经分发给某个worker的任务集合, hash类型. 空闲worker从celery队列中取得任务后放入这个集合, 等worker执行完毕后删除, 其field的值正是celery列表中的成员.

4. `unacked_index`: 如其名称, ta存储的是`unacked`的key的索引, `zset`类型(有序集合). ta表示被worker取走的任务id集合, 且有先后顺序, 可以通过ta的成员为key, 从`unacked`映射中取得相应的任务对象.

5. `unacked_mutex`: string类型, 我也不知道是干啥的.

6. `_kombu.binding.celeryev`: 貌似存的是worker的事件对象? 虽然不太明白什么是事件对象.