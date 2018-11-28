# nodejs获取当前文件路径

```
$ pwd
/root/test
$ tree 
.
├── main.js
└── subpath
    └── submain.js

1 directory, 2 files
```

`main.js`

```js
var path = require('path');
var sub = require('./subpath/submain.js');
console.log(path.resolve('../'));
console.log(__dirname);
sub();
```

`subpath/sumain.js`

```js
var path = require('path');

module.exports = () => {
    console.log(path.resolve('../'));
    console.log(__dirname);
}
```

执行`node main.js`有如下输出

```
$ node main.js 
/root
/root/test
/root
/root/test/subpath
```

而在`test`目录的父级目录执行会得到如下结果

```
$ cd ..
$ ls
test
$ node test/main.js 
/
/root/test
/
/root/test/subpath
```

也就是说, 单纯使用`../`这种相对路径解析到的绝对路径, 是和执行脚本时所在的目录有关的, 而且只表示入口脚本的路径, 子模块中的路径不会是你想要的.

而`__dirname`这个内置变量, 才是最符合我们期望的, 就表示当前文件所在路径, 与其他任何变量都不相关. 而且还不需要使用`resolve()`转换.