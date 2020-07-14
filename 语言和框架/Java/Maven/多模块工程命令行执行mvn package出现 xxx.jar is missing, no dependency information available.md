# 多模块工程命令行执行mvn package出现 xxx.jar is missing, no dependency information available

参考文章

1. [The POM for xxxx is missing, no dependency information available](https://blog.csdn.net/u013067420/article/details/53200361)

场景描述

多模块工程, 在命令行, 进入到启动类所在的子模块, 使用`mvn package`进行编译打包, 出现如下报错

```
[WARNING] The POM for com.sdzx:qgg-core:jar:0.0.1-SNAPSHOT is missing, no dependency information available
[WARNING] The POM for com.sdzx:qgg-service:jar:0.0.1-SNAPSHOT is missing, no dependency information available
```

其中报`missing`的那几个 jar 包是该工程的几个其他的子模块. 网上有说可能是因为内网镜像仓库的配置不正确, 虽然我这边的确是在无公网环境, 全部走内网仓库, 但我觉得应该不是这个原因.

后来找到了参考文章1, 说是因为需要事先在各子模块下使用`mvn install`将其安装到 maven 的本地仓库中. 我找了找, 发现本地仓库中的确没有这几个子类的 jar. 于是分别在各子类模块的根目录下(`pom.xml`所在目录)执行`mvn install`. 当然, 由于各子类间也是有相互依赖的, 所以必须先安装没有依赖其他模块的子模块, 一步一步来才行.

后来在`install`第2个子模块时, 发现出现了对父级模块的依赖报错.

```
[ERROR] Failed to execute goal on project k8s-middleware-utils: Could not resolve dependencies for project com.cmos: k8s-middleware-utils:jar:1.0.0-SNAPSHOT：Failed to collect dependencies at com.cmos:k8s-middleware-beans:jar:1.0.0-SNAPSHOT: Failed to read artifact descriptor for com.cmos:k8s-middleware-beans:jar:1.0.0-SNAPSHOT: Failure to find com.cmos:k8s-middleware-sv:pom:1.0.0-SNAPSHOT in http://192.168.21.14:25000/nexus/content/groups/public/ was cached in the local repository, resolution will not be reattempted until the update interval of nexus-public has elapsed or updates are forced -> [Help 1]
```

其中`192.168.21.14:25000`是我内网全局仓库的地址, `nexus-public`是这个仓库的名称, `k8s-middleware-sv`为父模块的名称.

我到本地仓库找了找, 没找到父级模块的jar包(不过父级模块根本没有代码, 也没想到ta竟然也要生成jar包...), 然后尝试在父级模块根目录下执行`mvn install`, 竟然成功了, 这一操作所有子模块都打包安装了...

不过运行的话仍然是对启动类的jar执行.
