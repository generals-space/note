# Java长注释

参考文章

1. [Java注释：类、方法和字段注释](http://c.biancheng.net/view/6114.html)

java的函数块注释要写成如下格式

```java
/**
 * 函数功能
 * @param xxx: 
 */
```

而不能写成

```java
/*
    函数功能
*/
```

后者在 vscode 中, 通过在调用函数的函数名处悬停是不会有代码提示的, 只能使用前者的格式.
