
```js
/*
 * 要想拥有_super()方法, 必须继承SuperExtend类.
 * 注意: 
 * 1. inherits方法中不要再在assign时为子类添加指向父类本身的属性了, 会出问题的.
 * 2. 当需要使用_super()方法调用父类的某个方法时, 必须要保证子类有同名方法, 需要通过子类的方法调用父类方法才行
 */
function SuperExtend(){}
SuperExtend.prototype._super = function(){
    // caller调用者应该会是子类成员方法对象或是子类构造函数本身
    var caller = arguments.callee.caller;
    // chain变量为继承链上的构造函数类对象, 
    // 通过chain.__proto__可以找到父类对象
    var chain = this.constructor;

    var parent = null;
    while(chain && chain.prototype){
        parent = chain.__proto__;
        // 如果调用者正好是构造函数类本身
        if(caller == chain){
            // 这里返回的父级构造函数类本身
            return parent;
        }

        var props = Object.getOwnPropertyNames(chain.prototype);
        for(var i = 0; i < props.length; i ++){
            // 这里虽然相等, 但有可能是当前类从上一层父类继承而来的属性, 
            // 我们需要进一步确认, 即确认父类原型上没有相同的属性.
            if(caller == chain.prototype[props[i]] && caller != parent.prototype[props[i]]){
                return parent.prototype;
            }
        }
        chain = parent;
    }
    return chain;
};

function inherits(subClass, superClass){
    Object.assign(subClass.prototype, superClass.prototype, {
        constructor: subClass,
    });
    // 建立这种联系后, 相当于subClass成了superClass的实例了
    // 基本等价于subClass.prototype = superClass
    Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; 
}

/*
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
    this._super().prototype.sayB.call(this);
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

*/

/*

// 测试用例1. 

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
*/

// 2
// 2 5
// 2 5 8
/////////////////////////////////////////////////////////////////////////////
// 测试用例2
/*
function A(){}
inherits(A, SuperExtend);
A.prototype.sayA = function(){
    console.log(123);
};

function B(){}
inherits(B, A);
B.prototype.sayB = function(){
    this._super().sayA.call(this);
    console.log(234);
};

// 单纯这样是不行的, 因为inherits建立继承关系后, 相当于把A.prototype上的属性全都附加到B.prototype上, 
// 那在_super的while循环中会错误判断, 我们必须添加一个同名属性, 通过调用子类自身的方法调用父类方法...

B.prototype.sayA = function(){
    this._super().sayA.call(this);
};
B.prototype.sayB = function(){
    this.sayA.();
    console.log(234);
};

*/
```