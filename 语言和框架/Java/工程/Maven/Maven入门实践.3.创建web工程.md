# Maven入门实践.3.创建web工程

## 1. 创建工程

```console
$ pwd
/home
$ mvn archetype:generate -DartifactId=my-web-app -DgroupId=com.company.app -DarchetypeArtifactId=maven-archetype-webapp -DinteractivMode=false
...
[INFO] Using property: groupId = com.company.app
[INFO] Using property: artifactId = my-web-app
## 这里会提示输入版本号, 示例中输入1.0.0
Define value for property 'version':  1.0-SNAPSHOT: :1.0.0
[INFO] Using property: package = com.company.app
Confirm properties configuration:
groupId: com.company.app
artifactId: my-web-app
version: 1.0.0
package: com.company.app
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

## 2. 打包

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

## 3. mvn部署war包到远程tomcat

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

## 4. mvn集成jetty(可跳过)

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
$ mvn jetty:run
...
```

访问地址`http://127.0.0.1:8080/my-web-app/index.jsp`.
