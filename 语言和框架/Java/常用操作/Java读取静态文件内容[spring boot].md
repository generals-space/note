# Java读取静态文件内容[spring boot]

参考文章

1. [java（包括springboot）读取resources下文件方式](https://www.cnblogs.com/whalesea/p/11677657.html)
    1. `File()` 使用项目相对路径 (打包后无法读取, 不通用)
    2. `org.springframework.util.ResourceUtils.getFile("classpath:resource.properties")` (Linux环境中无法读取, 不通用)
    3. `new org.springframework.core.io.ClassPathResource("resource.properties")`(通用)
    4. `org.springframework.core.io.ResourceLoader.getResource("classpath:resource.properties")`(通用)
2. [Java(springboot) 读取txt文本内容](https://www.cnblogs.com/strideparty/p/9517713.html)
    - `content = content.append(s)`

## 1. 

```java
File file = new File("src/main/resources/resource.properties");
```

## 2. 

```java
File file = ResourceUtils.getFile("classpath:resource.properties");
```

## 3. 

```java
Resource resource = new ClassPathResource("resource.properties");
```

## 4. 

```java
Resource resource = resourceLoader.getResource("classpath:resource.properties");
```

