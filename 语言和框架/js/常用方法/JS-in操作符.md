# JS-in操作符

参考文章

1. [js 关键字 in 的使用方法](http://www.cnblogs.com/fly-xfa/p/5968928.html)

除了`for..in..`, 单纯的`in`操作符可以用于判断目标对象中是否存在某属性, 不管是直接属性还是原型属性.

```js
obj = {
	a: 123,
	b: 234
}
// {a: 123, b: 234}
'a' in obj
// true
delete obj.a
// true
'a' in obj
// false
```