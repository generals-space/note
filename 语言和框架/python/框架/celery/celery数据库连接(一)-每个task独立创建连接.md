# celery数据库连接(一)-每个task独立创建连接

参考文章

1. [Using a database in a celery task](https://groups.google.com/forum/#!topic/celery-users/kBHq73y1YME)

本来我在任务目录下的`__init__.py`中定义了一个`conn_db()`函数, 返回`db_conn`和`db_curr`, 并且自行调用, 所有task都使用相同的`db_conn`和`db_curr`. 但是在执行时出现了如下问题.

```
DatabaseError: (psycopg2.DatabaseError) error with status PGRES_TUPLES_OK and no message from the libpq
```

搜索了很久, 没有找到有效的解决办法. 原因应该是在celery的多进程worker模型下使用同一个`db_conn`, 出现了类似死锁一样的情况. 

我尝试过保留全局的`db_conn`, 在不同的task函数内获取`db_curr`变量, 但是没有用.

目前先在不同task里手动建立数据库连接, 但是这样感觉会很低效, 频繁连接数据库对数据库本身也是不小的压力.

等以后能找到celery的跨task的全局变量再加以改进吧.

按照参考文章1的提示, 并且改进了一下连接代码, 基本如下.

```py
#!/usr/bin/env python3
from celery import Celery
from celery.task import Task

import psycopg2
import time

from . import db_config

class BaseTask(Task):
    _db = None

    def _connect_db(self):
        retry_times = 0
        while retry_times <= 3:
            try:
                db = psycopg2.connect(**db_config)
                return db
            except Exception as e:
                print('connect to database failed %d times: %s' % (retry_times + 1, e))
                time.sleep(3)
                retry_times += 1
        return None

    ## 没有把连接数据库的操作放在构造函数中, 因为不是所有task都要连接数据库
    @property
    def db(self):
        if self._db is not None: return self._db
        self._db = self._connect_db()
        return self._db

    '''
    after_return: 钩子函数, 其他钩子可以查阅官方文档
    '''
    def after_return(self, status, retval, task_id, args, kwargs, einfo):
        self.db.close()
        self._db = None

```

然后在常规任务函数中按照如下形式调用

```py
@celery_app.task(base = BaseTask, name = 'get_book_info')
def get_book_info(url):
    ## self貌似是可以为task的同名变量?
    self = get_book_info
    db_curr = self.db.cursor()
```
