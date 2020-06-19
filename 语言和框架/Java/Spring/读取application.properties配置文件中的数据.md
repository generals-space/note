# 读取application.properties配置文件中的数据

参考文章

1. [Spring Boot读取properties配置文件中的数据](https://blog.csdn.net/dkbnull/article/details/81953190)
    - `@Value`注解读入到类成员
    - `environment.getProperty()`方法获取指定字段值.
    - `@ConfigurationProperties`绑定指定前缀的所有字段到一个类
