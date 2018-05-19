# 在gunicorn环境下, 日志中看不代码使用print输出的调试信息

参考文章

1. [gunicorn配置文件官方文档](http://docs.gunicorn.org/en/latest/settings.html)

2. [Is there a way to log python print statements in gunicorn?](https://stackoverflow.com/questions/27687867/is-there-a-way-to-log-python-print-statements-in-gunicorn)

3. [python中stdout输出不缓存的设置方法](http://www.jb51.net/article/50508.htm)

使用`uwsgi`部署flask工程时, 代码中用`print`可以在日志中输出信息.

但用`gunicorn`部署时, 使用`print`打印的信息只会在进程停止后才会输出到日志文件. 配置如下

```py
loglevel    = 'info'
accesslog   = '/var/log/gunicorn.log'
errorlog    = '/var/log/gunicorn.log'
capture_output = True
```

按照参数文章2中的例子, 设置`PYTHONUNBUFFERED`这个环境变量.

这个环境变量可以实现`python -u 目标文件`中`-u`选项同样的功能(可以通过`man python`中查看更详细的信息), 不能在`gunicorn`的配置文件中设置, 而要在其执行前设置.

只要这个环境变量为非空字符串, 即可生效.

该变量的作用可以查看参考文章3.