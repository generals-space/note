# Django-定时任务Celery

参考文章

1. [结合Django+celery二次开发定时周期任务 - CSDN](http://www.cnblogs.com/huangxiaoxue/p/7266253.html)

2. [利用celery+django 在admin后台设置定时任务 - 51CTO](http://shineforever.blog.51cto.com/1429204/1737323)

~~放弃~~

~~太难而且太烦, 依赖太多服务, celery, redisg, 还要求数据库支持~~

---------------------------------------------------------------------------------------

实验环境

- django==1.11.4

- celery>=3.1.25,<4

- django-celery>=3.2.1

- celery-with-redis==3.0(这个也同时会安装上redis, 不过版本会比较老)

参考文章2中的步骤很清楚

首先`settings.py`文件中添加如下配置

```py
TIME_ZONE = 'Asia/Shanghai'
# CELERY STUFF
import djcelery
djcelery.setup_loader()
BROKER_URL = 'redis://localhost:6379'
CELERYBEAT_SCHEDULER = 'djcelery.schedulers.DatabaseScheduler' # 定时任务
CELERY_RESULT_BACKEND = 'djcelery.backends.database:DatabaseBackend'
CELERY_RESULT_BACKEND = 'redis://localhost:6379'
CELERY_ACCEPT_CONTENT = ['application/json']
CELERY_TASK_SERIALIZER = 'json'
CELERY_RESULT_SERIALIZER = 'json'
CELERY_TIMEZONE = 'Asia/Shanghai'

INSTALLED_APPS = (
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',
    'demo',
    'djcelery',
)
```

然后与`settings.py`相同目录下的`__init__.py`

```py
#! /usr/bin/env python
# coding: utf-8
 
from __future__ import absolute_import
 
# This will make sure the app is always imported when
# Django starts so that shared_task will use this app.
from .celery import app as celery_app
```

还是同级目录, 建立文件`celery.py`

```py
#! /usr/bin/env python
# coding: utf-8
 
from __future__ import absolute_import
import os
from celery import Celery
from django.conf import settings
 
# set the default Django settings module for the 'celery' program.
os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'saltops.settings')
app = Celery('saltops')
 
# Using a string here means the worker will not have to
# pickle the object when using Windows.
app.config_from_object('django.conf:settings')
app.autodiscover_tasks(lambda: settings.INSTALLED_APPS)

## 注册任务 
@app.task(bind=True)
def debug_task(self):
    print('Request: {0!r}'.format(self.request))
```

注意`debug_task`已经是注册任务, 接下来会在后台界面出现. 下面又是两个注册任务

```py
from __future__ import absolute_import
from celery import shared_task,task
 

@shared_task()
def add(x,y):
    # return x + y
    print x + y

@task(ignore_result=True,max_retries=1,default_retry_delay=10)
def just_print():
    print "Print from celery task"
```


不用`makemigrations`直接`migrate`就行.

```
(saltops) [root@localhost saltops]# python manage.py makemigrations
No changes detected
(saltops) [root@localhost saltops]# python manage.py migrate
Operations to perform:
  Apply all migrations: admin, auth, contenttypes, djcelery, sessions
Running migrations:
  Applying djcelery.0001_initial... OK
```

![](https://gitee.com/generals-space/gitimg/raw/master/af56402d6fb42485fadb83590562c168.png)

![](https://gitee.com/generals-space/gitimg/raw/master/cf25e80cdc912975b49a4ba9b015a7d8.png)
