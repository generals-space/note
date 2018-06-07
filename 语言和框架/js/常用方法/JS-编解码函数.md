# JS-编解码函数

参考文章

1. [js 中编码（encode）和解码（decode）的三种方法](http://blog.csdn.net/xingkong22star/article/details/39155739)

2. [javascript之url转义escape()、encodeURI()和decodeURI()](http://www.cnblogs.com/kissdodog/archive/2012/12/22/2829489.html)

3. [《Web前端黑客技术揭秘》6.7.2节]()

JS有3套编/解码函数, 分别为

- escape/unescape

- encodeURI/decodeURI

- encodeURIComponent/decodeURIComponent

分别使用3种函数对字符串`<Hello+World>`进入编码, 结果如下

|        编码函数        |          编码结果         |
|:------------------:|:---------------------:|
|       escape       |  `%3CHello+World%3E`  |
|      encodeURI     |  `%3CHello+World%3E`  |
| encodeURIComponent | `%3CHello%2BWorld%3E` |

我们可以看到3种编码方式近乎相同, 但还是有些许差别.

|        编码函数        |                           不编码字符范围                           |
|:------------------:|:-----------------------------------------------------------:|
|       escape       |            `a-zA-Z0-9`, `+-*/`, `.`, `@`, `_`共69个           |
|      encodeURI     | `a-zA-Z0-9`, `+-*/`, `()`, `~!@#$&*`, `,.'?`, `:;`, `=`共82个 |
| encodeURIComponent |   `a-zA-Z0-9`, `()`, `!`, `'`, `*`, `-`, `.`, `_`, `~`共71个  |

可以看到, 从编码范围上来说, `escape` > `encodeURIComponent` > `encodeURI`, 这其实与3者的适用范围有关. 

`escape`可以编码web应用中大部分不常见的字符, 包括中文, 使用`%uxxxx`双字节表示(...诶? 双字节还是4字节??). 它是一种程序级的编码方案.

`encodeURIComponent`一般用以编码url中的参数部分, 我们知道这部分的典型格式为`?a=val1&b=val2`. 没错, 它就是用来编码`val1`和`val2`部分的. OK, 为什么? 

如果url中有参数是一个合法的url地址? `?src=https://www.baidu.com?s=java&date=2017`, 其中`&date=2017`很可能将`src`参数截断从而认为`date`是与`src`同级的另一个参数. 为了防止这种情况, 我们需要`encodeURIComponent`.

至于`encodeURI`, 它要考虑的情况就更多了, 从名字上来看就知道, 它是应用于编码**整个url**的手段. 所以它不能对一个合法的url产生不好的影响, 比如`http\://`中的冒号和斜线, 邮件地址的`@`符号, 后缀分隔符`.`点号等, 都不会作处理.