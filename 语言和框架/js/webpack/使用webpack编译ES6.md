# 使用webpack编译ES6

参考文章

1. [【webpack】使用webpack进行ES6开发](http://blog.csdn.net/xiaozhuxmen/article/details/51597923)

2. [一小时包教会 —— webpack 入门指南](http://www.cnblogs.com/vajoy/p/4650467.html)

## 1. 环境依赖

webpack本身需要node环境, 预编译es6时需要`babel`插件(在babel中插件被称为`loader`).

```
$ cnpm instal webpack -g
$ cnpm instal webpack --save
$ cnpm install babel-loader babel-core babel-es2015 --save
```

## 2. webpack配置

工程目录结构

```
$ pwd
/root/es6
$ tree -I node_modules
.
├── index.js
├── package.json
└── webpack.config.js

0 directories, 4 files
```

`webpack.config.js`是webpack默认使用的配置文件, 内容如下

```js
var path = require('path');
module.exports = {
    entry: './index.js', 	// 入口文件
    output: {		//打包输出文件(包含路径配置)
    	path: __dirname,
        filename: 'bundle.js'
    },
    module: {
    	loaders: [
	    {
	    	test: /\.js$/,
		loader: 'babel-loader',
                //不打包exclude下的文件
		exclude: /node_modules/,
		query: {
		    presets: ['es2015'],
		},
	    },
	]
    },
    resolve: {
    	// 自动扩展文件后缀名, 即可以直接使用require('file')而不再写require('file.js')
        // 虽然有些教程中这个字段中第一个字段可以是一个空字符串'', 但我在实验时这样会报错.
    	extensions: ['.js'],
    }
}
```

然后同目录下`index.js`文件的内容为

```js
var a = 2;
{
    let a = 3;
    console.log(a);
}
console.log(a);
```

执行`webpack`, 会生成打包好的`bundle.js`.

查看`bundle.js`, 其中的关键代码为

```js
"use strict";


var a = 2;
{
    var _a = 3;
    console.log(_a);
}
console.log(a);
```

ok, 可以看到局部作用域与全局作用域内变量`a`不冲突的原因是...变量名变了...good.