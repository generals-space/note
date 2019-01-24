# celery数据库连接(二)-每个worker进程维护独立的数据库连接

参考文章

1. [Celery Worker Database Connection Pooling - ThatAintWorking的回答](https://stackoverflow.com/questions/14526249/celery-worker-database-connection-pooling)

2. [celery官方文档 signal](http://docs.celeryproject.org/en/latest/userguide/signals.html#worker-process-init)

3. [celery.signals.worker_process_init.connect示例集合](https://programtalk.com/python-examples/celery.signals.worker_process_init.connect/)

> 注意: 本文所提到的方案都是针对`prefork`并发模型来说的, 在使用`eventlet`或`gevent`时会出错.

```py
#!/usr/bin/env python3
from celery import Celery
from celery.signals import worker_process_init, worker_process_shutdown
import psycopg2
import time

from .config import CeleryConfig, logging_config, db_config

class SingletonObject():
    db_conn = None
    def connect_db(self):
        retry_times = 0
        while retry_times <= 3:
            try:
                self.db_conn = psycopg2.connect(**db_config)
                msg = 'connect to database success'
                print(msg)
                return self.db_conn
            except Exception as e:
                msg = 'connect to database failed %d times: %s' % (retry_times + 1, e)
                print(msg)
                time.sleep(3)
                retry_times += 1
        return None
    def close_db(self):
        self.db_conn.close()
        self.db_conn = None

    def reconnect_db(self):
        self.connect_db()

singleton_object = SingletonObject()

@worker_process_init.connect
def celery_worker_init(signal, sender, **kwargs):
    print('worker init...')
    global db_conn
    db_conn = singleton_object.connect_db()
    if db_conn == None: raise ValueError('数据库连接异常, worker启动失败')

@worker_process_shutdown.connect
def celery_worker_shutdown(sender, signal, pid, exitcode, **kwargs):
    global db_conn
    db_conn.close()
    print('worker shutdown...')
```

然后在具体的任务模块中, 可以导入`singleton_object`.

```py
@celery_app.task(base = BaseTask, name = 'get_book_info')
def get_book_info(url):
    print('get_book_info: %s' % url)
    db_conn = singleton_object.db_conn
    try:
        db_curr = db_conn.cursor()
    except psycopg2.InterfaceError as e:
        print(e)
        singleton_object.reconnect_db()
    finally:
        db_curr = db_conn.cursor()
```

在worker启动时, 每个进程会先执行`@worker_process_init.connect`装饰器下的方法, 初始化单例对象中的数据库连接. 然后task对象引入这个单例, 就可实现在同一个worker的生命周期中, 所有task都使用相同的数据库连接对象, 避免了频繁连接数据库造成的性能浪费.