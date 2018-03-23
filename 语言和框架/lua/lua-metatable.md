# Lua-metatable

参考文章

1. [解析Lua中关于Metatable学习笔记](http://mobile.51cto.com/iphone-285892.htm)

Lua中关于`Metatable`学习笔记是本文要介绍的内容，先来了解一下什么是Metatable，中Metatable这个概念, 国内将他翻译为**元表**. 元表为重定义Lua中任意一个对象(值)的默认行为提供了一种公开入口. 如同许多OO语言的**操作符重载**或**方法重载**. Metatable能够为我们带来非常灵活的编程方式.

具体的说, Lua中每种类型的值都有都有他的默认操作方式, 如, 数字可以做加减乘除等操作, 字符串可以做连接操作, 函数可以做调用操作, 表可以做表项的取值赋值操作. 他们都遵循这些操作的默认逻辑执行, 而这些操作可以通过Metatable来改变. 如, 你可以定义2个表如何相加等.

看一个最简单的例子, 重定义了2个表的加法操作. 这个例子中将c的`__add`域改写后**将a的Metatable设置为c**, 当执行到加法的操作时, Lua首先会检查a是否有Metatable并且Metatable中是否存在`__add`域, 如果有则调用, 否则将检查b的条件(和a相同), 如果都没有则调用默认加法运算, 而table没有定义默认加法运算, 则会报错.

```lua
-- test.lua
a = {5, 6}
b = {7, 8}

c = {}

c.__add = function(arg1, arg2)
    for _, item in pairs(arg2) do
        table.insert(arg1, item)
    end
    return arg1
end
-- 将c作为a的metatable
setmetatable(a, c)

d = a + b
for _, item in pairs(d) do
    print(item)
end
```

执行它, 得到如下输出

```
$ lua test.lua 
5
6
7
8
```

有了个感性的认识后, 我们看看Metatable的具体特性.

Metatable并不神秘, 他只是一个普通的table, 在table这个数据结构当中, Lua定义了许多重定义这些操作的入口. 他们均以双下划线开头为table的域, 如上面例子的`__add`, 表示`+`操作符(当然其他操作符也有类似的定义). 当你为一个值设置了Metatable, 并在Metatable中设置了重写了相应的操作域, 在这个值执行这个操作的时候就会触发重写的自定义操作. 当然每个操作都有每个操作的方法格式签名, 如`__add`会将加号两边的两个操作数做为参数传入并且要求一个返回值. 有人把这样的行为比作事件, 当xx行为触发会激活事件自定义操作.

Metatable中定义的操作add, sub, mul, div, mod, pow, unm, concat, len, eq, lt, le, tostring, gc, index, newindex, call...

在Lua中任何一个值都有Metatable, 不同的值可以有不同的Metatable也可以共享同样的Metatable, 但在Lua本身提供的功能中, 不允许你改变除了table类型值外的任何其他类型值的Metatable, 除非使用C扩展或其他库. `setmetatable`和`getmetatable`是唯一一组操作table类型的Metatable的方法.

## 1. Metatable与面向对象

Lua是个面向过程的语言, 但通过Metatable可以模拟出面向对象的样子. 其关键就在于`__index`这个域. 他提供了表的索引值入口. 这很像重写C#中的索引器, 当表要索引一个值时如`table[key]`, Lua会首先在table本身中查找key的值, 如果没有并且这个table存在一个带有`__index`属性的Metatable, 则Lua会按照`__index`所定义的函数逻辑查找. 仔细想想, 这不正为面向对象中的核心思想继承, 提供了实现方式么. Lua中实现面向对象的方式非常多, 但无论哪种都离不开`__index`.

这个例子中我使用了Programming In Lua中的实现OO的方式, 建立了Bird(鸟)对象, 拥有会飞的属性, 其他鸟对象基于此原型, Ostrich(鸵鸟)是鸟的一种但不会飞. 结果很明显, Bird和Ostrich分别有独立的状态.

```lua
-- bird.lua
-- 类定义方式
local Bird = { 
    flyable = true,
    -- self是必须的参数
    New = function(self, flyable)
        local o = {}
        o.flyable = flyable
        -- 把self对象的metatable赋给o...然后返回o, 就能创建新实例了???
        setmetatable(o, self)
        self.__index = self
        return o
    end 
}
local Ostrich = Bird:New(false)
print(Ostrich.flyable)
```

执行它得到如下输出

```
$ lua bird.lua 
false
```

类对象还可以这样定义, 作用相同, 结果一样.

```lua
-- bird.lua

local Bird = { 
    flyable = true,
}
function Bird:New(flyable)
    local o = {}
    o.flyable = flyable
    -- 把self对象的metatable赋给o...然后返回o, 就能创建新实例了???
    setmetatable(o, self)
    self.__index = self
    return o
end 
local Ostrich = Bird:New(false)
print(Ostrich.flyable)
```