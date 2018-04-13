# 参考文章

1. [python特性（八）：生成器对象的send方法](https://blog.csdn.net/hedan2013/article/details/56293173)

生成器对象是一个迭代器. 但是它比迭代器对象多了一些方法, 包括`send`, `throw`和`close`. **这些方法, 主要是用于外部与生成器对象的交互**. 

本文先介绍`send`方法. 

`send`方法有一个参数, 该参数指定的是上一次被挂起的`yield`语句的返回值. 这样说起来比较抽象, 看下面的例子. 

```py
def MyGenerator():
    value = yield 1
    value = yield value + 1
gen = MyGenerator()
print gen.next()
print gen.send(2)
print gen.send(3)
```

输出的结果如下

```
$  python gen.py 
1
3
Traceback (most recent call last):
  File "gen.py", line 7, in <module>
    print gen.send(3)
StopIteration
```

上面代码的运行过程如下. 

1. 当调用`gen.next()`方法时, python首先会执行`MyGenerator`方法到`yield 1`语句. 由于是一个`yield`语句, 因此方法的执行过程被挂起, 而`next`方法返回值为`yield`关键字后面表达式的值, 即为1. 

2. 当调用`gen.send(2)`方法时, python首先恢复`MyGenerator`方法的运行环境. 同时, 将表达式`value = yield 1`的返回值`value`定义为`send`方法参数的值, 即为3. 继续运行会遇到`value = yield value + 1`语句. 因此, `MyGenerator`方法再次被挂起. 同时, `send`方法也返回了, 它的返回值为`yield`关键字后面表达式的值, 也即`value`的值, 为3. 

3. 当调用`send(3)`方法时`MyGenerator`方法的运行环境. 同时, 将表达式`value = yield value + 1`的返回值定义为`send`方法参数的值, 即为3. 继续运行, `MyGenerator`方法执行完毕, 故而抛出`StopIteration`异常. 

总的来说, `send`方法和`next`方法唯一的区别是在执行`send`方法会首先把上一次挂起的`yield`语句的返回值通过参数设定, 从而实现与生成器方法的交互. 但是需要注意, 在一个生成器对象没有执行`next`方法之前, 由于没有`yield`语句被挂起, 所以执行`send`方法会报错. 例如

```py
def MyGenerator():
    value = yield 1
    value = yield value + 1
gen = MyGenerator()
print gen.send(2)
```

结果如下

```
$ python gen.py 
Traceback (most recent call last):
  File "gen.py", line 5, in <module>
    print gen.send(2)
TypeError: can't send non-None value to a just-started generator
```

当然, 下面的代码是可以接受的

```py
def MyGenerator():
    value = yield 1
    value = yield value + 1
gen = MyGenerator()
print gen.send(None)
```

因为当`send`方法的参数为`None`时, 它与`next`方法完全等价. 但是注意, 虽然上面的代码可以接受, 但是不规范. 所以, 在调用`send`方法之前, 还是先调用一次`next`方法为好. 

------

上面是参考文章1中的内容. 不过生成器中不一定会有`value = yield 值`这种表达式去专门取`yield`的返回值...