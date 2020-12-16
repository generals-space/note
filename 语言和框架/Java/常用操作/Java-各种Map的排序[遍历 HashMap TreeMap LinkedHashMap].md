# Java-各种Map的排序

参考文章

1. [java中实现HashMap中的按照key的字典顺序排序输出](https://www.cnblogs.com/bwgang/archive/2012/08/17/2947111.html)
    - 为`HashMap`排序的方法
2. [java 中HashMap排序](https://www.cnblogs.com/520yang/articles/4437852.html)
    - 比参考文章1更详细
3. [java中map里面的key是否可以按我们插入进去的顺序输出？](https://www.iteye.com/problems/69413)
    - `TreeMap`的key是自然顺序(如整数从小到大), 也可以指定比较函数.
    - `LinkedHashMap`内部有一个链表, 保持插入的顺序。迭代的时候, 也是按照插入顺序迭代, 而且迭代比HashMap快.

## 1. HashMap, TreeMap 字典序

```java
    HashMap<String, String> hashMap =new HashMap<String, String>();
    hashMap.put("b","0");
    hashMap.put("c","0");
    hashMap.put("d","0");
    hashMap.put("a","0");

    hashMap.forEach((k, v) -> {
        System.out.println("key: " + k + ", val: " + v);
    });
```

执行结果为

```
key: a, val: 0
key: b, val: 0
key: c, val: 0
key: d, val: 0
```

key为数字时也遵循从小到大的顺序, 同位数值也是按从小到来排的(比如11, 12, 23 这种, 但是非同位的, 如11会排在1前面).

不过我在实验时遇到一种特殊情况, 如下

```java
    HashMap<String, String> hashMap =new HashMap<String, String>();

    hashMap.put("2天","0");
    hashMap.put("3天","0");
    hashMap.put("1天","0");
    
    hashMap.forEach((k, v) -> {
        System.out.println("key: " + k + ", val: " + v);
    });
```

执行结果为

```
key: 3天, val: 0
key: 2天, val: 0
key: 1天, val: 0
```

## 2. LinkedHashMap 保持插入顺序

同样的代码, 将对象类型换成`LinkedHashMap`, 再次遍历时将会按照插入顺序排序.

```java
    LinkedHashMap<String, String> linkedHashMap =new LinkedHashMap<String, String>();
    linkedHashMap.put("b","0");
    linkedHashMap.put("c","0");
    linkedHashMap.put("d","0");
    linkedHashMap.put("a","0");

    linkedHashMap.forEach((k, v) -> {
        System.out.println("key: " + k + ", val: " + v);
    });
```

执行结果为

```
key: b, val: 0
key: c, val: 0
key: d, val: 0
key: a, val: 0
```
