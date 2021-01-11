# Java-Map字典遍历

参考文章

1. [Java中Map的 entrySet() 详解以及用法(四种遍历map的方式)](https://blog.csdn.net/q5706503/article/details/85122343)

```java
import java.util.Collection;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;
import java.util.HashMap;
import java.util.ArrayList;

public class Main{
    public static void main(String[] args) {
        HashMap<String, String> hashMap =new HashMap<String, String>();
        hashMap.put("2天","0");
        hashMap.put("3天","0");
        hashMap.put("1天","0");
    
        hashMap.forEach((k, v) -> {
            System.out.println("key: " + k + ", val: " + v);
        });

        // keySet() 得到目标字典中的 key 集合
        hashMap.keySet().forEach(k -> {
            System.out.println("key: " + k);
        });
        // entrySet() 返回的是 entry 集合, 一个 entry 就是一个 k-v 对,
        // 可以类比于 python 的 dict.items()
        hashMap.entrySet().forEach(entry -> {
            System.out.println("key: " + entry.getKey() + ", val: " + entry.getValue());
        });
        // Map.Entry 类型
        for(Map.Entry entry : hashMap.entrySet()){
            System.out.println("key: " + entry.getKey() + ", val: " + entry.getValue());
        }
    }
}

```
