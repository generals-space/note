# 压测工具Locust

参考文章

1. [Locust: 简介和基本用法](https://www.cnblogs.com/imyalost/p/9758189.html)
    locust的简单使用
2. [Locust官方文档 making-http-requests](https://docs.locust.io/en/stable/writing-a-locustfile.html#making-http-requests)

> locust: 蝗虫

使用locust要先创建测试脚本, 然后使用`locust`命令启动web界面, 在界面中点击运行, 请求数据会显示在web界面中展示的表格中.

`locust.py`

```py
from locust import HttpLocust, TaskSet, task

## 定义接口测试任务组, 继承TaskSet类
## 成员方法用@task装饰, 可以发起对指定路径的请求.
## 请求的客户端为self.client, 类似于requests库的使用方法.
class MyAPI(TaskSet):
    ## 钩子函数
    def on_start(self):
        print('start benchmark...')
    def on_stop(self):
        print('benchmark complete...')

    @task
    def login(self):
        json_data = {
            'id': 1,
            'name': 'general',
        }
        ##
        resp = self.client.post('/login', json = json_data)
        if resp.status_code == 200:
            ## 在这里打印没有意义
            ## print(resp.json)
            resp.success()
        else:
            resp.failure("Got wrong response")

## 这个是加载入口
class MainTester(HttpLocust):
    task_set = MyAPI
    host = 'http://localhost:8080'
    min_wait = 3000  # 单位为毫秒
    max_wait = 6000  # 单位为毫秒

```

启动

```
$ locust -f .\locust.py
```

默认会创建一个`http://localhost:8089`的web接口. 访问这个地址, 可以打开web界面.

输入如下两个值

- Number of users to simulate: 设置模拟的用户总数
- Hatch rate (users spawned/second): 每秒启动的虚拟用户数

Start swarming: 执行locust脚本.
