# Maven镜像仓库配置

参考文章

1. [Maven 镜像](https://developer.aliyun.com/mirror/maven)

阿里云镜像站中给出的配置有点不太一样, 下面是我正在用的, 至于区别, 以后再说吧.

```xml
    <mirror>
      <id>nexus-aliyun</id>
      <mirrorOf>*</mirrorOf>
      <name>Nexus aliyun</name>
      <url>https://maven.aliyun.com/nexus/content/groups/public</url>
    </mirror>
```
