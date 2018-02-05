# JS-ES5模拟super与多级继承(一)

参考文章

1. [js多层继承 super方法](http://blog.csdn.net/xiaotuwy/article/details/76152742)

本系列文章对js es5实现多级继承做一个学习和探究, 第三篇给出最终的模拟代码及测试用例.

## 简单的父-子继承

```js
// 父类A
function A(a){
    this.a = a;
}

A.prototype.sayA = function(){
    console.log(this.a);
};

// 子类B
function B(a, b){
    this._super.call(this, a);
    this.b = b;
}

// 原型链继承
Object.assign(B.prototype, A.prototype, {
    constructor: B, 
    _super: A
});

B.prototype.sayB = function(){
    console.log(this.b);
};

var b = new B(1, 2);
b.sayA();   // 1
b.sayB();   // 2
```

这里我用了`_super`关键字表示了继承的父类, `Object.assign()`方法可以将其附加到子类实例对象上, 用起来会方便一点.

但是, 比较致命的一点是, 这种方式不适用于多级继承, 我所定义的`_super`反而成了限制.

```js
// 父类A
function A(a){
    this.a = a;
}

A.prototype.sayA = function(){
    console.log(this.a);
};

// 子类B
function B(a, b){
    this._super.call(this, a);
    this.b = b;
}

// 原型链继承
Object.assign(B.prototype, A.prototype, {
    constructor: B, 
    _super: A
});

B.prototype.sayB = function(){
    console.log(this.b);
};

// 子类C
function C(a, b, c){
    this._super.call(this, a, b);
    this.c = c;
}

// 原型链继承
Object.assign(C.prototype, B.prototype, {
    constructor: C, 
    _super: B
});

C.prototype.sayC = function(){
    console.log(this.c);
};

var c = new C(1, 2, 3);
c.sayA();
c.sayB();
c.sayC();
```

上面的代码看起来似乎没什么错误, 但是执行时, 会栈溢出, 在B类函数体的`this._super.call(this, a);`这一行.

```
VM4484:10 Uncaught RangeError: Maximum call stack size exceeded
    at C.B [as _super] (<anonymous>:10:11)
    at C.B [as _super] (<anonymous>:11:17)
```

原因在于, c在实例化时构造函数调用父类B的构造函数, 但用的是`call`方法, B类构造函数在执行时this的值为c的实例, 而`this._super`的值又是B, 于是就在B的构造函数里一直循环.

要解决这个问题, `_super`变量就不能绑定在this上, 但是好像也没有好的方法绑定在子类本身, 除非在子类中用父类的类名显示调用父类的同名方法. 但这样耦合性太强, 稍不注意就会出错(尤其是代码复制时).

参考文章1中有错误, 不存在`__super__`属性, 但它给了我一个启示, `super`不一定非得是变量, 也可以是一个函数, 由函数的执行结果作为父类对象也是一种方法.