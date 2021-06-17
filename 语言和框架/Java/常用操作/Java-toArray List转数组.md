# Java-toArray List转数组

参考文章

1. [Java集合转有类型的数组之toArray(T[] a)](https://www.cnblogs.com/guanghe/p/10062975.html)

`List<String>` 和 `String[]`不能使用强转, 会报错. 

现有如下两种转换方法, 假设`list`为`String`类型的列表

## 1. 

```java
String[] array = new String[list.size()];
list.toArray(array);
```

## 2.

```java
String[] array= list.toArray(new String[list.size()]);
```
