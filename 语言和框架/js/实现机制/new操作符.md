# JS-new操作符

参考文章

1. [JS原型继承和类式继承](http://www.cnblogs.com/constantince/p/4754992.html)

2. [理解js中的new](http://rockyuse.iteye.com/blog/1426522)

JS中其实是没有类的概念的(所谓的类都是模拟出来的). 但是当我们是用`new`关键字创建对象的时候经常会混淆与其他语言中传统类的区别.

在其他语言里, new是用来创建对象的, 是类实例化的过程. 在JS里, 虽然new也是用来创建对象的, 但JS本就没有类, 自然也就没有实例的概念. 

在这时我们需要认识new的执行机制.

new的典型用法如下.

有"基类"base

```js
var base = function(){}
```

我们使用`new`创建`base`的实例.

```js
var obj = new base();
```

在这一操作中, new实际上做了3件事.

```js
var obj = {};
obj.__proto__ = base.prototype;
base.call(obj);
```

1. 创建一个空对象obj;

2. 将这个空对象的`__proto__`成员属性指向base的`prototype`成员对象;

3. 使用`call`函数在obj作用域中调用base函数. 至此, base函数中所有绑定在`this`对象上的属性, 都将重新绑定到obj对象上.

注意, new的使用是有条件的, 在js中所有对象都拥有`__proto__`属性, 但不是所有对象都有`prototype`属性, 只有`Function`类型才有. 所有base的角色只能由`Function`类型对象表示. 

> 不过, js中任意函数都可以当成'构造函数'使用(构造函数就是那种需要new实例化才能使用的函数), 这样显然不太好. 为了区分普通函数与构造函数, 最好还是首字母大写.

------

参考文章1中最后给出了使用`Object.create()`实现与`new`关键字同样的功能的方法.

```js
var obj = Object.create(base);
```

> `new`关键字掩盖了Javascript中真正的原型继承, 使得它更像是基于类的继承. 其实new关键字只是Javascript在为了获得流行度而加入与Java类似的语法时期留下来的一个残留物.

但这种方法只是将新base对象作为obj的原型创建, 而且

```js
obj.__proto__ === base
obj.prototype === base.prototype // 这个只有在base是Function类型才会出现
``` 

obj的`__proto__`属性直接就指向了`base`对象而不是`base`的`prototype`属性. 当然, 这并不影响实例化与继承机制的实现. 