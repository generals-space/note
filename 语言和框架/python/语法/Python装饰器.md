# Python装饰器

参考文章

1. [python装饰器通俗易懂的解释！](http://www.cnblogs.com/songyue/p/5196809.html)

2. [理解Python中的装饰器](http://www.cnblogs.com/rollenholt/archive/2012/05/02/2479833.html)

3. [Python装饰器学习（九步入门）](http://www.cnblogs.com/rhcad/archive/2011/12/21/2295507.html)

> 装饰器是一个很著名的设计模式，经常被用于有切面需求的场景，较为经典的有插入日志、性能测试、事务处理等。装饰器是解决这类问题的绝佳设计，有了装饰器，我们就可以抽离出大量函数中与函数功能本身无关的雷同代码并继续重用。概括的讲，装饰器的作用就是为已经存在的对象添加额外的功能。

装饰器在java spring中应该被称为`拦截器`, 意思是在执行某一操作A之前先进行另一操作B, 而B是比较通用的操作. 比如类型检查, 输出日志等. 编写成装饰器形式方便其他函数使用. 

## 关于装饰器通俗易懂的解释

小P闲来无事，随便翻看自己以前写的一些函数，忽然对一个最最最基础的函数起了兴趣: 

```py
#!/usr/bin/python
def sum1():
    sum = 1 + 2
    print(sum)
sum1()
```


小P想看看这个函数执行用了多长时间，所以写了几句代码插进去了. 运行之后，完美~~

```py
#!/usr/bin/python
import time

def sum1():
    start = time.time()
    sum = 1+2
    print(sum)
    end = time.time()
    print("time used:",end - start)
sum1()
```

可是随着继续翻看，小P对越来越多的函数感兴趣了，都想看下他们的运行时间如何，难道要一个一个的去改函数吗？当然不是！我们可以考虑重新定义一个函数timeit，将sum1的`引用`传递给他，然后在timeit中调用sum1并进行计时，这样，我们就达到了不改动sum1定义的目的，而且，不论小P看了多少个函数，我们都不用去修改函数定义了！

```py
#!/usr/bin/python
import time

def sum1():
    sum = 1+2
    print(sum)
def timeit(func):
    start = time.time()
    func()
    end = time.time()
    print("time used:",end - start)
timeit(sum1)
```

乍一看，没啥问题，可以运行！但是还是修改了一部分代码，把sum1() 改成了timeit(sum1)。这样的话，如果sum1在N处都被调用了，你就不得不去修改这N处的代码。所以，我们需要timeit(sum1)具有和sum1()一样的效果，于是将timeit赋值给sum1。可是timeit是有参数的，我们还需要timeit(sum1)具有和sum1()一样的调用方式，所以需要找个方法去统一参数，将timeit(sum1)的返回值（计算运行时间的函数）赋值给sum1。

```py
#!/usr/bin/python
#!coding:utf-8
import time

def sum1():
    sum = 1+2
    print(sum)
def timeit(func):
    ## 定义一个内嵌的包装函数, 给传入的函数加上计时功能的包装
    def wrappedFunc():
        start = time.time()
        func()
        end = time.time()
        print("time used:",end - start)
    return wrappedFunc

sum1 = timeit(sum1)
sum1()
```

这样一个简易的装饰器就做好了，我们只需要在定义sum1以后调用sum1之前，加上`sum1= timeit(sum1)`，就可以达到计时的目的，这也就是装饰器的概念，看起来像是sum1被timeit装饰了！

Python提供了一个语法糖来降低字符输入量。

```py
#!/usr/bin/python
#!coding:utf-8
import time

def timeit(func):
    ## 定义一个内嵌的包装函数, 给传入的函数加上计时功能的包装
    def wrappedFunc():
        start = time.time()
        func()
        end = time.time()
        print("time used:",end - start)
    return wrappedFunc

@timeit
def sum1():
    sum = 1+2
    print(sum)

sum1()
```

重点关注`@timeit`这一行，在定义上加上这一行与另外写`sum1 = timeit(sum1)`完全等价。`@`没有额外的作用, 除了字符输入少了一些，只有一个好处: 看上去更有`装饰`的感觉。

在这个例子中，函数进入和退出时需要计时，这被称为一个横切面(Aspect)，这种编程方式被称为面向切面的编程AOP(Aspect-Oriented Programming)。与传统编程习惯的从上往下执行方式相比较而言，像是在函数执行的流程中横向地插入了一段逻辑。在特定的业务领域里，能减少大量重复代码。面向切面编程还有相当多的术语，这里就不多做介绍，感兴趣的话可以去找找相关的资料。

执行这种被装饰过的函数时, 调用的实际上是`wrappedFunc`, 传入的参数也是传入`wrappedFunc`然后调用`func`时把这些参数传给`func`的.
