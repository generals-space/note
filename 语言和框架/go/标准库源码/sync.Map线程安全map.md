# sync.Map线程安全map

参考文章

1. [你不得不知道的sync.Map源码分析](https://segmentfault.com/a/1190000015242373)
    - sync.Map冗余数据结构read, dirty两个map的作用解释
    - 图解read, dirty两个map读写分离的原则
    - 提供了`Len()`的实现代码

如果你接触过大Java, 那你一定对`CocurrentHashMap`利用**锁分段技术**增加了锁的数目, 从而使争夺同一把锁的线程的数目得到控制的原理记忆深刻.  

那么Golang的`sync.Map`是否也是使用了相同的原理呢？`sync.Map`的原理很简单, 使用了空间换时间策略, 通过冗余的两个数据结构(`read`、`dirty`),实现加锁对性能的影响. 

通过引入两个`map`将读写分离到不同的`map`, 其中`read map`提供并发读和已存元素原子写, 而`dirty map`则负责读写.  这样read map就可以在不加锁的情况下进行并发读取, 当read map中没有读取到值时,再加锁进行后续读取,并累加未命中数,当未命中数大于等于dirty map长度,将dirty map上升为read map. 

从之前的结构体的定义可以发现, 虽然引入了两个map, 但是底层数据存储的是指针, 指向的是同一份值. 

## 结构体

### Map

```go
type Map struct {
    mu Mutex    //互斥锁，用于锁定dirty map

    read atomic.Value //优先读map,支持原子操作，注释中有readOnly不是说read是只读，而是它的结构体。read实际上有写的操作

    dirty map[interface{}]*entry // dirty是一个当前最新的map，允许读写

    misses int // 主要记录read读取不到数据加锁读取read map以及dirty map的次数，当misses等于dirty的长度时，会将dirty复制到read
}
```

### readOnly

readOnly 主要用于存储，通过原子操作存储在Map.read中元素。

```go
type readOnly struct {
    m       map[interface{}]*entry
    amended bool // 如果数据在dirty中但没有在read中，该值为true,作为修改标识
}
```

### entry

```go
type entry struct {
    // nil: 表示为被删除，调用Delete()可以将read map中的元素置为nil
    // expunged: 也是表示被删除，但是该键只在read而没有在dirty中，这种情况出现在将read复制到dirty中，即复制的过程会先将nil标记为expunged，然后不将其复制到dirty
    //  其他: 表示存着真正的数据
    p unsafe.Pointer // *interface{}
}
```

## 读写流程

### 1.

开始时`sync.Map`写入数据

```
X=1
Y=2
Z=3
```

`dirty map`主要接受写请求, `read map`没有数据, 此时`read map`与`dirty map`的数据如下.

- Read Map: []
- Dirty Map: [X = 1; Y = 2; Z = 3]

### 2. 

读数据时要从`read map`中读取, 此时`read map`并没有数据, 读取未命中会使`misses`加1.

读取`read map`失败的后续操作?

------

当`misses >= len(dirty map)`时, 将`dirty map`提升为`read map`并将`dirty map`清空, 且`misses`置0, `read map`结构中的`amended`成员值为false.

此时`read map`与`dirty map`数据如下.

- Read Map: [X = 1; Y = 2; Z = 3]
- Dirty Map: []

> 注意: 提升操作其实是对`dirty map`进行地址拷贝.

### 3. 

现在有需求将`Z`元素修改为`4`, `sync.Map`会直接修改`read map`中的元素.

- Read Map: [X = 1; Y = 2; Z = 4]
- Dirty Map: []

### 4.

新增元素`K = 5`, 需要操作`dirty map`, `sync.Map`会判断`read map`中`amended`成员的值做合适的操作.

- 如果其值为`false`, 则将`read map`中的内容拷贝到`dirty map`(但不清空`read map`), 并且将`amended`值置为`true`, 然后在`dirty map`中添加新值.
- 如果其值为`true`, 直接向`dirty map`中添加新值.

> 可以说, `admended`为`true`时, 表示了`dirty map`中包含着`read map`中没有的键, 前者是后者的超集. `amended`为`false`时, 一般`dirty map`处于提升为`read map`后的"空"状态.

此时`read map`与`dirty map`数据如下.

- Read Map: [X = 1; Y = 2; Z = 3]
- Dirty Map: [X = 1; Y = 2; Z = 3; K = 5]

之后如果`misses`值又超过了`dirty map`长度, 将会再次提升

- Read Map: [X = 1; Y = 2; Z = 3; K = 5]
- Dirty Map: []

### 5. 

如果需要删除元素`Z`, 需要分几种情况.

#### 5.1 

元素刚写入`dirty map`且未升级为`read map`, 直接调用golang的内置函数`delete`删除.


- Read Map: []
- Dirty Map: [X = 1; Y = 2; ~~Z = 3~~; K = 5]

#### 5.2 

`read map`中存在该元素且`dirty map`为空(此时`amended`为`false`), 直接将`read map`中的元素置为`nil`.

- Read Map: [X = 1; Y = 2; ~~Z = 3;~~ K = 5]
- Dirty Map: []

#### 5.3

`read map`和`dirty map`中同时存在该元素(此时`amended`应为`true`), 将`read map`中的元素置为nil即可. 因为`read map`和`dirty map`中使用的均为元素地址, 所以均被置为nil.

- Read Map: [X = 1; Y = 2; ~~Z = 3~~]
- Dirty Map: [X = 1; Y = 2; ~~Z = 3;~~ K = 5]

## 优化点

1. 空间换时间。通过冗余的两个数据结构(read、dirty),实现加锁对性能的影响。
2. 使用只读数据(read)，避免读写冲突。
3. 动态调整，miss次数多了之后，将dirty数据提升为read。
4. double-checking（双重检测）。
5. 延迟删除。 删除一个键值只是打标记，只有在提升dirty的时候才清理删除的数据。
6. 优先从read读取、更新、删除，因为对read的读取不需要锁。

## 总结

经过了上面的分析可以得到,sync.Map并不适合同时存在大量读写的场景,大量的写会导致read map读取不到数据从而加锁进行进一步读取,同时dirty map不断升级为read map。 从而导致整体性能较低,特别是针对cache场景.针对append-only以及大量读,少量写场景使用sync.Map则相对比较合适。