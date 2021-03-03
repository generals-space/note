# Python装饰器向被装饰函数注入变量

参考文章

1. [How to inject variable into scope with a decorator in python](https://stackoverflow.com/questions/17862185/how-to-inject-variable-into-scope-with-a-decorator-in-python)

常规的装饰器可以在被装饰函数执行前后执行一些代码, 而且采用的是闭包的形式, 但是一般无法影响到被装饰函数的执行过程, 只能在执行前可以改写参数, 或是处理一下执行结果.

```py
#!/usr/bin/python
#!coding:utf-8

def decorator(func):
    ## 定义一个内嵌的包装函数, 给传入的函数加上计时功能的包装
    def wrappedFunc():
        print('before...')
        func()
        print('after...')
    return wrappedFunc

@decorator
def sum1():
    sum = 1 + 2
    print(sum)

sum1()
```

最初的想法是, 在执行被装饰函数之前定义变量, 这样在执行目标函数时, 理论上应该能获取到这个变量, 毕竟被装饰函数在最内层的作用域, 像下面这样.

```py
def decorator(func):
    def wrappedFunc():
        counter = 1
        func()
    return wrappedFunc

@decorator
def show():
    print(counter)

show()
```

执行时却报如下错误

```
Traceback (most recent call last):
  File "test.py", line 11, in <module>
    show()
  File "test.py", line 4, in wrappedFunc
    func()
  File "test.py", line 9, in show
    print(counter)
NameError: global name 'counter' is not defined
```

...虽然可以通过向`func`将我们定义的变量当作参数一样传入, 但这并不是我们真正的目的. 参考文章1中(找了很久, 只有这一篇提到我们的这种作法)解释了原因.

> Scoped names (closures) are determined at compile time, you cannot add more at runtime.

> 闭包的作用域在编译时已经确定, 无法在运行时添加.

最高票答案提供了一个能够达到目的的方法, 就是使用被装饰函数对象的`__globals__`属性(注意, 是属性, 不是方法)手动将我们自定义的变量传入. 这其实与传入参数的方法类似.

```py
def decorator(func):
    def wrappedFunc():
        counter = 1
        g = func.__globals__  # use f.func_globals for py < 2.6
        sentinel = object()

        oldvalue = g.get('counter', sentinel)
        g['counter'] = counter

        try:
            res = func()
        finally:
            if oldvalue is sentinel:
                del g['counter']
            else:
                g['counter'] = oldvalue
        return res
    return wrappedFunc

@decorator
def show():
    print(counter)

show()          ## 输出为1
```

其中`func.__globals__`是函数`func`本身的全局命名空间, 上述操作其实是将`counter`变量写入到`func`全局作用域中.

python拥有`lexically scoped(词法作用域)`, 没有`clean`的方式来完成这样的要求...

关于**词法作用域**, 有点深奥, 没能理解, 先放一放.