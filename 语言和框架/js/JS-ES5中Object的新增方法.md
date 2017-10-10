# JS-ES5中Object的新增方法

参考文章

1. [ECMAScript5 Object的新属性方法](http://www.cnblogs.com/dolphinX/p/3348467.html)

## 1. Object.create(prototype[, descriptors])

这个方法用于创建一个对象，并把其`prototype`属性赋值为第一个参数，同时可以设置多个descriptors，关于decriptor下一个方法就会介绍这里先不说。

```js
var proto = {
    "say": function () {
        alert(this.name);
    },
    "name":"Byron"
}
var o = Object.create(proto);
```

这样的结果是

```js
o.__proto__ === proto  //true
```

`o`将通过隐式原型`__proto__`将原型指向proto对象, 同时`proto`作为普通对象也有其自己的原型链, 于是实现了继承.(原文中提到'原型链干净对象'我觉得不太合理).

## 2. Object.defineProperty(O, Prop, descriptor) / Object.defineProperties(O, descriptors)

想明白这两个函数必须明白descriptor是什么，在之前的JavaScript中对象字段是对象属性，是一个键值对，而在ECMAScript5中引入property，property有几个特征

1. `value`: 值，默认是`undefined`

2. `writable`: 是否是只读property，默认是false,有点像C++中的`const`

3. `enumerable`: 是否可以被枚举(如`for...in...`语句)，默认`false`

4. `configurable`: 是否可以被删除，默认false. 同样可以像C#、Java一样些get/set，不过这两个不能和value、writable属性同时使用

5. `get`:返回property的值得方法，默认是`undefined`

6. `set`: 为property设置值的方法，默认是`undefined`

```js
// 为对象o创建age属性, 其值为24, 可写可遍历可删除
Object.defineProperty(o,'age', {
    value: 24,
    writable: true,
    enumerable: true,
    configurable: true
});

Object.defineProperty(o, 'sex', {
    value: 'male',
    writable: false,
    enumerable: false,
    configurable: false
});

console.log(o.age); //24
o.age = 25;

for (var obj in o) {
    console.log(obj + ' : ' + o[obj]);
    /*
    age : 25  //没有把sex ： male 遍历出来
    say : function () {
        alert(this.name);
    } 
    name : Byron 
    */
}
delete o.age;
console.log(o.age);//undefined 

console.log(o.sex); //male
//o.sex = 'female'; //Cannot assign to read only property 'sex' of #<Object> 
delete o.age; 
console.log(o.sex); //male ,并没有被删除
```

也可以使用defineProperties方法同时定义多个property

```js
Object.defineProperties(o, {
    'age': {
        value: 24,
        writable: true,
        enumerable: true,
        configurable: true
    },
    'sex': {
        value: 'male',
        writable: false,
        enumerable: false,
        configurable: false
    }
});
```

## 3. Object.getOwnPropertyDescriptor(O,property)

这个方法用于获取defineProperty方法设置的property属性对象

```js
var props = Object.getOwnPropertyDescriptor(o, 'age');
console.log(props); //Object {value: 24, writable: true, enumerable: true, configurable: true}
```

## 4. Object.getOwnPropertyNames

获取所有的属性名，不包括`prototype`中的属性，返回一个数组

```js
console.log(Object.getOwnPropertyNames(o)); //["age", "sex"]
```

例子中可以看到prototype中的name属性没有获取到(本文第一节设置)

## 5. Object.keys()

和`getOwnPropertyNames`方法类似，但是获取所有的可枚举的属性，返回一个数组

console.log(Object.keys(o)); //["age"]

上面例子可以看出不可枚举的`sex`没有获取到

## 6. Object.preventExtensions(O) / Object.isExtensible

`Object.preventExtensions`方法用于锁住对象，使其不能够拓展，也就是不能增加新的属性，但是属性的值仍然可以更改，也可以把属性删除，`Object.isExtensible`用于判断对象是否可以被拓展

```js
console.log(Object.isExtensible(o)); //true
o.lastName = 'Sun';
console.log(o.lastName); //Sun, 此时对象可以拓展

Object.preventExtensions(o);
console.log(Object.isExtensible(o)); //false

o.lastName = "ByronSun";
console.log(o.lastName); //ByronSun，属性值仍然可以修改

//delete o.lastName;
console.log(o.lastName); //undefined仍可删除属性

o.firstname = 'Byron'; //Can't add property firstname, object is not extensible 不能够添加属性
```

## 7. Object.seal(O) / Object.isSealed

方法用于把对象密封，也就是让对象既不可以拓展也不可以删除属性（把每个属性的configurable设为false）,单数属性值仍然可以修改，Object.isSealed由于判断对象是否被密封

```js
Object.seal(o);
o.age = 25; //仍然可以修改
delete o.age; //Cannot delete property 'age' of #<Object>
```

## 8. Object.freeze(O) / Object.isFrozen

终极神器，完全冻结对象，在seal的基础上，属性值也不可以修改（每个属性的wirtable也被设为false）

```js
Object.freeze(o);
o.age = 25; //Cannot assign to read only property 'age' of #<Object>
```

------

最后

上面的代码都是在Chrome 29下一严格模式（`use strict`）运行的，而且提到的方法都是Object的静态函数，也就是在使用的时候应该是`Object.xxx(x)`，而不能以对象实例来调用。总体来说ES5添加的这些方法为javaScript面向对象设计提供了进一步的可配置性，用起来感觉很不错。