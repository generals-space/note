# Python装饰器与传参

参考文章

1. [六种装饰器示例](http://www.jianshu.com/p/a405814f8786)

2. [Python装饰器中，装饰函数带参数，被装饰函数的地址是如何传入的？](https://www.zhihu.com/question/64433992)

## 1. 被装饰函数带参数

```py
from time import ctime, sleep

def timefun(func):
    def wrappedfunc(*args, **kwargs):
        print("%s called at %s"%(func.__name__, ctime()))
        func(*args, **kwargs)
    return wrappedfunc

@timefun
def foo(a, b, c,num):
    print(a+b+c)
    print(num)

foo(3,5,7,num=10)
sleep(2)
foo(2,4,9,num=20)
```

执行它, 输出如下

```
运行结果如下：
foo called at Fri Jul 14 19:52:46 2017
15
10
foo called at Fri Jul 14 19:52:48 2017
15
```

因为`@timefun`的作用等价于`foo = timefun(foo)`, 所以在调用`foo`时, 实际上调的是`wrappedfunc`, 接收参数的也是它, 只要`wrappedfunc`能正确接收到传入的参数, 再代为传入`foo`就可以.

## 2. 装饰器带参数

这真是个欠打的问题...

参考文章1的例5给出了如下代码.

```py
from time import ctime, sleep

def timefun_arg(pre="hello"):
    def timefun(func):
        def wrappedfunc():
            print("%s called at %s %s"%(func.__name__, ctime(), pre))
            return func()
        return wrappedfunc
    return timefun

@timefun_arg("wangcai")
def foo():
    print("I am foo")

@timefun_arg("python")
def too():
    print("I am too")

foo()
too()
```

...装饰器函数里多了一层, 好像变得很复杂的样子.

最初看到这个示例时, 疑惑于`timefun_arg`如何接收`func`变量并把它传递给`timefun`的. 我还尝试了一下在`timefun`之前打印一下`func`的值, 结果报错了.

```py
def timefun_arg(pre="hello"):
    print(func)         ## NameError: global name 'func' is not defined
    def timefun(func):
    ...
```

按照参考文章2中`黄哥`的回答, 认识到`timefun_arg("python")`本身就是一个函数调用了...它返回了一个函数`timefun`, 正好用于装饰器. 简直了...

## 3. 为类的实例方法添加装饰器

因为修饰了类成员函数, 所以装饰器函数能取得类成员变量与方法的访问权限(这正是闭包的作用), 这里, 我们在执行类成员函数`now()`之前修改类成员变量`self.i`的值.

```py
def log(func):
    def wrapper(self, a):
        print 'arguments list : %s' % (a)
        self.i = a
        return func(self, a)
    return wrapper
    
class Date():
    @log
    def now(self, a):
        print self.i
        print '2013-12-25'

if __name__ == '__main__':
    date = Date()
    date.now(1)
```
