# Django静态文件404

参考文章

1. [Managing static files (e.g. images, JavaScript, CSS)](https://docs.djangoproject.com/en/1.11/howto/static-files/)

2. [django静态文件处理](http://www.cnblogs.com/jcli/archive/2011/09/18/2180504.html)

django: 1.11.4

在开发和测试环境, 通过django内置的`runserver`启动, 网页和静态文件都可以访问, 但在生产环境, 网页`template`可以访问, 但是静态文件`static`都404了.

在`settings.py`中相关配置如下.

```py
import os
BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
## 静态文件路径配置
STATICFILES_DIRS = (BASE_DIR + '/static',)
STATIC_URL = '/static/'
``

> `static`和`template`在工程根目录, 由各个app共享的.

之前在网上找了很多示例, 官方教程中有说要在`urls.py`里导入`static`的, 如下

```py
from django.conf import settings
from django.conf.urls.static import static

urlpatterns = [
    # ... the rest of your URLconf goes here ...
] + static(settings.STATIC_URL, document_root=settings.STATIC_ROOT)
```

但是依然不行.

经过对比, 两者的区别只在于`DEBUG`的取值不同, 开发和测试环境的`DEBUG = True`, 而生产环境的`DEBUG = False`. 参考文章2中解释了`DEBUG`选项的作用.

> 在`DEBUG`为true时我们只需要建立`static`目录后, 把静态资源放进去就可以访问. 在`DEBUG`为`False`时需要我们手动指定静态资源目录, 并配置映射关系. 在正式环境下建议不采用django处理静态资源文件, 这样对应用服务器压力较大, 也不好做cdn. 可以用nginx, apache部署静态资源.

简单来说, `DEBUG = False`时django就不再负责静态文件的访问请求了, 因为性能不够. 但它毕竟也能当做一个http服务器, 要想让它继续有效, 应该可以自行定义的. 参考文章2中给出了一种方法, 但已经不再适用, 这个问题以后再讨论.