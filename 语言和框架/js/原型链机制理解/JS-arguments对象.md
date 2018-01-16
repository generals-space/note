# JS-arguments对象

参考文章

1. [将arguments转换成数组的方法](http://www.cnblogs.com/AliceX-J/p/5400568.html)

## 1. arguments这个类数组

arguments是个类数组，除了有实参所组成的类似数组以外，还有自己的属性，如`callee`，`arguments.callee`就是当前正在执行的这个函数的引用，它只在函数执行时才存在。因为在函数开始执行时，才会自动创建第一个变量arguments对象。

1. 将实参以数组的形式保存着,还可以像数组一样访问实参,如arguments[0];

2. 也有自己独特的属性，如：callee，

3. `arguments`的长度(`arguments.length`属性)是**实参的个数**。补充：那`arguments.callee.length`又是什么呢？`arguments.callee`是当前正在执行的函数(主调函数)的引用，类似`function.length`，那就是**形参的个数**。

## 将arguments转换为真正的数组的方法

`apply`方法需要的参数列表为数组类型, 有时候的确需要把`arguments`转换成原生数组, 这里提供几个思路.

1. `Array.prototype.slice.apply(arguments)`: 这是运行效率比较快的方法（看别人资料说的）,为什么不是数组也可以，因为arguments对象有length属性，而这个方法会根据length属性,返回一个具有length长度的数组。若length属性不为`number`，则数组长度返回0; 所以其他对象只要有length属性也是可以的哟，如对象中有属性0, 对应的就是arr[0],即属性为自然数的number就是对应的数组的下标，若该值大于长度，当然要割舍啦。

2. `Array.prototype.concat.apply(thisArg,arguments)`: thisArg是新的空数组，apply方法将函数this指向thisArg，arguments做为类数组传参给apply。根据apply的方法的作用,即将Array.prototype.slice方法在指定的this为thisArg内调用，并将参数传给它。用此方法注意:若数组内有数组，会被拼接成一个数组。原因是apply传参的特性。

3. 我自己想了个方法，利用Array的构造函数,如`Array(1,2,3,4,5,6)`; 可以返回一个传入的参数的数组，那Array.apply(thisArg,arguments)也可以将arguments转化为数组，果然实验是可以的; 有没有什么影响呢，就是乱用了构造函数，但这也是js的特性嘛。构造函数也是函数。用此方法注意: 若数组内有数组，会被拼接成一个数组。原因是apply传参的特性。

4. 用循环，因为arguments类似数组可以使用arguments[0]来访问实参，那么将每项赋值给新的数组每项，直接复制比push要快，若实参有函数或者对象，就要深拷贝。