# nodejs-模块导出exports与module.exports的区别

参考文章

1. [对module.exports和exports的一些理解](https://www.cnblogs.com/wbxjiayou/p/5767632.html)

## 1. exports

`src.js`文件

```js
var User = {
    name: 'general',
    age: 21
};
exports.user = User;
```

另外的文件在`require`这个文件后, `User`对象的使用方法如下.

```js
var src = require('./src');
var User = src.user;    // 注意这里, src后要跟一个`user`属性才能取到src文件中的User对象.
```

## 2. module.exports

同样的场景

```js
var User = {
    name: 'general',
    age: 21
};
module.exports = User;
```

这种导出方式下, 其他文件引用`User`对象的方法如下.

```js
var User = require('./src');    // 这里就可以直接使用了.
```