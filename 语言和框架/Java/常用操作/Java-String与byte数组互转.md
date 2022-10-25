# Java-String与byte[]互转

参考文章

1. [Java中String与byte[]的转换](https://www.cnblogs.com/EasonJim/p/8142705.html)

```java
String s = "easonjim";          // String变量 
byte b[] = s.getBytes();        // String转换为byte[] 
String t = new String(b);       // bytep[]转换为String，支持传递编码
```
