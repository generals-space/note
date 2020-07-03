参考文章

1. [SpringBoot多模块项目实践（Multi-Module）](https://segmentfault.com/a/1190000011367492)
    - 有打包操作的过程示例
2. [IDEA 打jar包（IDEA自带的打包方式）](https://www.jianshu.com/p/97250dc28508)
    - 第二种方式打包成功, `META-INF`在启动类的`src`目录下, 与`java`目录同级.
3. [-jar参数运行应用时classpath的设置方法](https://www.cnblogs.com/aggavara/archive/2012/11/16/2773246.html)
    - `-Xbootclasspath/a`选项, `/a`应该表示`append`, 同理`-Xbootclasspath/p`类似于`previous`, 就像在写`$PATH`环境变量时的前后顺序会影响搜索顺序一样.
    - `-Xbootclasspath`将会完全重写搜索路径, 基本不会用到.
4. [【SpringBoot】项目打成 jar 包后关于配置文件的外部化配置](https://segmentfault.com/a/1190000015413754)
    - 工程打为`jar`包后如何设置参数, 给出了两种方法.
        1. `java -jar xxx.jar --server.port=10000`
        2. 在`xxx.jar`所在目录下建立`config`目录, 并在`config`目录下创建`application.properties`文件

Java工程中的打包, 应该与 golang 一样有两种. 一种是常规的`go build`, 创建工程的启动包; 另一种, 像通过`import github.com/vishvananda/netlink`这种直接引用的(当然还是可以通过`go build`进行编译, 不过因为没有 main 函数, 没有可执行的文件生成), 类似于linux中的共享库`.so`或是windows的动态链接库`.dll`.

Java中也是这两个模式, 但是前者又分 Jar 包和 War 包, 这两个的区别, 可以类比于 python 的 django. django 本身只是一个 web 开发框架, 但ta内置了一个 http server, 业务代码编写完成可以直接发布(虽然性能不怎么样), 作为启动包的 Jar 包就是这种. 还有就是需要部署额外的运行环境, 对于 python 是`uwsgi`, 或是`gunicorn`, 对于 Java 而言则是 tomcat.

Jar 包和 War 包可有适用场景, 前者适合于容器环境, 微服务等, 各工程独立部署运行, 互不干扰. 后者可以在一个运行环境中部署多个工程, 集中运维.

注意: 这里讲的打包(启动包)是指 Java Web, 尤其指像 Spring 这种工程, 最终是需要通过`java -jar XXX.jar`来执行的. 如果是普通的Java工程(像我之前大学写的Java爬虫), 可以见我关于Mave的几篇文章, 不过那里需要通过`mvn`启动, 不适用于生产环境(生产环境为什么要部署maven).

单一模块工程打包实验: [spring-boot-example/009.single-module-jar](https://github.com/generals-space/spring-boot-example/tree/master/009.single-module-jar)

多模块工程打包实验: [spring-boot-example/010.multi-module-jar](https://github.com/generals-space/spring-boot-example/tree/master/010.multi-module-jar)
