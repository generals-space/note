# Maven入门实践.2.创建普通工程

## 1. 创建工程

```console
$ pwd
/home
$ ls
## 空目录
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

## 2. 编译执行

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

## 3. 执行测试

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

## 4. 打包

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

## 5. 安装

作为一名Java开发者, 应该会在本地保留一份所有用到过的jar包, 尤其是使用maven管理工程时. maven将其称为本地仓库, 工程中所有依赖的第3方jar包, 或是自己开发的项目生成的jar包, 都会保留在这个目录下.

接下来我们将上一节生成的jar文件"安装"到本地仓库中, 呃, 其实就是复制(类似于`go install`, `make install`等).

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

这个仓库地址的配置在`$maven/conf/setting.conf`文件`localRepository`字段中.

```xml
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

## 6. 运行jar包

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
