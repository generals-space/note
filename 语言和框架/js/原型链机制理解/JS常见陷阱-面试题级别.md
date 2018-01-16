# JS常见陷阱

## 自由变量的取值作用域问题

```js
var a = 10;
function fn(){
	console.log(a); // a是自由变量, 函数创建时就确定了a要取值的作用域
}

function bar(f){
	var a = 20; // 注意, 这里的a是bar的局部变量
	f();
}
bar(fn); // 10
```

## 闭包函数的this问题

```js
var obj = {
	x: 10,
	fn: function(){
		console.log(this); // obj.fn();
		console.log(this.x); // 10
		function f(){
			console.log(this); // Window {stop: ƒ, open: ƒ, alert: ƒ, confirm: ƒ, prompt: ƒ, …}
			console.log(this.x); // undefined
		}
		f();
	}
};
obj.fn();
```

obj.fn中的闭包函数f的this变量...取到了window.

## 未实例化的prototype

```js
function Fn(){
	this.name = '王福朋';
	this.year = 1988;
}

Fn.prototype.getName = function(){
	console.log(this.name);
}

var f1 = new Fn();
f1.getName(); // 王福朋
Fn.getName(); // 报错
```

`Fn.getName()`报错的理由, `Fn`与其`prototype`是两个独立对象, 在未实例化之前, 属性搜索不会有原型链方面的联系. 

切记, prototype只与继承机制相关.

## 延原型链操作父级对象属性的问题

参考文章

1. [Object.create()对对象属性prototype和__proto__的影响](https://segmentfault.com/a/1190000005968121)

两道题

```js
var a = { name: 'kelen' };
var b = Object.create(a);
b.name = 'boke';
console.log(a.name);  // kelen
```

```js
var a = { person : { name: 'kelen' } };
var b = Object.create(a);
b.person.name = 'kobe';
console.log( a.person.name ); // kobe
```

...看第1题时信心满满觉得b对象name属性的赋值不会影响父级对象a, 看了第2题就有些萌B了.

首先, 要认识到a,b都是普通对象, 没有`prototype`属性, 但仍然通过`__proto__`隐式原型实现了继承机制.

`var b = Object.create(a);`一句实现了`b.name === b.__proto__.name === a.name`. 此时`b.name`是发现b对象本身没有`name`属性而沿原型链到a对象处找到的.

第1题中

`b.name = 'boke';`则是对b对象本身添加了`name`属性, 这将导致访问`b.name`时直接读取b本身的name属性而不再是沿原型链去查找了. 

```js
b.name //boke
a.name //kelen
```

如果使用`delete b.name`删掉b对象的`name`属性, 再次访问`b.name`你会发现结果又是`kelen`了. 也就是说, 直接对`b.name`赋值根本不会影响到a对象.

想要通过b修改a对象, 可以使用`b.__proto__.name = 'kobe'`.

第2题中

`b.person.name`引用到的倒的确是`a`对象的`person.name`了. 因为b对象在没有设置其本身的`person`属性时, `b.person`其实就是`a.person`. 由于js的引用传值特性, `b.person`的任何修改, 其实都量对`a.person`的修改.

## 循环中设置事件监听或定时任务的数据获取问题

```js
var select = ['a', 'b', 'c'];
for(var i = 0; i < select.length; i ++){
	$(document).on('change', 'select.' + select[i], function(event){
		console.log(i);
		console.log(select);
	});
}
```

本来我们的目的是为3个类名分别为`a`, `b`, `c`的`select`元素添加事件监听的. 但是当这个3`select`的`change`事件真的发生的时候, 输出的`i`的值却都变成了3, 正好大于`select`数组变量的长度. 为什么?

因为在事件发生的时候, `i`变量已经因为`for`循环结束变成了`3`, 所以事件其实是已经绑定在了3个元素上, 这个过程没有问题. 但是回调函数对`i`变量的引用却是在`for`循环的作用域内, 相当于是回调函数的域外变量, 所以没有得到想要的值.

尝试在回调函数内通过`var j = i`暂存这个`i`变量的值, 但没有效果.