# Maven入门实践.1.基本工程

参考文档

[maven3常用命令、java项目搭建、web项目搭建详细图解](http://blog.csdn.net/edward0830ly/article/details/8748986)

## 1. 写在前面

关于对maven的作用, 大部分介绍是以下两点

1. 项目构建
2. 依赖管理

**maven是一个命令行工具**. 不过`eclipse`与`idea`中都有创建maven工程的选项, 不必使用maven提供的命令手动创建. 

通过maven创建的工程, 拥有一定目录结构, 包括了源码, 测试部分, 使用maven编译后还会在特定路径生成目标class文件. 开发完成maven还整合了打包, 部署的功能. 整个项目结构清晰, 方便管理, 流程也更合理.

虽然有IDE包揽maven项目管理, 但是使用命令行了解的应该更为清晰一些, 需要有一点Java编程基础.

## 2. 环境安装

maven用来管理Java项目, 所以首先得有JDK. 然后是maven工具, 在maven官网可以下载. 还会介绍JavaWeb项目的创建, 所以需要tomcat.

- JDK: 1.8.0
- maven: 3.3.9(可执行文件包, 非源码编译)
- tomcat: 8.5.4

maven解压后, `maven/bin/mvn`为主要的maven命令, 将它软链接至`/usr/local/bin`目录下, 执行`mvn --version`输出maven版本信息即可.

## 3. maven工程的创建, 编译与运行

### 3.1 创建

```console
$ mvn archetype:generate -DartifactId=my-app -DgroupId=com.mycompany.app -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false
[INFO] Using property: groupId = com.mycompany.app
[INFO] Using property: artifactId = my-app
## 这里会提示输入版本号, 示例中输入1.0.0
Define value for property 'version' 1.0-SNAPSHOT: : 1.0.0
[INFO] Using property: package = com.mycompany.app
Confirm properties configuration:
groupId: com.mycompany.app
artifactId: my-app
version: 1.0.0
package: com.mycompany.app
 Y: : Y
```

这会创建一个名为`my-app`的目录, 上面的参数都会写入到生成工程的`pom.xml`文件中.

- `mvn archetype:generate`: 固定格式
- `-DartifactId`: 项目名称(这个决定生成的Java项目顶层目录的名称)
- `-DgroupId`: 组织标识(即包名, 这是项目的一个组件/模块, 这将出现在所创建的Java项目源文件中的`package`语句中)
- `-DarchetypeArtifactId`: 指定ArchetypeId(项目类型)
    - `maven-archetype-quickstart`: 普通的Java工程; 
    - `maven-archetype-webapp`, JavaWeb工程;
    - 更多类型在下文给出
- `-DinteractiveMode`: 是否使用交互模式, 默认为`false`.

进入到生成的`my-app`目录中, 其目录结构为

```
.
├── pom.xml
└── src
    ├── main
    │   └── java
    │       └── com
    │           └── mycompany
    │               └── app
    │                   └── App.java
    └── test
        └── java
            └── com
                └── mycompany
                    └── app
                        └── AppTest.java

11 directories, 3 files
```

`pom.xml`内容如下

```xml
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">
  <modelVersion>4.0.0</modelVersion>
  <groupId>com.mycompany.app</groupId>
  <artifactId>my-app</artifactId>
  <packaging>jar</packaging>
  <version>1.0.0</version>
  <name>my-app</name>
  <url>http://maven.apache.org</url>
  <dependencies>
    <dependency>
      <groupId>junit</groupId>
      <artifactId>junit</artifactId>
      <version>3.8.1</version>
      <scope>test</scope>
    </dependency>
  </dependencies>
</project>

```

### 3.2 编译

执行如下命令, 将会生成与`src`同级的`target`目录, 其下会得到`class`可执行文件.

```
mvn compile
```

```
.
├── pom.xml
├── src
│   ├── main
│   │   └── java
│   │       └── com
│   │           └── mycompany
│   │               └── app
│   │                   └── App.java
│   └── test
│       └── java
│           └── com
│               └── mycompany
│                   └── app
│                       └── AppTest.java
└── target
    ├── classes
    │   └── com
    │       └── mycompany
    │           └── app
    │               └── App.class
    └── maven-status
        └── maven-compiler-plugin
            └── compile
                └── default-compile
                    ├── createdFiles.lst
                    └── inputFiles.lst

20 directories, 6 files
```

### 3.3 测试

`mvn test-compile`: 将生成测试部分的`class`文件

`mvn test`: 进行测试(不会写测试代码, 不做分析...)

### 3.4 打包

```
mvn package
```

将当前项目打包成jar文件, 生成的jar包就在`target`目录下.

### 3.5 安装当前工程的输出文件到本地仓库

```
mvn install
```

将`package`生成的jar包复制到本地maven仓库.

### 3.6 清空

```
mvn clean
```

`compile`生成的可执行文件放在`target`目录, `clean`将删除这个目录, 清空所有`compile`操作生成的文件.

### 3.7 查看maven有哪些项目类型分类

```
$ mvn archetype:generate -DarchetypeCatalog=intrenal
maven-archetype-archetype  
maven-archetype-j2ee-simple
maven-archetype-plugin     
maven-archetype-plugin-site
maven-archetype-portlet    
maven-archetype-profiles   
maven-archetype-quickstart
maven-archetype-site       
maven-archetype-site-simple
maven-archetype-webapp

```

> 本文中只用到`maven-archetype-quickstart`与`maven-archetype-webapp`.
