# Python内存管理机制(一)引用计数与垃圾回收

参考文章

1. [Python内存管理机制及优化简析](http://kkpattern.github.io/2015/06/20/python-memory-optimization-zh.html)
2. [python之内存调试](https://blog.csdn.net/lj1404536198/article/details/80900549)
3. [使用 Gc、Objgraph 干掉 Python 内存泄露与循环引用！](http://www.cnblogs.com/xybaby/p/7491656.html)
    - 在阅读了参考文章1和2, 对python的内存管理机制, 包括引用计数, 垃圾回收(双向链表, 分代回收)等有了大概的概念后, 就可以阅读参考文章3了. 
    - 参考文章3给出了各种概念和场景的简洁并准备的描述, 尤其是有了参考文章1和2的基础后, 会觉得参考文章3的逻辑特别清晰.

Python有两种共存的内存管理机制: **引用计数**和**垃圾回收**. 引用计数是一种非常高效的内存管理手段, 当一个对象被引用时其引用计数增加1, 当其不再被一个变量引用时则计数减1. 当引用计数等于0时对象被删除.

为了方便解释, 本文使用了`gc`模块来辅助展示内存中的Python对象以及垃圾回收器的工作情况. 本文中具体使用到的接口包括:

- `gc.disable()`: 暂停自动垃圾回收.
- `gc.collect()`: 执行一次完整的垃圾回收, 返回垃圾回收所找到无法到达的对象的数量.
- `gc.set_threshold()`: 设置垃圾回收的阈值.
- `gc.set_debug()`: 设置垃圾回收的调试标记. 调试信息会被写入std.err.

同时我们还使用了`objgraph`库, 本文中具体使用到的接口包括:

`objgraph.count(typename)`: 对于给定类型typename, 返回垃圾回收器正在跟踪的对象个数.

> `objgraph`可以通过命令`pip install objgraph`安装.

## 1. 引用计数

```py
import gc
import objgraph

## 这一句主要是屏蔽gc对内存的影响, 实际上就算没有这句, 本例的结果也是相同的.
gc.disable()

class A():
    pass
class B():
    pass

def test1():
    a = A()
    b = B()
    print(objgraph.count('A')) ## 1
    print(objgraph.count('B')) ## 1

test1()

print(objgraph.count('A')) ## 0
print(objgraph.count('B')) ## 0
```

在`test1`中, 我们分别创建了类`A`和类`B`的对象, 并用变量`a`, `b`引用起来. 当`test1`调用结束后`objgraph.count('A')`返回0, 意味着内存中A的对象数量 没有增长. 同理B的对象数量也没有增长. 注意我们通过`gc.disable()`关闭了 Python的垃圾回收, 因此`test1`中生产的对象是在函数调用结束引用计数为0时被自动删除的.

引用计数的一个主要缺点是 **无法自动处理循环引用**.

### 循环引用

```py
import gc
import objgraph

## 这一句主要是屏蔽gc对内存的影响, 实际上就算没有这句, 本例的结果也是相同的.
## gc.disable()

class A():
    pass
class B():
    pass

def test1():
    a = A()
    b = B()
    a.child = b
    b.child = a
    print(objgraph.count('A')) ## 1
    print(objgraph.count('B')) ## 1

test1()

print(objgraph.count('A')) ## 1
print(objgraph.count('B')) ## 1

gc.collect()
print(objgraph.count('A')) ## 0
print(objgraph.count('B')) ## 0

```

与`test1`相比, `test2`的改变是将`A`和`B`的对象通过`child`和`parent`相互引用起来, 这就形成了一个循环引用. 当`test2`调用结束后, 表面上我们不再引用两个对象, 但由于两个对象相互引用着对方, 因此引用计数不为0, 则不会被自动回收. 

更糟糕的是由于现在没有任何变量引用他们, 我们无法再找到这两个变量并清除. Python使用垃圾回收机制来处理这样的情况. 执行`gc.collect()`, Python垃圾 回收器回收了两个相互引用的对象, 之后A和B的对象数又变为0.

## 2. 垃圾回收

本节将简单介绍Python的垃圾回收机制. 这篇文章[Garbage Collection for Python](http://arctrix.com/nas/python/gc/)以及Python垃圾回收源码中的注释进行了更详细的解释.

### 垃圾回收原理

在Python中, 所有能够引用其他对象的对象都被称为 **容器(container)**, 因此只有容器之间才可能形成循环引用. Python的垃圾回收机制利用了这个特点来寻找需要被释放的对象. 为了记录下所有的容器对象, Python将每一个容器都链到了一个双向链表中, 之所以使用双向链表是为了方便快速的在容器集合中插入和删除对象. 

有了这个维护了所有容器对象的双向链表以后, Python在垃圾回收时使用如下步骤来寻找需要释放的对象:

1. 为链表中的每一个新的容器对象, 设置一个`gc_refs`值, 表示该对象的引用计数值.
2. 进行垃圾回收时, 对于链表中每一个容器对象, 找到所有其引用的对象, 将被引用对象的gc_refs值减1.
执行完步骤2以后所有gc_refs值还大于0的对象都被非容器对象引用着, 至少存在一个非循环引用. 因此 不能释放这些对象, 将他们放入另一个集合.
在步骤3中不能被释放的对象, 如果他们引用着某个对象, 被引用的对象也是不能被释放的, 因此将这些 对象也放入另一个集合中.
此时还剩下的对象都是无法到达的对象. 现在可以释放这些对象了.

### __del__

### 分代回收

python中应用了 **分代回收机制**. 简单来说就是, 将存在时间短的对象容易死掉, 而老年的对象不太容易死, 这叫做 **弱代假说(weak generation hypothesis)**. 

这也很好理解, 一般生命周期长的对象往往是全局变量, 而短的多为局部变量或者临时定义的变量. 那么, 我们把当前的对象作为第0代, 我们每当allocation比deallocation多到某个阈值时, 就对这些对象做一次检查和清理, 没有被清理的那些就存活下来, 进入第1代, 第一代检查做若干次后, 对1代清理, 存活下来的进入第2代, 第二代也是如此. 这样就实现了分代回收的操作. 

------

对于垃圾回收, 有两个非常重要的术语, 那就是`reachable/unreachable`与`collectable/uncollectable`. 

`(un)reachable`是针对python对象而言, 如果从根集(root)能到找到对象, 那么这个对象就是`reachable`, 与之相反就是`unreachable`. 事实上就是只存在于循环引用中的对象, 无法通过变量名找到, 就是`unreachable`状态. Python的垃圾回收就是针对`unreachable`对象. 

而`(un)collectable`是针对`unreachable`对象而言, 正常情况下发生循环引用的对象能被gc回收, 那么就是`collectable`; 如果循环引用中的对象定义了`__del__`,  那么就是`uncollectable`. Python垃圾回收对`uncollectable`对象无能为力, 会造成事实上的内存泄露. 

## 总结

1. python使用引用计数和垃圾回收来释放(free)Python对象
2. 引用计数的优点是原理简单、将消耗均摊到运行时；缺点是无法处理循环引用
3. Python垃圾回收用于处理循环引用, 但是无法处理循环引用中的对象定义了`__del__`的情况, 而且每次回收会造成一定的卡顿
4. `gc`是python垃圾回收机制的接口模块, 可以通过该模块启停垃圾回收、调整回收触发的阈值、设置调试选项
5. 如果没有禁用垃圾回收, 那么Python中的内存泄露有两种情况：要么是对象被生命周期更长的对象所引用, 比如global作用域对象；要么是循环引用且造成循环引用的对象存在`__del__`
6. 使用`gc`、`objgraph`可以定位内存泄露, 定位之后, 解决很简单
7. 垃圾回收比较耗时, 因此在对性能和内存比较敏感的场景也是无法接受的, 如果能解除循环引用, 就可以禁用垃圾回收. 
8. 使用`gc`的`DEBUG`选项可以很方便的定位循环引用, 解除循环引用的办法要么是手动解除, 要么是使用`weakref`
