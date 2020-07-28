# 对象与json互转.1.fasterxml

参考文章

1. [利用fasterxml实现对象与JSON的转换](https://www.jianshu.com/p/0f6b14f9bae4)
    - 比较笼统
2. [Could not read JSON: Cannot construct instance of''类名""(no Creators, like default construct, exist)](https://www.cnblogs.com/shan-blog/p/12880733.html)
    - String 转 Object 时报错的解决方法: 无参的构造方法.

使用 springboot 的 web starter 貌似不需要在`pom.xml`中添加依赖, 直接就能引用.

我们的工程目录如下

![](https://gitee.com/generals-space/gitimg/raw/master/c5fb023f8083c946e94d154e6b1b220d.png)

默认启动类`DemoApplication.java`(因为用的是 web starter, 所以显得有点繁琐)

```java
package com.example.javastruct;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.LinkedList;
import java.util.List;

@SpringBootApplication
@RestController
public class DemoApplication {

	@GetMapping("/out")
	public String out() {
		List<VersionLog> list = new LinkedList<VersionLog>() {
			{
				add(new VersionLog("v0.0.1", new LinkedList<>() {
					{
						add("1. aaa");
						add("2. bbb");
					}
				}));
			}
		};
		VersionLogList vlList = new VersionLogList(list);
		return JsonUtil.pojoToString(vlList);
	}

	@GetMapping("/in")
	public String in() {
		String str = "{\"versionList\":[{\"title\":\"v0.0.1\",\"versionDescribe\":[\"1. aaa\",\"2. bbb\"]}]}";
		VersionLogList list = JsonUtil.stringToPojo(str, VersionLogList.class);
		return list.versionList.get(0).getTitle() + " " +
			list.versionList.get(0).getVersionDescribe();
		// return "hello you";
	}

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}
}

```

`JsonUtil.java`

```java
package com.example.javastruct;

import java.io.IOException;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

public class JsonUtil {
    private static final ObjectMapper mapper = new ObjectMapper();

    public static <T> T stringToPojo(String str, Class<T> clazz) {
        try {
            return mapper.readValue(str, clazz);
        } catch (IOException e) {
            System.out.println(e);
            return null;
        }
    }

    public static String pojoToString(Object object) {
        try {
            return mapper.writeValueAsString(object);
        } catch (JsonProcessingException e) {
            System.out.println(e);
            return null;
        }
    }
}

```

`VersionLog.java`

```java
package com.example.javastruct;

import java.util.List;

public class VersionLog {
    String title;
    List<String> versionDescribe;

    public VersionLog() {
    }

    public VersionLog(String title, List<String> versionDescribe) {
        this.title = title;
        this.versionDescribe = versionDescribe;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public List<String> getVersionDescribe() {
        return versionDescribe;
    }

    public void setVersionDescribe(List<String> versionDescribe) {
        this.versionDescribe = versionDescribe;
    }
}

```

`VersionLogList.java`

```java
package com.example.javastruct;

import java.util.LinkedList;
import java.util.List;

public class VersionLogList {
    public List<VersionLog> versionList = new LinkedList<VersionLog>();

    public VersionLogList() {
    }

    public VersionLogList(List<VersionLog> list) {
        this.versionList = list;
    }

    public List<VersionLog> getVersionList() {
        return versionList;
    }

    public void setVersionList(List<VersionLog> versionList) {
        this.versionList = versionList;
    }
}

```

这样, 访问`localhost:8080/out`, 将Java对象转换成 json 的结果如下

![](https://gitee.com/generals-space/gitimg/raw/master/81c8ff6bdad4546fa788f241469e1f28.png)

访问`localhost:8080/in`, 将 json 转换成 Java 对象的结果如下

![](https://gitee.com/generals-space/gitimg/raw/master/b0680dfa4a83bf08277af5f4e9650512.png)

1. 要使用`JsonUtil`中的`pojoToString()`方法, 那么 Java 对象必须拥有各属性的`getter/setter`方法, 没有`getter/setter`方法的属性将不会显示在返回的 String 结果中(我刚接触Java, 没有写`getter/setter`的习惯, 于是最初得到一个空对象...).
2. 要使用`JsonUtil`中的`stringToPojo()`方法, 则需要 Java 对象必须拥有一个无参构造函数, 否则会出错. 见参考文章2.

## FAQ

在进行 json 转 Java 对象的实验时, 访问接口, 日志中出现异常, 报错如下.

```
com.fasterxml.jackson.databind.exc.InvalidDefinitionException: Cannot construct instance of `com.example.javastruct.VersionLogList` (no Creators, like default constructor, exist): cannot deserialize from Object value (no delegate- or property-based Creator)
 at [Source: (String)"{"versionList":[{"title":"v0.0.1","versionDescribe":["1. aaa","2. bbb"]}]}"; line: 1, column: 2]
```

按照参考文章2的说法, 是因为我缺少了默认的构造函数(无参), 后来分别为`VersionLog`与`VersionLogList`补上无参构造函数后, 就正常了. 

> 注意: `VersionLog`与`VersionLogList`两个都要有无参构造函数.

