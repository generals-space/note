
extend()扩展方法

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