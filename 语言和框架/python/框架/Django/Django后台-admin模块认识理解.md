# Django后台-admin模块认识理解

参考文章

1. [Django 后台 - 自强学堂](http://code.ziqiangxuetang.com/django/django-admin.html)

Django版本: 1.11.4

启动Django后访问`/admin`可以看到后台的登录界面, 如下.

![](https://gitimg.generals.space/123fc1696d8676b2ab8898d16edfe816.png)

输入用户名密码后可以看到后台管理界面. 如果此时你还不知道用户名密码, 可以在命令行通过如下命令创建超级用户.

```
$ python manage.py createsuperuser
```

后台管理界面如下

![](https://gitimg.generals.space/b6168b798550705abb879a0e491f5a3c.png)

其实这个后台只是作为超级用户管理数据表的地方. 我们知道, 在`startapp`创建一个新的app时, 它的目录结构大概像这样

```
(saltops) [root@localhost saltops]# django-admin startapp mytest
(saltops) [root@localhost saltops]# cd mytest/
(saltops) [root@localhost mytest]# tree
.
├── admin.py
├── apps.py
├── __init__.py
├── migrations
│   └── __init__.py
├── models.py
├── tests.py
└── views.py

1 directory, 7 files
```

其中`models.py`我们可以定义自己需要的数据表结构, 如下

```py
# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import models

# Create your models here.
class Text(models.Model):
    title = models.CharField(u'标题', max_length=256)
    content = models.TextField(u'内容')
 
    pub_date = models.DateTimeField(u'发表时间', auto_now_add=True, editable = True)
    update_time = models.DateTimeField(u'更新时间',auto_now=True, null=True)

```

要普通的增删改查操作的话直接`import`就够了, 但是我们还可以通过修改`admin.py`文件上这个表出现在后台管理页面中.

```py
# -*- coding: utf-8 -*-
from __future__ import unicode_literals
from django.contrib import admin
# Register your models here.
from .models import Text

admin.site.register(Text)
```

然后把我们新的app-mytest添加到`settings.py`文件里

```py
INSTALLED_APPS = [
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',
    'mytest',
]
```

由于修改了`models`, 我们需要重建表结构

```
(saltops) [root@localhost saltops]# python manage.py makemigrations
Migrations for 'mytest':
  mytest/migrations/0001_initial.py
    - Create model Text
(saltops) [root@localhost saltops]# python manage.py migrate
Operations to perform:
  Apply all migrations: admin, auth, contenttypes, mytest, sessions
Running migrations:
  Applying mytest.0001_initial... OK
```

然后再次启动django并登录, 我们在`admin.py`中注册的数据表出现在这里.

![](https://gitimg.generals.space/f3185ec3a8f5d4c48a87938523f60850.png)

点进内容详情页你还可以看到查询和删除功能.

猜测django实现了自己ORM机制中各种列类型的基本样式, 所以可以根据`models.py`中的表结构定义显示一张页面, 虽然不是特别漂亮, 但也不丑了.

