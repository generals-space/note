# NodeJS使用笔记

## 1. NodeJS升级

无需编译新版源码包, 也无需重新apt/yum覆盖安装. node本身有一个`n`模块(名字真够短的), 用于管理NodeJS版本.

安装方法

```shell
sudo npm install -g n
```

之后系统中会出现`n`命令, 不过没有man手册, 使用`n --help`可以查看使用帮助, 用法很明了.

## 2. node服务自动重启

[参考文档](http://blog.csdn.net/haigenwong/article/details/48732267)

开发node应用时, 一但应用已经启动了, 这个时候如果你修改了服务端的文件, 那么要使这个修改起作用, 必须手动停止服务然后再重新启动. 这在开发过程中无疑是很烦人的一件事, 最好是有一个能够监控所有变动文件的脚本, 一单发现文件有变动则立即重启服务, 重新加载刚刚修改过的文件. 这里推荐一个: nodemon.

首先为了是这个命令全局可用, 最好我们进行全局安装:

```shell
npm install -g nodemon
```

然后进入项目根目录:  

```shell
nodemon  server.js
```

这样就可以启动应用了, 并且在文件有变化之后会自动重启服务.
