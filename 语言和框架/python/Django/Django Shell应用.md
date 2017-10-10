
## 2. Django Shell命令

### 2.1 基本命令

```
##更改服务器端口号
python manage.py runserver 8080

##启动交互界面
python manage.py shell

##创建一个app，名为books
python manage.py startapp books

##验证Django数据模型代码是否有错误
python manage.py validate

##为模型产生sql代码
python manage.py sqlall books

##运行sql语句，创建模型相应的Table
python manage.py syncdb

##启动数据库的命令行工具, 与执行`mysql -u用户名 -p密码`效果一样, 可以直接连接到当前工程指定的数据库
python manage.py dbshell

manage.py sqlall books
查看books这个app下所有的表

##同步数据库,生成管理界面使用的额外的数据库表
python manage.py syncdb
```

### 2.2 使用Django内置函数操作数据库

```
$ python manage.py shell
## 其中abc为和manage.py同级的目录名称, 导入它下面的models模块, 就可以像在views.py中一样对数据库中的数据进行对象化的操作
>>> import abc.models
## 查看指定对象在数据库中有多少条数据, 其中该对象名就是在models.py文件中定义的class名, 表示在数据库中的一个表名.
>>> abc.models.对象名.object.all()
## 过滤查询
>>> PermRole.objects.filter(name = '目标名称')
```

其实除了直接操作数据库, 也可以导入django工程目录下的其他模块, 然后直接执行其中的函数.
