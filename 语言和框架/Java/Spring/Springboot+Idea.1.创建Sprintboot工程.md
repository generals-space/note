# Springboot+Idea.1.创建Sprintboot工程

参考文章

1. [官方文档 - spring boot 版本索引](https://spring.io/projects/spring-boot#learn)
2. [官方文档 - Developing Your First Spring Boot Application](https://docs.spring.io/spring-boot/docs/2.3.1.RELEASE/reference/html/getting-started.html#getting-started-first-application)
3. [使用IDEA创建一个springboot项目](https://www.cnblogs.com/little-rain/p/11063967.html)
4. [Spring Boot : Whitelabel Error Page](https://blog.csdn.net/qq_37495786/article/details/82464294)

File -> New -> Project 

![](https://gitee.com/generals-space/gitimg/raw/master/195ae006444c9c93371ed3fc0cbc8da2.png)

窗口左侧列表中选择`SpringInitializr`, Service Url 使用默认, 点击 Next.

![](https://gitee.com/generals-space/gitimg/raw/master/0188ccfc10f01c48a12910cf1154d779.png)

有时候`start.spring.io`可能无法连接, 可以多试几次.

![](https://gitee.com/generals-space/gitimg/raw/master/cf88e0f7875213a7f58e0840726aa089.png)

成功连接后出现如下界面. 与创建普通的 maven 工程很像.

![](https://gitee.com/generals-space/gitimg/raw/master/248104036b9da072c35822537d2c8a7f.png)

按照自己的习惯填写工程信息(就像`npm init`时要填的信息).

![](https://gitee.com/generals-space/gitimg/raw/master/2765de8a3ee764acae90bf4aac49bf85.png)

> maven/gradle 的关系就类似于当初 govendor/godep 的关系, 或者说, npm/webpack 的关系.

然后选择依赖, 这里我们要建一个本地的 web 工程, 本来按照参考文章3要选 starter 组件的, 但是好像没有, 这里直接勾选`Spring Web`.

![](https://gitee.com/generals-space/gitimg/raw/master/d3cb7fe83c2cbbd994b021eb638fd057.png)

完成

![](https://gitee.com/generals-space/gitimg/raw/master/040471514f20f6504d62664bb8811c66.png)

最终生成的目标结构和入口代码如下.

![](https://gitee.com/generals-space/gitimg/raw/master/5f621b0d02d0c9d67ceeae02fe40fe19.png)

Idea会自动下载相关的依赖, 下载完成后, 右上角的运行按照处, 会出现可用的配置文件.

![](https://gitee.com/generals-space/gitimg/raw/master/7fa54f6fcfb458120c301d5cf107c2fc.png)

目前入口程序中还一个接口都没有, 这里我们加一下.

```java
package space.generals.java.myspring;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.RequestMapping;

@SpringBootApplication
public class MyspringApplication {
    @RequestMapping("/")
    String home(){
        return "hello world";
    }
    public static void main(String[] args) {
        SpringApplication.run(MyspringApplication.class, args);
    }
}

```

点击运行, 如下

![](https://gitee.com/generals-space/gitimg/raw/master/5c3daa41f49a7804348c5a8384d4f006.png)

...失败了, 404

![](https://gitee.com/generals-space/gitimg/raw/master/83dbd5910f96262b1d404c38800614f1.png)

网上找了找, 有人说是因为入口程序位置不正确引起的, 目录层级多了一层(比如参考文章4). 

右键入口程序`MyspringApplication` -> `Refactor` -> `Move Class...`, 将ta放到`space.generals.java`包中, 重新启动, 但是没用...

找了学Java小伙伴问了问, 人家一眼看出原因 -- 要在`MyspringApplication`加上`@RestController`注解(与`@RequestMapping`在同一个包内).

```java
@SpringBootApplication
@RestController
public class MyspringApplication {
    // ...
}
```

再次重启并访问, 有结果了.

![](https://gitee.com/generals-space/gitimg/raw/master/836f9f37ea546002427d1d6183a839cd.png)

...`@RestController`表示某个类将作为 controller 提供 uri 接口的服务, 但是 idea 在构建好的工程中, 将主类`MyspringApplication`作为`run()`的参数传入, 怎么想都应该自动为ta添加上`@RestController`注解才对吧...???
