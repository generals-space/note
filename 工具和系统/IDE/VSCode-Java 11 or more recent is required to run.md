# VSCode-Java 11 or more recent is required to run

参考文章

1. [VSCode showing “Java 11 or more recent is required to run. Please download and install a recent JDK”](https://stackoverflow.com/questions/63043585/vscode-showing-java-11-or-more-recent-is-required-to-run-please-download-and-i)
2. [JDK Requirements](https://github.com/redhat-developer/vscode-java/wiki/JDK-Requirements#setting-the-jdk)

前不久才刚刚会用 vscode 运行 java/maven/spring 项目, 今天突然就不行了, 点击`Run`按钮, 右下角就弹出如下提示框.

![](https://gitee.com/generals-space/gitimg/raw/master/7202d54c563aa5ca12e0d747d52b91fa.png)

```
Java 11 or more recent is required to run. Please download and install a recent JDK
```

按照参考文章2中的说法, vscode 的 Java 支持拓展了 eclipse 的工具, 而2020年开始, Oracle 规则 Java 11 为最低要求的版本, 所以就会不停地弹这个窗口.

但也不必真的切换到 Java 11, 参考文章1的采纳答案给出了解决方法.

首先要下载 Java 11, 下载完了解压(Java 8 是用 dmg 安装器安装的, 11 则有了 tar 包), 放到一个目录下就行了, 也不用配置什么环境变量, 因为也不打算用ta.

之后修改 vscode 的 settings.json 文件, 按照参考文章1中提问者给出的示例, 和采纳答案的回答, 有如下配置.

```json
    "java.configuration.updateBuildConfiguration": "disabled",
    "java.home": "/usr/local/jdk-11.0.8/Contents/Home",
    "java.configuration.runtimes": [
        {
            "name": "JavaSE-1.8",
            "path": "/Library/Java/JavaVirtualMachines/jdk1.8.0_251.jdk/Contents/Home",
            "default": true
        },
        {
            "name": "JavaSE-11",
            "path": "/usr/local/jdk-11.0.8/Contents/Home"
        }
    ],
```

`java.configuration.runtimes`中配置两个版本的 Java, 把 Java 8 作为默认选项, 然后添加`java.home`, 也指向 Java 11 的路径就可以了, 应该是用来"欺骗" Java 插件的.
