# uwsgi编程模块

参考文章

1. [uwsgi官方文档 - The uwsgi Python module](https://uwsgi-docs.readthedocs.io/en/latest/PythonModule.html)

uWSGI服务器会自动将`uwsgi`这个模块注入到你的python工程中. 这样便于开发者在程序中动态配置uWSGI服务, 使用其内部函数, 获取指定数据(比如我是不是运行在uWSGI服务下?)