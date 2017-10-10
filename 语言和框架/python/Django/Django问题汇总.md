# Django问题汇总

## 1.

```
django.core.exceptions.ImproperlyConfigured: Error loading either pysqlite2 or sqlite3 modules (tried in that order): No module named _sqlite3
```

问题描述:

- python: 2.7.12

- pip: 8.1.2

- django: 1.10.1

在`startproject`后执行`startapp`, 报这个错误.

问题分析: 

这明显是缺少`pysqlite2`及`sqlite3`这两个库, 但是使用`pip install`这两个库文件, 却报如下的错误

解决方法:

yum安装`sqlite-devel`库, 然后重新编译`python`, 没有别的办法. 

注意安装`sqlite-devel`库, 再次`./configure`依然不需要特殊指定额外的编译选项, 因为`sqlite`已经集成到python中, 安装有`sqlite-devel`库后python对sqlite的支持会自动开启.

## 2.

```
django.db.utils.OperationalError: unable to open database file
```

问题描述:

之前执行`python manage.py syncdb`, 虽然没有使用到`sqlite`所以没有配置sqlite, 并且创建时报了警告, 但是成功创建了`db.sqlite`文件, 然后执行`python manage.py runserver 0.0.0.0:80`正常, 虽然也报了数据库的警告...

然后git提交时, ignore了这个db.sqlite文件, 直接执行`python manage.py runserver 0.0.0.0:80`就报了上面的错.

解决方法:

为了关掉`sqlite`的配置, 在工程目录的同名app目录下, 找到`settings.py`, 找到`DATABASES`字段, 将其内容注释掉即可.

```
DATABASES = {
#    'default': {
#        'ENGINE': 'django.db.backends.sqlite3',
#        'NAME': os.path.join(BASE_DIR, 'db.sqlite3'),
#    }
}
```

## 3. django连接mysql数据库错误

```
django.core.exceptions.ImproperlyConfigured: Error loading MySQLdb module: No module named MySQLdb
```

情境描述

按照官网配置django中的mysql数据库, 打开`python manage.py shell`时报此错误

问题原因

未安装mysql的python接口(或者说驱动也行)

解决办法

按照[django.core.exceptions.ImproperlyConfigured: Error loading MySQLdb module: No module named MySQLdb](http://stackoverflow.com/questions/15312732/django-core-exceptions-improperlyconfigured-error-loading-mysqldb-module-no-mo)中所说, 使用pip下载

```
sudo pip install mysql-python
```

但是pip下载失败, 后来使用apt方式下载也可

```
sudo apt-get install python-mysqldb
```

## 4. 更新数据报错

```
MultiValueDictKeyError: "'subjectDesc'"
```

情境描述

前端ajax传入参数, 更新指定行的数据, (subjectDesc是POST方法的一个字段)报这个错误.

问题原因

可能是前端没有正确传入subjectDesc参数, 或是`request.GET[]/request.POST[]`没有取到该参数, 总之这个参数在update()中的值为空, 所以会报错.

解决方法

看看前端是否正确取到该值, 是否请求中正确带有该值, 后台是否正确接收到该值.

## 5. 函数命名与django系统冲突

情境描述

```python
url(r'^acceptinvitation$', chuang_views.acceptInvitation, name='acceptinvitation'),
url(r'^rejectinvitation$', chuang_views.rejectInvitation, name='rejectinvitation'),
```

上面的路由模式将会报出下面的错误

```
AttributeError: 'module' object has no attribute 'acceptInvitation'
AttributeError: 'module' object has no attribute 'rejectInvitation'
```

问题分析

可能是由于自定义的函数名称包含了系统关键字...嗯, **连包含也不行**...

解决方法

将`acceptInvitation`重命名为`accInvitation`, `rejectInvitation`重命名为`rejInvitation`, 不再包含系统关键字即可.


## 6.

```
$ python manage.py syncdb
Error loading MySQLdb module: No module named MySQLdb
```

问题分析

通过管理脚本创建数据库时, 报上述错误.

但是当前python环境明明是已经安装了`MySQLdb`模块的. 查看`MySQLdb`源码时发现, 这个错误出现在导入`Mysqldb`模块时, 导入失败就会报这个异常, 但应该不是因为此模块不存在.

在python交互命令行中导入此模块, 报错为缺少`mysqlclient`模块, 手动安装这个模块可解决.

解决办法

安装`mysqlclient`

```
$ pip install mysqlclient
```

## 7. Django提交表单时遇到403错误：CSRF verification failed

参考文章

1. [Django提交表单时遇到403错误：CSRF verification failed](http://blog.csdn.net/cauwu/article/details/52971819)

django1.11默认开启csrf认证, `settings.py`文件中`MIDDLEWARE`块的`django.middleware.csrf.CsrfViewMiddleware`字段.

为了解决这个问题, 我们需要在form表单中加入`{% csrf_token %}`标识, 然后在页面中引入django的`csrf.js`, 其他操作都不用做.