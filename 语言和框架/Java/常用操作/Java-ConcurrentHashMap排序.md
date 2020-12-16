# Java-ConcurrentHashMap排序

参考文章

1. [java中实现HashMap中的按照key的字典顺序排序输出](https://www.cnblogs.com/bwgang/archive/2012/08/17/2947111.html)
    - 为`HashMap`排序的方法
2. [java 中HashMap排序](https://www.cnblogs.com/520yang/articles/4437852.html)
    - 比参考文章1更详细

JDK: 1.8

参考文章1, 2 都给出了对`HashMap`对象的排序方法. 其实`HashMap`本身是按字典序排列的, key为数字时也遵循从小到大的顺序, 同位数值也是按从小到来排的(比如11, 12, 23 这种, 但是非同位的, 如11会排在1前面).

不过参考文章1中给出了一个特殊场景

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

除了这种场景的 HashMap, 还包括`ConcurrentHashMap`, 这种类型线程安全, 但是不保证遍历顺序, 对于这种情况, 需要进行一些额外的操作使其有序. 如下

```java
    HashMap<String, String> hashMap =new HashMap<String, String>();
    hashMap.put("2天","0");
    hashMap.put("3天","0");
    hashMap.put("1天","0");
    
    // 将 hashmap 中的 key 都取出来放到一个 list 中, 然后遍历 list 
    Collection<String> keyset = hashMap.keySet();
    List<String> list = new ArrayList<String>(keyset);
    Collections.sort(list);

    for(int i = 0; i < list.size(); i++){
        System.out.println("key: " + list.get(i) + ", val: " + hashMap.get(list.get(i)));
    }
```

执行结果为

```
key: 1天, val: 0
key: 2天, val: 0
key: 3天, val: 0
```
