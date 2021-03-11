# Java-switch..case不加break会匹配下一个case

参考文章

1. [java switch 不加 break 继续执行 下一个case(不用匹配条件) 这个设计是因为什么？](https://www.zhihu.com/question/27819136/answer/38212236)

最近在写 java 代码的时候发现一个奇怪的问题, `switch..case..`语句中每个`case`结尾都要加一个`break`, 否则会继续匹配之后的`case`.

```java
public class Main{
    public static void main(String[] args) {
        System.out.println("========== ");
        Integer var = 0;
        var = 1;
        switch(var){
            case 1:
                System.out.println("========== " + 1);
            case 2:
                System.out.println("========== " + 2);
            case 3:
                System.out.println("========== " + 3);
        }
    }
}
```

工程目录如下

![](https://gitee.com/generals-space/gitimg/raw/master/f2e542f006ee3749e7053d2c5481cb77.png)

这一点与我熟悉的 golang 是不同的, 如下代码只会匹配到`case 2`语句.

```go
	num := 2;
	switch(num){
	case 1:
		fmt.Printf("=========== %d", 1)
	case 2:
		fmt.Printf("=========== %d", 2)
	case 3:
		fmt.Printf("=========== %d", 3)
	}
```

打印结果为

```
=========== 2
```

如果要实现像`java`那样继续匹配之后的`case`, 还需要额外借助`fallthrough`关键字.

按照参考文章1的说法, 这是C, C++, Java 一直以来的传统, 不忘初心, 薪火相传...

