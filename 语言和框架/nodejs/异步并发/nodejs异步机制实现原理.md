# nodejs异步机制实现原理

参考文章

1. [深入剖析Nodejs的异步IO](http://blog.csdn.net/yezhenxu1992/article/details/51731237)

2. [解析NodeJS异步I/O的实现](http://www.jb51.net/article/111089.htm)

NodeJS的单线程, 原生支持异步IO的特性, 是在底层对操作系统接口进行了封装的缘故.

单线程的表现

NodeJS并没有在语言层面提供线程接口, 所以开发人员也根本无从选择.

关于异步IO

IO包括文件IO与网络IO, 只有IO操作才有异步的必要, 这一点应该很好理解.

在常规高级语言中, 文件操作代码类似这样

```
file = open(path, bytes, ...);
data = file.read()
print(data)
```

而node的实现方式是

```js
fs.readFile(path, ..., function (err, data) {
    console.log(data); //打印test01.txt文本内容
});
```

太方便了有没有.

这种实现机制是由语言提供的, 一般也只有网络操作(http请求, socket请求如数据库连接等)和文件操作会用到.node解析器会分析IO对象, 在其C++实现的底层中, 依然是通过多线程, 线程池、事件池等手段实现的.

具体的实现代码在`libv`这个底层库中.

------

再来说说python的异步框架, twisted/tornado, 后者貌似是通过纯python实现了异步机制, 也就是可以在python中在网络操作中使用异步功能. 我想它们应该是和nodejs的libv采用的相似的方法, 总逃不开线程池, 事件通知机制...

但是由于nodejs原生支持异步IO, 整个生态中(这里我指的是各种基础库, 各种数据库驱动, 网络框架等)各个层面中支持异步, 而tornado只是在应用层的模拟, 所以可能没有合适的异步数据库驱动? 你可能需要自己手动按照异步模式写这些驱动, 总之与python自身的生态并不兼容.

同样是基于线程的实现, 异步的好处就是, 不让IO成为性能瓶颈, 最大限度发挥CPU的潜力. 同样开10个线程, 异步方式相对同步方式有相当大的优势.