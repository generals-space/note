# Python迭代器和生成器(二) - 生成器

参考文章

1. [迭代器与生成器](http://python.jobbole.com/84527/)

2. [完全理解 Python 迭代对象、迭代器、生成器](http://python.jobbole.com/87805/)

3. [生成器 - 廖雪峰的官方网站](https://www.liaoxuefeng.com/wiki/001374738125095c955c1e6d8bb493182103fac9270762a000/00138681965108490cb4c13182e472f8d87830f13be6e88000)

简单来讲, 生成器就是迭代器的语法糖...我是这么理解的...

生成器一定是迭代器, 但反之就不一定.

## 1. 生成器通过`yield`语句创建迭代器

```py
def container(start, end):
    while start < end:
        yield start
        start += 1
c = container(0, 5)
print(type(c))                  ## <type 'generator'>
print(next(c))                  ## 0
next(c)
for i in c:
    print(i)                    ## 2 3 4
```

生成器通过`yield`语句快速生成迭代器，省略了复杂的 `__iter__()` & `next()` 方式. 相比我们上一节的迭代器示例, 还需要创建class, 定义成员方法, 实例化对象后才能得到, 方便了许多.

简单来说，`yield`语句可以让普通函数变成一个生成器，并且相应的`next()` 方法返回的是yield后面的值(上述代码中是`start`变量)。一种更直观的解释是：程序执行到`yield`会返回值并暂停，再次调用`next()`时会从上次暂停的地方继续开始执行. 而一般for语句的每次循环都相当于执行一次`next()`而已.

```py
def gen():
    yield 5
    yield "Hello"
    yield "World"
    yield 4
for i in gen():
    print(i)
```

```
5
Hello
World
4
```
