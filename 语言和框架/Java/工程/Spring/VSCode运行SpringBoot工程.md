# VSCode运行SpringBoot工程

参考文章

1. [VS Code开发Spring Boot应用](https://zhuanlan.zhihu.com/p/54358113)
    - 2个插件包: `Java Extension Pack`, `Spring Boot Extension Pack`
2. [VS Code打造一个完美的Springboot开发环境](https://blog.csdn.net/xiaocy66/article/details/82875770)
3. [Could not find or load main class - VS Code](https://stackoverflow.com/questions/57857855/could-not-find-or-load-main-class-vs-code)

按照参考文章1所说, 安装2个插件包(会自动安装多个插件), 然后就可以打开并运行 spring boot 工程了.

下面为maven添加了一些自定义的配置.

```json
{
    "java.configuration.maven.userSettings": "/usr/local/maven-3.6.3/conf/settings.xml",
    "maven.executable.path": "/usr/local/maven-3.6.3/bin/mvn",
}
```

我们需要打开项目中的主类文件, 然后打开左侧运行窗口中, 并点击`create a launch.json file`, ta会自动创建一个`launch.json`文件, 将打开文件中的主类填进去.

...但只是理论上, 实际上我在实践过程中在创建完成这个`launch.json`, 然后点击`Run and Debug`这个按钮时, 右下角总是弹出报错:

```log
Error: Main method not found in the file, please define the main method as: public static void main(String[] args)
```

ta认为我的主类文件中没有主类...

ok, 后来按照参考文章2中所说, 使用 vscode 重新创建了一个 spring boot 空工程, 期间出于网络问题失败了很多次, 只好设置全局代理(需要重启vscode).

当这个空工程启动成功后, 再回来打开我自己的工程, 重新创建一次`launch.json`, 就可以了.

见[generals-space/spring-boot-example](https://github.com/generals-space/spring-boot-example)

> `launch.json`是vscode对所有语言的 Run and Debug 的配置文件.

## FAQ

### Failed to resolve classpath: xxx does not exist

多模块项目中, 左侧`Run`功能区设置好主类后就出现了`Run`按钮, 但是点击`Run`的时候会报如下的错误.

![](https://gitee.com/generals-space/gitimg/raw/master/bb8e99e7e9a3edc57a6bb60b7a166358.png)

网上没找到答案...

这其实是因为图中所说的`k8s-middleware-beans`子模块下没有`.classpath`文件, 而直接打开父级工程并不会自动为所有子模块添加这个文件, 只会为父级模块和主类所在的子模块添加.

...所以需要手动用vscode打开各个子模块, 不要生成一份就想着拷贝给其他子模块, 因为每个模块生成的`.classpath`文件貌似还不太一样...

然后左侧`Explorer`功能区下面的`SPRING-BOOT-DASHBOARD`就可以看到所有的子模块了.

![](https://gitee.com/generals-space/gitimg/raw/master/0ef02865e36eca1211a8682b2083dc54.png)

现在应该可以执行了.

### 找不到主类

同一个工程, 之前还可以通过 vscode 运行, 后来做了个命令行使用 maven 打包的实验, 再打开执行的时候就说找不到主类了...

`launch.json`没变, `.classpath`文件也都还在, 很气愤.

找到参考文章3, 高票回答中, 执行`F1` -> `Clean the java language server workspace`, 重启了下 vscode, 竟然好了...

> 父级工程下不需要`.classpath`文件, 只有各子模块需要.

