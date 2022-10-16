# logging日志

参考文章

1. [Python打印log，包括行号，路径，方法名，文件](https://blog.csdn.net/wzhg0508/article/details/9364885)
2. [python logging详解及自动添加上下文信息](https://www.cnblogs.com/xybaby/p/9197032.html)
    - `logging`模块是python官方提供的、标准的日志模块, 借鉴了许多`log4j`中的概念和思想
    - Logger对象不是通过实例化`logging.Logger`类而来的, 都是通过`logging.getLogger(name)`获得一个与`name`关联的`logger`对象, `logging`内部会维护这个映射关系, 用同样的`name`反复调用`logging.getLogger`, 事实上返回的是同一个对象.
    - 一个logger可以包含多个`handler`, 多个`Filter`
    - `lazy logging`的概念
    - `logging`为了线程安全, 每个Handler都会持有一个互斥锁, 每次写日志的时候都会获取锁, 写完之后再释放锁(`logging.thread = None`)
    - `LoggerAdapter`自动输出上下文, `extra`参数添加自定义字段
3. [Python 日志logging模块初探及多线程踩坑(1)](https://blog.csdn.net/qq_41603102/article/details/89705421)
    - 字段的固定宽度, 对齐方式
    - 多线程多logger对象, 输出到同一文件日志丢失的问题
4. [Can't get this custom logging adapter example to work](https://stackoverflow.com/questions/39467271/cant-get-this-custom-logging-adapter-example-to-work)
    - `LoggerAdapter`的使用示例

```py
logging_config = {
    'level': logging.INFO,
    'format': '%(asctime)s %(levelname)s - %(filename)s:%(lineno)d - %(message)s',
}
logging.basicConfig(**logging_config)
logger = logging.getLogger(__name__)

```

其他文件引入`logger`即可.

`logging.basicConfig`有点像全局变量.

`logging.getLogger`得到的`logger`对象其实有点像单例, 可以在多个线程内使用, 不用加锁.

------

如果我们希望给`logging.getLogger()`传入不同的参数呢?

```py
logger = logging.getLogger(__name__)
logger.setLevel(logging.INFO)
handler = logging.FileHandler("log.log")
# handler.setLevel(logging.INFO)
logger_format = '%(asctime)s %(levelname)s - %(filename)s:%(lineno)d - %(message)s'
formatter = logging.Formatter(logger_format)
handler.setFormatter(fmt=formatter)
```
