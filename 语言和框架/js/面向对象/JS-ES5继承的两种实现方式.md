# JS-ES5继承的两种实现方式

参考文章

1. [ES5和ES6中的继承](http://keenwon.com/1524.html)

2. [backbone源码 - extend的实现](https://cdn.bootcss.com/backbone.js/1.3.3/backbone-min.js)

## 1. 术语约定

```js
function A(){                       // 父类构造函数类
    this.A_attr1 = 'A_val1';        // 实例属性, new实例可以取得, 但无法被继承
    var A_attr4 = 'A_val4';         // 无意义, 因为这种赋值只有在把A当作普通函数执行时(如`A()`), 才会有意义; 在继承机制中, 无法通过任何手段取得.
}

A.A_attr2 = 'A_val2';               // 静态属性/类属性, new操作时不会被实例取得.(这个时候你可以把A当成普通object对象来看, 就是它的键值属性嘛)
A.prototype.A_attr3 = 'A_val3';     // 原型属性, new实例可以取得, 也可以被继承
```

在这里, `A`本身等同于其他语言的类概念, 由于它是由JS中的构造函数(其实就是首字母大写的普通函数, 没什么特别的)实现的, 所以我把它称为**构造函数类**.

另外, 在`new`操作时, 构造函数类的函数体中的代码会自动执行. 所以函数体整块区域就等同于其他语言的**构造函数**(C++中的`Constructor`, python中的`__init__`). 在这里面编写的代码, 应该都是其他语言中构造函数的代码才对.

由于构造函数占用了`构造函数类`整个函数体区域, 所以只能把所有希望被继承的方法都赋值在原型属性上(如上面的`A.prototype.A_attr3 = 'A_val3';`), 这样不仅能实现继承, 在new实例时也可以取得. 其实也就相当其他语言中定义的成员属性/方法了, 只不过没法写在`构造函数类`内部而已, 这也是没办法的事.

## 2. 构造函数 + 原型链

```js
function B(){
    A.call(this);                               // 调用父类A的构造函数...其实是为了继承父类的实例属性而已
    this.B_attr1 = 'B_val1';
}
B.prototype = new A();                          // 原型对象赋值, 可以得到一个父类实例, 它上面的实例属性和原型属性都能得到
                                                // ...不过实例属性会因为上面的构造函数中的`A.call(this)`覆盖掉.
B.prototype.constructor = B;

var b = new B();
```

关于constructor, 在JS中, 一个函数对象的`prototype`成员对象最初只有一个`constructor`属性, 直接指向它所属的函数本身(参考《JavaScript高级程序设计第3版》第6.2.3节). 

```js
function a(){};
a.prototype.constructor === a;                  // true
```

## 3. extend()扩展方法

Backbone中如下方式可以直接扩展父类得到子类对象.

```js
var V = Backbone.View.extend(properties, [classProperties]);
var v = new V();
```

实际上, 在backbone源码中, `extend`是一个普通函数, 把它赋值为View, Model等模块的方法很简单.

```js
function extend(){...};
View.extend = Model.extend = extend;
```

然后就可以调用了.

### 简单继承

不考虑传入`properties`扩展子类的`extend`函数实现.

### step 1. 取得父级构造函数类的引用

首先我们要明白, `extend`是一个通用函数, 谁把它赋值为成员属性, 就能调用它. 但是我们没有办法在`extend`内显式确定父级构造函数的身份, 所以我们需要想办法找到它.

```js
function extend(){
    var parent = this;      // parent变量指向父级构造函数类本身, function类型.
    console.log(parent);
}
```

然后执行一下.

```
A.extend = extend;
A.extend();
```

我们必须要先把最初的this值取出来, 因为之后会发生冲突. 

### step 2. 继承父级实例属性

继承父类的实例属性比较简单, 还是像第一种继承中那样, 通过`call`/`apply`方法完成即可.

```js
function extend(){
    var parent = this;
    /*
     * function: 继承父类实例属性
     * child是子级构造函数类, 首先它必须要是一个函数, 等价于第一种继承方式中的B类.
     * 然后在它的函数体内, 需要以其本身的this变量调用父类构造函数中的代码.
     * 注意对比第一种继承方式, 等价于B类中的`A.call(this);`
     */
    var child = function(){
        parent.call(this);
    };
}
```

### step 3. 继承父级原型属性

OK, 继承了父类的实例属性, 我们还有继承它的原型属性. 实例化父类并赋值到`child`函数的原型即可.

哦, 还有constructor属性的指向, 一切都与第一种继承方式相匹配.

```js
function extend(){
    var parent = this;
    var child = function(){
        parent.call(this);
    };
    child.prototype = new parent();
    child.prototype.constructor = child;
    return child;
}
```

完成!

------

实验一下.

```js
A.extend = extend;
var B = A.extend();
var b = new B();
console.log(b.A_attr1);                 // A_val1
console.log(b.A_attr3);                 // A_val3
```

接下来还需要考虑构造函数参数的传递, extend()接受子类扩展属性, 以及静态属性的添加.