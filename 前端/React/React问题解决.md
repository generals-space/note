
参考文章

1. [React路上遇到的那些问题以及解决方案](http://blog.csdn.net/liangklfang/article/details/53694994)

## 1. 

```
Warning: React.createElement: type is invalid -- expected a string (for built-in components) or a class/function (for composite components) but got: undefined. 
You likely forgot to export your component from the file it's defined in. Check the render method of 'App'.
```

浏览器访问页面时报上述错.

原因分析: 目标模块或其子模块export接口错误. 实际上, 是我的一个子模块没有在结尾调用`export`语句.

## 2. 

```
TypeError: Super expression must either be null or a function, not undefined
```

解决：说明你extend的那个函数没有导出相应的属性(是我的React.Component单词拼写错误...)