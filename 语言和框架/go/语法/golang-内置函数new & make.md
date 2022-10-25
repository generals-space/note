# golang-内置函数new & make

参考文章

1. [深入学习golang(4)—new与make](https://www.cnblogs.com/hustcat/p/4004889.html)

好像在golang里, `new`返回的是一个对象的首地址指针, 而`make`返回的才是对象的引用...和python, java这种高级语言不太一样.

参考文章1中已经说了, 通俗一点来讲, `new`的作用类似于c里的`malloc + memset(0)`, 构造了内存空间, 清空, 但并未初始化.

而`make`才更像python/java中的`new`, 不过new出来的是动态类型, `slice`, `map`, `channel`这种数据结构.

或者说, 正是因为`new`和`make`两种方式同时存在(基础类型的指针与动态类型的引用), 才使得golang能同时兼备C和java两种特性. 虽然我还并不太能理解更深层的原理.
