# JS-ES5模拟super与多级继承(三)

参考文章

1. [js多层继承 super方法](http://blog.csdn.net/xiaotuwy/article/details/76152742)

参考文章1中提供了一个思路, `_super`不一定要是一个变量, 也可以是一个函数, 只要它能返回我们期望的父级对象就可以了. 下面是我对它给出的源码的一些修改和注释, 另外有3个测试示例.

```js
/*
 * @author: general
 * @github: https://gist.github.com/generals-space/a75cfca06e1f8d463022e0e02446c363
 */
/*
 * 要想拥有_super()方法, 必须继承SuperExtend类.
 * 注意: 
 * 1. inherits方法中不要再在assign时为子类添加指向父类本身的属性了, 会出问题的.
 * 2. 当需要使用_super()方法调用父类的某个方法时, 必须要保证子类有同名方法, 需要通过子类的方法调用父类方法才行
 */
function SuperExtend(){}
SuperExtend.prototype._super = function(){
    // caller调用者应该会是子类的成员方法对象, 或是子类构造函数本身
    var caller = arguments.callee.caller;
    // 这里先得到this所属的构造函数类
    var chain = this.constructor;
    var parent = null;
    // 沿继承链一直向上遍历, 至少要遍历到SuperExtend的第一个子类
    // 目标是**找到主调函数到底属于继承链上的哪一层级, 然后才能得到这个调用者的父类, 也就是我们需要的super对象**
    while(chain && chain.prototype){
        // 对象的隐式原型`__proto__`属性是一个指针, 它指向**构造本对象的**, **构造函数类**, **的原型**.
        // 但是由于inherits的自定义继承机制, chain.__proto__指向的是父级构造函数类(chain本身为子级构造函数类)
        parent = chain.__proto__;
        // 如果调用者正好是构造函数类本身, 说明是在构造函数类的函数体中调用的,
        // 直接返回父级构造函数类本身
        if(caller == chain) return parent;

        // 如果调用者不是子级构造函数类, 就应该是原型中的方法了.
        var props = Object.getOwnPropertyNames(chain.prototype);
        for(var i = 0; i < props.length; i ++){
            // 这里虽然相等, 但有可能是当前类从上一层父类继承而来的属性, 而当前类本身并没有定义过这个方法.
            // 需要进一步确认, 即确认父类原型上没有与它完全相同的方法(当然, 方法名可能一样).
            if(caller == chain.prototype[props[i]] && caller != parent.prototype[props[i]]){
                return parent.prototype;
            }
        }
        chain = parent;
    }
    return chain;
};
/*
 * function: 自定义通用继承方法.
 * 使用方法: inherits(子类, 父类)
 */ 
function inherits(subClass, superClass){
    Object.assign(subClass.prototype, superClass.prototype, {
        constructor: subClass,
    });
    // 建立这种联系后, 相当于subClass成了superClass的实例了
    // 基本等价于subClass.prototype = superClass
    Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; 
}
```

3个测试示例如下

### 测试用例1. 基本测试

```js
// 测试用例1. 基本测试
function A(a){
    this.a = a;
}

inherits(A, SuperExtend);

A.prototype.sayHi = function(){
    console.log(this.a);
};

function B(a, b){
    this._super().call(this, a);
    this.b = b;
}

inherits(B, A);

B.prototype.sayHi = function(){
    this._super().sayHi.call(this);
    console.log(this.a, this.b);
};

function C(a, b, c){
    // 这里得到的是父级构造函数类本身, 直接call调用即可
    this._super().call(this, a, b);
    this.c = c;
}

inherits(C, B);

C.prototype.sayHi = function(){
    // 这里得到的是父级构造函数类的原型对象
    this._super().sayHi.call(this);
    console.log(this.a, this.b, this.c);
};

var c = new C(2, 5, 8);
c.sayHi();
```

### 测试用例2. 验证同层级函数间调用的情况

```js
// 测试用例2. 验证同层级函数间调用的情况
function A(a){
    this.a = a;
}

inherits(A, SuperExtend);

A.prototype.sayA = function(){
    this.sayB();
};

A.prototype.sayB = function(){
    console.log(this.a);
};

function B(a, b){
    this._super().call(this, a);
    this.b = b;
}

inherits(B, A);

B.prototype.sayB = function(){
    this._super().sayB.call(this);
    console.log(this.a, this.b);
};

function C(a, b, c){
    // 这里得到的是父级构造函数类本身, 直接call调用即可
    this._super().call(this, a, b);
    this.c = c;
}

inherits(C, B);

var c = new C(2, 5, 8);
c.sayA();
```

### 测试用例3. 验证主调函数与被调函数不同名的情况

```js
// 测试用例3. 验证主调函数与被调函数不同名的情况
function A(a){
    this.a = a;
}

inherits(A, SuperExtend);

A.prototype.sayA = function(){
    console.log(this.a);    
};

function B(a, b){
    this._super().call(this, a);
    this.b = b;
}

inherits(B, A);

B.prototype.sayB = function(){
    this._super().sayA.call(this);
    console.log(this.a, this.b);
};

function C(a, b, c){
    // 这里得到的是父级构造函数类本身, 直接call调用即可
    this._super().call(this, a, b);
    this.c = c;
}

inherits(C, B);

var c = new C(2, 5, 8);
c.sayB();
```
