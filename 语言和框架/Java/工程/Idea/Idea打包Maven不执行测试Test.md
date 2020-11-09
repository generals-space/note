# Idea打包Maven不执行测试Test

参考文章

1. [关于idea 中maven打包不加载测试test](https://blog.csdn.net/gty931008/article/details/86071013)

```
<-------------> 0% EXECUTING [5m 1s]
> :compileJava > Resolve files of :compileClasspath > spring-beans-5.2.8.RELEASE.jar > 642 KB/672 KB downloaded
> root project > Resolve files of :classpath > jackson-module-parameter-names-2.11.2.jar
> IDLE
> IDLE
> :compileJava > Resolve files of :compileClasspath > spring-webmvc-5.2.8.RELEASE.jar > 277 KB/934 KB downloaded
> :compileJava > Resolve files of :compileClasspath > tomcat-embed-core-9.0.37.jar > 176 KB/3.23 MB downloaded
> IDLE
```


```
mvn package -DskipTests
```
