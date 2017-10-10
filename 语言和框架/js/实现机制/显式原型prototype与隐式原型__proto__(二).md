# 显式原型`prototype`与隐式原型`__proto__`(二)

参考文章

1. [理解js中的原型链，prototype与__proto__的关系](http://rockyuse.iteye.com/blog/1426510)

## `__proto__`由来

JavaScript中任意对象都有一个内置属性`[[prototype]]`, 但在ES5之前没有标准的方法访问这个内置属性, 但是大多数浏览器都支持通过`__proto__`来访问。ES5中有了对于这个内置属性标准的Get方法`Object.getPrototypeOf()`与`Object.setPrototypeOf()`, 用于替代直接对`__proto__`进行赋值的操作. 

`Object.setPrototypeOf(subClass, superClass)`与`subClass.__proto__ = superClass`等价.

设置`__proto__`与设置`prototype`的效果相同, 两者都是原型链的实现方式. 如下

```
function A(){}
> undefined
A.color = 'red';
> "red"
function B(){}
> undefined
Object.setPrototypeOf(B, A);
> function B(){}
B.color
> "red"
```

猜测, ES6中通过设置A与B间的`__proto__`指向来实现类的静态成员属性.

## 原型链的本质

> prototype只是一个假象, 它在实现原型链中只是起到了一个辅助作用, 换句话说, 它只是在new的时候有着一定的价值, 而原型链的本质, 其实在于`__proto__`!

关于这句话的解释, 参见《Javascript高级程序设计》第三版, 图6-5.

其中可以看出, `[[prototype]]`, 也就是`__proto__`, 其实是一个指针, 而`prototype`是一个对象. 这张图与数据结构中的链表十分相似. 

图中, Object, SuperType与SubType中只有`prototype`, 没有画出`__proto__`, 我猜想是因为作为函数类型对象, 本来就可以直接拥有`prototype`对象, 不必再通过`__proto__`来指向它了. 

图中没画出的是, `Object.prototype`的`__proto__`值, 其实是`null`. 我觉得, 这就是js原型链的源头.