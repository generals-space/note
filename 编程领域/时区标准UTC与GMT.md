# 时区标准UTC与GMT

参考文章

1. [UTC和GMT什么关系？](https://www.zhihu.com/question/27052407)
2. [ISO日期格式标准](http://www.zoucz.com/blog/2016/01/29/date-iso/)

## 1. 时区

UTC是根据原子钟来计算时间(世界上最精确的原子钟50亿年才会误差1秒), 是时间的计算依据.

而GMT是`Greenwich Mean Time`, 是划分时区的基准, 即零时区. GMT（格林威治时间）、CST（北京时间）、PST（太平洋时间）等等是具体的时区.

GMT: UTC +0    =    GMT: GMT +0
CST: UTC +8    =    CST: GMT +8
PST: UTC -8    =    PST: GMT -8

即, 其他时区都是以GMT为基准`+n`或`-n`来表示的.

## 2. ISO日期格式标准

```js
d = new Date();
Fri Jun 15 2018 16:34:24 GMT+0800 (China Standard Time)
d.toISOString();
"2018-06-15T08:34:24.091Z"
```

`toISOString()`方法返回的日期字符串, 中间的`T`表示`UTC`格式, 而最后的`Z`则表示时间偏移量. 它是4位的数字格式, 不写时默认为不偏移. 在js里, 这个方法只会将日期对象转换为从零时区表示的值, 所以`Z`后面是空的.

实际上`2018-06-15T08:34:24.091Z`和`2018-06-15T16:34:24.091+0800`表示的是同一个时间点, 只不过形式不一样而已. 

而`new Date()`默认以系统当前配置的时区创建时间对象, 它接受通用的字符串如`2018-06-15 16:34:24`这种. 当前也可以写成`ISO`标准格式, 这样就可以自定义时区了. 比如

```js
d = new Date()      // 当前时间
Fri Jun 15 2018 16:47:12 GMT+0800 (China Standard Time)
d = new Date('2018-06-15T16:47:12+0000')    // 创建零时区的时间对象
Sat Jun 16 2018 00:47:12 GMT+0800 (China Standard Time)
d.toISOString()                             // 输出形式还是零时区的话, 我们定义的时间字符串就相同了.
"2018-06-15T16:47:12.000Z"
```
