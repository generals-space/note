# JS-bind关键字

参考文章

1. [javascript中call(),apply().bind()方法的作用和异同](https://segmentfault.com/a/1190000011569075)

2. [Function.prototype.bind() - MDN官方文档](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Function/bind)

3. [偏函数 - 廖雪峰的官方网站](https://www.liaoxuefeng.com/wiki/001374738125095c955c1e6d8bb493182103fac9270762a000/001386819893624a7edc0e3e3df4d5d852a352b037c93ec000)

`bind`与`call`及`apply`的作用类似, 都是改变函数运行时的`this`对象指向.

不同的是, **`call`及`apply`出现的同时函数就开始执行了, 而`bind`在调用后会返回一个新的函数, 可以在其他需要调用的地方去调用它**.

> bind()的作用其实与call()以及apply()都是一样的，都是为了改变函数运行时的上下文，bind()与后面两者的区别是，call()和apply()在调用函数之后会立即执行，而bind()方法调用并改变函数运行时的上下文的之后，返回一个新的函数，在我们需要调用的地方去调用他。


基本用法

```js
var obj = {
    subFunc: function(arg){
        console.log(arg);
    }.bind(this, arg)
}
```

之后再调用`obj.subFunc`时, 其this的指向就不会再被改变了, 尤其是jquery的事件监听时.

参考文章1中非常简明准确地指出双方的区别, 然后再去看MDN官方文档.

> bind() 函数会创建一个新函数（称为绑定函数），新函数与被调函数（绑定函数的目标函数）具有相同的函数体（在 ECMAScript 5 规范中内置的call属性）。当新函数被调用时 this 值绑定到 bind() 的第一个参数，**该参数不能被重写**, 所以新函数不论怎么调用都有同样的 this 值。

## 关于偏函数

...偏函数的英文原文为`Partial Functions`, `Partial`意为`部分的, 偏爱的`, 我觉得取`部分`之意更合适. 因为偏函数就是对原函数进行了一下包装, 提供了部分参数作默认参数, 以保证执行时可以比较简单的调用而已, 具体含义可以查看参考文章3, 讲解得十分清晰.

python提供的`functools.partial`是创建偏函数的便捷方法, 等同于js的`bind`.

```js
function list() {
  return Array.prototype.slice.call(arguments);
}

var list1 = list(1, 2, 3); // [1, 2, 3]

// 偏函数leadingThirtysevenList
var leadingThirtysevenList = list.bind(undefined, 37);

// 像不像多态?
var list2 = leadingThirtysevenList(); // [37]
var list3 = leadingThirtysevenList(1, 2, 3); // [37, 1, 2, 3]
```

...不过好像没什么用, 本来创建偏函数就不用内置方法, 直接写个包装函数就行了, 所以`bind`的偏函数用法可以直接忽略.

## 作为构造函数使用的绑定函数

翻译的很差, 看了原文...没看懂. 咳...

简单来说, 就是如果对一个构造函数类调用`bind`创建一个新函数, 那么...没什么luan用.

```js
function Point(x, y) {
  this.x = x;
  this.y = y;
}

Point.prototype.toString = function() { 
  return this.x + ',' + this.y; 
};

var p = new Point(1, 2);
p.toString(); // '1,2'

var emptyObj = {};
// 这里绑定了一个空对象作为this, 按照最初的认知, 
// toString()方法中调用的this.x与this.y是找不到值的. 
// 但实际上不是.
var YAxisPoint = Point.bind(emptyObj, 0/*x*/);

var axisPoint = new YAxisPoint(5);
// 结果很正确, 说明bind绑定的this对象没生效, 但绑定的x参数的确是传入了的.
axisPoint.toString(); // '0,5'      

axisPoint instanceof Point; // true
axisPoint instanceof YAxisPoint; // true
new Point(17, 42) instanceof YAxisPoint; // true
```

其实在MDN手册最开始就有说明: **当绑定函数被调用时，第一个参数会作为原函数运行时的this指向。当使用new 操作符调用绑定函数时，该参数无效。**

## 实现原理

`bind`函数在 ECMA-262 第五版才被加入；它可能无法在所有浏览器上运行。参考文章2中有提供`bind`的模拟实现, 可以借鉴一下.

```js
if (!Function.prototype.bind) {
    // 如果没有bind方法, 则为其创建一个,
    // 参数oThis即为新this对象的指向, 其他参数在`arguments`对象中取得
    Function.prototype.bind = function(oThis) {
        // 只有函数对象拥有此方法
        if (typeof this !== 'function') {
            // closest thing possible to the ECMAScript 5
            // internal IsCallable function
            throw new TypeError('Function.prototype.bind - what is trying to be bound is not callable');
        }
        // 得到除oThis之外的其他参数, 数组类型
        var aArgs   = Array.prototype.slice.call(arguments, 1);
        var fToBind = this;
        var fNOP    = function() {};
        var fBound  = function() {
            return fToBind.apply(this instanceof fNOP ? this: oThis,
                // 获取调用时(fBound)的传参. bind 返回的函数入参往往是这么传递的
                aArgs.concat(Array.prototype.slice.call(arguments))
            );
        };

        // 维护原型关系
        if (this.prototype) {
            // Function.prototype doesn't have a prototype property
            fNOP.prototype = this.prototype; 
        }

        fBound.prototype = new fNOP();

        return fBound;
    };
}
```

关于`fBound`函数的返回值`apply`调用时第一个参数的选择问题.

还记得`bind`函数在一种情况下会失效么? 没错, 用于构造函数类时, new 操作创建的实例对象不会改变原this指向.

首先, 上述代码返回值为`fBound`函数, 即执行`bind(target)`后就得到`fBound`这个函数. 调用时即执行`fBound`的返回值.

而`this instanceof fNOP`这个判断依据, 就是, 

1. 当主调函数调用时, 其所处的上下文(即this指向)如果与`bind()`指定的不同, 就按照`oThis`的值去`apply`.

2. 当使用`new`操作符时, return语句中出现的`this`实际上指向实例对象本身, 而新实例对象即是`fBound`的实例对象(就是把普通函数也当构造函数类用了而已), 而`fBound`的原型链上是可以找到`fNOP`的, 所以这个判断将得到真.

> 注意: 原生`bind`得到的新函数没有`prototype`属性, 而上述模拟的`bind`实现了`prototype`.