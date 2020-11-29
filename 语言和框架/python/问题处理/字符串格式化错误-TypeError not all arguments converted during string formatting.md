# 字符串格式化错误-TypeError not all arguments converted during string formatting

参考文章

1. [Python学习笔记：TypeError: not all arguments converted during string formatting](https://blog.csdn.net/lvsehaiyang1993/article/details/80909984)
    - 还挺全的

我的代码是

```py
logger.info('Init webdriver complete, desired_capabilities: ', json.dumps(self.caps))
```

其中`self.caps`为一个字段类型, 结果报了如下错误

```
--- Logging error ---
Traceback (most recent call last):
  File "/Library/Frameworks/Python.framework/Versions/3.7/lib/python3.7/logging/__init__.py", line 1025, in emit
    msg = self.format(record)
  File "/Library/Frameworks/Python.framework/Versions/3.7/lib/python3.7/logging/__init__.py", line 869, in format
    return fmt.format(record)
  File "/Library/Frameworks/Python.framework/Versions/3.7/lib/python3.7/logging/__init__.py", line 608, in format
    record.message = record.getMessage()
  File "/Library/Frameworks/Python.framework/Versions/3.7/lib/python3.7/logging/__init__.py", line 369, in getMessage
    msg = msg % self.args
TypeError: not all arguments converted during string formatting
```

找到参考文章1时第一眼就看到的"低级错误", 就意识到了问题所在...把逗号`,`改成百分号`%`就可以了, 当然还要加个`%s`

```
logger.info('Init webdriver complete, desired_capabilities: %s' % json.dumps(self.caps))
```
