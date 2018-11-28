# npm install时报EACCES: permission denied错误

参考文章

1. [npm ERR! Error: EACCES: permission denied, access](https://blog.csdn.net/testcs_dn/article/details/78869419)

`npm install`安装package.json文件中的依赖时报这个错误, 当时屏幕闪得太快没看到, 最下方提示查看`~/.npm`目录下的日志, 没得到什么有用的信息. 胡乱找了好久才想起来往上翻.

解决方法如下.

在命令前加上 `sudo`

`sudo npm install`

不过这样子可能还是不行, 你需要这样

`sudo npm install --unsafe-perm=true`

或许你还是会遇到错误, 请尝试这样

`sudo npm install --unsafe-perm=true --allow-root`