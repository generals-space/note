# Maven入门实践

参考文档

[maven3常用命令、java项目搭建、web项目搭建详细图解](http://blog.csdn.net/edward0830ly/article/details/8748986)

## 1. 写在前面

关于对maven的作用, 大部分介绍是以下两点

- 项目构建

- 依赖管理

目前对于maven的理解是,

首先, **maven是一个命令行工具**.

不过eclipse与idea中都有创建maven工程的选项, 不必使用maven提供的命令手动创建. 通过maven创建的工程, 拥有一定目录结构, 包括了源码, 测试部分, 使用maven编译后还会在特定路径生成目标class文件. 开发完成maven还整合了打包, 部署的功能. 整个项目结构清晰, 方便管理, 流程也更合理.

虽然有IDE包揽maven项目管理, 但是使用命令行了解的应该更为清晰一些, 需要有一点Java编程基础.

## 2. 环境安装

maven用来管理Java项目, 所以首先得有JDK. 然后是maven工具, 在maven官网可以下载. 还会介绍JavaWeb项目的创建, 所以需要tomcat.

- JDK: 1.8.0

- maven: 3.3.9(可执行文件包, 非源码编译)

- tomcat: 8.5.4

maven解压后, `maven/bin/mvn`为主要的maven命令, 将它软链接至`/usr/local/bin`目录下, 执行`mvn --version`输出maven版本信息即可.

## 3. 常用命令

### 3.1 创建maven工程(即Java项目, project).

```
mvn archetype:generate -DartifactId=my-app -DgroupId=com.mycompany.app -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false
```

- mvn archetype:generate　　固定格式

- -DartifactId　　　　　　　　项目名称(这个决定生成的Java项目顶层目录的名称)

- -DgroupId　　　　　　　　　组织标识(即包名, 这是项目的一个组件/模块, 这将出现在所创建的Java项目源文件中的`package`语句中)

- -DarchetypeArtifactId　　  指定ArchetypeId(项目类型)： `maven-archetype-quickstart`, 创建一个普通的Java工程; `maven-archetype-webapp`，创建一个JavaWeb工程.

- -DinteractiveMode　　　　　　是否使用交互模式

### 3.2 编译源代码

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

然后执行如下命令, 将会生成与`src`同级的`target`目录, 其下会得到`class`可执行文件.

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

### 3.3 编译测试代码

```
mvn test-compile
```

将生成测试部分的`class`文件

### 3.4 运行测试

```
mvn test
```

不会写测试代码, 不做分析...

### 3.5 打包

```
mvn package
```

将当前项目打包成jar文件, jar包就在`target`目录下.

### 3.6 安装当前工程的输出文件到本地仓库

```
mvn install
```

将`package`生成的jar包复制到本地maven仓库.

### 3.7 清空

```
mvn clean
```

`compile`生成的可执行文件放在`target`目录, `clean`将删除这个目录, 清空所有`compile`操作生成的文件.

### 3.8 查看maven有哪些项目类型分类

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

本文中只用到`maven-archetype-quickstart`与`maven-archetype-webapp`.

## 4. 创建项目

首次执行`mvn`系列命令会下载一些依赖文件, 可能会因为网络环境不能一次下载成功而出现为找到依赖的问题, 需要多次尝试. 已经拥有的jar包不会再次下载, 也可以节省时间.

### 4.1 普通工程

#### 4.1.1 创建工程

```
$ pwd
/home
$ ls

$ mvn archetype:generate -DartifactId=my-app -DgroupId=com.company.app -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false
## 创建了一个名为my-app的项目
$ ls
my-app
## 项目基本结构
$ tree ./my-app
./my-app/
├── pom.xml
└── src
    ├── main
    │   └── java
    │       └── com
    │           └── company
    │               └── app
    │                   └── App.java
    └── test
        └── java
            └── com
                └── company
                    └── app
                        └── AppTest.java

11 directories, 3 files
```

可以看到, mvn在`/home`目录下创建了一个名为`my-app`的项目, 其中`pom.xml`文件是项目核心. pom即`project object model`. 其内容中`groupId`, `artifactId`与`version`唯一确定一个项目, 称为 **`项目坐标`**.

#### 4.1.2 编译执行

src...目录下的`App.java`文件内容为

```
$ cd my-app
$ cat ./src/main/java/com/company/app/App.java
```

```java
package com.company.app;

/**
 * Hello world!
 *
 */
public class App
{
    public static void main( String[] args )
    {
        System.out.println( "Hello World!" );
    }
}
```

接下来编译源程序

```
$ pwd
/home/my-app
$ ls
pom.xml src
## maven编译命令
$ mvn compile
...
$ ls
pom.xml src target
$ tree
.
├── pom.xml
├── src
│   ├── main
│   │   └── java
│   │       └── com
│   │           └── company
│   │               └── app
│   │                   └── App.java
│   └── test
│       └── java
│           └── com
│               └── company
│                   └── app
│                       └── AppTest.java
└── target
    ├── classes
    │   └── com
    │       └── company
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

可以看到, 当前目录下多了一个与`src`同级的`target`目录, 其下有`App.java`编译得到的`App.class`文件. 我们执行一下它, 将会有`Hello World!`的输出, 但输出的信息可能有点多, 需要仔细找找.

```
$ mvn exec:java -Dexec.mainClass="com.company.app.App"
...
Hello World!
...
```

#### 4.1.3 执行测试

因为并不理解Java的测试原理(`AppTest.java`文件中没有`main`函数, 没办法手动执行), 所以只里只写出命令, 对于测试的结果暂时没有办法分析. 测试操作会在`my-app/target`下生成`test-class`目录, 生成的可执行测试文件`AppTest.class`就在这个目录下.

```
$ mvn test
...
Running com.company.app.AppTest
Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.009 sec

Results :

Tests run: 1, Failures: 0, Errors: 0, Skipped: 0
...
```

#### 4.1.4 打包

打包的涵义是, 将class可执行文件转换成jar包, 作为`库函数`供其他地方调用. 从这种意义上来说, `jar`包类似于windows平台中的`dll(动态链接库)`与linux平台中的`so(共享对象)`的概念.

```
$ pwd
/home/my-app
$ mvn package
$ ls target/
classes  maven-archiver  maven-status  my-app-1.0-SNAPSHOT.jar  surefire-reports  test-classes
```

打包之前会先执行一遍`编译`和`测试`操作.

其中`my-app/target/my-app-1.0-SNAPSHOT.jar`就是生成的目标文件, `项目名-版本号.jar`.

在`my-app/pom.xml`文件中存在一个名为`version`的字段, 默认即是`1.0-SNAPSHOT`. 尝试编辑pom.xml文件, 将`version`字段的值改为`1.0.0`, 并再次执行`mvn package`, 你将得到`my-app/target/my-app-1.0.0.jar`.

#### 4.1.5 安装

作为一名Java开发者, 应该会在本地保留一份所有用到过的jar包, 尤其是使用maven管理工程时. maven将其称为本地仓库, 工程中所有依赖的第3方jar包, 或是自己开发的项目生成的jar包, 都会保留在这个目录下.

接下来我们将上一节生成的jar文件"安装"到本地仓库中, 呃, 其实时复制.

```
$ pwd
/home/my-app
$ mvn install
...
[INFO] Installing /home/my-app/target/my-app-1.0.0.jar to /root/.m2/repository/com/company/app/my-app/1.0.0/my-app-1.0.0.jar
[INFO] Installing /home/my-app/pom.xml to /root/.m2/repository/com/company/app/my-app/1.0.0/my-app-1.0.0.pom
...

$ ls /root/.m2/repository
antlr  backport-util-concurrent  com          commons-codec        commons-io    commons-logging  jdom   log4j  org  xml-apis
asm    classworlds               commons-cli  commons-collections  commons-lang  dom4j            junit  net    oro  xpp3
```

这样, 我们将jar包安装到了`/root/.m2/repository`对应的项目下.

这个仓库地址的配置在`maven/conf/setting.conf`文件`localRepository`字段中.

```
<!-- localRepository
   | The path to the local repository maven will use to store artifacts.
   |
   | Default: ${user.home}/.m2/repository
  <localRepository>/path/to/local/repo</localRepository>
  -->
```

尝试解开`localRepository`的注释并将其值修改为`/opt/java/repository`, 再次执行`mvn install`, 你将会在该目录找到你的Java项目.

```
$ cd /opt/java/repository
$ ls
backport-util-concurrent  classworlds  com  commons-cli  commons-lang  commons-logging  hsperfdata_root  junit  log4j  org
$ ls com/company/app/my-app/
1.0.0  maven-metadata-local.xml
$ ls com/company/app/my-app/1.0.0
my-app-1.0.0.jar  my-app-1.0.0.pom  _remote.repositories
```

#### 4.1.6 运行jar包

jar包执行方式与class文件相同.

```
$ pwd
/home/my-app
$ ls
pom.xml src target
$ java -cp target/my-app-1.0.0.jar com.company.app.App
Hello World!
## 或者
$ java -cp target/classes com.company.app.App
Hello World!
```

### 4.2 Web工程

#### 4.2.1 创建工程

```
$ pwd
/home
$ mvn archetype:generate -DartifactId=my-web-app -DgroupId=com.company.app -DarchetypeArtifactId=maven-archetype-webapp -DinteractivMode=false
...
[INFO] Using property: groupId = com.company.app
[INFO] Using property: artifactId = my-web-app
## 这里会提示输入版本号, 示例中输入1.0.0
Define value for property 'version':  1.0-SNAPSHOT: :1.0.0
## 确认项目信息
[INFO] Using property: package = com.company.app
Confirm properties configuration:
groupId: com.company.app
artifactId: my-web-app
version: 1.0.0
package: com.company.app
## 不知道小写的y可不可以
 Y: : Y
...

$ ls
my-app  my-web-app
## web项目基本结构
$ tree my-web-app
my-web-app/
├── pom.xml
└── src
    └── main
        ├── resources
        └── webapp
            ├── index.jsp
            └── WEB-INF
                └── web.xml

5 directories, 3 files

```

#### 4.2.2 打包

这个新建的空项目只有一个`index.jsp`页面. 其内容为

```jsp
<html>
<body>
<h2>Hello World!</h2>
</body>
</html>
```

不需要编译, 直接就能打包发布.

```
$ pwd
/home/my-web-app
$ mvn package
$ tree
.
├── pom.xml
├── src
│   └── main
│       ├── resources
│       └── webapp
│           ├── index.jsp
│           └── WEB-INF
│               └── web.xml
└── target
    ├── classes
    ├── maven-archiver
    │   └── pom.properties
    ├── my-web-app
    │   ├── index.jsp
    │   ├── META-INF
    │   └── WEB-INF
    │       ├── classes
    │       └── web.xml
    └── my-web-app.war

12 directories, 7 files
```

将`my-web-app/target/my-web-app.war`放到目标`tomcat/webapp`目录下就可以访问了, 注意url路径应为`http://127.0.0.1:8080/my-web-app/index.jsp`这种.

#### 4.2.3 mvn部署war包到远程tomcat

使用maven将war包直接部署到远程tomcat(而不是拷贝war包过去), 需要拥有目标tomcat的管理员权限, 说到这里, 目标tomcat还需要在webapps下存在manager工程, maven会通过它提供的接口进行部署.

假设目标tomcat所在服务器为A, tomcat监听端口为8080. 编辑`tomcat/conf/tomcat-users.xml`文件, 确认存在如下内容并且已解开注释.

```xml
<tomcat-users>
<user username="admin" password="111111" roles="manager-gui,manager-script,manager-jmx,manager-status"/>
</tomcat-users>
```

现在目标tomcat拥有管理员帐号`admin`, 密码为`111111`.

接下来需要配置本地的maven, 将远程tomcat的管理员帐号密码写到其配置文件中, 其位置在`$MAVEN_HOME/conf/settings.xml`.

```xml
  <servers>
    <server>
      <id>Tomcat_A</id>
      <username>admin</username>
      <password>111111</password>
    </server>
  </servers>
```

其中`servers`下`server`字段可以存在多个, 认证方式也可以有私钥+密码等, 但是目标tomcat的IP与端口却不是在这里配置的, 之后可以再做深究.

然后还要配置工程的`pom.xml`文件.

```xml
<build>
  <finalName>my-web-app</finalName>

  <plugins>
    <plugin>
      <groupId>org.apache.tomcat.maven</groupId>
      <artifactId>tomcat7-maven-plugin</artifactId>
      <version>2.2</version>
      <configuration>
         <!--目标tomcat的manager地址, 注意manager/text这一段是必须的-->
         <url>http://Tomcat_A的IP地址:8080/manager/text</url>
         <!--这里server的值要与上面maven中server字段的id相匹配-->
         <server>Tomcat_A</server>
         <!--工程的访问地址-->
         <path>/my-web-app</path>
      </configuration>
    </plugin>
  </plugins>
</build>
```

接下来可以使用mvn命令执行部署了.

```
$ pwd
/home/my-web-app
$ ls
pom.xml src
$ mvn tomcat7:deploy
```

然后可以通过`http://Tomcat_A的IP:8080/my-web-app/index.jsp`访问工程.

#### 4.2.4 mvn集成jetty

嗯...这个例子除了开发的时候可能比较方便, 其他好像没什么用, 不过方便也算是一个理由了. JavaWeb项目直接使用jetty作为web容器.

修改`my-web-app/pom.xml`文件, 添加`jetty`作为项目插件

```xml
<build>
  <finalName>my-web-app</finalName>
  <!--有的教程中有这个字段, 但我试验时出现Unrecognised tag: 'pluginManagement'
  <pluginManagement>
  -->
    <!--添加Jetty配置-->
    <plugins>
      <plugin>
       <groupId>org.mortbay.jetty</groupId>   
       <artifactId>maven-jetty-plugin</artifactId>
      </plugin>
    </plugins>
  <!--<pluginManagement>-->
</build>
```

```
$ pwd
/home/my-web-app
$ tree
.
├── pom.xml
└── src
    └── main
        ├── resources
        └── webapp
            ├── index.jsp
            └── WEB-INF
                └── web.xml

5 directories, 3 files
mvn jetty:run
...
```

访问地址`http://127.0.0.1:8080/my-web-app/index.jsp`.

## 5. 本地Maven仓库配置

## 6. 远程镜像仓库配置

阿里云maven镜像

```xml
 <mirrors>
    <mirror>
      <id>alimaven</id>
      <name>aliyun maven</name>
      <url>http://maven.aliyun.com/nexus/content/groups/public/</url>
      <mirrorOf>central</mirrorOf>        
    </mirror>
  </mirrors>
```