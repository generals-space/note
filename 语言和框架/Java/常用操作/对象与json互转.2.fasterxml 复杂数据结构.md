# 对象与json互转.2.fasterxml 复杂数据结构

参考文章

1. [jackson构建复杂Map的序列化](https://www.cnblogs.com/zhshlimi/p/10300159.html)
    - 类型较多, 但是使用示例只有一个, 不太直观, 但其实很详细
2. [Jackson 处理复杂类型(List,map)两种方法](https://www.cnblogs.com/surge/p/9046223.html)

## 1. 引言

本文的示例延续上一篇, 但是`VersionLogList`的处理方式有点不一样.

json数据如下

```json
{
    "versionList": [
        {
            "title": "v0.0.1",
            "versionDescribe": [
                "1. aaa",
                "2. bbb"
            ]
        }
    ]
}
```

由于我们希望在`VersionLogList`的构造函数中将其读取并构造自身, 就不能再用`List<VersionLog> versionList`这个成员属性了. 先看看下面的代码的异常之处

```java
public class VersionLogList {
    public List<VersionLog> versionList = new LinkedList<VersionLog>();

    public VersionLogList() {
        // 假设这里是上面的 json 字符串.
        String str = "" 
        // 那么问题来了, 这里是不是怪怪的???
        this.versionList = JsonUtil.stringToPojo(str, VersionLogList.class);
    }
    // 省略 getter/setter
}
```

`versionList`是 List<VersionLog> 类型, 这样赋值是绝对不行的.

但是能像`this = `直接对 this 对象赋值吗? 好像也不行.

那么能把右侧改成`stringToPojo(str, List<VersionLog>.class)`吗? 想想更不可能.

所以只能变一下了, 我们为`VersionLogList`添加`private Map<String, List> info;`成员, 希望将 json 数据先格式化成`Map`类型.

## 2. 代码

默认启动类`DemoApplication.java`

```java
package com.example.javastruct;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import com.example.javastruct.svc.JsonUtil;
import com.example.javastruct.svc.VersionLogList;

@SpringBootApplication
@RestController
public class DemoApplication {
	@Autowired
	private VersionLogList versionList;

	@GetMapping("/")
	public String home() {
		return JsonUtil.pojoToString(versionList.getInfo());
	}

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}
}

```

`JsonUtil.java`

```java
package com.example.javastruct.svc;

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

    public static <T>T objectToPojo(Object object, Class<T> clazz){
        String str = pojoToString(object);
        return stringToPojo(str, clazz);
    }
}

```

`VersionLog.java`, 这个没有变动

```java
package com.example.javastruct.svc;

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
package com.example.javastruct.svc;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import com.fasterxml.jackson.databind.JavaType;
import com.fasterxml.jackson.databind.ObjectMapper;

import org.springframework.core.io.ClassPathResource;
import org.springframework.core.io.Resource;
import org.springframework.stereotype.Service;

@Service
public class VersionLogList {
    private Map<String, List> info;

    private static final ObjectMapper jsonMapper = new ObjectMapper();
    private static final JavaType strType = jsonMapper.getTypeFactory().constructType(
        String.class
    );
    private static final JavaType listType = jsonMapper.getTypeFactory().constructCollectionType(
        ArrayList.class,
        VersionLog.class
    );
    private static final JavaType strListMap = jsonMapper.getTypeFactory().constructMapType(
        HashMap.class, 
        strType,
        listType
    );

    public VersionLogList() throws Exception {
        Resource resource = new ClassPathResource("static/version_list.json");
        InputStream is = resource.getInputStream();
        InputStreamReader isr = new InputStreamReader(is);
        BufferedReader br = new BufferedReader(isr);
        StringBuffer sb = new StringBuffer();

        String s = "";
        while ((s = br.readLine()) != null) {
            sb.append(s);
        }
        String content = sb.toString();

        br.close();
        isr.close();
        is.close();
        
        System.out.println(content);
        info = jsonMapper.readValue(content, strListMap);
		// 虽然在定义 listType 的时候指定了 List 的成员类型就是 VersionLog, 但还是需要进行显式的类型转换.
		VersionLog vLog = (VersionLog)info.get("versionList").get(0);
        System.out.println("=====================");
        System.out.println(vLog.getTitle());
    }

    public Map<String, List> getInfo() {
        return info;
    }

    public void setInfo(Map<String, List> info) {
        this.info = info;
    }
}
```

## 3. 运行

![](https://gitee.com/generals-space/gitimg/raw/master/fe05efe5c5bdc9f3e6e6cffc2a599997.png)

由于 info 其实是`Map<String,List>`类型, 所以从里面取出`VersionLog`成员还需要经过显式的类型转换, 如果使用比较频繁, 可以考虑~~把`List<VersionLog> versionList`成员再回回来, 然后用`JsonUtil.objectToPojo(info.get("versionList"), )`~~ 当我没说...
