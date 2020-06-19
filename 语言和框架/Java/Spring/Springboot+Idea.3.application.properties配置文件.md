# Springboot+Idea.3.application.properties配置文件

参考文章

1. [SpringBoot项目修改访问端口和访问路径](https://blog.csdn.net/qq_40087415/article/details/82497668)
    - `src/main/resources/application.properties`与`src/main/resources/application.yml`
    - 修改端口与访问路径(类似`/api`的全局前缀)
    - 修改访问路径的字段`context-path`应该修改为`server.servlet.context-path`

本文基于前两篇文章所建的工程.

## 1. 端口与访问前缀

在不创建任何额外的文件时, 运行一个最基本的 spring boot 工程所使用的端口为 `:8080`, 我们从修改端口, 访问uri前缀来学习Java中配置文件的使用方法.

![](https://gitee.com/generals-space/gitimg/raw/master/95bbd3fccc850f6810d2f04e5ec8b2d5.png)

Idea 很智能, 猜到我们可能希望创建`resources`目录.

![](https://gitee.com/generals-space/gitimg/raw/master/84e92fc440c6f6b4308e69af1a42bc63.png)

然后再在该目录下创建`application.properties`, 内容如下

```
server.address = 127.0.0.1
server.port = 8090
server.servlet.context-path = /api
```

重启服务就可生效.

## 2. dev, pro环境选择

需要一个键

```
spring.profiles.active = pro
```

然后`application.properties`同目录再创建一个`application-pro.properties`即可. 同理, 如果`active`选择为`dev`, 也需要创建一个`application-dev.properties`文件.

