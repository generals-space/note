# nodejs-path路径解析

不只是nodejs的path库, python的pathlib, golang的os/path, 功能都差不多, 只是路径字符的解析操作. 

path不能判断(文件/目录)路径是否存在, 不能判断是否拥有读写权限, 更不能完成创建/删除的操作. 

path常用的操作差不多就是如下几种

预设变量`targetPath = 'a/b/c';`

## 1. 

一种是basename, dirname, 与linux系统中的同名命令作用相同.

```js
path.basename(targetPath); // c
path.dirname(targetPath); // a/b
path.extname('d.txt') // .txt
```

## 2. 

然后是parse和format, 作用类似于golang的time库中parse和format方法, `format()`是`parse()`的逆操作.

```js
path.parse('/home/user/dir/file.txt');
// 返回:
// { root: '/',
//   dir: '/home/user/dir',
//   base: 'file.txt',
//   ext: '.txt',
//   name: 'file' }
```

## 3.

然后是路径拼接, resolve和join, 这个操作还是比较常用的.

```js
path.resolve(targetPath, 'd.txt'); // /Users/general/Code/playground/a/b/c/d.txt
path.join(targetPath, 'd.txt'); // a/b/c/d.txt
```

上面可以看出, resolve返回的结果是绝对路径, 而join的结果是相对路径.

但ta们两者的区别可不只于此.

```js
path.resolve('/a/b', '/c', 'd.txt'); // /c/d.txt
path.join('/a/b', '/c', 'd.txt'); // /a/b/c/d.txt
```

在做路径拼接时, join对于参数中的根路径成员并不能得到正确的结果, 而resolve则可以.

在nodejs中, url库没有路径拼接的操作(只有parse和format, 还有一个resolve但是只能接受两个参数), 所以多个路径片段拼接时可能需要使用path库, 到时重新赋值Url对象中的path成员, 再调用toString()就可以得到结果了.
