# Python异常处理[exception repr]

参考文章

1. [Python中获取异常（Exception）信息](https://www.cnblogs.com/klchang/p/4635040.html)

```py
try:
    pass
except Exception as e:
    print(e)
```

只会打印异常的 message 消息体, 没有异常信息的类型.

需要使用`print(repr(e))`.
