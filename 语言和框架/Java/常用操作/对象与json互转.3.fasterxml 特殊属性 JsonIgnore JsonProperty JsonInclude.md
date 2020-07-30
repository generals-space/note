# 对象与json互转.3.fasterxml 特殊属性 JsonIgnore JsonProperty JsonInclude

参考文章

1. [fasterxml Jackson 各种使用场景以及序列化、反序列化的时候忽略不必要的Properties](https://blog.csdn.net/aHardDreamer/article/details/89235188)
    - `@JsonIgnore`注解, 在格式化为 json 可以忽略指定属性, 类似于 golang 的 `omitempty`.
2. [Jackson 序列化/反序列化时忽略某属性](https://www.iteye.com/blog/wwwcomy-2397340)
3. [Setting default values to null fields when mapping with Jackson](https://stackoverflow.com/questions/18805455/setting-default-values-to-null-fields-when-mapping-with-jackson)
    - 默认值

## 1. Java 对象 -> json 时忽略某些字段

有时定义的 Class 在将其格式化为 json 时, 希望某些字段不要出现在结果中(比如`password`这种), 需要让`fasterxml`忽略这些字段.

参考文章1中提出可以使用`@JsonIgnore`为指定字段添加注解.

但是参考文章2提到, `@JsonIgnore`这个注解在1.9版本后, 会使得该属性的`getter`方法无法获得正确的值, 就是说, 除了序列化成json的时候这个字段不会出现外, 使用`getter`方法从对象中取值时也会得到`null`, 这太坑了.

我用的是`2.8.0`, 按照参考文章2所说的解决方法, 使用了`@JsonProperty(access = Access.WRITE_ONLY)`, 这样可以在 json 结果中隐藏该字段, 也不会影响 getter 的结果, 从字符串格式化成对象的时候也没有问题.

貌似还有一些其他的 bug, 这里先不讨论了.

> `@JsonProperty(access = Access.WRITE_ONLY)`用于修饰类成员字段.

```java
public class VersionLog {
    @JsonProperty(access = Access.WRITE_ONLY)
    String title = "默认title";
}
```

------

上面是显式指定哪些值可以被格式化到 json, 如果要想像 golang 那样(`omitempty`), 值为`null`的字段就不要显示, 则需要另一个属性.

```java
@JsonInclude(JsonInclude.Include.NON_NULL)
public class VersionLog {
    String title = "默认title";
}
```

## 2. json -> Java 对象时字段默认值

前面的示例中, 在将 json 字符串格式化成 Java 对象时, 如果某些字段缺失, 在 Java 对象中该字段将会为`null`(不管该字段是`String`, `Integer`还是`List`, `Map`什么的), 不像 golang 样还有个默认值.

不过, `fasterxml`貌似没有提供这样的注解, 按照参考文章3的采纳答案, 只能在目标类的字段定义时显式写明默认值.

```java
public class VersionLog {
    String title = "默认title";
}
```

## 3. json -> Java 对象时未知字段

如果 json 字符串中包含 Java Class 中未定义的字段, 那么在解析的时候可能会报如下错误.

```
Caused by: com.fasterxml.jackson.databind.exc.UnrecognizedPropertyException: Unrecognized field "new_key" (class com.example.javastruct.svc.VersionLog), not marked as ignorable (6 known properties: "versionDescribe", "involvedResource", "arch", "title", "age", "version"])
```

如果想要忽略 json 中的多余字段, 可以使用`@JsonIgnoreProperties(ignoreUnknown = true)`修饰目标类对象.

```java
@JsonIgnoreProperties(ignoreUnknown = true)
public class VersionLog {
}
```
