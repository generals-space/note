参考文章

1. [pip download 只下载不安装命令的使用方法](https://blog.csdn.net/qq_27565769/article/details/84262760)


pip freeze: 查看当前 pip 环境已安装的依赖列表
pip install -r requirements.txt: 安装列表中的依赖
pip download requests==2.24.0: 只下载不安装
pip download -r requirements.txt: 安装列表中的依赖

## 

按照参考文章1所说, 使用`download`子命令下载依赖时, 报如下错误(明明`install`就执行得好好的)

```console
$ pip3 download -r ./require.txt
Looking in indexes: http://mirrors.aliyun.com/pypi/simple/
WARNING: The repository located at mirrors.aliyun.com is not a trusted or secure host and is being ignored. If this repository is available via HTTPS we recommend you use HTTPS instead, otherwise you may silence this warning and allow it anyway with '--trusted-host mirrors.aliyun.com'.
ERROR: Could not find a version that satisfies the requirement certifi==2020.6.20 (from -r ./require.txt (line 1)) (from versions: none)
ERROR: No matching distribution found for certifi==2020.6.20 (from -r ./require.txt (line 1))
WARNING: You are using pip version 20.1.1; however, version 20.2.2 is available.
You should consider upgrading via the '/Users/general/Public/pyvenv/xunjian/bin/python3 -m pip install --upgrade pip' command.
```

本来以为是使用方法有问题, 结果按照`WARNING`提示, 添加上`--trusted-host`参数就好了.

```
pip3 download -r ./require.txt --trusted-host mirrors.aliyun.com
```
