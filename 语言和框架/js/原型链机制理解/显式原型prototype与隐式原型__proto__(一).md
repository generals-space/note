# JS-显式原型`prototype`与隐式原型`__proto__`

参考文章

1. [js中__proto__和prototype的区别和关系？](https://www.zhihu.com/question/34183746)

一句话概括: 

js中一切皆对象, 所有对象都有隐式原型`__proto__`, 只有函数(Function)类型对象有(正统)原型prototype(当然, 作为对象它也有`__proto__`).

> 当然, 凡事有例外, 但基本上可以涵盖大部分情况了.

------

## 1. `__proto__`

ok, 一切皆对象好说, 但是一个字符串也有可能有`__proto__`吗? 

答案是肯定的, 不过需要特殊的方法去发现.

```js
var a = 'abc';
console.dir(a);             // 这样是看不出什么来的
> abc
a.__proto__
> String {length: 0, constructor: function, charAt: function, charCodeAt: function, concat: function…}
```

对象的隐式原型`__proto__`属性是一个指针, 它指向**构造本对象的**, **构造函数类**, **的原型**.

在这里, 变量`a`是一个`String`类型(也可以说是对象)的实例, 它的`__proto__`就指向`String`的(显式)原型. 

不信? 试试

```
a = '';
> ""
a.__proto__ === String.prototype
> true
```

那么, 隐式原型存在的意义呢? 

你以为`a`这么一个普通的变量是怎么拥有String定义的方法的? `slice`, `split`...? 究其根本, 就是因为`__proto__`的存在, **它保证了实例能够访问在构造函数原型中定义的属性和方法**

感觉`__proto__`和`prototype`分别定义了广义上类与实例, 类与类之间的联系方式.

## 2. prototype

函数对象在创建之初就默认~~继承了`Function`基类~~, 错, 不是继承关系, 而是类与实例的关系, 如下

```js
function a(){}
a instanceof Function       // 判断a是不是Function的实例
> true
```

要知道, 这个时候还没prototype什么事呢...

于是`a.__proto__ === Function.prototype`的结果为`true`也不会再让人疑惑了.

...那`Function`基类本身也算是对象啊, 它的`__proto__`指向谁?

这个问题貌似就有点烧脑了诶, 因为`Function.__proto__ === Function.prototype`...

所有内置包装类, 包括`Array`, `String`等的原型对象的`__proto__`都指向`Function.prototype`. 在这个时刻, `Array`等都被看作是函数, 虽然是构造函数...所以`Object`也被看作是函数, 所以`Object.__proto__ === Function.prototype`...

先撇开这个不谈, `Function.prototype`原型对象的`__proto__`又是谁?

答案是`Function.prototype.__proto__ === Object.prototype`, 同样, `Array.prototype.__proto__ === Object.prototype`.

而`Object.prototype.__proto__ === null`.

现在我们可以大胆猜测, js语言的源头其实是原型, Object是最初始的类型, 其他的内置对象, 都是在此基础上构造而来的.

下面一张图给出所有解释.

![](https://gitee.com/generals-space/gitimg/raw/master/e46508fbcd140db304232aba89f41c83.jpg)

