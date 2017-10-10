# session生命周期内的全局变量设置

参考文章

1. [Django session 序列化对象](https://my.oschina.net/esdn/blog/880279)

2. [Django session](http://code.ziqiangxuetang.com/django/django-session.html)

除了`application`生命周期内的变量, 就是`session`生命周期内的变量可称得上是全局变量了.

在用户登录成功的同时设置. 方法如下

```py
def login(request):
    request.session[键名] = 键值
```

`request.session.get(键名, 默认值)`: 这种方法可以在目标键可能不存在时使用, 可以设置一个默认值. 不然就是一个异常.

`del request.session[键名]`: 删除某键

`request.session`貌似可以用`keys()`和`items()`方法

Session字典中以下划线开头的key值是Django内部保留key值。 框架只会用很少的几个下划线 开头的session变量，除非你知道他们的具体含义，而且愿意跟上Django的变化，否则，最好 不要用这些下划线开头的变量，它们会让Django搅乱你的应用。

------

一般情况下, session中存储的都是简单数据, 不会存在实例对象. 但是如果有这样的需求, 则要事先在`settings`文件中(随便什么位置)写入如下配置, 否则http访问请求会报错.

```
SESSION_SERIALIZER = 'django.contrib.sessions.serializers.PickleSerializer'
```

报的错, 就是不是合法的JSON数据, 看来django默认把session对象当成普通键值字典看待的

```
    raise TypeError(repr(o) + " is not JSON serializable")
TypeError: <saltstack.saltapi.SaltAPI object at 0x7f26c81c7950> is not JSON serializable
```

> 不只是Django, Java Web中要添加实例对象类型的Session成员也是需要序列化操作的.