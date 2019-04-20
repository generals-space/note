# Windows下的npm设置

参考文章

1. [设置Nodejs NPM全局路径](http://blog.csdn.net/carechere/article/details/51279789)

2. [nodejs 中的 NODE_PATH](http://www.cnblogs.com/zxktxj/p/5051831.html)

我们想设置的, 无非是两点: 全局包路径, 源.

全局包路径默认在`C:\Users\general\AppData\Roaming\npm\node_modules`, `npm install -g xxx`时可以验证.

好像在3个配置文件

1. `C:\Users\general\AppData\Roaming\npm\etc\npmrc` 不过好像没用.

2. `C:\Users\.npmrc` 可用, 但不想用

3. node安装路径下的`node_modules\npm\npmrc` 也可用

在其中定义`registry`, `cache`字段, 不难理解.

另外一个是`prefix`, 它的值加上`node_modules`得到的路径就是全局包的安装路径.

> cmd中用`npm config edit/set`命令操作的就是第2个配置文件...没什么意思

但是, 安装了却无法调用(如会找不到`cnpm`命令), 也没办法在node文件或命令行中用`require`导入. 这是因为除了`node.exe`所在的路径要添加到PATH变量为, 还需要再添加一个`NODE_PATH`变量, 其值为npm的全局包路径, 即`prefix`值 + `node_modules`. 然后问题就都解决了.
