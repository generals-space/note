# JS-ASCII字符与数字互转

不同于`parseInt()`与`parseFloat()`可以将字符串`"123",`, `"12.3"`转换成`123`, `12.3`这种, 我们希望得到字符`A`在ASCII表中的序号, 也希望能相反地也能够得到结果.

js提供这样的方法.

```js
var char = 'A';
char.charCodeAt()   // 得到65
```

```js
String.fromCharCode(65) // 得到A
```

...错了, `charCodeAt()`可编码的其实不只是ASCII表的字符, 貌似还包含了所有utf8的字符(还是unicode来着?) 起码可以得到中文在编码表中的位置, 而且`fromCharCode`可以成功反查.

```js
var a = '中'
a.charCodeAt()  // 20013
String.fromCharCode(20013) // "中"
```

------

字符串对象还有一个`charAt(n)`方法, 它的作为是返回目标字符串的第n个字符.

```js
var str = 'ABCD';
str.charAt(0)   // "A"
str.charAt(2)   // "C"
```