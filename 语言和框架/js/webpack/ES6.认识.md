`const`关键字在编译时即确定其地位, 如果在ES6代码中试图修改一个`const`类型变量, 编译会出错.

`let`关键字在声明局部块变量时, 转换成ES5时的妥协做法是...换个名字, 依然还是全局变量, 于是这样与编译前同名的变量就不同名了, 也就没有冲突了.
