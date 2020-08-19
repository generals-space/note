# SpringBoot Post请求体大小限制

参考文章

1. [Java SpringBoot post请求大小限制](https://blog.csdn.net/t1014336028/article/details/81675621)
2. [spring boot 设置tomcat post参数限制](https://www.ancii.com/ammbvqm4/)
    - `tomcat`的`server.xml`配置, 添加`maxPostSize`字段, 默认`1M`.

springboot工程编译成jar包, 内置的web容器默认是tomcat, 要修改post请求体的大小限制, 需要在`application.properties`文件中添加`server.tomcat.max-http-post-size=-1`.

网上也有说`spring.http.multipart.max-file-size=-1`, 这个应该是上传文件的请求对文件大小的限制, 两者是有区别的.
