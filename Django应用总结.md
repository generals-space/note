# Django应用总结

## 1. 模板转义

在使用artTemplate.js, 以ajax方式将响应渲染到页面时, 需要在html文件中输出`{{`与`}}`, 然而django在生成html响应时会自动对其html模板中的`{%`与`%}` 进行变量替换. 

而类似下面这样的错误, 是因为`django`在遇到`{{`的第一个`{`就开始进行变量的赋值解析操作了, 但又找不到与之匹配的`%`, 所以会出错.

```
{{each list as value }}    #出错行

TemplateSyntaxError at /posts
Could not parse the remainder: ' list as value' from 'each list as value'
```

解决方法

使用`{% templatetag openvariable %}`与`{% templatetag closevariable %}`标签, 将分别输出`{{`与`}}`到html而不是把'templatetag'等当作变量来替换. 于是`{% templatetag openvariable %} each list as value {% templatetag closevariable %}`将会被django正常地输出为`{{each list as value }}`到静态html文件, 之后就可以被`artTemplate`编译渲染了.

另外, 其他的渲染模板是根据`{%`与`%}`作为变量标签的, 鉴于这些情况, django还提供了其他对应的输出方案:

```
openvariable {{
closevariable }}
openblock {%
closeblock %}
openbrace {
closebrace }
opencomment {#
closecomment #}
```

类比上面的示例, 不难得到`templatetag`的正确使用方法.

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

## 3. Django获取前端checkbox值

多选情况下, 前端代码为

```html
<form action="" method="POST">
    <input type="checkbox" value="1" name="check_box_list"/>苹果
    <input type="checkbox" value="2" name="check_box_list"/>梨
    <input type="checkbox" value="3" name="check_box_list"/>杏
    <input type="checkbox" value="4" name="check_box_list"/>桃子
    <input type="submit" value="提交">
</form>
```

后端django代码

```python
if request.method == 'POST':
    ## check_box_list是一个list对象, 选中的值会出现在里面, 如[2, 4]
    check_box_list = request.REQUEST.getlist('notuseldap')
```