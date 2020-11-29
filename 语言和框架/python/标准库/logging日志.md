# logging日志

参考文章

1. [Python打印log，包括行号，路径，方法名，文件](https://blog.csdn.net/wzhg0508/article/details/9364885)

```py
logging_config = {
    'level': logging.INFO,
    'format': '%(asctime)s %(levelname)s - %(filename)s:%(lineno)d - %(message)s',
}
logging.basicConfig(**logging_config)
logger = logging.getLogger(__name__)

```

其他文件引入`logger`即可.

