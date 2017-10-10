# Django内置登录机制

参考文章

1. [Django用户认证系统　authentication system](http://blog.csdn.net/feelang/article/details/24992693)

2. [Django URL name详解](http://code.ziqiangxuetang.com/django/django-url-name.html)

环境: django 1.11.4

本文的目的, 是使用django自带的登录模块完成认证过程, 即自定义登录页, 自定义登录成功后的跳转页, 但是用户数据是直接用django内置的admin数据库进行认证的.

## 1. django内置登录流程

首先, 默认`/admin`页的路由配置如下

```py
from django.conf.urls import url
from django.contrib import admin

urlpatterns = [
    url(r'^admin/', admin.site.urls)
]
```

我们到django包所在路径去查看`/admin`路由的处理方法, 因为我们访问这个路径就会被重定向到一个登录页`/admin/login/`, 先到`$DJANGO_HOME/contrib/admin/...`

![](https://gitimg.generals.space/54e181029ee23ae664a10fa3ef1ad5b9.png)

好吧, 没找到`admin.site.urls`的路径...根本就不对嘛. 

我们先不考虑这个黑盒, 按照参考文章1的提示, 可以确定该路由的实际处理方法其实是在`$DJANGO_HOME/contrib/auth/views.py`的`login`函数. (同目录下的`urls.py`中也刚好有`/login`的路由映射).

`login`的函数原型如下(总之用到的参数就是下面这些, 只是经过了层层包装, 这里不深究)

```py
class LoginView(SuccessURLAllowedHostsMixin, FormView):
    """
    Displays the login form and handles the login action.
    """
    form_class = AuthenticationForm
    authentication_form = None
    redirect_field_name = REDIRECT_FIELD_NAME
    template_name = 'registration/login.html'
    redirect_authenticated_user = False
    extra_context = None
```

其中, 参数`template_name`为登录页的模板(可以`find`一下这个`login.html`在哪)

另一个比较有用的参数为`redirect_field_name`, 它的值默认为变量`REDIRECT_FIELD_NAME`, `grep`一下可以发现这个值为`next`, 字符串类型.

```
[root@localhost django]# grep -r 'REDIRECT_FIELD_NAME' ./*
...省略
./contrib/auth/__init__.py:REDIRECT_FIELD_NAME = 'next'
...省略
```

这个`next`其实是一个变量名, 在访问一个被`@login_required`装饰器修饰的url时会被强行重定向到登录页, 在地址栏中可以看到有一个`next`参数保留着之前要访问的url路径, 当登录成功后会再跳转回去. 类似如下

```
http://172.32.100.232:9000/login?next=/dashboard/
```

## 2. 自定义登录页

好了, 现在我们自定义登录url及模板页面.

```py
from django.contrib.auth.views import login
...
    url(r'^login$', login, kwargs = {'template_name': 'login.html'}, name = 'login')
...
```

其中`kwargs`是传入login方法的参数, 普通对象类型.

`login.html`可以放在`templates`目录下的合适地方, 自定义, 它的内容可以为

```html
<form method="post" action="{% url 'login' %}">
    {% csrf_token %}  
    <fieldset>
        <label class="block clearfix">
            {{ form.username }}
        </label>
        <label class="block clearfix">
            {{ form.password }}
        </label>

        <button type="submit" class="btn btn-sm btn-primary btn-block">
            登录
        </button>
        <input type="hidden" name="next" value="{{ next }}" /> 
        <div class="space-4"></div>
        {% if form.errors %}  
            <p>Your username and password didn't match. Please try again.</p>  
        {% endif %}  
    </fieldset>
</form>
```

`{{ form.username }}`和`{{form.password}}`会被渲染为对应的`input`元素, 当`登录`按钮按下时, 提交表单. 

提交的地址正是`form`元素`action`属性所指向的url, 即名为`login`的url. 

什么? 它怎么知道名为`login`是哪个? 

在`urls.py`中我们不是指定了login页的url映射吗? 其中不正好有一个`name`参数? 就是这个啦. 关于django的url name讲解, 去看看参数文章2吧.

这样就基本完成了我们的定制.

访问一个被`@login_required`修改的url, 会跳转到我们的登录页, 输入正确的用户名和密码, 也能正常跳转.

再自由一点, 不如把`{{ form.username }}`和`{{form.password}}`都写成实际的html元素, 因为其实只有input元素的name属性是有作用的, 我们只要知道为它们命名成合适的名称就行了, 这样我们就可以自定义`input`框样式了.

------

不过貌似还有一个问题, 直接访问`login`页时, `next`参数是没有值的, 这种情况下完成登录会默认被重定向到`/accounts/profile/`, 也许我们还要再做一个重定向, 把`/accounts/profile/`重定向我们自己的主页.

但是仍然没办法添加诸如验证码, 记住我等功能, 需要重写django逻辑, 暂时不考虑.

## 3. 注销

```py
url(r'^logout$', logout, kwargs = {'template_name': 'login.html'}, name = 'logout'),
```

django里的注销还有一个单独的注销页, 不太符合我们平时的习惯...注销后直接跳转到登录页嘛.

所以好一点解决方法是, 把注销按钮的链接写为`/logout?next=/dashboard`, 这样跳转到登录页后重新登录又可以回到主页了.