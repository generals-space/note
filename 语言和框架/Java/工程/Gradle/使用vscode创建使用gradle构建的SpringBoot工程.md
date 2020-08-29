# 使用vscode创建使用gradle构建的SpringBoot工程

安装了`Spring Boot Extension Pack`后, `Ctrl`+`Shift`+`P`, 调出`Spring Initializr`, 会有如下两个选项

- Spring Initializr: Generate a Gradle Project
- Spring Initializr: Generate a Maven Project

使用`Generate a Gradle Project`, 即可创建一个使用`Gradle`构建的 SpringBoot 工程, 初始结构如下

![](https://gitee.com/generals-space/gitimg/raw/master/557d8b52169842695516f637f70550de.png)

然后添加`.vscode/launch.json`文件, 可以正常启动.

也可以在命令行使用`gradle build`进行打包, 先修改`build.gradle`

```
repositories {
	maven { url 'http://maven.aliyun.com/nexus/content/groups/public'}
	// mavenCentral()
}
```

然后执行`gradle build`(需要事件下载`gradle`包)

```console
$ gradle build

Welcome to Gradle 6.6!

Here are the highlights of this release:
 - Experimental build configuration caching
 - Built-in conventions for handling credentials
 - Java compilation supports --release flag

For more details see https://docs.gradle.org/6.6/release-notes.html

Starting a Gradle Daemon (subsequent builds will be faster)

> Task :test
2020-08-26 14:22:04.973  INFO 52271 --- [extShutdownHook] o.s.s.concurrent.ThreadPoolTaskExecutor  : Shutting down ExecutorService 'applicationTaskExecutor'

BUILD SUCCESSFUL in 21s
5 actionable tasks: 5 executed
```

可以在`build/libs`目录下生成相关的`jar`包, 可以启动.

