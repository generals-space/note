# java-语法错误Local variable xxx defined in an enclosing scope must be final or effectively final

参考文章

1. [Local variable flag defined in an enclosing scope must be final or effectively final](https://blog.csdn.net/weixin_38883338/article/details/89195749)

```java
boolean flag = true;
if(xxx){
    flag = false;
}

list.stream().forEach(e -> {
    boolean newFlag = false;
	if (flag) {
		newFlag = flag; // 这里 flag 会出现红色波浪线
	}
});
```

参考文章1中的解释说明和处理方法写得都不错, 就是有点麻烦.

原因在于外部的`flag`在定义时拥有一个初始值`true`, 之后又在if条件中进行修改, 这样在`forEach`函数中被认为该值不再是"final"了(没懂是什么意思), 所以报错.

出现这个问题的条件有3个吧

1. 变量有初始值
2. 变量值被修改
3. 内部类/内部函数中引用了该变量

解决方法其实不用参考文章1说的那么麻烦, 声明`flag`变量时不再赋初始值即可.

```java
boolean flag;
if(xxx){
    flag = false;
} else {
    flag = false;
}
```

这样, 内部函数中就不会再报错了.
