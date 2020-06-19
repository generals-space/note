# Springboot+Idea.2.由空工程构建

参考文章

1. [官方文档 - spring boot 版本索引](https://spring.io/projects/spring-boot#learn)
2. [官方文档 - Developing Your First Spring Boot Application](https://docs.spring.io/spring-boot/docs/2.3.1.RELEASE/reference/html/getting-started.html#getting-started-first-application)

按照参考文章2, 首先创建一个工目录, 其下编写`pom.xml`, 内容如下.

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>space.generals.java</groupId>
    <artifactId>myspring</artifactId>
    <version>0.0.1-SNAPSHOT</version>

    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.3.1.RELEASE</version>
    </parent>

    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
    </dependencies>

    <description/>
    <developers>
        <developer/>
    </developers>
    <licenses>
        <license/>
    </licenses>
    <scm>
        <url/>
    </scm>
    <url/>
    
    <!-- Additional lines to be added here... -->

</project>

```

上面的 pom 中, 最主要的就是`parent`块, Maven 工程会继承 parent 块依赖的所有依赖(如果有子模块, 子模块也会继承父模块的依赖). 此时已经可以在命令行执行`maven install`安装依赖, 或是直接`maven package`进行打包了. 

然后使用 Idea 打开此工程, Idea 会自动识别其中的 pom.xml, 将其视为 maven 工程.

![](https://gitee.com/generals-space/gitimg/raw/master/7e4a43215ea385310ecccd74f6832bfb.png)

右键左侧工程目录, New -> Module..., 之后的流程就和上一篇文档相同了, 这里我们手动完成这些工作.

命令行创建一系列目录及入口文件.

```
mkdir -p src/main/java/space/generals/java
touch src/main/java/space/generals/java/MySpring.java
```

> Java对文件名及文件中的类名做了绑定, 所以文件名也要用驼峰命名法.

`MySpring.java`的内容如下

```java
package space.generals.java;

import org.springframework.boot.SpringApplication;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;

@RestController
@EnableAutoConfiguration
public class MySpring {
    @RequestMapping("/")
    public String home(){
        return "hello world";
    }
    public static void main(String[] args) {
        SpringApplication.run(MySpring.class, args);
    }
}
```

接下来要尝试运行, 目前还没有配置文件, 同时左侧工程树的目录结构与常规的 maven 项目也有所不同(`space.generals.java`在显示时应该可以用点号`.`连接起来, 不分层级).

![](https://gitee.com/generals-space/gitimg/raw/master/0a8ef4ebe1db3153dc9f1934ae186245.png)

首先, 点击右侧的`Maven`标签, 在弹出的窗口中对工程目录点击右键, 选择`Reimport`.

![](https://gitee.com/generals-space/gitimg/raw/master/9acf2789a0080249a924fdea3760eb87.png)

当所有依赖都检查完毕后, 左侧工程树就发生了变化.

![](https://gitee.com/generals-space/gitimg/raw/master/d1a4935a711f6983168c0b16d3bbec4f.png)

然后点击右上角`Add Configuration...`.

![](https://gitee.com/generals-space/gitimg/raw/master/cd6fbed049495ca6f0a801aa7766b1f5.png)

点击弹出框左侧的`+`号, 在显示的模板中找到`Spring Boot`.

![](https://gitee.com/generals-space/gitimg/raw/master/9d90c8d9bb31026ac93e9526f5e0074a.png)

> 记得把`Name`从`Unnamed`改成其他名字, 比如`MySpring`.

ok, 现在可以点击运行了, 访问`localhost:8080`可以见到"hello world"的输出, 这里就不截图了.
