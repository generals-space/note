# pip search报错-Fault -32500：RuntimeError：PyPI's XMLRPC API is currently disabled due to unmanageable load and will be deprecated in the near future.

参考文章

1. [Getting error with pip search and pip install](https://stackoverflow.com/questions/66375972/getting-error-with-pip-search-and-pip-install)

```console
$ pip3 search s3cmd
ERROR: Exception:
Traceback (most recent call last):
  ## ...省略
  File "/usr/lib64/python3.7/xmlrpc/client.py", line 656, in close
    raise Fault(**self._stack[0])
xmlrpc.client.Fault: <Fault -32500: "RuntimeError: PyPI's XMLRPC API is currently disabled due to unmanageable load and will be deprecated in the near future. See https://status.python.org/ for more information.">
```

网络没问题, 也添加了阿里的镜像源, 不过 search 还是报上面的错误.

按照参考文章1所说, python官网`python.org`把`pip search1`的请求屏蔽了(压力太大), 所以要查找一个 package, 只能去官网的 Pypi 网页上操作.

不过`pip install`还是正常的.

------

完事之后, 我又尝试把 search 流量指到阿里云镜像源, 不过也报错了.

```console
$ pip3 search s3cmd -i http://mirrors.aliyun.com/pypi/simple/ --trusted-host mirrors.aliyun.com
ERROR: Exception:
Traceback (most recent call last):
## ...省略
pip._internal.exceptions.NetworkConnectionError: 404 Client Error: Not Found for url: http://mirrors.aliyun.com/pypi/simple/

During handling of the above exception, another exception occurred:

Traceback (most recent call last):
## ...省略
AssertionError
```

算了, 先不追究了.
