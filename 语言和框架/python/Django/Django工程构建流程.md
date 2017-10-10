# Django工程构建流程

参考文章

1. [Django 基本命令](http://code.ziqiangxuetang.com/django/django-basic.html)

## 1. 创建工程目录脚手架

```
$ django-admin startproject saltops
```

这个简单框架中存在一个简单的app, 我将其称为`模块`, 其名称也为`saltops`

```
(saltops) [root@localhost saltops]# tree
.
├── manage.py
└── saltops
    ├── __init__.py
    ├── settings.py
    ├── urls.py
    ├── wsgi.py

```

下面的命令可以创建一个与`saltops`同级的模块.

```
$ django-admin startapp 新模块名称
```

可以看到, 工程根目录下只有一个`manage.py`和各个模块, 所以在使用`python manage.py runserver 0.0.0.0:8000`启动时, 实际上执行的是项目同名模块`saltops`目录下的`uwsgi.py`, 它才是实际的工程入口.

## 2. 数据库的配置与表创建

如果有创建数据库的需要, 可以在各个模块目录下`models.py`文件中定义你自己的表对象. 然后执行如下代码

```
Django 1.7.1及以上 用以下命令
# 1. 创建更改的文件
$ python manage.py makemigrations
# 2. 将生成的py文件应用到数据库
$ python manage.py migrate


旧版本的Django 1.6及以下用
$ python manage.py syncdb
```

当然, 这需要事先创建好数据库配置, 默认使用

> 我一直觉得, makemigrations就是查看工程目录的表结构变动情况, 然后创建有效的SQL语句, 有些类似于源码编译的configure过程; migrate则是实际写数据库操作.

## 4. 模板

首先要在模板文件开头添加`{% load staticfiles %}`标记