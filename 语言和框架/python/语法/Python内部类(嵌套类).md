# Python内部类, 嵌套类

参考文章

1. [Python学习笔记二十---- 对象高级 - 163博客](http://blog.163.com/qimeizhen8808@126/blog/static/16511951820127220173667)

2. [python内部类](http://blog.csdn.net/u013551220/article/details/19154047)

3. [Python有没有"内部类"这一说法?内部类如何访问外部类的成员? - CSDN论坛](http://bbs.csdn.net/topics/90479518)

## 1. 内部类应用场景

我们知道类表示一种对象的模板, 书中教程都包括`People`, `Bird`等例子. 以`Car`为例, 它可能有`引擎`, `轮胎`等配件, 如果再细分, 轮胎也可以有自己的属性(如轮毂)和方法(转动)等.

mmp, 例子举的不好...简单粗暴地说, 就当是类成员对象也有自己的一组成员时, 需要如何构建? 

有一种方法就是把这个内部对象简单的当作一个字典对象处理, 但是如果这个对象需要有自己的成员方法呢? 

首先是预热活动, python不像js那样可以在定义对象字面量时直接赋值为函数.

```js
var obj = {
    func: function(){console.log(123)}
}
```

...放弃吧, 在python里, 字典成员值作为函数调用我们只能这么干.

```py
#!/usr/bin/env python
#!encoding: utf-8

def say():
    print('hello world')

dic = {
    'func': say
}

dic['func']()           ## hello world
```

python需要先定义一个函数, 再在字典对象中将其赋值. 

这还算好的...如果是在类中呢? 我们需要先定义内部函数, 才能给这个子对象使用...呵呵哒

```py
#!/usr/bin/env python
#!encoding: utf-8

class Car():
    def __init__(self):
        self.door = 'a door'
    def run():
        print('i am running')
    wheel = {
        'run': run
    }

car = Car()
car.wheel['run']()      ## i am running
```

...不觉得别扭吗? 这样封装还有什么意义?

我个人觉得内部类就是为了这种情况出现的. 可以将子类封装的更完美.

```py
#!/usr/bin/env python
#!encoding: utf-8

class Car():
    def __init__(self):
        self.door = 'a door'
    class Wheel():
        def run(self):
            print('i am running')

car = Car()             ## 实例化外部类
wheel = car.Wheel()     ## 实例化内部类
wheel.run()             ## i am running
```

其实直接实例化内部类也可以, 通过类名直接调用即可. 但是这样的话, 你可能无法在内部类中取得外部类实例的成员对象. 在很多情况下, 这都是常见需求.

```py
wheel2 = Car.Wheel()    
wheel2.run()            ## i am running
```

我想类似这种应用场景不仅限于python, 内部类的设计与实现应该就是为了应对这种需要的.

## 2. 深入理解

上面我们通过内部类完成了轮胎Wheel对象的封装, 但调用时感觉很奇怪.

不管是通过外部类名, 还是外部类实例找到内部类名再实例化, 都会有一种他们两者根本没什么必然联系的感觉...反正不是父子关系

```py
#!/usr/bin/env python
#!encoding: utf-8

class Car():
    def __init__(self):
        self.door = 'a door'
        self.wheel = self.Wheel()       ## 在外部类方法中可以实例化内部类
    class Wheel():
        def run(self):
            print('i am running')

car = Car()
car.wheel.run()

wheel1 = car.Wheel()                    ## 通过外部类实例, 实例化内部类
wheel1.run()      ## i am running

wheel2 = Car.Wheel()                    ## 通过外部类名直接得到内部类实例
wheel2.run()     ## i am running
```

> 注意: 这里需要明确一点, 内部类本就**不适合描述类之间的组合关系**，而应把Door，Wheel类的对象作为类的属性使用

关于作用域, 我们在外部类方法中实例化内部类后, 就可以通过外部类实例访问内部类实例的成员属性或方法. 

如果要在内部类实例中访问外部类实例的成员呢? 在外部类方法中实例化内部类时传`self`进去吧, 这样内部类就可以取到外部类的实例对象, 进而可以访问其成员属性和方法了.

```py
#!/usr/bin/env python
#!encoding: utf-8

class Car():
    def __init__(self):
        self.door = 'a door'
        self.wheel = self.Wheel(self)       ## 实例化内部类时传入self, 代表外部类实例对象
    class Wheel():
        def __init__(self, outObj = None):
            self.outObj = outObj            ## 得到外部类实例对象
        def run(self):
            print(self.outObj.door)

car = Car()                                 ## 需要实例化外部类才行
car.wheel.run()                             ## a door
```

但要明白, 这种情况下相当于实现了一个闭包, 内部类实例持有外部类句柄, 不销毁它的话外部实例所占内存也不会释放, 希望不要滥用.