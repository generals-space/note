# JS-typeof与instanceof区别及实现原理

参考文章

1. [instanceof - MDN官方文档](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Operators/instanceof)

2. [js中typeof和instanceof用法区别](http://blog.csdn.net/u014421556/article/details/52083215)

## 1. 区别

`typeof`的返回值只有 number, boolean, string, object, undefined, function这几种基本类型.

而`instanceof`的选择就有很多了.

除了`typeof`没有办法直接分辨数组和对象, 简直完全没什么好说的. 这时直接用`instanceof`就行了, 不用纠结其他的.

## 2. instanceof实现原理

首先, `instanceof`的语法为

```js
a instanceof A
```

往简单了说, 就是用来判断`a`是否是`A`类(或者子类)的实例对象的.

如

```js
var arr = new Array();
arr instanceof Array;       // true
```

专业一点讲, 其判断依据就是`A.prototype`对象是否在`a`对象的原型链上. 

比如`var arr = new Array();` 那么`arr`就会拥有`Array.prototype`上的属性, 即`Object.getPrototypeOf(arr) === Array.prototype`.

我们知道, 一个对象的隐式原型`__proto__`指向它所属的构造函数类的原型, 即`prototype`属性. 这也是原型链机制实现的真正原理. 只要能通过原型链找到`A.prototype`, 结果就为true.

## 3. 破坏方法

JS的继承过程不比静态语言, 原型属性是可以动态修改的, 所以`instanceof`的结果也不是完全不变的. 只要修改`a`的`__proto__`指向, 或是修改`A`类的`prototype`对象, 都会打破这样的关系.

第1种, `__proto__`

```js
function A(){};
var a = new A();
a instanceof A;         //true
a.__proto__ = {};       // 本来a.__proto__ === A.prototype
a instanceof A;         //false
```

第2种, `prototype`

```js
function A(){};
var a = new A();
a instanceof A;         //true
A.prototype = {};
a instanceof A;         //false
```