# JS对象合并

参考文章

1. [Javascript 对象（object）合并](http://www.cnblogs.com/yes-V-can/p/5631645.html)

2. [[译]ECMAScript 6: 使用 Object.assign() 合并对象](http://www.tuicool.com/articles/VF3Uf2E)

3. [Object.assign()](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Object/assign)

> 本文的对象单指字典列表对象, 不包括数组类型对象.

## 1. 基本方法

从一个对象复制所有的属性到另一个对象是一个常见的操作, 这个操作在Javascript生态系统中有一个专属名称, 叫做 **extend**.

有很多库都实现了这个方法

- Prototype: [Object.extend(destination, source)](http://prototypejs.org/doc/latest/language/Object/extend/) (Prototype 是第一个使用'extend'这个名称的库)

- Underscore.js: [_.extend(destination, *sources)](http://underscorejs.org/#extend)

- jquery: [$.extend[deep], target, object1, [objectN]](http://jquery.cuishifeng.cn/jQuery.extend.html)

新版的ES6也有原生的解决方案

- [Object.assign(target, ...sources)](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Object/assign)

它们的使用方法大同小异

```js
var obj1 = {a: 'abc'};
var obj2 = {b: 123};
var obj = Object.assign(obj1, obj2);
console.log(obj);
console.log(obj1);
```

```
Object {a: "abc", b: 123}
    a : "abc"
    b : 123
    __proto__ : Object
Object {a: "abc", b: 123}
    a : "abc"
    b : 123
    __proto__ : Object
```

上面介绍到的4种不同的合并方法, 第一个参数都是`target`, 所以合并后的对象会将`obj1`覆盖. 如果不希望发生这样的情况, 可以把第一个参数改为`{}`, 如下.

```js
var obj1 = {a: 'abc'};
var obj2 = {b: 123};
var obj = Object.assign({}, obj1, obj2);
console.log(obj);
console.log(obj1);
```

输出为

```
Object {a: "abc", b: 123}
    a : "abc"
    b : 123
    __proto__ : Object
Object {a: "abc"}
    a : "abc"
    __proto__ : Object
```

## 2. 深浅拷贝

```js
var obj1 = {
    a: 'abc',
    subobj: {
        a: 'abc',
        b: 'abc'
    },
};
var obj2 = {
    b: 123,
    subobj: {
        a: 123,
    },
};
var obj = Object.assign({}, obj1, obj2);
console.log(obj);
```

输出为

```
Object {a: "abc", subobj: Object, b: 123}
    a : "abc"
    b : 123
    subobj : Object
        a : 123
        __proto__ : Object
    __proto__ : Object
```

使用`assign`方法合并后, `obj2.subobj`的值覆盖了`obj1.subobj`的值, 而不是对`subobj`字段继续进行合并, 这应该相当于浅拷贝.

至于深拷贝, 自然就是目标对象的子对象实现递归合并了. 我只知道`jquery.extend()`方法的第一个参数可以控制是否进行深拷贝, 至于`Prototype.js`与`Undersocre.js`的实现, 不深究.

```js
var obj1 = {
    a: 'abc',
    subobj: {
        a: 'abc',
        b: 'abc'
    },
};
var obj2 = {
    b: 123,
    subobj: {
        a: 123,
    },
};
// 普通浅拷贝
// var obj = $.extend({}, obj1, obj2);
// 递归深拷贝
var obj = $.extend(true, {}, obj1, obj2);
console.log(obj);
```

结果如下

```
Object {a: "abc", subobj: Object, b: 123}
    a : "abc"
    b : 123
    subobj : Object
        a : 123
        b : "abc"
        __proto__ : Object
    __proto__ : Object
```

## 3. 继承中原型链扩展的问题

`Object.assign`不能应用于原型扩展继承, 可能是我的用法不对, 总之始终没想到错在哪里了.

```js
function A(){}

A.prototype.sayA = function(){
    console.log('I am sayA...');
};

function B(){}

$.extend(B.prototype, new A());
// Object.assign(B.prototype, new A());

var b = new B();
b.sayA();
```

`$.extend()`能正确扩展, 但是`Object.assign()`不行, 会报错说找不到`sayA`方法...

理论上`new A()`将得到一个实例对象, `sayA`将是它的一个属性, 但却没能合并到`B.prototype`对象上, 就算A的实例`sayA`只是挂在原型链上, 那`Object.assign(B.prototype, A)`也没法扩展, 实在搞不懂. 留个疑问<???>