# celery问题处理

## 1. 无法启动worker(启动后2分钟内挂掉)

参考文章

1. [Unrecoverable error: WorkerLostError('Could not start worker processes',)](https://github.com/celery/celery/issues/2966)

**情境描述**

在docker容器里开了100个worker, 启动过程中就不断有worker进程报下面的错误.

```
billiard.exceptions.WorkerLostError: Could not start worker processes
```

按照参考文章1中官方的回答, 是因为分配给容器的资源太少, 以致于没有办法启动更多的worker进程.