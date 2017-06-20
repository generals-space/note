# Python包错误

## 1. 

```
  File "/usr/lib/python2.6/site-packages/pkg_resources.py", line 546, in resolve
    raise DistributionNotFound(req)
pkg_resources.DistributionNotFound: meld3>=0.6.5
```

情景描述

使用`pip`安装`supervisor`, 执行`supervisord -h`时报上述错误(实际上, `supervisor`附带的3个命令`supervisord`, `supervisorctl`和`echo_supervisord_conf`都会报这个错误).

```
$ pip freeze | grep meld3
meld3==1.0.2
```

原因分析与解决办法

貌似是因为pip形式安装的`meld3`有问题, 在python官网的pypi上下载[meld3-1.0.2.tar.gz](https://pypi.python.org/packages/45/a0/317c6422b26c12fe0161e936fc35f36552069ba8e6f7ecbd99bbffe32a5f/meld3-1.0.2.tar.gz#md5=3ccc78cd79cffd63a751ad7684c02c91), 然后pip安装这个本地包能够解决.

或许可以尝试使用`yum`安装这个包`python-meld3`.