# JS-in与for..in

参考文章

1. [js 关键字 in 的使用方法](http://www.cnblogs.com/fly-xfa/p/5968928.html)

## 1. for..in

`for..in`可以对数组和对象进行遍历(js原生是没有`foreach`操作符的).

数组遍历

```js
let list = [1, 2, 3, 4];
for(let idx in list){
	console.log(idx, list[idx]);
}
```

对象遍历

```js
let dict = {a: 1, b: 3, c: 7};
for(let key in dict){
	console.log(key, dict[key]);
}
```

## 2. in操作符

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