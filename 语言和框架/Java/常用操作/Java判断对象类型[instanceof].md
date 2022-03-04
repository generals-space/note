# Java判断对象类型

参考文章

1. [Java:判断Object类型,并返回长度](https://zhuanlan.zhihu.com/p/136997161)

```java
		System.out.println("cluster.name");
		System.out.println(map.get("cluster.name") instanceof String);
		System.out.println(map.get("cluster.name") instanceof Integer);
		System.out.println("http.port");
		System.out.println(map.get("http.port") instanceof String);
		System.out.println(map.get("http.port") instanceof Integer);
		System.out.println("bootstrap.memory_lock");
		System.out.println(map.get("bootstrap.memory_lock") instanceof String);
		System.out.println(map.get("bootstrap.memory_lock") instanceof Boolean);
		System.out.println("cluster.initial_master_nodes");
		System.out.println(map.get("cluster.initial_master_nodes") instanceof String);
		System.out.println(map.get("cluster.initial_master_nodes") instanceof List);
```

`instanceof`判断`List`, `Map`, 无法具体到`List<String>`还是`List<Integer>`.
