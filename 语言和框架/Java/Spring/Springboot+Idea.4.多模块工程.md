# Springboot+Idea.4.多模块工程

参考文章

1. [IDEA创建Web项目（基于Maven多模块）](https://www.cnblogs.com/shuaishuai1993/p/9795227.html)
    - 过程详细, 有图示
    - 父子模块, 两个子模块, 其中 web 子模块实现 http 服务
    - **父级工程只用来管理依赖, 所以 Idea 自动生成的 src 目录可以移除**
2. [Maven多模块项目，无法引用另一模块的路径、文件问题](https://www.jianshu.com/p/601f7c7c7f28)

通常后端服务基本遵循MVC的结构, 分为 用于接收请求并验证的 Controller, 提供实际服务的 Service, 以及与数据库抽象的 Model 层.

前面我们创建的 hello world 工程可以说只有一个 Controller, 且是单一模块, 这里我们通过创建 Controller 和一个 Service 进一步熟悉 Maven(Spring) 的工程结构.

首先按照第一篇文章创建一个 Spring 工程, 工程名为`myproject`, 生成的`pom.xml`内容如下.

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.3.1.RELEASE</version>
        <relativePath/> <!-- lookup parent from repository -->
    </parent>
    <groupId>space.generals.java</groupId>
    <artifactId>myproject</artifactId>
    <version>0.0.1-SNAPSHOT</version>
    <name>myproject</name>
    <description>Demo project for Spring Boot</description>

    <properties>
        <java.version>1.8</java.version>
    </properties>

    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
            <exclusions>
                <exclusion>
                    <groupId>org.junit.vintage</groupId>
                    <artifactId>junit-vintage-engine</artifactId>
                </exclusion>
            </exclusions>
        </dependency>
    </dependencies>

    <build>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
            </plugin>
        </plugins>
    </build>
</project>
```

由于顶层工程的`pom.xml`只是用于管理第三方依赖, 所以将`src`和`target`目录移除.

然后在左侧的工程目录上点击右键 -> New -> Module..., 选择 Maven 工程.

![](https://gitee.com/generals-space/gitimg/raw/master/7caaed41e4c2dd72456537959b7b1edd.png)

![](https://gitee.com/generals-space/gitimg/raw/master/8f0a0aa31f5d2bfb67f906885b37dee9.png)

完成后, 根目录下的`pom.xml`会多出一个`modules`块, 子模块`mysvc`下的`pom.xml`中也会有父模块中的引用, 如下.

![](https://gitee.com/generals-space/gitimg/raw/master/ad1645d0a427670e1d9499669f58da17.png)

同样的方法创建`myctr`子模块.

此时工程结构如下, `myctr`和`mysvc`还没有源文件.

![](https://gitee.com/generals-space/gitimg/raw/master/e4467fabcd1b20e845e47f1e15259f01.png)

我们手动创建.

```
mkdir -p myctr/src/main/java/space/generals/java
touch myctr/src/main/java/space/generals/java/MyCtr.java
mkdir -p mysvc/src/main/java/space/generals/java
touch mysvc/src/main/java/space/generals/java/MySvc.java
```

我们需要在`mysvc`模块的`MySvc.java`写一个方法, 由 myctr 的`MyCtr.java`调用. 在`MySvc.java`写入以下内容.

```java
package space.generals.java;

public class MySvc {
    public String GetName(){
        return "hello general";
    }
}
```

按照参考文章1所说, `myctr`想要调用`mysvc`中的方法, 需要在自己的`pom.xml`中添加`mysvc`的依赖.

```xml
    <dependencies>
        <dependency>
            <groupId>space.generals.java</groupId>
            <artifactId>mysvc</artifactId>
            <version>0.0.1-SNAPSHOT</version>
        </dependency>
    </dependencies>
```

> `dependencies`块与`parent`同级.

然后在`MyCtr.java`中写入

```java
package space.generals.java;

import org.springframework.boot.SpringApplication;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;

import space.generals.java.MySvc;

@RestController
@EnableAutoConfiguration
public class MyCtr {
    MySvc mySvc = new MySvc();

    @RequestMapping("/")
    public String home(){
        return mySvc.GetName();
    }
    public static void main(String[] args) {
        SpringApplication.run(MyCtr.class, args);
    }
}

```

但是总是无法引入`MySvc`, 直接写入还会报错.

![](https://gitee.com/generals-space/gitimg/raw/master/a8785f76033b81a1d52587ea8a0032e6.png)

------

后来找到了参考文章3, 下面是解决方法.

Idea 对子模块间的依赖貌似并不是通过`pom.xml`中的`dependencies`实现的.

点击 File -> Project Structure...

![](https://gitee.com/generals-space/gitimg/raw/master/c462fbb384f0cccc894c407e645df47a.png)

按照上图中的指示, 选择"3. Module Dependency...", 可以看到弹出如下窗口.

![](https://gitee.com/generals-space/gitimg/raw/master/0c8dc35d2a2c2403706346fdbc1b88ef.png)

选择`mysvc`并确定, ta就会出现在`myctr`模块的`dependency`列表中, 点击确定.

![](https://gitee.com/generals-space/gitimg/raw/master/7d95f774d0929413351be2b28ed8752b.png)

然后就可以引入`MySvc`了.

![](https://gitee.com/generals-space/gitimg/raw/master/d27b14d8e102b70147d6be62d04fe457.png)

ok, 还没完, 因为目前的主类还是我们创建`myproject`父级模块的时候`src`下的主类`MyProjectApplication`, 不过我们已经把ta移除了...

现在我们要把启动类指定`MyCtr`, ta基于 spring boot 运行一个 webserver.

![](https://gitee.com/generals-space/gitimg/raw/master/7d221fa59b60813a76a267999093b50f.png)

首先选择 classpath of module.

![](https://gitee.com/generals-space/gitimg/raw/master/bf4ebb8194cdfa31252475b1d2602279.png)

然后点击`Main Class`右侧的图标, 可以找到`MyCtr`, 点击选择, 最上面的`Name`会随着变成`MyCtr`.

![](https://gitee.com/generals-space/gitimg/raw/master/c78ac9c591554b017f43fd7717111570.png)

现在可以运行了, 访问`localhost:8080`有如下结果.

![](https://gitee.com/generals-space/gitimg/raw/master/6e0ad2aac6ce6c011ac46b4167ecd9a8.png)

------

我试了试, 把`myctr`模块中`pom.xml`的`dependencies`块移除了也没有任何影响, 目前还不清楚具体的运行机制, 以后再说吧...???
