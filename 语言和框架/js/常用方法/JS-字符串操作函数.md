# JS-字符串操作函数

参考文章

1. [Node.js ECMAScript6 字符串的扩展函数,字符串的拼接,字符串以...开头/结尾](https://blog.csdn.net/houyanhua1/article/details/80260762)

假设有如下字符串

```js
var testStr = 'hello world!';
```

```js
testStr.startsWith('hello'); // true 是否以...开头
testStr.endsWith('world!'); // true 是否以...结尾
testStr.includes('lo wo'); // true 是否包含
testStr.repeat(3); // "hello world!hello world!hello world!" 将原字符串重复n次并返回
```
