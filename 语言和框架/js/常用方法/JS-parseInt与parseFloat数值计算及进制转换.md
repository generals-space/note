# JS-parseInt与parseFloat数值计算及进制转换

## 1. 从字符串提取整数/浮点数

```js
//字符串转换为数字
parseInt('12')
12
parseFloat('3.245')
3.245

//可以无视空格和其他字符干扰, 但只能匹配到第1个合法的整数/浮点数
parseInt(' 34 e')
34
parseInt('34 4')
34

//另外, 如果是其他字符的话, 只有在合法数值的后面出现才可以, 如果出现在数值前面, 会返回NaN
parseFloat('12.34abc')
12.34
parseFloat('a 12.34')
NaN
```

## 2. 使用parseInt()进行进制转换

`parseInt()`的输入参数可以是任意进制的字符串, 返回值是十进制, 即可以将任意进制的数字转换为十进制.

```js
parseInt('11000000', 2)
192
parseInt('FF', 16)
255
//纯数字的话, 参数格式可以直接是数字, 但超过十进制需要有符号代替时, 需要加上引号才行, 否则会报错
parseInt(11000000, 2)
192
parseInt(ff, 16)
Uncaught ReferenceError: ff is not defined
```

------

上面的方法是已知某进制下的具体数值, 要转换成十进制的. 但大多数情况下我们有的是一个十进制数值, 要转换成目标进制. 这种情况下要使用数值类型的`toString()`方法...多年来第一次知道还可以这么用.

```js
var a = 45;
a.toString(2); // '101101'
a.toString(16); // '2d'
```

得到的结果为字符串类型.

## 3. 使用parseFloat()进行科学计数法格式转换

`parseFloat()`输入科学计数法的格式, 返回值为正常格式. (这个功能貌似有点鸡肋, 程序中几乎没见过, 而实际使用时自己也能数0的个数嘛).

```js
parseFloat('314e-2')
3.14
parseFloat('0.0314E+2')
3.14
//...貌似经过parseFloat计算后的结果不能超过float的表示极限
parseFloat("10E+19")
100000000000000000000
parseFloat("10E+20")
1e+21
parseFloat("100E+18")
100000000000000000000
parseFloat("100E+19")
1e+21
```
