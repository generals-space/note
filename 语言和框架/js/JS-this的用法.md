# JS-this的用法

参考文章

1. [JS中this的四种用法](http://www.cnblogs.com/pabitel/p/5922511.html)

## 1. 普通方法中, this指代全局变量

```js
var a = 123;
function test(){
	console.log(this.a);
	this.a = 'abc';
	console.log(this.a);
}
test();
```

输出如下 

```
123
abc
```

当然, 这种情况下, 其实不用this也能引用并修改`a`的值的.

## 2. 在对象的成员方法使用, this指代上级对象

这一节举的3个例子, 不要萌B, 你会发现它们说的是同一个意思.

```js
var obj = {
	setName: function(name){
		this.name = name;
	},
	getName: function(){
		console.log(this.name);
	}
};
obj.setName('general');
obj.getName();
console.log(this.name);
```

输出如下

```
general     ## 这里的name与window的作用域不同.
            ## 空行表示最后一行的`console.log()`输出, 因为window对象的name成员变量并未设置过.
```

> 注意: 由于对象变量的成员一般都是键值对, 所以在普通成员变量中使用`this`是不合理的, 即this只能在成员**方法**中使用. 比如, 这句就不对`var obj = {this.a = 123};`.

### 2.1 嵌套对象的成员方法中this的指向

```js
var obj = {
	setName: function(name){
		this.name = name;
	},
	getName: function(){
		console.log(this.name);
	},
	subobj: {
		setName: function(name){
			this.name = name;
		},
		getName: function(){
			console.log(this.name);
		}
	}
};
obj.setName('general');
obj.getName();
obj.subobj.setName('hehe');
obj.subobj.getName();
obj.getName();
```

上述代码的输出为

```
general
hehe
general
```

ok, 可以看出`obj`与`obj.subobj`中引用的`this`是互不影响的, 倒是很符合认知.

### 2.2 函数对象的属性方法中this的指向

在第1节我们已经知道在普通函数体内部定义的`this`取到的是`window`, 但是下面这种就有点区别.

```js
function A(){};
A.extend = function(){
	console.log(this);				// ƒ A(){}
	console.log(typeof this);		// function
}

A.extend();
```

原因在于, 这种情况下, A是当作普通`object`类型对象来看待的, 而不是`function`类型.

直接定义在`A`本身的属性, 被称为类属性, 或者静态属性...其实就是普通的对象属性. 所以要匹配的还是我们本节(也就是第2节)的场景.

## 3. 在构造函数中使用，this指代new出的实例对象

```js
function Person(){
    // 赋初始值
	this.name = 'general';
	this.getName = function(){
		console.log(this.name);
	};
	this.setName = function(name){
		this.name = name;
	}
};
person = new Person();
person.getName();
person.setName('hehe');
person.getName();
```

得到如下输出

```
general
hehe
```

## 4. 通过apply()方法调用

`apply()`方法的作用是改变当前方法的主调对象, 也即this对象. (所以貌似应该只对使用了this的方法有效???).

```js
function Person(){
    // 赋初始值
	this.name = 'general';
	this.getName = function(){
		console.log(this.name);
	};
	this.setName = function(name){
		this.name = name;
	}
};
person = new Person();
person.getName();
person.setName('hehe');
person.getName();
window.name = 'win10';
// 参数为空时, 默认主调者为全局对象, 在浏览器中就是window对象.
person.getName.apply();
person.getName.apply(person);
```

得到如下输出.

```
general
hehe
win10
hehe
```