# JS-字符串截取

参考文章

1. [JavaScript截取字符串的Slice、Substring、Substr函数详解和比较](http://www.jb51.net/article/48257.htm)

按照索引截取.

```js
var testStr = 'hello world!';
```

## 1. slice

语法: `testStr.slice(start, [stop])`

1. `start`: 表示起始位置

2. `stop`: 默认直接到末尾

注意: **左开右闭区间**(即`start`为2, 但不包括第2个字符; `stop`为3, 包括第3个)

```js
testStr.slice(2)
"llo world!"
testStr.slice(2, 3)
"l"
```

`start`和`stop`都可以为负数, 表示从倒数第几个开始或结束. 但这个时候就不再是左开右闭区间了. 

说了也容易忘, 用的时候试验一下吧, 免得出错.

```js
testStr.slice(-2)         // 从倒数第2个字符开始到字符串结束
"d!"
testStr.slice(-3, -1)    // 从倒数第3个字符开始, 倒数第1个字符结束
"ld"
testStr.slice(3, -1)        // 从第3个字符开始, 到倒数第1个字符结束
"lo world"
```

> 其实`slice`主要用于数组的切片式截取...尤其可以像python一样可以用负数逆向截取.

## 2. substring

语法: `testStr.substring(indexA [, indexB])`

`substring()`的参数没有明确指明`start`与`stop`, 因为它的两个参数没有先后顺序. 总是以较小的值为开始, 较大的值为结束.

```js
testStr.substring(2, 5)
"llo"
testStr.substring(5, 2)
"llo"
```

但它的两个参数就取负数了, 取负值将被看作是0. 嗯...两个都取负数的时候是得不到什么东西的.

```js
testStr.substring(5, -2)
"hello"
testStr.substring(-2, 5)
"hello"
testStr.substring(-2, -5)
""
```

## 3. substr

语法

```js
testStr.substr(start [, length])
```

除了第二个参数表示截取长度之外, 其他的与slice相同, start可正可负, 负数时表示倒数.

```js
testStr.substr(2)
"llo world!"
testStr.substr(2, 3)
"llo"
testStr.substr(-5, 3)
"orl"
```

> length参数可不能为负数哦~

不过按照参数文章1中所说, `substr 为Web浏览器附加的ECMAScript特性，不建议使用时 start 索引为负值`

未及验证, 使用时注意.
