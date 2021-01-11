# snakeyaml

参考文章

1. [SnakeYaml快速入门](https://www.jianshu.com/p/d8136c913e52)
2. [24. Externalized Configuration](https://www.docs4dev.com/docs/zh/spring-boot/2.1.1.RELEASE/reference/boot-features-external-config.html)
    - `spring-boot-starter`自动提供`SnakeYAML`(不需要额外添加`dependency`块)
3. [SnakeYAML Documentation](https://bitbucket.org/asomov/snakeyaml/wiki/Documentation#markdown-header-snakeyaml-documentation)
    - 高级, 详细, 值得收藏
    - `SafeConstructor()`, 只处理值类型为Java内置类型的数据, 其他的都将解析成`null`
    - `putListPropertyType()`与`putMapPropertyType()`将要被废弃, 可以使用`addPropertyParameters()`方法代替.

## 1. 简单使用

```yaml
cluster.name: elasticsearch 
node.name: es-01
network.host: 0.0.0.0
http.port: 9200
bootstrap.memory_lock: false
cluster.initial_master_nodes: ["es-01"]
## 中文注释
xpack.monitoring.exporters.mylocal:
  type: local
```

```java
    Yaml yaml = new Yaml();

    Object object = yaml.load(content.toString());
    // {cluster.name=elasticsearch, node.name=es-01, network.host=0.0.0.0, http.port=9200, bootstrap.memory_lock=false, cluster.initial_master_nodes=[es-01], xpack.monitoring.exporters.mylocal={type=local}}
    System.out.println(object);

    Map<String, Object> map = (Map<String, Object>)object;
    // {cluster.name=elasticsearch, node.name=es-01, network.host=0.0.0.0, http.port=9200, bootstrap.memory_lock=false, cluster.initial_master_nodes=[es-01], xpack.monitoring.exporters.mylocal={type=local}}
    System.out.println(map);
```

`load()`方法默认生成一个`Map<String, Object>`类型的对象, 几乎可以解析所有层级的(只要是基本类型).

## 2. 解析yaml到自定义对象类型

```yaml
name: general
age: 24
gender: 男
lesson:
  - 语文
  - 数学
  - 英语

```

```yaml
import java.util.List;

public class Person {
    String name;
    Integer age;
    String gender;
    List<String> lesson;
    // getter 与 setter 方法
}
```

```java
import org.yaml.snakeyaml.constructor.Constructor;
        // 这里是重点!!!
        Yaml yaml = new Yaml(new Constructor(Person.class));

        Person person = yaml.load(content.toString());
        // com.example.demo.Person@7c3fb849
        System.out.println(person);
        // {"name":"general","age":24,"gender":"男","lesson":["语文","数学","英语"]}
        System.out.println(JsonUtil.pojoToString(person));

```

## 3. 自定义类型对象嵌套

```yaml
name: general
age: 24
gender: 男
lesson:
  - 语文
  - 数学
  - 英语
car:
  - {model: 比亚迪}
  - {model: 特斯拉}

## 这种是不正确的... Map 嵌套要多一层才可以
## myCar: 
##   model: 小三轮儿
myCar: 
  car1: 
    model: 小三轮儿
  car2: 
    model: 小黄车
```

看了看参考文章3对自定义类型嵌套的示例, 本来还有点萌b, ta就给出了`List`和`Map`两种方式. 在实验的时候才意识到, 嵌套使用真就只有这两种方式...

```java

import java.util.List;
import java.util.Map;

public class Person {
    String name;
    Integer age;
    String gender;
    List<String> lesson;
    List<Car> car;
    Map<String, Car> myCar;
    // getter 与 setter 方法
}
```

```java
public class Car {
    private String model;
    // getter 与 setter 方法
}
```

```java
        Constructor constructor = new Constructor(Person.class);
        TypeDescription typeDesc = new TypeDescription(Person.class);
    // 将要废弃, 使用 addPropertyParameters() 代替.
        // typeDesc.putListPropertyType("car", Car.class);
        // typeDesc.putMapPropertyType("myCar", String.class, Car.class);
        typeDesc.addPropertyParameters("car", Car.class);
        typeDesc.addPropertyParameters("myCar", String.class, Car.class);
        constructor.addTypeDescription(typeDesc);

        Yaml yaml = new Yaml(constructor);

        Person person = yaml.load(content.toString());
        System.out.println(person);
        // {"name":"general","age":24,"gender":"男","lesson":["语文","数学","英语"],"car":[{"model":"比亚迪"},{"model":"特斯拉"}],"myCar":{"car1":{"model":"小三轮儿"},"car2":{"model":"小黄车"}}}
        System.out.println(JsonUtil.pojoToString(person));

```

## 4. 关于输出

使用`dump()`方法将一个 map 格式化成字符串时, 得到的是类似json的东西, 如下

```
{key1: value1, key2: 123}
```

这种本质上其实也是`yaml`, 能够被正常解析, 但可能不是常规yaml文件的格式.
