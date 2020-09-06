参考文章

1. [Parsing a YAML document with a map at the root using snakeYaml](https://stackoverflow.com/questions/28551081/parsing-a-yaml-document-with-a-map-at-the-root-using-snakeyaml)
2. [How to load a list of custom objects with SnakeYaml](https://stackoverflow.com/questions/56187845/how-to-load-a-list-of-custom-objects-with-snakeyaml)

## 1. 引言

最近要处理一份格式比较特殊的`yaml`配置, 很棘手, 所以找了找能不能通过高端一点的手段解决, 在这个过程中发现了自定义`constructor`的方法.

首先, 代解析的yaml格式如下

```yaml
name: general
age: 24
gender: {xingbie}
```

其中`{xingbie}`是用一个大括号包裹的, 类似占位符(或是shell中`${var}`)的格式, 在使用`snakeyaml`解析的过程时候, 还没有办法得到这个变量的值.

等到过一段时间, 知道了这个变量的值后, 会使用shell脚本对这个变量进行替换.

```
sed -i 's/{xingxie}/男/g' yaml
```

但是在使用 snakeyaml 对这个格式的 yaml 进行解析的时候, snakeyaml 发现`gender`字段的值用`{}`包裹, 就会将 {xingbie} 也当作一个字典对象处理, 然后将`xingbie`当作一个`key`, 但又找不到`value`, 就当成`null`了.

```java
		Yaml yaml = new Yaml();
		Object object = yaml.load(content.toString());
		Map<String, Object> map = (Map<String, Object>)object;
        // {name=general, age=24, gender={xingbie=null}}
		System.out.println(map);
```

按照这种解析得到的结果, 用`sed`就不好再次进行替换了.

我想通过`constructor`让`snakeyaml`把以`{`开头, `}`结尾, 且中间没有`:`的字段, 直接当成字符串, 就不要再尝试解析ta的键值了.

## 2. constructor

按照参考文章1中, `Narcoleptic Snowman`给出的答案, 我自己编写了一个`constructor`.

```java
import org.yaml.snakeyaml.constructor.Constructor;
import org.yaml.snakeyaml.nodes.MappingNode;
import org.yaml.snakeyaml.nodes.Node;
import org.yaml.snakeyaml.nodes.NodeTuple;
import org.yaml.snakeyaml.nodes.ScalarNode;
import org.yaml.snakeyaml.nodes.Tag;

public class YamlConstructor<T> extends Constructor {
    private Map<String, Object> clazz;

    /*
    private Class<T> clazz;
    public YamlConstructor(Class<T> clazz) {
        this.clazz = clazz;
    }
    */

    @Override
    protected Object constructObject(Node node) {
        // System.out.println("======================");
        /*
            // eyaml.nodes.ScalarNode (tag=tag:yaml.org,2002:str, value=比亚迪)>> })>
            // node 对象有 value 成员, 但是没有能直接获取 value 的方法.
            System.out.println(node);
            // tag:yaml.org,2002:int, tag:yaml.org,2002:str, 
            // tag:yaml.org,2002:seq, tag:yaml.org,2002:map 等
            System.out.println(node.getTag());
            // class java.lang.Object ...还没见过其他类型
            System.out.println(node.getType());
            // scalar(str), sequence(seq), mapping(map)
            System.out.println(node.getNodeId());
        */
        // 我们只处理字典对象, 因为我们的目标是 key: {value} 这样的键值对.
        if(node.getTag() == Tag.MAP) {
            MappingNode mNode = (MappingNode)node;
            for (NodeTuple item : mNode.getValue()) {
                Node keyNode = item.getKeyNode();
                Node valNode = item.getValueNode();
                String keyNodeType = item.getKeyNode().getTag().getClassName();
                String valNodeType = item.getValueNode().getTag().getClassName();
                // 这里的 if 条件中, 第1个条件是 value, 类型为 str, 第2个条件是 null
                if(keyNodeType.equals("str") && valNodeType.equals("null")) {
                    System.out.println(JsonUtil.pojoToString(item.getKeyNode()));
                    System.out.println(JsonUtil.pojoToString(item.getValueNode()));

                    // valNode.setType(String.class);
                    // 这里的 snippet 的值为 'key: {value}'
                    System.out.println(keyNode.getStartMark().get_snippet());
                    // 如果发现了这样的情况, 就只返回 keyNode, 即 key 本身, 而不再包括 null 的部分了.
                    return super.constructObject(keyNode);
                }
            }
        }

        // In all other cases, use the default constructObject.
        return super.constructObject(node);
    }
}

```

对于这个`constructor`, 按照如下的方法使用.

```java
		Yaml yaml = new Yaml(new YamlConstructor());
		// Yaml yaml = new Yaml(new YamlConstructor<>(Map.class));
		Object object = yaml.load(content.toString());
		System.out.println(object);
```

其中最重要的是构造函数的定义方式.

因为这个`YamlConstructor`类并没有确切的字段, 不像上一篇文档中的`Person`与`Car`, 毕竟如果这个类的字段能够定下来, 我也不需要这么麻烦了.

`Class<T>`那个可以用, 但其实不需要那么复杂, 直接定义一个`Map<String, Object>`类型的成员即可.
