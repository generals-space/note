# 命令行编译运行java程序

参考文章

1. [命令行执行java程序](http://blog.csdn.net/lee_decimal/article/details/5885406)
2. [在CMD下使用Java命令执行class文件](http://ivan0513.iteye.com/blog/982445)

## 1. 源文件无package语句

```console
$ pwd
/home/App
$ ls
App.java
less App.java
```

```java
//App.java源文件内容
public class App
{
    public static void main( String[] args )
    {
        System.out.println( "Hello World!" );
    }
}
```

```console
$ javac ./App.java
$ ls
App.class  App.java
$ java App
Hello World!
```

## 2. 源文件存在package语句

```console
$ pwd
/home/App/space/generals/java
$ ls
App.java
$ less App.java
```

```java
//App.java源文件内容
//注意package语句与源文件路径的关系
package space.generals.java;
public class App
{
    public static void main( String[] args )
    {
        System.out.println( "Hello World!" );
    }
}
```

```console
$ javac App.java
$ ls
App.class App.java
$ pwd
/home/App/space/generals/java
$ cd /home/App
$ tree
.
└── space
    └── generals
        └── java
            └── App.java
## 注意执行此类java程序的方法, 以及当前目录
$ java space.generals.java.App
```

------

以上两种都是在Java项目根目录, 指定了要执行的Java程序. 想要在任意目录执行目标程序的话, 需要在命令行指定`-cp/-classpath`参数, 代表目标程序的位置.

以上面两者为例.

```console
$ pwd
/tmp
## 无package语句时
$ java -cp /home/App App
## 有package语句时
$ java -cp /home/App space.generals.java.App
```

当然, 编译的话, 只要指定待编译的源文件路径即可.
