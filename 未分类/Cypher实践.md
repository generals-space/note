# Neo4j Cypher实践

## 1. 查询

### 1.1 查询所有节点

将会得到所有节点, 包括有联系的节点以及孤岛节点.

```
match (a) return a
```

### 1.2 查询所有有联系的节点

```
match (a)-[*]-(b) return a, b
```

### 1.3 多重查询

```
match (a {name: 'general'}), (b {name: 'lianqia'}) return a, b
```

### 1.4 查询节点ID

```cypher
create (a: server {ip: '192.168.1.1'})
create (a: server {ip: '192.168.1.1'})
match (a: server {ip: '192.168.1.1'}) return a
```

在neo4j中, 我们可以创建多个`Label`相同, `Property`相同的节点, 而且可以被同时查询出来. 但是它们是不同的. 当我们需要一个像mysql中那样的自增主键唯一标识一个节点对象时怎么办?

答案是, neo4j本身在创建节点时会为其指定一个id, 也是自增ID, 查询方法是cypher提供的`ID()`函数.

```
create (a: server {ip: '192.168.1.1'}) return ID(a)
match (a: server {ip: '192.168.1.1'}) return a, ID(a)
match (a: server {ip: '192.168.1.1'}) where ID(a) = 目标节点id return a
```


## 2. 创建

### 2.1 创建联系

为两个已知节点创建联系, **联系必须有一个箭头存在**.

```
match (a {name: 'general'}), (b {name: 'lianqia'}) create (a)-[:relateTo]->(b)
```
