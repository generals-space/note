# JS-继承的实现方式

参考文章

1. [ES5和ES6中的继承](http://keenwon.com/1524.html)

2. [backbone源码 - extend的实现](https://cdn.bootcss.com/backbone.js/1.3.3/backbone-min.js)

## 1. 术语约定

```js
// 父类构造函数类
function A(){
    // 实例属性, new实例可以取得, 但无法被继承
    this.A_attr1 = 'A_val1';
    // 无意义, 因为这种赋值只有在把A当作普通函数执行时(如`A()`), 才会有意义; 在继承机制中, 无法通过任何手段取得.
    var A_attr4 = 'A_val4';         
}

// 静态属性/类属性, new操作时不会被实例取得.(这个时候你可以把A当成普通object对象来看, 就是它的键值属性嘛)
A.A_attr2 = 'A_val2';
// 原型属性, new实例可以取得, 也可以被继承
A.prototype.A_attr3 = 'A_val3';
```

在这里, `A`本身等同于其他语言中**类**的概念, 由于它是由JS中的构造函数(其实就是首字母大写的普通函数, 没什么特别的)实现的, 所以我把它称为**构造函数类**.

另外, 在`new`操作时, 构造函数类的函数体中的代码会自动执行. 所以函数体整块区域就等同于其他语言的**构造函数**(C++中的`Constructor`, python中的`__init__`). 在这里面编写的代码, 应该都是其他语言中构造函数的代码才对.

由于构造函数占用了`构造函数类`整个函数体区域, 所以只能把所有希望被继承的方法都赋值在原型属性上(如上面的`A.prototype.A_attr3 = 'A_val3';`), 这样不仅能实现继承, 在new实例时也可以取得. 其实也就相当其他语言中定义的成员属性/方法了, 只不过没法写在`构造函数类`内部而已, 这也是没办法的事.

## 2. 构造函数 + 原型链

```js
function B(){
    // 调用父类A的构造函数...其实是为了继承父级构造函数类内部定义的实例属性而已
    A.call(this);
    this.B_attr1 = 'B_val1';
}
// 原型对象赋值, 可以得到一个父类实例, 它上面的实例属性和原型属性都能得到
// ...不过实例属性会因为上面的构造函数中的`A.call(this)`覆盖掉.
B.prototype = new A();                          
B.prototype.constructor = B;

var b = new B();
```

> 关于`constructor`, 在JS中, 一个函数对象的`prototype`成员对象最初只有一个`constructor`属性, 直接指向它所属的函数本身(参考《JavaScript高级程序设计第3版》第6.2.3节). 

```js
function a(){};
a.prototype.constructor === a;                  // true
```

## 3. 原型链继承机制核心代码的几种实现

**第1种**

最原始的形式, 就是

```js
B.prototype = new A();
B.prototype.constructor = B;
```

不过这样有点不太好的就是, 在声明B的继承语句的时候要实例化A一次, 其实不太喜欢这么做.

**第2种**

下面的语句也可以实现同样的功能, 把`B.prototype`对象本身的原型指向`A.prototype`对象(虽然没有`B.prototype.prototype`, 但不影响原型链的继承), 而是像上面那样直接覆盖, 所以不必修改`constructor`的指向.

```js
B.prototype.__proto__ = A.prototype;
```

**第3种**

```js
B.prototype = Object.create(A.prototype);
B.prototype.constructor = B;
```

**第4种**

```js
$.extend(B.prototype, new A());
B.prototype.constructor = B;
```

需要注意的, 虽然表面看起来js原生方法`Object.assign()`能实现与`$.extend()`相似的功能, 但是`Object.assign(B.prototype, new A())`没法合并A类`prototype`的属性到B类实例中...mmp. 但是可以用下面的方式扩展.

```js
Object.assign(B.prototype, A.prototype);
B.prototype.constructor = B;
```