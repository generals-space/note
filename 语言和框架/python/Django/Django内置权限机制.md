# Django内置权限机制

参考文章

1. [django权限管理 - 博客园](http://www.cnblogs.com/huangxm/p/5770735.html)

2. [python django模型内部类meta详细解释 - 博客园](http://www.cnblogs.com/lcchuguo/p/4754485.html)

3. [Django models中的meta选项 - 博客园](http://www.cnblogs.com/ccorz/p/Django-models-zhong-demeta-xuan-xiang.html)

4. [Django权限机制的实现 - 伯乐在线](http://python.jobbole.com/84086/)

## 1. 内置权限机制解读

### 1.1 可选权限

Django内置一个简单权限机制, 对每一个表对象, 都有**增加**, **修改**, **删除**3种权限.

作为超级用户, 拥有对所有(已注册的)表对象的增删改权限.

为了证明上述结论, 我们以超级用户身份登录后台管理系统, 打开`User`表, 查看超级用户本身(这里称为root)的权限. 如下

![](https://gitimg.generals.space/00412ee5eb350d5894887e3749cf6e38.png)

图中所示, 以竖线`|`分隔的分别是`app名`, `表名`, `权限名`(关于这3者的关系, 看一下Django后台的笔记就可以理解).

```
mytest | text | Can add text
mytest | text | Can change text
mytest | text | Can delete text
```

这3个默认权限我们没有显式定义过, 可以认为是所有Model对象的父级对象`models.Model`附带的.

> Three basic permissions -- add, change and delete -- are automatically created for each Django model.   -- 来自$DJANGO/contrib/auth/models.py中的Permisson类

可以说, 每个permission都是`$DJANGO.contrib.auth.Permission`类的实例.

### 1.2 权限使用

仔细想想, 可选的权限其实并不是单纯附加在用户对象或是已注册表对象上的, 而单独存在, 并分配到用户或用户组对象上的.

比如, 对于`用户`表, 我们可以对不同的用户分配不同的增删改权限. 当我们想要创建自定义表`Paper`, 需要配置增删改查的权限呢? 尤其是查阅的权限.

首先我们要能在Django系统权限表中增加这4种(针对Paper表的)权限, 然后可以在后台为指定用户分配权限, 至于对不同权限的把控与验证的问题, 另说.

现在我们看看django默认的权限表, 在数据库中是如何存储的吧. 以默认sqlite数据库示例

```
$ sqlite3 db.sqlite3
sqlite> .tables
auth_group                  django_admin_log          
auth_group_permissions      django_content_type       
auth_permission             django_migrations         
auth_user                   django_session            
auth_user_groups            minion_text               
auth_user_user_permissions  mytest_text               
sqlite> .header on
sqlite> select * from auth_permission;
id|content_type_id|codename|name
1|1|add_logentry|Can add log entry
2|1|change_logentry|Can change log entry
3|1|delete_logentry|Can delete log entry
4|2|add_group|Can add group
5|2|change_group|Can change group
6|2|delete_group|Can delete group
7|3|add_permission|Can add permission
8|3|change_permission|Can change permission
9|3|delete_permission|Can delete permission
...省略
```

`add_`, `change_`, `delete_`是前缀, 每在admin模块注册一个model, 都会在这张表里创建3个权限行.

接下来我们要尝试自定义我们的权限.

### 1.3 自定义权限

自定义权限的目的, 还是要为了更好的管理我们的自定义model, 所以我们需要自建新的model, 这里以`Paper`表为例, 依然在`mytest`应用下创建.

`models.py`文件

```py
class Paper(models.Model):
    title = models.CharField(u'标题', max_length=256)
    content = models.TextField(u'内容')
    class Meta:
        verbose_name = 'Paper权限列表'
        verbose_name_plural = verbose_name
        ## 权限信息，这里定义的权限的名字，后面是描述信息，描述信息是在django admin中显示权限用的
        permissions = (
            ("view_paper", "查看Paper信息"),
            ("view_paper_detail", "查看Paper详细信息"),
        )
```

django内置的`Meta`类, 可以在这里定义权限, 就是`permissions`字段.

`admin.py`文件

```py
# -*- coding: utf-8 -*-
from __future__ import unicode_literals
from django.contrib import admin

# Register your models here.
from .models import Paper

admin.site.register(Paper)
```

然后写入数据库

```
python manage.py makemigrations
Migrations for 'mytest':
  mytest/migrations/0002_auto_20170910_0250.py
    - Create model Paper
python manage.py migrate
Operations to perform:
  Apply all migrations: admin, auth, contenttypes, mytest, sessions
Running migrations:
  Applying mytest.0002_auto_20170910_0250... OK
```

登录后台, 查看权限列表

![](https://gitimg.generals.space/459d4561505ba9d7ec39c459cefdfbba.png)

这次再看, 以竖线分隔的其实是`app名`, 表中Meta元类的verbose_name字段值, 和permission的元组值.

### 1.4 自定义权限使用

默认存在的增删改权限是为了方便后台模块超级用户直接分配表级权限的, 也就是说, 这种权限是给后台模块用的. 但是实际中我们要的权限控制应该是粒度更细, 更灵活. 

我们需要用户表中存储有每个用户对象的权限列表, 然后在进行操作时, 查询它是否拥有相配的权限.

User对象的user_permission字段管理用户的权限

```py
myuser.user_permissions = [permission_list]
myuser.user_permissions.add(permission, permission, ...) #增加权限
myuser.user_permissions.remove(permission, permission, ...) #删除权限
myuser.user_permissions.clear() #清空权限
myuser.has_perm('myapp.fix_car')                    ## 检查是否拥有目标权限, 注意前缀的存在
##############################################################
# 注：上面的permission为django.contrib.auth.Permission类型的实例
##############################################################
```

