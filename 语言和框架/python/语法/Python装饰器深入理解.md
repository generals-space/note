# Python装饰器深入理解

```python
def log(func):
    print 'entered log'
    def wrapper():
        print 'call %s():' % func.__name__
        return func()
    return wrapper

@log
def now():
    print '2013-12-25'

if __name__ == '__main__':
    now()

```

过程分析:

```
## 经过装饰器修改的函数名now可以看作是log(now)
now         = log(now)
## 之后是函数执行阶段, 调用now().
## 函数名(参数列表){内部操作}()有些类似js中的函数操作. 不过在python中没法这样实现, 因为python是根据缩进来判断函数结束的
## 在js中可以尝试如下: function show(a){console.log(a)}(2);
now(){}()   = log(now){...}()
            ## 在log函数作用域内已取得func = now, 然后func的值可以被内部的wrapper函数取到了
            ## 执行log函数内部操作
            = print 'entered log'
## 函数输出
>>> entered log
            ## 返回wrapper, wrapper函数将接力
            = wrapper(){...}()
            ## 然后开始执行wrapper内的操作
            = print 'call now'
## 函数输出
>>> call now():
            ## 然后返回func, 这时func为now, 即, 开始执行原now函数的操作, 装饰器函数完成
            = now(){...}()
            = print '2013-12-25'
## 函数输出
>>> 2013-12-25
```

示例2: 尝试捕获并修改被装饰函数的参数列表

```python 
#!/bin/python

def log(func):
    def wrapper(a):
        print 'arguments list : %s' % (a)
        a = a + 1
        return func(a)
    return wrapper

@log
def now(a):
    print a
    print '2013-12-25'

if __name__ == '__main__':
    now(1)
```

过程分析

```
now             = log(now)
now(a){}(1)     = log(now){...}(1)
                ## 这种函数执行方式貌似只存在于js中, 不过原理是一样的, (1)可以传入函数参数并执行
                = wrapper(a){...}(1)
                = print 'arguments list : %s' % (a)
## 函数输出
arguments list : 1
                = a = a + 1, 即 a = 2
                = func(a){...}(2)
                = now(a){...}(2)
                = print a
## 函数输出
2
                = print '2013-12-25'
## 函数输出
2013-12-25
```

示例3: 类成员函数的装饰器

类成员函数的定义都显式包含一个`self`参数, 所以装饰器返回的函数, 其参数列表也必须包含这个self.

```python
def log(func):
    def wrapper(self, a):
        print 'arguments list : %s' % (a)
        a = a + 1
        return func(self, a)
    return wrapper
class Date():
    @log
    def now(self, a):
        print a
        print '2013-12-25'

if __name__ == '__main__':
    date = Date()
    date.now(1)

```

```
date = Date()
date.now(1) = log(date.now){}(1)
            ## 这里尤其要注意, 因为装饰器最终要返回原来的函数, 所以wrapper一定要获得原类成员函数的self参数, 以传给func函数, 这样, 返回的func才可以接受self参数
            = wrapper(self, a){...}(1)
            = print 'arguments list : %s' % (a)
## 函数输出
arguments list : 1
            = a = a + 1, 即a = 2
            = func(self, a), 即date.now(a)
            = print a
## 函数输出
2
            = print '2013-12-25'
## 函数输出
2013-12-25
```

## 示例4: 扩展, 为类成员函数执行前修改类成员变量.

因为修饰了类成员函数, 所以装饰器函数能取得类成员变量与方法的访问权限(这正是闭包的作用), 这里, 我们在执行类成员函数now()之前修改类成员变量`self.i`的值.

```python
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